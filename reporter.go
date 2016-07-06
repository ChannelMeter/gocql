package gocql

import "github.com/channelmeter/go-statsd-client/statsd"

type StatsReporter struct {
	statter statsd.Statter
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
