package cmd

import (
	"fmt"
	"net"

	"github.com/ark1790/alpha/backend"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start API server",
	Long:  `Start the API server`,
	Run:   serve,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

		portStr := viper.GetString("PORT")
		// portStr := "9000"
		lsnr, err := net.Listen("tcp", ":"+portStr)
		if err != nil {
			return fmt.Errorf("Port %s is not available", portStr)
		}
		_ = lsnr.Close()
		return nil
	},
}

func init() {
	// serveCmd.PersistentFlags().IntP("port", "p", 8080, "port on which the server will listen")
	// serveCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yml", "config file")
	// viper.BindPFlag("port", serveCmd.PersistentFlags().Lookup("port"))
	RootCmd.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, args []string) {
	backend.NewServer().Serve()
}
