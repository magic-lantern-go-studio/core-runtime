/*
 * @file MlePQElement.go
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
	"strconv"
)

/**
 * An element as stored in the heap tree or the result list.
 * The sorting key is implemented as an <b>int</b>.
 *
 * Implements IMleElement.
 */
type MlePQElement struct {
    /** The sorting key. */
    Key int
    /** The associated data. */
    Data Object
}

/**
 * The default constructor.
 */
func NewMlePQElement() *MlePQElement {
	p := new(MlePQElement)
	p.Data = nil
	return p
}

/**
 * A constructor that initializes the element.
 * 
 * @param key The sorting key.
 * @param data The element's associated data.
 */
func NewMlePQElementWithKey(key int, data Object) *MlePQElement {
	p := new(MlePQElement)
    p.Key = key
	p.Data = data
	return p
}

/**
 * Determine if the specified element <b>e</b> is greater
 * than <b>this</b> element.
 * 
 * @param e The element to test.
 * 
 * @return <b>true</b> is returned if <b>this</b> element is greater
 * than element <b>e</b>. Otherwise, <b>false</b> will be returned.
 */
func (pe *MlePQElement) IsGreaterThan(e IMleElement) bool {
	return pe.Key > e.(*MlePQElement).Key
}
   
/**
 * Get the string representation of the element.
 * 
 * @return A <code>String</code> is returned.
 */
func (pe *MlePQElement) ToString() string {
	v := int64(pe.Key)
	return strconv.FormatInt(v, 10)
}
