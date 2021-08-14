package main

import "fmt"

type song struct{
	title string
	artist string

}
// buatlah struct song
// dengan title, artist string

type playlist struct {
	genre string
	songTitles []string
	songArtist []string
	songs []song
}

// buatlah playlist struct
// dengan genre string, songTitles []string
// songArtist []string, dan songs []song


func main() {

	// Buat dua variabel song1, dan song2 menggunakan struct song
    // misalkan artist BTS dengan 2 lagu yang berbeda
	var song1 = song{"Making Go", "BTS"}
	var song2 = song{"Making Lang", "BTS"}
	fmt.Printf("song1: %+v\nsong2: %+v\n", song1, song2)
	// copy song2 kedalam song1 
	song1 = song2
	
      // gunakan if untuk membandingkan title dan artis
      // jika sama print songs are equal
      // kalau tidak print songs are not equal
	  if song1.artist == song2.artist && song1.title == song2.title {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}

	// buatlah slice of songs
	sliceOfSongs := []song{
		song1,
		song2,
		{"making java", "BTS"},
	}
	// buatlah playlist misalkan rock dan masukkan
      // slice of songs tersebut ke dalam playlist rock
	rock := playlist{
		genre: "Rock",
		songTitles: []string{"A", "B"},
		songArtist: []string{"C", "D"},
		songs: sliceOfSongs,
	}


	// iterasi songs di dalam rock kemudian print
	for _,v := range rock.songs {
		fmt.Println(v)
	}
}