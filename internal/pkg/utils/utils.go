package utils

import "github.com/SnaffCon/ShnafflerPoint/internal/pkg/logging"

// dumb utils to avoid good codding practices

// ErrCheck determines whether an error occurred
// returns true if an error, false otherwise while adding lots of overhead cause why not
// user ErrCheckLog to log the error
// use ErrCheckFatal to log and terminate
func ErrCheck(err error) bool {
	if err != nil {
		return true
	}
	return false
}

// ErrCheckLog will log an error if it exists and allow the program to continue
// returns true if an error was encountered
// Use ErrCheckFatal to log a fatal error
func ErrCheckLog(custom string, err error) bool {
	if err != nil {
		logging.Errr(custom, err)
		return true
	}
	return false
}

// ErrCheckFatal will log an error if it exists exit
func ErrCheckFatal(custom string, err error) {
	if err != nil {
		logging.Fatal(custom, err)
	}
}
