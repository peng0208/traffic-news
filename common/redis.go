package common

//
//var kv *redis.Pool
//
//func InitRedis() {
//	host := Cfg.Section("redis").Key("host").String()
//	port := Cfg.Section("redis").Key("port").String()
//	pass := Cfg.Section("redis").Key("password").String()
//	kv = newRedis(host, port , pass)
//}
//
//func newRedis(host, port, pass string) *redis.Pool {
//	dsn := fmt.Sprintf("%s:%s", host, port)
//	pool := &redis.Pool{
//		MaxIdle:     10,
//		MaxActive:   100,
//		IdleTimeout: 180 * time.Second,
//		Dial: func() (redis.Conn, error) {
//			c, err := redis.Dial("tcp", dsn)
//			if err != nil {
//				return nil, err
//			}
//			c.Do("auth", pass)
//			//c.Do("SELECT", 0)
//			return c, nil
//		},
//	}
//
//	return pool
//}
//
//func Set(k, v, t string) bool {
//	if _, err := kv.Get().Do("set", k, v, "EX 1800"); err != nil {
//		return false
//	}
//	return true
//}
//
//func SetWithTTL(k, v, t string) bool {
//	if _, err := kv.Get().Do("set", k, v, "EX", t); err != nil {
//		return false
//	}
//	return true
//}
//
//func Get(k string) string {
//	v, err := redis.String(kv.Get().Do("get", k))
//	if err != nil {
//		return ""
//	}
//	return v
//}