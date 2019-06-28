/**
 * @file MleScheduler_test.go
 * Created on June 18, 2019. (msm@wizzerworks.com)
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
	"testing"
	"time"

	mle_core "github.com/mle/runtime/core"
	mle_sched "github.com/mle/runtime/scheduler"
)

var numPhases int = 4
var numTasks  int = 5

type testMleScheduler_ThreadTest1 struct {
	mName string
	t *testing.T
}

func testMleScheduler_NewThreadTest1(t *testing.T) *testMleScheduler_ThreadTest1 {
	p := new(testMleScheduler_ThreadTest1)
	p.mName = "Thread Test 1"
	p.t = t
	return p
}

func (task *testMleScheduler_ThreadTest1) Run(done chan bool) {
    for i := 0; i < 10; i++ {
		mle_core.MleLogInfo("Thread Test 1: " + strconv.Itoa(i), false)
	}
	// Signal completion.
	if done != nil {
		task.t.Logf("testMleScheduler_ThreadTest1 signaling completion.")
	    done <- true
	}
}

func (task *testMleScheduler_ThreadTest1) String() string {
	return task.mName
}

type testMleScheduler_ThreadTest2 struct {
	mName string
	t *testing.T
}

func testMleScheduler_NewThreadTest2(t *testing.T) *testMleScheduler_ThreadTest2 {
	p := new(testMleScheduler_ThreadTest2)
	p.mName = "Thread Test 2"
	p.t = t
	return p
}

func (task *testMleScheduler_ThreadTest2) Run(done chan bool) {
    for i := 0; i < 10; i++ {
		mle_core.MleLogInfo("Thread Test 2: " + strconv.Itoa(i), false)
	}
	// Signal completion.
	if done != nil {
		task.t.Logf("testMleScheduler_ThreadTest2 signaling completion.")
	    done <- true
	}
}

func (task *testMleScheduler_ThreadTest2) String() string {
	return task.mName
}

type testMleScheduler_ThreadTest3 struct {
	mName string
	t *testing.T
}

func testMleScheduler_NewThreadTest3(t *testing.T) *testMleScheduler_ThreadTest3 {
	p := new(testMleScheduler_ThreadTest3)
	p.mName = "Thread Test 3"
	p.t = t
	return p
}

func (task *testMleScheduler_ThreadTest3) Run(done chan bool) {
    for i := 0; i < 10; i++ {
		mle_core.MleLogInfo("Thread Test 3: " + strconv.Itoa(i), false)
	}
	// Signal completion.
	if done != nil {
		task.t.Logf("testMleScheduler_ThreadTest3 signaling completion.")
	    done <- true
	}
}

func (task *testMleScheduler_ThreadTest3) String() string {
	return task.mName
}

type testMleScheduler_ThreadTest4 struct {
	mName string
	t *testing.T
}

func testMleScheduler_NewThreadTest4(t *testing.T) *testMleScheduler_ThreadTest4 {
	p := new(testMleScheduler_ThreadTest4)
	p.mName = "Thread Test 4"
	p.t = t
	return p
}

func (task *testMleScheduler_ThreadTest4) Run(done chan bool) {
    for i := 0; i < 10; i++ {
		mle_core.MleLogInfo("Thread Test 4: " + strconv.Itoa(i), false)
	}
	// Signal completion.
	if done != nil {
		task.t.Logf("testMleScheduler_ThreadTest4 signaling completion.")
	    done <- true
	}
}

func (task *testMleScheduler_ThreadTest4) String() string {
	return task.mName
}

type testMleScheduler_ThreadTest5 struct {
	mName string
	t *testing.T
}

func testMleScheduler_NewThreadTest5(t *testing.T) *testMleScheduler_ThreadTest5 {
	p := new(testMleScheduler_ThreadTest5)
	p.mName = "Thread Test 5"
	p.t = t
	return p
}

func (task *testMleScheduler_ThreadTest5) Run(done chan bool) {
    for i := 0; i < 10; i++ {
		mle_core.MleLogInfo("Thread Test 5: " + strconv.Itoa(i), false)
	}
	// Signal completion.
	if done != nil {
		task.t.Logf("testMleScheduler_ThreadTest5 signaling completion.")
	    done <- true
	}
}

func (task *testMleScheduler_ThreadTest5) String() string {
	return task.mName
}

/**
 * Test the Scheduler by registering three phases with two tasks to run in each phase.
 * The scheulder is then run to completion and some of the tasks are removed.
 */
func TestSchdulerTest1(t *testing.T) {
    scheduler := mle_sched.NewMleScheduler()
		 
	/* Register three phases. */
	for i := 0; i < 3; i++ {
		name := "Test Phase " + strconv.Itoa(i+1)
		phase := mle_sched.NewMlePhaseWithName(name)
 
		/** Create and register task 1 */
		one := testMleScheduler_NewThreadTest1(t)
		task := mle_sched.NewMleTaskWithName(one,"Test 1")
		phase.AddTask(task)
			 
		/** Create and register task 2 */
		two := testMleScheduler_NewThreadTest2(t)
		task = mle_sched.NewMleTaskWithName(two,"Test 2")
		phase.AddTask(task)
			 
		/** Register phase */
		scheduler.AddPhase(phase);
	}

	/* Print number of registered phases. */
	msg := "Number of scheduled phases: " + strconv.Itoa(scheduler.GetNumberOfPhases())
	t.Logf(msg)

	msg = "*** Test 1 ***"
	t.Logf(msg)

	if scheduler.GetNumberOfPhases() != 3 {
		t.Errorf("TestSchdulerTest1: GetNumberOfPhases() expected 3, got %d", scheduler.GetNumberOfPhases())
	}
	for i := 0; i < 3; i++ {
		phase := scheduler.GetPhase(i)
		str := "Test Phase " + strconv.Itoa(i+1)
		if phase.GetName() != str {
			t.Errorf("TestSchdulerTest1: expected %s, got %s", str, phase.GetName())
		}
		task1 := phase.GetTask(0)
		str = "Test 1"
		if task1.GetName() != str {
			t.Errorf("TestSchdulerTest1: expected %s, got %s", str, task1.GetName())
		}
		task2 := phase.GetTask(1)
		str = "Test 2"
		if task2.GetName() != str {
			t.Errorf("TestSchdulerTest1: expected %s, got %s", str, task2.GetName())
		}
	}
	
	/* Print tasks. */
	//scheduler.Dump()
        
	/* Invoke tasks, that is, execute them. */
    go func() {
		// Wait 2 seconds, and then trigger an exit. Otherwise the scheduler
		// Run method called below will never exit.
		time.Sleep(2000 * time.Millisecond)
		scheduler.SetExitOk()
	}()
	scheduler.Run(nil)
}