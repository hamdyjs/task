// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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

package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := bolt.Open("tasks.db", 0666, nil)
		if err != nil {
			fmt.Println("DB ERROR:", err)
			return
		}
		defer db.Close()

		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("tasks"))
			if b == nil {
				fmt.Println("There are no tasks to complete")
				return nil
			}

			n, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			var t task
			jsonTask := b.Get(itob(n))
			if jsonTask == nil {
				fmt.Println("There are no task with the given id")
				return nil
			}

			json.Unmarshal(jsonTask, &t)
			t.Status = 1

			jsonUpdatedTask, err := json.Marshal(t)
			if err != nil {
				fmt.Println("Error completing task")
				return err
			}

			return b.Put(itob(n), jsonUpdatedTask)
		})
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
