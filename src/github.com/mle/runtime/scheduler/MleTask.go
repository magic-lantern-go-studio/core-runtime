/*
 * @file MleTask.go
 * Created on June 14, 2019. (msm@wizzerworks.com)
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
	"sync"

	mle_util "github.com/mle/runtime/util"
)

/**
 * <code>MleTask</code> is a class that manages a named thread.
 * <p>
 * MleTask encapsulates an autonamous action, using the
 * <code>channels</code> class, for the Magic Lantern scheduler,
 * <code>MleScheduler</code>.
 * </p>
 *
 * @see MlePhase
 * @see MleScheduler
 */
type MleTask struct {
    // The runnable object.
    m_task mle_util.Runnable
    // A thread that executes the Runnable.
    m_thread *mle_util.Thread
    // The name of the task.
    m_name string
    // A flag indicating whether the task is running.
	m_running bool
	// The task work group.
	m_wg sync.WaitGroup
}

/**
 * Creates a new MleTask with a reference
 * to the specified Runnable, <i>task</i>.
 *
 * @param task An instance of a Runnable object.
 */
func NewMleTask(task mle_util.Runnable) *MleTask {
	p := new(MleTask)
	p.m_task = task
	p.m_running = false
	return p
}

/**
 * Creates a new MleTask with the specified <i>name</i> and a reference
 * to the specified Runnable, <i>task</i>.
 *
 * @param task An instance of a Runnable object.
 * @param name The name of the task; an instance of String.
 */
func NewMleTaskWithName(task mle_util.Runnable, name string) *MleTask {
	p := new(MleTask)
	p.m_task = task
	p.m_name = name
	p.m_running = false
	return p
}

/**
 * Retrieves the name of this task.
 *
 * @return The name of the task as a String.
 */
func (t *MleTask) GetName() string {
	return t.m_name
}

/**
 * Executes task by starting a thread with the Runnable
 * specified during construction.
 */
func (t *MleTask) Invoke() {
	t.m_running = true
	if t.m_name != "" {
		t.m_thread = mle_util.NewThreadWithRunnableAndName(t.m_task, t.m_name)
	} else {
		t.m_thread = mle_util.NewThreadWithRunnable(t.m_task)
	}
	t.m_thread.Start(&t.m_wg)
}

/**
 * Checks if this task is still running.
 *
 * @return A boolean indicating whether the task is still running.
 * If the task is running, <b>true</b> is returned;
 * otherwise, <b>false</b> will be returned.
 */
func (t *MleTask) IsRunning() bool {
	var status bool
	 
	if t.m_running == false {
		status = false
	} else {
		if (t.m_thread.IsAlive()) {
			status = true
		} else {
			t.m_running = false
			status = false
	    }
	}
	return status
}

// String implements the IObject interface.
func (t *MleTask) String() string {
    return t.m_name
}
