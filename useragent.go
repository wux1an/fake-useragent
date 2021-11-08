package ua

import (
	"crypto/rand"
	"embed"
	_ "embed"
	"encoding/json"
	"math/big"
)

var (
	//go:embed static/*
	fs embed.FS
)

// uaType user-agent type
type uaType string

const (
	Android uaType = "android"
	Chrome  uaType = "chrome"
	Desktop uaType = "desktop"
	Firefox uaType = "firefox"
	IE      uaType = "ie"
	IOS     uaType = "ios"
	IPad    uaType = "ipad"
	IPhone  uaType = "iphone"
	Linux   uaType = "linux"
	MacOSX  uaType = "mac_os_x"
	Mobile  uaType = "mobile"
	Safari  uaType = "safari"
)

var (
	uaCache = map[uaType]*[]string{
		Android: {},
		Chrome:  {},
		Desktop: {},
		Firefox: {},
		IE:      {},
		IOS:     {},
		IPad:    {},
		IPhone:  {},
		Linux:   {},
		MacOSX:  {},
		Mobile:  {},
		Safari:  {},
	}
)

// load data if empty
func load(name uaType) {
	var ua = uaCache[name]
	if loaded := len(*ua) > 0; loaded {
		return
	}

	var data, _ = fs.ReadFile("static/" + string(name) + ".json")
	_ = json.Unmarshal(data, ua)
}

// Random return a random user-agent
func Random() string {
	var count = 0
	for key, value := range uaCache {
		load(key)
		count += len(*value)
	}

	var n, _ = rand.Int(rand.Reader, big.NewInt(int64(count)))
	var rnd = int(n.Int64())

	for _, value := range uaCache {
		var l = len(*value)
		if skip := l <= rnd; skip {
			rnd -= l
			continue
		}

		return (*value)[rnd]
	}

	return "" // Unreachable code
}

// RandomType return a random user-agent of the specified type
func RandomType(t uaType) string {
	var ua = uaCache[t]

	if valid := ua != nil; !valid {
		return Random()
	}

	load(t)
	var n, _ = rand.Int(rand.Reader, big.NewInt(int64(len(*ua))))
	return (*ua)[n.Int64()]
}
