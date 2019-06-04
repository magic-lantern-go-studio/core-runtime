/**
 * @file Observable.go
 * Created on May 29, 2019. (msm@wizzerworks.com)
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

// IObservable defines an interface for the Observable/Observer pattern.
type IObservable interface {
	AddObserver(observer IObserver)
	DeleteObserver(observer IObserver)
	NotifyObservers()
	NotifyObserversWithObject(arg IObject)
	ClearChanged()
	SetChanged()
	HasChanged() bool
	CountObservers() int
}

// Observable is a class that manages a list of observers.
type Observable struct {
	observers []IObserver
	changed   bool
}

/**
 * The default constructor.
 */
func NewObservable() *Observable {
	p := new(Observable)
	p.changed = false
	return p
}

/**
 * Adds an observer to the set of observers for this object,
 * provided that it is not the same as some observer already in the set.
 * The order in which notifications will be delivered to multiple observers
 * is not specified. See the class comment.
 *
 * @param observer An observer to be added.
 */
func (source *Observable) AddObserver(observer IObserver) {
	source.observers = append(source.observers, observer)
}

/**
 * Deletes an observer from the set of observers of this object.
 * Passing null to this method will have no effect.
 *
 * @param observer The observer to be deleted.
 */
func (source *Observable) DeleteObserver(observer IObserver) {
	var index = -1

	// Find the specified listener.
	for i := 0; i <= len(source.observers); i++ {
		next := (source.observers)[i]
		if next == observer {
			index = i
			break
		}
	}

	if index != -1 {
		// Found matching observers. Now remove it.
		copy(((*source).observers)[index:], ((*source).observers)[index+1:])
		((*source).observers)[len((*source).observers)-1] = nil // or the zero value of T
		(*source).observers = ((*source).observers)[:len((*source).observers)-1]
	}
}

/**
 * Clears the observer list so that this object no longer has any observers.
 */
func (source *Observable) DeleteObservers() {
	for i := 0; i <= len(source.observers); i++ {
		(source.observers)[i] = nil
	}
}

/**
 * If the object has changed, as indicated by the HasChanged method,
 * then notify all of its observers and then call the ClearChanged method
 * to indicate that this object has no longer changed.
 *
 * Each observer has its update method called with two arguments:
 * this observable object and nil. In other words, this method is equivalent to:
 *
 * NotifyObservers(nil)
 */
func (source *Observable) NotifyObservers() {
	for _, observer := range source.observers {
		if observer != nil {
			observer.Update(source, nil)
		}
	}

	source.ClearChanged()
}

/**
 * If this object has changed, as indicated by the HasChanged method,
 * then notify all of its observers and then call the ClearChanged method
 * to indicate that this object has no longer changed.
 *
 * Each observer has its Update method called with two arguments:
 * this observable object and the arg argument.
 *
 * @param arg any object.
 */
func (source *Observable) NotifyObserversWithObject(arg IObject) {
	for _, observer := range source.observers {
		if observer != nil {
			observer.Update(source, arg)
		}
	}

	source.ClearChanged()
}

/**
 *
 * Indicates that this object has no longer changed,
 * or that it has already notified all of its observers of its most recent change,
 * so that the HasChanged method will now return false. This method is called
 * automatically by the NotifyObservers methods.
 *
 * @see SetChanged(), HasChanged()
 */
func (source *Observable) ClearChanged() {
	source.changed = false
}

/**
 * Marks this Observable object as having been changed;
 * the HasChanged method will now return <b>true</b>.
 *
 * @see ClearChanged(), HasChanged()
 */
func (source *Observable) SetChanged() {
	source.changed = true
}

/**
 * Tests if this object has changed.
 *
 * @return <b>true</b> if and only if the SetChanged method has been called more
 * recently than the ClearChanged method on this object; <b>false</b> otherwise.
 *
 * @see SetChanged(), ClearChanged()
 */
func (source *Observable) HasChanged() bool {
	return source.changed
}

/**
 * Returns the number of observers of this Observable object.
 *
 * @return The number of observers of this object.
 */
func (source *Observable) CountObservers() int {
	return len(source.observers)
}
