package sessions
import (
	"github.com/go-redis/redis"
	"bridge/utils"
)
type Session struct { }
type RedisStore struct {
	Client *redis.Client
}
type Store interface {
	Get(string) (Session, error)
	Set(string, Session) error
}
func NewMemoryStore() Store {
	client := redis.newClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0
	})
	_, err := client.Ping().Result()
	utils.CheckErr("Error connecting to redis:", err)
	return &RedisStore{client}
}
func(r RedisStore) Get(id string) (Session, error){
    var session Session
    bson, err := r.client.Get(id).Bytes()
    if err != nil {
        return session, err
    }
    if err := json.Unmarshal(bson, &session); err != nil {
        return session, err
    }
    return session, nil
}
func(r RedisStore) Set(s Session, id string) error {
	bson, err := json.Marshal(s)
    if err != nil {
        return err
    }
    if err := r.client.Set(id, bson, 0).Err(); err != nil {
        return err
    }
    return nil
}