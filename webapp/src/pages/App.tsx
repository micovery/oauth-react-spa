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


import './App.css'
import {useAuth} from "react-oidc-context";
import {getName} from "../util.ts";

function App() {
  const auth = useAuth();

  if (auth.isLoading) {
    return <div>Loading...</div>;
  }

  async function handleLogin() {
    await auth.signinRedirect();
  }
  async function handleLogout() {
    await auth.signoutRedirect({post_logout_redirect_uri:window.location.origin});
    await auth.removeUser();
  }

  async function handleRefresh() {
      await auth.signinSilent()
  }

  return (
    <>
      <h2>OAuth Demo</h2>
      {auth.user ? <h3>Welcome {getName(auth.user?.profile)} !</h3> : ""}
      {!auth.user ? <button style={{margin: "10px"}} onClick={handleLogin}>Log in</button> : ""}
      {auth.user ? <button style={{margin: "10px"}} onClick={handleLogout}>Log out</button> : ""}
      {auth.user?.refresh_token ? <button style={{margin: "10px"}} onClick={handleRefresh}>Refresh</button> : ""}
      {auth.user ? <pre className="pre">{JSON.stringify(auth.user, null, 2)}</pre> : ""}
      {auth?.error?<pre className="pre">{JSON.stringify(auth.error, null, 2)}</pre>:""}

    </>
  );
}

export default App
