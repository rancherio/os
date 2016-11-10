package integration

import . "github.com/cpuguy83/check"

func (s *QemuSuite) TestNetworkFromUrl(c *C) {
	c.Parallel()
	netArgs := []string{"-net", "nic,vlan=0,model=virtio"}
	args := []string{"--cloud-config", "./tests/assets/test_10/cloud-config.yml"}
	for i := 0; i < 7; i++ {
		args = append(args, netArgs...)
	}
	err := s.RunQemu(c, args...)
	defer s.stopQemu(c)
	c.Assert(err, IsNil)

	s.CheckCall(c, `
cat > test-merge << "SCRIPT"
set -x -e

ip link show dev br0
ip link show dev br0.100 | grep br0.100@br0
ip link show dev eth1.100 | grep 'master br0'

SCRIPT
sudo bash test-merge`)

	s.CheckCall(c, `
cat > test-merge << "SCRIPT"
set -x -e

cat /etc/resolv.conf | grep "search mydomain.com example.com"
cat /etc/resolv.conf | grep "nameserver 208.67.222.123"
cat /etc/resolv.conf | grep "nameserver 208.67.220.123"

SCRIPT
sudo bash test-merge`)
}
