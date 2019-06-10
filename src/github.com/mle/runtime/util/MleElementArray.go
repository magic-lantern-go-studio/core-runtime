/*
 * @file MleElementArray.go
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

// Import go packages.
import (
	"bytes"
)

/** Defualt size queue grows by. */
const MLE_INC_QSIZE int = 64

/**
 * Base class of <code>MleHeapArray</code>.
 * <p>
 * Contains an array of elements that can be compared and swapped.
 * It provides the type of array used by classical sorting algorithms like heapsort.
 * Functionality includes:
 * <ul>
 * <li>specialized access methods to elements</li>
 * <li>compare elements</li>
 * <li>swap elements</li>
 * <li>add an element</li>
 * <li>String of element strings</li>
 * </ul>
 */
 type MleElementArray struct {
	// Array holds heap tree followed by the result list.
	m_array []IMleElement
	// Index of last node in array.
	m_lastElement int
 }

 /**
  * The default constructor.
  */
func NewMleElementArray() *MleElementArray {
	p := new(MleElementArray)
	p.m_array = make([]IMleElement, 0)
	p.m_lastElement = -1
	return p
}

/**
 * A constructor that initializes with an empty array of length <b>len</b>.
 * 
 * @param len The capacity of the array.
 */
func NewMleElementArrayWithSize(size int) *MleElementArray {
	p := new(MleElementArray)
	p.m_array = make([]IMleElement, size)
	p.m_lastElement = -1
	return p
}

/**
 * Get the array of elements as a string.
 * 
 * @return A <code>String</code> is returned.
 */
func (ea *MleElementArray) ToString() string {
	var buf bytes.Buffer

	buf.WriteString("( ")
	var size = ea.m_lastElement + 1
	for i := 0; i < size; i++ {
		buf.WriteString(ea.m_array[i].ToString())
		buf.WriteString(" ")
	}
	buf.WriteString(")")
	
	return buf.String()
}
 
/**
 * Add an element after the last element in the array.
 * 
 * @param n The element to add.
 */
func (ea *MleElementArray) AddElement(n IMleElement) {
    if (ea.m_lastElement == (len(ea.m_array) - 1)) {
	    ea.Grow(MLE_INC_QSIZE);
	}
	ea.m_lastElement++
	ea.m_array[ea.m_lastElement] = n;
}
	
/**
 * Get the element at the specified index.
 * 
 * @param k The index of the element to get.
 * 
 * @return A reference to an <code>IMleElement</code> is returned.
 */
func (ea *MleElementArray) GetElementAt(k int) IMleElement {
	var v IMleElement
	if k < 0 || k > ea.m_lastElement {
		v = nil
	} else {
        v = ea.m_array[k]
	}
    return v
}

/**
 * Get the number of items in the queue.
 * 
 * @return The number of elements in the queue is returned.
 */
func (ea *MleElementArray) GetNumElements() int {
    return ea.m_lastElement + 1;
}
	
/** 
 * Register that the last element has been removed.
 *
 * The last element in the array is dropped.
 */
func (ea *MleElementArray) DecrementNumElements() {
    ea.m_array[ea.m_lastElement] = nil
    ea.m_lastElement--
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
func (ea *MleElementArray) IsGreaterThan(a int, b int) bool {
	return ea.m_array[a].IsGreaterThan(ea.m_array[b])
}

/**
  * Swap the elements at the indices a and b in the array.
  * 
  * @param a The first element index.
  * @param b The second element index.
  */
func (ea *MleElementArray) Swap(a int, b int) {
	var t IMleElement = ea.m_array[a]
	ea.m_array[a] = ea.m_array[b]
	ea.m_array[b] = t
}

/**
 * Increase the size of the array.
 * 
 * @param size The size to increase the array by.
 * 
 * @return <b>true</b> is returned if the array grew successfully.
 * Otherwise <b>false</b> is returned.
 */
func (ea *MleElementArray) Grow(size int) bool {
	result := true
		 
	// Allocate space for the new array size.
	newSize := len(ea.m_array) + size
	newArray := make([]IMleElement, newSize)
		 
	if (newArray != nil) {
		if (ea.m_array != nil) {
			//System.arraycopy(m_array,0,newArray,0,m_array.length)
			num := copy(newArray, ea.m_array)
			_ = num  // Don't do anything yet with the number of elements copied.
			for i := 0; i < ea.m_lastElement; i++ {
				// Free up current elements in the array.
				ea.m_array[i] = nil
			}
			ea.m_array = nil
		}
		ea.m_array = newArray
	}
 
	return result
}
