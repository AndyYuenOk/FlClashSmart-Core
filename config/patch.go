package config

import (
	"sync"
)

var (
	proxyNameListMu sync.RWMutex
	proxyNameList   []string
)

func setProxyNameList(names []string) {
	proxyNameListMu.Lock()
	defer proxyNameListMu.Unlock()
	proxyNameList = append(proxyNameList[:0], names...)
}

func GetProxyNameList() []string {
	proxyNameListMu.RLock()
	defer proxyNameListMu.RUnlock()
	out := make([]string, len(proxyNameList))
	copy(out, proxyNameList)
	return out
}
