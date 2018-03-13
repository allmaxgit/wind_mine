package main

type Watcher struct {
	OnNewRate func(string, string, float64)
}

func NewWatcher() *Watcher {
	return &Watcher{}
}
