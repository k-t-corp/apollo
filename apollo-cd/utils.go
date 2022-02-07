package main

import (
	"archive/tar"
	"compress/gzip"
	"github.com/coreos/go-systemd/v22/dbus"
	"io"
	"os"
	"path/filepath"
)

// Untargz takes a destination path and a reader; a tar reader loops over the tarfile
// creating the file structure at 'dst' along the way, and writing any files
// https://gist.github.com/sdomino/635a5ed4f32c93aad131#file-untargz-go
func untargz(file string, dst string) error {
	fr, err := os.Open(file)
	if err != nil {
		return err
	}

	gzr, err := gzip.NewReader(fr)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()

		switch {

		// if no more files are found return
		case err == io.EOF:
			return nil

		// return any other error
		case err != nil:
			return err

		// if the header is nil, just skip it (not sure how this happens)
		case header == nil:
			continue
		}

		// the target location where the dir/file should be created
		target := filepath.Join(dst, header.Name)

		// the following switch could also be done using fi.Mode(), not sure if there
		// a benefit of using one vs. the other.
		// fi := header.FileInfo()

		// check the file type
		switch header.Typeflag {

		// if its a dir and it doesn't exist create it
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}

		// if it's a file create it
		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			// copy over contents
			if _, err := io.Copy(f, tr); err != nil {
				return err
			}

			// manually close here after each file operation; defering would cause each file close
			// to wait until all operations have completed.
			f.Close()
		}
	}
}

func systemctlStop(unitName string) error {
	conn, err := dbus.NewSystemdConnection()
	if err != nil {
		return nil
	}
	defer conn.Close()
	_, err = conn.StopUnit(unitName, "replace", nil)
	return err
}

func systemctlStart(unitName string) error {
	conn, err := dbus.NewSystemdConnection()
	if err != nil {
		return nil
	}
	defer conn.Close()
	_, err = conn.StartUnit(unitName, "replace", nil)
	return err
}

// https://github.com/gutengo/fil/blob/6109b2e0b5cfdefdef3a254cc1a3eaa35bc89284/file.go#L27-L34
func chownR(path string, uid, gid int) error {
	return filepath.Walk(path, func(name string, info os.FileInfo, err error) error {
		if err == nil {
			err = os.Chown(name, uid, gid)
		}
		return err
	})
}
