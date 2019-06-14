/**
 * @file Runnable_test.go
 * Created on June 13, 2019. (msm@wizzerworks.com)
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
	"testing"
)

type testRunnable_myRunnable struct {
	mName string
	t *testing.T
}

func testRunnable_newMyRunnable(t *testing.T) *testRunnable_myRunnable {
	p := new(testRunnable_myRunnable)
	p.mName = "MyRunnable"
	p.t = t
	return p
}

func (r *testRunnable_myRunnable) Run(done chan bool) {
	if done != nil {
	    done <- true
	}
}

func TestRunnable(t *testing.T) {
	runnable := testRunnable_newMyRunnable(t)
	var done = make(chan bool)

	// Execute run with done notification.
	go runnable.Run(done)
	<-done

	// Execute run without a done notification.
	go runnable.Run(nil)
}
