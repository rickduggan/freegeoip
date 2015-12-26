// Copyright 2009-2015 The freegeoip authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"log"

	"github.com/fiorix/freegeoip/apiserver"
)

func main() {
	if err := apiserver.Run(); err != nil {
		log.Fatal(err)
	}
}
