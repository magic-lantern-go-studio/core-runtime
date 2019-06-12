/*
 * @file IMleEventCallback.go
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
	mle_util "github.com/mle/runtime/util"
	mle_core "github.com/mle/runtime/core"
)

/** The event is to be dispatched immediately, without delay. */
const MLE_EVENT_IMMEDIATE int16 = 0x0001

/** The event is to be dispatched at a later date, it is queued. */
const MLE_EVENT_DELAYED int16 = 0x0002

/**
 * This interface is used to declare a callback handler for a <code>MleEvent</code>.
 *
 * @author Mark S. Millard
 * @version 1.0
 */
type IMleEventCallback interface {
	/** Extend with IMleCallback interface. */
	mle_core.IMleCallback

	/**
	 * Dispatch the event using the specified event object and client data.
	 * <p>
	 * Callbacks should implement this method with functionality
	 * that executes the work to be done when the event is dispatched.
	 * </p>
	 *
	 * @param event The event object for the event being dispatched.
	 * @param clientdata The data associated with this callback by the client
	 * that registered the event with the event dispatcher.
	 *
	 * @return If the event is successfully dispatched, then <b>true</b>
	 * is returned. Otherwise, <b>false</b> should be returned.
	 */
	Dispatch(event MleEvent, clientdata mle_util.IObject) bool
}
