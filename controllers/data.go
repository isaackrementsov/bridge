package controllers
import "bridge/models/sessions"
var RedisInstance sessions.RedisStore
func InitializeSessions(store RedisStore){
	RedisInstance = store
}