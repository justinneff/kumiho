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
	"path"

	"github.com/justinneff/kumiho/providers"
	"github.com/justinneff/kumiho/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// functionCmd represents the function command
var functionCmd = &cobra.Command{
	Use:   "function <name>",
	Short: "Adds a template function script file",
	Long: `Creates a new function script file in the functions directory.

For example:
kumiho add function my_function

This would create the file ./db/functions/my_function.sql. If the provider,
supports schemas and has a default schema then the created file would be at
./db/{defaultSchema}/functions/my_function.sql.

To assign the function to a database schema other than the provider default,
include the --schema flag.

kumiho add function my_function --schema Sales

To create a function that returns a table instead of a scalar value , include
the --table flag`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a function name")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := providers.GetProvider(viper.GetString("Provider"))
		cobra.CheckErr(err)

		schema, err := p.ResolveSchema(addCmd.PersistentFlags().Lookup("schema").Value.String())
		cobra.CheckErr(err)

		outDir, err := utils.GetOutDir("functions", schema)
		cobra.CheckErr(err)

		name := args[0]
		filename := path.Join(outDir, fmt.Sprintf("%s.sql", name))

		isTable, err := cmd.Flags().GetBool("table")
		cobra.CheckErr(err)

		var content string

		if isTable {
			content, err = p.GenerateTableFunction(schema, name)
		} else {
			content, err = p.GenerateScalarFunction(schema, name)
		}
		cobra.CheckErr(err)

		err = utils.WriteOutFile(filename, content)
		cobra.CheckErr(err)

		fmt.Printf("Created function %s\n", filename)
		return nil
	},
}

func init() {
	addCmd.AddCommand(functionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// functionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	functionCmd.Flags().BoolP("table", "t", false, "Create a table function instead of a scalar function")
}
