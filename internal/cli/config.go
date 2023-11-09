package cli

import (
	"os"

	"github.com/spf13/cobra"
)

var name, bucket string
var dev, rootless bool

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the simulator cli",
	Run: func(cmd *cobra.Command, args []string) {
		if name != "" {
			cfg.Name = name
		}

		if bucket != "" {
			cfg.Bucket = bucket
		}

		if dev {
			cfg.Cli.Dev = true
			cfg.Container.Image = "controlplane/simulator:dev"

			baseDir, err := os.Getwd()
			cobra.CheckErr(err)

			cfg.BaseDir = baseDir
		} else {
			cfg.Cli.Dev = false
			cfg.Container.Image = "controlplane/simulator:latest"
			cfg.BaseDir = ""
		}

		cfg.Container.Rootless = rootless
	},
}

func init() {
	configCmd.PersistentFlags().StringVar(&name, "name", "simulator", "the name for the infrastructure")
	configCmd.PersistentFlags().StringVar(&bucket, "bucket", "", "the s3 bucket used for storage")
	configCmd.PersistentFlags().BoolVar(&dev, "dev", false, "developer mode")
	configCmd.PersistentFlags().BoolVar(&rootless, "rootless", false, "docker running in rootless mode")

	simulatorCmd.AddCommand(configCmd)
}