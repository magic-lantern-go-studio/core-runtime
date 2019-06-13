/**
 * @file MleTables.go
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

// The Singleton instance of the tables.
var g_theTables *MleTables

/**
 * This table gives the playprint chunk number for the corresponding
 * set in the mleRTSet table.
 */
var g_mleRTSetChunk []int

/**
 * This table gives the playprint chunk number for the corresponding
 * group in the mleRTGroup table.
 */
var g_mleRTGroupChunk []int

/**
 * This table gives the playprint chunk number for the corresponding
 * scene in the mleRTScene table.
 */
var g_mleRTSceneChunk []int

/**
 * The boot Scene.
 */
var g_mleBootScene int

/**
 * This class is used to define and register the tables generated for a
 * Magic Lantern application.
 *
 * @author Mark S. Millard
 */
type MleTables struct {
	/** The collection of Actor properties. */
	g_mleRTActorProperties *mle_util.Vector
	/** The collection of Set properties. */
	g_mleRTSetProperties *mle_util.Vector
	/** The collection of Actor classes. */
	g_mleRTActorClass *mle_util.Vector
	/** The collection of Role classes. */
	g_mleRTRoleClass *mle_util.Vector
	/** The collection of Set classes. */
	g_mleRTSetClass *mle_util.Vector
	/** The collection of Sets. */
	g_mleRTSet *mle_util.Vector
	/** The collection of Group classes. */
	g_mleRTGroupClass *mle_util.Vector
	/** The collection of Scene classes. */
	g_mleRTSceneClass *mle_util.Vector
	/** The collection of MediaRef classes. */
	g_mleRTMediaRefClass *mle_util.Vector
	/** The collection of MediaRefs. */
	g_mleRTMediaRef *mle_util.Vector

	/** A reference to an Observable pattern. */
	m_observable *mle_util.Observable
}

/**
 * The default constructor.
 * <p>
 * Since the MleTables is supposed to be a Singleton instance,
 * this function should not be called outside of GetMleTablesInstance().
 * </p>
 */
func newMleTables() *MleTables {
	p := new(MleTables)
	p.g_mleRTActorProperties = mle_util.NewVector()
	p.g_mleRTSetProperties = mle_util.NewVector()
	p.g_mleRTActorClass = mle_util.NewVector()
	p.g_mleRTRoleClass = mle_util.NewVector()
	p.g_mleRTSetClass = mle_util.NewVector()
	p.g_mleRTSet = mle_util.NewVector()
	p.g_mleRTGroupClass = mle_util.NewVector()
	p.g_mleRTSceneClass = mle_util.NewVector()
	p.g_mleRTMediaRefClass = mle_util.NewVector()
	p.g_mleRTMediaRef = mle_util.NewVector()

	p.m_observable = mle_util.NewObservable()
	return p
}

/**
 * Get the Singleton instance of the table manager.
 *
 * @return A reference to the <code>MleTables</code> is returned.
 */
func GetMleTablesInstance() *MleTables {
	if g_theTables == nil {
		g_theTables = newMleTables()
		g_mleRTSetChunk = nil
		g_mleRTGroupChunk = nil
		g_mleRTSceneChunk = nil
		g_mleBootScene = -1
	}
	return g_theTables
}

/**
 * This class is a runtime Actor/Set Property Table Entry.
 */
type MleRTPropertyEntry struct {
	// The name of the Class.
	m_classname string
	// The name of the Field.
	m_fieldname string
}

func NewMleRTPropertyEntry() *MleRTPropertyEntry {
	p := new(MleRTPropertyEntry)
	p.m_classname = ""
	p.m_fieldname = ""
	return p
}

func NewMleRTPropertyEntryWithClassAndField(classname string, fieldname string) *MleRTPropertyEntry {
	p := new(MleRTPropertyEntry)
	p.m_classname = classname
	p.m_fieldname = fieldname
	return p
}

// String implements IObject interface.
func (entry *MleRTPropertyEntry) String() string {
	// ToDo: return a usable string value.
	return ""
}

/**
 * Get the property field.
 *
 * @return The name of the field is returned.
 */
func (prop *MleRTPropertyEntry) GetProperty() string {
	return prop.m_fieldname
}

/**
 * This class is a runtime Actor Type Table Entry.
 */
type MleRTActorClassEntry struct {
	/** The class name for invoking the default constructor. */
	m_classname string
	/** The property table offset. */
	m_offset int
}

func NewMleRTActorClassEntry() *MleRTActorClassEntry {
	p := new(MleRTActorClassEntry)
	p.m_classname = ""
	p.m_offset = 0
	return p
}

func NewMleRTActorClassEntryWithClassAndOffset(classname string, offset int) *MleRTActorClassEntry {
	p := new(MleRTActorClassEntry)
	p.m_classname = classname
	p.m_offset = offset
	return p
}

// String implements IObject interface.
func (acentry *MleRTActorClassEntry) String() string {
	// ToDo: return a usable string value.
	return ""
}

// CreateActor creates an instance of an Actor based on an ActorClassEntry.
// A Class object must have been registered with the ClassFactory.
func (acentry *MleRTActorClassEntry) CreateActor() (*mle_util.Object, *MleError) {
	var newActor mle_util.Object
	var mlerr *MleError
	var obj interface{}

	// See if class exists in our registry.
	var found = false
	for key, data := range mle_util.GClassRegistry {
		if key == acentry.m_classname {
			// Found registered class, make sure that it can create a new instance.
			found = mle_util.MethodExists(data, "NewInstance")
			if found {
				// The method name exists and we can proceed.
				obj = data
				break
			}
		}
	}

	if !found {
		// Return an error.
		msg := "CreateActor: class " + acentry.m_classname + " not found."
		mlerr = NewMleError(msg, 0, nil)
	} else {
		// Call method to create an Actor. There are no input or output parameters.
		var err error
		newActor, err = mle_util.Invoke(obj, "NewInstance")
		if err != nil {
			// Calling method on Class object failed.
			mlerr = NewMleError(err.Error(), 0, err)
		}
	}

	return &newActor, mlerr
}

/**
 * This class is a runtime Role Type Table Entry.
 */
type MleRTRoleClassEntry struct {
	/** The class name for invoking the default constructor. */
	m_classname string
}

func NewMleRTRoleClassEntry() *MleRTRoleClassEntry {
	p := new(MleRTRoleClassEntry)
	p.m_classname = ""
	return p
}

func NewMleRTActorRoleEntryWithClass(classname string) *MleRTRoleClassEntry {
	p := new(MleRTRoleClassEntry)
	p.m_classname = classname
	return p
}

// String implements IObject interface.
func (rcentry *MleRTRoleClassEntry) String() string {
	// ToDo: return a usable string value.
	return ""
}

// CreateRole creates an instance of a Role based on a RoleClassEntry.
// A Class object must have been registered with the ClassFactory.
func (rcentry *MleRTRoleClassEntry) CreateRole(actor *MleActor) (*mle_util.Object, *MleError) {
	var newRole mle_util.Object
	var mlerr *MleError
	var obj interface{}

	// See if class exists in our registry.
	var found = false
	for key, data := range mle_util.GClassRegistry {
		if key == rcentry.m_classname {
			// Found registered class, make sure that it can create a new instance.
			found = mle_util.MethodExists(data, "NewInstance")
			if found {
				// The method name exists and we can proceed.
				obj = data
				break
			}
		}
	}

	if !found {
		// Return an error.
		msg := "CreateRole: class " + rcentry.m_classname + " not found."
		mlerr = NewMleError(msg, 0, nil)
	} else {
		// Call method to create a Role. There are no input or output parameters.
		var err error
		newRole, err = mle_util.Invoke(obj, "NewInstance")
		if err != nil {
			// Calling method on Class object failed.
			mlerr = NewMleError(err.Error(), 0, err)
		} else {
			// Set the Actor on the new Role.
			//role := newRole.(MleRole)
			//role.SetActor(actor)
			newRole.(*MleRole).SetActor(actor)
		}
	}

	return &newRole, mlerr
}

/**
 * This class is a runtime Set Type Table Entry.
 */
type MleRTSetClassEntry struct {
	/** The class name for invoking the default constructor. */
	m_classname string
	/** The property table offset. */
	m_offset int
}

func NewMleRTSetClassEntry() *MleRTSetClassEntry {
	p := new(MleRTSetClassEntry)
	p.m_classname = ""
	p.m_offset = 0
	return p
}

func NewMleRTSetClassEntryWithClassAndOffset(classname string, offset int) *MleRTSetClassEntry {
	p := new(MleRTSetClassEntry)
	p.m_classname = classname
	p.m_offset = offset
	return p
}

// String implements IObject interface.
func (scentry *MleRTSetClassEntry) String() string {
	// ToDo: return a usable string value.
	return ""
}

// CreateSet creates an instance of a Set based on a SetClassEntry.
// A Class object must have been registered with the ClassFactory.
func (scentry *MleRTSetClassEntry) CreateSet() (*mle_util.Object, *MleError) {
	var newSet mle_util.Object
	var mlerr *MleError
	var obj interface{}

	// See if class exists in our registry.
	var found = false
	for key, data := range mle_util.GClassRegistry {
		if key == scentry.m_classname {
			// Found registered class, make sure that it can create a new instance.
			found = mle_util.MethodExists(data, "NewInstance")
			if found {
				// The method name exists and we can proceed.
				obj = data
				break
			}
		}
	}

	if !found {
		// Return an error.
		msg := "CreateSet: class " + scentry.m_classname + " not found."
		mlerr = NewMleError(msg, 0, nil)
	} else {
		// Call method to create a Set. There are no input or output parameters.
		var err error
		newSet, err = mle_util.Invoke(obj, "NewInstance")
		if err != nil {
			// Calling method on Class object failed.
			mlerr = NewMleError(err.Error(), 0, err)
		}
	}

	return &newSet, mlerr
}

/**
 * This class is a runtime Group Type Table Entry.
 */
type MleRTGroupClassEntry struct {
	/** The class name for invoking the default constructor. */
	m_classname string
}

func NewMleRTGroupClassEntry() *MleRTGroupClassEntry {
	p := new(MleRTGroupClassEntry)
	p.m_classname = ""
	return p
}

func NewMleRTGroupEntryWithClass(classname string) *MleRTGroupClassEntry {
	p := new(MleRTGroupClassEntry)
	p.m_classname = classname
	return p
}

// String implements IObject interface.
func (gcentry *MleRTGroupClassEntry) String() string {
	// ToDo: return a usable string value.
	return ""
}

// CreateGroup creates an instance of a Group based on a GroupClassEntry.
// A Class object must have been registered with the ClassFactory.
func (gcentry *MleRTGroupClassEntry) CreateGroup() (*mle_util.Object, *MleError) {
	var newGroup mle_util.Object
	var mlerr *MleError
	var obj interface{}

	// See if class exists in our registry.
	var found = false
	for key, data := range mle_util.GClassRegistry {
		if key == gcentry.m_classname {
			// Found registered class, make sure that it can create a new instance.
			found = mle_util.MethodExists(data, "NewInstance")
			if found {
				// The method name exists and we can proceed.
				obj = data
				break
			}
		}
	}

	if !found {
		// Return an error.
		msg := "CreateGroup: class " + gcentry.m_classname + " not found."
		mlerr = NewMleError(msg, 0, nil)
	} else {
		// Call method to create a Group. There are no input or output parameters.
		var err error
		newGroup, err = mle_util.Invoke(obj, "NewInstance")
		if err != nil {
			// Calling method on Class object failed.
			mlerr = NewMleError(err.Error(), 0, err)
		}
	}

	return &newGroup, mlerr
}

/**
 * This class is a runtime MediaRef Type Table Entry.
 */
type MleRTMediaRefClassEntry struct {
	/** The class name for invoking the default constructor. */
	m_classname string
}

func NewMleRTMediaRefClassEntry() *MleRTMediaRefClassEntry {
	p := new(MleRTMediaRefClassEntry)
	p.m_classname = ""
	return p
}

func NewMleRTMediaRefClassEntryWithClass(classname string) *MleRTMediaRefClassEntry {
	p := new(MleRTMediaRefClassEntry)
	p.m_classname = classname
	return p
}

// String implements IObject interface.
func (mcentry *MleRTMediaRefClassEntry) String() string {
	// ToDo: return a usable string value.
	return ""
}

// CreateMediaRef creates an instance of a MediaRef based on a MediaRefClassEntry.
// A Class object must have been registered with the ClassFactory.
func (mcentry *MleRTMediaRefClassEntry) CreateMediaRef() (*mle_util.Object, *MleError) {
	var newMediaRef mle_util.Object
	var mlerr *MleError
	var obj interface{}

	// See if class exists in our registry.
	var found = false
	for key, data := range mle_util.GClassRegistry {
		if key == mcentry.m_classname {
			// Found registered class, make sure that it can create a new instance.
			found = mle_util.MethodExists(data, "NewInstance")
			if found {
				// The method name exists and we can proceed.
				obj = data
				break
			}
		}
	}

	if !found {
		// Return an error.
		msg := "CreateMediaRef: class " + mcentry.m_classname + " not found."
		mlerr = NewMleError(msg, 0, nil)
	} else {
		// Call method to create a MediaRef. There are no input or output parameters.
		var err error
		newMediaRef, err = mle_util.Invoke(obj, "NewInstance")
		if err != nil {
			// Calling method on Class object failed.
			mlerr = NewMleError(err.Error(), 0, err)
		}
	}

	return &newMediaRef, mlerr
}

/**
 * This class is a runtime Scene Type Table Entry.
 */
type MleRTSceneClassEntry struct {
	/** The class name for invoking the default constructor. */
	m_classname string
}

func NewMleRTSceneClassEntry() *MleRTSceneClassEntry {
	p := new(MleRTSceneClassEntry)
	p.m_classname = ""
	return p
}

func NewMleRTSceneEntryWithClass(classname string) *MleRTSceneClassEntry {
	p := new(MleRTSceneClassEntry)
	p.m_classname = classname
	return p
}

// String implements IObject interface.
func (scentry *MleRTSceneClassEntry) String() string {
	// ToDo: return a usable string value.
	return ""
}

// CreateScene creates an instance of a Scene based on a SceneClassEntry.
// A Class object must have been registered with the ClassFactory.
func (scentry *MleRTSceneClassEntry) CreateScene() (*mle_util.Object, *MleError) {
	var newScene mle_util.Object
	var mlerr *MleError
	var obj interface{}

	// See if class exists in our registry.
	var found = false
	for key, data := range mle_util.GClassRegistry {
		if key == scentry.m_classname {
			// Found registered class, make sure that it can create a new instance.
			found = mle_util.MethodExists(data, "NewInstance")
			if found {
				// The method name exists and we can proceed.
				obj = data
				break
			}
		}
	}

	if !found {
		// Return an error.
		msg := "CreateScene: class " + scentry.m_classname + " not found."
		mlerr = NewMleError(msg, 0, nil)
	} else {
		// Call method to create a Scene. There are no input or output parameters.
		var err error
		newScene, err = mle_util.Invoke(obj, "NewInstance")
		if err != nil {
			// Calling method on Class object failed.
			mlerr = NewMleError(err.Error(), 0, err)
		}
	}

	return &newScene, mlerr
}

type MleRTSetEntry struct {
	/** The class name for invoking the default constructor. */
	m_classname string
	/** A reference to the instance. */
	m_theSet *MleSet
}

func NewMleRTSetEntryWithClassAndSet(classname string, set *MleSet) *MleRTSetEntry {
	p := new(MleRTSetEntry)
	p.m_classname = classname
	p.m_theSet = set
	return p
}

// String implements IObject interface.
func (sentry *MleRTSetEntry) String() string {
	// ToDo: return a usable string value.
	return ""
}

type MleRTMediaRefEntry struct {
	/** The class name for invoking the default constructor. */
	m_classname string
}

func NewMleRTMediaRefEntryWithClass(classname string) *MleRTMediaRefEntry {
	p := new(MleRTMediaRefEntry)
	p.m_classname = classname
	return p
}

// String implements IObject interface.
func (sentry *MleRTMediaRefEntry) String() string {
	// ToDo: return a usable string value.
	return ""
}

/**
 * Add an Actor Property entry.
 *
 * @param property The entry to add.
 *
 * @return If the property is successfully added, then <b>true</b>
 * will be returned. Otherwise, <b>false</b> will be returned.
 *
 * @throws MleRuntimeException This exception is thrown if the entry
 * is not an instance of <code>MleRTPropertyEntry</code>.
 */
func (tables *MleTables) addActorProperty(property *MleRTPropertyEntry) (bool, *MleError) {
	added := true

	if !mle_util.InstanceOf(property, (*MleRTPropertyEntry)(nil)) {
		msg := "addActorProperty: Not an Actor property."
		return false, NewMleError(msg, 0, nil)
	}
	tables.g_mleRTActorProperties.AddElement(property)

	// Notify observers of change.
	tables.m_observable.SetChanged()
	tables.m_observable.NotifyObserversWithObject(property)

	return added, nil
}

/**
 * Remove an Actor Property entry.
 *
 * @param property The entry to remove.
 *
 * @return If the property is successfully removed, then <b>true</b>
 * will be returned. Otherwise, <b>false</b> will be returned.
 *
 * @throws MleRuntimeException This exception is thrown if the entry
 * is not an instance of <code>MleRTPropertyEntry</code>.
 */
func (tables *MleTables) removeActorProperty(property *MleRTPropertyEntry) (bool, *MleError) {
	retValue := false

	if !mle_util.InstanceOf(property, (*MleRTPropertyEntry)(nil)) {
		msg := "removeActorProperty: Not an Actor property."
		return false, NewMleError(msg, 0, nil)
	}
	tables.g_mleRTActorProperties.RemoveElement(property)

	// Notify observers of change.
	tables.m_observable.SetChanged()
	tables.m_observable.NotifyObserversWithObject(property)

	return retValue, nil
}

/**
 * Add a Set Property entry.
 *
 * @param property The entry to add.
 *
 * @return If the property is successfully added, then <b>true</b>
 * will be returned. Otherwise, <b>false</b> will be returned.
 *
 * @throws MleRuntimeException This exception is thrown if the entry
 * is not an instance of <code>MleRTPropertyEntry</code>.
 */
func (tables *MleTables) addSetProperty(property *MleRTPropertyEntry) (bool, *MleError) {
	added := true

	if !mle_util.InstanceOf(property, (*MleRTPropertyEntry)(nil)) {
		msg := "addSetProperty: Not a Set property."
		return false, NewMleError(msg, 0, nil)
	}
	tables.g_mleRTSetProperties.AddElement(property)

	// Notify observers of change.
	tables.m_observable.SetChanged()
	tables.m_observable.NotifyObserversWithObject(property)

	return added, nil
}

/**
 * Remove a Set Property entry.
 *
 * @param property The entry to remove.
 *
 * @return If the property is successfully removed, then <b>true</b>
 * will be returned. Otherwise, <b>false</b> will be returned.
 *
 * @throws MleRuntimeException This exception is thrown if the entry
 * is not an instance of <code>MleRTPropertyEntry</code>.
 */
func (tables *MleTables) removeSetProperty(property *MleRTPropertyEntry) (bool, *MleError) {
	retValue := false

	if !mle_util.InstanceOf(property, (*MleRTPropertyEntry)(nil)) {
		msg := "removeSetProperty: Not a Set property."
		return false, NewMleError(msg, 0, nil)
	}
	tables.g_mleRTSetProperties.RemoveElement(property)

	// Notify observers of change.
	tables.m_observable.SetChanged()
	tables.m_observable.NotifyObserversWithObject(property)

	return retValue, nil
}

/**
 * Add an Actor Class entry.
 *
 * @param clazz The entry to add.
 *
 * @return If the class is successfully added, then <b>true</b>
 * will be returned. Otherwise, <b>false</b> will be returned.
 *
 * @throws MleRuntimeException This exception is thrown if the entry
 * is not an instance of <code>MleRTActorClassEntry</code>.
 */
func (tables *MleTables) addActorClass(clazz *MleRTActorClassEntry) (bool, *MleError) {
	retValue := true

	if !mle_util.InstanceOf(clazz, (*MleRTActorClassEntry)(nil)) {
		msg := "addActorClass: Not a Actor class."
		return false, NewMleError(msg, 0, nil)
	}
	tables.g_mleRTActorClass.AddElement(clazz)

	// Notify observers of change.
	tables.m_observable.SetChanged()
	tables.m_observable.NotifyObserversWithObject(clazz)

	return retValue, nil
}

/**
 * Remove an Actor Class entry.
 *
 * @param clazz The entry to remove.
 *
 * @return If the class is successfully removed, then <b>true</b>
 * will be returned. Otherwise, <b>false</b> will be returned.
 *
 * @throws MleRuntimeException This exception is thrown if the entry
 * is not an instance of <code>MleRTActorClassEntry</code>.
 */
func (tables *MleTables) removeActorClass(clazz *MleRTActorClassEntry) (bool, *MleError) {
	retValue := true

	if !mle_util.InstanceOf(clazz, (*MleRTActorClassEntry)(nil)) {
		msg := "removeActorClass: Not a Actor class."
		return false, NewMleError(msg, 0, nil)
	}
	tables.g_mleRTActorClass.RemoveElement(clazz)

	// Notify observers of change.
	tables.m_observable.SetChanged()
	tables.m_observable.NotifyObserversWithObject(clazz)

	return retValue, nil
}

/**
 * Add a Role Class entry.
 *
 * @param clazz The entry to add.
 *
 * @return If the class is successfully added, then <b>true</b>
 * will be returned. Otherwise, <b>false</b> will be returned.
 *
 * @throws MleRuntimeException This exception is thrown if the entry
 * is not an instance of <code>MleRTRoleClassEntry</code>.
 */
func (tables *MleTables) addRoleClass(clazz *MleRTRoleClassEntry) (bool, *MleError) {
	retValue := true

	if !mle_util.InstanceOf(clazz, (*MleRTRoleClassEntry)(nil)) {
		msg := "addRoleClass: Not a Role class."
		return false, NewMleError(msg, 0, nil)
	}
	tables.g_mleRTRoleClass.AddElement(clazz)

	// Notify observers of change.
	tables.m_observable.SetChanged()
	tables.m_observable.NotifyObserversWithObject(clazz)

	return retValue, nil
}

/**
 * Remove a Role Class entry.
 *
 * @param clazz The entry to remove.
 *
 * @return If the class is successfully removed, then <b>true</b>
 * will be returned. Otherwise, <b>false</b> will be returned.
 *
 * @throws MleRuntimeException This exception is thrown if the entry
 * is not an instance of <code>MleRTRoleClassEntry</code>.
 */
func (tables *MleTables) removeRoleClass(clazz *MleRTRoleClassEntry) (bool, *MleError) {
	retValue := true

	if !mle_util.InstanceOf(clazz, (*MleRTRoleClassEntry)(nil)) {
		msg := "removeRoleClass: Not a Role class."
		return false, NewMleError(msg, 0, nil)
	}
	tables.g_mleRTRoleClass.RemoveElement(clazz)

	// Notify observers of change.
	tables.m_observable.SetChanged()
	tables.m_observable.NotifyObserversWithObject(clazz)

	return retValue, nil
}

/**
 * Add a Set Class entry.
 *
 * @param clazz The entry to add.
 *
 * @return If the class is successfully added, then <b>true</b>
 * will be returned. Otherwise, <b>false</b> will be returned.
 *
 * @throws MleRuntimeException This exception is thrown if the entry
 * is not an instance of <code>MleRTSetClassEntry</code>.
 */
func (tables *MleTables) addSetClass(clazz *MleRTSetClassEntry) (bool, *MleError) {
	retValue := true

	if !mle_util.InstanceOf(clazz, (*MleRTSetClassEntry)(nil)) {
		msg := "addSetClass: Not a Set class."
		return false, NewMleError(msg, 0, nil)
	}
	tables.g_mleRTSetClass.AddElement(clazz)

	// Notify observers of change.
	tables.m_observable.SetChanged()
	tables.m_observable.NotifyObserversWithObject(clazz)

	return retValue, nil
}

/**
 * Remove a Set Class entry.
 *
 * @param clazz The entry to remove.
 *
 * @return If the class is successfully removed, then <b>true</b>
 * will be returned. Otherwise, <b>false</b> will be returned.
 *
 * @throws MleRuntimeException This exception is thrown if the entry
 * is not an instance of <code>MleRTSetClassEntry</code>.
 */
func (tables *MleTables) removeSetClass(clazz *MleRTSetClassEntry) (bool, *MleError) {
	retValue := true

	if !mle_util.InstanceOf(clazz, (*MleRTSetClassEntry)(nil)) {
		msg := "removeSetClass: Not a Set class."
		return false, NewMleError(msg, 0, nil)
	}
	tables.g_mleRTSetClass.RemoveElement(clazz)

	// Notify observers of change.
	tables.m_observable.SetChanged()
	tables.m_observable.NotifyObserversWithObject(clazz)

	return retValue, nil
}

/**
 * Add a Group Class entry.
 *
 * @param clazz The entry to add.
 *
 * @return If the class is successfully added, then <b>true</b>
 * will be returned. Otherwise, <b>false</b> will be returned.
 *
 * @throws MleRuntimeException This exception is thrown if the entry
 * is not an instance of <code>MleRTGroupClassEntry</code>.
 */
func (tables *MleTables) addGroupClass(clazz *MleRTGroupClassEntry) (bool, *MleError) {
	retValue := true

	if !mle_util.InstanceOf(clazz, (*MleRTGroupClassEntry)(nil)) {
		msg := "addGroupClass: Not a Group class."
		return false, NewMleError(msg, 0, nil)
	}
	tables.g_mleRTGroupClass.AddElement(clazz)

	// Notify observers of change.
	tables.m_observable.SetChanged()
	tables.m_observable.NotifyObserversWithObject(clazz)

	return retValue, nil
}

/**
 * Remove a Group Class entry.
 *
 * @param clazz The entry to remove.
 *
 * @return If the class is successfully removed, then <b>true</b>
 * will be returned. Otherwise, <b>false</b> will be returned.
 *
 * @throws MleRuntimeException This exception is thrown if the entry
 * is not an instance of <code>MleRTGroupClassEntry</code>.
 */
func (tables *MleTables) removeGroupClass(clazz *MleRTGroupClassEntry) (bool, *MleError) {
	retValue := true

	if !mle_util.InstanceOf(clazz, (*MleRTGroupClassEntry)(nil)) {
		msg := "removeGroupClass: Not a Group class."
		return false, NewMleError(msg, 0, nil)
	}
	tables.g_mleRTGroupClass.RemoveElement(clazz)

	// Notify observers of change.
	tables.m_observable.SetChanged()
	tables.m_observable.NotifyObserversWithObject(clazz)

	return retValue, nil
}

/**
 * Add a Scene Class entry.
 *
 * @param clazz The entry to add.
 *
 * @return If the class is successfully added, then <b>true</b>
 * will be returned. Otherwise, <b>false</b> will be returned.
 *
 * @throws MleRuntimeException This exception is thrown if the entry
 * is not an instance of <code>MleRTScemeClassEntry</code>.
 */
func (tables *MleTables) addSceneClass(clazz *MleRTSceneClassEntry) (bool, *MleError) {
	retValue := true

	if !mle_util.InstanceOf(clazz, (*MleRTSceneClassEntry)(nil)) {
		msg := "addSceneClass: Not a Scene class."
		return false, NewMleError(msg, 0, nil)
	}
	tables.g_mleRTSceneClass.AddElement(clazz)

	// Notify observers of change.
	tables.m_observable.SetChanged()
	tables.m_observable.NotifyObserversWithObject(clazz)

	return retValue, nil
}

/**
 * Remove a Scene Class entry.
 *
 * @param clazz The entry to remove.
 *
 * @return If the class is successfully removed, then <b>true</b>
 * will be returned. Otherwise, <b>false</b> will be returned.
 *
 * @throws MleRuntimeException This exception is thrown if the entry
 * is not an instance of <code>MleRTSceneClassEntry</code>.
 */
func (tables *MleTables) removeSceneClass(clazz *MleRTSceneClassEntry) (bool, *MleError) {
	retValue := true

	if !mle_util.InstanceOf(clazz, (*MleRTSceneClassEntry)(nil)) {
		msg := "removeSceneClass: Not a Scene class."
		return false, NewMleError(msg, 0, nil)
	}
	tables.g_mleRTSceneClass.RemoveElement(clazz)

	// Notify observers of change.
	tables.m_observable.SetChanged()
	tables.m_observable.NotifyObserversWithObject(clazz)

	return retValue, nil
}

/**
 * Add a Set entry.
 *
 * @param set The Set entry to add.
 *
 * @return If the set is successfully added, then <b>true</b>
 * will be returned. Otherwise, <b>false</b> will be returned.
 *
 * @throws MleRuntimeException This exception is thrown if the entry
 * is not an instance of <code>MleRTSetEntry</code>.
 */
func (tables *MleTables) addSet(set *MleRTSetEntry) (bool, *MleError) {
	retValue := true

	if !mle_util.InstanceOf(set, (*MleRTSetEntry)(nil)) {
		msg := "addSet: Not a Set."
		return false, NewMleError(msg, 0, nil)
	}
	tables.g_mleRTSet.AddElement(set)

	// Notify observers of change.
	tables.m_observable.SetChanged()
	tables.m_observable.NotifyObserversWithObject(set)

	return retValue, nil
}

/**
 * Remove a Set entry.
 *
 * @param set The Set entry to remove.
 *
 * @return If the set is successfully removed, then <b>true</b>
 * will be returned. Otherwise, <b>false</b> will be returned.
 *
 * @throws MleRuntimeException This exception is thrown if the emtry
 * is not an instance of <code>MleRTSetClassEntry</code>.
 */
func (tables *MleTables) removeSet(set *MleRTSetEntry) (bool, *MleError) {
	retValue := true

	if !mle_util.InstanceOf(set, (*MleRTSetEntry)(nil)) {
		msg := "removeSet: Not a Set."
		return false, NewMleError(msg, 0, nil)
	}
	tables.g_mleRTSet.RemoveElement(set)

	// Notify observers of change.
	tables.m_observable.SetChanged()
	tables.m_observable.NotifyObserversWithObject(set)

	return retValue, nil
}

/**
 * Add a Media Reference entry.
 *
 * @param mediaref The Media Reference entry to add.
 *
 * @return If the media reference is successfully added, then <b>true</b>
 * will be returned. Otherwise, <b>false</b> will be returned.
 *
 * @throws MleRuntimeException This exception is thrown if the entry
 * is not an instance of <code>MleRTMediaRefEntry</code>.
 */
func (tables *MleTables) addMediaRef(mediaref *MleRTMediaRefEntry) (bool, *MleError) {
	retValue := true

	if !mle_util.InstanceOf(mediaref, (*MleRTMediaRefEntry)(nil)) {
		msg := "addMediaRef: Not a MediaRef."
		return false, NewMleError(msg, 0, nil)
	}
	tables.g_mleRTMediaRef.AddElement(mediaref)

	// Notify observers of change.
	tables.m_observable.SetChanged()
	tables.m_observable.NotifyObserversWithObject(mediaref)

	return retValue, nil
}

/**
 * Remove a Media Reference entry.
 *
 * @param mediaref The Media Reference entry to remove.
 *
 * @return If the media reference is successfully removed, then <b>true</b>
 * will be returned. Otherwise, <b>false</b> will be returned.
 *
 * @throws MleRuntimeException This exception is thrown if the entry
 * is not an instance of <code>MleRTMediaRefEntry</code>.
 */
func (tables *MleTables) removeMediaRef(mediaref *MleRTMediaRefEntry) (bool, *MleError) {
	retValue := true

	if !mle_util.InstanceOf(mediaref, (*MleRTMediaRefEntry)(nil)) {
		msg := "removeMediaRef: Not a MediaRef."
		return false, NewMleError(msg, 0, nil)
	}
	tables.g_mleRTMediaRef.RemoveElement(mediaref)

	// Notify observers of change.
	tables.m_observable.SetChanged()
	tables.m_observable.NotifyObserversWithObject(mediaref)

	return retValue, nil
}

/**
 * Register the specified Magic Lantern Object.
 * <p>
 * Note that this class does not actually manage the registration
 * of the Magic Lantern Object. Instead it passes the Object on to any
 * Observers so that they can determine whether to register the
 * Object or not.
 * </p>
 *
 * @param obj The Magic Lantern Object to register.
 */
func (tables *MleTables) RegisterObject(obj IMleObject) {
	if tables.m_observable.CountObservers() > 0 {
		// Notify observers of change.
		tables.m_observable.SetChanged()
		tables.m_observable.NotifyObserversWithObject(obj)
	}
}

/**
 * Unregister the specified Magic Lantern Object.
 * <p>
 * Note that this class does not actually manage the registration
 * of the Magic Lantern Object. Instead it passes the Object on to any
 * Observers so that they can determine whether to unregister the
 * Object or not.
 * </p>
 *
 * @param obj The Magic Lantern Object to unregister.
 */
func (tables *MleTables) UnregisterObject(obj IMleObject) {
	if tables.m_observable.CountObservers() > 0 {
		// Notify observers of change.
		tables.m_observable.SetChanged()
		tables.m_observable.NotifyObserversWithObject(obj)
	}
}
