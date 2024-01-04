package main

import "fmt"

type Music struct {
	name  string
	plays int
}

func searchMostPlayed(musics []Music) int {
	most := 0

	for index, value := range musics {
		if value.plays > musics[most].plays {
			most = index
		}
	}

	return most
}

func ordered_selection() {
	musics := []Music{
		{name: "Rocket Skates", plays: 20},
		{name: "Fear of the Dark", plays: 10},
		{name: "SICKO MODE", plays: 40},
	}

	var ordered []Music

	for range musics {
		mostPlayed := searchMostPlayed(musics)
		ordered = append(ordered, musics[mostPlayed])
		musics = append(musics[:mostPlayed], musics[mostPlayed + 1 :]...)
	}

	fmt.Println("-->", ordered)
}