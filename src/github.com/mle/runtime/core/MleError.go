/**
 * @file MleError.go
 * Created on April 30, 2019. (msm@wizzerworks.com)
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
	"fmt"
	"time"
)

// Import Magic Lantern packages.

// MleError is an error implementation that includes a timestamp, message,
// and code identifier.
type MleError struct {
	When  time.Time // A timestamp for when the error occurred.
	What  string    // A JSON formatted response "{ code: <value>, message: <string> }"
	Value int       // The HTTP status code
	Err   error     // An internal error
}

// NewError constructs a MleError.
func NewMleError(msg string, value int, err error) *MleError {
	p := new(MleError)
	p.What = msg
	p.Value = value
	p.Err = err
	p.When = time.Now()
	return p
}

func (e MleError) Error() string {
	return fmt.Sprintf("%v: %v - %v", e.Value, e.When, e.What)
}
