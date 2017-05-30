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

func Uint64ToWord256(i uint64) Word256 {
	buf := [8]byte{}
	PutUint64BE(buf[:], i)
	return LeftPadWord256(buf[:])
}

func Int64ToWord256(i int64) Word256 {
	buf := [8]byte{}
	PutInt64BE(buf[:], i)
	return LeftPadWord256(buf[:])
}

func RightPadWord256(bz []byte) (word Word256) {
	copy(word[:], bz)
	return
}

func LeftPadWord256(bz []byte) (word Word256) {
	copy(word[32-len(bz):], bz)
	return
}

func Uint64FromWord256(word Word256) uint64 {
	buf := word.Postfix(8)
	return GetUint64BE(buf)
}

func Int64FromWord256(word Word256) int64 {
	buf := word.Postfix(8)
	return GetInt64BE(buf)
}