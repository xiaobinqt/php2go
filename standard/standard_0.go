/**
 * @Author: v_bivwei
 * @Description:
 * @Date: 2021/1/1 19:14
 */
package standard

import "time"

func Sleep(seconds int64) {
    time.Sleep(time.Duration(seconds) * time.Second)
}

func Usleep(microseconds int64) {
    time.Sleep(time.Duration(microseconds) * time.Microsecond)
}

func TimeNanoSleep(seconds, nanoseconds int64) {
    time.Sleep(time.Duration(seconds)*time.Second + time.Duration(nanoseconds)*time.Nanosecond)
}
