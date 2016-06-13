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
	"fmt"

	"github.com/shirou/gopsutil/disk"
	"github.com/spf13/cobra"
)

// pList is a list of partitions tobe checked for inode usage
var pList []string

// InodeDetails represents detailed info about the inodess of a specific mountpoint
type InodeDetails struct {
	InodesTotal       uint64
	InodesUsed        uint64
	InodesFree        uint64
	InodesUsedPercent float64
	Mountpoint        string
}

// Represents all the mountpoints and that have been checkled
var inodes []*InodeDetails

// InodeInfo will return the inode details for a given mountpoint to later be checked
func (i InodeDetails) InodeInfo(d string) *InodeDetails {
	u, _ := disk.Usage(d)
	i.InodesTotal = u.InodesTotal
	i.InodesUsed = u.InodesUsed
	i.InodesFree = u.InodesFree
	i.InodesUsedPercent = u.InodesUsedPercent
	i.Mountpoint = d
	return &i
}

// checkDiskInodeCmd represents the checkDiskInode command
var checkDiskInodeCmd = &cobra.Command{
	Use:   "checkDiskInode",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(sensupluginsdisk *cobra.Command, args []string) {
		pList = ListPartitions()
		for _, p := range pList {

			in := new(InodeDetails)
			in = in.InodeInfo(p)
			inodes = append(inodes, in)

		}
		fmt.Println(pList)
		fmt.Println()
	},
}

func init() {
	RootCmd.AddCommand(checkDiskInodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkDiskInodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkDiskInodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
