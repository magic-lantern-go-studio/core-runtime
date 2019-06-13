/**
 * @file MleHeapArray_test.go
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
package mle_test

// Import go packages.
import (
	"testing"

	mle_util "github.com/mle/runtime/util"
)

type ha_MyElement struct {
	name string
	id int
}

func newHeapMyElement(name string, id int) *ha_MyElement {
	p := new(ha_MyElement)
	p.name = name
	p.id = id
	return p
}

func (e *ha_MyElement) String() string {
	return e.name
}

func (e *ha_MyElement)  IsGreaterThan(element mle_util.IMleElement) bool {
	e1 := e.id
	e2 := element.(*ha_MyElement).id
	if e1 > e2 {
		return true
	}
	return false
}

// The MleHeapArray object unit test.
func TestNewMleHeapArray(t *testing.T) {
	a := mle_util.NewMleHeapArray()
	if a == nil {
		t.Errorf("TestNewMleHeapArray: NewMleHeapArray() returned nil")
	}
}

// Test algorithm with a fixed input.
func TestHeapSort(t *testing.T) {
	t.Logf("TestHeapSort: Heapsort Algorithm\n")
	
	var tk []int
	tk = []int{3, 2, 1, 5, 4, 6, 8, 7, 8}
	aha := mle_util.NewMleHeapSortArrayWithSize(len(tk))
    for k := 0; k < len(tk); k++ {
		aha.AddElement(mle_util.NewMlePQElementWithKey(tk[k], nil))
    }

	t.Logf("TestHeapSort: Unsorted list %s\n", aha.String())
    aha.Heapsort()

	t.Logf("TestHeapSort: Sorted list %s\n", aha.String())
	var str = aha.String()
	if str != "( 1 2 3 4 5 6 7 8 8 )" {
		t.Errorf("TestHeapSort: want elements = ( 1 2 3 4 5 6 7 8 8 ), got %s", str)
	}
}
