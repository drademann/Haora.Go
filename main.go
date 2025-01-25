//
// Copyright 2024-2025 The Haora Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import "github.com/drademann/haora/cmd/root"

// imports to initialize commands

import _ "github.com/drademann/haora/cmd/add"
import _ "github.com/drademann/haora/cmd/edit"
import _ "github.com/drademann/haora/cmd/finish"
import _ "github.com/drademann/haora/cmd/list"
import _ "github.com/drademann/haora/cmd/pause"
import _ "github.com/drademann/haora/cmd/remove"
import _ "github.com/drademann/haora/cmd/version"

func main() {
	root.Execute()
}
