// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc/db"
	"github.com/concourse/atc/radar"
)

type FakeLocker struct {
	AcquireResourceCheckingLockStub        func() (db.Lock, error)
	acquireResourceCheckingLockMutex       sync.RWMutex
	acquireResourceCheckingLockArgsForCall []struct{}
	acquireResourceCheckingLockReturns struct {
		result1 db.Lock
		result2 error
	}
	AcquireReadLockStub        func(names []string) (db.Lock, error)
	acquireReadLockMutex       sync.RWMutex
	acquireReadLockArgsForCall []struct {
		names []string
	}
	acquireReadLockReturns struct {
		result1 db.Lock
		result2 error
	}
	AcquireWriteLockStub        func(names []string) (db.Lock, error)
	acquireWriteLockMutex       sync.RWMutex
	acquireWriteLockArgsForCall []struct {
		names []string
	}
	acquireWriteLockReturns struct {
		result1 db.Lock
		result2 error
	}
}

func (fake *FakeLocker) AcquireResourceCheckingLock() (db.Lock, error) {
	fake.acquireResourceCheckingLockMutex.Lock()
	fake.acquireResourceCheckingLockArgsForCall = append(fake.acquireResourceCheckingLockArgsForCall, struct{}{})
	fake.acquireResourceCheckingLockMutex.Unlock()
	if fake.AcquireResourceCheckingLockStub != nil {
		return fake.AcquireResourceCheckingLockStub()
	} else {
		return fake.acquireResourceCheckingLockReturns.result1, fake.acquireResourceCheckingLockReturns.result2
	}
}

func (fake *FakeLocker) AcquireResourceCheckingLockCallCount() int {
	fake.acquireResourceCheckingLockMutex.RLock()
	defer fake.acquireResourceCheckingLockMutex.RUnlock()
	return len(fake.acquireResourceCheckingLockArgsForCall)
}

func (fake *FakeLocker) AcquireResourceCheckingLockReturns(result1 db.Lock, result2 error) {
	fake.AcquireResourceCheckingLockStub = nil
	fake.acquireResourceCheckingLockReturns = struct {
		result1 db.Lock
		result2 error
	}{result1, result2}
}

func (fake *FakeLocker) AcquireReadLock(names []string) (db.Lock, error) {
	fake.acquireReadLockMutex.Lock()
	fake.acquireReadLockArgsForCall = append(fake.acquireReadLockArgsForCall, struct {
		names []string
	}{names})
	fake.acquireReadLockMutex.Unlock()
	if fake.AcquireReadLockStub != nil {
		return fake.AcquireReadLockStub(names)
	} else {
		return fake.acquireReadLockReturns.result1, fake.acquireReadLockReturns.result2
	}
}

func (fake *FakeLocker) AcquireReadLockCallCount() int {
	fake.acquireReadLockMutex.RLock()
	defer fake.acquireReadLockMutex.RUnlock()
	return len(fake.acquireReadLockArgsForCall)
}

func (fake *FakeLocker) AcquireReadLockArgsForCall(i int) []string {
	fake.acquireReadLockMutex.RLock()
	defer fake.acquireReadLockMutex.RUnlock()
	return fake.acquireReadLockArgsForCall[i].names
}

func (fake *FakeLocker) AcquireReadLockReturns(result1 db.Lock, result2 error) {
	fake.AcquireReadLockStub = nil
	fake.acquireReadLockReturns = struct {
		result1 db.Lock
		result2 error
	}{result1, result2}
}

func (fake *FakeLocker) AcquireWriteLock(names []string) (db.Lock, error) {
	fake.acquireWriteLockMutex.Lock()
	fake.acquireWriteLockArgsForCall = append(fake.acquireWriteLockArgsForCall, struct {
		names []string
	}{names})
	fake.acquireWriteLockMutex.Unlock()
	if fake.AcquireWriteLockStub != nil {
		return fake.AcquireWriteLockStub(names)
	} else {
		return fake.acquireWriteLockReturns.result1, fake.acquireWriteLockReturns.result2
	}
}

func (fake *FakeLocker) AcquireWriteLockCallCount() int {
	fake.acquireWriteLockMutex.RLock()
	defer fake.acquireWriteLockMutex.RUnlock()
	return len(fake.acquireWriteLockArgsForCall)
}

func (fake *FakeLocker) AcquireWriteLockArgsForCall(i int) []string {
	fake.acquireWriteLockMutex.RLock()
	defer fake.acquireWriteLockMutex.RUnlock()
	return fake.acquireWriteLockArgsForCall[i].names
}

func (fake *FakeLocker) AcquireWriteLockReturns(result1 db.Lock, result2 error) {
	fake.AcquireWriteLockStub = nil
	fake.acquireWriteLockReturns = struct {
		result1 db.Lock
		result2 error
	}{result1, result2}
}

var _ radar.Locker = new(FakeLocker)
