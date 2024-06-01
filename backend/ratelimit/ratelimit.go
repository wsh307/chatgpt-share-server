package ratelimit

import (
	"sync"
	"time"
)

type RateLimiter struct {
	requests sync.Map
	duration time.Duration
	maxCount int
}

type requestInfo struct {
	timestamp time.Time
	count     int
}

func NewRateLimiter(duration, maxCount int) *RateLimiter {
	return &RateLimiter{
		duration: time.Duration(duration) * time.Second,
		maxCount: maxCount,
	}
}

func (rl *RateLimiter) Allow(key string) bool {
	v, loaded := rl.requests.LoadOrStore(key, &requestInfo{timestamp: time.Now(), count: 1})
	if !loaded {
		return true
	}

	ri := v.(*requestInfo)
	if time.Since(ri.timestamp) > rl.duration {
		ri.timestamp = time.Now()
		ri.count = 1
		return true
	}

	if ri.count < rl.maxCount {
		ri.count++
		return true
	}

	return false
}
