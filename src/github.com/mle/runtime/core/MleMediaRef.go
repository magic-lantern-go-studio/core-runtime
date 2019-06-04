/**
 * @file MleMediaRef.go
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
	"runtime"
)

/**
 * A container for managing media.
 */
type MleMediaRefBuffer struct {
	/** Flags associated with the buffer. */
	m_flags int32
	/** The size of the buffer. */
	m_bufferSize int
	/** The media reference buffer. */
	m_buffer []byte

	// The next buffer in a linked list.
	m_next *MleMediaRefBuffer
}

/**
 * The default constructor.
 */
func NewMleMediaRefBuffer() *MleMediaRefBuffer {
	p := new(MleMediaRefBuffer)
	p.m_bufferSize = 0
	p.m_buffer = nil
	p.m_next = nil
	return p
}

/**
 * <code>MleMediaRef</code> is a class that manages references to
 * media assets.
 * <p>
 * This is the base class for all Magic Lantern media references.
 * </p><p>
 * Use Init() to initialize the media reference.
 * </p>
 *
 * @see MleMediaRefBuffer
 * @see MleMediaRefConverter
 *
 * @author  Mark S. Millard
 * @version 1.0
 */
type MleMediaRef struct {
	/** The media references. */
	m_references *MleMediaRefBuffer
	/** The number of media references. */
	m_numReferences int
	/** The reference converter to change a buffer into a local representation. */
	m_converter *MleMediaRefConverter
}

/**
 * The default constructor.
 */
func NewMleMediaRef() *MleMediaRef {
	p := new(MleMediaRef)
	p.m_references = nil
	p.m_numReferences = 0
	// Create a default converter.
	p.m_converter = NewMleMediaRefConverter()
	return p
}

/**
 * Initialize the media reference.
 * <p>
 * This method is a hook to do any initialization of the media reference.
 * Typically, the media reference may schedule itself.
 * The base init() function does nothing.
 * </p>
 *
 * @throws MleRuntimeException This exception is thrown if the
 * media reference can not be successfully initialized.
 */
func (mediaref *MleMediaRef) Init() {}

/**
 * Dispose all resources associated with the MediaRef.
 *
 * @throws MleRuntimeException This exception is thrown if the
 * media reference can not be successfully disposed.
 */
func (mediaref *MleMediaRef) Dispose() {}

// ToString implements IObject interface.
func (mediaref *MleMediaRef) ToString() string {
	return ""
}

/**
 * Register the media for this reference.
 * <p>
 * This method is used to register the media data that is
 * referenced.
 * </p>
 */
func (mediaref *MleMediaRef) RegisterMedia(flags int32, size int, media []byte) bool {
	var newRef, nextRef *MleMediaRefBuffer
	var status bool = true

	// Allocate a new reference.
	newRef = NewMleMediaRefBuffer()
	newRef.m_flags = flags
	newRef.m_bufferSize = size
	newRef.m_buffer = make([]byte, size)
	copy(newRef.m_buffer, media)

	// Attach next reference.
	nextRef = mediaref.m_references
	if nextRef == nil {
		// First entry on the list.
		mediaref.m_references = newRef
	} else {
		// Add entry to the list.
		for nextRef.m_next != nil {
			nextRef = nextRef.m_next
		}
		nextRef.m_next = newRef
	}

	return status
}

/**
 * Clear the registry of media references.
 * <p>
 * This method is used to unload the media data from
 * the media registry.
 * </p>
 */
func (mediaref *MleMediaRef) ClearRegistry() {
	var nextRef, tmp *MleMediaRefBuffer

	nextRef = mediaref.m_references
	for i := 0; i < mediaref.m_numReferences; i++ {
		tmp = nextRef.m_next
		nextRef.m_buffer = nil
		nextRef = nil
		nextRef = tmp
	}
	mediaref.m_references = nil
	mediaref.m_numReferences = 0

	// Invoke garbage collection. Note that this blocks execution.
	runtime.GC()
}

/**
 * Get the next media reference relative to the specifed
 * load reference.
 *
 * @param loadReference The <code>MleMediaRefBuffer</code> which marks
 * the current media reference. May be <b>nil</b>.
 *
 * @return A <code>MleMediaRefBuffer</code> is returned.
 * If <i>loadReference</i> was specified as <b>nil</b>, then the first
 * <code>MleMediaRefBuffer</code> in the list will be returned.
 */
func (mediaref *MleMediaRef) GetNextMediaRef(loadReference *MleMediaRefBuffer) *MleMediaRefBuffer {
	if loadReference != nil {
		return loadReference.m_next
	} else {
		return mediaref.m_references
	}
}

/**
 * Get the media reference flags for the specified load
 * reference.
 *
 * @param loadReference The <code>MleMediaRefBuffer</code> which marks
 * the current media reference.
 *
 * @return The flags will be returned.
 *
 * @throws MleRuntimeException This exception is thrown if the specified
 * load reference is <b>nil</b>.
 */
func (mediaref *MleMediaRef) GetMediaRefFlags(loadReference *MleMediaRefBuffer) (int32, *MleError) {
	if loadReference == nil {
		msg := "Media Reference Buffer is nil."
		err := NewMleError(msg, 0, nil)
		return 0, err
	}

	return loadReference.m_flags, nil
}

/**
 * Get the media reference size for the specified load
 * reference.
 *
 * @param loadReference The <code>MleMediaRefBuffer</code> which marks
 * the current media reference.
 *
 * @return The size will be returned.
 *
 * @throws MleRuntimeException This exception is thrown if the specified
 * load reference is <b>nil</b>.
 */
func (mediaref *MleMediaRef) GetMediaRefBufferSize(loadReference *MleMediaRefBuffer) (int, *MleError) {
	if loadReference == nil {
		msg := "Media Reference Buffer is nil."
		err := NewMleError(msg, 0, nil)
		return 0, err
	}

	return loadReference.m_bufferSize, nil
}

/**
 * Get the media reference buffer for the specified load
 * reference.
 *
 * @param loadReference The <code>MleMediaRefBuffer</code> which marks
 * the current media reference.
 *
 * @return A byte array will be returned.
 *
 * @throws MleRuntimeException This exception is thrown if the specified
 * load reference is <b>null</b>.
 */
func (mediaref *MleMediaRef) GetMediaRefBuffer(loadReference *MleMediaRefBuffer) ([]byte, *MleError) {
	if loadReference == nil {
		msg := "Media Reference Buffer is nil."
		err := NewMleError(msg, 0, nil)
		return nil, err
	}

	return loadReference.m_buffer, nil
}

/**
 * Set the reference converter.
 *
 * @param converter The media reference converter to
 * set for this media reference.
 */
func (mediaref *MleMediaRef) SetMediaRefConverter(converter *MleMediaRefConverter) {
	mediaref.m_converter = converter
}

/**
 * Get the reference converter.
 *
 * @return A reference to a <code>MleMediaRefConverter</code> is returned.
 * May be <b>null</b>.
 */
func (mediaref *MleMediaRef) GetMediaRefConverter() *MleMediaRefConverter {
	return mediaref.m_converter
}

/**
 * Remove the reference converter.
 */
func (mediaref *MleMediaRef) eleteMediaRefConverter() {
	if mediaref.m_converter != nil {
		mediaref.m_converter.Dispose()
		mediaref.m_converter = nil
	}
}
