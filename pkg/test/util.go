/*
 * Copyright 2018 SUSE LINUX GmbH, Nuernberg, Germany..
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
 package test
 import (
	"flag"
	"log"
	"os"
	"testing"
)

// SkipTestIfShort should skip this specific test if we are running with the short flag
// This is intended for skipping integration tests
func SkipTestIfShort(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping tests in short mode")
	}
}

// SkipSetupIfShort should skip the test setup if we are running with the short flag
// This is intended for skipping integration tests
func SkipSetupIfShort(m *testing.M) {
	flag.Parse()
	if testing.Short() {
		log.Print("Skipping tests in short mode")
	}
	os.Exit(0)
}
