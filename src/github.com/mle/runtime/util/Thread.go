/*
 * @file Thread.go
 * Created on June 12, 2019. (msm@wizzerworks.com)
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
package util

// Import go packages.
import (
	"sync"
)

type Thread struct {
	// The name of the thread.
	m_name string
	// The object that will do the thread execution.
	m_runnable Runnable
	// The object is alive and running.
	m_alive bool
	// Channel for observing when a thread has completed.
	m_done chan bool
}

func NewThread() *Thread {
	p := new(Thread)
	p.m_name = "Unknown Thread"
	p.m_runnable = nil
	p.m_alive = false
	p.m_done = make(chan bool)
	return p
}

func NewThreadWithRunnable(runnable Runnable) *Thread {
	p := new(Thread)
	p.m_name = "Unknown Thread"
	p.m_runnable = runnable
	p.m_alive = false
	p.m_done = make(chan bool)
	return p
}

func NewThreadWithRunnableAndName(runnable Runnable, name string) *Thread {
	p := new(Thread)
	p.m_name = name
	p.m_runnable = runnable
	p.m_alive = false
	p.m_done = make(chan bool)
	return p
}

// Run will execute the thread.
//
// If this thread was constructed using a separate Runnable run object,
// then that Runnable object's run method is called; otherwise,
// this method does nothing and returns.
func (t *Thread) Run(done chan bool) {
	if t.m_runnable != nil {
		t.m_alive = true
		go t.m_runnable.Run(done)
		if done != nil {
			// Wait for Run goroutine to complete.
			<-done
		}
		t.m_alive = false
	}
}

// Start will begin the thread execution.
//
// Parameters
//   wg - A reference to a synchronization WaitGroup.
func (t *Thread) Start(wg *sync.WaitGroup) {
	if t.m_runnable != nil {
		// Start the runnable.
		t.m_alive = true
		wg.Add(1)
		go t.m_runnable.Run(t.m_done)

		// Establish a goroutine to indicate when the thread has
		// completed running.
		go func(waitgroup *sync.WaitGroup) {
			// Wait for runnable to complete.
			<-t.m_done
			defer waitgroup.Done()
			t.m_alive = false
		}(wg)
	}
}

// IsAlive can be used to determine if the Thread is alive and active.
func (t *Thread) IsAlive() bool {
	return t.m_alive
}

// String returns a string representation of this thread, including the thread's name,
// priority, and thread group.
func (t *Thread) String() string {
	return t.m_name
}
