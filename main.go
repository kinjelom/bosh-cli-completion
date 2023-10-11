package main

import (
	"bosh-cli-completion/cmd/completion"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshlogfile "github.com/cloudfoundry/bosh-utils/logger/file"
	boshsys "github.com/cloudfoundry/bosh-utils/system"

	bilog "github.com/cloudfoundry/bosh-cli/v7/logger"
	boshui "github.com/cloudfoundry/bosh-cli/v7/ui"
)

func main() {
	logger := newLogger()

	if completion.IsItCompletionCommand(os.Args[1:]) {
		// completion support
		blindUi := boshui.NewWrappingConfUI(&completion.BlindUI{}, logger) // only completion can write to stdout
		bc := completion.NewBoshComplete(blindUi, logger)
		if _, err := bc.Execute(os.Args[1:]); err != nil {
			fail(err, logger)
		}
	} else {
		// other commands
		fail(fmt.Errorf("unsuppored command"), logger)
	}
}

func newLogger() boshlog.Logger {
	level := boshlog.LevelNone

	logLevelString := os.Getenv("BOSH_LOG_LEVEL")

	if logLevelString != "" {
		var err error
		level, err = boshlog.Levelify(logLevelString)
		if err != nil {
			err = bosherr.WrapError(err, "Invalid BOSH_LOG_LEVEL value")
			logger := boshlog.NewLogger(boshlog.LevelError)
			fail(err, logger)
		}
	}

	logPath := os.Getenv("BOSH_LOG_PATH")
	if logPath != "" {
		return newSignalableFileLogger(logPath, level)
	}

	return newSignalableLogger(boshlog.NewLogger(level))
}

func newSignalableLogger(logger boshlog.Logger) boshlog.Logger {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)
	signalableLogger, _ := bilog.NewSignalableLogger(logger, c)
	return signalableLogger
}

func newSignalableFileLogger(logPath string, level boshlog.LogLevel) boshlog.Logger {
	// Log file logger errors to the STDERR logger
	logger := boshlog.NewLogger(boshlog.LevelError)
	fs := boshsys.NewOsFileSystem(logger)

	// Log file will be closed by process exit
	// Log file readable by all
	logfileLogger, _, err := boshlogfile.New(level, logPath, boshlogfile.DefaultLogFileMode, fs)
	if err != nil {
		logger := boshlog.NewLogger(boshlog.LevelError)
		fail(err, logger)
	}

	return newSignalableLogger(logfileLogger)
}

func fail(err error, logger boshlog.Logger) {
	if err != nil {
		logger.Error("CLI", err.Error())
	}
	os.Exit(1)
}
