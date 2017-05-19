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

package packing

import (
	gethAbi "github.com/ethereum/go-ethereum/accounts/abi"
)

// NOTE: this does not work for nested nested types; because
// as far as Ive looked there is no differentiator per bug 
// in go-ethereum/abi
func isNestedType(inputType *gethAbi.Type) bool {
	// type is nested if it is either a slice "[]";
	// or an array "[n]"
	return ((inputType.IsSlice || inputType.IsArray) &&
		// but exceptions are the type "bytes" 
		!(inputType.T == gethAbi.BytesTy ||
		// or fixed length bytes like "bytes1 .. bytes32"
		inputType.T == gethAbi.FixedBytesTy ||
		// or function identifiers which are effectively "bytes24"
		inputType.T == gethAbi.FunctionTy))
}