/**
 * @file MleEvent.go
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

// Import go packages.
import (
	"fmt"
	mle_util "github.com/mle/runtime/util"
)

/** Constant defining invalid id. */
const MLE_EVENT_INVALID_ID int = 0xffffffff

/**
 * <code>MleEvent</code> encapsulates a Magic Lantern event.
 *
 * @author  Mark S. Millard
 * @version 1.0
 */
type MleEvent struct {
	m_event *mle_util.EventObject

	/** The event id. */
	m_id int
	/** The event dispatch type. */
	m_type int16
	/** Call data assoicated with the event. */
	m_calldata mle_util.IObject
}

/**
 * Make a composite identifier based on the specified group
 * and event ids.
 *
 * @param group The group that the event belongs to.
 * @param id The unique identifier for the event.
 *
 * @return The composite event identifier is returned.
 */
func MakeId(group int16, id int16) int {
	var newId int = (int)((group << 16) | id)
	return newId
}

/**
 * Get the group id for the specified composite identifier.
 *
 * @param cid The composite event identifier.
 *
 * @return The group that the event belongs to is returned.
 */
func GetGroupId(cid int) int16 {
	var groupId int16 = (int16)(cid >> 16)
	return groupId
}

/**
 * Get the event id for the specified composite identifier.
 *
 * @param cid The composite event identifier.
 *
 * @return The event id is returned.
 */
func GetEventId(cid int) int16 {
	var eventId int16 = (int16)(cid & 0x00FF)
	return eventId
}

/**
 * Create a MleEvent for the specified object.
 *
 * @param source The object on which the Event initially occurred.
 */
func NewMleEvent(source mle_util.Object) *MleEvent {
	p := new(MleEvent)
	p.m_event = mle_util.NewEventObjectWithSource(source)
	p.m_id = MLE_EVENT_INVALID_ID
	p.m_type = MLE_EVENT_DELAYED
	p.m_calldata = nil
	return p
}

/**
 * Create a MleEvent for the specified object with the specified
 * identifier. By default the event will be dispatched in delayed
 * mode.
 *
 * @param source The object on which the Event initially occurred.
 * @param id The event id.
 */
func NewMleEventWithId(source mle_util.Object, id int) *MleEvent {
	p := new(MleEvent)
	p.m_event = mle_util.NewEventObjectWithSource(source)
	p.m_id = id
	p.m_type = MLE_EVENT_DELAYED
	p.m_calldata = nil
	return p
}

/**
 * Create a MleEvent for the specified object with the specified
 * identifier. Set the event dispatching type upon construction.
 *
 * @param source The object on which the Event initially occurred.
 * @param id The event id.
 * @param evType The event dispatching type. Valid types include:
 * <ul>
 *   <li>IMleEventCallback.MLE_EVENT_IMMEDIATE</li>
 *   <li>IMleEventCallback.MLE_EVENT_DELAYED</li>
 * </ul>
 */
func NewMleEventWithIdEvtype(source mle_util.Object, id int, evType int16) *MleEvent {
	p := new(MleEvent)
	p.m_event = mle_util.NewEventObjectWithSource(source)
	p.m_id = id
	p.m_type = evType
	p.m_calldata = nil
	return p
}

/**
 * Create a MleEvent for the specified object with the specified
 * identifier. By default the event will be dispatched in delayed
 * mode.
 *
 * @param source The object on which the Event initially occurred.
 * @param id The event id.
 * @param calldata The caller data associated with this event.
 */
func NewMleEventWithIdCalldata(source mle_util.Object, id int, calldata mle_util.IObject) *MleEvent {
	p := new(MleEvent)
	p.m_event = mle_util.NewEventObjectWithSource(source)
	p.m_id = id
	p.m_type = MLE_EVENT_DELAYED
	p.m_calldata = calldata
	return p
}

/**
 * Create a MleEvent for the specified object with the specified
 * identifier. Set the event dispatching type upon construction.
 *
 * @param source The object on which the Event initially occurred.
 * @param id The event id.
 * @param evType The event dispatching type. Valid types include:
 * <ul>
 *   <li>IMleEventCallback.MLE_EVENT_IMMEDIATE</li>
 *   <li>IMleEventCallback.MLE_EVENT_DELAYED</li>
 * </ul>
 * @param calldata The caller data associated with this event.
 */
func NewMleEventWithIdEvTypeCalldata(source mle_util.Object, id int, evType int16, calldata mle_util.IObject) *MleEvent {
	p := new(MleEvent)
	p.m_event = mle_util.NewEventObjectWithSource(source)
	p.m_id = id
	p.m_type = evType
	p.m_calldata = calldata
	return p
}

/**
 * Get the event id.
 *
 * @return An integer value representing the event id is returned.
 */
func (event *MleEvent) GetId() int {
	return event.m_id
}

/**
 * Get the event dispatching type.
 *
 * @return A short is returned identifying the type of event dispatching
 * used for this event. Valid types include:
 * <ul>
 *   <li>IMleEventCallback.MLE_EVENT_IMMEDIATE</li>
 *   <li>IMleEventCallback.MLE_EVENT_DELAYED</li>
 * </ul>
 */
func (event *MleEvent) GetType() int16 {
	return event.m_type
}

/**
 * Set the event dispatching type.
 *
 * @param evType A short identifying the type of event dispatching
 * to be used for this event. Valid types include:
 * <ul>
 *   <li>IMleEventCallback.MLE_EVENT_IMMEDIATE</li>
 *   <li>IMleEventCallback.MLE_EVENT_DELAYED</li>
 * </ul>
 */
func (event *MleEvent) SetType(evType int16) {
	event.m_type = evType
}

/**
 * Get the caller data associated with this event.
 *
 * @return An Object is returned encapsulating the client data.
 */
func (event *MleEvent) GetCallData() mle_util.IObject {
	return event.m_calldata
}

/**
 * Set the caller data for this event.
 *
 * @param calldata An Object containing the caller data.
 */
func (event *MleEvent) SetCallData(calldata mle_util.IObject) {
	event.m_calldata = calldata
}

// GetSource is used to obtain the source oject that is associated with
// generating the event.
//
// Return
//
//   An Object interface is returned
func (event *MleEvent) GetSource() mle_util.Object {
    return event.m_event.GetSource()
}

/**
 * Create a String representation of this MleEvent Object.
 *
 * @return A <code>string</code> representation is returned.
 */
func (event *MleEvent) String() string {
	var evType string

	if event.m_type == MLE_EVENT_IMMEDIATE {
		evType = "MLE_EVENT_IMMEDIATE"
	} else if event.m_type == MLE_EVENT_DELAYED {
		evType = "MLE_EVENT_DELAYED"
	} else {
		evType = "UNKNOWN"
	}

	str := fmt.Sprintf("%s%d%s%s", "MleEvent: id=", event.m_id, " : type=", evType)
	return str
}

func (event *MleEvent) ToString() string {
	return event.String()
}