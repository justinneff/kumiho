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
	"fmt"
	"io/fs"
	"os"
	"path"

	"github.com/justinneff/kumiho/providers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// procedureCmd represents the procedure command
var procedureCmd = &cobra.Command{
	Use:   "procedure <name>",
	Short: "Adds a template stored procedure script file",
	Long: `Creates a new stored procedure script file in the procedures directory.

For example:
kumiho add procedure my_procedure

This would create the file ./db/procedures/my_procedure.sql. If the provider,
supports schemas and has a default schema then the created file would be at
./db/{defaultSchema}/procedures/my_procedure.sql.

To assign the procedure to a database schema other than the provider default,
include the --schema flag.

kumiho add procedure my_procedure --schema Sales

This would create the file ./db/Sales/procedures/my_procedure.sql.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := providers.GetProvider(viper.GetString("Provider"))
		cobra.CheckErr(err)

		schema, err := p.ResolveSchema(addCmd.PersistentFlags().Lookup("schema").Value.String())
		cobra.CheckErr(err)

		cwd, err := os.Getwd()
		cobra.CheckErr(err)

		var outDir string

		if len(schema) > 0 {
			outDir = path.Join(cwd, viper.GetString("Dir"), schema, "procedures")
		} else {
			outDir = path.Join(cwd, viper.GetString("Dir"), "procedures")
		}

		name := args[0]
		filename := path.Join(outDir, fmt.Sprintf("%s.sql", name))

		if fs.ValidPath(filename) {
			return fmt.Errorf("procedure already exists at %s", filename)
		}

		content, err := p.GenerateProcedure(schema, name)
		cobra.CheckErr(err)

		err = os.MkdirAll(outDir, 0755)
		cobra.CheckErr(err)

		err = os.WriteFile(filename, []byte(content), 0777)
		cobra.CheckErr(err)

		fmt.Printf("Create procedure: %s\n", filename)
		return nil
	},
}

func init() {
	addCmd.AddCommand(procedureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// procedureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// procedureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
