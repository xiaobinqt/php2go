package time

import "time"

// Sleep .
func Sleep(seconds int64) {
    time.Sleep(time.Duration(seconds) * time.Second)
}

// Usleep .
func Usleep(microseconds int64) {
    time.Sleep(time.Duration(microseconds) * time.Microsecond)
}

// TimeNanoSleep .
func TimeNanoSleep(seconds, nanoseconds int64) {
    time.Sleep(time.Duration(seconds)*time.Second + time.Duration(nanoseconds)*time.Nanosecond)
}
