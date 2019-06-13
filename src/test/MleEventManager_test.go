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

func testMleEventManager_setUp() {
	manager := mle_event.NewMleEventManager()
	manager.Clear()
}

func testMleEventManager_tearDown() {
	// Do nothing for now.
}

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
	testMleEventManager_setUp()

	manager := mle_event.NewMleEventManager()
	if manager == nil {
		t.Errorf("TestAddEvent: NewMleEventManager() returned nil")
	}

	// Create a new event in group 0.
    var event = mle_event.CreateEvent(0)
	manager.AddEvent(event, "")

	// Try to add it again, should not be able to.
	manager.AddEvent(event, "");
	if manager.Size() != 1 {
		// There should be four default events plus the one we just added.
		t.Errorf("TestAddEvent: EventManager size should be 5, got %d", manager.Size())
	}

	testMleEventManager_tearDown()
}

func TestRemoveEvent(t *testing.T) {
	testMleEventManager_setUp()

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

	if manager.Size() != 2 {
		// There should be two events.
		t.Errorf("TestAddEvent: EventManager size should be 2, got %d", manager.Size())
	}

	// Remove the last event.
	manager.RemoveEvent(event)
	if manager.Size() != 1 {
		// There should be one event less.
		t.Errorf("TestAddEvent: EventManager size should be 1, got %d", manager.Size())
	}

	testMleEventManager_tearDown()
}

/**
 * Test creating new events.
 */
func TestCreateEvent(t *testing.T) {
	testMleEventManager_setUp()

	manager := mle_event.NewMleEventManager()
	if manager == nil {
		t.Errorf("TestCreateEvent: NewMleEventManager() returned nil")
	}

	var event = mle_event.CreateEvent(0)
	//assertEquals(event, 0);
	if event != 0 {
		t.Errorf("TestCreateEvent: expecting event 0, got %d", event)
	}
	manager.AddEvent(event, "")
	event = mle_event.CreateEvent(0)
	//assertEquals(event, 1);
	if event != 1 {
		t.Errorf("TestCreateEvent: expecting event 1, got %d", event)
	}
	manager.AddEvent(event, "")
	event = mle_event.CreateEvent(0)
	//assertEquals(event, 2);
	if event != 2 {
		t.Errorf("TestCreateEvent: expecting event 2, got %d", event)
	}
	manager.AddEvent(event, "")
	event = mle_event.CreateEvent(0)
	//assertEquals(event, 3);
	if event != 3 {
		t.Errorf("TestCreateEvent: expecting event 3, got %d", event)
	}
	manager.AddEvent(event, "")
	event = mle_event.CreateEvent(0)
	//assertEquals(event, 4);
	if event != 4 {
		t.Errorf("TestCreateEvent: expecting event 4, got %d", event)
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
	if event != 5 {
		t.Errorf("TestCreateEvent: expecting event 5, got %d", event)
	}
	
	//assertEquals(MleEventManager.getInstance().size(), 10);
	if manager.Size() != 10 {
		// There should be ten events.
		t.Errorf("TestCreateEvent: EventManager size should be 19, got %d", manager.Size())
	}

	testMleEventManager_tearDown()
}

func TestHasEvent(t *testing.T) {
	testMleEventManager_setUp()

    manager := mle_event.NewMleEventManager()
	if manager == nil {
		t.Errorf("TestHasEvent: NewMleEventManager() returned nil")
	}

	var event = mle_event.CreateEvent(0)
	if event != 0 {
		t.Errorf("TestHasEvent: expecting event 0, got %d", event)
	}
    manager.AddEvent(event, "")
	event = mle_event.CreateEvent(0)
	if event != 1 {
		t.Errorf("TestHasEvent: expecting event 1, got %d", event)
	}
	manager.AddEvent(event, "")
	if !manager.HasEvent(0) {
		t.Errorf("TestHasEvent: expecting event true, got false")
	}
	if !manager.HasEvent(1) {
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

	testMleEventManager_tearDown()
}

func TestHasEventWithName(t *testing.T) {
	testMleEventManager_setUp()

	manager := mle_event.NewMleEventManager()
	if manager == nil {
		t.Errorf("TestHasEventWithName: NewMleEventManager() returned nil")
	}
	
	var event = mle_event.CreateEvent(0)
	manager.AddEvent(event, "")
	event = mle_event.CreateEvent(0)
	manager.AddEvent(event, "One")
	if !manager.HasEvent(0) {
		t.Errorf("TestHasEvent: expecting event true, got false")
	}
	if !manager.HasEvent(1) {
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

	testMleEventManager_tearDown()
}
