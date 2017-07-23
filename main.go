package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ryanuber/columnize"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("please provide a file to chk")
	}

	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]

		// check it's a proper file we can touch and everything
		if _, err := os.Stat(arg); err != nil {
			continue
		}

		// get the hashes for the file
		hashes, err := hash(arg)
		if err != nil {
			// if anything happened, just skip to the next file
			continue
		}

		var lines []string

		// pop everything into lines to columnize the output
		lines = append(lines, fmt.Sprintf("file: | %s", arg))
		for key, value := range hashes {
			lines = append(lines, fmt.Sprintf("%s: | %s\n", key, value))
		}
		r := columnize.SimpleFormat(lines)
		fmt.Println(r)

		// if we are doing multiple files, separate them
		if len(os.Args)-1 > 1 {
			fmt.Println("---")
		}
	}
}

func hash(filename string) (map[string]string, error) {
	var hashes = make(map[string]string)

	f, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return hashes, err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Println(err)
		return hashes, err
	}
	hashes["md5"] = fmt.Sprintf("%x", h.Sum(nil))

	h2 := sha1.New()
	if _, err := io.Copy(h2, f); err != nil {
		log.Println(err)
		return hashes, err
	}
	hashes["sha1"] = fmt.Sprintf("%x", h2.Sum(nil))

	h3 := sha256.New()
	if _, err := io.Copy(h3, f); err != nil {
		log.Println(err)
		return hashes, err
	}
	hashes["sha256"] = fmt.Sprintf("%x", h3.Sum(nil))

	return hashes, nil
}
