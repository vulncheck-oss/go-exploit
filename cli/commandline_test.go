package cli

import (
	"testing"

	"github.com/vulncheck-oss/go-exploit/c2"
	"github.com/vulncheck-oss/go-exploit/config"
)

func TestCodeExecutionCmdLineParse(t *testing.T) {
	conf := config.New(config.CodeExecution, []c2.Impl{c2.SimpleShellServer}, "test product", "CVE-2023-1270", 1270)
	conf.Rhost = "rcetest"

	success := CodeExecutionCmdLineParse(conf)

	if conf.Rhost != "" {
		t.Fatal("Rhost should have no default value")
	}
	if conf.Rport != 1270 {
		t.Fatal("Rport should default to passed in value")
	}
	if conf.SSL != false {
		t.Fatal("SSL should default to false")
	}
	if conf.DoVerify != false {
		t.Fatal("verify should default to false")
	}
	if conf.DoVersionCheck != false {
		t.Fatal("version check should default to false")
	}
	if conf.DoExploit != false {
		t.Fatal("exploit should default to false")
	}
	if success != false {
		t.Fatal("parsing should have failed")
	}
	if conf.ThirdPartyC2Server != false {
		t.Fatal("outside should default to false")
	}
	if conf.C2Timeout != 30 {
		t.Fatal("timeout should default to 30")
	}
}

func TestCommonValidate(t *testing.T) {
	conf := config.New(config.CodeExecution, []c2.Impl{c2.SimpleShellServer}, "test product", "CVE-2023-1270", 1270)
	var rhosts string
	var rports string
	var rhostsFile string

	if commonValidate(conf, rhosts, rports, rhostsFile) {
		t.Fatal("commonValidate should fail with an empty Rhost")
	}

	conf.Rhost = "10.9.49.99"
	if commonValidate(conf, rhosts, rports, rhostsFile) {
		t.Fatal("commonValidate should fail with no supplied action")
	}

	conf.DoVerify = true
	if !commonValidate(conf, rhosts, rports, rhostsFile) {
		t.Fatal("commonValidate should succeed with rhost, rport, and doVerify")
	}

	conf.Rhost = ""
	if !commonValidate(conf, "127.0.0.1", "1270,1280", rhostsFile) {
		t.Fatal("commonValidate should have succeeded")
	}

	if !commonValidate(conf, "127.0.0.1,127.0.0.2", rports, rhostsFile) {
		t.Fatal("commonValidate have succeeded")
	}
}

func TestRhostsParsing(t *testing.T) {
	conf := config.New(config.CodeExecution, []c2.Impl{c2.SimpleShellServer}, "test product", "CVE-2023-1270", 1270)

	if !handleRhostsOptions(conf, "127.0.0.1,127.0.0.2", "80,443", "") {
		t.Fatal("commonValidate should succeed")
	}
	if len(conf.RhostsNTuple) != 4 {
		t.Fatal("Failed to parse rhosts")
	}
	if conf.RhostsNTuple[0].Rhost != "127.0.0.1" || conf.RhostsNTuple[1].Rhost != "127.0.0.1" ||
		conf.RhostsNTuple[2].Rhost != "127.0.0.2" || conf.RhostsNTuple[3].Rhost != "127.0.0.2" {
		t.Fatal("Failed to parse rhosts")
	}
	if conf.RhostsNTuple[0].Rport != 80 || conf.RhostsNTuple[1].Rport != 443 {
		t.Fatal("Failed to parse rports")
	}
	conf.RhostsNTuple = make([]config.RhostTriplet, 0)

	if !handleRhostsOptions(conf, "127.0.0.3", "443", "") {
		t.Fatal("commonValidate should succeed")
	}
	if len(conf.RhostsNTuple) != 1 {
		t.Fatal("Failed to parse rhosts")
	}
	if conf.RhostsNTuple[0].Rhost != "127.0.0.3" {
		t.Fatal("Failed to parse rhosts")
	}
	if conf.RhostsNTuple[0].Rport != 443 {
		t.Fatal("Failed to parse rports")
	}
	conf.RhostsNTuple = make([]config.RhostTriplet, 0)

	conf.Rhost = "127.0.0.4"
	if !handleRhostsOptions(conf, "", "443,80,8080", "") {
		t.Fatal("commonValidate should succeed")
	}
	if len(conf.RhostsNTuple) != 3 {
		t.Fatal("Failed to parse rhosts")
	}
	if conf.RhostsNTuple[0].Rhost != "127.0.0.4" {
		t.Fatal("Failed to parse rhosts")
	}
	if conf.RhostsNTuple[0].Rport != 443 {
		t.Fatal("Failed to parse rports")
	}
	if conf.RhostsNTuple[1].Rport != 80 {
		t.Fatal("Failed to parse rports")
	}
	if conf.RhostsNTuple[2].Rport != 8080 {
		t.Fatal("Failed to parse rports")
	}
	conf.Rhost = ""

	conf.RhostsNTuple = make([]config.RhostTriplet, 0)
	if !handleRhostsOptions(conf, "192.168.1.0/24", "80", "") {
		t.Fatal("commonValidate should succeed")
	}
	if len(conf.RhostsNTuple) != 256 {
		t.Fatal("Failed to parse rhosts")
	}

	conf.RhostsNTuple = make([]config.RhostTriplet, 0)
	if !handleRhostsOptions(conf, "192.168.1.0/24", "80,8080", "") {
		t.Fatal("commonValidate should succeed")
	}
	if len(conf.RhostsNTuple) != 512 {
		t.Fatal("Failed to parse rhosts")
	}
}
