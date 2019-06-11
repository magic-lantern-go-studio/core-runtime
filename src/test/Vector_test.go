/**
 * @file Vector_test.go
 * Created on April 30, 2019. (msm@wizzerworks.com)
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
package util_test

import (
	"testing"

	mle_util "github.com/mle/runtime/util"
)

// The Vector object unit test.
func TestNewVector(t *testing.T) {
	// Construct an empty vector.
	v := mle_util.NewVector()
	if v == nil {
		t.Errorf("TestNewVector: NewVector() returned nil")
	}
	if len(*v) != 0 {
		t.Errorf("TestNewVector: NewVector() length should be empty (0)")
	}
}

func TestAppendVector(t *testing.T) {
	// Construct an empty vector.
	v := mle_util.NewVector()
	if v == nil {
		t.Errorf("TestAppendVector: NewVector() returned nil")
	}
	if len(*v) != 0 {
		t.Errorf("TestAppendVector: NewVector() length should be empty (0)")
	}

	type element struct {
		id int
	}

	// Append first element.
	a := new(element)
	a.id = 1
	v.AppendVector(a)
	size := len(*v)
	if size != 1 {
		t.Errorf("TestAppendVector: want length = 1, got %d", size)
	}

	// Append 9 more elements.
	for i := 0; i < 9; i++ {
		b := new(element)
		b.id = i
		v.AppendVector(b)
	}
	size = len(*v)
	if size != 10 {
		t.Errorf("TestAppendVector: want length = 10, got %d", size)
	}

	// Append 5 more elements.
	b := [5]element{}
	b[0].id = 1
	b[1].id = 2
	b[2].id = 3
	b[3].id = 4
	b[4].id = 5
	v.AppendVector(b[0], b[1], b[2], b[3], b[4])
	size = len(*v)
	if size != 15 {
		t.Errorf("TestAppendVector: want length = 15, got %d", size)
	}
}

func TestElementAt(t *testing.T) {
	// Construct an empty vector.
	v := mle_util.NewVector()
	if v == nil {
		t.Errorf("TestElementAt: NewVector() returned nil")
	}
	if len(*v) != 0 {
		t.Errorf("TestElementAt: NewVector() length should be empty (0)")
	}

	type element struct {
		id int
	}

	// Append 5 elements.
	b := [5]element{}
	b[0].id = 1
	b[1].id = 2
	b[2].id = 3
	b[3].id = 4
	b[4].id = 5
	v.AppendVector(b[0], b[1], b[2], b[3], b[4])

	// Check Vector contents.
	size := len(*v)
	for i := 0; i < size; i++ {
		var e element = v.ElementAt(i).(element)
		if e.id != i+1 {
			t.Errorf("TestElementAt: want id = %d, got %d ", i+1, e.id)
		}
	}
}

func TestCopy(t *testing.T) {
	// Construct an empty vector.
	v := mle_util.NewVector()
	if v == nil {
		t.Errorf("TestCopy: NewVector() returned nil")
	}
	if len(*v) != 0 {
		t.Errorf("TestCopy: NewVector() length should be empty (0)")
	}

	type element struct {
		id int
	}

	// Append 5 elements.
	b := [5]element{}
	b[0].id = 1
	b[1].id = 2
	b[2].id = 3
	b[3].id = 4
	b[4].id = 5
	v.AppendVector(b[0], b[1], b[2], b[3], b[4])

	// Create a new Vector to copy into.
	c := mle_util.NewVector()
	if c == nil {
		t.Errorf("TestCopy: NewVector() returned nil")
	}
	if len(*c) != 0 {
		t.Errorf("TestCopy: NewVector() length should be empty (0)")
	}
	v.Copy(*c)

	// Check copied Vector contents.
	size := len(*c)
	for i := 0; i < size; i++ {
		var e element = c.ElementAt(i).(element)
		if e.id != i+1 {
			t.Errorf("TestCopy: want id = %d, got %d ", i+1, e.id)
		}
	}
}

func TestCut(t *testing.T) {
	// Construct an empty vector.
	v := mle_util.NewVector()
	if v == nil {
		t.Errorf("TestCut: NewVector() returned nil")
	}
	if len(*v) != 0 {
		t.Errorf("TestCut: NewVector() length should be empty (0)")
	}

	type element struct {
		id int
	}

	// Append 5 elements.
	b := [5]element{}
	b[0].id = 1
	b[1].id = 2
	b[2].id = 3
	b[3].id = 4
	b[4].id = 5
	v.AppendVector(b[0], b[1], b[2], b[3], b[4])

	// Cut out elements 4 and 5; note first index is non-inclusive.
	v.Cut(3, 5)

	// Check Vector contents.
	size := len(*v)
	if size != 3 {
		t.Errorf("TestCut: want size = %d, got %d ", 2, size)
	}
	for i := 0; i < size; i++ {
		var e element = v.ElementAt(i).(element)
		//fmt.Printf("element %d: id = %d\n", i, e.id)
		t.Logf("TestCut: element %d: id = %d\n", i, e.id)
		if e.id != i+1 {
			t.Errorf("TestCut: want id = %d, got %d ", i+1, e.id)
		}
	}
}

func TestDelete(t *testing.T) {
	// Construct an empty vector.
	v := mle_util.NewVector()
	if v == nil {
		t.Errorf("TestDelete: NewVector() returned nil")
	}
	if len(*v) != 0 {
		t.Errorf("TestDelete: NewVector() length should be empty (0)")
	}

	type element struct {
		id int
	}

	// Append 5 elements.
	b := [5]element{}
	b[0].id = 1
	b[1].id = 2
	b[2].id = 3
	b[3].id = 4
	b[4].id = 5
	v.AppendVector(b[0], b[1], b[2], b[3], b[4])

	// Delete element 3.
	v.Delete(3)

	// Check Vector contents.
	size := len(*v)
	if size != 4 {
		t.Errorf("TestDelete: want size = %d, got %d ", 4, size)
	}
	capacity := cap(*v)
	if capacity != 5 {
		t.Errorf("TestDelete: want capacity = %d, got %d ", 5, capacity)
	}
}

func TestExtend(t *testing.T) {
	// Construct an empty vector.
	v := mle_util.NewVector()
	if v == nil {
		t.Errorf("TestExtend: NewVector() returned nil")
	}
	if len(*v) != 0 {
		t.Errorf("TestExtend: NewVector() length should be empty (0)")
	}
	if cap(*v) != 0 {
		t.Errorf("TestExtend: NewVector() capacity should be empty (0)")
	}

	type element struct {
		id int
	}

	// Append 5 elements.
	b := [5]element{}
	b[0].id = 1
	b[1].id = 2
	b[2].id = 3
	b[3].id = 4
	b[4].id = 5
	v.AppendVector(b[0], b[1], b[2], b[3], b[4])

	// Extend Vector capacity. New element entries will be nil.
	v.Extend(3)

	// Check Vector contents.
	size := len(*v)
	if size != 8 {
		t.Errorf("TestExtend: want size = %d, got %d ", 8, size)
	}
	capacity := cap(*v)
	if capacity <= 8 {
		t.Errorf("TestExtend: want capacity > %d, got %d ", 8, capacity)
	}
	for i := 0; i < size; i++ {
		var iface interface{} = v.ElementAt(i)
		if iface != nil {
			var e element = iface.(element)
			//fmt.Printf("element %d: id = %d\n", i, e.id)
			t.Logf("TestExtend: element %d: id = %d\n", i, e.id)
			if e.id != i+1 {
				t.Errorf("TestExtend: want id = %d, got %d ", i+1, e.id)
			}
		}
	}
}

func TestExpand(t *testing.T) {
	// Construct an empty vector.
	v := mle_util.NewVector()
	if v == nil {
		t.Errorf("TestExpand: NewVector() returned nil")
	}
	if len(*v) != 0 {
		t.Errorf("TestExpand: NewVector() length should be empty (0)")
	}
	if cap(*v) != 0 {
		t.Errorf("TestExpand: NewVector() capacity should be empty (0)")
	}

	type element struct {
		id int
	}

	// Append 5 elements.
	b := [5]element{}
	b[0].id = 1
	b[1].id = 2
	b[2].id = 3
	b[3].id = 4
	b[4].id = 5
	v.AppendVector(b[0], b[1], b[2], b[3], b[4])

	// Expand Vector capacity. New element entries will start at index 3 and should be nil.
	v.Expand(3, 3)

	// Check Vector contents.
	size := len(*v)
	if size != 8 {
		t.Errorf("TestExpand: want size = %d, got %d ", 8, size)
	}
	capacity := cap(*v)
	if capacity <= 8 {
		t.Errorf("TestExpand: want capacity > %d, got %d ", 8, capacity)
	}
	for i := 0; i < size; i++ {
		var iface interface{} = v.ElementAt(i)
		if iface != nil {
			var e element = iface.(element)
			//fmt.Printf("element %d: id = %d\n", i, e.id)
			t.Logf("TestExpand: element %d: id = %d\n", i, e.id)
			if i < 3 {
				if e.id != i+1 {
					t.Errorf("TestExpand: want id = %d, got %d ", i+1, e.id)
				}
			}
			if i > 5 {
				if e.id != i-2 {
					t.Errorf("TestExpand: want id = %d, got %d ", i-2, e.id)
				}
			}
		}
	}
}

func TestInsert(t *testing.T) {
	// Construct an empty vector.
	v := mle_util.NewVector()
	if v == nil {
		t.Errorf("TestInsert: NewVector() returned nil")
	}
	if len(*v) != 0 {
		t.Errorf("TestInsert: NewVector() length should be empty (0)")
	}
	if cap(*v) != 0 {
		t.Errorf("TestInsert: NewVector() capacity should be empty (0)")
	}

	type element struct {
		id int
	}

	// Append 5 elements.
	b := [5]element{}
	b[0].id = 1
	b[1].id = 2
	b[2].id = 3
	b[3].id = 4
	b[4].id = 5
	v.AppendVector(b[0], b[1], b[2], b[3], b[4])

	// Extend Vector capacity. New element entries will be nil.
	c := new(element)
	c.id = 6
	v.Insert(3, *c)

	// Check Vector contents.
	size := len(*v)
	if size != 6 {
		t.Errorf("TestInsert: want size = %d, got %d ", 6, size)
	}
	for i := 0; i < size; i++ {
		var iface interface{} = v.ElementAt(i)
		if iface != nil {
			var e element = iface.(element)
			//fmt.Printf("element %d: id = %d\n", i, e.id)
			t.Logf("TestInsert: element %d: id = %d\n", i, e.id)
			if i == 3 {
				if e.id != 6 {
					t.Errorf("TestInsert: want id = %d, got %d ", 6, e.id)
				}
			}
		}
	}
}

func TestInsertVector(t *testing.T) {
	// Construct an empty vector.
	v := mle_util.NewVector()
	if v == nil {
		t.Errorf("TestInsertVector: NewVector() returned nil")
	}
	if len(*v) != 0 {
		t.Errorf("TestInsertVector: NewVector() length should be empty (0)")
	}
	if cap(*v) != 0 {
		t.Errorf("TestInsertVector: NewVector() capacity should be empty (0)")
	}

	type element struct {
		id int
	}

	// Append 5 elements.
	b := [5]element{}
	b[0].id = 1
	b[1].id = 2
	b[2].id = 3
	b[3].id = 4
	b[4].id = 5
	v.AppendVector(b[0], b[1], b[2], b[3], b[4])

	// Extend Vector by inserting 5 more elements, starting at index 3.
	c := [5]element{}
	c[0].id = 6
	c[1].id = 7
	c[2].id = 8
	c[3].id = 9
	c[4].id = 10
	// Note: We have to copy element slice into an interface slice
	// because the two types do not have the same representation in memory.
	s := make([]interface{}, len(c))
	for i, v := range c {
		s[i] = v
	}
	v.InsertVector(3, s)

	// Check Vector contents.
	size := len(*v)
	if size != 10 {
		t.Errorf("TestInsertVector: want size = %d, got %d ", 10, size)
	}
	for i := 0; i < size; i++ {
		var iface interface{} = v.ElementAt(i)
		if iface != nil {
			var e element = iface.(element)
			//fmt.Printf("element %d: id = %d\n", i, e.id)
			t.Logf("TestInsertVector: element %d: id = %d\n", i, e.id)
			if i > 2 && i < 8 {
				if e.id != i+3 {
					t.Errorf("TestInsertVector: want id = %d, got %d ", i+3, e.id)
				}
			}
		}
	}
}

func TestPush(t *testing.T) {
	// Construct an empty vector.
	v := mle_util.NewVector()
	if v == nil {
		t.Errorf("TestPush: NewVector() returned nil")
	}
	if len(*v) != 0 {
		t.Errorf("TestPush: NewVector() length should be empty (0)")
	}
	if cap(*v) != 0 {
		t.Errorf("TestPush: NewVector() capacity should be empty (0)")
	}

	type element struct {
		id int
	}

	// Append 5 elements.
	b := [5]element{}
	b[0].id = 1
	b[1].id = 2
	b[2].id = 3
	b[3].id = 4
	b[4].id = 5
	v.AppendVector(b[0], b[1], b[2], b[3], b[4])

	// Push a new element.
	c := new(element)
	c.id = 6
	v.Push(*c)

	// Check Vector contents.
	size := len(*v)
	if size != 6 {
		t.Errorf("TestPush: want size = %d, got %d ", 6, size)
	}
	for i := 0; i < size; i++ {
		var iface interface{} = v.ElementAt(i)
		if iface != nil {
			var e element = iface.(element)
			//fmt.Printf("element %d: id = %d\n", i, e.id)
			t.Logf("TestPush: element %d: id = %d\n", i, e.id)
			if e.id != i+1 {
				t.Errorf("TestPush: want id = %d, got %d ", i+1, e.id)
			}
		}
	}
}

func TestPop(t *testing.T) {
	// Construct an empty vector.
	v := mle_util.NewVector()
	if v == nil {
		t.Errorf("TestPop: NewVector() returned nil")
	}
	if len(*v) != 0 {
		t.Errorf("TestPop: NewVector() length should be empty (0)")
	}
	if cap(*v) != 0 {
		t.Errorf("TestPop: NewVector() capacity should be empty (0)")
	}

	type element struct {
		id int
	}

	// Append 5 elements.
	b := [5]element{}
	b[0].id = 1
	b[1].id = 2
	b[2].id = 3
	b[3].id = 4
	b[4].id = 5
	v.AppendVector(b[0], b[1], b[2], b[3], b[4])

	// Pop the last element.
	v.Pop()

	// Check Vector contents.
	size := len(*v)
	if size != 4 {
		t.Errorf("TestPop: want size = %d, got %d ", 4, size)
	}
	for i := 0; i < size; i++ {
		var iface interface{} = v.ElementAt(i)
		if iface != nil {
			var e element = iface.(element)
			//fmt.Printf("element %d: id = %d\n", i, e.id)
			t.Logf("TestPop: element %d: id = %d\n", i, e.id)
			if e.id != i+1 {
				t.Errorf("TestPop: want id = %d, got %d ", i+1, e.id)
			}
		}
	}
}

func TestPeek(t *testing.T) {
	// Construct an empty vector.
	v := mle_util.NewVector()
	if v == nil {
		t.Errorf("TestPeek: NewVector() returned nil")
	}
	if len(*v) != 0 {
		t.Errorf("TestPeek: NewVector() length should be empty (0)")
	}
	if cap(*v) != 0 {
		t.Errorf("TestPeek: NewVector() capacity should be empty (0)")
	}

	type element struct {
		id int
	}

	// Append 5 elements.
	b := [5]element{}
	b[0].id = 1
	b[1].id = 2
	b[2].id = 3
	b[3].id = 4
	b[4].id = 5
	v.AppendVector(b[0], b[1], b[2], b[3], b[4])

	// Peek to see if the element exists in the Vector.
	index := v.Peek(b[1])
	if index != 1 {
		t.Errorf("TestPeek: want index = %d, got %d ", 1, index)
	}
}

func TestPushFront(t *testing.T) {
	// Construct an empty vector.
	v := mle_util.NewVector()
	if v == nil {
		t.Errorf("TestPushFront: NewVector() returned nil")
	}
	if len(*v) != 0 {
		t.Errorf("TestPushFront: NewVector() length should be empty (0)")
	}
	if cap(*v) != 0 {
		t.Errorf("TestPushFront: NewVector() capacity should be empty (0)")
	}

	type element struct {
		id int
	}

	// Append 5 elements.
	b := [5]element{}
	b[0].id = 1
	b[1].id = 2
	b[2].id = 3
	b[3].id = 4
	b[4].id = 5
	v.AppendVector(b[0], b[1], b[2], b[3], b[4])

	// Push a new element into the Vector and shift the index locations.
	c := new(element)
	c.id = 6
	v.PushFront(*c)

	// Check Vector contents.
	size := len(*v)
	if size != 6 {
		t.Errorf("TestPushFront: want size = %d, got %d ", 6, size)
	}
	for i := 0; i < size; i++ {
		var iface interface{} = v.ElementAt(i)
		if iface != nil {
			var e element = iface.(element)
			//fmt.Printf("element %d: id = %d\n", i, e.id)
			t.Logf("TestPushFront: element %d: id = %d\n", i, e.id)
			if i == 0 {
				if e.id != 6 {
					t.Errorf("TestPushFront: want id = %d, got %d ", 6, e.id)
				}
			} else {
				if e.id != i {
					t.Errorf("TestPushFront: want id = %d, got %d ", i, e.id)
				}
			}
		}
	}
}

func TestPopFront(t *testing.T) {
	// Construct an empty vector.
	v := mle_util.NewVector()
	if v == nil {
		t.Errorf("TestPopFront: NewVector() returned nil")
	}
	if len(*v) != 0 {
		t.Errorf("TestPopFront: NewVector() length should be empty (0)")
	}
	if cap(*v) != 0 {
		t.Errorf("TestPopFront: NewVector() capacity should be empty (0)")
	}

	type element struct {
		id int
	}

	// Append 5 elements.
	b := [5]element{}
	b[0].id = 1
	b[1].id = 2
	b[2].id = 3
	b[3].id = 4
	b[4].id = 5
	v.AppendVector(b[0], b[1], b[2], b[3], b[4])

	// Pop the first element in the Vector (b[0]).
	v.PopFront()

	// Check Vector contents.
	size := len(*v)
	if size != 4 {
		t.Errorf("TestPopFront: want size = %d, got %d ", 4, size)
	}
	for i := 0; i < size; i++ {
		var iface interface{} = v.ElementAt(i)
		if iface != nil {
			var e element = iface.(element)
			//fmt.Printf("element %d: id = %d\n", i, e.id)
			t.Logf("TestPopFront: element %d: id = %d\n", i, e.id)
			if e.id != i+2 {
				t.Errorf("TestPopFront: want id = %d, got %d ", i+2, e.id)
			}
		}
	}
}
