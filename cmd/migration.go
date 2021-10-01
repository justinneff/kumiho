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
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"time"

	"github.com/justinneff/kumiho/providers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// migrationCmd represents the migration command
var migrationCmd = &cobra.Command{
	Use:   "migration <name>",
	Short: "Adds a template migration script file",
	Long: `Create a new migration script file in the migrations directory.

For example:
kumiho add migration add_column_to_table

Would add the file ./db/migrations/{yyyyMMddHHmmss}_add_column_to_table.sql`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a migration name")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		viper.GetString("Provider")
		p, err := providers.GetProvider(viper.GetString("Provider"))
		cobra.CheckErr(err)

		name := fmt.Sprintf("%s_%s", time.Now().Format("20060102150405"), args[0])
		cwd, err := os.Getwd()
		cobra.CheckErr(err)

		outDir := path.Join(cwd, viper.GetString("Dir"), "migrations")
		filename := path.Join(outDir, fmt.Sprintf("%s.sql", name))

		if fs.ValidPath(filename) {
			return fmt.Errorf("migration already exists at %s", filename)
		}

		content, err := p.GenerateMigration(name)
		cobra.CheckErr(err)

		err = os.MkdirAll(outDir, 0755)
		cobra.CheckErr(err)

		err = os.WriteFile(filename, []byte(content), 0777)
		cobra.CheckErr(err)

		fmt.Printf("Created migration: %s\n", filename)
		return nil
	},
}

func init() {
	addCmd.AddCommand(migrationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
