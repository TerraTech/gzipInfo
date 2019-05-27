//==============================================================================
// Copyright (c) 2017, FutureQuest, Inc.
//   https://www.FutureQuest.net
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//==============================================================================

package gzipInfo

import (
	"fmt"
	"os"
)

const whence_end = 2

// IsGzip verifies the GZIP header magic bytes: 0x1F 0x8B
func IsGzip(name string) bool {
	fGzip, err := os.Open(name)
	if err != nil {
		return false
	}
	defer fGzip.Close()

	magic := make([]byte, 2)
	_, err = fGzip.Read(magic)
	if err != nil {
		return false
	}

	return magic[0] == 0x1f && magic[1] == 0x8b
}

// UncompressedSize returns the gzip (RFC 1952) uncompressed file size.
func UncompressedSize(name string) (uint32, error) {
	if !IsGzip(name) {
		return 0, fmt.Errorf("file is not gzip compressed")
	}

	fGzip, err := os.Open(name)
	if err != nil {
		return 0, err
	}

	_, err = fGzip.Seek(-4, whence_end)
	if err != nil {
		return 0, err
	}

	isize := make([]byte, 4)
	_, err = fGzip.Read(isize)
	if err != nil {
		return 0, err
	}

	err = fGzip.Close()
	if err != nil {
		return 0, err
	}

	return uint32(isize[0]) | uint32(isize[1])<<8 | uint32(isize[2])<<16 | uint32(isize[3])<<24, nil
}
