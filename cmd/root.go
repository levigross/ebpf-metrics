package cmd

import (
	"os"

	"github.com/levigross/bpf-metrics/pkg/metrics"
	"github.com/spf13/cobra"
)

var cfg *metrics.Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ebpf-metrics",
	Short: "An exporter for ebpf and perf metrics",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		return cfg.Run()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVar(&cfg.EnableBPFMetrics, "enable-ebpf-metrics", true, "Try and enable ebpf metrics")
	rootCmd.Flags().BoolVar(&cfg.DisableMetricsOnShutdown, "cleanup-on-shutdown", true, "Disable ebpf metrics when we close gracefully")
	rootCmd.Flags().Uint16Var(&cfg.Port, "metrics-port", 2312, "The port to use for prometheus metrics")
}