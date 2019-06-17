/*
 * @file MleScheduler.go
 * Created on June 17, 2019. (msm@wizzerworks.com)
 */

// COPYRIGHT_BEGIN
//
// The MIT License (MIT)
//
// Copyright (c) 2019 Wizzer Works
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//
//  For information concerning this source file, contact Mark S. Millard,
//  of Wizzer Works at msm@wizzerworks.com.
//
//  More information concerning Wizzer Works may be found at
//
//      http://www.wizzerworks.com
//
// COPYRIGHT_END

// Declare package.
package scheduler

// Import go packages.
import (
	"fmt"
	"strconv"
	"bytes"
	mle_util "github.com/mle/runtime/util"
)

/**
 * The <code>MleScheduler</code> class is used to schedule phases of execution
 * which might be required by a runtime engine. For example, a 3D game engine
 * might require four phases of execution:
 * <ol>
 * <li> a behavioral pass on a scene graph to update an actor's behavior
 * <li> a pre-render pass on a scene graph to calculate transformations
 * <li> a render pass on a scene graph to render a scene to an offscreen buffer
 * <li> a post-render pass to blit the offscreen buffer to a screen
 * </ol>
 * <p>
 * Each phase of the scheduler must be registered by calling
 * <code>addPhase</code> with an instance of the <code>MlePhase</code> class.
 * The <code>MlePhase</code> class is a collection of tasks which must be
 * executed during that phase.
 * The tasks are actually <code>java.lang.Thread</code> objects encapsulated by
 * the <code>MleTask</code> class.
 * </p>
 *
 * @see MlePhase
 * @see MleTask
 */
type MleScheduler struct {
    // The number of phases.
    m_phases *mle_util.Vector
    // Flag indicating that it is ok to exit.
    m_exitOK bool
}

// NewMleScheduler is the default constructor.
func NewMleScheduler() *MleScheduler {
	p := new(MleScheduler)
	p.m_phases = mle_util.NewVector()
	p.m_exitOK = false
	return p
}

/**
 * Flags scheduler that it is Ok to exit.
 */
func (s *MleScheduler) SetExitOk() {
	s.m_exitOK = true
}

/**
 * Gets the number of registered phases for this scheduler.
 *
 * @return An integer representing the number of phases which have been
 * registered with this scheduler.
 */
func (s *MleScheduler) GetNumberOfPhases() int {
	return len(*s.m_phases)
}

/**
 * Adds a phase to this scheduler.
 *
 * @param phase An instance of <code>MlePhase</code>.
 *
 * @return A boolean value is returned indicating whether the phase was
 * successfully added. If the phase was successfully added, then
 * <b>true</b> is returned; otherwise, <b>false</b> will be returned.
 *
 * @see MlePhase
 */
func (s *MleScheduler) AddPhase(phase *MlePhase) bool {
	s.m_phases.AddElement(phase)
	return true
}

/**
 * Removes a phase from this scheduler.
 *
 * @param phase An instance of <code>MlePhase</code>.
 *
 * @return A boolean value is returned indicating whether the phase was
 * successfully deleted. If the phase was successfully deleted, then
 * <b>true</b> is returned; otherwise, <b>false</b> will be returned.
 *
 * @see MlePhase
 */
func (s *MleScheduler) DeletePhase(phase *MlePhase) bool {
	s.m_phases.RemoveElement(phase)
	return true
}

/**
 * Gets the phase at index <code>n</code>.
 *
 * @param n An integer identifying the nth phase in this scheduler.
 *
 * @return If a phase is found at the specifed location, <code>n</code>,
 * then an object of type <code>MlePhase</code> will be returned.
 * Otherwise, <b>null</b> will be returned.
 *
 * @see MlePhase
 */
func (s *MleScheduler) GetPhase(n int) *MlePhase {
	var phase *MlePhase

	e := s.m_phases.ElementAt(n)
	if e != nil {
		phase = e.(*MlePhase)
	} else {
		phase = nil
	}
	return phase
}

/**
 * Gets the phase with the specified <code>name</code>
 *
 * @param name The name of the phase to retireve from this
 * scheduler. It is an instance of <code>java.lang.String</code>
 *
 * @return If a phase is found with the specifed <code>name</code>,
 * then an object of type <code>MlePhase</code> will be returned.
 * Otherwise, <b>null</b> will be returned.
 *
 * @see MlePhase
 */
func (s *MleScheduler) GetPhaseWithName(name string) *MlePhase {
	var phase *MlePhase

	for i := 0; i < len(*s.m_phases); i++ {
		curPhase := s.m_phases.ElementAt(i).(*MlePhase)
		curName := curPhase.GetName()
		if name == curName {
			/** Names are equal */
			phase = curPhase
			break
		}
	}
	return phase
}

/**
 * Adds a task to the specified phase. The phase must have been previously
 * registered with this scheduler.
 *
 * @param phase The phase which the <code>task</code> will be added to.
 * @param task The task which will be added to the specifed
 * <code>phase</code>.
 *
 * @return A boolean value is returned indicating whether the
 * <code>task</code> was successfully added to the specifed
 * <code>phase</code>. If the task was successfully added, then
 * <b>true</b> is returned; otherwise, <b>false</b> will be returned.
 *
 * @see MlePhase
 * @see MleTask
 */
func (s *MleScheduler) AddTask(phase *MlePhase, task *MleTask) bool {
	if s.m_phases.Contains(phase) {
		return phase.AddTask(task)
	}
	return false
}

/**
 * Removes the a task from the specified phase. The phase must be
 * registered with this scheduler.
 *
 * @param phase The phase which the <code>task</code> will be removed from.
 * @param task The task which will be removed from the specifed
 * <code>phase</code>.
 *
 * @return A boolean value is returned indicating whether the
 * <code>task</code> was successfully deleted from the specifed
 * <code>phase</code>. If the task was successfully deleted, then
 * <b>true</b> is returned; otherwise, <b>false</b> will be returned.
 *
 * @see MlePhase
 * @see MleTask
 */
func (s *MleScheduler) DeleteTask(phase *MlePhase, task *MleTask) bool {
	if s.m_phases.Contains(phase) {
		return phase.DeleteTask(task)
	}
	return false
}

/**
 * Executes scheduled phases. The order of execution is based on the order
 * in which the phases were registered with this scheduler. A phase must
 * complete before the next phase is executed. A phase is complete when all
 * tasks registered with that phase have completed and are no longer
 * running.
 * <p>
 * When all phases have been processed, the execution continues again with
 * the first phase. Thus the phases are executed using a round-robin
 * algorithm.
 * </p><p>
 * To discontinue execution, the scheduler must be flagged by specifying
 * that it is Ok to exit. This is done using the <code>setExitOk</code>
 * method. The scheduler will finish executing the current phase and exit
 * prior to starting the next phase.
 * </p><p>
 * No state is maintained between phases. Therefore an application which
 * has discontinued execution by setting the exit Ok flag should not expect
 * the scheduler to start up again at the next phase if this method is
 * invoked again. <code>run</code> will always invalidate the exit Ok flag
 * and start at the first scheduled phase.
 */
func (s *MleScheduler) Run(done chan bool) {
	s. m_exitOK = false
	for i := 0; i < len(*s.m_phases); i++ {
		/* Fork off tasks in task list scheduled for this phase. */ 
		phase := s.m_phases.ElementAt(i).(* MlePhase)
 
		/* Wait for phase[i] to complete. */
		phase.Run(nil)
		if s.m_exitOK {
			break
		}
	}
}

/**
 * Run the specified phase.
 * 
 * @param phase The phase to execute.
 */
func (s *MleScheduler) Go(phase *MlePhase) {
	phase.Run(nil)
}

/**
 * Dumps the list of registered phases for this scheduler. It will also
 * list the tasks associated with each phase.
 */
func (s *MleScheduler) Dump() {
	var buf bytes.Buffer

	for i := 0; i < s.GetNumberOfPhases(); i++ {
		phase := s.GetPhase(i)

		buf.WriteString("Phase ")
		buf.WriteString(strconv.Itoa(i+1))
		buf.WriteString(": ")
		buf.WriteString(phase.GetName())
			 
		for j := 0; j < phase.GetNumberOfTasks(); j++ {
			task := phase.GetTask(j)

			buf.WriteString("\tTask ")
			buf.WriteString(strconv.Itoa(j+1))
			buf.WriteString(": ")

			if (task.GetName() != "") {
				buf.WriteString(task.GetName())
			} else {
				buf.WriteString("empty")
			}
		}
	}

	fmt.Println(buf.String())
}
