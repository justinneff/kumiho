package publishing

import (
	"io/fs"
	"path/filepath"

	"github.com/justinneff/kumiho/cache"
	"github.com/justinneff/kumiho/entities"
)

func getDatabaseObjectPaths(dbDir string) ([]string, error) {
	objectsDir := filepath.Join(dbDir, "objects")

	var objectFiles []string

	err := filepath.WalkDir(objectsDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if filepath.Ext(path) == ".sql" {
			objectFiles = append(objectFiles, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return objectFiles, nil
}

func resolveDependencies(obj *entities.DatabaseObject, otherObjects []*entities.DatabaseObject, provider entities.Provider) ([]string, error) {
	var deps []string
	for _, other := range otherObjects {
		if obj.Hash != other.Hash {
			matched, err := provider.IsDependency(obj.Content(), other.Schema, other.Name)
			if err != nil {
				return nil, err
			}
			if matched {
				deps = append(deps, other.FullName())
			}
		}
	}
	return deps, nil
}

func allDependenciesIncluded(dependencies []string, objects []*entities.DatabaseObject) bool {
	for _, dep := range dependencies {
		found := false

		for _, obj := range objects {
			if dep == obj.FullName() {
				found = true
			}
		}

		if !found {
			return false
		}
	}
	return true
}

func sortObjects(remaining []*entities.DatabaseObject, sorted []*entities.DatabaseObject) []*entities.DatabaseObject {
	if len(remaining) == 0 {
		return sorted
	}

	var nextRemaining []*entities.DatabaseObject

	for _, obj := range remaining {
		if allDependenciesIncluded(obj.Dependencies, sorted) {
			sorted = append(sorted, obj)
		} else {
			nextRemaining = append(nextRemaining, obj)
		}
	}

	return sortObjects(nextRemaining, sorted)
}

func LoadDatabaseObjects(dbDir string, provider entities.Provider) ([]*entities.DatabaseObject, error) {
	objectPaths, err := getDatabaseObjectPaths(dbDir)
	if err != nil {
		return nil, err
	}

	var objects []*entities.DatabaseObject

	for _, objPath := range objectPaths {
		obj, err := entities.NewDatabaseObject(objPath, provider)
		if err != nil {
			return nil, err
		}
		objects = append(objects, obj)
	}

	for i, obj := range objects {
		cacheObj, err := cache.ReadDatabaseObject(obj.Hash)
		if err != nil {
			return nil, err
		}

		if cacheObj != nil {
			objects[i].Dependencies = cacheObj.Dependencies
		} else {
			deps, err := resolveDependencies(obj, objects, provider)
			if err != nil {
				return nil, err
			}
			objects[i].Dependencies = deps

			cache.WriteDatabaseObject(*objects[i])
		}
	}

	var sortedObjects []*entities.DatabaseObject
	sortedObjects = sortObjects(objects, sortedObjects)

	return sortedObjects, nil
}
