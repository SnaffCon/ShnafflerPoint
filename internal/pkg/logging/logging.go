package logging

// TODO: add log file output with mutex

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/fatih/color"
)

var (
	// singleton
	one sync.Once

	// stuff for writing to a file
	file sync.Mutex
	loc  = pathfinder()

	// logger handles default
	h handles = handles{
		debug: setupLogger(os.Stdout, green, pDebug),
		log:   setupLogger(os.Stdout, green, pLog),
		warn:  setupLogger(os.Stdout, magenta, pWarn),
		error: setupLogger(os.Stderr, red, pErr),
	}

	// colours
	green   = color.New(color.FgGreen, color.Bold)
	yellow  = color.New(color.FgYellow, color.Bold)
	red     = color.New(color.FgRed, color.Bold)
	magenta = color.New(color.FgMagenta, color.Bold)
	blue    = color.New(color.FgBlue, color.Bold)

	// prefixes
	pDebug string = "[-] "
	pLog   string = "[*] "
	pWarn  string = "[!] "
	pErr   string = "[ERROR] "
)

type (
	// handles wraps a number of log handles
	handles struct {
		// logHandles
		debug *log.Logger
		log   *log.Logger
		warn  *log.Logger
		error *log.Logger
	}
)

func pathfinder() string {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("Unable to grok CWD. Are you using a computer?")
		os.Exit(1)
	}
	return path
}

// NewBasicLogger sets up logging based on the log leve
// fatality will ALWAYS output logs
// dateStamps = true will add dates to all log output
// timeStamps = true will add times to all log output
// logLevel >=4 outputs all logs; =3 outputs log,warn,error; =2 outputs warn,error; =1 outputs error, =0 suppresses all
func newLogger(debug io.Writer, log io.Writer, warn io.Writer, error io.Writer) {
	one.Do(func() {
		setupLoggers(debug, log, warn, error)
	})
}

func Debug(s string) {
	h.debug.Println(s)
}

func Info(s string) {
	h.log.Println(s)
}

func Warn(s string) {
	h.warn.Println(s)
}

func Errr(s string, err error) {
	if len(s) == 0 {
		h.error.Println(err)
	} else {
		h.error.Println(fmt.Sprintf("%s | %s", s, err))
	}
}

func Fatal(s string, err error) {
	if len(s) == 0 {
		h.error.Fatalln(err)
	} else {
		h.error.Fatalln(fmt.Sprintf("%s | %s", s, err))
	}
}

// ChangePrefixes sets up the log prefixes
// setting params to "" or "default" will use default values
func changePrefixes(debug string, log string, warn string, err string) {
	if debug == "default" || len(debug) == 0 {
		pDebug = "[-] "
	} else {
		pDebug = debug
	}
	if log == "default" || len(log) == 0 {
		pLog = "[*] "
	} else {
		pLog = log
	}
	if warn == "default" || len(warn) == 0 {
		pWarn = "[!] "
	} else {
		pWarn = warn
	}
	if err == "default" || len(err) == 0 {
		pErr = "[ERROR] "
	} else {
		pErr = err
	}
}

func setupLoggers(debug io.Writer, log io.Writer, warn io.Writer, error io.Writer) *handles {
	return &handles{
		debug: setupLogger(ioutil.Discard, green, pDebug),
		log:   setupLogger(os.Stdout, green, pLog),
		warn:  setupLogger(os.Stdout, magenta, pWarn),
		error: setupLogger(os.Stderr, red, pErr),
	}
}

func setupLogger(outputWriter io.Writer, colouriser *color.Color, prefixText string) (logger *log.Logger) {
	return log.New(outputWriter, colouriser.Sprintf(prefixText), 0)
}
