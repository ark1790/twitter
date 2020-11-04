package cmd

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/ark1790/alpha/backend"
	"github.com/ark1790/alpha/model"
	"github.com/ark1790/alpha/repo/mongo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	db, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString("MONGO_URI")))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := db.Connect(ctx); err != nil {
		panic(err)
	}

	appDB := db.Database(viper.GetString("DB_NAME"))

	users := mongorepo.NewUser(appDB, "users")
	if err := users.EnsureIndices(&model.User{}); err != nil {
		panic(err)
	}

	fls := mongorepo.NewFollow(appDB, "follows")
	if err := fls.EnsureIndices(&model.Follow{}); err != nil {
		panic(err)
	}

	twts := mongorepo.NewTweet(appDB, "tweets")
	if err := twts.EnsureIndices(&model.Tweet{}); err != nil {
		panic(err)
	}

	fds := mongorepo.NewFeed(appDB, "feeds")
	if err := fds.EnsureIndices(&model.Feed{}); err != nil {
		panic(err)
	}

	backend.NewServer(users, fls, twts, fds).Serve()
}
