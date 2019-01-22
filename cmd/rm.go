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
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a task",
	Long:  `Remove a task`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := bolt.Open("tasks.db", 0666, nil)
		if err != nil {
			fmt.Println("DB ERROR:", err)
			return
		}

		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("tasks"))
			if b == nil {
				fmt.Println("There are no created tasks to remove")
				return nil
			}

			n, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			return b.Delete(itob(n))
		})
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
