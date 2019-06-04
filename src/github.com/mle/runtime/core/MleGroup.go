/**
 * @file MleGroup.go
 * Created on May 21, 2019. (msm@wizzerworks.com)
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

import (
	mle_util "github.com/mle/runtime/util"
)

/**
 * Base class for runtime actor groups.
 * <p>
 * The group is the minimal unit of actor loading at runtime.
 * A group may contain one or more actors and every actor must be loaded as part of a group
 * defined in a MleDwpGroup item in the Digital Workprint.
 * A group may  be loaded directly at runtime by calling the MleDppLoader.mleLoadGroup() method.
 * </p><p>
 * At runtime the MleGroup object holds an array of refernces to
 * all the actors it has loaded.  When the group is deleted, it
 * deletes all actors it refers to.  Similarly, a MleGroup loaded as
 * part of an MleScene is deleted when the containing scene is deleted.
 * </p>
 *
 * @see MleActor
 * @see MleScene
 *
 * @author Mark S. Millard
 * @version 1.0
 */
type MleGroup struct {
	/** The collection of Actors belonging to this group. */
	m_actors *mle_util.Vector
}

/**
 * Default constructor.
 */
func NewMleGroup() *MleGroup {
	p := new(MleGroup)
	p.m_actors = mle_util.NewVector()
	return p
}

/**
 * Initialize the group.
 * <p>
 * The class-specific initialization to be called after the group is
 * loaded and its actors' init() methods are called.
 * </p>
 *
 * @throws MleRuntimeException This exception is thrown if the
 * group can not be successfully initialized.
 */
func (group *MleGroup) Init() *MleError {
	return nil
}

/**
 * Dispose all resources associated with the Group.
 *
 * @throws MleRuntimeException This exception is thrown if the
 * group can not be successfully disposed.
 */
func (group *MleGroup) Dispose() *MleError {
	// Remove all elements from the Vector.
	for index := range *group.m_actors {
		group.m_actors.Delete(index)
	}
	group.m_actors = nil
	return nil
}

// ToString implements IObject interface.
func (group *MleGroup) ToString() string {
	return ""
}

/**
 * Add an Actor to the Group.
 *
 * @param actor The <code>MleActor</code> to add.
 */
func (group *MleGroup) Add(actor *MleActor) {
	group.m_actors.AppendVector(actor)
}

/**
 * Remove the specified Actor from the Group.
 *
 * @param actor The <code>MleActor</code> to remove.
 */
func (group *MleGroup) Remove(actor *MleActor) {
	index := group.m_actors.Peek(actor)
	if index >= 0 {
		group.m_actors.Delete(index)
	}
}
