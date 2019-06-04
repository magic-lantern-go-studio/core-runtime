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

/**
 * The Event class is the abstract root class from which all event state objects shall be derived.
 *
 * All Event's are constructed with a reference to the object, the "source", that is logically
 * deemed to be the object upon which the Event in question initially occurred upon. 
 */
type EventObject struct {
	/** The source of the event. */
	m_source *IObject
}

/**
 * The default constructor.
 */
func NewEventObject() *EventObject {
	p := new(EventObject)
	p.m_source = nil
	return p
}

/**
 * Constructs a prototypical Event.
 *
 * @param source The object that the Event occurred upon.
 */
func NewEventObjectWithObject(source *IObject) *EventObject {
	p := new(EventObject)
	p.m_source = source
	return p
}

/**
 * Get the source of the Event.
 *
 * @return The object that the Event initially occurred upon.
 */
func (event *EventObject) GetSource() *IObject {
	return event.m_source
}

/**
 * Implement IObject interface.
 */
func ToString() string {
	// ToDo: implement string.
	return ""
}