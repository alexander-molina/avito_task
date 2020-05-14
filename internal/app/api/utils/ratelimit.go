package utils

import (
	"sync"
	"time"
)

func init() {
	Limiter = &RateLimiter{
		trackers: make(map[string]*addressTracker),
		mux:      &sync.RWMutex{},
	}
}

const (
	maxRequests       int           = 100
	trackingTimeLimit time.Duration = time.Minute
	blockTime         time.Duration = time.Minute * 2
)

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

	if tracker.requestCount >= maxRequests {
		tracker.blockedUntil = time.Now().Add(blockTime)
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
