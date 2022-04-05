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
// [START drive_activity_v2_quickstart]
package main

import (
	"context"
	"fmt"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/driveactivity/v2"
	"google.golang.org/api/option"
	"log"
	"reflect"
)

// Returns the name of a set property in an object, or else "unknown".
func getOneOf(m interface{}) string {
	v := reflect.ValueOf(m)
	for i := 0; i < v.NumField(); i++ {
		if !v.Field(i).IsNil() {
			return v.Type().Field(i).Name
		}
	}
	return "unknown"
}

// Returns a time associated with an activity.
func getTimeInfo(activity *driveactivity.DriveActivity) string {
	if activity.Timestamp != "" {
		return activity.Timestamp
	}
	if activity.TimeRange != nil {
		return activity.TimeRange.EndTime
	}
	return "unknown"
}

// Returns the type of action.
func getActionInfo(action *driveactivity.ActionDetail) string {
	return getOneOf(*action)
}

// Returns user information, or the type of user if not a known user.
func getUserInfo(user *driveactivity.User) string {
	if user.KnownUser != nil {
		if user.KnownUser.IsCurrentUser {
			return "people/me"
		}
		return user.KnownUser.PersonName
	}
	return getOneOf(*user)
}

// Returns actor information, or the type of actor if not a user.
func getActorInfo(actor *driveactivity.Actor) string {
	if actor.User != nil {
		return getUserInfo(actor.User)
	}
	return getOneOf(*actor)
}

// Returns information for a list of actors.
func getActorsInfo(actors []*driveactivity.Actor) []string {
	actorsInfo := make([]string, len(actors))
	for i := range actors {
		actorsInfo[i] = getActorInfo(actors[i])
	}
	return actorsInfo
}

// Returns the type of target and an associated title.
func getTargetInfo(target *driveactivity.Target) string {
	if target.DriveItem != nil {
		return fmt.Sprintf("driveItem:\"%s\"", target.DriveItem.Title)
	}
	if target.Drive != nil {
		return fmt.Sprintf("drive:\"%s\"", target.Drive.Title)
	}
	if target.FileComment != nil {
		parent := target.FileComment.Parent
		if parent != nil {
			return fmt.Sprintf("fileComment:\"%s\"", parent.Title)
		}
		return "fileComment:unknown"
	}
	return getOneOf(*target)
}

// Returns information for a list of targets.
func getTargetsInfo(targets []*driveactivity.Target) []string {
	targetsInfo := make([]string, len(targets))
	for i := range targets {
		targetsInfo[i] = getTargetInfo(targets[i])
	}
	return targetsInfo
}

func main() {
	ctx := context.Background()
	/* Load pre-authorized user credentials from the environment.
	   TODO(developer) - See https://developers.google.com/identity  and
	     https://cloud.google.com/docs/authentication/production for
	    guides on implementing OAuth2 for your application.
	*/
	client, err := google.DefaultClient(ctx, driveactivity.DriveActivityReadonlyScope)
	if err != nil {
		log.Fatalf("Error creating Google client: %v", err)
	}

	srv, err := driveactivity.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve driveactivity Client %v", err)
	}

	q := driveactivity.QueryDriveActivityRequest{PageSize: 10}
	r, err := srv.Activity.Query(&q).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve list of activities. %v", err)
	}

	fmt.Println("Recent Activity:")
	if len(r.Activities) > 0 {
		for _, a := range r.Activities {
			time := getTimeInfo(a)
			action := getActionInfo(a.PrimaryActionDetail)
			actors := getActorsInfo(a.Actors)
			targets := getTargetsInfo(a.Targets)
			fmt.Printf("%s: %s, %s, %s\n", time, actors, action, targets)
		}
	} else {
		fmt.Print("No activity.")
	}
}

// [END drive_activity_v2_quickstart]
