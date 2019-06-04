/**
 * @file MleSize.go
 * Created on May 23, 2019. (msm@wizzerworks.com.com)
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

/**
 * This class is used as a container for a managed size element.
 *
 * @author Mark S. Millard
 */
type MleSize struct {
	// The width.
	m_width uint32
	// The height.
	m_height uint32
}

/**
 * The default constructor.
 *
 * @param width The width of the size element.
 * @param height The height of the size element.
 */
func NewMleSizeWithWidthAndHeight(width uint32, height uint32) *MleSize {
	p := new(MleSize)
	p.m_width = width
	p.m_height = height
	return p
}

/**
 * Get the width.
 *
 * @return The width is returned as a <b>long</b>.
 */
func (size *MleSize) GetWidth() uint32 {
	return size.m_width
}

/**
 * Get the height.
 *
 * @return The height is returned as a <b>long</b>.
 */
func (size *MleSize) GetHeight() uint32 {
	return size.m_height
}
