package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

const tmpfile = "models.sqlite"

func main() {
	logrus.SetOutput(os.Stderr)

	out := flag.String("o", "models", "Set output directly")
	sql := flag.String("sql", "schema.sql", "Set schema SQL file")
	pkg := flag.String("pkg", "models", "Set pkg name")
	flag.Parse()

	dir, err := ioutil.TempDir("", "")
	if err != nil {
		logrus.Error("Cannot create temporary file")
		return
	}
	defer os.RemoveAll(dir)

	var buf bytes.Buffer
	sqlite := exec.Command("sqlite3", filepath.Join(dir, tmpfile))
	f, err := os.Open(*sql)
	if err != nil {
		logrus.Error("Failed to open sql file:\n")
		return
	}
	sqlite.Stdin = f
	sqlite.Stdout = ioutil.Discard
	sqlite.Stderr = &buf
	if err := sqlite.Run(); err != nil {
		logrus.Error("Failed to run sqlite3:\n", err.Error(), buf.String())
		return
	}

	buf.Reset()

	os.Setenv("SQLITE3_DBNAME", filepath.Join(dir, tmpfile))
	sqlboiler := exec.Command("sqlboiler", "sqlite3", "--output", *out, "--wipe", "--pkgname", *pkg)
	sqlboiler.Env = append(sqlboiler.Env, os.Environ()...)
	sqlboiler.Stdin = os.Stdin
	sqlboiler.Stdout = os.Stdout
	sqlboiler.Stderr = &buf
	if err := sqlboiler.Run(); err != nil {
		logrus.Error("Failed to run sqlboiler:\n", err.Error(), buf.String())
		return
	}
}
