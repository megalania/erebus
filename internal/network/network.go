package network

import (
	"context"
	"io"
	"net"
	"net/http"
	"net/url"

	"github.com/refraction-networking/utls"
)

func GetPrivateIPv4Address() (net.IP, error) {
	udpConn, err := net.Dial(
		"udp",
		"1.1.1.1:80",
	)
	if err != nil {
		return nil, err
	}
	defer udpConn.Close()
	return udpConn.LocalAddr().(*net.UDPAddr).IP, nil
}

func GetPublicIPv4Address() (net.IP, error) {
	httpClient := &http.Client{
		Transport: &http.Transport{
			DialTLSContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				var netDialer net.Dialer
				tcpConn, err := netDialer.DialContext(
					ctx,
					network,
					addr,
				)
				if err != nil {
					return nil, err
				}
				tlsConn := tls.UClient(
					tcpConn,
					&tls.Config{
						InsecureSkipVerify: false,
						ServerName: "api.ipify.org",
					},
					tls.HelloRandomizedNoALPN,
				)
				err = tlsConn.HandshakeContext(
					ctx,
				)
				if err != nil {
					defer tcpConn.Close()
					return nil, err
				}
				return tlsConn, nil
			},
		},
	}
	httpResponse, err := httpClient.Do(
		&http.Request{
			Method: "GET",
			URL: &url.URL{
				Scheme: "https",
				Host: "api.ipify.org:443",
				Path: "/",
			},
		},
	)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()
	responseBody, err := io.ReadAll(
		httpResponse.Body,
	)
	return net.ParseIP(
		string(
			responseBody,
		),
	), nil
}
