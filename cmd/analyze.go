/*
Copyright © 2021 Justin Neff <neffjustin@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/justinneff/kumiho/providers"
	"github.com/justinneff/kumiho/publishing"
	"github.com/justinneff/kumiho/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// analyzeCmd represents the analyze command
var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Generates a dependency tree for database objects",
	Long: `Inspects all database objects and generates a dependency tree for the
objects. The results are printed showing the order objects will be published to
the database. Results are cached to the .kumiho folder.`,
	Run: func(cmd *cobra.Command, args []string) {
		p, err := providers.GetProvider(viper.GetString("Provider"))
		cobra.CheckErr(err)

		dbDir, err := utils.GetDatabaseDir()
		cobra.CheckErr(err)

		objectPaths, err := publishing.GetDatabaseObjectPaths(dbDir)
		cobra.CheckErr(err)

		for _, obj := range objectPaths {
			item, err := publishing.CreateDatabaseObject(obj, p)
			cobra.CheckErr(err)
			itemJson, err := json.Marshal(item)
			cobra.CheckErr(err)
			fmt.Println(string(itemJson))
		}
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// analyzeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// analyzeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
