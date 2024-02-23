package repo

import (
	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr/v2"
	"os"
	"path/filepath"
)

const packPath = "../config"

func Initialize(repoRoot string) error {
	box := packr.New(packPath, packPath)
	if err := box.Walk(func(s string, file packd.File) error {
		p := filepath.Join(repoRoot, s)
		dir := filepath.Dir(p)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err := os.MkdirAll(dir, 0755)
			if err != nil {
				return err
			}
		}

		return os.WriteFile(p, []byte(file.String()), 0644)
	}); err != nil {
		return err
	}

	return nil
}
