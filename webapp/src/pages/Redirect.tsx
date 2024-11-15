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


import {useEffect, useState} from "react";
import {ErrorResponse} from "../util.ts";

import { useNavigate } from "react-router-dom";

function Redirect() {
  const [message, setMessage] = useState<string>("")
  const [errors, setErrors] = useState<any[]>([])
  const navigate = useNavigate();

  function sleep(ms:number) {
    return new Promise(resolve => setTimeout(resolve, ms));
  }

  useEffect( ()  => {
    (async () => {
      try {
        setMessage("Getting OAuth access token ...")
        await sleep(1000);

        setMessage("Getting user info ...")

        setMessage("Exchange complete, redirecting ...")
        await sleep(3000);
        navigate("/");

      } catch(ex : any) {
        setErrors([...errors, (ex instanceof ErrorResponse)? ex:ex.message])
      }
    })();

    return () => {
      // Component unmount code.
    };
  }, []);

  return (
    <>
      <div className="card">{message}</div>
      <pre className="pre">{errors.length > 0 ? JSON.stringify({errors}, null, 2) : ""}</pre>
    </>
  )
}

export default Redirect
