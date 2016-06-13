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

var pList []string

type InodeDetails struct {
	InodesTotal       uint64
	InodesUsed        uint64
	InodesFree        uint64
	InodesUsedPercent float64
	Mountpoint        string
}

var inodes []InodeDetails

func ListPartitions() []string {
	var pl []string
	// Setting Partitions to true will also give you virtual and non-user filesystems.
	p, _ := disk.Partitions(false)

	for _, d := range p {
		pl = append(pl, d.Mountpoint)
	}
	return pl
}

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

			inodes = append(inodes, InodeInfo(p))

		}
		fmt.Println(pList)
		fmt.Println(inodes)
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
