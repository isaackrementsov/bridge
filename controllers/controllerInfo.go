package controllers
import (
	"bridge/models/sessions"
	"time"
)
var SessionInstance sessions.SessionStore
var viewDir string = "../src/bridge"
var CookieExpires time.Duration
func SetViewDir(dir string){
	viewDir += dir + "/"
}
func getCookieExpiration() time.Time {
	return time.Now().Add(CookieExpires)
}

