package api

import (
	"strings"
	"testing"

	"github.com/SparkPost/go-sparkpost/test"
)

func TestRecipients(t *testing.T) {
	cfgMap, err := test.LoadConfig()
	if err != nil {
		t.Error(err)
		return
	}
	cfg, err := NewConfig(cfgMap)
	if err != nil {
		t.Error(err)
		return
	}

	var client Client
	err = client.Init(cfg)
	if err != nil {
		t.Error(err)
		return
	}

	list, _, err := client.RecipientLists()
	if err != nil {
		t.Error(err)
		return
	}

	strs := make([]string, len(*list))
	for idx, rl := range *list {
		strs[idx] = rl.String()
	}
	t.Errorf("%s\n", strings.Join(strs, "\n"))
}
