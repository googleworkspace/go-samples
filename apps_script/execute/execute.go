/**
 * @license
 * Copyright Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package main

func main() {
  // [START apps_script_api_execute]

  scriptId := "ENTER_YOUR_SCRIPT_ID_HERE"
  client := getClient(ctx, config)

  // Generate a service object.
  srv, err := script.New(client)
  if err != nil {
    log.Fatalf("Unable to retrieve script Client %v", err)
  }

  // Create an execution request object.
  req := script.ExecutionRequest{Function:"getFoldersUnderRoot"}

  // Make the API request.
  resp, err := srv.Scripts.Run(scriptId, &req).Do()
  if err != nil {
    // The API encountered a problem before the script started executing.
    log.Fatalf("Unable to execute Apps Script function. %v", err)
  }

  if resp.Error != nil {
    // The API executed, but the script returned an error.

    // Extract the first (and only) set of error details and cast as a map.
    // The values of this map are the script's 'errorMessage' and
    // 'errorType', and an array of stack trace elements (which also need to
    // be cast as maps).
    error := resp.Error.Details[0].(map[string]interface{})
    fmt.Printf("Script error message: %s\n", error["errorMessage"]);

    if (error["scriptStackTraceElements"] != nil) {
      // There may not be a stacktrace if the script didn't start executing.
      fmt.Printf("Script error stacktrace:\n")
      for _, trace := range error["scriptStackTraceElements"].([]interface{}) {
        t := trace.(map[string]interface{})
        fmt.Printf("\t%s: %d\n", t["function"], int(t["lineNumber"].(float64)))
      }
    }
  } else {
    // The result provided by the API needs to be cast into the correct type,
    // based upon what types the Apps Script function returns. Here, the
    // function returns an Apps Script Object with String keys and values, so
    // must be cast into a map (folderSet).
    r := resp.Response.(map[string]interface{})
    folderSet := r["result"].(map[string]interface{})
    if len(folderSet) == 0 {
      fmt.Printf("No folders returned!\n")
    } else {
      fmt.Printf("Folders under your root folder:\n")
      for id, folder := range folderSet {
        fmt.Printf("\t%s (%s)\n", folder, id)
      }
    }
  }
  // [END apps_script_api_execute]
}