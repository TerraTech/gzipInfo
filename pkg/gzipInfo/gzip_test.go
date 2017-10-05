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

package gzipInfo_test

import (
	"path/filepath"
	"os"
	"testing"

	"github.com/TerraTech/gzipInfo/pkg/gzipInfo"

	"github.com/stretchr/testify/assert"
)

var d_base = filepath.Join(os.Getenv("GOPATH"), "src/github.com/TerraTech/gzipInfo")

func TestUncompressedSize(t *testing.T) {
	var expectations = []struct{
		name string
		size uint32
	}{
		{"test-26.gz", 26},
		{"test-1500.gz", 1500000000},
	}

	for _, expected := range expectations {
		isize, err := gzipInfo.UncompressedSize(filepath.Join(d_base, "files", expected.name))
		if !assert.NoError(t, err) {
			return
		}
		if !assert.Equal(t, expected.size, isize) {
			return
		}
	}
}
