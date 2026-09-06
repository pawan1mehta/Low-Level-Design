package main

type Result struct {
	Allowed    bool
	Remaining  int64
	RetryAfter int
}
