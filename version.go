//==============================================================================
// Copyright (c) 2019, FutureQuest, Inc.
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

// +build tools

package main

import (
	_ "futurequest.net/FQgolibs/FQversion"
	_ "futurequest.net/FQgolibs/tools"
)

//go:generate go run $P_GENVERSION -package $GOPACKAGE -prog "$PROG" -version "$VERSION" -build "$BUILD" -import_FQversion $IMPFQVERSION -lib "$LIB"
