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
	mle_core "github.com/mle/runtime/core"
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

// GTheEventManager is the singleton instance of the event manager.
var GTheEventManager *MleEventManager
    
// GOkToExit is a global flag indicating whether it is Ok to exit.
var GOkToExit bool

// EventSetItem is a utility class used to manage named events.
//
// Implements Comparable interface.
type EventSetItem struct {
    /** The event identifier. */
    m_id types.Int
    /** The named event. May be <b>nil</b>. */
    m_name string
}

// NewEventSetItem is a constructor that initializes the event id and name.
//
// Parameters
//   id - The event identifier.
//   name - The name of the event.
//
// Return
//   A reference to a new EventSetItem is returned.
func NewEventSetItem(id int, name string) *EventSetItem {
	p := new(EventSetItem)
	p.m_id = types.Int(id)
	p.m_name = name
	return p
}

// Equals will determine if the EventSetITem is equal to the specified object.
//
// Parameters
//   obj - The Object to test against.
//
// Return
//   true will be returned if the obj is equal to this EventSetITem.
//   Otherwise false will be returned.
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

// CompareTo implements the Comparable interface.
//
// Return
//   positive integer, if the current object is greater than the specified object.
//   negative integer, if the current object is less than the specified object.
//   zero, if the current object is equal to the specified object.
func (esi *EventSetItem) CompareTo(item EventSetItem) int {
	if esi.m_id < item.m_id {
		return -1
	} else if (esi.m_id > item.m_id) {
		return 1
	}
	return 0
}

// EventSet is a helper class managing registered events.
type EventSet struct {
	m_tree types.TreeMap
}

// NewEventSet is the default constructor.
func NewEventSet() *EventSet {
	p := new(EventSet)
	return p
}

// Add will put the specified item in the EventSet collection.
//
// Parameters
//  item - A reference to an EventSetItem to add
//
func (es *EventSet) Add(item *EventSetItem) {
    es.m_tree.Put(item.m_id, item)
}

// Remove will delete the specified item from the EventSet collection.
///
// Parameters
//  item - A reference to an EventSetItem to remove
//
func (es *EventSet) Remove(item *EventSetItem) {
	es.m_tree.Remove(item.m_id)
}

// Clear will delete the entire EventSet collection.
func (es *EventSet) Clear() {
	for k, next := es.m_tree.Keys()(); next != nil; k, next = next() {
		key := k
		es.m_tree.Remove(key)
	}
}

// Contains will determine if the EventSet collection has the specified item.
//
// Parameters
//  item - A reference to an EventSetItem
//
// Return
//   true will be returned if the specified item is in the EventSet
//   collection, Otherwise, false will be returned.
func (es *EventSet) Contains(item *EventSetItem) bool {
    return es.m_tree.Has(item.m_id)
}

// Size will get the size of the EventSet collection.
//
// Return
//   The size of the EventSet collection will be returned as an integer.
func (es *EventSet) Size() int {
	return es.m_tree.Size()
}

// FindByName will find the named EventSetItem in the EventSet collection.
//
// Parameters
//   name - The name of the EventSetItem
//
// Return
//   A reference to the found EventSetItem will be returned. If the named
//   item does not exist in the collection, then nil will be returned.
func (es *EventSet) FindByName(name string) *EventSetItem {
	var found *EventSetItem = nil

	// Iterate through the tree to find a matching EventSetItem.
    for k, v, next := es.m_tree.Iterate()(); next != nil; k, v, next = next() {
		_ = k
		var item = v.(*EventSetItem)
		if item.m_name == name {
			found = item
			break
		}
	}

	return found
}

/**
 * Find the registered <code>EventSetItem</code> for the specified event
 * identifier.
 * 
 * @param id The event identifier.
 * 
 * @return The registered <code>EventSetItem</code> will be returned.
 */
func (es *EventSet) Find(id int) *EventSetItem {
	var found *EventSetItem = nil

	// Iterate through the tree to find a matching EventSetItem.
    for k, v, next := es.m_tree.Iterate()(); next != nil; k, v, next = next() {
		_ = k
		var item = v.(*EventSetItem)
		if item.m_id == types.Int(id) {
			found = item
			break
		}
	}

	return found
}

/**
 * Get the last event for the specified event group.
 * 
 * @param group The group identifier for the event category.
 * 
 * @return The registered <code>EventSetItem</code> will be returned for
 * the last event in the specified group. <b>null</b> may be returned
 * if there is no registered event.
 */
func (es *EventSet) Last(group int16) *EventSetItem {
	var found *EventSetItem = nil
	var lastEid int16 = 0

	// Iterate through the tree to find a matching EventSetItem.
	for k, v, next := es.m_tree.Iterate()(); next != nil; k, v, next = next() {
		_ = k
		var item = v.(*EventSetItem)
		gid := GetGroupId(int(item.m_id))
		if (gid == group) {
			eid := GetEventId(int(item.m_id))
			if (eid >= lastEid) {
				lastEid = eid
				found = item
			}
		}
	}	
						 
	return found
}

/**
 * Get the first event for the specified event group.
 * 
 * @param group The group identifier for the event category.
 * 
 * @return The registered <code>EventSetItem</code> will be returned for
 * the first event in the specified group. <b>null</b> may be returned
 * if there is no registered event.
 */
func (es *EventSet) First(group int16) *EventSetItem {
	var found *EventSetItem = nil

	// Iterate through the tree to find a matching EventSetItem.
	for k, v, next := es.m_tree.Iterate()(); next != nil; k, v, next = next() {
		_ = k
		var item = v.(*EventSetItem)
		gid := GetGroupId(int(item.m_id))
		if (gid == group) {
			found = item
			break
		}
	}
	
	return found
}
  
// MleEventManager is a class used to help manage events. A registry is kept for bookkeeping,
// in order to generate unique event identifiers on-the-fly.
type MleEventManager struct {
    // The event registry.
    m_eventRegistry *EventSet
}

// NewMleEventManager is the default constructor that returns a singleton
// instance of the Event Manager.
func NewMleEventManager() *MleEventManager {
	if GTheEventManager == nil {
		p := new(MleEventManager)
		p.m_eventRegistry = NewEventSet()
		// Add default events.
		p.AddEvent(MLE_PAINT, "")
		p.AddEvent(MLE_SIZE, "")
		p.AddEvent(MLE_RESIZEPAINT, "")
		p.AddEvent(MLE_QUIT, "")
		GTheEventManager = p

	}
	return GTheEventManager
}

/**
 * Register an event with the event manager.
 * 
 * @param id The event identifier. Nothing will be done if the event has
 * previously been registered.
 * @param name The named event; must be unique or <b>null</b>.
 * 
 * @throws MleRuntimeException This exception is thrown if an event
 * already exists for the specified <i>name</i>.
 */
func (evm *MleEventManager) AddEvent(id int, name string)  *mle_core.MleError {
	item := NewEventSetItem(id, name)
 
	if (evm.m_eventRegistry.Contains(item)) {
		return nil
	}
	if (name != "") && (evm.m_eventRegistry.FindByName(name) != nil) {
		msg := "Named event already exists."
		err := mle_core.NewMleError(msg, 0, nil)
		return err
	}
 
	evm.m_eventRegistry.Add(item)

	return nil
}

/**
 * Determine if the event manager already contains the named event.
 * 
 * @param name The event name to test.
 * 
 * @return <b>true</b> will be returned if the event manager already
 * has the specified event registered. Otherwise, <b>false</b> will be
 * returned.
 */
func (evm *MleEventManager) HasEventByName(name string) bool {
	if evm.m_eventRegistry.FindByName(name) != nil {
		return true
	}
	return false
}

/**
 * Determine if the event manager already contains the specified event.
 * 
 * @param id The event identifier to test.
 * 
 * @return <b>true</b> will be returned if the event manager already
 * has the specified event registered. Otherwise, <b>false</b> will be
 * returned.
 */
func (evm *MleEventManager) HasEvent(id int) bool {
	item := NewEventSetItem(id, "");
	if (evm.m_eventRegistry.Contains(item)) {
		return true
	}
	return false;
}

/**
 * Get the event identifier for the named event.
 * 
 * @param name The named event.
 * 
 * @return If the named event is registered, then the value of the
 * event identifier will be returned. If the named event does not
 * exist, then <code>MleEvent.MLE_EVENT_INVALID_ID</code> will be
 * returned.
 */
 func (evm *MleEventManager) GetEventId(name string) int  {
	item := evm.m_eventRegistry.FindByName(name)
	if item != nil {
		return int(item.m_id)
	}
	return MLE_EVENT_INVALID_ID;
}

/**
 * Get the name for the specified event.
 * 
 * @param id The event identifier.
 * 
 * @return If the event is registered, then the name of the
 * event will be returned. This value may be <b>null</b> if
 * the event was not registered with a name.
 * If the event does not exist, then <b>""</b> will be
 * returned.
 */
func (evm *MleEventManager) GetEventName(id int) string {
	item:= evm.m_eventRegistry.Find(id)
	if item != nil {
				 return item.m_name
	}
	return ""
}

/**
 * Unregister an event from the event manager.
 * 
 * @param id The event identifier. Nothing happens if the registry does not
 * contain the specified event.
 */
 func (evm *MleEventManager) RemoveEvent(id int) {
	item := NewEventSetItem(id, "")
 
	if evm.m_eventRegistry.Contains(item) {
		event := evm.m_eventRegistry.Find(id)
		evm.m_eventRegistry.Remove(event)
	}
}
	 
/**
 * Clear the event manager registry of all events.
 */
func (evm *MleEventManager) Clear() {
	evm.m_eventRegistry.Clear()
}

/**
 * Get the number of events registered with the event manager.
 * 
 * @return The number of registered events is returned.
 */
func (evm *MleEventManager) Size() int {
	return evm.m_eventRegistry.Size()
}
	 
/**
 * Determine whether it is ok to exit.
 * 
 * @return <b>true</b> will be returned if it is Ok to exit.
 * Otherwise, <b>false</b> will be returned.
 */
func OkToExit() bool {
    return GOkToExit
}

/**
 * Set the exit state.
 * 
 * @param status This parameter should be <b>true</b> if it is
 * Ok to exit the application/title. If it is not Ok to exit,
 * then the status should be set to <b>false</b>.
 */
func SetExitStatus(status bool) {
	GOkToExit = status
}

/**
 * Create an event identifier for the specified group.
 * 
 * @param group The group identifier to create an event for.
 * 
 * @return A new event identifier will be returned. Note that the returned event
 * has not yet been added to the event manager (see addEvent(int)).
 * 
 * @see MleEvent.makeId(short, short)
 */
func CreateEvent(group int16) int {
	var id int
 
	lastEvent := NewMleEventManager().m_eventRegistry.Last(group)
	if lastEvent == nil {
		// First event in the group.
		id = MakeId(group, 0)
	} else {
		// Create the event identifier following the last one registered.
		eid := GetEventId(int(lastEvent.m_id))
		eid++;
		id = MakeId(group, eid)
	}
		 
	return id
}
