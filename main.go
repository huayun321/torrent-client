package main

import (
	"fmt"
	"log"
	"os"
	"torrent-client/torrentfile"
)

func main() {
	file, err := os.Open("debian-mac-11.3.0-amd64-netinst.iso.torrent")
	if err != nil {
		log.Fatalln(err)
	}
	bto, err := torrentfile.Open(file)
	if err != nil {
		log.Fatalln(err)
	}

	tf, err := bto.ToTorrentFile()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("pieces len ", len(tf.PieceHashes))
	fmt.Println("length len ", tf.Length)
	fmt.Println("name ", tf.Name)
	fmt.Println("CreationDate ", bto.CreationDate)
	//fmt.Printf("%+v", tf)
}
