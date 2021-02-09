/*
Copyright © 2021 Golem

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/GolemHerry/v2sub/pkg/subscribe"
	"github.com/spf13/cobra"
)

// subscribeCmd represents the subscribe command
var subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("subscribe called")
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			log.Fatalln(err)
		}

		if err := subscribe.Update(url); err != nil {
			log.Fatal(err)
		}
		fmt.Println("SUCCESS")
	},
}

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("info called")

		if err := subscribe.Info(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("SUCCESS")
	},
}

func init() {
	rootCmd.AddCommand(subscribeCmd)

	subscribeCmd.Flags().String("url", "", "subscription address")
	subscribeCmd.MarkFlagRequired("url")

	subscribeCmd.AddCommand(infoCmd)
}
