package sessions
import (
	"github.com/go-redis/redis"
    "bridge/utils"
    "encoding/json"
)
type Session struct { 
    NoMethods
}
type RedisStore struct {
	Client *redis.Client
}
type NoMethods interface { }
func NewMemoryStore() RedisStore {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 0})
	_, err := client.Ping().Result()
	utils.CheckErr("Error connecting to redis:", err)
	return RedisStore{client}
}
func(r RedisStore) Get(id string) (Session, error){
    var session Session
    bson, err := r.Client.Get(id).Bytes()
    if err != nil {
        return session, err
    }
    if err := json.Unmarshal(bson, &session); err != nil {
        return session, err
    }
    return session, nil
}
func(r RedisStore) Set(id string, s Session) error {
	bson, err := json.Marshal(s)
    if err != nil {
        return err
    }
    if err := r.Client.Set(id, bson, 0).Err(); err != nil {
        return err
    }
    return nil
}