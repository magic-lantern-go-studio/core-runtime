/*
 * @file MleEventCallback.go
 * Created on June 4, 2019. (msm@wizzerworks.com)
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
 * This class may be used to register a callback handler for a specific event.
 * 
 * @author Mark S. Millard
 * @version 1.0
 */
type MleEventCallback struct {
	m_enabled bool
}

/**
 * The default constructor.
 */
func NewMleEventCallback() *MleEventCallback {
	p := new(MleEventCallback)
	p.m_enabled = false
	return p
}

/**
 * The callback dispatch method. Extenders must implement this method
 * to provide functionality specific to the handler.
 * 
 * @param event The event that is being dispatched by the handler.
 * @param clientdata Client data registered with this handler.
 * 
 * @return If the event is successfully dispatched, then <b>true</b> should be
 * returned. Otherwise, <b>false</b> should be returned.
 * 
 * @see com.wizzer.mle.runtime.event.IMleEventCallback#dispatch(com.wizzer.mle.runtime.event.MleEvent, java.lang.Object)
 */
func Dispatch(event *MleEvent, clientdata *mle_util.Object) bool {
	// ToDo: Log something here. This method is supposed to be abstract (i.e. no implemented)
	return false
}
        
/**
 * Enable the callback handler.
 * 
 * @param enable <b>true</b> should be used if the callback is to be enabled. Otherwise,
 * <b>false</b> should be used to disable the callback.
 */
func (cb *MleEventCallback) Enable(enable bool) {
    cb.m_enabled = enable;
}
		 
/**
 * Determine whether the callback is enabled.
 * 
 * @return <b>true</b> is returned if the callback is enabled. Otherwise,
 * <b>false</b> will be returned.
 */
func (cb *MleEventCallback) IsEnabled() bool {
    return cb.m_enabled;
}
