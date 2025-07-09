// Copyright 2016 Peter Mattis.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License. See the AUTHORS file
// for names of contributors.

//go:build (amd64 || amd64p32) && gc && go1.5
// +build amd64 amd64p32
// +build gc
// +build go1.5

package goid

// func Get() int64


//go:build !windows && !plan9
// +build !windows,!plan9

package pid

import (
    "bytes"
    "runtime"
    "strconv"
)

// Get returns the ID of the current goroutine.
// NOTE: This is a hack and not officially supported by Go.
func Get() int {
    var buf [64]byte
    n := runtime.Stack(buf[:], false)
    stack := bytes.TrimPrefix(buf[:n], []byte("goroutine "))
    idField := stack[:bytes.IndexByte(stack, ' ')]
    id, err := strconv.Atoi(string(idField))
    if err != nil {
        return -1
    }
    return id
}


