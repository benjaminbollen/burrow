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

import (
	"fmt"
	"math"
	"math/big"
	"strconv"

	"github.com/hyperledger/burrow/common/sanity"
	"github.com/hyperledger/burrow/word256"
)

type AbiUint struct {
	representation interface{}
	size           IntegerSize
}

func NewAbiUint(size IntegerSize) (*AbiUint, error) {
	newAbiUint := new(AbiUint)

	switch size {
	case IntSize8:
		newAbiUint.representation = interface{}(uint8(0))
		newAbiUint.size = IntSize8
	case IntSize16:
		newAbiUint.representation = interface{}(uint16(0))
		newAbiUint.size = IntSize16
	case IntSize32:
		newAbiUint.representation = interface{}(uint32(0))
		newAbiUint.size = IntSize32
	case IntSize64:
		newAbiUint.representation = interface{}(uint64(0))
		newAbiUint.size = IntSize64
	case IntSize256:
		newAbiUint.representation = interface{}(big.NewInt(0))
		newAbiUint.size = IntSize256
	default:
		return nil, fmt.Errorf("Failed to initiate ABI unsigned integer for size %d", size)
	}
	return newAbiUint, nil
}

// TODO: accept hex encoded string and big integer to set values
func (abiUint *AbiUint) Set(value interface{}) error {
	if abiUint.representation == nil {
		sanity.PanicSanity("ABI unsigned integer must always have a representation for Set().")
	}

	switch abiUint.representation.(type) {
	case uint8:
		convertedValue, err := convertToUint8(value)
		if err != nil {
			return err
		}
		abiUint.representation = convertedValue
	case uint16:
		convertedValue, err := convertToUint16(value)
		if err != nil {
			return err
		}
		abiUint.representation = convertedValue
	case uint32:
		convertedValue, err := convertToUint32(value)
		if err != nil {
			return err
		}
		abiUint.representation = convertedValue
	case uint64:
		convertedValue, err := convertToUint64(value)
		if err != nil {
			return err
		}
		abiUint.representation = convertedValue
	case *big.Int:
		convertedValue, err := convertToUint256(value)
		if err != nil {
			return err
		}
		abiUint.representation = convertedValue
	default:
		sanity.PanicSanity("ABI unsigned integer represented by unhandled type for Set().")
	}
	return nil
}

// Bytes returns the word256 byte representation of Uint
func (abiUint *AbiUint) Bytes() []byte {
	if abiUint.representation == nil {
		sanity.PanicSanity("ABI unsigned integer must always have a representation for Bytes().")
	}

	switch t := abiUint.representation.(type) {
	case uint8:
		return word256.Uint64ToWord256(uint64(t)).Bytes()
	case uint16:
		return word256.Uint64ToWord256(uint64(t)).Bytes()
	case uint32:
		return word256.Uint64ToWord256(uint64(t)).Bytes()
	case uint64:
		return word256.Uint64ToWord256(t).Bytes()
	case *big.Int:
		return word256.Uint256ToWord256(t).Bytes()
	default:
		sanity.PanicSanity("ABI unsigned integer represented by unhandled type for Bytes().")
	}
	return nil
}

func (abiUint *AbiUint) String() string {
	return fmt.Sprintf("%s%d", TypeStringUint, abiUint.size)
}

func checkIntegerUint(t float64) (uint64, error) {
	var y uint64
	if math.Abs(t) <= math.MaxUint64 && t >= 0 {
		y = uint64(t)
		if t == float64(y) {
			return y, nil
		} else {
			return 0, fmt.Errorf("Failed to convert float64 %v to uint: non-integer value", t)
		}
	} else {
		return 0, fmt.Errorf("Failed to convert float to uint: bigger than max uint64 or negative")
	}
}

func checkOverflowUint8(x uint8, t uint64) (uint8, error) {
	if uint64(x) == t {
		return x, nil
	} else {
		return 0, fmt.Errorf("Failed to convert int %v to uint8: overflow", t)
	}
}

func checkOverflowUint16(x uint16, t uint64) (uint16, error) {
	if uint64(x) == t {
		return x, nil
	} else {
		return 0, fmt.Errorf("Failed to convert int %v to uint16: overflow", t)
	}
}

func checkOverflowUint32(x uint32, t uint64) (uint32, error) {
	if uint64(x) == t {
		return x, nil
	} else {
		return 0, fmt.Errorf("Failed to convert int %v to uint32: overflow", t)
	}
}

func checkUnderflowUint64(x uint64, t int64) (uint64, error) {
	if t >= 0 {
		return x, nil
	} else {
		return 0, fmt.Errorf("Failed to convert int %v to uint64: underflow", t)
	}
}

func checkUnderflowUint256(x uint64, t int64) (*big.Int, error) {
	if t >= 0 {
		return new(big.Int).SetUint64(x), nil
	} else {
		return nil, fmt.Errorf("Failed to convert int %v to uint256: underflow", t)
	}
}

func convertToUint8(value interface{}) (uint8, error) {
	switch t := value.(type) {
	case int:
		return checkOverflowUint8(uint8(t), uint64(t))
	case int8:
		return checkOverflowUint8(uint8(t), uint64(t))
	case int16:
		return checkOverflowUint8(uint8(t), uint64(t))
	case int32:
		return checkOverflowUint8(uint8(t), uint64(t))
	case int64:
		return checkOverflowUint8(uint8(t), uint64(t))
	case uint:
		return checkOverflowUint8(uint8(t), uint64(t))
	case uint8:
		return t, nil
	case uint16:
		return checkOverflowUint8(uint8(t), uint64(t))
	case uint32:
		return checkOverflowUint8(uint8(t), uint64(t))
	case uint64:
		return checkOverflowUint8(uint8(t), t)
	case float32:
		y, err := checkIntegerUint(float64(t))
		if err != nil {
			return 0, err
		}
		return convertToUint8(y)
	case float64:
		y, err := checkIntegerUint(t)
		if err != nil {
			return 0, err
		}
		return convertToUint8(y)
	case string:
		// NOTE: add hex
		// do not allow for floating point conversions,
		// as rounding may be lossy
		y, err := strconv.ParseUint(t, 10, 64)
		if err != nil {
			return 0, err
		}
		return convertToUint8(y)
	case bool:
		if t {
			return uint8(1), nil
		} else {
			return uint8(0), nil
		}
	case complex64, complex128:
		return 0, fmt.Errorf("Failed to convert complex type to uint")
	default:
		return 0, fmt.Errorf("Failed to convert unhandled type to uint")
	}
}

func convertToUint16(value interface{}) (uint16, error) {
	switch t := value.(type) {
	case int:
		return checkOverflowUint16(uint16(t), uint64(t))
	case int8:
		return checkOverflowUint16(uint16(t), uint64(t))
	case int16:
		return checkOverflowUint16(uint16(t), uint64(t))
	case int32:
		return checkOverflowUint16(uint16(t), uint64(t))
	case int64:
		return checkOverflowUint16(uint16(t), uint64(t))
	case uint:
		return checkOverflowUint16(uint16(t), uint64(t))
	case uint8:
		return uint16(t), nil
	case uint16:
		return t, nil
	case uint32:
		return checkOverflowUint16(uint16(t), uint64(t))
	case uint64:
		return checkOverflowUint16(uint16(t), t)
	case float32:
		y, err := checkIntegerUint(float64(t))
		if err != nil {
			return 0, err
		}
		return convertToUint16(y)
	case float64:
		y, err := checkIntegerUint(t)
		if err != nil {
			return 0, err
		}
		return convertToUint16(y)
	case string:
		// NOTE: add hex
		// do not allow for floating point conversions,
		// as rounding may be lossy
		y, err := strconv.ParseUint(t, 10, 64)
		if err != nil {
			return 0, err
		}
		return convertToUint16(y)
	case bool:
		if t {
			return uint16(1), nil
		} else {
			return uint16(0), nil
		}
	case complex64, complex128:
		return 0, fmt.Errorf("Failed to convert complex type to uint")
	default:
		return 0, fmt.Errorf("Failed to convert unhandled type to uint16")
	}
}

func convertToUint32(value interface{}) (uint32, error) {
	switch t := value.(type) {
	case int:
		return checkOverflowUint32(uint32(t), uint64(t))
	case int8:
		return checkOverflowUint32(uint32(t), uint64(t))
	case int16:
		return checkOverflowUint32(uint32(t), uint64(t))
	case int32:
		return checkOverflowUint32(uint32(t), uint64(t))
	case int64:
		return checkOverflowUint32(uint32(t), uint64(t))
	case uint:
		return checkOverflowUint32(uint32(t), uint64(t))
	case uint8:
		return uint32(t), nil
	case uint16:
		return uint32(t), nil
	case uint32:
		return t, nil
	case uint64:
		return checkOverflowUint32(uint32(t), uint64(t))
	case float32:
		y, err := checkIntegerUint(float64(t))
		if err != nil {
			return 0, err
		}
		return convertToUint32(y)
	case float64:
		y, err := checkIntegerUint(t)
		if err != nil {
			return 0, err
		}
		return convertToUint32(y)
	case string:
		// NOTE: add hex
		// do not allow for floating point conversions,
		// as rounding may be lossy
		y, err := strconv.ParseUint(t, 10, 64)
		if err != nil {
			return 0, err
		}
		return convertToUint32(y)
	case bool:
		if t {
			return uint32(1), nil
		} else {
			return uint32(0), nil
		}
	case complex64, complex128:
		return 0, fmt.Errorf("Failed to convert complex type to uint32")
	default:
		return 0, fmt.Errorf("Failed to convert unhandled type to uint32")
	}
}

func convertToUint64(value interface{}) (uint64, error) {
	switch t := value.(type) {
	case int:
		return checkUnderflowUint64(uint64(t), int64(t))
	case int8:
		return checkUnderflowUint64(uint64(t), int64(t))
	case int16:
		return checkUnderflowUint64(uint64(t), int64(t))
	case int32:
		return checkUnderflowUint64(uint64(t), int64(t))
	case int64:
		return checkUnderflowUint64(uint64(t), t)
	case uint:
		return uint64(t), nil
	case uint8:
		return uint64(t), nil
	case uint16:
		return uint64(t), nil
	case uint32:
		return uint64(t), nil
	case uint64:
		return t, nil
	case float32:
		return checkIntegerUint(float64(t))
	case float64:
		return checkIntegerUint(t)
	case string:
		// NOTE: add hex
		// do not allow for floating point conversions,
		// as rounding may be lossy
		return strconv.ParseUint(t, 10, 64)
	case bool:
		if t {
			return uint64(1), nil
		} else {
			return uint64(0), nil
		}
	case complex64, complex128:
		return 0, fmt.Errorf("Failed to convert complex type to uint64")
	default:
		return 0, fmt.Errorf("Failed to convert unhandled type to uint64")
	}
}

func convertToUint256(value interface{}) (*big.Int, error) {
	switch t := value.(type) {
	case int:
		return checkUnderflowUint256(uint64(t), int64(t))
	case int8:
		return checkUnderflowUint256(uint64(t), int64(t))
	case int16:
		return checkUnderflowUint256(uint64(t), int64(t))
	case int32:
		return checkUnderflowUint256(uint64(t), int64(t))
	case int64:
		return checkUnderflowUint256(uint64(t), t)
	case uint:
		return new(big.Int).SetUint64(uint64(t)), nil
	case uint8:
		return new(big.Int).SetUint64(uint64(t)), nil
	case uint16:
		return new(big.Int).SetUint64(uint64(t)), nil
	case uint32:
		return new(big.Int).SetUint64(uint64(t)), nil
	case uint64:
		return new(big.Int).SetUint64(t), nil
	case *big.Int:
		if err := word256.CheckOverflowUint256(t); err != nil {
			return nil, err
		} else {
			return word256.U256(t), nil
		}
	case float32:
		y, err := checkIntegerUint(float64(t))
		if err != nil {
			return nil, err
		}
		return convertToUint256(y)
	case float64:
		y, err := checkIntegerUint(t)
		if err != nil {
			return nil, err
		}
		return convertToUint256(y)
	case string:
		// NOTE: add hex
		// do not allow for floating point conversions,
		// as rounding may be lossy
		y, err := strconv.ParseUint(t, 10, 64)
		if err != nil {
			return nil, err
		}
		return convertToUint256(y)
	case bool:
		if t {
			return big.NewInt(1), nil
		} else {
			return big.NewInt(0), nil
		}
	case complex64, complex128:
		return nil, fmt.Errorf("Failed to convert complex type to uint256")
	default:
		return nil, fmt.Errorf("Failed to convert unhandled type to uint256")
	}
}
