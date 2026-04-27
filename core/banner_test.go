package core

import (
	"regexp"
	"testing"
)

func TestBannerName(t *testing.T) {
	if Name != "evilcap" {
		t.Fatalf("expected '%s', got '%s'", "evilcap", Name)
	}
}
func TestBannerWebsite(t *testing.T) {
	if Website != "https://github.com/dest4590/evilcap/" {
		t.Fatalf("expected '%s', got '%s'", "https://github.com/dest4590/evilcap/", Website)
	}
}

func TestBannerVersion(t *testing.T) {
	match, err := regexp.MatchString(`\d+.\d+`, Version)
	if err != nil {
		t.Fatalf("unable to perform regex on Version constant")
	}
	if !match {
		t.Fatalf("expected Version constant in format '%s', got '%s'", "X.X", Version)
	}
}

func TestBannerAuthor(t *testing.T) {
	if Author != "Simone 'evilsocket' Margaritelli" {
		t.Fatalf("expected '%s', got '%s'", "Simone 'evilsocket' Margaritelli", Author)
	}
}
