/**
 * @file MleMediaRefConverter.go
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
	"reflect"

	mle_util "github.com/mle/runtime/util"
)

/**
 * <code>MleMediaRefConverter</code> is a class that determines a local
 * representation from a media reference, which may be a file, URL, or some
 * other type of identifier.
 * <p>
 * This is the base class for all Magic Lantern media reference converters.
 * It treats the reference as a local file name.
 * </p>
 *
 * @see MleMediaRef
 *
 * @author  Mark S. Millard
 * @version 1.0
 */
type MleMediaRefConverter struct {
	/** The local reference. */
	m_reference interface{}
	/** Flag indicating conversion is complete. */
	m_converted bool
}

/**
 * The default constructor.
 */
func NewMleMediaRefConverter() *MleMediaRefConverter {
	p := new(MleMediaRefConverter)
	p.m_reference = nil
	p.m_converted = false
	return p
}

/**
 * Set the media reference for this converter.
 *
 * @param reference The media reference to convert.
 */
func (converter *MleMediaRefConverter) SetReference(reference interface{}) {
	converter.m_reference = reference
}

/**
 * Get the media reference for this converter.
 *
 * @return The media reference is returned.
 */
func (converter *MleMediaRefConverter) GetReference() interface{} {
	return converter.m_reference
}

/**
 * Treat the media reference as a filename.
 *
 * @return The media reference is returned as a <code>string</code>.
 * <b>nil</b> will be returned if no reference has ever been set
 * for conversion.
 *
 * @throws MleRuntimeException This exception will be thrown if an
 * error occurs while attempting to get the filename.
 *
 * @see SetReference()
 */
func (converter *MleMediaRefConverter) GetFilename() (string, *MleError) {
	if converter.m_reference != nil {
		converter.m_converted = true
		rt := reflect.TypeOf(converter.m_reference)
		if rt.Kind() == reflect.Array {
			// ToDo: How do we know that this is an array of bytes?
			filename := string(converter.m_reference.([]byte))
			return filename, nil
		} else {
			// Expecting a generic IObject if not a byte array.
			filename := converter.m_reference.(mle_util.IObject).ToString()
			return filename, nil
		}
	}

	converter.m_converted = false
	// ToDo: create a MleError here and return it.
	return "", nil
}

/**
 * Determine if the conversion is complete.
 *
 * @return <b>true</b> will be returned if the conversion has completed
 * successfully. Otherwise, <b>false</b> will be returned.
 *
 * @throws MleRuntimeException This exception is thrown if an error
 * occured while converting the media reference data.
 */
func (converter *MleMediaRefConverter) ConversionComplete() bool {
	return converter.m_converted
}

/**
 * Dispose of converter resources.
 */
func (converter *MleMediaRefConverter) Dispose() {
	// Does nothing for now.
}
