/*
 * @file MleHeapSortArray.go
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
package util

/**
 * Sorts an array of element keys with the heapsort algorithm.
 *
 * Extends MleHeapArray.
 */
type MleHeapSortArray struct {
	m_array *MleHeapArray
}

/**
 * The default constructor.
 */
func NewMleHeapSortArray() *MleHeapSortArray {
	p := new(MleHeapSortArray)
	p.m_array = NewMleHeapArray()
	return p
}

/**
 * A constructor that initializes with an empty array of length <b>size</b>.
 * 
 * @param size The capacity of the array.
 */
func NewMleHeapSortArrayWithSize(size int) *MleHeapSortArray {
	p := new(MleHeapSortArray)
	p.m_array = NewMleHeapArrayWithSize(size)
	return p
}

/**
 * Sort the input keys in increasing order.
 * This is the entry function for the heapsort algorithm.
 */
func (ha *MleHeapSortArray) Heapsort() {
	ha.m_array.buildHeap()
	ha.sort()
}

/**
 * Move all elements in decreasing key size order into the result list.
 * <p>
 * Precondition: the tree must be a heap.
 * </p>
 */
func (ha *MleHeapSortArray) sort() {
	for ! ha.m_array.heapEmpty() {
		ha.m_array.moveMax();
	}
}

/**
 * Add an element to the array.
 * 
 * @param n The element to add.
 */
 func (ha *MleHeapSortArray) AddElement(n IMleElement) {
	ha.m_array.AddElement(n)
 }

 /**
 * Get the element at the specified index.
 * 
 * @param k The index of the element to get.
 * 
 * @return A reference to an <code>IMleElement</code> is returned.
 */
 func (ha *MleHeapSortArray) GetElementAt(k int) IMleElement {
    return ha.m_array.GetElementAt(k)
}

/**
 * Get the number of items in the queue.
 * 
 * @return The number of elements in the queue is returned.
 */
 func (ha *MleHeapSortArray) GetNumElements() int {
    return ha.m_array.GetNumElements()
}

 /**
 * Clear the queue.
 */
func (ha *MleHeapSortArray) Clear() {
	ha.m_array.Clear()
}

/**
 * Determines if element at index a greater than element at index b.
 * 
 * @param a The first element index.
 * @param b The second element index.
 * 
 * @return <b>true</b> is returned if the first element at index <b>a</b>
 * is greater than the second element at index <b>b</b>.
 */
 func (ha *MleHeapSortArray) IsGreaterThan(a int, b int) bool {
	return ha.m_array.IsGreaterThan(a, b)
}

/**
  * Swap the elements at the indices a and b in the array.
  * 
  * @param a The first element index.
  * @param b The second element index.
  */
 func (ha *MleHeapSortArray) Swap(a int, b int) {
	ha.m_array.Swap(a, b)
}

/**
 * Increase the size of the array.
 * 
 * @param size The size to increase the array by.
 * 
 * @return <b>true</b> is returned if the array grew successfully.
 * Otherwise <b>false</b> is returned.
 */
func (ha *MleHeapSortArray) Grow(size int) bool {
	return ha.m_array.Grow(size)
}

// String implements the IObject interface.
func (ha *MleHeapSortArray) String() string {
	return ha.m_array.String()
}