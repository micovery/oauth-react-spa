<!--
 Copyright 2024 Google LLC

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
-->
# OAuth (PKCE) with React Single Page App 

This is a sample [React](https://react.dev/) web application (built in [TypeScript](https://www.typescriptlang.org/) with [Vite](https://vite.dev/)), that serves to demo the OAuth [Authorization Code Flow](https://oauth.net/2/grant-types/authorization-code/) with [PKCE](https://oauth.net/2/pkce/)

This app can be used as a starting point for building more complex react apps that need to authenticate users with an OpenID Connect provider.

This app uses the [authTS](https://github.com/authts) libraries.

## Run the App Locally (with Node.js)

This method of running the application is useful for development.
It allows you to try the application on your own machine without deploying it to a cloud environment.
It also allows you to modify the application live, and see the changes immediately.

Open the [config.js](/webapp/public/config.js) file, and edit the `oidc` section

e.g.

```shell
const config = {
  oidc: {
    authority: "...",
    client_id: "...",
    ...
  }
};
```

You need to set both the `authority` and `client_id` fields.

* **authority** - This is the OpenID Connect Authorization Server's issuer URL. If you take this URL, and append `/.well-known/openid-configuration` you should see a JSON file containing config parameters like `authroization_endpoint` and `token_endpoint`.
* **client_id** - This is the Client ID for the application that is registered at the OpenID Connect Authorization server. This application must be configured to allow the [authorization_code](https://oauth.net/2/grant-types/authorization-code/) grant type with [PKCE](https://oauth.net/2/pkce/).


Then, switch to the `webapp` directory, and run:

```shell
npm run dev
```

You can the access the application in your browser at [http://localhost:8080](http://localhost:8080)

## Run the App Locally (with Docker)

This method of running the application is useful for quickly trying it out.

First, build the docker image:
```shell
docker build -t oauth-react-spa .
```

Then, set the following environment variables:

```shell
export AUTHORITY="https://example.com/path/to/as"
export CLIENT_ID="your-application-client-id"
```

Finally, run the application:

```shell
docker run --rm -it \
   -e "AUTHORITY=${AUTHORITY}" \
   -e "CLIENT_ID=${CLIENT_ID}"  \
   -p 8080:8080 oauth-react-spa
```

You can the access the application in your browser at [http://localhost:8080](http://localhost:8080)


## Deploy the App with Cloud Run

Make sure you have installed the Google Cloud CLI (gcloud), and login:

```shell
gcloud auth login
```

Then, set the following environment variables:

```shell
export PROJECT_ID="YOUR_GCP_PROJECT_ID"
export AUTHORITY="https://example.com/path/to/as"
export CLIENT_ID="your-application-client-id"
```

Finally, run the following script:

```shell
./deploy-cloud-run.sh
```

Once the script completes, it will tell you the URL where the application is running on.

## Disclaimer
This application is not an official Google product, nor is it part of an official Google product.