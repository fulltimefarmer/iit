package main

import "fmt"

type Song struct {
	name   string
	artist string
	next   *Song
}

func main() {
	_, _, _, _, artists := createArtistMap()
	printArtistInfo(artists)
}

func printArtistInfo(artists map[string][]*Song) {
	// print out each artist in the map, and all of the songs that belong to that artist
	fmt.Println("")
	for artist, songs := range artists {
		fmt.Printf("Artist: %s\n", artist)
		for _, song := range songs {
			fmt.Printf("%s\n", song.name)
		}
		fmt.Println("")
	}
}

func createArtistMap() ([]*Song, []*Song, []*Song, []*Song, map[string][]*Song) {
	artists := make(map[string][]*Song)

	bobDylanSongs := make([]*Song, 0)
	song0 := &Song{name: "Like a Rolling Stone", artist: "Bob Dylan"}
	song1 := &Song{name: "Gimme Shelter", artist: "Bob Dylan"}
	song2 := &Song{name: "Hurricane", artist: "Bob Dylan"}
	bobDylanSongs = append(bobDylanSongs, song0, song1, song2)

	artists["Bob Dylan"] = bobDylanSongs

	rollingStonesSongs := make([]*Song, 0)
	song3 := &Song{name: "Paint it Black", artist: "Rolling Stones"}
	song4 := &Song{name: "Gimme Shelter", artist: "Rolling Stones"}
	song5 := &Song{name: "Angie", artist: "Rolling Stones"}
	rollingStonesSongs = append(rollingStonesSongs, song3, song4, song5)

	artists["Rolling Stones"] = rollingStonesSongs

	beatlesSongs := make([]*Song, 0)
	song6 := &Song{name: "Hey Jude", artist: "The Beatles"}
	song7 := &Song{name: "Come Together", artist: "The Beatles"}
	song8 := &Song{name: "Here Comes the Sun", artist: "The Beatles"}
	beatlesSongs = append(beatlesSongs, song6, song7, song8)

	artists["The Beatles"] = beatlesSongs

	nirvanaSongs := make([]*Song, 0)
	song9 := &Song{name: "Smells Like Teen Spirit", artist: "Nirvana"}
	song10 := &Song{name: "Come as You Are", artist: "Nirvana"}
	song11 := &Song{name: "About a Girl", artist: "Nirvana"}
	nirvanaSongs = append(nirvanaSongs, song9, song10, song11)

	artists["Nirvana"] = nirvanaSongs

	return bobDylanSongs, rollingStonesSongs, beatlesSongs, nirvanaSongs, artists
}
