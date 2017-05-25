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
	"reflect"

	gethAbi "github.com/ethereum/go-ethereum/accounts/abi"
)

// A valued argument holds a value for a specific argument
type ValuedArgument struct{
	Name string

	// value holds the value after succcesful conversion
	// to the required type.  If nil, value is not yet set,
	// or failed to be set.
	Value interface{}

	// infered base type is the base golang type
	// required for the argument to be before packing
	InferedBaseType reflect.Type

	// Rank is supported for 0 scalar and 1 vectors
	Rank uint
	// Multiplicity is the dimension of the vector;
	// for a vector the multiplicity is zero if undetermined
	// for scalar the multiplicity is always one
	Multiplicity uint
}

func NewValuedArgument(name string, inputType *gethAbi.Type) (*ValuedArgument, error) {
	inferedType, err := inferType(inputType)
	if err != nil {
		return nil, fmt.Errorf("Failed to infer base type for %s: %s", inputType, err)
	}
	inferedRank, err := inferRank(inputType)
	if err != nil {
		return nil, fmt.Errorf("Failed to infer rank of base type for %s: %s", inputType, err)
	}

	return &ValuedArgument{
			Name: name,
			Value: nil,
			InferedBaseType: inferedType,
			Rank: inferedRank,
			Multiplicity: 1,
		}, nil
}