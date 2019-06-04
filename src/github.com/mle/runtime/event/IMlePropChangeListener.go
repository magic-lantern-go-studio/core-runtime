/*
 * @file IMlePropChangeListener.go
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
package event

// Import Magic Lantern packages.
import (
	mle_core "github.com/mle/runtime/core"
)

/**
 * A "PropChange" event gets fired whenever an <code>Actor</code> changes a "bound" property.
 * You can register a <code>IMlePropChangeListener</code> with a source <code>Actor</code>
 * so as to be notified of any bound property updates.
 *
 * @author Mark S. Millard
 */
type IMlePropChangeListener interface {
	mle_core.IMleListener

	/**
	 * This method gets called when a bound property is changed.
	 *
	 * @param event A <code>MlePropChangeEvent</code> object describing the event
	 * source and the property that has changed.
	 */
	PropChangedEvent(event MlePropChangeEvent)
}
