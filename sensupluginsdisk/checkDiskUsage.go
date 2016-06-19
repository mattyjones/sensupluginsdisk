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

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/disk"
	"github.com/spf13/cobra"
	"github.com/yieldbot/sensuplugin/sensuutil"
)

// UsageDetails represents detailed info about the disk usage of a specific mountpoint
type UsageDetails struct {
	Total       uint64
	Used        uint64
	Free        uint64
	UsedPercent float64
	Mountpoint  string
}

// Represents all the mountpoints and that have been checkled
var disks []*UsageDetails

// DiskInfo will return the disk usage details for a given mountpoint to later be checked
func (d UsageDetails) DiskInfo(p string) *UsageDetails {
	u, _ := disk.Usage(p)
	d.Total = u.Total
	d.Used = u.Used
	d.Free = u.Free
	d.UsedPercent = u.UsedPercent
	d.Mountpoint = p
	return &d
}

// checkDiskUsageCmd represents the checkDiskUsage command
var checkDiskUsageCmd = &cobra.Command{
	Use:   "checkDiskUsage",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		condition := ""
		var msg, pList []string

		pList = ListPartitions()
		for _, p := range pList {

			du := new(UsageDetails)
			du = du.DiskInfo(p)
			disks = append(disks, du)

		}
		// Need to make sure the highest condition rules the day
		// Would be nice to give the threshold as well.
		for _, d := range disks {
			condition = CheckThreshold(d.UsedPercent, warnThreshold, critThreshold)
			if condition != ok {
				var buffer bytes.Buffer
				buffer.WriteString(d.Mountpoint)
				buffer.WriteString(" is above the threshold. Usage is")
				buffer.WriteString(" ")
				buffer.WriteString(strconv.FormatUint(d.Used, 10))
				buffer.WriteString("/")
				buffer.WriteString(strconv.FormatUint(d.Total, 10))

				msg = append(msg, buffer.String())
			}
		}

		// exit with the right code
		switch condition {
		case warning:
			sensuutil.Exit(condition, strings.Join(msg, ","))
		case critical:
			sensuutil.Exit(condition, strings.Join(msg, ","))
		default:
			sensuutil.Exit(condition)
		}
	},
}

func init() {
	RootCmd.AddCommand(checkDiskUsageCmd)

}
