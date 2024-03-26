/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gotes",
	Short: "A Go note taking CLI app for Markdown notes.",
	Long: `Gotes - A Go note taking CLI app for Markdown notes.
	
	This CLI app was created while studying the Go language. It solves a problem that I had of context switching
for taking meaningful notes. I used to use Obsidian for note taking but now I switched
to plain markdown files, and this app aims to facilitate this workflow. It can use the ChatGPT API to create AI-Powered notes
from small prompts.

	Example of usage:

	- gotes new --name "My first note" --subject "My first subject" --content "My first item;My second item"
		
		The result of this command will be a markdown file with the name "My first note" and the content:
		
		# My first subject
		
		## Summary

		- My first item
		- My second item

	- gotes new --name "My first note" --subject "My first subject" --content "My first item;My second item" --ai
		
		The result of this command will be a markdown file with the name "My first note" and the content:

		Whatever ChatGPT feels like it should create.`,
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	cfgFile     string
	userLicense string
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file for Gotes (default is $HOME/.gotes.yaml)")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "Caio Gallo <caiogallo88@gmail.com>")
	viper.SetDefault("license", "Apache")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".gotes")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error loading config file, the gotes CLI might not work as expected.")
	}
}
