package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/charmbracelet/log"
	"github.com/mitchellh/go-homedir"

	// "github.com/rahuldesar/yt-cli/ui"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/bro/.cobra.yaml)")
	rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home + "/.config/yt-custom/")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err := viper.WriteConfig()
			if err != nil {
				// TODO: setup config file here
				fmt.Println("TODO: setup config file here")
				log.Error("", err)
			}
		} else {
			log.Error("", err)
		}
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "yt-cli",
	Short: "ytcli short desc",
	Long:  `ytcli long desc`,
	Run: func(cmd *cobra.Command, args []string) {
		// viperhelper.ShowAllViperSettings()
		// fmt.Println("this should run bubbletea?")

		// lipgloss.SetHasDarkBackground(termenv.HasDarkBackground())
		// INFO: oau
		// p := tea.NewProgram(
		// 	tabs.InitialModel(),
		// 	tea.WithAltScreen(),
		// 	// tea.WithReportFocus(),
		// )
		// if _, err := p.Run(); err != nil {
		// 	log.Fatal("Failed starting the TUI", err)
		// }

		testRun()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// ME: TEST HERE
func testRun() {
	type Thumbnail struct {
		Url    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	}

	type VideoEntry struct {
		Title               string      `json:"title"`
		Availability        string      `json:"availability"`
		Channel             string      `json:"channel"`
		Channel_id          string      `json:"channel_id"`
		Channel_is_verified bool        `json:"channel_is_verified"`
		Channel_url         string      `json:"channel_url"`
		Description         string      `json:"description"`
		Duration            float32     `json:"duration"`
		Id                  string      `json:"id"`
		Ie_key              string      `json:"ie_key"`
		Live_status         string      `json:"live_status"`
		Release_timestamp   string      `json:"release_timestamp"`
		Thumbnails          []Thumbnail `json:"thumbnails"`
		Timestamp           int         `json:"timestamp"`
		Uploader            string      `json:"uploader"`
		Uploader_id         string      `json:"uploader_id"`
		Uploader_url        string      `json:"uploader_url"`
		Url                 string      `json:"url"`
		View_count          int         `json:"view_count"`
	}

	type response struct {
		Entries []VideoEntry `json:"entries"`
	}

	// ytOut, err := utils.RunCommand("yt-dlp",
	// 	"https://www.youtube.com/feed/recommended",
	// 	"-J",
	// 	"--flat-playlist",
	// 	"--extractor-args", "youtubetab:approximate_date",
	// 	"--playlist-start", "1",
	// 	"--playlist-end", "20",
	// 	"--cookies-from-browser", "chrome:Profile_1",
	// )
	// if err != nil {
	// 	fmt.Println(err)
	// }

	myJson, err := os.Open("./dummy-jsons/feed.json")
	if err != nil {
		fmt.Println(err)
	}
	defer myJson.Close()
	ytOut, err := io.ReadAll(myJson)
	if err != nil {
		fmt.Printf("Ioutil ReadAll error: %v", err)
	}

	var myValue response
	err = json.Unmarshal(ytOut, &myValue)
	if err != nil {
		fmt.Println(err)
	}
	// for _, v := range myValue.Entries {
	//
	// 	fmt.Println("Title:", v.Title)
	// 	fmt.Println("Url:", v.Url)
	// 	fmt.Println("Channel:", v.Channel)
	// 	fmt.Println("Uploader:", v.Uploader)
	// 	fmt.Println("Description:", v.Description)
	// 	fmt.Println("Thumbnails:", v.Thumbnails[len(v.Thumbnails)-1].Url)
	//
	// 	fmt.Println("========")
	//
	// }
	fmt.Printf("%+v", myValue.Entries[0])
}
