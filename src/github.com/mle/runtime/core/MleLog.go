/*
 * @file MleLog.go
 * Created on June 14, 2019. (msm@wizzerworks.com)
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
	"bytes"
	"log"
)

// GMleLogger is the singleton instance of the Magic Lantern logger.
var GMleLogger *MleLog

// MleLog is a convenience class for Magic Lantern logging.
type MleLog struct {
	// The logger buffer.
	mBuf bytes.Buffer
	// A logger
	mLogger *log.Logger
}

// NewMleLog is a default constructor that will allocate a singleton
// instance of the Magic Lantern logger.
//
// Return
//   A reference to the global Magic Lantern logger is returned.
func NewMleLog() *MleLog {
	if GMleLogger == nil {
	    p := new(MleLog)
		p.mLogger = log.New(&p.mBuf, "MLE: ", log.Lshortfile)
		GMleLogger = p
	}

	return GMleLogger
}

// Log may be used to log a generic string with the logger.
//
// Parameters
//   msg - The message to log.
func (l *MleLog) Log(msg string) {
	l.mLogger.Print(msg)
}

// Info may be used to log informational messages with the logger.
//
// Parameters
//   msg - The message to log.
func (l *MleLog) Info(msg string) {
	str := "INFO: " + msg
	l.mLogger.Print(str)
}

// Warn may be used to log warning messages with the logger.
//
// Parameters
//   msg - The message to log.
func (l *MleLog) Warn(msg string) {
	str := "WARN: " + msg
	l.mLogger.Print(str)
}

// Error may be used to log error messages with the logger.
//
// Parameters
//   msg - The message to log.
func (l *MleLog) Error(msg string) {
	str := "ERROR: " + msg
	l.mLogger.Print(str)
}

// MleLogInfo may be used to log informational messages with the
// global Magic Lantern logger.
//
// Parameters
//   msg - The message to log.
func MleLogInfo(msg string, print bool) {
	var logger = NewMleLog()
	logger.Info(msg)

	if print {
		fmt.Print(logger.mBuf)
	}
}

// MleLogWarn may be used to log warning messages with the
// global Magic Lantern logger.
//
// Parameters
//   msg - The message to log.
func MleLogWarn(msg string, print bool) {
	var logger = NewMleLog()
	logger.Warn(msg)

	if print {
		fmt.Print(logger.mBuf)
	}
}

// MleLogError may be used to log informational messages with the
// global Magic Lantern logger.
//
// Parameters
//   msg - The message to log.
func MleLogError(msg string, print bool) {
	var logger = NewMleLog()
	logger.Error(msg)

	if print {
		fmt.Print(logger.mBuf)
	}
}
