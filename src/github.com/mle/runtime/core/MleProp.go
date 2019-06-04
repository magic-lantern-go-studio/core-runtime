/**
 * @file MleProp.go
 * Created on May 23, 2019. (msm@wizzerworks.com)
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
)

/**
 * This class implements a generic property with an unknown data
 * type (by default). MleProp implements the IMleProp interface.
 *
 * @author Mark S. Millard
 */
type MleProp struct {
	/** The property type. */
	m_type int
	/** The property length. */
	m_length int
	/** The data input stream. */
	m_stream io.ByteReader
}

/**
 * Default constructor.
 */
func NewMleProp() *MleProp {
	p := new(MleProp)
	p.m_type = PROP_TYPE_UNKNOWN
	p.m_length = 0
	p.m_stream = nil
	return p
}

/**
 * A constructor that initializes the data managed by the property.
 *
 * @param length The expected length of the data stream.
 * @param data The data input stream for the property.
 */
func NewMlePropWithLengthAndData(length int, data io.ByteReader) *MleProp {
	p := new(MleProp)
	p.m_type = PROP_TYPE_UNKNOWN
	p.m_length = length
	p.m_stream = data
	return p
}

/**
 * Set the property type.
 *
 * @param propType An integer representing the property type.
 */
func (prop *MleProp) SetType(propType int) {
	prop.m_type = propType
}

/**
 * Get the property type.
 * <p>
 * Valid types include:
 * <ul>
 * <li>PROP_TYPE_UNKNOWN</li>
 * <li>PROP_TYPE_MEDIAREF</li>
 * </ul>
 * </p>
 *
 * @return The property's type is returned.
 *
 * @see IMleProp.GetType()
 */
func (prop *MleProp) GetType() int {
	return prop.m_type
}

/**
 * Get the length of the property data.
 *
 * @return The size of the property data is returned.
 *
 * @see IMleProp.GetLength()
 */
func (prop *MleProp) GetLength() int {
	return prop.m_length
}

/**
 * Get the property data as an input stream.
 *
 * @return An input stream is returned.
 *
 * @see IMleProp.GetStream()
 */
func (prop *MleProp) GetStream() io.ByteReader {
	return prop.m_stream
}
