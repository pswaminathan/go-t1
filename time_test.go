package t1

// Copyright 2016 MediaMath <http://www.mediamath.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"testing"
	"time"
)

var notJSONEncodableTimes = []struct {
	time AdamaTime
	want string
}{
	{AdamaTime(time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC)), "Time.MarshalJSON: year outside of range [0,9999]"},
	{AdamaTime(time.Date(-1, 1, 1, 0, 0, 0, 0, time.UTC)), "Time.MarshalJSON: year outside of range [0,9999]"},
}

func TestNotJSONEncodableTime(t *testing.T) {
	for _, tt := range notJSONEncodableTimes {
		_, err := tt.time.MarshalJSON()
		if err == nil || err.Error() != tt.want {
			t.Errorf("%v MarshalJSON error = %v, want %v", tt.time, err, tt.want)
		}
	}
}
