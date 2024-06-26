package main

import (
	note "easy_note/kitex_gen/note/note/noteservice"
	"log"
)

func main() {
	svr := note.NewServer(new(NoteServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
