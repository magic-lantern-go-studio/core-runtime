/**
 * @file MleEventManager_test.go
 * Created on June 13, 2019. (msm@wizzerworks.com)
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
package mle_test

import (
	"testing"

	mle_event "github.com/mle/runtime/event"
)

func TestNewMleEventManager(t *testing.T) {
	// Construct an empty queue.
	manager := mle_event.NewMleEventManager()
	if manager == nil {
		t.Errorf("TestNewMleEventManager: NewMleEventManager() returned nil")
	}
}

/**
 * Test adding events to the event manager.
 */
func TestAddEvent(t *testing.T) {
	manager := mle_event.NewMleEventManager()
	if manager == nil {
		t.Errorf("TestAddEvent: NewMleEventManager() returned nil")
	}

	// Create a new event in group 0.
    var event = mle_event.CreateEvent(0)
	manager.AddEvent(event, "")

	// Try to add it again, should not be able to.
	manager.AddEvent(event, "");
	if manager.Size() != 5 {
		// There should be four default events plus the one we just added.
		t.Errorf("TestAddEvent: EventManager size should be 5, got %d", manager.Size())
	}
}

func TestRemoveEvent(t *testing.T) {
	manager := mle_event.NewMleEventManager()
	if manager == nil {
		t.Errorf("TestAddEvent: NewMleEventManager() returned nil")
	}

	// Create a new event in group 0.
    var event = mle_event.CreateEvent(0)
	manager.AddEvent(event, "")

	// Create another new event in group 0.
	event = mle_event.CreateEvent(0)
	manager.AddEvent(event, "")

	if manager.Size() != 6 {
		// There should be four default events plus the two we just added.
		t.Errorf("TestAddEvent: EventManager size should be 6, got %d", manager.Size())
	}

	// Remove the last event.
	manager.RemoveEvent(event)
	if manager.Size() != 5 {
		// There should be one event less.
		t.Errorf("TestAddEvent: EventManager size should be 5, got %d", manager.Size())
	}
}

/**
 * Test creating new events.
 */
func TestCreateEvent(t *testing.T) {
	manager := mle_event.NewMleEventManager()
	if manager == nil {
		t.Errorf("TestCreateEvent: NewMleEventManager() returned nil")
	}

	var event = mle_event.CreateEvent(0)
	//assertEquals(event, 0);
	if event != 4 {
		t.Errorf("TestCreateEvent: expecting event 4, got %d", event)
	}
	manager.AddEvent(event, "")
	event = mle_event.CreateEvent(0)
	//assertEquals(event, 1);
	if event != 5 {
		t.Errorf("TestCreateEvent: expecting event 5, got %d", event)
	}
	manager.AddEvent(event, "")
	event = mle_event.CreateEvent(0)
	//assertEquals(event, 2);
	if event != 6 {
		t.Errorf("TestCreateEvent: expecting event 6, got %d", event)
	}
	manager.AddEvent(event, "")
	event = mle_event.CreateEvent(0)
	//assertEquals(event, 3);
	if event != 7 {
		t.Errorf("TestCreateEvent: expecting event 7, got %d", event)
	}
	manager.AddEvent(event, "")
	event = mle_event.CreateEvent(0)
	//assertEquals(event, 4);
	if event != 8 {
		t.Errorf("TestCreateEvent: expecting event 8, got %d", event)
	}
	manager.AddEvent(event, "")
	event = mle_event.CreateEvent(1)
	//assertEquals(event, 65536);
	if event != 65536 {
		t.Errorf("TestCreateEvent: expecting event 65536, got %d", event)
	}
	manager.AddEvent(event, "")
	event = mle_event.CreateEvent(1)
	//assertEquals(event, 65537);
	if event != 65537 {
		t.Errorf("TestCreateEvent: expecting event 65537, got %d", event)
	}
	manager.AddEvent(event, "")
	event = mle_event.CreateEvent(1)
	//assertEquals(event, 65538);
	if event != 65538 {
		t.Errorf("TestCreateEvent: expecting event 65538, got %d", event)
	}
	manager.AddEvent(event, "")
	event = mle_event.CreateEvent(1)
	//assertEquals(event, 65539);
	if event != 65539 {
		t.Errorf("TestCreateEvent: expecting event 65539, got %d", event)
	}
	manager.AddEvent(event, "")
	event = mle_event.CreateEvent(1)
	//assertEquals(event, 65540);
	if event != 65540 {
		t.Errorf("TestCreateEvent: expecting event 65540, got %d", event)
	}
	manager.AddEvent(event, "")
	event = mle_event.CreateEvent(2)
	//assertEquals(event, 131072);
	if event != 131072 {
		t.Errorf("TestCreateEvent: expecting event 131072, got %d", event)
	}
	event = mle_event.CreateEvent(0)
	//assertEquals(event, 5);
	if event != 9 {
		t.Errorf("TestCreateEvent: expecting event 9, got %d", event)
	}
	
	//assertEquals(MleEventManager.getInstance().size(), 10);
	if manager.Size() != 14 {
		// There should be four default events plus the 10 we just added.
		t.Errorf("TestCreateEvent: EventManager size should be 14, got %d", manager.Size())
	}
}

func TestHasEvent(t *testing.T) {
    manager := mle_event.NewMleEventManager()
	if manager == nil {
		t.Errorf("TestHasEvent: NewMleEventManager() returned nil")
	}

	var event = mle_event.CreateEvent(0)
	if event != 4 {
		t.Errorf("TestHasEvent: expecting event 4, got %d", event)
	}
    manager.AddEvent(event, "")
	event = mle_event.CreateEvent(0)
	if event != 5 {
		t.Errorf("TestHasEvent: expecting event 5, got %d", event)
	}
	manager.AddEvent(event, "")
	if !manager.HasEvent(4) {
		t.Errorf("TestHasEvent: expecting event true, got false")
	}
	if !manager.HasEvent(5) {
		t.Errorf("TestHasEvent: expecting event true, got false")
	}
	event = mle_event.CreateEvent(1)
	if event != 65536 {
		t.Errorf("TestHasEvent: expecting event 5, got %d", event)
	}
	manager.AddEvent(event, "")
	if !manager.HasEvent(65536) {
		t.Errorf("TestHasEvent: expecting event true, got false")
	}

	// Test for an event that is not registered.
	if manager.HasEvent(131072) {
		t.Errorf("TestHasEvent: expecting event false, got true")
	}
}

func TestHasEventWithName(t *testing.T) {
	manager := mle_event.NewMleEventManager()
	if manager == nil {
		t.Errorf("TestHasEventWithName: NewMleEventManager() returned nil")
	}
	
	var event = mle_event.CreateEvent(0)
	manager.AddEvent(event, "")
	event = mle_event.CreateEvent(0)
	manager.AddEvent(event, "One")
	if !manager.HasEvent(4) {
		t.Errorf("TestHasEvent: expecting event true, got false")
	}
	if !manager.HasEvent(5) {
		t.Errorf("TestHasEvent: expecting event true, got false")
	}
	if !manager.HasEventByName("One") {
		t.Errorf("TestHasEvent: expecting event true, got false")
	}
	event = mle_event.CreateEvent(1)
	//	assertEquals(event, 65536);
	manager.AddEvent(event, "Big")
	if !manager.HasEvent(65536) {
		t.Errorf("TestHasEvent: expecting event true, got false")
	}
	if !manager.HasEventByName("Big") {
		t.Errorf("TestHasEvent: expecting event true, got false")
	}
	if manager.HasEventByName("Two") {
		t.Errorf("TestHasEvent: expecting event false, got true")
	}
}
