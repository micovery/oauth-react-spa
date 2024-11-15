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


class ErrorResponse {
  public message: string;
  public data : any

  constructor(message: string, data: any) {
    this.message = message;
    this.data = data;
  }
}

function getName(profile: any) : string {
  if (!profile) {
    return "Unknown User";
  }

  if (profile.name) {
    return profile.name;
  }

  if (profile.given_name && profile.family_name) {
    return `${profile.given_name} ${profile.family_name}`
  }

  if (profile.given_name) {
    return profile.given_name
  }

  if (profile.email) {
    return profile.email
  }

  if (profile.sub) {
    return profile.sub
  }
  return ""
}

export {
  getName,
  ErrorResponse
}