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
	"fmt"
	"math/big"
	"reflect"

	gethAbi "github.com/ethereum/go-ethereum/accounts/abi"
)

// inferType returns the golang type that matches the abi type,
// a boolean that specifies whether it is an array or slice, uint for size
// 
func inferType(abiType *gethAbi.Type) (reflect.Type, error) {

	if isNestedType(abiType) {

	}

	if abiType.Elem != nil {
		return inferType(abiType.Elem)
	}


	switch abiType.T {
	// for integers Type.
	// NOTE: for nested type T is not initialised and falsely defaulted to 0 IntTy
	case gethAbi.IntTy:
		return inferIntegerType(true, abiType.Size)
	case gethAbi.UintTy:
		return inferIntegerType(false, abiType.Size)
	case gethAbi.BoolTy:
		return reflect.TypeOf(false), nil
	case gethAbi.StringTy:
		return reflect.TypeOf(""), nil
	// NOTE: SliceTy is never assigned by go-eth/abi
	// case gethAbi.SliceTy:

	// case gethAbi.AddressTy:
	// case gethAbi.FixedBytesTy:
	// case gethAbi.BytesTy:
	// case gethAbi.HashTy:
	// case gethAbi.FixedpointTy:
	// case gethAbi.FunctionTy:
	default:

	}
	return nil, fmt.Errorf("MARMOT IN PROGRESS")
}

func inferIntegerType(signed bool, varSize int) (reflect.Type, error) {
	if varSize > 256 {
		return nil, fmt.Errorf("Failed to infer integer for size %d (max 256 bits)",
			varSize)
	}
	if varSize < 1 {
		return nil, fmt.Errorf("Failed to infer integer for size %d (must be strictly positive)",
			varSize)
	}
	
	switch varSize {
	case 8:
		if signed {
			return reflect.TypeOf(int8(0)), nil
		} else {
			return reflect.TypeOf(uint8(0)), nil
		}
	case 16:
		if signed {
			return reflect.TypeOf(int16(0)), nil
		} else {
			return reflect.TypeOf(uint16(0)), nil
		}
	case 32:
		if signed {
			return reflect.TypeOf(int32(0)), nil
		} else {
			return reflect.TypeOf(uint32(0)), nil
		}
	case 64:
		if signed {
			return reflect.TypeOf(int64(0)), nil
		} else {
			return reflect.TypeOf(uint64(0)), nil
		}
default:
		return reflect.TypeOf(big.Int{}), nil
	}
}

// inferRank determines whether the type is scalar or vector valued
// and returns zero for scalar, one for vector, or error if unhandled.
func inferRank(abiType *gethAbi.Type) (uint, error) {

	if isNestedType(abiType) {

	}
	return 0, nil
}