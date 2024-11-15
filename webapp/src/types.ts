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


interface OAuthContext {
  access_token: string,
  refresh_token: string,
  expires_in: number,
  expires_at: number | undefined,
  refresh_expires_in: number,
  refresh_expires_at: number | undefined,
  id_token: string
}

interface UserInfo {
  sub: string
  name: string,
  given_name: string,
  family_name: string,
  preferred_username: string,
  email: string,
  picture: string
}


export type { OAuthContext, UserInfo };
