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

// postDeployCmd represents the postDeploy command
var postDeployCmd = &cobra.Command{
	Use:   "postDeploy <name>",
	Short: "Adds a template post deploy script file",
	Long: `Adds a new post deployment script that will be executed after
migrations and database objects have been executed.

For example:
kumiho add postDeploy populate_some_data

Would add the file ./db/postDeploy/{yyyyMMddHHmmss}_populate_some_data.sql`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a post deploy script name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		p, err := providers.GetProvider(viper.GetString("Provider"))
		cobra.CheckErr(err)

		name := fmt.Sprintf("%s_%s", time.Now().Format("20060102150405"), args[0])

		outDir, err := utils.GetOutDir("postDeploy", "")
		cobra.CheckErr(err)

		filename := path.Join(outDir, fmt.Sprintf("%s.sql", name))

		content, err := p.GeneratePostDeploy(name)
		cobra.CheckErr(err)

		err = utils.WriteOutFile(filename, content)
		cobra.CheckErr(err)

		fmt.Printf("Created post-deployment: %s\n", filename)
	},
}

func init() {
	addCmd.AddCommand(postDeployCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// postDeployCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// postDeployCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
