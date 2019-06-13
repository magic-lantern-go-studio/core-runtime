/*
 * @file IMleObject.go
 * Created on April 25, 2019. (msm@wizzerworks.com)
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
package core

// Import go packages.
import (
	"io"

	mle_util "github.com/mle/runtime/util"
)

// IMleObject is an interface used to define functionality that is shared across
// Magic Lantern Runtime classes.
//
// The IMleObject interface specifies a common way of setting and
// retrieving properties.
type IMleObject interface {
	// Inherit root Object interface.
	mle_util.IObject

	/**
	 * Get the value of the property with the specified name.
	 *
	 * @param name The name of the property as a <code>string</code>.
	 *
	 * @return The property value is returned as an <code>Object</code>.
	 *
	 * @throws MleRuntimeException This exception is thrown if the specified
	 * property can not be retrieved.
	 */
	GetProperty(name string) IMleProp

	/**
	 * Set the value of the property with the specified name.
	 *
	 * @param name The name of the property as a <code>string</code>.
	 * @param property The property to set.
	 *
	 * @throws MleRuntimeException This exception is thrown if the specified
	 * property can not be set.
	 */
	SetProperty(name string, property IMleProp)

	/**
	 * Set the value of the property with the specified name.
	 *
	 * @param name The name of the property as a <code>string</code>.
	 * @param length The length of each property, in bytes.
	 * @param nElements The number of elements in the array, each of size <b>length</b>.
	 * @param value The value of the property to set.
	 *
	 * @throws MleRuntimeException This exception is thrown if the specified
	 * property can not be set.
	 */
	SetPropertyArray(name string, length int, nElements int, value io.ByteReader)

	/**
	 * Report a bound property change. If <i>oldValue</i> and <i>newValue</i> are not
	 * equal and the <code>MlePropChangeEvent</code> listener list isn't empty,
	 * then fire a <code>MlePropChangeEvent</code> event to each listener.
	 *
	 * @param name The name of the property that was changed.
	 * @param oldProperty The old value of the property (as an Object).
	 * @param newProperty The new value of the property (as an Object).
	 *
	 * @throws MleRuntimeException This exception is thrown if registered property
	 * listeners can not be notified.
	 */
	NotifyPropertyChange(name string, oldProperty IMleProp, newProperty IMleProp)

	/**
	 * Adds a <code>IMlePropChangeListener</code> for a specific property.
	 * The listener will be invoked only when a call on <code>NotifyPropertyChange</code>
	 * names that specific property. If listener is <b>null</b>, no exception is thrown
	 * and no action is performed.
	 *
	 * @param name The name of the property to listen on.
	 * @param listener The <code>IMlePropChangeListener</code> to be added.
	 *
	 * @throws MleRuntimeException This exception is thrown if the property
	 * listener can not be added. It is also thrown if the <i>name</i> argument
	 * is <b>null</b>.
	 */
	AddPropertyChangeListener(name string, listener IMleListener) MleError

	/**
	 * Removes a <code>IMlePropChangeListener</code> for a specific property.
	 * If listener is <b>null</b>, no exception is thrown and no action is performed.
	 *
	 * @param name The name of the property that was listened on.
	 * @param listener The <code>IMlePropChangeListener</code> to be removed.
	 *
	 * @throws MleRuntimeException This exception is thrown if the property
	 * listener can not be removed. It is also thrown if the <i>name</i> argument
	 * is <b>null</b>.
	 */
	RemovePropertyChangeListener(name string, listener IMleListener) MleError
}
