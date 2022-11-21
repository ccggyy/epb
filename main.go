package main

import (
	"archive/zip"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	version = ""
)

func main() {
	// init cmd parameters
	eCmd := &ECmd{}
	err := eCmd.InitWithFlag()
	if err != nil {
		log.Fatal(err)
	}

	// init app information with package json file.
	app := &EApp{}
	err = app.InitWithJsonFile(eCmd.FullConfigFilePath())
	if err != nil {
		log.Fatal(err)
	}
	// increment version
	switch eCmd.VersionPart {
	case VersionPartLeading:
		err := app.IncreaseLeadingVersion()
		if err != nil {
			log.Fatal(err)
		}
	case VersionPartCenter:
		err := app.IncreaseCenterVersion()
		if err != nil {
			log.Fatal(err)
		}
	case VersionPartTrailing:
		err := app.IncreaseTrailingVersion()
		if err != nil {
			log.Fatal(err)
		}
	}
	app.IncrBuilds()
	// rewrite package json file
	err = app.WriteToFile(eCmd.FullConfigFilePath())
	if err != nil {
		log.Fatal(err)
	}

	// build zip file
	zFile, err := os.Create(path.Join(eCmd.OutputPath, app.OutputZipFileName()))
	if err != nil {
		log.Fatal(err)
	}

	archive := zip.NewWriter(zFile)
	defer func(archive *zip.Writer) {
		_ = archive.Close()
	}(archive)

	// zip bundle files
	err = handleZip(archive, eCmd.BundlePath)
	if err != nil {
		log.Fatal(err)
	}
	// zip dist files
	err = handleZip(archive, eCmd.DistPath)
	if err != nil {
		log.Fatal(err)
	}

	// done
	log.Println("Done.")
}

func handleZip(archive *zip.Writer, src string) error {
	return filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = strings.TrimPrefix(path, src)
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}

		return err
	})
}
