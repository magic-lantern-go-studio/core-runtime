/**
 * @file MleEventDispatcher.go
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
	hash_tbl   "github.com/timtadh/data-structures/hashtable"
	hash_types "github.com/timtadh/data-structures/types"
	mle_util   "github.com/mle/runtime/util"
	mle_core   "github.com/mle/runtime/core"
)

// Event callback node definition. Implements IMleCallbackId.
type _EventCBNode struct {
	/** The callback handler. */
	m_callback *IMleEventCallback
	/** The associated client data. */
	m_clientData *mle_util.Object
	/** Flag indicating whether event is enabled. */
	m_isEnabled bool
}

/**
 * The default constructor.
 */
func _NewEventCBNode() *_EventCBNode {
	p := new(_EventCBNode)
	p.m_callback = nil
	p.m_clientData = nil
	p.m_isEnabled = false
	return p
}

/**
 * Determine if the callback associated with this identifier is enabled or not.
 * 
 * @return <b>true</b> is returned if the callback is enabled. Otherwise,
 * <b>false</b> will be returned.
 */
func (cbnode *_EventCBNode) IsEnabled() bool {
    return cbnode.m_isEnabled
}
 
 /**
  * Get the callback associated with this identifier.
  * 
  * @return A reference to an <code>IMleCallback</code> is returned.
  */
func (cbnode *_EventCBNode) GetCallback() mle_core.IMleCallback {
    return *cbnode.m_callback
}

// Event node definition.
type _EventNode struct {
    // The event composite value.
    m_event int
    // A priority queue of callbacks.
    m_callbacks *mle_util.MlePQ
    // Flag indicating whether event is enabled.
    m_isEnabled bool
    // Next event node in a linked list.
    m_next *_EventNode
    // The owning group.
    m_group *_EventGroupNode
} // end of EventNode

// Construct a new Event Node.
func _NewEventNode() *_EventNode {
	p := new(_EventNode)
	p.m_event = MLE_EVENT_INVALID_ID
	p.m_callbacks = nil
	p.m_isEnabled = false
	p.m_next = nil
	p.m_group = nil
	return p
}

// Event group node definition.
type _EventGroupNode struct {
    // The first event node in a group.
	m_head *_EventNode
	// The last event node in a group.
	m_tail *_EventNode
}

// Construct a new Event Group Node.
func _NewEventGroupNode() *_EventGroupNode {
	p := new(_EventGroupNode)
	p.m_head = nil
	p.m_tail = nil
	return p
}

// Link in the specified node.
func (egn *_EventGroupNode) linkEventNode(node *_EventNode) (bool, *mle_core.MleError) {
    if node == nil {
		msg := "MleEventDispatcher: Node is nil."
		err := mle_core.NewMleError(msg, 0, nil)
		return false, err
	}
        
    if egn.m_head == nil {
        egn.m_head = node
        egn.m_tail = node
    } else {
        egn.m_tail.m_next = node;
        egn.m_tail = node
        node.m_next = nil
    }
        
    return true, nil
}

// Unlink the specified node.
func (egn *_EventGroupNode) unlinkEventNode(node *_EventNode) (bool, *mle_core.MleError) {
    var prevNode *_EventNode

    if node == nil {
		msg := "MleEventDispatcher: Node is nil."
		err := mle_core.NewMleError(msg, 0, nil)
		return false, err
	}

    // Find previous node.
    prevNode = egn.findPrevEventNode(egn.m_head, node)
    if (prevNode != nil) {
        return false, nil
	}

    // Unlink node from linked-list.
    if egn.m_head == node {
        // Node is first one in list.
        egn.m_head = node.m_next
        if egn.m_tail == node {
	        // Node is also last one in list.
			egn.m_tail = nil
		}
    } else {
        prevNode.m_next = node.m_next
        if egn.m_tail == node {
	        // Node is last one in list.
			egn.m_tail = prevNode
		}
    }

    return true, nil
}

// Find the event node in a linked list of nodes.
func (egn *_EventGroupNode) findEventNode(event int) *_EventNode {
	var node *_EventNode
	var found = false

	node = egn.m_head
	for node != nil	{
		if node.m_event == event {
			found = true
			break
		}
		node = node.m_next
	}

	if ! found {
		node = nil
	}
	return node
}

// Find the previous node to the one specified.
func (egn *_EventGroupNode) findPrevEventNode(firstNode *_EventNode, node *_EventNode) *_EventNode {
	var prevNode, nextNode *_EventNode
	var found = false

	prevNode = firstNode
	nextNode = firstNode
	for nextNode != nil {
		if nextNode == node	{
			found = true
			break
		} else {
			prevNode = nextNode
			nextNode = nextNode.m_next
		}
	}

	if ! found {
		return nil
	}
	return prevNode;
}

// A helper class to encapsulate a queue element.
type _EventQueueElement struct {
    /** The event. */
    m_event *MleEvent
}

/**
 * Construct an queue element to hold a <code>MleEvent</code>.
 *
 * @param event A Magic Lantern event.
 */
func _NewEventQueueElement(event *MleEvent) (*_EventQueueElement) {
	p := new(_EventQueueElement)
	p.m_event = event
	return p
}

/**
 * Get the Magic Lantern event associated with this queue element.
 *
 * @return A Magic Lantern event is returned.
 */
func (eqe *_EventQueueElement) _GetEvent() *MleEvent {
	return eqe.m_event
}

/**
 * <code>MleEventDispatcher</code> is used to synchronize the dispatching of
 * Magic Lantern runtime events.
 * 
 * @author Mark S. Millard
 */
type MleEventDispatcher struct {
    // The event callback nodes organized by groups.
    m_eventGroups *hash_tbl.Hash
    // The queued events.
    m_eventQueue *mle_util.MlePQ
    // Registry of event listeners.
    m_eventListeners *mle_util.Vector
}

/**
 * The default constructor.
 * <p>
 * The default number of groups is set to <b>10</b>.
 * </p>
 */
func NewMleEventDispatcher() *MleEventDispatcher {
	p := new(MleEventDispatcher)
	p.m_eventGroups = hash_tbl.NewHashTable(10)
	p.m_eventQueue = mle_util.NewMlePQWithSize(mle_util.MLE_INC_QSIZE)
	p.m_eventListeners = mle_util.NewVector()
	return p
}

/**
 * A constructor that initializes the capacity of the number
 * of event groups.
 * 
 * @param capacity The number of event groups.
 */
 func NewMleEventDispatcherWithCapacity(capacity int) *MleEventDispatcher {
	p := new(MleEventDispatcher)
	p.m_eventGroups = hash_tbl.NewHashTable(capacity)
	p.m_eventQueue = mle_util.NewMlePQWithSize(mle_util.MLE_INC_QSIZE)
	p.m_eventListeners = mle_util.NewVector()
	return p
}

// Find the event node based on the composite event id.
func (dispatcher *MleEventDispatcher) findEventNode(id int) *_EventNode {
	var node *_EventNode = nil
	var groupId int16 = GetGroupId(id)
	var key hash_types.Int16 = hash_types.Int16(groupId)
	 
	var value interface{}
	var err error
	value, err =  dispatcher.m_eventGroups.Get(key)
	_ = err
	var group *_EventGroupNode = value.(*_EventGroupNode)
	if group != nil {
		node = group.findEventNode(id)
	}
	
	return node
}

// Add a node. Note that this method does not check for duplicate entries.
func (dispatcher *MleEventDispatcher) addEventNode(node *_EventNode) {
	var groupId int16 = GetGroupId(node.m_event)
	var key hash_types.Int16 = hash_types.Int16(groupId)

	// Get the associated group.
	var value interface{}
	var err error
	value, err = dispatcher.m_eventGroups.Get(key)
	_ = err
    var group *_EventGroupNode = value.(*_EventGroupNode)
    if group == nil {
        // No group by this id, create a new one.
        group = _NewEventGroupNode()
        dispatcher.m_eventGroups.Put(key, group)
                
        node.m_group = group
    }
         
    group.linkEventNode(node)
}
     
// Remove the specified node.
func (dispatcher *MleEventDispatcher) removeEventNode(node *_EventNode) {
	var groupId int16 = GetGroupId(node.m_event)
	var key hash_types.Int16 = hash_types.Int16(groupId)

	// Get the associated group.
	var value interface{}
	var err error
	value, err = dispatcher.m_eventGroups.Get(key)
	_ = err
	var group *_EventGroupNode = value.(*_EventGroupNode)
    if group != nil {
        group.unlinkEventNode(node)
    }         
}

// Find the event callback node.
func (dispatcher *MleEventDispatcher) findEventCBNode(node *_EventNode, id *_EventCBNode) int {
	var result int = -1

	for i := 0; i < node.m_callbacks.GetNumElements(); i++ {
		item := node.m_callbacks.Peek(i)
		if item.Data == id {
			result = i
			break
		}
	}

	return result
}

/**
 * Install a callback for the specified event.
 * 
 * @param event The composite event identifier.
 * @param callback The callback to install.
 * @param clientData Client data associated with the dispatch
 * of the callback.
 * 
 * @return A callback identifier is returned. This may be used to
 * uniquely identify a specific callback for a particular event.
 * 
 * @throws MleRuntimeException This exception is thrown if the
 * callback can not be installed successfully.
 */
func (dispatcher *MleEventDispatcher) InstallEventCB(event int, callback IMleEventCallback, clientData *mle_util.Object) (mle_core.IMleCallbackId, *mle_core.MleError) {
    var node *_EventNode
 // Check if event node already exists.
	node = dispatcher.findEventNode(event)
    if node == nil {
	    // It doesn't, so create a new one.
	    node = _NewEventNode()
	    if node != nil {
		    // Initialize the node.
		    node.m_event = event;
		    node.m_callbacks = mle_util.NewMlePQWithSize(mle_util.MLE_INC_QSIZE)
		    node.m_isEnabled = true
			
		    dispatcher.addEventNode(node)
	    } else {
		    var msg string = "MleEventDispatcher: Unable to install event callback."
		    err := mle_core.NewMleError(msg, 0, nil)
		    return nil, err
	    }
    }

    // Install callback.
    cbNode := _NewEventCBNode();
    if cbNode != nil {
	    cbNode.m_callback = &callback
	    cbNode.m_clientData = clientData
	    cbNode.m_isEnabled = true

	    // Add callback node to priority queue.
	    item := mle_util.NewMlePQElementWithKey(0, cbNode)
	    node.m_callbacks.Insert(item)
    } else  {
	    var msg string = "MleEventDispatcher: Unable to install event callback."
	    err := mle_core.NewMleError(msg, 0, nil)
	    return nil, err
    }

    return cbNode, nil
}

/**
 * Uninstall the specified callback associated with the given
 * <b>event</b>.
 * 
 * @param event The composite event identifier.
 * @param id The identifier for the callback to uninstall.
 * 
 * @return <b>true</b> is returned if the callback is successfully uninstalled.
 * Otherwise, <b>false</b> will be returned.
 */
func (dispatcher *MleEventDispatcher) UninstallEventCB(event int, id mle_core.IMleCallbackId) bool {
	var node *_EventNode
 
	// Find event node.
	node = dispatcher.findEventNode(event)
	if node == nil {
		return false
	} else {
		var index int
 
		// Find callback node.
		index = dispatcher.findEventCBNode(node,id.(*_EventCBNode))
		if index == -1 {
			return false
		} else {
			// Destroy priority queue item.
			node.m_callbacks.DestroyItem(index)
		}
	}
 
	return true
}

/**
 * Enable the specified callback associated with the given
 * <b>event</b>.
 * 
 * @param event The composite event identifier.
 * @param id The identifier for the callback to enable.
 * 
 * @return <b>true</b> is returned if the callback is successfully enabled.
 * Otherwise, <b>false</b> will be returned.
 */
func (dispatcher *MleEventDispatcher) EnableEventCB(event int, id mle_core.IMleCallbackId) bool {
	var node *_EventNode
	var cbNode *_EventCBNode
 
	// Find event node.
	node = dispatcher.findEventNode(event)
	if node == nil {
		return false
	} else {
		// Find callback node.
		cbNode = id.(*_EventCBNode)
		if cbNode == nil {
			return false
		} else {
			cbNode.m_isEnabled = true
		}
	}
 
	return true
}

/**
     * Disable the specified callback associated with the given
     * <b>event</b>.
     * 
     * @param event The composite event identifier.
     * @param id The identifier for the callback to disable.
     * 
     * @return <b>true</b> is returned if the callback is successfully enabled.
     * Otherwise, <b>false</b> will be returned.
     */
func (dispatcher *MleEventDispatcher) DisableEventCB(event int, id mle_core.IMleCallbackId) bool {
	var node *_EventNode
	var cbNode *_EventCBNode
 
	// Find event node.
	node = dispatcher.findEventNode(event)
	if node == nil {
		return false;
	} else {
		// Find callback node.
		cbNode = id.(*_EventCBNode)
		if cbNode == nil {
			return false
		} else {
			cbNode.m_isEnabled = false
		}
	}
 
	return true
}
 
/**
 * Uninstall all callbacks associated with the specified event.
 *
 * @param event The composite event identifier.
 * 
 * @return <b>true</b> is returned if the event is successfully uninstalled.
 * Otherwise, <b>false</b> will be returned indicating that the specified
 * event is invalid or could not be uninstalled.
 * 
 * @throws MleRuntimeException This exception is thrown if the callbacks
 * can not be removed from the priority queue.
 */
 func (dispatcher *MleEventDispatcher) UninstallEvent(event int) (bool, *mle_core.MleError) {
    var node *_EventNode

	// Check if event already exists.
	node = dispatcher.findEventNode(event)
    if node == nil {
        return false, nil
    } else {
        // Destroy callback nodes.
        if node.m_callbacks != nil {
            numCallbacks := node.m_callbacks.GetNumElements()
            for i := 0; i < numCallbacks; i++ {
                var item *mle_util.MlePQElement = node.m_callbacks.Remove()
                item.Data = nil
            }
            node.m_callbacks = nil
        }

        // Free node.
        dispatcher.removeEventNode(node)
    }	  
	  
    return true, nil	  
}	  
	  
/**	  
 * Enable the specified event f	  or dispatching.
 *
 * @param event The composite event identifier.
 * 
 * @return <b>true</b> is returned if the event is succussfull enabled.
 * Otherwise, <b>false</b> will be returned.
 */
func (dispatcher *MleEventDispatcher) EnableEvent(event int) bool {
	var node *_EventNode
 
	// Check if event already exists.
	node = dispatcher.findEventNode(event)
	if node == nil {
		return false	
	} else {
		node.m_isEnabled = true
	}
 
	return true
}

/**
 * Disable the specified event for dispatching.
 *
 * @param event The composite event identifier.
 * 
 * @return <b>true</b> is returned if the event is succussfull disabled.
 * Otherwise, <b>false</b> will be returned.
 */
func (dispatcher *MleEventDispatcher) DisableEvent(event int) bool {
	var node *_EventNode
 
	// Check if event already exists.
	node = dispatcher.findEventNode(event)
	if node == nil {
		return false	
	} else {
		node.m_isEnabled = false
	}
 
	return true
}

/**
 * Change the priority of the callback for the specified event.
 * 
 * @param event The composite event identifier.
 * @param id The callback identifier.
 * @param key The new priority.
 * 
 * @return If the callback priority is successfully changed, then
 * <b>true</b> will be returned. Otherwise, <b>false</b> will be
 * returned.
 */
func (dispatcher *MleEventDispatcher) changeCBPriority(event int, id mle_core.IMleCallbackId, key int) bool {
	var node *_EventNode
	var result bool
 
	// Find event node.
	node = dispatcher.findEventNode(event)
	if node == nil {
		result = false
	} else {
		var index int
 
		// Find callback node.
		index = dispatcher.findEventCBNode(node, id.(*_EventCBNode))
		if index == -1 {
			result = false
		} else {
			result = node.m_callbacks.ChangeItem(index, key)
		}
	}
 
	return result
}

/**
 * Change the dispatch priority for the specified event.
 * 
 * @param event The composite event identifier.
 * @param key The new priority.
 * 
 * @return If the event priority is successfully changed, then
 * <b>true</b> will be returned. Otherwise, <b>false</b> will be
 * returned.
 */
func (dispatcher *MleEventDispatcher) ChangeEventPriority(event int, key int) bool {
	var result = false
		 
	for i := 0; i < dispatcher.m_eventQueue.GetNumElements(); i++ {
		element := dispatcher.m_eventQueue.GetElementAt(i)
		if event == element.Data.(MleEvent).m_id {
			result = dispatcher.m_eventQueue.ChangeItem(event, key)
			break;
		}
	}
 
	return result
}

/**
 * Dispatch all events that have been placed in the delayed queue.
 */
 func (dispatcher *MleEventDispatcher) DispatchEvents() {
    // Retrieve the size of the queue.
    size := dispatcher.m_eventQueue.GetNumElements();
        
    // Dispatch all events currently on the queue.
    for i := 0; i < size; i++ {
		var element *_EventQueueElement
		element = dispatcher.PopEvent()
        if element != nil {
            event := element._GetEvent()
                
            // Find the event node is our registry.
            node := dispatcher.findEventNode(event.GetId())
            if (node != nil) && (node.m_isEnabled) {
                // Execute each callback that has been installed for this event
                // a priori.
                    
                // Copy the queue into one we can process.
                processQ := mle_util.NewMlePQWithElements(node.m_callbacks.CopyQueue())
                for ! processQ.IsEmpty() {
                    item := processQ.Remove()
                    cbNode := item.Data.(*_EventCBNode)
                    if (cbNode != nil) && (cbNode.m_isEnabled) {
						// Invoke callback.
						cb := cbNode.m_callback
                        (*cb).Dispatch(*event, cbNode.m_clientData)
                    }
                }
    
                // Notify listeners.
                for j := 0; j < len(*dispatcher.m_eventListeners); j++ {
					listener := dispatcher.m_eventListeners.ElementAt(j).(IMleEventListener)
					listener.EventDispatched(event)
				}
            }
        }
    }
}

/**
 * Process the event specified by the event id.
 * <p>
 * If the event type associated with the id is MLE_EVENT_IMMDIATE,
 * then the event will dispatched immediately. If the event type
 * is MLE_EVENT_DELAYED, then the event will be placed on a priority
 * queue to be processed at a later date (when dispatchEvents() is
 * called).
 * </p><p>
 * The event will be dispatched with the specified priority if the event
 * type is MLE_EVENT_DELAYED.
 * </p>
 *
 * @param id The composite event identifier.
 * @param calldata An Object containing the data to
 * be processed along with the event.
 * @param type The type of dispatching to use.
 * @param priority The event dispatch priority.
 *
 * @return If the event is successfully processed, then
 * <b>true</b> will be returned. Otherwise, <b>false</b> will be
 * returned.
 */
func (dispatcher *MleEventDispatcher) ProcessEventWithPriority(id int, calldata *mle_util.Object, evType int16, priority int) bool {
	var status = false
	//Todo: Figure out how to make the dispatcher the source of the event.
	// i.e. var source *mle_util.Object = dispatcher.(mle_util.Object)
	var source *mle_util.Object
    var event = NewMleEventWithIdEvTypeCalldata(source, id, evType, calldata)
        
    if event != nil {
        if (evType == MLE_EVENT_IMMEDIATE) {
            // Dispatch event immediately.
            node := dispatcher.findEventNode(id)
            if (node != nil) && (node.m_isEnabled) {
                // Execute each callback that has been installed for this event
                // a priori.

                // Copy the queue into one we can process.
                processQ := mle_util.NewMlePQWithElements(node.m_callbacks.CopyQueue())
                for ! processQ.IsEmpty() {
                    var item *mle_util.MlePQElement = processQ.Remove()
                    var cbNode = item.Data.(*_EventCBNode)
                    if (cbNode != nil) && (cbNode.IsEnabled()) {
                        cb := cbNode.m_callback
                        status = (*cb).Dispatch(*event, cbNode.m_clientData)
                    }
                }
                    
                // Notify listeners.
                for i := 0; i < len(*dispatcher.m_eventListeners); i++ {
					listener := dispatcher.m_eventListeners.ElementAt(i).(IMleEventListener)
					listener.EventProcessed(event)
				}
            }
        } else if (evType == MLE_EVENT_DELAYED) {
            /* Push event onto delayed queue. */
            status = dispatcher.PushEvent(event, calldata, priority)
        }
    }
        
	return status
}

/**
 * Process the event specified by the event id.
 * <p>
 * If the event type associated with the id is MLE_EVENT_IMMEDIATE,
 * then the event will dispatched immediately. If the event type
 * is MLE_EVENT_DELAYED, then the event will be placed on a priority
 * queue to be processed at a later date (when DispatchEvents() is
 * called).
 * </p><p>
 * The event will be dispatched with the default priority if the event
 * type is MLE_EVENT_DELAYED.
 * </p>
 *
 * @param id The composite event identifier.
 * @param calldata An Object containing the data to
 * be processed along with the event.
 * @param type The type of dispatching to use.
 *
 * @return If the event is successfully processed, then
 * <b>true</b> will be returned. Otherwise, <b>false</b> will be
 * returned.
 */
 func (dispatcher *MleEventDispatcher) ProcessEvent(id int, calldata *mle_util.Object, evType int16) bool {
		 return dispatcher.ProcessEventWithPriority(id, calldata, evType, 0)
	 }
 
/*
 * Push the event onto the queue for delayed dispatching.
 *
 * @param event The event to be pushed onto the queue.
 * @param calldata The call data to be associated with the event.
 * @param key The event priority.
 *
 * @return If the event is successfully pushed onto the queue,
 * then <b>true</b> will be returned. Otherwise <b>false</b> will be
 * returned.
 */
func (dispatcher *MleEventDispatcher) PushEvent(event *MleEvent, calldata *mle_util.Object, priority int) bool {
	var queueElement = _NewEventQueueElement(event)
	var element = mle_util.NewMlePQElementWithKey(priority, queueElement)
	dispatcher.m_eventQueue.Insert(element)
	return true
}
	 
/*
 * Pop the event from the queue.
 *
 * @return If the event is successfully popped from the queue,
 * then an <code>EventQueueElement</code> will be returned.
 * Otherwise <b>null</b> will be returned.
 */
func (dispatcher *MleEventDispatcher) PopEvent() *_EventQueueElement {
	var element *_EventQueueElement = nil
		 
	// Check if the queue is empty.
	if ! dispatcher.m_eventQueue.IsEmpty() {
		/* There is at least one element in the queue. */
		queueElement := dispatcher.m_eventQueue.Remove()
		element = queueElement.Data.(*_EventQueueElement)
	}
		 
	return element
}

/**
 * Flush the delayed dispatch queue. All pending events will be
 * ignored.
 */
func (dispatcher *MleEventDispatcher) Flush() {
	dispatcher.m_eventQueue.Clear()
}

/**
 * Add an event listener.
 * 
 * @param listener The event listener to notify when events
 * are processed and/or dispatched.
 *
 */
func (dispatcher *MleEventDispatcher) AddListener(listener IMleEventListener) {
	if listener == nil {
		return
	}
	dispatcher.m_eventListeners.AddElement(listener)
}
	 
/**
 * Remove an event listener.
 * 
 * @param listener The event listener to remove.
 */
func (dispatcher *MleEventDispatcher) RemoveListener(listener IMleEventListener) {
	if listener == nil {
		return
	}
	dispatcher.m_eventListeners.RemoveElement(listener);
}

// Implement IObject interface.
func (dispatcher *MleEventDispatcher) ToString() string {
	// Todo: return something relavent for Dispatcher.
	return " "
}