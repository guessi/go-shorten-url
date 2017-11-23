package main

import (
	"io/ioutil"
	"testing"
)

func TestGetRedirection(t *testing.T) {
	// load configuration
	keywords, err := ioutil.ReadFile("config/redirections.json")
	if err != nil {
		t.Errorf("Failed to open configuration file")
	}

	// initialized
	expected := ""
	result := ""
	fallback_url := "https://github.com/guessi/go-shorten-url"

	// store
	expected = "https://www.amazon.com"
	_, result = getRedirection(keywords, "store", "default")
	if expected != result {
		t.Errorf("Error, expected: %s, get: %s", expected, result)
	}

	expected = "https://play.google.com"
	_, result = getRedirection(keywords, "store", "AndroidOS")
	if expected != result {
		t.Errorf("Error, expected: %s, get: %s", expected, result)
	}

	expected = "https://www.apple.com/ios/app-store"
	_, result = getRedirection(keywords, "store", "iOS")
	if expected != result {
		t.Errorf("Error, expected: %s, get: %s", expected, result)
	}

	// github
	expected = "https://github.com"
	_, result = getRedirection(keywords, "github", "default")
	if expected != result {
		t.Errorf("Error, expected: %s, get: %s", expected, result)
	}

	expected = "https://github.com"
	_, result = getRedirection(keywords, "github", "not-exist")
	if expected != result {
		t.Errorf("Error, expected: %s, get: %s", expected, result)
	}

	// appleonly
	expected = "https://www.apple.com"
	_, result = getRedirection(keywords, "appleonly", "iOS")
	if expected != result {
		t.Errorf("Error, expected: %s, get: %s", expected, result)
	}

	expected = fallback_url
	_, result = getRedirection(keywords, "appleonly", "not-exist")
	if expected != result {
		t.Errorf("Error, expected: %s, get: %s", expected, result)
	}

	// androidonly
	expected = "https://source.android.com/"
	_, result = getRedirection(keywords, "androidonly", "AndroidOS")
	if expected != result {
		t.Errorf("Error, expected: %s, get: %s", expected, result)
	}

	expected = fallback_url
	_, result = getRedirection(keywords, "androidonly", "not-exist")
	if expected != result {
		t.Errorf("Error, expected: %s, get: %s", expected, result)
	}

	// nodefault
	expected = "https://www.apple.com"
	_, result = getRedirection(keywords, "nodefault", "iOS")
	if expected != result {
		t.Errorf("Error, expected: %s, get: %s", expected, result)
	}

	expected = "https://source.android.com/"
	_, result = getRedirection(keywords, "nodefault", "AndroidOS")
	if expected != result {
		t.Errorf("Error, expected: %s, get: %s", expected, result)
	}

	expected = fallback_url
	_, result = getRedirection(keywords, "nodefault", "not-exist")
	if expected != result {
		t.Errorf("Error, expected: %s, get: %s", expected, result)
	}

	// example
	expected = "https://www.google.com/"
	_, result = getRedirection(keywords, "example", "iOS")
	if expected != result {
		t.Errorf("Error, expected: %s, get: %s", expected, result)
	}

	expected = "https://www.google.com/"
	_, result = getRedirection(keywords, "example", "AndroidOS")
	if expected != result {
		t.Errorf("Error, expected: %s, get: %s", expected, result)
	}

	expected = "https://www.google.com/"
	_, result = getRedirection(keywords, "example", "default")
	if expected != result {
		t.Errorf("Error, expected: %s, get: %s", expected, result)
	}

	expected = "https://www.google.com/"
	_, result = getRedirection(keywords, "example", "not-exist")
	if expected != result {
		t.Errorf("Error, expected: %s, get: %s", expected, result)
	}

	// not-exist
	expected = fallback_url
	_, result = getRedirection(keywords, "not-exist", "not-exist")
	if expected != result {
		t.Errorf("Error, expected: %s, get: %s", expected, result)
	}
}
