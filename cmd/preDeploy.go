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
	"time"

	"github.com/justinneff/kumiho/providers"
	"github.com/justinneff/kumiho/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// preDeployCmd represents the preDeploy command
var preDeployCmd = &cobra.Command{
	Use:   "preDeploy <name>",
	Short: "Adds a template pre deploy script file",
	Long: `Adds a new pre deployment script that will be executed before
migrations are executed.

For example:
kumiho add preDeploy do_something

Would add the file ./db/preDeploy/{yyyyMMddHHmmss}_do_something.sql`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a pre deploy script name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		p, err := providers.GetProvider(viper.GetString("Provider"))
		cobra.CheckErr(err)

		name := fmt.Sprintf("%s_%s", time.Now().Format("20060102150405"), args[0])

		outDir, err := utils.GetScriptDir("preDeploy")
		cobra.CheckErr(err)

		filename := path.Join(outDir, fmt.Sprintf("%s.sql", name))

		content, err := p.GeneratePreDeploy(name)
		cobra.CheckErr(err)

		err = utils.WriteOutFile(filename, content)
		cobra.CheckErr(err)

		fmt.Printf("Created pre-deployment: %s\n", filename)
	},
}

func init() {
	addCmd.AddCommand(preDeployCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// preDeployCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// preDeployCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
