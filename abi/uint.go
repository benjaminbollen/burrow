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
		i := new(big.Int)
		newAbiUint.representation = interface{}(i)
		newAbiUint.size = IntSize256
	default:
		return nil, fmt.Errorf("Failed to initiate ABI unsigned integer for size %d", size)
	}
	return newAbiUint, nil
}

func (abiUint *AbiUint) Set(value interface{}) error {
	if abiUint.representation == nil {
		sanity.PanicSanity("ABI unsigned integer must always have a representation.")
	}

	switch abiUint.representation.(type) {
	case uint8:
	case uint16:
	case uint32:
	case uint64:
	case *big.Int:
	}
	return nil
}

func (abiUint *AbiUint) Bytes() ([]byte, error) {
	return nil, nil
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
		return 0, fmt.Errorf("Failed to reduce int %v to uint8: overflow", t)
	}
}

func checkOverflowUint16(x uint16, t uint64) (uint16, error) {
	if uint64(x) == t {
		return x, nil
	} else {
		return 0, fmt.Errorf("Failed to reduce int %v to uint16: overflow", t)
	}
}

func checkOverflowUint32(x uint32, t uint64) (uint32, error) {
	if uint64(x) == t {
		return x, nil
	} else {
		return 0, fmt.Errorf("Failed to reduce int %v to uint32: overflow", t)
	}
}

func checkOverflowUint64(x uint64, t int64) (uint64, error) {
	if t >= 0 {
		return x, nil
	} else {
		return 0, fmt.Errorf("Failed to reduce int %v to uint64: overflow", t)
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
		return checkOverflowUint8(uint8(t), uint64(t))
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
		return checkOverflowUint16(uint16(t), uint64(t))
	case uint16:
		return t, nil
	case uint32:
		return checkOverflowUint16(uint16(t), uint64(t))
	case uint64:
		return checkOverflowUint16(uint16(t), uint64(t))
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
		return 0, fmt.Errorf("Failed to convert unhandled type to uint")
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
		return checkOverflowUint32(uint32(t), uint64(t))
	case uint16:
		return checkOverflowUint32(uint32(t), uint64(t))
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
		return 0, fmt.Errorf("Failed to convert complex type to uint")
	default:
		return 0, fmt.Errorf("Failed to convert unhandled type to uint")
	}
}

func convertToUint64(value interface{}) (uint64, error) {
	switch t := value.(type) {
	case int:
		return checkOverflowUint64(uint64(t), int64(t))
	case int8:
		return checkOverflowUint64(uint64(t), int64(t))
	case int16:
		return checkOverflowUint64(uint64(t), int64(t))
	case int32:
		return checkOverflowUint64(uint64(t), int64(t))
	case int64:
		return checkOverflowUint64(uint64(t), t)
	case uint:
		return checkOverflowUint64(uint64(t), int64(t))
	case uint8:
		return checkOverflowUint64(uint64(t), int64(t))
	case uint16:
		return checkOverflowUint64(uint64(t), int64(t))
	case uint32:
		return checkOverflowUint64(uint64(t), int64(t))
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
		return 0, fmt.Errorf("Failed to convert complex type to uint")
	default:
		return 0, fmt.Errorf("Failed to convert unhandled type to uint")
	}
}
