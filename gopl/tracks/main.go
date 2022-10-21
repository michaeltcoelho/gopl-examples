package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type byArtist []*Track

func (p byArtist) Len() int           { return len(p) }
func (p byArtist) Less(i, j int) bool { return p[i].Artist < p[j].Artist }
func (p byArtist) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (p customSort) Len() int           { return len(p.t) }
func (p customSort) Less(i, j int) bool { return p.less(p.t[i], p.t[j]) }
func (p customSort) Swap(i, j int)      { p.t[i], p.t[j] = p.t[j], p.t[i] }

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Tittle", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "-----", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // Caculates columns and display table
}

func main() {
	sort.Sort(byArtist(tracks))

	printTracks(tracks)

	sort.Sort(sort.Reverse(byArtist(tracks)))

	printTracks(tracks)

	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Artist != y.Artist {
			return x.Artist < y.Artist
		}
		return false
	}})

	printTracks(tracks)
}
