package zeroconf_poc

import (
	"context"
	"github.com/brutella/dnssd"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDnssdZeroconf_ResolvesSMGW(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	lookupCtx, lookupCancel := context.WithTimeout(ctx, 10*time.Second)

	service, err := dnssd.LookupInstance(lookupCtx, "smgw._theben._tcp.local.")

	assert.NoError(t, err)
	assert.NotNil(t, service)

	assert.Equal(t, "smgw._theben._tcp.local", service.Name)
	assert.Equal(t, "_theben._tcp", service.Type)
	assert.Equal(t, "local.", service.Domain)
	assert.Equal(t, 9999, service.Port)
	assert.NotEmpty(t, service.IPs)

	defer cancel()
	defer lookupCancel()
}

//func TestDnssdZeroconf_ResolvesSMGWById(t *testing.T) {
//	resolver, err := zeroconf.NewResolver(nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	lookupChan := make(chan *zeroconf.ServiceEntry)
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
//	defer cancel()
//	err = resolver.Lookup(ctx, "abc1234", "_theben._tcp", "local.", lookupChan)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	select {
//	case <-ctx.Done():
//		t.Fatal("Timed out waiting for service")
//	case lookup := <-lookupChan:
//		fmt.Printf("%#v", lookup)
//		assert.NotNil(t, lookup)
//		assert.Equal(t, "abc1234", lookup.Instance)
//		assert.Equal(t, "_theben._tcp", lookup.Service)
//		assert.Equal(t, "local.", lookup.Domain)
//		assert.Equal(t, 1234, lookup.Port)
//		assert.NotNil(t, lookup.AddrIPv4)
//	}
//}
//
//func TestDnssdZeroconf_DoesntResolveInvalidName(t *testing.T) {
//	resolver, err := zeroconf.NewResolver(nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	lookupChan := make(chan *zeroconf.ServiceEntry)
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
//	defer cancel()
//	err = resolver.Lookup(ctx, "invalid", "_theben._tcp", "local.", lookupChan)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	select {
//	case <-ctx.Done():
//		assert.Error(t, ctx.Err())
//	case lookup := <-lookupChan:
//		fmt.Printf("%#v", lookup)
//		t.Fatal("Erroneously resolved service")
//	}
//}
