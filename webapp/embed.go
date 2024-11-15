// Copyright 2024 Google LLC
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

//go:generate npm install --no-progress
//go:generate npm run build

package webapp

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed dist/*
var Fs embed.FS

func HTTPFileSystem() (_ http.FileSystem, err error) {
	var distFS fs.FS
	if distFS, err = fs.Sub(Fs, "dist"); err != nil {
		return nil, err
	}

	return http.FS(distFS), nil
}
