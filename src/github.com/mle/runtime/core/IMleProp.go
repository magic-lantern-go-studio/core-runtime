/*
 * @file IMleProp.go
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
import "io"

/** An unknown property type. */
const PROP_TYPE_UNKNOWN int = -1

/** The Media Reference property. */
const PROP_TYPE_MEDIAREF int = 10

/**
 * This interface identifies the contract for dealing with Magic Lantern
 * properties in a consistent manner.
 *
 * @author Mark S. Millard
 */
type IMleProp interface {
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
	 */
	GetType() int

	/**
	 * Get the length of the property data.
	 *
	 * @return The size of the property data is returned.
	 */
	GetLength() int

	/**
	 * Get the property data as an input stream.
	 *
	 * @return An input stream is returned.
	 */
	GetStream() io.ByteReader
}
