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

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

var allFlag, yFlag, nFlag bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List created tasks",
	Long:  `List created tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := bolt.Open("tasks.db", 0666, nil)
		if err != nil {
			fmt.Println("DB ERROR:", err)
			return
		}
		defer db.Close()

		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("tasks"))
			if b == nil {
				fmt.Println("There are no tasks created")
				return nil
			}

			c := b.Cursor()

			if k, _ := c.First(); k == nil {
				fmt.Println("There are no tasks created")
				return nil
			}

			for k, v := c.First(); k != nil; k, v = c.Next() {
				var t task
				json.Unmarshal(v, &t)
				status := "Y"
				if t.Status == 0 {
					status = "N"
				}
				if allFlag || (yFlag && t.Status == 1) || (nFlag && t.Status == 0 && !yFlag) {
					fmt.Printf("TASK %d[%s]: %s\n", t.ID, status, t.Text)
				}
			}
			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&nFlag, "pending", "n", true, "Show pending tasks only")
	listCmd.Flags().BoolVarP(&yFlag, "completed", "y", false, "Show completed tasks only")
	listCmd.Flags().BoolVarP(&allFlag, "all", "a", false, "Show all tasks")
}
