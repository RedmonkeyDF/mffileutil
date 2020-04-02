# MFMODFILEUTIL

Simple file utilities.  I wrote these as I was tired of typing the same code into my projects just to check 
if a directory or file exists.

`func DirectoryExists(adirpath string) (bool, error)`

Returns true if the directory exists, false if it does not.  Returns an error if an unexpected error is encountered.

`func RegularfileExists(afilepath string) (bool, error)`

Returns true if the file exists, and is a regular file (i.e. not a symlink or such).  Otherwise returns false.  Returns an error if an unexpected error is encountered.

#### TODO

Specialfileexists

Exists

???