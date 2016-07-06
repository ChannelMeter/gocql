package gocql

import "github.com/channelmeter/go-statsd-client/statsd"

type StatsLogger interface {
	Printf(format string, v ...interface{})
}

type StatsReporter struct {
	statter statsd.Statter
	slowLog StatsLogger
}

func (s *Session) reportingInit() {
	s.reporter.statter, _ = statsd.NewNoop()
}

func (s *Session) ReportingEnable(host string, prefix string) error {

	if statsdClient, err := statsd.New(host, prefix); err == nil {
		s.reporter.statter = statsd.Statter(statsdClient)
		return nil
	} else {
		return err
	}
}

func (s *Session) ReportingSlowLog(logger StatsLogger) {
	s.reporter.slowLog = logger
}
