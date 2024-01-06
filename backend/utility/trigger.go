package utility

import (
	"sync"
	"time"
)

type StatsInstance struct {
	ID    string
	Calls []time.Time
}

var (
	statsMap = make(map[string]*StatsInstance)
	mu       = &sync.Mutex{}
)

func init() {
	ticker := time.NewTicker(1 * time.Hour)
	go func() {
		for range ticker.C {
			clearOldCalls()
		}
	}()
}

func GetStatsInstance(id string) *StatsInstance {
	mu.Lock()
	defer mu.Unlock()

	if instance, ok := statsMap[id]; ok {
		return instance
	}
	instance := &StatsInstance{
		ID: id,
	}
	statsMap[id] = instance
	return instance
}

func (s *StatsInstance) RecordCall() {
	now := time.Now()
	s.Calls = append(s.Calls, now)
}

func (s *StatsInstance) GetCallCount() int {
	return len(s.Calls)
}

func clearOldCalls() {
	mu.Lock()
	defer mu.Unlock()

	now := time.Now()
	threeHoursAgo := now.Add(-3 * time.Hour)

	for _, instance := range statsMap {
		for i, call := range instance.Calls {
			if call.After(threeHoursAgo) {
				instance.Calls = instance.Calls[i:]
				break
			}
		}
	}
}
