package main

type Watcher struct {
	OnNewRate func(uint8, string, float64)
}

func NewWatcher() *Watcher {
	return &Watcher{}
}
