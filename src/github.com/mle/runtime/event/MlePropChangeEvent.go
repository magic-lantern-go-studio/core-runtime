/**
 * @file MleEvent.go
 * Created on April 26, 2019. (msm@wizzerworks.com)
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
package event

// Import Magic Lantern packages.
import (
	mle_util "github.com/mle/runtime/util"
)

/**
 * A property change event.
 * <p>
 * This class is used by <code>IMleObject</code>s (i.e. <code>Actor</code>s and
 * <code>Set</code>s) to notify <code>IMlePropChangeListeners</code> of changes
 * made to their properties.
 * </p>
 *
 * @author Mark S. Millard
 */
type MlePropChangeEvent struct {
	/** Common event data. */
	m_event MleEvent
	/** The source that caused this event to be fired. */
	m_source *mle_util.Object

	/** The name of the property. */
	m_name string
	/** The old value of the property. */
	m_oldValue interface{}
	/** The new value of the property. */
	m_newValue interface{}
}

/**
 * Constructs a new <code>MlePropChangeEvent.</code>
 *
 * @param source The object that caused this event to be fired.
 * @param propertyName the name of the property that has changed.
 * @param oldValue The old value of the property.
 * @param newValue The new value of the property.
 */
func NewMlePropChangeEvent(source *mle_util.Object, propertyName string, oldValue interface{}, newValue interface{}) *MlePropChangeEvent {
	p := new(MlePropChangeEvent)
	p.m_source = source
	p.m_name = propertyName
	p.m_oldValue = oldValue
	p.m_newValue = newValue
	return p
}

/**
 * Gets the name of the property that was changed.
 *
 * @return The name of the property that was changed is returned.
 */
func (event *MlePropChangeEvent) GetPropertyName() string {
	return event.m_name
}

/**
 * Gets the old value for the property, expressed as an <code>Object</code>.
 *
 * @return The old value for the property, expressed as an <code>Object</code>
 * is returned. May be <b>null</b> if multiple values have changed, as in a
 * property array.
 */
func (event *MlePropChangeEvent) GetOldValue() interface{} {
	return event.m_oldValue
}

/**
 * Gets the new value for the property, expressed as an <code>Object</code>.
 *
 * @return The new value for the property, expressed as an <code>Object</code>
 * is returned. May be <b>null</b> if multiple values have changed, as in a
 * property array.
 */
func (event *MlePropChangeEvent) GetNewValue() interface{} {
	return event.m_newValue
}
