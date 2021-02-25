//
//    Copyright 2020 Satvik Reddy
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
//

package server

import "encoding/binary"

// Uint32ToByteArr takes a 32 bit integer as input, converts it to a byte slice, and returns the byte slice
func Uint32ToByteArr(num uint32) []byte {
	// Create slice
	// Since there are 8 bits per byte, and there are 32 bits, they length of the slice
	// must be 32 / 8
	bytes := make([]byte, 32/8)

	// Put num into slice
	binary.LittleEndian.PutUint32(bytes, num)

	return bytes
}

func Reverse(nums []int64) []int64 {
	for i := 0; i < len(nums); i++ {
		j := len(nums) - i - 1
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}
