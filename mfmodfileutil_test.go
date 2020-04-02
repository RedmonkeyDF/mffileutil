package mfmodfileutil

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestDirectoryExists_Direxists(t *testing.T) {

	_, testfile, _, _ := runtime.Caller(0)
	adir := filepath.Dir(testfile)

	fi, errstat := os.Stat(adir)
	if errstat != nil {

		if os.IsNotExist(errstat) {

			t.Fatalf("%s test failed.  Internal test error.  How the hell did you get here?", t.Name())
		} else {

			t.Fatalf("%s test failed.  Internal test error.  Unexpected test error \"%s\"", t.Name(), errstat)
		}
	}

	if fi.IsDir() {

		de, errde := DirectoryExists(adir)
		if errde != nil {

			t.Fatalf("%s test failed.  Unexpected error \"%s\" returned.", t.Name(), errde)
		}

		if !de {

			t.Fatalf("%s test failed.  Expected true.", t.Name())
		}

	} else {

		t.Fatalf("%s test failed.  Internal test error.  Again, how the hell did you get here?", t.Name())
	}
}

func TestDirectoryExists_DirDoesntExist(t *testing.T) {

	_, testfile, _, _ := runtime.Caller(0)
	adir := filepath.Join(filepath.Dir(testfile), "ishouldnotexist")

	_, errstat := os.Stat(adir)
	if errstat != nil {

		if os.IsNotExist(errstat) {

			de, errde := DirectoryExists(adir)
			if errde != nil {

				t.Fatalf("%s test failed.  Unexpected error \"%s\" returned.", t.Name(), errde)
			}

			if de {

				t.Fatalf("%s test failed.  Expected false.", t.Name())
			}

		} else {

			t.Fatalf("%s test failed.  Internal test error.  Unexpected test error \"%s\"", t.Name(), errstat)
		}
	} else {

		t.Fatalf("%s test failed.  Internal test error.  No error returned when one was expected.", t.Name())
	}
}

func TestRegularfileExists_FileExists(t *testing.T) {

	_, testfile, _, _ := runtime.Caller(0)

	fi, errstat := os.Stat(testfile)

	if errstat != nil {

		if os.IsNotExist(errstat) {

			t.Fatalf("%s - internal test error.  Wut?", t.Name())
		} else {

			t.Fatalf("%s - internal test error.  Unexpected error \"%s\".", t.Name(), errstat)
		}
	}

	if fi.Mode().IsDir() {

		t.Fatalf("%s - internal test error.  It's a directory.", t.Name())
	}

	if fi.Mode().IsRegular() {

		ex, errex := RegularfileExists(testfile)
		if errex != nil {

			t.Fatalf("%s - unexpected error \"%s\"", t.Name(), errex)
		}

		if !ex {

			t.Fatalf("%s - Expected true.", t.Name())
		}

	} else {

		t.Fatalf("%s - internal test error.  It's not a regular file.", t.Name())
	}
}

func TestRegularfileExists_FileDoesntExist(t *testing.T) {

	_, testfile, _, _ := runtime.Caller(0)

	thefile := filepath.Join(filepath.Dir(testfile), "ishouldnotexist.file")

	_, errstat := os.Stat(thefile)

	if errstat != nil {

		if os.IsNotExist(errstat) {

			ex, errex := RegularfileExists(thefile)
			if errex != nil {

				t.Fatalf("%s failed.  Unexpected error encountered \"%s\"", t.Name(), errex)
			}

			if ex {

				t.Fatalf("%s failed.  Expected false.", t.Name())
			}

		} else {

			t.Fatalf("%s - internal test error.  Unexpected error \"%s\".", t.Name(), errstat)
		}
	}
}