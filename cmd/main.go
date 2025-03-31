/*
 *    Copyright 2025 A Jentleman
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package main

import (
	"context"
	"log/slog"
	"os"

	"golang.org/x/sync/errgroup"
)

func main() {
	grp, grpCtx := errgroup.WithContext(context.Background())
	grp.Go(func() error { return rootCmd.ExecuteContext(grpCtx) })
	if err := grp.Wait(); err != nil {
		slog.Error("exiting on error", "error", err)
		os.Exit(1)
	}
}
