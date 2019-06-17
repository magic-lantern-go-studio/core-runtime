/**
 * @file Vector.go
 * Created on April 26, 2019. (msm@wizzerworks.com)
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
	"strconv"
	"bytes"
	"fmt"
)

// Import Magic Lantern packages.

/**
 * A vector object that uses slices to emulate the old container/vector package.
 *
 * @see https://github.com/golang/go/wiki/SliceTricks
 */
type Vector []interface{}

// NewVector is the default constructor for creating a Vector.
func NewVector() *Vector {
	p := new(Vector)
	return p
}

// ElementAt returns the Vector element located at the specified index.
func (vector *Vector) ElementAt(index int) interface{} {
	return (*vector)[index]
}

// AddElement adds an elment into the Vector.
func (vector *Vector) AddElement(element interface{}) {
	vector.AppendVector(element)
}

// RemoveElement removes the specified element from the Vector.
func (vector *Vector) RemoveElement(element interface{}) {
	index := vector.Peek(element)
	if index >= 0 {
		// Only try to remove the element if it exists in the Vector.
		vector.Delete(index)
	}
}

// AppendVector adds the specified elements to the Vector.
func (vector *Vector) AppendVector(vs ...interface{}) {
	*vector = append(*vector, vs...)
	// Note that this method is not checking for "uniqueness" of the element being added.
}

// Copy will copy this Vector into the specified destination Vector.
func (vector *Vector) Copy(v []interface{}) {
	v = make([]interface{}, len(*vector))
	copy(v, *vector)
	// or
	// b = append([]T(nil), a...)
	// or
	// b = append(a[:0:0], a...)  // See https://github.com/go101/go101/wiki
}

// Cut is used to remove elements from the Vector.
func (vector *Vector) Cut(i int, j int) {
	//a = append(a[:i], a[j:]...)
	// NOTE If the type of the element is a pointer or a struct with pointer fields,
	// which need to be garbage collected, the above implementation of Cut has a potential
	// memory leak problem: some elements with values are still referenced by slice a and
	// thus can not be collected. The following code can fix this problem:

	copy((*vector)[i:], (*vector)[j:])
	for k, n := len(*vector)-j+i, len(*vector); k < n; k++ {
		(*vector)[k] = nil // or the zero value of T
	}
	*vector = (*vector)[:len(*vector)-j+i]
}

// Delete is used to remove the element specified at the given index.
func (vector *Vector) Delete(i int) {
	//a = append(a[:i], a[i+1:]...)
	// or
	//a = a[:i+copy(a[i:], a[i+1:])]
	// NOTE If the type of the element is a pointer or a struct with pointer fields,
	// which need to be garbage collected, the above implementation of Delete has a potential
	// memory leak problem: some elements with values are still referenced by slice a and
	// thus can not be collected. The following code can fix this problem:

	copy((*vector)[i:], (*vector)[i+1:])
	(*vector)[len(*vector)-1] = nil // or the zero value of T
	*vector = (*vector)[:len(*vector)-1]
}

// Expand is used to increase the size of the Vector.
func (vector *Vector) Expand(i int, j int) {
	*vector = append((*vector)[:i], append(make([]interface{}, j), (*vector)[i:]...)...)
}

// Extend is used to increase the size of the vector.
func (vector *Vector) Extend(j int) {
	*vector = append(*vector, make([]interface{}, j)...)
}

// Insert is used to add a specified element to the Vector at the given index location.
func (vector *Vector) Insert(i int, element interface{}) {
	//a = append(a[:i], append([]T{x}, a[i:]...)...)
	// NOTE The second append creates a new slice with its own underlying storage and copies
	// elements in a[i:] to that slice, and these elements are then copied back to slice a
	// (by the first append). The creation of the new slice (and thus memory garbage) and the
	// second copy can be avoided by using an alternative way:

	*vector = append(*vector, 0 /* use the zero value of the element type */)
	copy((*vector)[i+1:], (*vector)[i:])
	(*vector)[i] = element
}

// InsertVector is used to add a collection of elements, another Vector, to this Vector
// at the given index location.
func (vector *Vector) InsertVector(i int, v []interface{}) {
	*vector = append((*vector)[:i], append(v, (*vector)[i:]...)...)
}

// Push is used to add an element to the Vector. The element will be added
// to the end of the collection, v[n+1].
func (vector *Vector) Push(element interface{}) {
	*vector = append(*vector, element)
}

// Pop is used to remove an element from the Vector. The element removed will be
// the one located in the first index, v[0].
func (vector *Vector) Pop() {
	var element interface{}
	element, *vector = (*vector)[len(*vector)-1], (*vector)[:len(*vector)-1]
	// Note: dropping element here.
	_ = element
}

// Peek will determine if the specified element is contained in the Vector.
//
// Parameters
//   element - The element to test for containership.
//
// Return
//   The index into the Vector where the contained element resides will
//   be returned. Otherwise a value of -1 will be returned.
func (vector *Vector) Peek(element interface{}) int {
	for i := 0; i <= len(*vector); i++ {
		next := (*vector)[i]
		if next == element {
			return i
		}
	}

	return -1
}

// Contains will determine if the specified element is contained in the
// Vector.
//
// Parameters
//   element - The element to test for containership.
//
// Return
//   true will be returned if the specified element is contained in the
//   Vector. Otherwise false will be returned.
func (vector *Vector) Contains(element interface{}) bool {
	var index int = vector.Peek(element)
	if index == -1 {
		return false
	}
	return true
}

// PushFront pushes the specified element onto the Vector into the first
// location, v[0]. The remaining elements are shifted up by 1 (i.e. v[1]
// will now reference what was previously located at v[0]).
func (vector *Vector) PushFront(element interface{}) {
	*vector = append([]interface{}{element}, *vector...)
}

// PopFront removes the first element in the Vector, shifting the remaining
// elements down by 1 (i.e. v[0] is removed and v[0] now references the
// element that was previously at v[1]).
func (vector *Vector) PopFront() {
	var element interface{}
	element, *vector = (*vector)[0], (*vector)[1:]
	// Note: dropping element here.
	_ = element
}

// PrintVector prints the contents of the Vector.
func (vector *Vector) PrintVector() {
	for i := 0; i < len(*vector); i++ {
		fmt.Println("vector[", i, "] = ", (*vector)[i])
	}
}

// String implements the IObject interface.
func (vector *Vector) String() string {
	var buf bytes.Buffer

	for i := 0; i < len(*vector); i++ {
		//var str string = "vector[" + strconv.Itoa(i) + "] = " + strconv.Itoa((*vector)[i].String())
		var str string = "vector[" + strconv.Itoa(i) + "]"
		buf.WriteString(str)
	}

	return buf.String()
}