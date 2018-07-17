package controllers
import "bridge/models/sessions"
var redisInstance sessions.RedisStore
func InitializeSessions(store sessions.RedisStore){
	redisInstance = store
}