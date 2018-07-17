package sessions
import (
	"github.com/go-redis/redis"
    "bridge/utils"
    "encoding/json"
)
type Session struct { 
    NoMethods
}
type SessionStore struct {
	Client *redis.Client
}
type NoMethods interface { }
func NewMemoryStore() SessionStore {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 0})
	_, err := client.Ping().Result()
	utils.CheckErr("Error connecting to redis:", err)
	return SessionStore{client}
}
func(r SessionStore) Get(session Session, id string) (Session, error){
    bson, err := r.Client.Get(id).Bytes()
    if err != nil {
        return session, err
    }
    if err := json.Unmarshal(bson, &session); err != nil {
        return session, err
    }
    return session, nil
}
func(r SessionStore) Set(id string, s Session) error {
	bson, err := json.Marshal(s)
    if err != nil {
        return err
    }
    if err := r.Client.Set(id, bson, 0).Err(); err != nil {
        return err
    }
    return nil
}