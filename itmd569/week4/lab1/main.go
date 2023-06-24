package main

import "fmt"

type Song struct {
	name   string
	artist string
	next   *Song
}

type Playlist struct {
	name             string
	head             *Song // first song in the playlist
	currentlyPlaying *Song
}

func (playlist *Playlist) addSong(name string, artist string) error { 
	newSong := &Song{name, artist, nil}
	if playlist.head == nil { // if the playlist is empty, add the song as the first song
		playlist.head = newSong // set the head of the playlist to the new song
	} else {
		currentSong := playlist.head // start at the first song in the playlist
		for currentSong.next != nil { // loop through the songs until the end is reached (the next song is nil)
			currentSong = currentSong.next // set the current song to the next song
		}
		currentSong.next = newSong // add the new song to the end of the playlist
	}
	return nil
}

func (playlist *Playlist) listAllSongs() error {
	currentSong := playlist.head // start at the first song in the playlist

	if currentSong == nil { // if the playlist is empty, print out a message
		fmt.Println("The playlist is empty! You need to add some songs.")
		return nil
	}

	fmt.Println(currentSong.name)
	for currentSong.next != nil { // loop through the songs until the end is reached (the next song is nil)
		currentSong = currentSong.next
		fmt.Println(currentSong.name) // print out the song
	}
	return nil
}

func main() {
	playlist := &Playlist{name: "MyCoolPlayList"}
	
	playlist.addSong("SongName", "Artist")
	playlist.addSong("SongName2", "Artist2")

	playlist.listAllSongs()
}
