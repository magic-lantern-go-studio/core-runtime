/*
 * @file EventObject.go
 * Created on June 4, 2019. (msm@wizzerworks.com)
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

// EventObject is the abstract root class from which all event state objects shall be derived.
//
// All Event's are constructed with a reference to the object, the "source", that is logically
// deemed to be the object upon which the Event in question initially occurred upon. 
type EventObject struct {
	/** The source of the event. */
	m_source Object
}

// NewEventObject is the default constructor.
func NewEventObject() *EventObject {
	p := new(EventObject)
	p.m_source = nil
	return p
}

// NewEventObjectWithSource constructs a prototypical Event.
//
// Parameters
//   source - The object that the Event occurred upon.
func NewEventObjectWithSource(source Object) *EventObject {
	p := new(EventObject)
	p.m_source = source
	return p
}

// GetSource returns the source of the Event.
//
// Return
//   The object that the Event initially occurred upon. nil
//   may be returned.
func (event *EventObject) GetSource() Object {
	return event.m_source
}

// ToString implements the IObject interface.
//
// The ToString method on the Event source will be invoked, if it exists.
// If the method doesn't exist, then "" will be returned.
//
// Return
//   The value of the Event source ToString method will be returned.
//   If the Event source does not have a ToString method, then ""
//   will be returned.
func (event *EventObject) ToString() string {
    if event.m_source != nil {
	    if MethodExists(event.m_source, "ToString") {
		    value, _ := Invoke(event.m_source, "ToString")
			var str = value.Interface().(string)
		    return str
	    }
    }
	return ""
}