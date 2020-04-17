package libs

import "testing"

func TestHostIpManager_LocalHostName(t *testing.T) {
	m := HostManager{}
	hostName, _ := m.LocalHostName()
	if hostName == "" {
		t.Errorf("HostIpManager.LocalHostName failed")
	}
}

func TestHostIpManager_LocalIp(t *testing.T) {
	m := HostManager{}
	ip, _ := m.LocalIp()
	if ip == "" {
		t.Errorf("HostIpManager.LocalIp failed")
	}
}

func TestHostManager_ExternalIp(t *testing.T) {
	m := HostManager{}
	extIp, _ := m.ExternalIp()
	if extIp == "" {
		t.Errorf("HostIpManager.ExternalIp failed")
	}
}