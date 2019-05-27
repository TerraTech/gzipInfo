//==============================================================================
// This file is part of FQgolibs
// Copyright (c) 2017-2019, FutureQuest, Inc.
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

// +build ignore

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	gott "text/template"
	"time"
)

type lType int // License Type
type vType int // Version Type

const (
	LICENSE_FQ lType = iota
	LICENSE_APACHE
)

const (
	VTYPE_APP vType = iota
	VTYPE_LIB
)

const (
	IMPORT_FMT = "fmt"
	IMPORT_NL  = "\n"
)

const (
	default_import_FQversion = "futurequest.net/FQgolibs/FQversion"
)

var license = map[lType]string{
	LICENSE_FQ: "// Copyright (c) {{ .YEAR }}, FutureQuest, Inc.",
	LICENSE_APACHE: `//==============================================================================
// This file is part of {{ .PACKAGE }}
// Copyright (c) {{ .YEAR }}, FutureQuest, Inc.
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
//==============================================================================`,
}

type VersionTemplate struct {
	Timestamp        time.Time
	YEAR             int
	PACKAGE          string
	PROG             string
	VERSION          string
	BUILD            string
	LIB              string
	License          lType
	Imports          []string
	Import_FQversion string
	FQversion        string
}

func usage() {
	fmt.Printf("usage: %s -package -prog -version -build [-import_FQversion] [-lib]\n", filepath.Base(os.Args[0]))
	os.Exit(1)
}

func main() {
	var licenseFQ bool
	var vtype vType = VTYPE_APP
	vt := &VersionTemplate{
		Timestamp: time.Now(),
	}
	vt.YEAR = vt.Timestamp.Year()

	flag.BoolVar(&licenseFQ, "licenseFQ", false, "Emit short FQ Copyright only")
	flag.StringVar(&vt.PACKAGE, "package", "", "package name")
	flag.StringVar(&vt.PROG, "prog", "", "prog/lib name")
	flag.StringVar(&vt.VERSION, "version", "", "version string")
	flag.StringVar(&vt.BUILD, "build", "", "build string")
	flag.StringVar(&vt.Import_FQversion, "import_FQversion", "", "override 'import .../FQversion'")
	flag.StringVar(&vt.LIB, "lib", "", "type name for Version() method receiver")
	flag.Parse()

	mustExist := func(name, val string) {
		if val == "" {
			fmt.Printf("%s must be specified\n", name)
			usage()
		}
	}

	// must exist for all
	mustExist("package", vt.PACKAGE)
	mustExist("prog", vt.PROG)
	mustExist("version", vt.VERSION)
	mustExist("build", vt.BUILD)

	if vt.Import_FQversion == "" {
		vt.Import_FQversion = default_import_FQversion
	}

	if vt.LIB != "" {
		vtype = VTYPE_LIB
	}

	if vt.PACKAGE != "FQversion" {
		vt.FQversion = "FQversion."
	}

	if licenseFQ {
		vt.License = LICENSE_FQ
	} else {
		vt.License = LICENSE_APACHE
	}

	template := vt.getTemplate(vtype)
	if template == nil {
		log.Fatal("template parsing failed: %s", vtype)
	}

	f, err := os.Create("version_autogen.go")
	die(err)
	defer f.Close()

	template.Execute(f, vt)
}

func (v *vType) String() string {
	switch *v {
	case VTYPE_APP:
		return "app"
	case VTYPE_LIB:
		return "lib"
	}
	return ""
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (vt *VersionTemplate) getTemplate(vtype vType) *gott.Template {
	var template string
	switch vtype {
	case VTYPE_APP:
		template = vt.getTemplateApp()
	case VTYPE_LIB:
		template = vt.getTemplateLib()
	default:
		log.Fatal("unrecognized template type: %s", vtype)
	}

	return gott.Must(gott.New("").Parse(template))
}

func (vt *VersionTemplate) getTemplateLib() string {
	if vt.PACKAGE != "FQversion" {
		vt.Imports = append(vt.Imports, vt.Import_FQversion)
	}
	template := getTemplateHeader(license[vt.License])
	template += `
func init() {
	{{ .FQversion }}Register(PROG, VERSION, BUILD)
}
{{ if .FQversion }}
func ({{ .LIB }}) Version() string {
	return FQversion.ProgVersion(PROG, VERSION, BUILD)
}
{{ end -}}
`

	return template
}

func (vt *VersionTemplate) getTemplateApp() string {
	if vt.PROG != "FQgolibs" {
		vt.Imports = append(vt.Imports,
			IMPORT_FMT,
			IMPORT_NL,
		)
	}
	vt.Imports = append(vt.Imports,
		vt.Import_FQversion,
	)
	template := getTemplateHeader(license[vt.License])
	template += `
func Version() string {
	return FQversion.ProgVersion(PROG, VERSION, BUILD)
}
{{ if ne .PROG "FQgolibs" }}
func printVersions() {
	fmt.Println(FQversion.ShowVersionsAligned(PROG, VERSION, BUILD))
}
{{ end -}}
`
	return template
}

func getTemplateHeader(license string) string {
	if license != "" {
		license += "\n\n"
	}

	return license + `// Code generated by go generate; DO NOT EDIT.
// This file was generated by genVersion at
// {{ .Timestamp }}

package {{ .PACKAGE }}
{{ if .Imports }}
import (
	{{- range .Imports }}
{{ if eq . "\n"}}{{else}}	"{{ . }}"{{end}}
	{{- end }}
)
{{ end }}
var (
	PROG    string = "{{ .PROG }}"
	VERSION string = "{{ .VERSION }}"
	BUILD   string = "{{ .BUILD }}"
)
`
}
