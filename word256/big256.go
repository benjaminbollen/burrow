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

package word256

import (
	"fmt"
	"math/big"
)

// To256
//
// "cast" the big int to a 256 big int (i.e., limit to)
var tt256 = new(big.Int).Lsh(big.NewInt(1), 256)
var tt256m1 = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))
var tt255 = new(big.Int).Lsh(big.NewInt(1), 255)

// U256 restricts big.Int to Uint256
func U256(x *big.Int) *big.Int {
	x.And(x, tt256m1)
	return x
}

// S256 returns a new big.Int as the (signed) two's complement
// of the unsigned uint256 value;
// NOTE: undefined for big.Int > tt256
func S256(x *big.Int) *big.Int {
	if x.Cmp(tt255) < 0 {
		return x
	} else {
		// We don't want to modify x, ever
		return new(big.Int).Sub(x, tt256)
	}
}

// Uint256ToWord256 copies the big-endian unsigned bytes into Word256
// Truncates big.Int to 256 bits if too big
func Uint256ToWord256(i *big.Int) Word256 {
	// if big.Int has more than 32 bytes, truncate with U256
	return LeftPadWord256(U256(i).Bytes())
}

func CheckOverflowUint256(x *big.Int) error {
	if x.Cmp(tt256) < 0 && big.NewInt(0).Cmp(x) <= 0 {
		return nil
	} else {
		return fmt.Errorf("big.Int out of bounds of Uint256")
	}
}
