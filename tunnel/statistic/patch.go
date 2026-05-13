package statistic

type RequestNotify func(c Tracker)

var DefaultRequestNotify RequestNotify

func (m *Manager) TotalTraffic(onlyProxy bool) (up, down int64) {
	// smart manager has no dedicated proxy-only counters yet.
	return m.Total()
}

func (m *Manager) NowTraffic(onlyProxy bool) (up, down int64) {
	// smart manager has no dedicated proxy-only counters yet.
	return m.Now()
}
