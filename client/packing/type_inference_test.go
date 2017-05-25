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
	"math/big"
	"reflect"
	"testing"

	gethAbi "github.com/ethereum/go-ethereum/accounts/abi"
)

func TestInference(t *testing.T) {
	inferenceTests := []struct{
		abiKind      string
		goType       reflect.Type
		rank         uint
		multiplicity uint
	}{
		{"int", reflect.TypeOf(big.Int{}), 0, 1},
		{"int8", reflect.TypeOf(int8(0)), 0, 1},
		{"int16", reflect.TypeOf(int16(0)), 0, 1},
		{"int32", reflect.TypeOf(int32(0)), 0, 1},
		{"int64", reflect.TypeOf(int64(0)), 0, 1},
		{"int256", reflect.TypeOf(big.Int{}), 0, 1},
		{"uint", reflect.TypeOf(big.Int{}), 0, 1},
		{"uint8", reflect.TypeOf(uint8(0)), 0, 1},
		{"uint16", reflect.TypeOf(uint16(0)), 0, 1},
		{"uint32", reflect.TypeOf(uint32(0)), 0, 1},
		{"uint64", reflect.TypeOf(uint64(0)), 0, 1},
		{"uint256", reflect.TypeOf(big.Int{}), 0, 1},
		{"string", reflect.TypeOf(""), 0, 1},

		{"uint[2]", reflect.TypeOf(big.Int{}), 1, 2},
		{"uint16[2]", reflect.TypeOf(uint16(0)), 1, 2},
		{"uint16[]", reflect.TypeOf(uint16(0)), 1, 0},

		{"bytes", reflect.TypeOf(uint8(0)), 1, 0},
		{"bytes4", reflect.TypeOf(uint8(0)), 1, 4},
		{"bytes32", reflect.TypeOf(uint8(0)), 1, 32},
	}

	for _, infer := range inferenceTests {
		abiType, err := gethAbi.NewType(infer.abiKind)
		if err != nil {
			t.Errorf("Failed to create abi type for %s: %s", infer.abiKind, err)
		}
		inferedType, err := inferType(&abiType)
		if err != nil {
			t.Errorf("Failed to infer type from %s: %s", abiType.String(), err)
		}
		if !reflect.DeepEqual(infer.goType, inferedType) {
			t.Errorf("Infered type %s does not match expected type %s",
				inferedType, infer.goType)
		}
	}
}