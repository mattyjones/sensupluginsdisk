package sensupluginsdisk

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInodeUsage(t *testing.T) {
	var alarm string
	var usage float64

	Convey("When checking the usage", t, func() {
		Convey("If the usage is 3.14%, the warning threshold is 4%, and the critical threshold is 6%", func() {
			usage = 3.14
			warnThreshold = 4
			critThreshold = 6
			alarm = CheckThreshold(usage, warnThreshold, critThreshold)

			Convey("The alarm should be ok", func() {
				So(alarm, ShouldEqual, "ok")
			})
			Convey("The alarm not should be 'warning'", func() {
				So(alarm, ShouldNotEqual, "warning")
			})
			Convey("The alarm should not be 'critical'", func() {
				So(alarm, ShouldNotEqual, "critical")
			})
		})

		Convey("If the usage is 4.14%, the warning threshold is 4, and the critical threshold is 6%", func() {

			usage = 4.14
			warnThreshold = 4
			critThreshold = 6
			alarm = CheckThreshold(usage, warnThreshold, critThreshold)

			Convey("The alarm should be not be ok", func() {
				So(alarm, ShouldNotEqual, "ok")
			})
			Convey("The alarm should be 'warning'", func() {
				So(alarm, ShouldEqual, "warning")
			})
			Convey("The alarm should not be 'critical'", func() {
				So(alarm, ShouldNotEqual, "critical")
			})
		})

		Convey("If the usage is 6.14%, the warning threshold is 4, and the critical threshold is 6%", func() {

			usage = 6.14
			warnThreshold = 4
			critThreshold = 6
			alarm = CheckThreshold(usage, warnThreshold, critThreshold)

			Convey("The alarm should be not be ok.", func() {
				So(alarm, ShouldNotEqual, "ok")
			})
			Convey("The alarm should not be 'warning'", func() {
				So(alarm, ShouldNotEqual, "warning")
			})
			Convey("The alarm should be 'critical'", func() {
				So(alarm, ShouldEqual, "critical")
			})
		})
	})
}

// Need to check exit code, status, and msg
