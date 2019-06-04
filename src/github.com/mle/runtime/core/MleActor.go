/**
 * @file MleActor.go
 * Created on April 26, 2019. (msm@wizzerworks.com)
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

/**
 * <code>MleActor</code> is a class that encapsulates a title element
 * that is independent of a specific target platform.
 * <p>
 * This is the base class for all Magic Lantern actors.
 * MleActor provides the baseline runtime and rehearsal-time
 * functionality of the actors.
 * </p><p>
 * For runtime, MleActor simply provides the ability
 * to associate an actor instance to a role instance.
 * For rehearsal time, MleActor provides many hooks
 * that allow Magic Lantern authoring tools, like the
 * Scene Editor, to monitor, control and edit the actor
 * instance in its runtime environment.
 * </p><p>
 * Use init() to initialize the actor after the actor's
 * member variables (i.e., properties) have been
 * loaded by MleDppLoader.loadGroup() or MleDppLoader.loadScene().
 * Use getRole() to get the actor's role instance,
 * </p>
 *
 * @see MleRole
 * @see com.wizzer.mle.runtime.dpp.MleDppLoader#mleLoadGroup(int)
 * @see com.wizzer.mle.runtime.dpp.MleDppLoader#mleLoadScene(int)
 *
 * @author  Mark S. Millard
 * @version 1.0
 */
type MleActor struct {
	/** A reference to the role for this actor to play. */
	m_role *MleRole
	/** The collection of "PropChange" event listeners, per property. */
	//protected HashMap<String,Vector<IMlePropChangeListener>> m_propChangeListeners;
	m_propChangeListeners map[string](mle_util.Vector)
}

/**
 * The default constructor.
 */
func NewMleActor() *MleActor {
	p := new(MleActor)
	// The role should be set to null.
	p.m_role = nil
	//m_propChangeListeners = new HashMap<String,Vector<IMlePropChangeListener>>()
	p.m_propChangeListeners = make(map[string]mle_util.Vector)
	return p
}

/**
 * Get the actor's associated role.
 * <p>
 * Returns a pointer to the actor's role, if any.
 * Use getRole() to get the actor's role instance,
 * if any. When loading an actor instance using MleDppLoader.loadGroup()
 * or MleLaod.loadScene(), the role, if any, will be set
 * after the actor's constructor is called and before
 * the actor's init() function is called. Therefore,
 * the result of calling getRole() before init()
 * is called is undefined.
 * </p>
 *
 * @return The actor's associcated role is returned.
 */
func (actor *MleActor) GetRole() *MleRole {
	return actor.m_role
}

/**
 * Remove the role from the actor.
 * <p>
 * The attached role provides notification when it is destroyed.
 * This method is made protected to allow sub-class to provide
 * notification to other classes, or for the actor to possibly
 * delete self.
 * </p><p>
 * Note that only the role instance should call this method
 * since it can clean-up the actor/role relationship.
 * </p>
 */
func (actor *MleActor) RemoveRole() {
	actor.m_role = nil
}

/**
 * Attach a role to the actor.
 * <p>
 * A role can only be added if the actor currently does not have one.
 * Note that only the role instance should call this method.
 * </p>
 *
 * @param role The role to be attached to this actor.
 */
func (actor *MleActor) AttachRole(role *MleRole) {
	if actor.m_role == nil {
		actor.m_role = role
	}
}

/**
 * Initialize the actor.
 * <p>
 * This method is a hook to do any initialization after property
 * values are inserted into the actor when the actor is directly loaded
 * by MleDppLoader.loadGroup() or indirectly by MleDppLoader.loadScene().
 * Typically, the actor may schedule itself or initialize its
 * role here.  The base init() function does nothing.
 * Use init() to initialize the actor after the actor's
 * member variables (i.e., properties) have been
 * loaded by MleDppLoader.loadGroup() or MleDppLoader.loadScene().
 * init() is called after the actor's data has
 * been loaded into memory, so this is the safest
 * time to perform initialization and synchronization
 * with the rest of the environment.
 * </p>
 *
 * @throws MleRuntimeException This exception is thrown if the
 * actor can not be successfully initialized.
 */
func (actor *MleActor) Init() {}

/**
 * Dispose all resources associated with the Actor.
 *
 * @throws MleRuntimeException This exception is thrown if the
 * actor can not be successfully disposed.
 */
func (actor *MleActor) Dispose() {}

// ToString implements IObject interface.
func (actor *MleActor) ToString() string {
	return ""
}

// Implement IMleObject interface.

func (actor *MleActor) GetProperty(name string) IMleProp {
	// TBD - log something here.
	return nil
}

func (actor *MleActor) SetProperty(name string, property IMleProp) {
	// TBD - log something here.
}

func (actor *MleActor) SetPropertyArray(name string, length int, nElements int, value io.ByteReader) {
	// TBD - log something here.
}

func (actor *MleActor) AddPropertyChangeListener(name string, listener IMleListener) MleError {
	var err *MleError

	// ToDo: can we validate that the listener is an IMlePropChangeListener?

	if name == "" {
		err := NewMleError("Property name must not be empty.", 0, nil)
		return *err
	}

	//Vector<IMlePropChangeListener> listeners;
	var listeners *mle_util.Vector

	value, found := actor.m_propChangeListeners[name]
	if !found {
		// Add a new container to collect the listeners for the named property.
		listeners = mle_util.NewVector()
		actor.m_propChangeListeners[name] = *listeners
	} else {
		// Use the existing collection container.
		listeners = &value
	}

	// Add the property change listener.
	listeners.AppendVector(listener)

	return *err
}

func (actor *MleActor) RemovePropertyChangeListener(name string, listener IMleListener) MleError {
	var err *MleError

	// ToDo: can we validate that the listener is an IMlePropChangeListener?

	if name == "" {
		err := NewMleError("Property name must not be empty.", 0, nil)
		return *err
	}

	value, found := actor.m_propChangeListeners[name]
	if found {
		listeners := value
		index := listeners.Peek(listener)
		if index >= 0 {
			listeners.Delete(index)
		}
	}

	return *err
}

func (actor *MleActor) NotifyPropertyChange(name string, oldProperty IMleProp, newProperty IMleProp) {
	// Create a new "event".
	args := make(map[string]interface{})
	args["property"] = name
	args["old_property"] = oldProperty
	args["new_property"] = newProperty

	value, found := actor.m_propChangeListeners[name]
	if found {
		listeners := value
		for i := 0; i < len(listeners); i++ {
			// The expectation is that the listener is an instance of IMlePropChangeListener.
			// The method PropChangedEvent will be called upon recieving the SendEvent.
			listener := listeners.ElementAt(i).(IMleListener)
			listener.SendEvent(actor, args)
		}
	}
}
