package trash

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const (
	// _trashGlobPattern is the glob pattern to find all trash items
	_trashGlobPattern = "storage/users/spaces/*/*/trash/*/*/*/*"
)

// PurgeTrashOrphanedPaths purges orphaned paths in the trash
func PurgeTrashOrphanedPaths(p string, dryRun bool) error {
	// we have all trash nodes in all spaces now
	dirs, err := filepath.Glob(filepath.Join(p, _trashGlobPattern))
	if err != nil {
		return err
	}

	if len(dirs) == 0 {
		return errors.New("no trash found. Double check storage path")
	}

	for _, d := range dirs {
		if err := removeEmptyFolder(d, dryRun); err != nil {
			return err
		}
	}
	return nil
}

func removeEmptyFolder(path string, dryRun bool) error {
	if dryRun {
		f, err := os.ReadDir(path)
		if err != nil {
			return err
		}
		if len(f) < 1 {
			fmt.Println("would remove", path)
		}
		return nil
	}
	if err := os.Remove(path); err != nil {
		return nil
	}
	nd := filepath.Dir(path)
	if filepath.Base(nd) == "trash" {
		return nil
	}
	return removeEmptyFolder(nd, dryRun)
}
