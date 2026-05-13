package tunnel

import C "github.com/metacubex/mihomo/constant"

func ProxiesWithProviders() map[string]C.Proxy {
	allProxies := make(map[string]C.Proxy)
	for name, proxy := range proxies {
		allProxies[name] = proxy
	}
	for _, p := range providers {
		for _, proxy := range p.Proxies() {
			name := proxy.Name()
			allProxies[name] = proxy
		}
	}
	return allProxies
}
