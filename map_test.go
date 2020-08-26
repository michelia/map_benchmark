package main

import (
	"sync"
	"testing"
)

func BenchmarkMapM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(num)
		for i := 0; i < num; i++ {
			go func() {
				mm.GetRandom()
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkMapRWM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(num)
		for i := 0; i < num; i++ {
			go func() {
				rwm.GetRandom()
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkSyncMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(num)
		for i := 0; i < num; i++ {
			go func() {
				sm.GetRandom()
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkCMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(num)
		for i := 0; i < num; i++ {
			go func() {
				cm.GetRandom()
				wg.Done()
			}()
		}
		wg.Wait()
	}
}
