package anonymizer

import (
	"net"
	"testing"
)

func TestAnonymizeIPV4Local(t *testing.T) {
	ip := "127.0.0.1"
	out := AnonymizeIP(ip)
	if out != "127.0.0.0" {
		t.Errorf("Expected 127.0.0.0 got %s", out)
	}
}

func TestAnonymizeIPV4LocalPrivate(t *testing.T) {
	ip := "192.168.0.12"
	out := AnonymizeIP(ip)
	if out != "192.168.0.0" {
		t.Errorf("Expected 192.168.0.0 got %s", out)
	}
}

func TestAnonymizeIPV4Remote(t *testing.T) {
	ip := "216.58.207.67"
	out := AnonymizeIP(ip)
	if out != "216.58.207.0" {
		t.Errorf("Expected 216.58.207.0 got %s", out)
	}
}

func TestAnonymizeIPV6Local(t *testing.T) {
	ip := "::1"
	out := AnonymizeIP(ip)
	if out != "::" {
		t.Errorf("Expected :: got %s", out)
	}
}

func TestAnonymizeIPV6Remote(t *testing.T) {
	ip := "2001:db8:85a3:8d3:1319:8a2e:370:7348"
	out := AnonymizeIP(ip)
	if out != "2001:db8:85a3:8d3:" {
		t.Errorf("Expected 2001:db8:85a3:8d3: got %s", out)
	}
}

func TestAnonymizeIPV6RemoteShort(t *testing.T) {
	ip := "2001:db8:12"
	out := AnonymizeIP(ip)
	if out != "2001:db8:" {
		t.Errorf("Expected 2001:db8: got %s", out)
	}
}

func TestAnonymizeIPV4NetRemote(t *testing.T) {
	ip := net.ParseIP("216.58.207.67")
	out := AnonymizeNetIP(ip)
	if out != "216.58.207.0" {
		t.Errorf("Expected 216.58.207.0 got %s", out)
	}
}

func TestAnonymizeIPV6NetRemote(t *testing.T) {
	ip := net.ParseIP("2001:db8:85a3:8d3:1319:8a2e:370:7348")
	out := AnonymizeNetIP(ip)
	if out != "2001:db8:85a3:8d3:" {
		t.Errorf("Expected 2001:db8:85a3:8d3: got %s", out)
	}
}
