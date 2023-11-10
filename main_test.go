package zeroconf_poc_test

import (
	"context"
	"fmt"
	"github.com/grandcat/zeroconf"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestZeroconf_ResolvesSMGW(t *testing.T) {
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		t.Fatal(err)
	}

	lookupChan := make(chan *zeroconf.ServiceEntry)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	err = resolver.Lookup(ctx, "smgw", "_theben._tcp", "local.", lookupChan)
	if err != nil {
		t.Fatal(err)
	}

	select {
	case <-ctx.Done():
		t.Fatal("Timed out waiting for service")
	case lookup := <-lookupChan:
		fmt.Printf("%#v", lookup)
		assert.NotNil(t, lookup)
		assert.Equal(t, "smgw", lookup.Instance)
		assert.Equal(t, "_theben._tcp", lookup.Service)
		assert.Equal(t, "local.", lookup.Domain)
		assert.Equal(t, 9999, lookup.Port)
		assert.NotNil(t, lookup.AddrIPv4)
	}
}

func TestZeroconf_ResolvesSMGWById(t *testing.T) {
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		t.Fatal(err)
	}

	lookupChan := make(chan *zeroconf.ServiceEntry)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	err = resolver.Lookup(ctx, "abc1234", "_theben._tcp", "local.", lookupChan)
	if err != nil {
		t.Fatal(err)
	}

	select {
	case <-ctx.Done():
		t.Fatal("Timed out waiting for service")
	case lookup := <-lookupChan:
		fmt.Printf("%#v", lookup)
		assert.NotNil(t, lookup)
		assert.Equal(t, "abc1234", lookup.Instance)
		assert.Equal(t, "_theben._tcp", lookup.Service)
		assert.Equal(t, "local.", lookup.Domain)
		assert.Equal(t, 1234, lookup.Port)
		assert.NotNil(t, lookup.AddrIPv4)
	}
}

func TestZeroconf_DoesntResolveInvalidName(t *testing.T) {
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		t.Fatal(err)
	}

	lookupChan := make(chan *zeroconf.ServiceEntry)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err = resolver.Lookup(ctx, "invalid", "_theben._tcp", "local.", lookupChan)
	if err != nil {
		t.Fatal(err)
	}

	select {
	case <-ctx.Done():
		assert.Error(t, ctx.Err())
	case lookup := <-lookupChan:
		fmt.Printf("%#v", lookup)
		t.Fatal("Erroneously resolved service")
	}
}

func TestZeroconf_ResolvesViaBrowse(t *testing.T) {
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		t.Fatal(err)
	}

	entries := make([]zeroconf.ServiceEntry, 0)

	lookupChan, lookupDone := context.WithCancel(context.Background())
	resultsChan := make(chan *zeroconf.ServiceEntry)
	go func(results <-chan *zeroconf.ServiceEntry) {
		for entry := range results {
			entries = append(entries, *entry)
			log.Println(entry)
		}
		lookupDone()
	}(resultsChan)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err = resolver.Browse(ctx, "_theben._tcp", "local.", resultsChan)
	if err != nil {
		t.Fatal(err)
	}

	select {
	case <-ctx.Done():
		t.Fatal("Timed out waiting for service")
	case <-lookupChan.Done():
		assert.NotEmpty(t, entries)
	}
}
