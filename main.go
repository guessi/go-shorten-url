package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Shaked/gomobiledetect"
	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"

	flag "github.com/spf13/pflag"
)

func main() {
	// parse config options
	var config *string = flag.StringP("config", "c", "config/redirections.json", "redirection rules in json format")
	var port *int = flag.IntP("port", "p", 8080, "port number")
	var withfallbackurl *bool = flag.BoolP("with-fallback-url", "k", false, "enable fallback url feature")
	var nocolor *bool = flag.BoolP("no-color", "n", true, "disable color output")
	var debug *bool = flag.BoolP("debug", "d", false, "debug mode")
	var help *bool = flag.BoolP("help", "h", false, "display help messages")

	flag.Parse()

	// print help messages and exit
	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	// load redirction rules from json
	keywords, err := ioutil.ReadFile(*config)
	if err != nil {
		fmt.Println("Failed to open configuration file (", *config, ")")
		os.Exit(1)
	}

	if *nocolor {
		// turn off console color output
		gin.DisableConsoleColor()
	}

	if !*debug {
		// reduce console log output
		gin.SetMode(gin.ReleaseMode)
	}

	// initiate web server
	svc := gin.Default()

	// route handling
	svc.GET("/:uri", func(c *gin.Context) {
		uri := c.Param("uri")

		// set default user agent string
		var userAgent = "default"

		// detect and set user agent
		detect := mobiledetect.NewMobileDetect(c.Request, nil)

		if detect.IsMobile() {
			if detect.Is("ios") {
				userAgent = "iOS"
			} else {
				// known issue: detect.Is("android") not work as expected
				userAgent = "AndroidOS"
			}
		} else {
			userAgent = "default"
		}

		// return redirection
		rc, rule := getRedirection(keywords, uri, userAgent)
		if rc != 404 || *withfallbackurl {
			c.Redirect(302, rule)
		} else {
			c.String(404, "404 page not found")
		}
	})

	// start web server
	svc.Run(fmt.Sprintf(":%d", *port))
}

func getRedirection(config []byte, query string, user_agent string) (int, string) {
	r, err := jsonparser.GetString(config, query, user_agent)
	if err != nil {
		r2, err := jsonparser.GetString(config, query, "default")
		if err != nil {
			r3, _ := jsonparser.GetString(config, "__fallback_url", "default")
			return 404, r3
		}
		return 302, r2
	}
	return 302, r
}
