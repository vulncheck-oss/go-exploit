package fileplant_test

import (
	"testing"

	"github.com/vulncheck-oss/go-exploit/payload/fileplant"
)

func TestEncodeSelfRemovingCron(t *testing.T) {
	cron, xploit := fileplant.Cron.SelfRemovingCron("root", "/etc/cron.d/hi", "/tmp/test", "id")

	if cron != "* * * * * root /bin/sh /tmp/test\n" {
		t.Fatal(cron)
	}

	if xploit != "#!/bin/sh\n\nrm -f /etc/cron.d/hi\nrm -f /tmp/test\nid\n" {
		t.Fatal(xploit)
	}
}
