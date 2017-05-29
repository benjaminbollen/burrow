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
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUintString(t *testing.T) {
	abiUint8, err := NewAbiUint(8)
	if assert.NoError(t, err) {
		assert.Equal(t, abiUint8.String(), "uint8")
	}	
	abiUint16, err := NewAbiUint(16)
	if assert.NoError(t, err) {
		assert.Equal(t, abiUint16.String(), "uint16")
	}

	abiUint32, err := NewAbiUint(32)
	if assert.NoError(t, err) {
		assert.Equal(t, abiUint32.String(), "uint32")
	}

	abiUint64, err := NewAbiUint(64)
	if assert.NoError(t, err) {
		assert.Equal(t, abiUint64.String(), "uint64")
	}

	abiUint256, err := NewAbiUint(256)
	if assert.NoError(t, err) {
		assert.Equal(t, abiUint256.String(), "uint256")
	}

	_, err = NewAbiUint(9)
	assert.Error(t, err)
}

func TestUint8Conversion(t *testing.T) {
	conversionUint8Tests := []struct {
		shouldSucceed bool
		intention     uint8
		input         interface{}
	}{
		{true, 1, 1},
		{true, 1, "1"},
		{true, 1, float32(1.0)},
		
		{true, 2, uint(2)},
		{true, 3, uint8(3)},
		{true, 4, uint16(4)},
		{true, 5, uint32(5)},
		{true, 6, uint64(6)},

		{true, 2, int(2)},
		{true, 3, int8(3)},
		{true, 4, int16(4)},
		{true, 5, int32(5)},
		{true, 6, int64(6)},
		
		{false, 0, uint(math.MaxUint8 + 1)},
		{false, 0, int8(-3)},

		{false, 1, float64(1.0001)},
		{false, 1, "1.0"},
	}

	for _, conversion := range conversionUint8Tests {
		if conversion.shouldSucceed {
			output, err := convertToUint8(conversion.input)
			if assert.NoError(t, err) {
				assert.Equal(t, conversion.intention, output, "Failed to convert %v to uint8: %s", conversion.input, err)
			}
		} else {
			_, err := convertToUint8(conversion.input)
			assert.Error(t, err)
		}
	}
}

func TestUint16Conversion(t *testing.T) {
	conversionUint16Tests := []struct {
		shouldSucceed bool
		intention     uint16
		input         interface{}
	}{
		{true, 1, 1},
		{true, 1, "1"},
		{true, 1, float32(1.0)},
		
		{true, 2, uint(2)},
		{true, 3, uint8(3)},
		{true, 4, uint16(4)},
		{true, 5, uint32(5)},
		{true, 6, uint64(6)},

		{true, 2, int(2)},
		{true, 3, int8(3)},
		{true, 4, int16(4)},
		{true, 5, int32(5)},
		{true, 6, int64(6)},
		
		{false, 0, uint(math.MaxUint16 + 1)},
		{false, 0, int16(-3)},

		{false, 1, float64(1.0001)},
		{false, 1, "1.0"},
	}

	for _, conversion := range conversionUint16Tests {
		if conversion.shouldSucceed {
			output, err := convertToUint16(conversion.input)
			if assert.NoError(t, err) {
				assert.Equal(t, conversion.intention, output, "Failed to convert %v to uint16: %s", conversion.input, err)
			}
		} else {
			_, err := convertToUint16(conversion.input)
			assert.Error(t, err)
		}
	}
}

func TestUint32Conversion(t *testing.T) {
	conversionUint32Tests := []struct {
		shouldSucceed bool
		intention     uint32
		input         interface{}
	}{
		{true, 1, 1},
		{true, 1, "1"},
		{true, 1, float32(1.0)},
		
		{true, 2, uint(2)},
		{true, 3, uint8(3)},
		{true, 4, uint16(4)},
		{true, 5, uint32(5)},
		{true, 6, uint64(6)},

		{true, 2, int(2)},
		{true, 3, int8(3)},
		{true, 4, int16(4)},
		{true, 5, int32(5)},
		{true, 6, int64(6)},
		
		{false, 0, uint(math.MaxUint32 + 1)},
		{false, 0, int32(-3)},

		{false, 1, float64(1.0001)},
		{false, 1, "1.0"},
	}

	for _, conversion := range conversionUint32Tests {
		if conversion.shouldSucceed {
			output, err := convertToUint32(conversion.input)
			if assert.NoError(t, err) {
				assert.Equal(t, conversion.intention, output, "Failed to convert %v to uint32: %s", conversion.input, err)
			}
		} else {
			_, err := convertToUint32(conversion.input)
			assert.Error(t, err)
		}
	}
}

func TestUint64Conversion(t *testing.T) {
	conversionUint64Tests := []struct {
		shouldSucceed bool
		intention     uint64
		input         interface{}
	}{
		{true, 1, 1},
		{true, 1, "1"},
		{true, 1, float32(1.0)},
		
		{true, 2, uint(2)},
		{true, 3, uint8(3)},
		{true, 4, uint16(4)},
		{true, 5, uint32(5)},
		{true, 6, uint64(6)},

		{true, 2, int(2)},
		{true, 3, int8(3)},
		{true, 4, int16(4)},
		{true, 5, int32(5)},
		{true, 6, int64(6)},
		
		{false, 0, int64(-3)},

		{false, 1, float64(1.0001)},
		{false, 1, "1.0"},
	}

	for _, conversion := range conversionUint64Tests {
		if conversion.shouldSucceed {
			output, err := convertToUint64(conversion.input)
			if assert.NoError(t, err) {
				assert.Equal(t, conversion.intention, output, "Failed to convert %v to uint64: %s", conversion.input, err)
			}
		} else {
			_, err := convertToUint64(conversion.input)
			assert.Error(t, err, "Failed at %v", conversion.input)
		}
	}
}