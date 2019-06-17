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
	//"sync"

	mle_util "github.com/mle/runtime/util"
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
 * Sets the name of this phase.
 *
 * @param name A <code>java.lang.String</code> identifying the name
 * of this phase.
 */
func (p *MlePhase) SetName(name string) {
	p.m_name = name
}

/**
 * Gets the name of this phase.
 *
 * @return A <code>java.lang.String</code> identifying the name
 * of this phase.
 */
func (p *MlePhase) GetName() string {
	return p.m_name
}
	 
/**
 * Gets the number of registered tasks for this phase.
 *
 * @return An integer representing the number of tasks belonging
 * to this phase.
 */
func (p *MlePhase) GetNumberOfTasks() int	 {
	return len(*p.m_tasks)
}


/**
 * Adds a task to this phase.
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
 * Removes a task from this phase.
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
	//f := &foo{}
	p.m_tasks.RemoveElement(task)
	return true
}
 
// String implements the IObject interface.
func (p *MlePhase) String() string {
    return p.m_name
}
