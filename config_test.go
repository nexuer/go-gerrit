package gerrit

import (
	"context"
	"testing"
)

func TestConfigService_GetVersion(t *testing.T) {
	client := NewClient(testPasswordCredential, &Options{
		Debug: true,
	})

	reply, err := client.Config.GetVersion(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("version: %v", reply)
}

func TestConfigService_GetServerInfo(t *testing.T) {
	client := NewClient(testPasswordCredential, &Options{
		Debug: true,
	})

	reply, err := client.Config.GetServerInfo(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("info: %+v", reply)
}
