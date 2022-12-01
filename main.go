package main

import (
	"fmt"
	"gips/ips"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) > 3 || len(os.Args) < 3 {
		fmt.Println("Usage: gips [patch] [rom]")
		os.Exit(1)
	}

	i, err := ips.New(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	rom_ext := filepath.Ext(os.Args[2])
	rom_root := filepath.Dir(os.Args[2])
	rom_name := strings.TrimSuffix(os.Args[2], rom_ext)

	if rom_root == "." {
		rom_root = rom_root[1:]
	}

	err = i.Apply(os.Args[2], rom_root+rom_name+"_patched"+rom_ext)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Patching done!")
}
