package basic

import "time"

type WorkRequest struct {
	Name string
	Delay time.Duration
}
