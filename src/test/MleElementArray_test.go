/**
 * @file MleElementArray_test.go
 * Created on June 9, 2019. (msm@wizzerworks.com)
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
package core_test

import (
	"strconv"
	"testing"

	mle_util "github.com/mle/runtime/util"
)

type MyElement struct {
	name string
	id int
}

func NewMyElement(name string, id int) *MyElement {
	p := new(MyElement)
	p.name = name
	p.id = id
	return p
}

func (e *MyElement) ToString() string {
	return e.name
}

func (e *MyElement)  IsGreaterThan(element mle_util.IMleElement) bool {
	e1 := e.id
	e2 := element.(*MyElement).id
	if e1 > e2 {
		return true
	}
	return false
}

// The MleElementArray object unit test.
func TestNewMleElementArray(t *testing.T) {
	a := mle_util.NewMleElementArray()
	if a == nil {
		t.Errorf("TestNewMleElementArray: NewMleElementArray() returned nil")
	}
}

func TestNewMleElementArrayWithSize(t *testing.T) {
	a := mle_util.NewMleElementArrayWithSize(10)
	if a == nil {
		t.Errorf("TestNewMleElementArrayWithSize: NewMleElementArrayWithSize() returned nil")
	}
}

func TestAddElement(t *testing.T) {
	a := mle_util.NewMleElementArray()
	if a == nil {
		t.Errorf("TestAddElement: NewMleElementArray() returned nil")
	}

	e := NewMyElement("", 0)
	if e == nil {
		t.Errorf("TestAddElement: NewElement() returned nil")
	}

	n := a.GetNumElements()
	if n != 0 {
		t.Errorf("TestAddElement: want number of elements = 0, got %d", n)
	}
	a.AddElement(e)
	n = a.GetNumElements()
	if n != 1 {
		t.Errorf("TestAddElement: want number of elements = 1, got %d", n)
	}

	for i := 0; i < 9; i++ {
		e = NewMyElement("", 0)
		a.AddElement(e)
	}
	n = a.GetNumElements()
	if n != 10 {
		t.Errorf("TestAddElement: want number of elements = 10, got %d", n)
	}
}

func TestGetNumElements(t *testing.T) {
	a := mle_util.NewMleElementArray()
	if a == nil {
		t.Errorf("TestGetNumElements: NewMleElementArray() returned nil")
	}

	var e mle_util.IMleElement

	for i := 0; i < 10; i++ {
		e = NewMyElement(strconv.Itoa(i), i)
		a.AddElement(e)
	}
	n := a.GetNumElements()
	if n != 10 {
		t.Errorf("TestGetNumElements: want number of elements = 10, got %d", n)
	}
}

func TestGetElementAt(t *testing.T) {
	a := mle_util.NewMleElementArray()
	if a == nil {
		t.Errorf("TestGetElementAt: NewMleElementArray() returned nil")
	}

	var e mle_util.IMleElement

	for i := 0; i < 10; i++ {
		e = NewMyElement(strconv.Itoa(i), i)
		a.AddElement(e)
	}
	n := a.GetNumElements()
	if n != 10 {
		t.Errorf("TestGetElementAt: want number of elements = 10, got %d", n)
	}

	var v1 = a.GetElementAt(3).(*MyElement)
	if v1.name != "3" {
		t.Errorf("TestGetElementAt: want element = 3, got %s", v1.name)
	}

	var v2 = a.GetElementAt(25)
	if v2 != nil {
		t.Errorf("TestGetElementAt: want element = nil")
	}
}

func TestDecrementNumElements(t *testing.T) {
	a := mle_util.NewMleElementArray()
	if a == nil {
		t.Errorf("TestDecrementNumElements: NewMleElementArray() returned nil")
	}

	var e mle_util.IMleElement

	for i := 0; i < 10; i++ {
		e = NewMyElement(strconv.Itoa(i), i)
		a.AddElement(e)
	}
	n := a.GetNumElements()
	if n != 10 {
		t.Errorf("TestDecrementNumElements: want number of elements = 10, got %d", n)
	}

	for i := 0; i < 3; i++ {
	    a.DecrementNumElements()
	}
	n = a.GetNumElements()
	if n != 7 {
		t.Errorf("TestDecrementNumElements: want number of elements = 7, got %d", n)
	}
}

func TestToString(t *testing.T) {
	a := mle_util.NewMleElementArray()
	if a == nil {
		t.Errorf("TestToString: NewMleElementArray() returned nil")
	}

	var e mle_util.IMleElement

	for i := 0; i < 10; i++ {
		e = NewMyElement(strconv.Itoa(i), i)
		a.AddElement(e)
	}
	n := a.GetNumElements()
	if n != 10 {
		t.Errorf("TestToString: want number of elements = 10, got %d", n)
	}

	var str = a.ToString()
	if str != "( 0 1 2 3 4 5 6 7 8 9 )" {
		t.Errorf("TestToString: want elements = ( 0 1 2 3 4 5 6 7 8 9 ), got %s", str)
	}
}

func TestIsGreaterThan(t *testing.T) {
	a := mle_util.NewMleElementArray()
	if a == nil {
		t.Errorf("TestIsGreaterThan: NewMleElementArray() returned nil")
	}

	var e mle_util.IMleElement

	for i := 0; i < 10; i++ {
		e = NewMyElement(strconv.Itoa(i), i)
		a.AddElement(e)
	}
	n := a.GetNumElements()
	if n != 10 {
		t.Errorf("TestIsGreaterThan: want number of elements = 10, got %d", n)
	}

	status := a.IsGreaterThan(3, 2)
	if ! status {
		t.Errorf("TestIsGreaterThan: want false, got true")
	}
	status = a.IsGreaterThan(1, 9)
	if status {
		t.Errorf("TestIsGreaterThan: want false, got true")
	}
}