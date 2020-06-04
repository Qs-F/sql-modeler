package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

const tmpfile = "models.sqlite"

func main() {
	logrus.SetOutput(os.Stderr)

	out := "models"
	if len(os.Args) > 1 {
		out = os.Args[1]
	}
	logrus.Println(out)

	dir, err := ioutil.TempDir("", "")
	if err != nil {
		logrus.Error("Cannot create temporary file")
		return
	}
	defer os.RemoveAll(dir)

	var buf bytes.Buffer
	sqlite := exec.Command("sqlite3", filepath.Join(dir, tmpfile))
	sqlite.Stdin = os.Stdin
	sqlite.Stdout = ioutil.Discard
	sqlite.Stderr = &buf
	if err := sqlite.Run(); err != nil {
		logrus.Error("Failed to run sqlite3:\n", err.Error(), buf.String())
		return
	}

	buf.Reset()

	sqlboiler := exec.Command("sqlboiler", "sqlite3", "--output", out, "--wipe")
	sqlboiler.Env = append(os.Environ(), "SQLITE3_DBNAME="+filepath.Join(dir, tmpfile))
	sqlboiler.Stdin = os.Stdin
	sqlboiler.Stdout = os.Stdout
	sqlboiler.Stderr = &buf
	if err := sqlboiler.Run(); err != nil {
		logrus.Error("Failed to run sqlboiler:\n", err.Error(), buf.String())
		return
	}
}
