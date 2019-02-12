package kv_test

import (
	"context"
	"testing"

	"github.com/influxdata/influxdb"
	"github.com/influxdata/influxdb/kv"
	influxdbtesting "github.com/influxdata/influxdb/testing"
)

func TestBoltTelegrafService(t *testing.T) {
	influxdbtesting.TelegrafConfigStore(initBoltTelegrafService, t)
}

func TestInmemTelegrafService(t *testing.T) {
	influxdbtesting.TelegrafConfigStore(initInmemTelegrafService, t)
}

func initBoltTelegrafService(f influxdbtesting.TelegrafConfigFields, t *testing.T) (influxdb.TelegrafConfigStore, func()) {
	s, closeBolt, err := NewTestBoltStore()
	if err != nil {
		t.Fatalf("failed to create new kv store: %v", err)
	}

	svc, closeSvc := initTelegrafService(s, f, t)
	return svc, func() {
		closeSvc()
		closeBolt()
	}
}

func initInmemTelegrafService(f influxdbtesting.TelegrafConfigFields, t *testing.T) (influxdb.TelegrafConfigStore, func()) {
	s, closeBolt, err := NewTestInmemStore()
	if err != nil {
		t.Fatalf("failed to create new kv store: %v", err)
	}

	svc, closeSvc := initTelegrafService(s, f, t)
	return svc, func() {
		closeSvc()
		closeBolt()
	}
}

func initTelegrafService(s kv.Store, f influxdbtesting.TelegrafConfigFields, t *testing.T) (influxdb.TelegrafConfigStore, func()) {
	svc := kv.NewService(s)
	svc.IDGenerator = f.IDGenerator

	ctx := context.Background()
	if err := svc.Initialize(ctx); err != nil {
		t.Fatalf("error initializing user service: %v", err)
	}

	for _, tc := range f.TelegrafConfigs {
		if err := svc.PutTelegrafConfig(ctx, tc); err != nil {
			t.Fatalf("failed to populate telegraf config: %s", err.Error())
		}
	}

	for _, m := range f.UserResourceMappings {
		if err := svc.CreateUserResourceMapping(ctx, m); err != nil {
			t.Fatalf("failed to populate user resource mapping")
		}
	}

	return svc, func() {
		for _, tc := range f.TelegrafConfigs {
			if err := svc.DeleteTelegrafConfig(ctx, tc.ID); err != nil {
				t.Logf("failed to remove telegraf config: %v", err)
			}
		}
	}
}
