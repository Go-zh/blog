// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build appengine

// This file implements an App Engine blog server.

package main

import (
	"net/http"

	"github.com/Go-zh/tools/blog"
)

func init() {
	config.ContentPath = "content/"
	config.TemplatePath = "template/"
	s, err := blog.NewServer(config)
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; preload")
		s.ServeHTTP(w, r)
	})
}
