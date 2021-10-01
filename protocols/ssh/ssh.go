package ssh

import (
	"github.com/wzshiming/bridge"
	"github.com/wzshiming/sshproxy"
)

// SSH ssh://[username:password@]{address}[?identity_file=path/to/file]
func SSH(dialer bridge.Dialer, addr string) (bridge.Dialer, error) {
	d, err := sshproxy.NewDialer(addr)
	if err != nil {
		return nil, err
	}
	if dialer != nil {
		d.ProxyDial = dialer.DialContext
	}
	return d, nil
}
