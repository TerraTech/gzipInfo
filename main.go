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

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/TerraTech/gzipInfo/pkg/gzipInfo"
)

var (
	flagLShowVersion bool
	flagShowVersion  bool
	flagShowUsage    bool
)

func usage() {
	fmt.Printf("usage: %s [flags] file.gz...\n\n", filepath.Base(os.Args[0]))
	fmt.Println("flags: -h")
	fmt.Println("       -V | -version")
	os.Exit(1)
}

func init() {
	flag.BoolVar(&flagLShowVersion, "version", false, "Show version")
	flag.BoolVar(&flagShowVersion, "V", false, "Show version")
	flag.BoolVar(&flagShowUsage, "h", false, "Show help")
}

func main() {
	handleFlags()

	var sum uint32

	for _, fn := range os.Args[1:] {
		if filepath.Ext(fn) != ".gz" {
			fmt.Println("[FATAL] filename given was not a gz file")
			usage()
		}

		us, err := gzipInfo.UncompressedSize(fn)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		sum += us
	}

	fmt.Println(sum)
}

func handleFlags() {
	flag.Parse()

	if flagShowVersion || flagLShowVersion {
		fmt.Println(Version())
		os.Exit(0)
	}

	if flagShowUsage || len(flag.Args()) == 0 {
		usage()
	}
}
