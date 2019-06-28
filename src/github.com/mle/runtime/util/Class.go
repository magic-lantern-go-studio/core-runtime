/**
 * @file Class.go
 * Created on May 24, 2019. (msm@wizzerworks.com)
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
package util

// Import go packages.
import (
	"fmt"
	"reflect"
)

// GClassRegistry is a global collection of classes.
//
// Since go does not have a "Class" object to determine available classes that can
// be retrieved using a method like "forName", we create a global list of registered
// "classes".
var GClassRegistry map[string]interface{}

// MethodExists will determine if a method exists on a specified interface.
//
// Parameters
//   any  - The interface to determine if the named method exists on.
//   name - The name of the interface to test.
//
// Return
//   true is returned if the named method exists. Otherwise, false will be
//   returned.
func MethodExists(any interface{}, name string) bool {
	method := reflect.ValueOf(any).MethodByName(name)
	if method.IsValid() {
		return true
	}
	return false
}

// FieldExists will determine if a field exists on a specified interface.
//
// Parameters
//   any  - The interface to determine if the named field exists on.
//   name - The name of the interface to test.
//
// Return
//   true is returned if the named field exists. Otherwise, false will be
//   returned.
func FieldExists(any interface{}, name string) bool {
	field := reflect.ValueOf(any).FieldByName(name)
	if field.IsValid() {
		return true
	}
	return false
}

// Invoke is used to call a method on a specified interface.
//
// Parameters
//   any - The interface to invoke the method on.
//   name - The name of the method to invoke.
//   args - The list of parameters to the method being invoked.
//
// Return
//   The return values of the invoked method will be returned.
//
// Example
//   firstResult, err := Invoke(AnyStructInterface, MethodName, Params...)
func Invoke(any interface{}, name string, args ...interface{}) (reflect.Value, error) {
	method := reflect.ValueOf(any).MethodByName(name)
	methodType := method.Type()
	numIn := methodType.NumIn()
	if numIn > len(args) {
		return reflect.ValueOf(nil), fmt.Errorf("Method %s must have minimum %d params. Have %d", name, numIn, len(args))
	}
	if numIn != len(args) && !methodType.IsVariadic() {
		return reflect.ValueOf(nil), fmt.Errorf("Method %s must have %d params. Have %d", name, numIn, len(args))
	}
	in := make([]reflect.Value, len(args))
	for i := 0; i < len(args); i++ {
		var inType reflect.Type
		if methodType.IsVariadic() && i >= numIn-1 {
			inType = methodType.In(numIn - 1).Elem()
		} else {
			inType = methodType.In(i)
		}
		argValue := reflect.ValueOf(args[i])
		if !argValue.IsValid() {
			return reflect.ValueOf(nil), fmt.Errorf("Method %s. Param[%d] must be %s. Have %s", name, i, inType, argValue.String())
		}
		argType := argValue.Type()
		if argType.ConvertibleTo(inType) {
			in[i] = argValue.Convert(inType)
		} else {
			return reflect.ValueOf(nil), fmt.Errorf("Method %s. Param[%d] must be %s. Have %s", name, i, inType, argType)
		}
	}
	return method.Call(in)[0], nil
}

// InstanceOf determines if an interface is an instance of, or type of, a specified class.
//
// Parameters
//   objectPtr - A reference to the object being tested.
//   typePtr   - A reference to the type to compare against.
//
// Return
//   true will be returned if the objectPtr is an instance of the typePtr.
//   Otherwise false will be returned.
func InstanceOf(objectPtr interface{}, typePtr interface{}) bool {
	aType := reflect.TypeOf(objectPtr)
	bType := reflect.TypeOf(typePtr)
	if aType == bType {
		return true
	}
	return false
}
