package main

import (
	"container/list"
	"fmt"
)

type Playable interface {
	Duration() int64
}

type MusicTrack struct {
	trackName     string
	artistName    string
	trackDuration int64
}

func (m MusicTrack) Duration() int64 {
	return m.trackDuration
}

type PodcastEpisode struct {
	episodeName     string
	hostName        string
	episodeDuration int64
}

func (p PodcastEpisode) Duration() int64 {
	return p.episodeDuration
}

type AudioBook struct {
	bookName     string
	authorName   string
	bookDuration int64
}

func (a AudioBook) Duration() int64 {
	return a.bookDuration
}

type Playlist struct {
	items         *list.List
	wholeDuration int64
}

func (p *Playlist) Add(item Playable) {
	p.items.PushFront(item)
	p.wholeDuration += item.Duration()
}

func (p *Playlist) TotalDuration() int64 {
	return p.wholeDuration
}

func (p *Playlist) getWholeDuration() int64 {
	return p.wholeDuration
}

func (p *Playlist) removeCertainItem(item Playable) {
	for e := p.items.Front(); e != nil; e = e.Next() {
		if e.Value == item {
			p.items.Remove(e)
			p.wholeDuration -= item.Duration()
			break
		}
	}
}

func (p *Playlist) playAll() {
	for e := p.items.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func main() {
	var p = Playlist{items: list.New(), wholeDuration: 0}
	p.Add(MusicTrack{"test", "test", 25})
	podcast := PodcastEpisode{"test", "test", 25}
	p.Add(podcast)
	p.Add(AudioBook{"reck", "rock", 124})

	p.playAll()
	fmt.Println("test")
	p.removeCertainItem(podcast)

	p.playAll()

	fmt.Println(p.getWholeDuration())

	fmt.Println()
}