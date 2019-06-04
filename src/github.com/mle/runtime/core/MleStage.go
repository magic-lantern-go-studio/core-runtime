/**
 * @file MleStage.go
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
 * The reference to the stage, restricting each process to a single
 * stage.
 */
var g_theStage *MleStage

/**
 * <code>MleStage</code> is a class that manages the rendering
 * for a specific target platform.
 * <p>
 * This is the base class for all Magic Lantern stages.
 * </p><p>
 * Use Init() to initialize the stage. Use GetSize() to get the dimensions
 * of the stage.
 * </p>
 *
 * @author  Mark S. Millard
 * @version 1.0
 */
type MleStage struct{}

/**
 * The default constructor.
 */
func NewMleStage() *MleStage {
	p := new(MleStage)
	return p
}

/**
 * Get the Singleton instance of the Stage.
 *
 * @return A <code>MleStage</code> is returned.
 */
func GetMleStageInstance() *MleStage {
	// Set the global stage reference.
	if g_theStage == nil {
		g_theStage = NewMleStage()
	}
	return g_theStage
}

/**
 * Initialize the stage.
 * <p>
 * Typically, the stage may register itself with the scheduler.
 * The base init() function does nothing.
 * </p>
 *
 * @throws MleRuntimeException This exception is thrown if the
 * stage can not be successfully initialized.
 */
func (stage *MleStage) Init() {}

/**
 * Dispose all resources associated with the Stage.
 *
 * @throws MleRuntimeException This exception is thrown if the
 * stage can not be successfully disposed.
 */
func (stage *MleStage) Dispose() {}

// ToString implements IObject interface.
func (stage *MleStage) ToString() string {
	return ""
}

/**
 * Get the size of the stage.
 * <p>
 * Returns size of stage's window. Magic Lantern 1.0
 * supports one window per stage: this is the
 * default window.
 * </p>
 */
func (stage *MleStage) GetSize() *MleSize {
	return nil
}
