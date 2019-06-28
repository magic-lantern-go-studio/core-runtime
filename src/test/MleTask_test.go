/**
 * @file MleScheduler_test.go
 * Created on June 28, 2019. (msm@wizzerworks.com)
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
	"time"
	"strconv"
	"testing"

	mle_core "github.com/mle/runtime/core"
	mle_sched "github.com/mle/runtime/scheduler"
)

type testMleTask_TaskTest1 struct {
	mName string
	t *testing.T
}

func testMleTask_NewTaskTest1(t *testing.T) *testMleTask_TaskTest1 {
	p := new(testMleTask_TaskTest1)
	p.mName = "Task Test 1"
	p.t = t
	return p
}

func (task *testMleTask_TaskTest1) Run(done chan bool) {
    for i := 0; i < 10; i++ {
		mle_core.MleLogInfo("Task Test 1: " + strconv.Itoa(i), false)
	}
	// Signal completion.
	if done != nil {
		task.t.Logf("testMleTask_TaskTest1 signaling completion.")
	    done <- true
	}
}

func (task *testMleTask_TaskTest1) String() string {
	return task.mName
}

func TestNewMleTask(t *testing.T) {
	task := mle_sched.NewMleTask(nil)
	if task == nil {
		t.Errorf("TestNewMleTask: NewMleTask() returned nil.")
	}
}

func TestMleTaskInvoke(t *testing.T) {
	task := mle_sched.NewMleTask(nil)
	if task == nil {
		t.Errorf("TestMleTaskInvoke: NewMleTask() returned nil.")
	}
	// Runnable is nil, so invocation should return cleanly.
	t.Logf("testing nil runnable.")
	task.Invoke()

	// Create and run task with runnable.
	runnable := testMleTask_NewTaskTest1(t)
	task = mle_sched.NewMleTask(runnable)
	t.Logf("testing simple runnable.")
	task.Invoke()

	// Wait for task completion.
	t.Logf("waiting for simple runnable to complete.")
	time.Sleep(5000 * time.Millisecond)
}