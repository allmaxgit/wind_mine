package main

type Watcher struct {
	OnNewRate func(string, float64)
}

func NewWatcher() *Watcher {
	return &Watcher{}
}
