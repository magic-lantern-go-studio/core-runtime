/**
 * @file Thread_test.go
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
package mle_test

import (
	"strconv"
	"sync"
	"time"
	"testing"

	mle_util "github.com/mle/runtime/util"
)

type testThread_myRunnable struct {
	mName string
	t *testing.T
}

func testThread_newMyRunnable(t *testing.T) *testThread_myRunnable {
	p := new(testThread_myRunnable)
	p.mName = "MyThreadRunnable"
	p.t = t
	return p
}

func (r *testThread_myRunnable) Run(done chan bool) {
	r.t.Logf("Running thread " + r.mName)
	time.Sleep(100 * time.Millisecond)
	// Signal completion.
	if done != nil {
	    done <- true
	}
}

func (r *testThread_myRunnable) String() string {
	return r.mName
}

func TestNewThread(t *testing.T) {
	thread := mle_util.NewThread()
	if thread == nil {
		t.Errorf("TestNewThread: NewThread() returned nil")
	}
	thread.Run(nil)
}

func TestNewThreadWithRunnable(t *testing.T) {
    r := testThread_newMyRunnable(t)

	thread := mle_util.NewThreadWithRunnableAndName(r, "MyThread")
	if thread == nil {
		t.Errorf("TestNewThreadWithRunnable: NewThreadWithRunnable() returned nil")
	}

	// Execute thread but don't wait for completion.
	thread.Run(nil)

	// Execute thread and wait for completion based on done channel.
	var done = make(chan bool)
	thread.Run(done)

	time.Sleep(1000 * time.Millisecond)
}

func TestThreadStart(t *testing.T) {
	r := testThread_newMyRunnable(t)

	thread := mle_util.NewThreadWithRunnableAndName(r, "MyThread")
	if thread == nil {
		t.Errorf("TestStart: NewThreadWithRunnable() returned nil")
	}

	// Start the thread execution.
	var wg sync.WaitGroup
	thread.Start(&wg)
	// And wait for it to complete.
	wg.Wait()

	time.Sleep(1000 * time.Millisecond)
}

func TestThreadMutlipleThreads(t *testing.T) {
	r := testThread_newMyRunnable(t)

	var threads [10](*mle_util.Thread)
	for i := 0; i < 10; i++ {
		tname := "Thread-" + strconv.Itoa(i)
		threads[i] = mle_util.NewThreadWithRunnableAndName(r, tname)
	}

	// Start the threads' execution.
	var wg sync.WaitGroup
	for i := 0; i < len(threads); i++ {
		threads[i].Start(&wg)
	}
	// And wait for the threads to complete.
	wg.Wait()
}