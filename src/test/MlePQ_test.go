/**
 * @file MlePQ_test.go
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
package util_test

// Import go packages.
import (
	"strconv"
	"testing"

	mle_util "github.com/mle/runtime/util"
)

func TestNewPQ(t *testing.T) {
	// Construct an empty queue.
	q := mle_util.NewMlePQ()
	if q == nil {
		t.Errorf("TestNewPQ: NewMlePQ() returned nil")
	}
	if ! q.IsEmpty() {
		t.Errorf("TestNewPQ: NewMlePQ() length should be empty (0)")
	}
}

/**
 *  Test element insertion/removal.
 */
func TestInsertionRemoval(t *testing.T) {
	pq := mle_util.NewMlePQWithSize(10)
 
	t.Logf("TestInsertionRemoval: Simple Test\n")
		   
	pq.Insert(mle_util.NewMlePQElementWithKey(3,nil))
	var n3 mle_util.IMleElement = pq.Remove()
	t.Logf("TestInsertionRemoval: element n3 %s\n", n3.ToString())
	//TestCase.assertEquals(new Integer(3).toString(),n3.toString())
	if n3.ToString() != "3" {
		t.Errorf("TestInsertionRemoval: want element = 3, got %s", n3.ToString())
	}
		   
	pq.Insert(mle_util.NewMlePQElementWithKey(5,nil))
	pq.Insert(mle_util.NewMlePQElementWithKey(2,nil))
	var n5 *mle_util.MlePQElement = pq.Remove()
	t.Logf("TestInsertionRemoval: element n5 %s\n", n5.ToString())
	//TestCase.assertEquals(new Integer(5).toString(),n5.toString())
	if n5.ToString() != "5" {
		t.Errorf("TestInsertionRemoval: want element = 5, got %s", n5.ToString())
	}
	var n2 *mle_util.MlePQElement = pq.Remove()
	t.Logf("TestInsertionRemoval: element n2 %s\n", n2.ToString())
	//TestCase.assertEquals(new Integer(2).toString(),n2.toString())
	if n2.ToString() != "2" {
		t.Errorf("TestInsertionRemoval: want element = 2, got %s", n2.ToString())
	}
		   
	pq.Insert(mle_util.NewMlePQElementWithKey(4,nil))
	pq.Insert(mle_util.NewMlePQElementWithKey(7,nil))
	var n7 *mle_util.MlePQElement = pq.Remove()
	t.Logf("TestInsertionRemoval: element n7 %s\n", n7.ToString())
	var n4 *mle_util.MlePQElement = pq.Remove()
	t.Logf("TestInsertionRemoval: element n7 %s\n", n4.ToString())
	//TestCase.assertEquals(new Integer(7).toString(),n7.toString())
	if n7.ToString() != "7" {
		t.Errorf("TestInsertionRemoval: want element = 7, got %s", n7.ToString())
	}
	//TestCase.assertEquals(new Integer(4).toString(),n4.toString())
	if n4.ToString() != "4" {
		t.Errorf("TestInsertionRemoval: want element = 4, got %s", n4.ToString())
	}
}

/**
 *  Test queue growth.
 */
func TestQueueGrowth(t *testing.T) {
	pq := mle_util.NewMlePQWithSize(mle_util.MLE_INC_QSIZE)
 
	t.Logf("TestQueueGrowth: Queue Growth Test\n")
			 
	for i := 0; i < mle_util.MLE_INC_QSIZE; i++	{
		pq.Insert(mle_util.NewMlePQElementWithKey(i,nil))
	}
	pq.Insert(mle_util.NewMlePQElementWithKey(mle_util.MLE_INC_QSIZE,nil))
	cap := pq.Capacity()
	if cap != 128 {
		t.Errorf("TestQueueGrowth: want capacity = 128, got %d", cap)
	}
			 
	for i :=  mle_util.MLE_INC_QSIZE; i >= 0;  i-- {
		element := pq.Remove()
		t.Logf("TestQueueGrowth: element %d %s\n", i, element.ToString())
		//TestCase.assertEquals(new Integer(i).toString(),element.toString())
		if element.ToString() != strconv.Itoa(i) {
			t.Errorf("TestQueueGrowth: want element = %s, got %s", strconv.Itoa(i), element.ToString())
		}
	}
}

/**
 * Test the clear() method.
 */
func TestClear(t *testing.T) {
	pq := mle_util.NewMlePQWithSize(mle_util.MLE_INC_QSIZE)

 	t.Logf("TestClear: Clear Test\n")
			 
	for i := 0; i < mle_util.MLE_INC_QSIZE; i++ {
		pq.Insert(mle_util.NewMlePQElementWithKey(i,nil))
	}
			 
	pq.Clear()
			 
	//TestCase.assertEquals(0,pq.getNumElements());
	if pq.GetNumElements() != 0 {
		t.Errorf("TestClear: want number of elements = 0, got %d", pq.GetNumElements())
	}
}

/**
 * Test the remove() methods.
 */
func TestRemove(t *testing.T) {
	pq := mle_util.NewMlePQWithSize(10)

 	t.Logf("TestRemove: Remove Item Test\n")
			 
	for i := 0; i < 10; i++ {
		pq.Insert(mle_util.NewMlePQElementWithKey(i,nil))
	}
	pq.Insert(mle_util.NewMlePQElementWithKey(5,nil))
	pq.Insert(mle_util.NewMlePQElementWithKey(5,nil))
	pq.Insert(mle_util.NewMlePQElementWithKey(5,nil))
	pq.Insert(mle_util.NewMlePQElementWithKey(5,nil))
			 
	element := pq.Remove()
	//TestCase.assertEquals(new Integer(9).toString(),element.toString())
	if element.ToString() != "9" {
		t.Errorf("TestRemove: want element = 9, got %s", element.ToString())
	}
	
	// Removing all elements with priority 5.
	elements := pq.RemoveWithPriority(5)
	for i := 0; i < 5; i++ {
		//TestCase.assertEquals(new Integer(5).toString(),elements[i].toString())
		if elements[i].ToString() != "5" {
			t.Errorf("TestRemove: want element = 5, got %s", elements[i].ToString())
		}
	}
}

/**
 * Test the Join() method.
 */
func TestJoin(t *testing.T) {
	pq1 := mle_util.NewMlePQ()
	pq2 := mle_util.NewMlePQ()
	var result *mle_util.MlePQ
 
	t.Logf("TestJoin: Join Queues Test\n")
			 
	for i := 0; i < 10; i++ {
		pq1.Insert(mle_util.NewMlePQElementWithKey(i,nil))
		pq2.Insert(mle_util.NewMlePQElementWithKey(i+25,nil))
	}
			 
	result = mle_util.Join(*pq1, *pq2)
			 
	element := result.Remove()
	//TestCase.assertEquals(new Integer(34).toString(),element.toString());
	if element.ToString() != "34" {
		t.Errorf("TestJoin: want element = 34, got %s", element.ToString())
	}
}

/**
 * Test the changeItem() method.
 */
func TestChangeItem(t *testing.T) {
	pq := mle_util.NewMlePQWithSize(10)
			 
	t.Logf("TestJoin: Change Item Test\n")
			 
	for i := 0; i < 10; i++ {
		pq.Insert(mle_util.NewMlePQElementWithKey(i,nil))
	}
 
	// Crank priority to 56.
	pq.ChangeItem(pq.FindItemWithPriority(5),56)
			 
	element := pq.Remove()
	//TestCase.assertEquals(new Integer(56).toString(),element.toString());
	if element.ToString() != "56" {
		t.Errorf("TestJoin: want element = 56, got %s", element.ToString())
	}
}

/**
 * Test the Destroy() methods.
 */
func TestDestroy(t *testing.T) {
	pq := mle_util.NewMlePQWithSize(10)

	t.Logf("TestDestroy: Destroy Test\n")
			 
	for i := 0; i < 10; i++ {
		pq.Insert(mle_util.NewMlePQElementWithKey(i,nil))
	}
 
	// Destroy item with priority 5.
	pq.DestroyItemWithPriority(5)
			 
	k := pq.FindItemWithPriority(5)
	if k != -1 {
		t.Errorf("TestDestroy: want element = -1, got %d", k)
	}
			 
	// Destroy the top 3 priority items.
	pq.Destroy()
	pq.Destroy()
	pq.Destroy()
		 
	element := pq.Remove()
	//TestCase.assertEquals(new Integer(6).toString(),element.toString());
	if element.ToString() != "6" {
		t.Errorf("TestDestroy: want element = 6, got %s", element.ToString())
	}
}

/**
 * Test the destroyItem() method.
 */
func TestDestoryItem(t *testing.T) {
	pq := mle_util.NewMlePQWithSize(mle_util.MLE_INC_QSIZE)
 
	t.Logf("TestDestroyItem: Destroy Item Test\n")
			 
	for i := 0; i < mle_util.MLE_INC_QSIZE; i++ {
		pq.Insert(mle_util.NewMlePQElementWithKey(i,nil))
	}
			 
	pq.DestroyItem(pq.FindItemWithPriority(10))
	pq.DestroyItem(pq.FindItemWithPriority(20))
	pq.DestroyItem(pq.FindItemWithPriority(30))
	pq.DestroyItem(pq.FindItemWithPriority(40))
	//TestCase.assertEquals(60,pq.getNumElements());
	if pq.GetNumElements() != 60 {
		t.Errorf("TestDestroyItem: want number of elements = 60, got %d", pq.GetNumElements())
	}
			 
	k := pq.FindItemWithPriority(10)
	//TestCase.assertEquals(-1,k);
	if k != -1 {
		t.Errorf("TestDestroyItem: want element = -1, got %d", k)
	}
	k = pq.FindItemWithPriority(20);
	//TestCase.assertEquals(-1,k);
	if k != -1 {
		t.Errorf("TestDestroyItem: want element = -1, got %d", k)
	}
	k = pq.FindItemWithPriority(30);
	//TestCase.assertEquals(-1,k);
	if k != -1 {
		t.Errorf("TestDestroyItem: want element = -1, got %d", k)
	}
	k = pq.FindItemWithPriority(40);
	//TestCase.assertEquals(-1,k);
	if k != -1 {
		t.Errorf("TestDestroyItem: want element = -1, got %d", k)
	}
}
