/**
Package gigasecond implements a function for adding a gigasecond to a Time.
*/
package gigasecond

import "time"

/**
Function gigasecond takes a Time object and returns that time with a gigasecond added to it.
*/
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1000000000 * time.Second)
}
