package mfmodfileutil

import "os"

func DirectoryExists(adirpath string) (bool, error) {

	di, errexist :=  os.Stat(adirpath)

	if errexist != nil {

		if os.IsNotExist(errexist) {

			return false, nil
		}

		return false, errexist
	}

	return di.IsDir(), nil
}

func RegularfileExists(afilepath string) (bool, error) {

	fi, errstat := os.Stat(afilepath)

	if errstat != nil {

		if os.IsNotExist(errstat) {

			return false, nil
		}

		return false, errstat
	}

	if fi.IsDir() {

		return false, nil
	}

	return fi.Mode().IsRegular(), nil
}