/**
 * @file MleScene.go
 * Created on May 23, 2019. (msm@wizzerworks.com.com)
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
	mle_util "github.com/mle/runtime/util"
)

// The global Scene.
var g_globalScene *MleScene

// The current Scene.
var g_currentScene *MleScene

/**
 * Base class for Magic Lantern Runtime scenes.
 * <p>
 * The scene is the maximal  unit of actor loading at run time.
 * A scene may contain one or more groups or group references.
 * A scene may be loaded directly at runtime by calling the
 * <code>MleDppLoader.mleLoadScene()</code> method.
 * </p><p>
 * Two distinguished scenes may be active at once in a title.  The
 * <i>current</i> scene is set by scene loading, unloading, or
 * changing methods.  The <i>global</i> scene is under user control
 * and is intended to hold a scene that will govern transitions between
 * two current scenes, for instance during a level change in a game.
 * </p><p>
 * At run time the MleScene object holds an array of references to
 * all the groups it has loaded.  When the scene is deleted, it
 * deletes all groups it refers to.
 * </p>
 *
 * @see MleGroup
 *
 * @author Mark S. Millard
 * @version 1.0
 */
type MleScene struct {
	// The collection of Groups belonging to this Scene.
	m_groups *mle_util.Vector
}

/**
 * The default constructor.
 */
func NewMleScene() *MleScene {
	p := new(MleScene)
	p.m_groups = mle_util.NewVector()
	return p
}

/**
 * Get the global scene.
 *
 * @return A reference to the global Scene is returned.
 */
func GetGlobalScene() *MleScene {
	return g_globalScene
}

/**
 * Set the caller as the global scene.
 */
func (scene *MleScene) SetGlobalScene() {
	g_globalScene = scene
}

/**
 * Get the current scene.
 *
 * @return A reference to the current Scene is returned.
 */
func GetCurrentScene() *MleScene {
	return g_currentScene
}

/**
 * Set the caller as the current scene.
 */
func (scene *MleScene) SetCurrentScene() {
	g_currentScene = scene
}

// Clear the current scene.
func ClearCurrentScene() {
	g_currentScene = nil
}

/**
 * Initialize the scene.
 * <p>
 * The class-specific initialization to be called after the scene is
 * loaded and its groups' Onit() methods are called.
 * </p>
 *
 * @throws MleRuntimeException This exception is thrown if the
 * scene can not be successfully initialized.
 */
func (scene *MleScene) Init() {}

/**
 * Dispose all resources associated with the Scene.
 *
 * @throws MleRuntimeException This exception is thrown if the
 * scene can not be successfully initialized.
 */
func (scene *MleScene) Dispose() {}

/**
 * Delete the global Scene.
 *
 * @throws MleRuntimeException This exception is always thrown becuase it has not
 * yet been implemented.
 */
func DeleteGlobalScene() *MleError {
	msg := "MleScene: DeleteGlobalScene not implemented."
	err := NewMleError(msg, 0, nil)
	return err
}

/**
 * Delete the current Scene.
 *
 * @throws MleRuntimeException This exception is thrown if the
 * scene can not be successfully deleted.
 */
func DeleteCurrentScene() {
	// Clear current Scene.
	old := GetCurrentScene()
	ClearCurrentScene()

	// Get rid of old scene and all its contents by disposing it.
	if old != nil {
		old.Dispose()
	}
}

/**
 * Changes the current scene to that passed in. In this default
 * implementation, it simply disposes the old currently active scene,
 * replacing it with the new one.
 *
 * @param newScene The new Scene to switch to.
 *
 * @throws MleRuntimeException This exception is thrown if the
 * old scene can not be successfully disposed of.
 */
func ChangeCurrentScene(newScene *MleScene) *MleScene {
	// Swap old Scene for new Scene.

	// Do it this way so that we\'ve set the new Scene before deleting
	// the old one, in case we\'re the old one.
	old := GetCurrentScene()
	newScene.SetCurrentScene()

	// Get rid of old scene and all its contents by disposing it.
	if old != nil {
		old.Dispose()
	}

	return newScene
}

/**
 * Add a Group to the Scene.
 *
 * @param group The <code>MleGroup</code> to add.
 */
func (scene *MleScene) Add(group *MleGroup) {
	scene.m_groups.AddElement(group)
}

/**
 * Remove the specified Group from the Scene.
 *
 * @param group The <code>MleGroup</code> to remove.
 */
func (scene *MleScene) Remove(group *MleGroup) {
	scene.m_groups.RemoveElement(group)
}
