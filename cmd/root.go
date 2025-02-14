package cmd

import (
	"github.com/NETWAYS/go-check"
	"github.com/spf13/cobra"
	"os"
)

var (
	timeout = 30
)

var rootCmd = &cobra.Command{
	Use:   "check_elasticsearch",
	Short: "Icinga check plugin to check Elasticsearch",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		go check.HandleTimeout(timeout)
	},
	Run: Help,
}

func Execute(version string) {
	defer check.CatchPanic()

	rootCmd.Version = version
	rootCmd.VersionTemplate()

	if err := rootCmd.Execute(); err != nil {
		check.ExitError(err)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.DisableAutoGenTag = true

	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})

	pfs := rootCmd.PersistentFlags()
	pfs.StringVarP(&cliConfig.Hostname, "hostname", "H", "localhost",
		"Hostname of the Elasticsearch instance")
	pfs.IntVarP(&cliConfig.Port, "port", "p", 9200,
		"Port of the Elasticsearch instance")
	pfs.StringVarP(&cliConfig.Username, "username", "U", "",
		"Username if authentication is required")
	pfs.StringVarP(&cliConfig.Password, "password", "P", "",
		"Password if authentication is required")
	pfs.BoolVarP(&cliConfig.TLS, "tls", "S", false,
		"Use a HTTPS connection")
	pfs.BoolVar(&cliConfig.Insecure, "insecure", false,
		"Skip the verification of the server's TLS certificate")
	pfs.IntVarP(&timeout, "timeout", "t", timeout,
		"Timeout in seconds for the CheckPlugin")

	rootCmd.Flags().SortFlags = false
	pfs.SortFlags = false
}

func Help(cmd *cobra.Command, strings []string) {
	_ = cmd.Usage()

	os.Exit(3)
}
