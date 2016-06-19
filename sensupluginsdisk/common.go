// Copyright © 2016 Yieldbot <devops@yieldbot.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package sensupluginsdisk

import "github.com/shirou/gopsutil/disk"

var ok = "ok"
var warning = "warning"
var critical = "critical"

// ListPartitions will generate a slice of all partitions on a system.
// Setting Partitions to true will also give you virtual and non-user filesystems.
func ListPartitions() []string {
	var pl []string
	p, _ := disk.Partitions(false)

	for _, d := range p {
		pl = append(pl, d.Mountpoint)
	}
	return pl
}

// CheckThreshold will determine if a given number is greater than either a
// warning or critical threshold. It will return an exit code that can be
func CheckThreshold(val float64, warnT float64, critT float64) string {
	switch {
	case val >= critT:
		return critical
	case val >= warnT:
		return warning
	default:
		return ok
	}
}
