/**
 * @file MleSet.go
 * Created on April 26, 2019. (msm@wizzerworks.com.com)
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
package core

// Import go packages.
import (
	"io"

	mle_util "github.com/mle/runtime/util"
)

// Import Magic Lantern packages.

/**
 * This global holds the current set, which is valid during a group
 * load.
 */
var g_currentSet *MleSet

/**
 * <code>MleSet</code> is a class that encapsulates a
 * actor/role policy for platform-specific behavior.
 * <p>
 * This is the base class for all Magic Lantern sets.
 * MleSet provides the baseline runtime and rehearsal-time
 * functionality of the sets.
 * </p><p>
 * Use init() to initialize the set after the sets's
 * member variables (i.e., properties) have been
 * loaded by MleDppLoader.loadGroup() or MleDppLoader.loadScene().
 * Use attachRoles() to maintain a hierarchy of roles that
 * are managed by the set.
 * </p>
 *
 * @see MleRole
 * @see com.wizzer.mle.runtime.dpp.MleDppLoader#mleLoadGroup(int)
 * @see com.wizzer.mle.runtime.dpp.MleDppLoader#mleLoadScene(int)
 *
 * @author  Mark S. Millard
 * @version 1.0
 */
type MleSet struct {
	/** The collection of "PropChange" event listeners, per property. */
	//protected HashMap<String,Vector<IMlePropChangeListener>> m_propChangeListeners;
	m_propChangeListeners map[string](mle_util.Vector)
}

/**
 * The default constructor.
 */
func NewMleSet() *MleSet {
	p := new(MleSet)
	//m_propChangeListeners = new HashMap<String,Vector<IMlePropChangeListener>>()
	p.m_propChangeListeners = make(map[string]mle_util.Vector)
	return p
}

// String implements IObject interface.
func (stage *MleSet) String() string {
	return ""
}

/**
 * Get the current Set.
 * <p>
 * Actors or roles that are instantiated by a group load should only
 * refer to the current Set in their init() function
 * (to get a reference to the set that contains them).
 * </p>
 *
 * @return The current Set is returned. <b>null</b> may be returned if no
 * Set has been made current.
 */
func (set *MleSet) GetCurrentSet() *MleSet {
	return g_currentSet
}

/**
 * Set the current Set.
 *
 * @param set The actor/role behavior policy to set.
 */
func (set *MleSet) SetCurrentSet() {
	g_currentSet = set
}

/**
 * Attach a child role to its parent role.
 * <p>
 * The implementations of this function will generally cast the
 * MleRole arguments to the type of roles which the given
 * set use and then perform role-specific operations to perform
 * the attach.
 * </p>
 *
 * @param parent The role to attach the child role to.
 * @param child The role which is being attached.
 *
 * @throws MleRuntimeException This exception is thrown if the
 * set can not be successfully attach the roles.
 */
func (set *MleSet) AttachRoles(parent *MleRole, child *MleRole) {}

/**
 * Initialize the set.
 * <p>
 * Typically, the set may register itself with the scheduler.
 * </p>
 *
 * @throws MleRuntimeException This exception is thrown if the
 * set can not be successfully initialized.
 */
func (set *MleSet) Init() {}

/**
 * Dispose all resources associated with the Set.
 *
 * @throws MleRuntimeException This exception is thrown if the
 * set can not be successfully disposed.
 */
func (set MleSet) Dispose() {}

// Implement IMleObject interface.

func (set *MleSet) GetProperty(name string) IMleProp {
	// TBD - log something here.
	return nil
}

func (set *MleSet) SetProperty(name string, property IMleProp) {
	// TBD - log something here.
}

func (set *MleSet) SetPropertyArray(name string, length int, nElements int, value io.ByteReader) {
	// TBD - log something here.
}

func (set *MleSet) AddPropertyChangeListener(name string, listener IMleListener) MleError {
	var err *MleError

	if name == "" {
		err = NewMleError("Property name must not be empty.", 0, nil)
		return *err
	}

	var listeners *mle_util.Vector

	value, found := set.m_propChangeListeners[name]
	if !found {
		// Add a new container to collect the listeners for the named property.
		listeners = mle_util.NewVector()
		set.m_propChangeListeners[name] = *listeners
	} else {
		// Use the existing collection container.
		*listeners = value
	}

	// Add the property change listener.
	listeners.AppendVector(listener)

	return *err
}

func (set *MleSet) RemovePropertyChangeListener(name string, listener IMleListener) MleError {
	var err *MleError

	if name == "" {
		err = NewMleError("Property name must not be empty.", 0, nil)
		return *err
	}

	value, found := set.m_propChangeListeners[name]
	if found {
		listeners := value
		index := listeners.Peek(listener)
		if index >= 0 {
			listeners.Delete(index)
		}
	}

	return *err
}

func (set *MleSet) NotifyPropertyChange(name string, oldProperty IMleProp, newProperty IMleProp) {
	// Create a new "event".
	args := make(map[string]interface{})
	args["property"] = name
	args["old_property"] = oldProperty
	args["new_property"] = newProperty

	value, found := set.m_propChangeListeners[name]
	if found {
		listeners := value
		for i := 0; i < len(listeners); i++ {
			// The expectation is that the listener is an instance of IMlePropChangeListener.
			// The method PropChangedEvent will be called upon recieving the SendEvent.
			listener := listeners.ElementAt(i).(IMleListener)
			listener.SendEvent(set, args)
		}
	}
}
