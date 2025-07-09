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

import (
    "bytes"
    "runtime"
    "strconv"
)

func GetGoID() int64 {
    var buf [64]byte
    n := runtime.Stack(buf[:], false)
    // 栈信息格式类似： "goroutine 123 [running]:\n"
    fields := bytes.Fields(buf[:n])
    if len(fields) < 2 {
        return -1
    }
    id, err := strconv.ParseInt(string(fields[1]), 10, 64)
    if err != nil {
        return -1
    }
    return id
}

