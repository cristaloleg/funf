package htmlmetaparse

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	ss := []string{`<!DOCTYPE html>
	<html lang="en">
	
	<script>window.addEventListener('error', window.__err=function f(e){f.p=f.p||[];f.p.push(e)});</script>
	<meta charset="utf-8"><meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<meta name="Description" content="Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.">
	<meta class="js-gtmID" data-gtmid="GTM-W8MVQXG">
	<link href="https://fonts.googleapis.com/css?family=Work+Sans:600|Roboto:400,500,700|Source+Code+Pro" rel="stylesheet">
	<link href="/static/css/stylesheet.css?version=2020-09-22t12-24-165bdd93e41588a06e7cee511f6ffbcfaa7e5bb6" rel="stylesheet">
	
	  <link href="/static/css/sidenav.css?version=2020-09-22t12-24-165bdd93e41588a06e7cee511f6ffbcfaa7e5bb6" rel="stylesheet">
	
	<link href="/third_party/dialog-polyfill/dialog-polyfill.css?version=2020-09-22t12-24-165bdd93e41588a06e7cee511f6ffbcfaa7e5bb6" rel="stylesheet">
	<title>strings package · pkg.go.dev</title>
	`}

	for _, s := range ss {
		res, err := Parse(s)
		if err != nil {
			t.Error(err)
		}

		fmt.Printf("\n%v\n", res)
	}
}

func TestTag(t *testing.T) {
	ss := []string{`<meta charset="utf-8">`}

	for _, s := range ss {
		res, err := parseTag(s)
		if err != nil {
			t.Error(err)
		}

		fmt.Printf("\n%v\n", res)
	}
}
