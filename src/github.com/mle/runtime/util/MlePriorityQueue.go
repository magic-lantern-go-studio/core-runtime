/*
 * @file MlePriorityQueue.go
 * Created on June 5, 2019. (msm@wizzerworks.com)
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
 * Implements a Priority Queue using a heap tree.
 *
 * Extends MleHeapArray.
 */
type MlePriorityQueue struct {
    m_array *MleHeapArray  // MlePriorityQueue extends MleHeapArray.
}

/**
 * The default constructor.
 */
func NewMlePriorityQueue() *MlePriorityQueue {
	p := new(MlePriorityQueue)
	p.m_array = NewMleHeapArray()
	return p
}

/**
 * A constructor that initializes with an empty array of length <b>size</b>.
 * 
 * @param size The initial capacity of the array.
 */
func NewMlePriorityQueueWithSize(size int) *MlePriorityQueue {
	p := new(MlePriorityQueue)
	p.m_array = NewMleHeapArrayWithSize(size)
	return p
}

/**
 * Add an element to the queue.
 * 
 * @param n The element to add.
 */
func (pq *MlePriorityQueue) AddElement(n IMleElement) {
	pq.m_array.AddElement(n)
	pq.flowUp(pq.m_array.m_lastHeap)
}

/**
 * Get and remove maximum element.
 * 
 * @return If queue is empty, <b>null</b> is returned.
 */
func (pq *MlePriorityQueue) GetMaxElement() IMleElement {
	if pq.m_array.heapEmpty() {
		return nil
	}
		   
	maxNode := pq.m_array.GetElementAt(pq.m_array.moveMax())
	pq.m_array.decrementNumElements()
		   
	return maxNode
}

/**
 * Peek into queue for specified item.
 * 
 * @param k The index of the element to check.
 * 
 * @return A reference to the item in the queue is returned
 * without actually removing it from the queue.
 */
func (pq *MlePriorityQueue) Peek(k int) IMleElement {
    var result IMleElement = nil
		 
	if k >= 0 && k < pq.m_array.GetNumElements() {
	    return pq.m_array.GetElementAt(k)
	}
		 
	return result
}

/**
 * Move up a key until it satisfies the heap order.
 * 
 * @param k The index of the element to move up.
 */
func (pq *MlePriorityQueue) flowUp(k int) {
	// Swap the key at k with its parents along the path to the root
	// until it finds the place where the heap order is fulfilled.
	for pq.m_array.parent(k) != pq.m_array.m_nullIndex &&
	    pq.m_array.IsGreaterThan(k, pq.m_array.parent(k)) {
		pq.m_array.Swap(pq.m_array.parent(k), k)
		k = pq.m_array.parent(k)
	}
}

/**
 * Move down a key until it satisfies the heap order.
 * 
 * @param k The index of the element to move down.
 */
func (pq *MlePriorityQueue) flowDown(k int) {
	pq.m_array.heapify(k)
}

/**
 * Clear the queue.
 */
func (pq *MlePriorityQueue) Clear() {
	pq.m_array.Clear()
}

/**
 * Get the element at the specified index.
 * 
 * @param k The index of the element to get.
 * 
 * @return A reference to an <code>IMleElement</code> is returned.
 */
 func (pq *MlePriorityQueue) GetElementAt(k int) IMleElement {
    return pq.m_array.GetElementAt(k)
}

/**
 * Get the number of items in the queue.
 * 
 * @return The number of elements in the queue is returned.
 */
 func (pq *MlePriorityQueue) GetNumElements() int {
    return pq.m_array.GetNumElements()
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
func (pq *MlePriorityQueue) IsGreaterThan(a int, b int) bool {
	return pq.m_array.IsGreaterThan(a, b)
}

/**
  * Swap the elements at the indices a and b in the array.
  * 
  * @param a The first element index.
  * @param b The second element index.
  */
func (pq *MlePriorityQueue) Swap(a int, b int) {
	pq.m_array.Swap(a, b)
}

/**
 * Determine if the queue is empty.
 *
 * <b>true</b> is returned if the queue is empty. Otherwise,
 * <b>false</b> is returned.
 */
func (pq *MlePriorityQueue) IsEmpty() bool {
	return pq.m_array.heapEmpty()
}

// Capacity is used to identify the capacity of the MlePriorityQueue.
func (pq *MlePriorityQueue) Capacity() int {
	return pq.m_array.Capacity()
}
