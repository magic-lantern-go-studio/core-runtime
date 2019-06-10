/**
 * @file MleEventManager.go
 * Created on June 10, 2019. (msm@wizzerworks.com)
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
	"github.com/timtadh/data-structures/types"
	mle_util "github.com/mle/runtime/util"
)

/** The event group reserved by the manager. */
const MLE_SYSTEM_GROUP int16 = 0

/** The paint event. */
var MLE_PAINT int = MakeId(MLE_SYSTEM_GROUP, 0)
/** The size/resize event. */
var MLE_SIZE int = MakeId(MLE_SYSTEM_GROUP, 1)
/** The resize paint event. */
var MLE_RESIZEPAINT int = MakeId(MLE_SYSTEM_GROUP, 2)
/** The exit event. */
var MLE_QUIT int = MakeId(MLE_SYSTEM_GROUP, 3)


/** The first composite event in the range of reserved events. */
var MLE_FIRST_EVENT int = MLE_PAINT
/** The last composite event in the range of reserved events. */
var MLE_LAST_EVENT int = MLE_QUIT

//  Constants for event dispatching priorities. Note that this ordering is
//  important. The Stage resize callback must be executed prior to the
//  Set resize callback because the Set uses the new off-screen buffer
//  which is reallocated by the Stage resize callback. Also, the Stage
//  resize paint callback is executed after all other resize events have
//  been processed. The resize paint callback redraws the world in the new
//  off-screen buffer.
        
/** The priority for dispatching the MLE_RESIZE event for a Stage. */
const MLE_RESIZE_STAGE_PRIORITY int = 10
/** The priority for dispatching the MLE_RESIZE event for a Set. */
const MLE_RESIZE_SET_PRIORITY int = 9
/** The priority for dispatching the MLE_RESIZEPAINT event for a Stage. */
const MLE_RESIZEPAINT_STAGE_PRIORITY int = -1

// The singleton instance of the event manager.
var g_theEventManager *MleEventManager
    
// Flag indicating whether it is Ok to exit.
var g_okToExit bool	

// EventSetItem is a utility class used to manage named events.
//
// Implements Comparable interface.
type EventSetItem struct {
    /** The event identifier. */
    m_id types.Int
    /** The named event. May be <b>nil</b>. */
    m_name string
}

func NewEventSetItem(id int, name string) *EventSetItem {
	p := new(EventSetItem)
	p.m_id = types.Int(id)
	p.m_name = name
	return p
}

func (esi *EventSetItem) Equals(obj mle_util.Object) bool {
	retValue := false
    if mle_util.InstanceOf(obj, (*EventSetItem)(nil)) {
        var item *EventSetItem = obj.(*EventSetItem)
        if item.m_id.Equals(esi.m_id) {
			retValue = true
		}
    }
    return retValue
}

func (esi *EventSetItem) CompareTo(item EventSetItem) int {
	if esi.m_id < item.m_id {
		return -1
	} else if (esi.m_id > item.m_id) {
		return 1
	}
	return 0
}

/**
 * This class is used to help manage events. A registry is kept for bookkeeping,
 * in order to generate unique event identifiers on-the-fly.
 * 
 * @author Mark S. Millard
 */

type MleEventManager struct {

}

// NewMleEventManager is the default constructor that returns a singleton
// instance of the Event Manager.
func NewMleEventManager() *MleEventManager {
	if g_theEventManager == nil {
		g_theEventManager = new(MleEventManager)
	}
	return g_theEventManager
}