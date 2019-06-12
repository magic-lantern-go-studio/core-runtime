/**
 * @file MleEventDispatcher_test.go
 * Created on June 12, 2019. (msm@wizzerworks.com)
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
package util_test

// import go packages.
import (
	"bytes"
	"testing"
	"strconv"

	mle_util "github.com/mle/runtime/util"
	mle_core "github.com/mle/runtime/core"
	mle_event "github.com/mle/runtime/event"
)

func TestNewMleEventDispatcher(t *testing.T) {
	// Construct an empty queue.
	q := mle_event.NewMleEventDispatcher()
	if q == nil {
		t.Errorf("TestNewMleEventDispatcher: NewMleEventDispatcher() returned nil")
	}
}

func TestNewMleEventDispatcherWithCapacity(t *testing.T) {
	// Construct an empty queue.
	q := mle_event.NewMleEventDispatcherWithCapacity(256)
	if q == nil {
		t.Errorf("TestNewMleEventDispatcher: NewMleEventDispatcher() returned nil")
	}
}

/*
 * Extends MleEventCallback.
 */
type _State struct {
	t *testing.T
	mCb *mle_event.MleEventCallback
	mName string
    mState int
}

/*
 * The default constructor.
 */
func _NewState(t *testing.T) *_State {
	p := new(_State)
	p.t = t
	p.mCb = mle_event.NewMleEventCallback()
	p.mName = "UNKNOWN"
    p.mState = mle_event.MLE_EVENT_INVALID_ID
	return p
}

/*
 * A constructor that initializes the state.
 * 
 * @param name The state name.
 */
func _NewStateWithName(name string, t *testing.T) *_State {
	p := new(_State)
	p.t = t
	p.mCb = mle_event.NewMleEventCallback()
	p.mName = name
    p.mState = mle_event.MLE_EVENT_INVALID_ID
	return p
}

/*
 * Execute this state.
 * 
 * @param event The event that is being dispatched by the handler.
 * @param clientData Client data registered with this handler.
 * 
 * @return If the event is successfully dispatched, then <b>true</b> will be
 * returned. Otherwise, <b>false</b> will be returned.
 */
func (s *_State) Dispatch(event mle_event.MleEvent, clientData mle_util.IObject) bool {
	var retValue = true
			 
	// Do some work.
	//System.out.println("Executing Callback: " + m_name + "\n" +
	//					"\tCaller Data: " + event.getCallData() + "\n" +
	//					"\tClient Data: " + clientdata)
	var buf bytes.Buffer
	buf.WriteString("Executing Callback: ")
	buf.WriteString(s.mName + "\n")
	buf.WriteString("\tCaller Data: ")
	//buf.WriteString(event.GetCallData().ToString() + "\n")
	buf.WriteString("\tClient Data: ")
	buf.WriteString(clientData.ToString())
	s.t.Logf(buf.String())
	s.mState = event.GetId()
				 
	return retValue
}

func (s *_State) Enable(enable bool) {
	s.mCb.Enable(enable)
}

func (s *_State) IsEnabled() bool {
	return s.mCb.IsEnabled();
}

var _machine *mle_event.MleEventDispatcher
var _states [5]int
var _ids [5]mle_core.IMleCallbackId

// Setup the test case.
var setupComplete = false
func setup() {
	// Create the event dispatcher.
	_machine = mle_event.NewMleEventDispatcher()
            
	// Initialize the composite event identifiers.
	// State 0.
	_states[0] = mle_event.MakeId(0x0000, 0x0001)
	// State 1.
	_states[1] = mle_event.MakeId(0x0000, 0x0002)
	// State 2.
	_states[2] = mle_event.MakeId(0x0000, 0x0003)
	// State 3.
	_states[3] = mle_event.MakeId(0x0000, 0x0004)
	// State 4.
	_states[4] = mle_event.MakeId(0x0000, 0x0005)

	setupComplete = true
}

// Execute mainloop, processing each event in immediate mode.
func mainloop() {
	for i := 0; i < len(_states); i++ {
		var buf bytes.Buffer
		buf.WriteString("Processing State = " + strconv.Itoa(_states[i]))
		_machine.ProcessEvent(_states[i], buf.String(), mle_event.MLE_EVENT_IMMEDIATE);
	}
}

type clientData struct {
	mData string
}

func newClientData(str string) *clientData {
	p := new(clientData)
	p.mData = str
	return p
}

func (cd *clientData) ToString() string {
	return cd.mData
}

// Initialize the state machine.
func initStateMachine(cb *_State) {
    if ! setupComplete {
		setup()
	}

	// Add the handlers to the event dispatcher.
	clientData := newClientData("Next State = " + strconv.Itoa(_states[1]))
	_ids[0], _ = _machine.InstallEventCB(_states[0],cb,clientData)
	clientData = newClientData("Next State = " + strconv.Itoa(_states[2]))
	_ids[1], _ = _machine.InstallEventCB(_states[1],cb,clientData)
	clientData = newClientData("Next State = " + strconv.Itoa(_states[3]))
	_ids[2], _ = _machine.InstallEventCB(_states[2],cb,clientData)
	clientData = newClientData("Next State = " + strconv.Itoa(_states[4]))
	_ids[3], _ = _machine.InstallEventCB(_states[3],cb,clientData)
	clientData = newClientData("Next State = " + strconv.Itoa(_states[0]))
	_ids[4], _ = _machine.InstallEventCB(_states[4],cb,clientData)
	
	// Enable the handlers.
	_machine.EnableEventCB(_states[0],_ids[0])
	_machine.EnableEventCB(_states[1],_ids[1])
	_machine.EnableEventCB(_states[2],_ids[2])
	_machine.EnableEventCB(_states[3],_ids[3])
	_machine.EnableEventCB(_states[4],_ids[4])
}

/**
 * Test the MleEventDispatcher using a simple state machine.
 * Process the events in MLE_IMMEDIATE_MODE;
 */
func TestImmediateMode(t *testing.T) {
	// Create the state machine callback handlers.
	var stateCB = _NewState(t)
			 
	initStateMachine(stateCB)
			 
	// Process the state machine.
	t.Logf("TestImmediateMode: processing events.")
	mainloop();
}
 