package time_manager

import "time"

type TimeManager struct {
	DeltaTime   float64
	Accumulator float64
	lastTime    time.Time
}

func (t *TimeManager) Update() {
	now := time.Now()
	t.DeltaTime = now.Sub(t.lastTime).Seconds()
	t.lastTime = now

	t.Accumulator += t.DeltaTime
}

var Time = TimeManager{
	lastTime: time.Now(),
}
