package utils

import (
	"sync"
	"time"

	"github.com/alexander-molina/avito_task/internal/app/config"
)

var (
	trackingTimeLimit time.Duration = time.Minute
)

func init() {
	Limiter = &RateLimiter{
		trackers: make(map[string]*addressTracker),
		mux:      &sync.RWMutex{},
	}
}

var (
	// Limiter global rate limiter
	Limiter *RateLimiter
)

// RateLimiter controlls incoming requests from current subnet
type RateLimiter struct {
	trackers map[string]*addressTracker
	mux      *sync.RWMutex
}

// controll current address access
type addressTracker struct {
	requestCount int
	timer        time.Time
	blockedUntil time.Time
}

// GetLimiter return global limiter
func GetLimiter() *RateLimiter {
	return Limiter
}

// AllowRequests check if current address has reached request limit.
// If not return true, else return false
func (r *RateLimiter) AllowRequests(address string) bool {
	appConfig := config.GetConfig()
	r.mux.Lock()
	defer r.mux.Unlock()
	tracker, ok := r.trackers[address]
	if !ok {
		r.trackers[address] = newTracker()
		return true
	}

	// If tracker for current address already exist and
	// tracking time limit expired restart tracker
	if time.Now().After(tracker.timer.Add(trackingTimeLimit)) {
		delete(r.trackers, address)
		r.trackers[address] = newTracker()
		return true
	}

	// if now is less than block time expiration return false
	if time.Now().Before(tracker.blockedUntil) {
		return false
	}

	if tracker.requestCount >= appConfig.ReqestLimit {
		tracker.blockedUntil = time.Now().Add(appConfig.BlockTime)
		r.trackers[address] = tracker
		return false
	}

	tracker.requestCount++
	r.trackers[address] = tracker

	return true
}

// ResetTrackers delete tracker for requested addresses
func (r *RateLimiter) ResetTrackers(addresses []string) {
	r.mux.Lock()
	defer r.mux.Unlock()

	for _, a := range addresses {
		delete(r.trackers, a)
	}
}

// creates a correctly initialized tracker
func newTracker() *addressTracker {
	return &addressTracker{
		requestCount: 1,
		timer:        time.Now(),
		blockedUntil: time.Time{},
	}
}
