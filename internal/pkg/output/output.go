package output

// TODO: add log file output with mutex (probably add that to the logging library?)

import (
	"context"

	"github.com/SnaffCon/ShnafflerPoint/internal/pkg/logging"
)

var (
	loglevel      = TRACE
	interestlevel = GREEN
)

// Init the UI thread output structure
func Init(ctx context.Context, msgChan <-chan *Message, logLevel MsgClass, interestLevel TriageClass) {

	// set the lowestMsg level to output
	loglevel = logLevel
	interestlevel = interestLevel

	// TODO: figure out if we need more workers on the UI
	go messageWorker(ctx, msgChan)
}

// AdjustLogLevel which is returned in output
func AdjustLogLevel(logLevel MsgClass) {
	loglevel = logLevel
}

// AdjustInterestLevel which is returned in output
func AdjustInterestLevel(interestLevel TriageClass) {
	interestlevel = interestLevel
}

func messageWorker(ctx context.Context, msgChan <-chan *Message) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-msgChan:
			processMessage(msg)
		}
	}
}

func processMessage(msg *Message) {

	// ignore this if it falls below a threshold
	if msg.MsgType < loglevel {
		logging.Debug("skipped message of level " + string(msg.MsgType))
		return
	}
	if msg.Triaged < interestlevel {
		logging.Debug("skipped message of level " + string(msg.Triaged))
		return
	}

	// err check
	if msg.Err != nil {
		if msg.ErrType == FATAL {
			logging.Fatal(msg.Msg, msg.Err)
		} else if msg.ErrType == WARN {
			logging.Warn(msg.Msg)
		} else if msg.ErrType == SUPPRESS {
			return
		} else {
			// TODO: some sort of "how did we get here?" message or something, I dunnno
			return
		}
	}

	// break out by log level
	if msg.MsgType == TRACE {
		trace(msg)
	} else if msg.MsgType == DEGUB {
		degub(msg)
	} else if msg.MsgType == INFO {
		info(msg)
	} else if msg.MsgType == IMPORTANT {
		important(msg)
	}

}

func trace(msg *Message) {
	// TODO: add log file output with mutex
	logging.Warn("NOT IMPLEMENTED")
}

func degub(msg *Message) {
	// TODO: add log file output with mutex
	logging.Warn("NOT IMPLEMENTED")
}

func info(msg *Message) {
	// TODO: add log file output with mutex
	logging.Warn("NOT IMPLEMENTED")
}

func important(msg *Message) {
	// TODO: add log file output with mutex
	logging.Warn("NOT IMPLEMENTED")
}
