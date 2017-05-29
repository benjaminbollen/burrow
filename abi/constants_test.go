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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstantIntegerSize(t *testing.T) {
	assert.Equal(t, IntSize8, IntegerSize(8), "ABI integer size must be 8.")
	assert.Equal(t, IntSize16, IntegerSize(16), "ABI integer size must be 16.")
	assert.Equal(t, IntSize32, IntegerSize(32), "ABI integer size must be 32.")
	assert.Equal(t, IntSize64, IntegerSize(64), "ABI integer size must be 64.")
	assert.Equal(t, IntSize256, IntegerSize(256), "ABI integer size must be 256.")
}
