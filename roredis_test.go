package roredis

import (
	"testing"
	"time"
)

const testKey = "tkey"
const testBogusKey = "bogusKey"
const testVal1 = "abc123"
const testKeyDB1 = "tkeyDB1"
const testVal1DB1 = "abc123DB1"

// A Redis instance is required for the tests here
var testCfg = RedisCfg{
	Host: "localhost",
	DB:   0,
}

func TestInitRedis(t *testing.T) {
	InitRedis(testCfg)
	if rclient == nil {
		t.Fatal("Failed to initialize Redis client")
	}
}

func TestPing(t *testing.T) {
	InitRedis(testCfg)
	ret := Ping()
	t.Log("Redis Ping returned", ret)
}

func TestSet(t *testing.T) {
	InitRedis(testCfg)
	err := Set(testKey, testVal1, 20*time.Second)
	if err != nil {
		t.Error("Set failed", err)
	}
}

func TestGetExistent(t *testing.T) {
	InitRedis(testCfg)
	TestSet(t)
	ret, err := Get(testKey)
	if err != nil {
		t.Error("Get failed", err)
		return
	}
	if ret != testVal1 {
		t.Error("Get: returned value did not match Set value")
	}
}

func TestGetNonExistent(t *testing.T) {
	InitRedis(testCfg)
	_, err := Get(testBogusKey)
	if err != nil {
		t.Log("Get failed for non-existent key", err)
		return
	}
}

func TestDel(t *testing.T) {
	InitRedis(testCfg)
	err := Del(testKey)
	if err != nil {
		t.Error("Delete failed", err)
	}
}

// Test Second DB

// A Redis instance is required for the tests here
var testCfgDB1 = RedisCfg{
	Host: "localhost",
	DB:   1,
}

func TestInitRedisDB1(t *testing.T) {
	InitRedis(testCfgDB1)
	if rclient == nil {
		t.Fatal("Failed to initialize Redis client")
	}
}

func TestPingDB1(t *testing.T) {
	InitRedis(testCfgDB1)
	ret := Ping()
	t.Log("Redis Ping returned", ret)
}

func TestSetDB1(t *testing.T) {
	InitRedis(testCfgDB1)
	err := Set(testKeyDB1, testVal1DB1, 0)
	if err != nil {
		t.Error("Set failed", err)
	}
}

func TestGetExistentDB1(t *testing.T) {
	InitRedis(testCfgDB1)
	TestSetDB1(t)
	ret, err := Get(testKeyDB1)
	if err != nil {
		t.Error("Get failed", err)
		return
	}
	if ret != testVal1DB1 {
		t.Error("Get: returned value did not match Set value")
	}
}

func TestGetNonExistentDB1(t *testing.T) {
	InitRedis(testCfgDB1)
	_, err := Get(testKey) // We should not see the key from DB0
	if err != nil {
		t.Log("Get failed for non-existent key", testKey, err)
		return
	}
}
