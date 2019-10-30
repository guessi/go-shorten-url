package main

import (
	"io/ioutil"
	"testing"
)

var (
	userAgentUnsupported = "unknown"
	fallbackURL         = "https://github.com/guessi/go-shorten-url"
)

func runTest(t *testing.T, keywords []byte, keyword, expected, userAgent string) {
	_, result := getRedirection(keywords, keyword, userAgent)
	if expected != result {
		t.Errorf("Error, expected: %s, get: %s", expected, result)
	}
}

func TestSuite1(t *testing.T) {
	// load configuration
	keywords, err := ioutil.ReadFile("config/redirections.json")
	if err != nil {
		t.Errorf("Failed to open configuration file")
	}

	runTest(t, keywords, "__fallbackURL", fallbackURL, userAgentUnsupported)
	runTest(t, keywords, "__fallbackURL", fallbackURL, userAgentDefault)
	runTest(t, keywords, "__fallbackURL", fallbackURL, userAgentAndroidOS)
	runTest(t, keywords, "__fallbackURL", fallbackURL, userAgentIOS)
}

func TestSuite2(t *testing.T) {
	// load configuration
	keywords, err := ioutil.ReadFile("config/redirections.json")
	if err != nil {
		t.Errorf("Failed to open configuration file")
	}

	runTest(t, keywords, "store", "https://www.amazon.com", userAgentUnsupported)
	runTest(t, keywords, "store", "https://www.amazon.com", userAgentDefault)
	runTest(t, keywords, "store", "https://play.google.com", userAgentAndroidOS)
	runTest(t, keywords, "store", "https://www.apple.com/ios/app-store", userAgentIOS)
}

func TestSuite3(t *testing.T) {
	// load configuration
	keywords, err := ioutil.ReadFile("config/redirections.json")
	if err != nil {
		t.Errorf("Failed to open configuration file")
	}

	runTest(t, keywords, "github", "https://github.com", userAgentUnsupported)
	runTest(t, keywords, "github", "https://github.com", userAgentDefault)
	runTest(t, keywords, "github", "https://github.com", userAgentAndroidOS)
	runTest(t, keywords, "github", "https://github.com", userAgentIOS)
}

func TestSuite4(t *testing.T) {
	// load configuration
	keywords, err := ioutil.ReadFile("config/redirections.json")
	if err != nil {
		t.Errorf("Failed to open configuration file")
	}

	runTest(t, keywords, "appleonly", fallbackURL, userAgentUnsupported)
	runTest(t, keywords, "appleonly", fallbackURL, userAgentDefault)
	runTest(t, keywords, "appleonly", fallbackURL, userAgentAndroidOS)
	runTest(t, keywords, "appleonly", "https://www.apple.com", userAgentIOS)
}

func TestSuite5(t *testing.T) {
	// load configuration
	keywords, err := ioutil.ReadFile("config/redirections.json")
	if err != nil {
		t.Errorf("Failed to open configuration file")
	}

	runTest(t, keywords, "androidonly", fallbackURL, userAgentUnsupported)
	runTest(t, keywords, "androidonly", fallbackURL, userAgentDefault)
	runTest(t, keywords, "androidonly", "https://source.android.com/", userAgentAndroidOS)
	runTest(t, keywords, "androidonly", fallbackURL, userAgentIOS)
}

func TestSuite6(t *testing.T) {
	// load configuration
	keywords, err := ioutil.ReadFile("config/redirections.json")
	if err != nil {
		t.Errorf("Failed to open configuration file")
	}

	runTest(t, keywords, "nodefault", fallbackURL, userAgentUnsupported)
	runTest(t, keywords, "nodefault", fallbackURL, userAgentDefault)
	runTest(t, keywords, "nodefault", "https://source.android.com/", userAgentAndroidOS)
	runTest(t, keywords, "nodefault", "https://www.apple.com", userAgentIOS)
}

func TestSuite7(t *testing.T) {
	// load configuration
	keywords, err := ioutil.ReadFile("config/redirections.json")
	if err != nil {
		t.Errorf("Failed to open configuration file")
	}

	runTest(t, keywords, "example", "https://www.google.com/", userAgentUnsupported)
	runTest(t, keywords, "example", "https://www.google.com/", userAgentDefault)
	runTest(t, keywords, "example", "https://www.google.com/", userAgentAndroidOS)
	runTest(t, keywords, "example", "https://www.google.com/", userAgentIOS)
}

func TestSuite8(t *testing.T) {
	// load configuration
	keywords, err := ioutil.ReadFile("config/redirections.json")
	if err != nil {
		t.Errorf("Failed to open configuration file")
	}

	runTest(t, keywords, "not-defined", fallbackURL, userAgentUnsupported)
	runTest(t, keywords, "not-defined", fallbackURL, userAgentDefault)
	runTest(t, keywords, "not-defined", fallbackURL, userAgentAndroidOS)
	runTest(t, keywords, "not-defined", fallbackURL, userAgentIOS)
}
