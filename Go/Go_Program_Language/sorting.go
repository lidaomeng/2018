package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

// 结构定义
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

// 全局变量
var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m30s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("2m8s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

// 打印输出
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)

	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "-----", "-----", "-----", "-----")

	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}

	tw.Flush()
}

// 排序
type byArtist []*Track

func (p byArtist) Len() int {
	return len(p)
}

func (p byArtist) Less(i, j int) bool {
	return p[i].Artist < p[j].Artist
}

func (p byArtist) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// main 主函数
func main() {
	sort.Sort(byArtist(tracks))

	printTracks(tracks)
}
