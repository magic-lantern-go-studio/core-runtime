/*
 * @file MleHeapArray.go
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

/**
 * This class contains a heap tree and result list.
 * It provides the following functionality:
 * <ul>
 * <li>specialized access methods to elements</li>
 * <li>test for emptiness of heaptree</li>
 * <li>swap</li>
 * </ul>
 */
type MleHeapArray struct {
	m_array *MleElementArray  // MleHeapArray extends MleElementArray extends

	// Heap tree root stored at index 0.
	m_rootIndex int
	m_nullIndex int
	// Points to last element in heap. Initialize for empty heap.
	m_lastHeap int
}

/**
 * The default constructor.
 */
func NewMleHeapArray() *MleHeapArray {
	p := new(MleHeapArray)
	p.m_array = NewMleElementArray()
	p.m_rootIndex = 0
	p.m_nullIndex = p.m_rootIndex - 1
	p.m_lastHeap = p.m_nullIndex
	return p
}

/**
 * A constructor that initializes with an empty array of length <b>size</b>.
 * 
 * @param size The initial capacity of the array.
 */
 func NewMleHeapArrayWithSize(size int) *MleHeapArray {
	p := new(MleHeapArray)
	p.m_array = NewMleElementArrayWithSize(size)
	p.m_rootIndex = 0
	p.m_nullIndex = p.m_rootIndex - 1
	p.m_lastHeap = p.m_nullIndex
	return p
}

/**
 * Add an element to the array.
 * 
 * @param n The element to add.
 */
func (ha *MleHeapArray) AddElement(n IMleElement) {
	ha.m_array.AddElement(n)
	ha.incrementHeapsize()
}

// For debug: return heapTree formatted as a tree.
func (ha *MleHeapArray) heapTreeString() string {
    ns := ha.m_rootIndex   // Index of start element in current level.
    ne := ha.leftIndex(ns) // Index of start element in next level.
    spa := 7 // Number of spaces before first element.
    spb := 1 // Number of spaces between elements.
      
	//StringBuffer s = new StringBuffer()
	var s bytes.Buffer
    for level := 0; level <= 3; level++ {
		//s.append(ha.space(spa)); // Space before first element.
		s.WriteString(ha.space(spa))
        for ni := ns; ni < ne; ni++ {  // ni: indices of elements in current level.
            // Get node's string representation; space if no node any more.
			//String node = at(ni) != m_nullIndex ? getElementAt(ni).toString() : " ";
			var node string
			if ha.at(ni) != ha.m_nullIndex {
                node = ha.GetElementAt(ni).ToString()
			} else {
			    node = " "
			}
			//s.append(node + ha.space(spb)); // append node and space between nodes.
			s.WriteString(node)
			s.WriteString(ha.space(spb))
        } 
        s.WriteString("\n")
        ns = ha.leftIndex(ns)
        ne = ha.leftIndex(ne)
        spb = spa
        spa = (spa+1) / 2 -1
    }
      
    return s.String()
}

// For debug: return a string of size spaces.
func (ha *MleHeapArray) space(size int) string {
	var buf bytes.Buffer
	for i := 0; i < size; i++ {
		buf.WriteString(" ")
	}
	return buf.String()
}

// For debug: left index of child element.
func (ha *MleHeapArray) leftIndex(k int) int {
    // Because root is at 0. Root at 1 gives 2*k.
    return 2 * k + 1;
}

/**
 * Utiltiy for determining validity of index.
 * 
 * @param index The index to test.
 * 
 * @return If index is in heap, return it unchanged, otherwise return nullIndex.
 */
func (ha *MleHeapArray) at(index int) int {
	//return index >= ha.m_rootIndex && index <= ha.m_lastHeap ? index : ha.m_nullIndex
	if index >= ha.m_rootIndex && index <= ha.m_lastHeap {
		return index
	}
	return ha.m_nullIndex
}

/**
 * Build the heap order for the whole tree.
 */
func (ha *MleHeapArray) buildHeap() {
    // Start at last non-leaf node. Process all nodes in reverse storage
	// order until we reach the root by this order, we ensure that heapify's
	// precondition is fulfilled.
	for k := ha.lastInnerElement(); k != ha.m_nullIndex; k = ha.predesessor(k) {
		ha.heapify(k)
	}
}

/**
 * Heapify the subtree located at index <b>k</b>.
 * <p>
 * Precondition: both of k's child trees are heap ordered.
 * </p>
 * 
 * @param k The index of the subtree.
 */
func (ha *MleHeapArray) heapify(k int) { 
	// Move the key down the tree till we're done.
	//while( k != m_nullIndex)
	for k != ha.m_nullIndex {
		k = ha.heapifyLocally(k)
	}    
}

/**
 * Heap order element k with respect to both its children.
 * 
 * @param k The index of the subtree.
 * 
 * @return If keys had to be swapped, return the element where k.m_key now is.
 * Otherwise, <m>null</m> will be returned.
 */
func (ha *MleHeapArray) heapifyLocally(k int) int {
	mc := ha.maxChild(k)
			
	if mc == ha.m_nullIndex {
		// k is leaf, we're done.
		return ha.m_nullIndex
	}
			
	if ha.IsGreaterThan(mc,k) {
		// If max child has bigger key then swap.
		ha.Swap(k,mc)
		return mc
	}
	return ha.m_nullIndex
}

/** 
 * Move the maximum node from the heap into the result list.
 * <p>
 * Precondition: the tree must be a heap.
 * </p>
 * 
 * @return Return index where node was placed in result list
 */
func (ha *MleHeapArray) moveMax() int {
	ha.Swap(ha.m_rootIndex,ha.m_lastHeap) // Move maximum element to result list.
	ha.decrementHeapsize()                // Heap is now one element smaller.
	ha.heapify(ha.m_rootIndex)            // Restore the heap.
	return ha.m_lastHeap + 1
}
		
/**
 * Determine if heap is empty.
 * 
 * @return <b>true</b> is returned if the heap is empty. Otherwise,
 * <b>false</b> is returned.
 */
func (ha *MleHeapArray) heapEmpty() bool {
	return ha.m_lastHeap == ha.m_nullIndex
}

/**
 * A node has been moved out of the heap. Register this by decrementing the heapsize counter.
 */
func (ha *MleHeapArray) decrementHeapsize() {
	ha.m_lastHeap--
}
 
/**
 * Increment the heap size.
 */
func (ha *MleHeapArray) incrementHeapsize() {
	ha.m_lastHeap++
}

/**
    * Get the last inner element in the tree.
    * 
    * @return Return last inner element (has no children) in heap,
    * nullIndex is returned if there is no such element. 
    */
func (ha *MleHeapArray) lastInnerElement() int {
	if ha.heapEmpty() {
		return ha.m_nullIndex
	}
	return ha.parent(ha.m_lastHeap)
}

/**
    * Get the predessor of the specified element.
    * 
    * @param k The index of the element to get the predessor of.
    * 
    * @return Return the predessor of element k in storage order,
    * nullIndex is returned if k is the root.
    */ 
func (ha *MleHeapArray) predesessor(k int) int {
	return ha.at(k - 1);
}
	
/**
 * Get the child on the left of the heap tree.
 * 
 * @param k The index of the parent element.
 * 
 * @return Return left child of element k, nullIndex if there is no such child.
 */
func (ha *MleHeapArray) leftChild(k int) int {
	// Because root is at 0. Root at 1 gives 2*k.
	return ha.at(2 * k + 1);
}

/**
 * Get the child on the right of the heap tree.
 * 
 * @param k The index of the parent element.
 * 
 * @return Return right child of element k, nullIndex if there is no such child.
 */
func (ha *MleHeapArray) rightChild(k int) int {
	// Because root is at 0. Root at 1 gives 2*k +1.
	return ha.at(2 * k + 2);
}
	
/** 
 * Get the parent element of the specified element as index <b>k</b>.
 * 
 * @param k The index of the element to get the parent of.
 * 
 * @return Return parent element of element k. Return nullIndex if k is the root.
 */
func (ha *MleHeapArray) parent(k int) int {
	// Because root is at 0. Root at 1 gives k/2.
	// parent(root) is now root. This makes ?: necessary
	//return k == m_rootIndex ? k == m_rootIndex : at((k-1) / 2)
	if k == ha.m_rootIndex {
		return ha.m_rootIndex
	}
	return ha.at((k-1) / 2)
}

/**
 * Get the child with the highest priority.
 * <p>
 * Precondition: heap must be complete.
 * </p>
 * 
 * @param k The element to obtain the child from.
 * 
 * @return Return the bigger child of element k. Return nullIndex if k is a leaf.
 */
func (ha *MleHeapArray) maxChild(k int) int {
	rc := ha.rightChild(k)
	lc := ha.leftChild(k)
	  
	if (rc == ha.m_nullIndex) {
		return lc
	}

	// Because heap is complete there must be a left child.
	if ha.IsGreaterThan(lc, rc) {
		return lc
	}
	return rc
}
 
/**
 * Clear the queue.
 */
func (ha *MleHeapArray) Clear() {
	if ha.heapEmpty() {
		return
	}
		 
	node := ha.GetElementAt(ha.moveMax());
	/*
	do {
		ha.decrementNumElements()
		if ha.heapEmpty() {
			break
		}
	} while ((node = ha.GetElementAt(ha.moveMax())) != null)
	*/
	for {
        ha.decrementNumElements()
		if ha.heapEmpty() {
			break
		}
		node = ha.GetElementAt(ha.moveMax())
        if node == nil {
            break
		}
	}
}

/**
 * Get the element at the specified index.
 * 
 * @param k The index of the element to get.
 * 
 * @return A reference to an <code>IMleElement</code> is returned.
 */
 func (ha *MleHeapArray) GetElementAt(k int) IMleElement {
    return ha.m_array.GetElementAt(k)
}

/**
 * Get the number of items in the queue.
 * 
 * @return The number of elements in the queue is returned.
 */
 func (ha *MleHeapArray) GetNumElements() int {
    return ha.m_array.GetNumElements()
}
	
/** 
 * Register that the last element has been removed.
 */
func (ha *MleHeapArray) decrementNumElements() {
    ha.m_array.DecrementNumElements()
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
func (ha *MleHeapArray) IsGreaterThan(a int, b int) bool {
	return ha.m_array.IsGreaterThan(a, b)
}

/**
  * Swap the elements at the indices a and b in the array.
  * 
  * @param a The first element index.
  * @param b The second element index.
  */
func (ha *MleHeapArray) Swap(a int, b int) {
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
func (ha *MleHeapArray) Grow(size int) bool {
	return ha.m_array.Grow(size)
}

// ToString implements the IObject interface.
func (ha *MleHeapArray) ToString() string {
	return ha.m_array.ToString()
}