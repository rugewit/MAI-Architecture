package circuit_breaker

import (
	"errors"
	"log"
	"sync"
	"time"
)

type State int

const (
	Closed State = iota
	Open
	HalfOpen
)

type CircuitBreaker struct {
	state            State
	failures         int
	maxFailures      int
	resetTimeout     time.Duration
	lastFailureTime  time.Time
	lock             sync.Mutex
	halfOpenMaxCalls int
	successes        int
}

// NewCircuitBreaker initializes a new CircuitBreaker
func NewCircuitBreaker(maxFailures int, resetTimeout time.Duration, halfOpenMaxCalls int) *CircuitBreaker {
	return &CircuitBreaker{
		state:            Closed,
		failures:         0,
		maxFailures:      maxFailures,
		resetTimeout:     resetTimeout,
		halfOpenMaxCalls: halfOpenMaxCalls,
		successes:        0,
	}
}

// Call tries to execute the function wrapped by the circuit breaker
func (cb *CircuitBreaker) Call(reqFunc func() error) error {
	cb.lock.Lock()

	// Check circuit breaker state
	switch cb.state {
	case Open:
		// Check if timeout has passed for half-open trial period
		if time.Since(cb.lastFailureTime) >= cb.resetTimeout {
			log.Println("CircuitBreaker is now HalfOpen")
			cb.state = HalfOpen
			cb.successes = 0 // reset successes for half-open trial
		} else {
			cb.lock.Unlock()
			return errors.New("circuit breaker is open, request not allowed")
		}
	case HalfOpen:
		// Limit the number of calls in half-open state
		if cb.successes >= cb.halfOpenMaxCalls {
			cb.state = Closed
			cb.failures = 0 // reset failures on successful half-open test
			log.Println("CircuitBreaker is now Closed")
		}
	default:
		// do nothing
	}

	cb.lock.Unlock()

	// Execute the request
	err := reqFunc()

	cb.lock.Lock()
	defer cb.lock.Unlock()

	if err != nil {
		// Log the failure
		log.Printf("Request failed: %v", err)

		cb.failures++
		if cb.failures >= cb.maxFailures {
			cb.state = Open
			cb.lastFailureTime = time.Now()
			log.Println("Circuit breaker transitioned to OPEN state")
		}
		return err
	}

	// Success case
	if cb.state == HalfOpen {
		cb.successes++
	}

	return nil
}
