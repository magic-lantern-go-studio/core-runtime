/*
 * @file MlePQ.go
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

/** The maximum priority. */
const MLE_MAX_QPRIORITY int = 32767
/** The minimum priority. */
const MLE_MIN_QPRIORITY int = ^32767

/**
 * This class implements a priority queue.
 * 
 * Extends MlePriorityQueue.
 */
type MlePQ struct {
	m_queue *MlePriorityQueue 
}

/**
 * The default constructor.
 */
func NewMlePQ() *MlePQ {
	p := new(MlePQ)
	p.m_queue = NewMlePriorityQueue()
	return p
}

/**
 * Construct a priority queue from the specified elements.
 * 
 * @param elements The items to place in the queue.
 */
func NewMlePQWithElements(elements []MlePQElement) *MlePQ {
	p := new(MlePQ)
	p.m_queue = NewMlePriorityQueueWithSize(len(elements))
	for i := 0; i < len(elements); i++ {
		p.Insert(&elements[i])
	}
	return p
}

func NewMlePQWithSize(size int) *MlePQ {
	p := new(MlePQ)
	p.m_queue = NewMlePriorityQueueWithSize(size)
	return p
}

/**
 * Get the element at the specified index.
 * 
 * @param k The index of the element to get.
 * 
 * @return A reference to an <code>MlePQElement</code> is returned.
 */
func (pq *MlePQ) GetElementAt(k int) *MlePQElement {
	element := pq.m_queue.GetElementAt(k)
    return element.(*MlePQElement)
}

/**
 * Get the number of items in the queue.
 * 
 * @return The number of elements in the queue is returned.
 */
func (pq *MlePQ) GetNumElements() int {
    return pq.m_queue.GetNumElements()
}

// Capacity is used to identify the capacity of the MlePQ.
func (pq *MlePQ) Capacity() int {
	return pq.m_queue.Capacity()
}

/**
 * Insert a new item.
 * <p>
 * This method will grow the queue if necessary.
 * Once extended, the queue will never shrink back to its original size
 * (it remains extended).
 * </p>
 *
 * @param item The item to insert into the queue.
 */
func (pq *MlePQ) Insert(item *MlePQElement) {
	pq.m_queue.AddElement(item)
}

/**
 * Remove the highest priority item.
 *
 * @return If an item is successfully removed, then it will be
 * returned. Otherwise, <b>null</b> will be returned.
 */
func (pq *MlePQ) Remove() *MlePQElement {
	element := pq.m_queue.GetMaxElement()
	return element.(*MlePQElement)
}

/**
 * Remove all items with the specified priority.
 * 
 * @param priority The priority to match.
 * 
 * @return If any items are successfully removed, then they will be
 * returned in an array. Otherwise, <b>null</b> will be returned.
 */
func (pq *MlePQ) RemoveWithPriority(priority int) []MlePQElement {
	var k int
	var numFound int = 0
    var foundQ []MlePQElement

    // Allocate enough space for potential "hit" list.
    foundQ = make([]MlePQElement, pq.GetNumElements())
 
    // Find matching items.
    //while((k = findItem(priority)) != -1) {
	k = pq.FindItemWithPriority(priority)
	for k != -1 {
		foundQ[numFound] = *pq.GetElementAt(k)
		numFound++
		// Remove the item from the queue.
		pq.DestroyItem(k)
		// Find the next item of given priority.
		k = pq.FindItemWithPriority(priority)
    }
 
    if numFound > 0 {
	    result := make([]MlePQElement, numFound)
		//System.arraycopy(foundQ,0,result,0,numFound);
		copy(result,foundQ) // The number copied should be equal to numFound
	    return result
    }
	
	return nil
}

/**
 * Check whether the queue is empty.
 * 
 * @return <b>true</b> is returned if the queue is empty. Otherwise,
 * <b>false</b> will be returned.
 */
func (pq *MlePQ) IsEmpty() bool {
	return pq.m_queue.IsEmpty()
}

/**
 * Delete the specified item.
 * 
 * @param k The item to destroy.
 */
func (pq *MlePQ) DestroyItem(k int) {
	var prevKey int
 
	// Check to see if there are any items in the queue.
	if (k < 0) || (k > pq.GetNumElements()) || pq.IsEmpty() {
	    return
	} else if (pq.GetNumElements() == 1) {
	    pq.m_queue.m_array.decrementHeapsize()
	    pq.m_queue.m_array.decrementNumElements()
	    return
	}
 
	//prevKey = ((MlePQElement)getElementAt(k)).m_key
	prevKey = pq.GetElementAt(k).Key
	pq.m_queue.Swap(k, pq.GetNumElements() - 1)
	pq.m_queue.m_array.decrementHeapsize()

	if pq.GetElementAt(k).Key < prevKey {
		// if (((MlePQElement)getElementAt(k)).m_key < prevKey) flowDown(k)
		pq.m_queue.flowDown(k)
	} else if pq.GetElementAt(k).Key > prevKey {
		//else if (((MlePQElement)getElementAt(k)).m_key > prevKey) flowUp(k)
		pq.m_queue.flowUp(k)
	}
	pq.m_queue.m_array.decrementNumElements()
}

/**
 * Destroy the first element in the queue.
 */
func (pq *MlePQ) Destroy() {
	pq.DestroyItem(0)
}

/**
 * Destroy the elements in the queue with the specified priority.
 *
 * @param priority The priority of the elements to destroy.
 */
func (pq *MlePQ) DestroyItemWithPriority(priority int) {
	var k int

	// Delete all items of specified weight.
	//while((k = findItem(priority)) != -1)
	k = pq.FindItemWithPriority(priority)
	for k != -1 {
		pq.DestroyItem(k)
		k = pq.FindItemWithPriority(priority)
	}
}

/**
 * Change the priority of an item.
 * 
 * @param k The item to change.
 * @param priority The new priority.
 */
func (pq *MlePQ) ChangeItem(k int, priority int) bool {
	var result = true
	item := NewMlePQElementWithKey(priority, nil)
 
	// Check to see if there are any items in the queue.
	if (k < 0) || (k > pq.GetNumElements()) || pq.IsEmpty() {
		result = false
	} else {
		item.Key = priority
		item.Data = pq.GetElementAt(k).Data
		pq.DestroyItem(k)
		pq.Insert(item)
	}

	return result
}

/**
 * Find an item with specified priority.
 * 
 * @param priority The priority to look for.
 * 
 * @return The index to the found item is returned.
 */
func (pq *MlePQ) FindItemWithPriority(priority int) int {
	var notFound = true
	var i int = 0
 
	// Find first item with specified priority.
	//while((notFound) && (i < getNumElements()))
	for notFound && i < pq.GetNumElements() {
		if pq.GetElementAt(i).Key == priority {
			notFound = false
		} else {
			i++
		}
	}
 
	if notFound {
		return -1
	}
	return i
}

/**
 * Find an item matching the specified element.
 * 
 * @param item The element used to match.
 * 
 * @return The index to the found item is returned.
 */
func (pq *MlePQ) FindItem(item MlePQElement) int {
	var notFound bool = true
	var i int = 0
 
	// Find first item with specified priority.
	//while((notFound) && (i < getNumElements()))
	for notFound && i < pq.GetNumElements() {
		if (pq.GetElementAt(i).Key == item.Key) &&
		   (pq.GetElementAt(i).Data == item.Data) {
			notFound = false
		} else {
			i++;
		}
	}
 
	if notFound {
	    return -1	
	}
	return i
}

/**
 * Copy the elements in the queue.
 * 
 * @return An array of the copied elements is returned.
 */
func (pq *MlePQ) CopyQueue() []MlePQElement {
	var queue []MlePQElement 
		 
	if pq.GetNumElements() > 0 {
		//queue = new MlePQElement[getNumElements()]
		queue = make([]MlePQElement, pq.GetNumElements())
		for i := 0; i < pq.GetNumElements(); i++ {
            queue[i] = *pq.GetElementAt(i)
		}
	} else {
		queue = nil
	}
 
	return queue
}

/**
 * Determine if an element exits in the queue
 * with the specified priority.
 * 
 * @param priority The priority to check for.
 * 
 * @return <b>true</b> is returned if an element exists in
 * the queue with the specified priority. Otherwise <b>false</b>
 * will be returned.
 */
func (pq *MlePQ) InQueueWithPriority(priority int) bool {
	var result = true
 
	if pq.FindItemWithPriority(priority) == -1 {
		result = false
	}
 
	return result
}

/**
 * Determine if an element exits in the queue
 * matching the specified element.
 * 
 * @param item The element to check for.
 * 
 * @return <b>true</b> is returned if an element exists in
 * the queue matching the specified item. Otherwise <b>false</b>
 * will be returned.
 */
func (pq *MlePQ) InQueue(item MlePQElement) bool {
	var result = true
 
	if (pq.FindItem(item) == -1) {
		result = false
	}
 
	return result
}

/**
 * Peek into queue for specified item.
 * 
 * @param k The index of the element to check.
 * 
 * @return A reference to the item in the queue is returned
 * without actually removing it from the queue.
 */
func (pq *MlePQ) Peek(index int) *MlePQElement {
	return pq.m_queue.Peek(index).(*MlePQElement)
}

/**
 * Clear the queue.
 */
func (pq *MlePQ) Clear() {
	pq.m_queue.Clear()
}

/**
 * Join two priority queues into one larger one.
 * 
 * @param pq1 The first queue.
 * @param pq2 The second queue.
 * 
 * @return A new priority queue is returned containing elements
 * from queues <b>pq1</b> and <b>pq2</b>.
 */
func Join(pq1 MlePQ, pq2 MlePQ) *MlePQ {
	var numItems int
	var heap *MlePQ
	var tmpQ1, tmpQ2 []MlePQElement
 
	// Construct a new queue.
	numItems = pq1.GetNumElements() + pq2.GetNumElements()
	if numItems > 0 {
		heap = NewMlePQWithSize(numItems)
		tmpQ1 = pq1.CopyQueue()
		if tmpQ1 != nil {
			for i := 0; i < pq1.GetNumElements(); i++ {
				heap.Insert(&tmpQ1[i])
			}
		}
		
		tmpQ2 = pq2.CopyQueue()
		if tmpQ2 != nil {
			for i := 0; i < pq2.GetNumElements(); i++ {
				heap.Insert(&tmpQ2[i])
			}
		}
	} else {
		heap = nil
	}
 
	return heap
}
