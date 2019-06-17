/*
 * @file MlePhase.go
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
	"bytes"

	mle_util "github.com/mle/runtime/util"
	mle_core "github.com/mle/runtime/core"
)

/**
 * The <code>MlePhase</code> class is used to manage a set of tasks. Each task
 * is represented by an instance of the <code>MleTask</code> class.
 * For convenience the <code>MlePhase</code> may be identified with a
 * <code>String</code>; however, by default the <code>MlePhase</code> is
 * unnamed.
 * <p>
 * This class is used to register a task or group of tasks which must be
 * executed in context of a scheduled phase. In addition to providing methods
 * for adding and deleting a task, the <code>MlePhase</code> class provides a
 * mechanism for running all of the associated tasks.
 * When the tasks are executed, they will be invoked asynchronously.
 * However, <code>MlePhase</code> will not return until all of the tasks have
 * been completed.
 * </p>
 *
 * @see MleTask
 * @see MleScheduler
 */
type MlePhase struct {
	// The list of tasks.
	m_tasks *mle_util.Vector
	// The name of the phase.
	m_name string
}

/**
 * Creates a new MlePhase.
 */
func NewMlePhase() *MlePhase {
	p := new(MlePhase)
	p.m_name = ""
	p.m_tasks = mle_util.NewVector()
	return p
}

/**
 * Creates a new MlePhase with the specified <i>name</i>.
 *
 * @param name A <code>java.lang.String</code> identifying the name
 * of the phase.
 */
func NewMlePhaseWithName(name string) *MlePhase {
	p := new(MlePhase)
	p.m_name = name
	p.m_tasks = mle_util.NewVector()
	return p
}

/**
 * Set the name of this phase.
 *
 * @param name A <code>java.lang.String</code> identifying the name
 * of this phase.
 */
func (p *MlePhase) SetName(name string) {
	p.m_name = name
}

/**
 * Get the name of this phase.
 *
 * @return A <code>java.lang.String</code> identifying the name
 * of this phase.
 */
func (p *MlePhase) GetName() string {
	return p.m_name
}
	 
/**
 * Get the number of registered tasks for this phase.
 *
 * @return An integer representing the number of tasks belonging
 * to this phase.
 */
func (p *MlePhase) GetNumberOfTasks() int	 {
	return len(*p.m_tasks)
}


/**
 * Add a task to this phase.
 *
 * @param task An instance of the MleTask class.
 *
 * @return A boolean indicating whether the task was successfully added.
 * If the task was added successfully,
 * <b>true</b> is returned; otherwise, <b>false</b> will be returned.
 *
 * @see MleTask
 */
func (p *MlePhase) AddTask(task *MleTask) bool {
	p.m_tasks.AddElement(task)
	return true
}
	 
/**
 * Remove a task from this phase.
 *
 * @param task An instance of the MleTask class.
 *
 * @return A boolean indicating whether the task was successfully deleted.
 * If the task was successfully removed,
 * <b>true</b> is returned; otherwise, <b>false</b> will be returned.
 *
 * @see MleTask 
 */
func (p *MlePhase) DeleteTask(task *MleTask) bool {
	p.m_tasks.RemoveElement(task)
	return true
}

/**
 * Remove the nth task from this phase.
 *
 * @param n An integer identifying the nth task in this phase.
 *
 * @return A boolean indicating whether the task was successfully deleted.
 * If the task was successfully removed,
 * <b>true</b> is returned; otherwise, <b>false</b> will be returned.
 *
 * @see MleTask
 */
func (p *MlePhase) DeleteTaskAt(n int) bool {
	task := p.GetTask(n)
		 
	if (task != nil){
		return p.DeleteTask(task)
	}

	return false
}

/**
 * Remove the task with the specified name from this phase.
 *
 * @param name A <code>java.lang.String</code> identifying the name
 * of the task to be removed.
 *
 * @return A boolean indicating whether the task was successfully deleted.
 * If the task was successfully removed,
 * <b>true</b> is returned; otherwise, <b>false</b> will be returned.
 */
func (p *MlePhase) DeleteTaskWithNAme(name string) bool {
	task := p.GetTaskWithName(name)
		 
	if task != nil {
		return p.DeleteTask(task)
	}
	
	return false
}
 

/**
 * Retrieves the task at index n. The task is not removed from this phase.
 *
 * @param n An integer identifying the nth task in this phase.
 *
 * @return If a task is found at the specifed location, <code>n</code>,
 * then an object of type <code>MleTask</code> will be returned.
 * Otherwise, <b>null</b> will be returned.
 *
 * @see MleTask
 */
func (p *MlePhase) GetTask(n int) *MleTask {
	var task *MleTask
	
	e := p.m_tasks.ElementAt(n)
	if e == nil {
		// Task not found at index n.
		return nil
	} else {
	    task = e.(*MleTask)
	}
	return task
}

/**
 * Retrieves the task with the specified <code>name</code.
 *
 * @param name A <code>java.lang.String</code> identifying the name
 * of the task to be retrieved.
 *
 * @return If a task is found matching the specifed <code>name</code>
 * then an object of type <code>MleTask</code> will be returned.
 * Otherwise, <b>null</b> will be returned.
 *
 * @see MleTask
*/
func (p *MlePhase) GetTaskWithName(name string) *MleTask {
	var task *MleTask

	for i := 0; i < len(*p.m_tasks); i++ {
		curTask := p.m_tasks.ElementAt(i).(*MleTask)
		curName := curTask.GetName()
		if (name == curName) {
			/* Names are equal */
			task = curTask
			break
		}
	}
	return task
}

/**
 * Executes the tasks registered with this phase. <code>Run</code>
 * will not return until all tasks have been completed.
 */
func (p *MlePhase) Run(done chan bool) {
	var buf bytes.Buffer
	buf.WriteString("*** EXECTUING PHASE ")
	buf.WriteString(p.m_name)
	buf.WriteString(" ***")
	mle_core.MleLogInfo(buf.String(), false)
		 
	/* Invoke tasks which have been registered. */
	for i := 0; i < len(*p.m_tasks); i++	{
		task := p.m_tasks.ElementAt(i).(*MleTask)
		task.Invoke()
	}
		 
	/* Wait for all tasks to complete before returning */
	tasksCompleted := false
	waitForTaskCompletion:
	for ! tasksCompleted {
		for i := 0; i < len(*p.m_tasks); i++	{
			task := p.m_tasks.ElementAt(i).(*MleTask)
			if (task.IsRunning()) {
				continue waitForTaskCompletion
			}
		}
	    tasksCompleted = true;
	}
} 
 
// String implements the IObject interface.
func (p *MlePhase) String() string {
    return p.m_name
}
