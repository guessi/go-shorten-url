package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	mobiledetect "github.com/Shaked/gomobiledetect"
	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"

	flag "github.com/spf13/pflag"
)

const (
	userAgentDefault   = "default"
	userAgentIOS       = "iOS"
	userAgentAndroidOS = "AndroidOS"
)

var (
	withfallbackurl *bool   = flag.BoolP("with-fallback-url", "k", false, "enable fallback url feature")
	nocolor         *bool   = flag.BoolP("no-color", "n", true, "disable color output")
	debug           *bool   = flag.BoolP("debug", "d", false, "debug mode")
	help            *bool   = flag.BoolP("help", "h", false, "display help messages")
	port            *int    = flag.IntP("port", "p", 8080, "port number")
	config          *string = flag.StringP("config", "c", "config/redirections.json", "redirection rules in json format")
)

func main() {
	// parse config options
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

	// turn off console color output
	if *nocolor {
		gin.DisableConsoleColor()
	}

	// reduce console log output
	if !*debug {
		gin.SetMode(gin.ReleaseMode)
	}

	// initiate web server
	svc := gin.Default()

	// route handling
	svc.GET("/:uri", func(c *gin.Context) {
		uri := c.Param("uri")

		// set default user agent string
		userAgent := userAgentDefault

		// detect and set user agent
		detect := mobiledetect.NewMobileDetect(c.Request, nil)

		if detect.IsMobile() {
			if detect.Is("ios") {
				userAgent = userAgentIOS
			} else {
				// known issue: detect.Is("android") not work as expected
				userAgent = userAgentAndroidOS
			}
		} else {
			userAgent = userAgentDefault
		}

		// return redirection
		rc, rule := getRedirection(keywords, uri, userAgent)
		if rc != http.StatusNotFound || *withfallbackurl {
			c.Redirect(http.StatusFound, rule)
		} else {
			c.String(http.StatusNotFound, "404 page not found")
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
			return http.StatusNotFound, r3
		}
		return http.StatusFound, r2
	}
	return http.StatusFound, r
}
