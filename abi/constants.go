// Copyright 2017 Monax Industries Limited
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package abi

type IntegerSize uint

// IntSize defines the available integer sizes in bytes for ABI
// integer types
const (
	IntSize8   IntegerSize = 8
	IntSize16  IntegerSize = 16
	IntSize32  IntegerSize = 32
	IntSize64  IntegerSize = 64
	IntSize256 IntegerSize = 256
)

const (
	TypeStringInt    = "int"
	TypeStringUint   = "uint"
	TypeStringBytes  = "bytes"
	TypeStringString = "string"
)
