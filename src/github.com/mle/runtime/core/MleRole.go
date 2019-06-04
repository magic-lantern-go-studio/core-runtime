/**
 * @file MleRole.go
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

// Import Magic Lantern packages.

/**
 * <code>MleRole</code> is a class that is used to
 * specify the platform-specific interface of an actor.
 * <p>
 * MleRole implements the platform-specific code
 * for an actor. Each instance of MleActor can have
 * zero or one corresponding role instance.
 * At runtime, the only space overhead for MleRole
 * is a reference to the role's actor instance.
 * </p><p>
 * getActor() returns the role instance's corresponding
 * actor instance. init() can be used to initialize the
 * role after it has been attached to its actor
 * and to the set/role hierarchy.
 * </p><p>
 * MleRole instances are created by MleDppLoader.loadGroup() and
 * Mle.loadScene() when actors are loaded from
 * either the workprint or playprint. MleDppLoader.loadGroup()
 * (which is called by MleDppLoader.loadScene()) goes through
 * the following sequence to initialize the role/actor
 * pair:
 * </p><p>
 * <ol>
 *   <li> When an actor is found inside a group that
 *        is being loaded, the actor is created
 *        (i.e., memory is allocated and its constructor is called).
 *   </li>
 *   <li> The values of the actor's properties
 *        in the workprint (or playprint) are inserted into the
 *        appropiate actor member variable locations in memory.
 *   </li>
 *   <li> If the workprint (or playprint) has a
 *        role class binding for the actor, then an
 *        instance of that role class is created
 *        (the role's constructor is called).
 *   </li>
 *   <li> The role instance is bound to its actor.
 *   </li>
 *   <li> The role is attached either to another
 *        role or to the current set, depending
 *        on whether the role class supports
 *        role-to-role attachment. addChild() is
 *        used to add a role to another role
 *        whenever the roles support attachment.
 *   </li>
 *   <li> The role's actor init() member function
 *        is called. This allows the actor to do
 *        final initialization since all of its data
 *        has been loaded and its role has been
 *        attached. The role actor might at this
 *        point call the role's init() function.
 *   </li>
 * </ol>
 * </p><p>
 * MleRole's instance variable <i>m_set</i> points to the
 * role instance's set (i.e., the set to which
 * it is directly or indirectly attached).
 * </p>
 *
 * @see MleActor
 * @see com.wizzer.mle.runtime.dpp.MleDppLoader#mleLoadGroup(int)
 * @see com.wizzer.mle.runtime.dpp.MleDppLoader#mleLoadScene(int)
 *
 * @author  Mark S. Millard
 * @version 1.0
 */
type MleRole struct {
	/** A reference to the actor that is rendered by this role. */
	m_actor *MleActor
	/** A reference to the role's set. */
	m_set *MleSet
}

/**
 * A factory method for creating a role.
 *
 * @return A new MleRole is returned.
 */
func _mlCreateMleRole(actor *MleActor) *MleRole {
	p := NewMleRoleWithActor(actor)
	return p
}

/**
 * The default constructor.
 */
func NewMleRole() *MleRole {
	p := new(MleRole)
	p.m_actor = nil
	// Place role in current set.
	//p.m_set = MleSet.GetCurrentSet()
	p.m_set = g_currentSet
	return p
}

/**
 * A constructor that is used to assign the associated actor.
 * <p>
 * MleRole constructor sets the actor to point to this role
 * because both sides of the link must be maintained at all times.
 * </p>
 *
 * @param actor The actor for this role.
 */
func NewMleRoleWithActor(actor *MleActor) *MleRole {
	p := NewMleRole()
	p.m_actor = actor
	actor.AttachRole(p)

	// Place role in current set.
	//p.m_set = MleSet.GetCurrentSet()
	p.m_set = g_currentSet
	return p
}

/**
 * Set the actor for this role.
 *
 * @param actor The actor to set.
 */
func (role *MleRole) SetActor(actor *MleActor) {
	role.m_actor = actor
	actor.AttachRole(role)
}

/**
 * Get the actor for this role.
 * <p>
 * This method should not be called inside the role's
 * constructor because the role's actor might not be necessarily
 * bound to the role when the role's constructor is called.
 * </p>
 *
 * @return A reference to the actor for this role is returned.
 */
func (role *MleRole) GetActor() *MleActor {
	return role.m_actor
}

/**
 * Initialize the role.
 * <p>
 * This method is a hook to initialize the role.  This may
 * involve initial data "pull" from the actor, scheduling, etc.
 * The base class implementation does nothing.
 * </p>
 */
func (role *MleRole) Init() {}

/**
 * Dispose all resources associated with the role.
 */
func (role *MleRole) Dispose() {}

// ToString implements IObject interface.
func (role *MleRole) ToString() string {
	return ""
}

/**
 * Add a child to this role.
 *
 * This method is used to attach other roles
 * to this role.  It should be overridden by sub-classes
 * of MleRole for which there's a semantic policy for attachment.
 * The base class implemetation does nothing.
 */
func (role *MleRole) AddChild(child *MleRole) {}
