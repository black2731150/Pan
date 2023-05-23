package global

import (
	"sync"
	"time"
)

var EmailCodeMap map[string]string = make(map[string]string, 100)
var mu sync.Mutex

func AddToMap(key, value string, duration time.Duration) {
	mu.Lock()
	EmailCodeMap[key] = value
	mu.Unlock()

	time.AfterFunc(duration, func() {
		mu.Lock()
		delete(EmailCodeMap, key)
		mu.Unlock()
	})
}

func GetEmailCodeFromMap(key string) (string, bool) {
	mu.Lock()
	value, ok := EmailCodeMap[key]
	mu.Unlock()
	return value, ok
}
