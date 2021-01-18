package roredis

import "testing"

const testKey = "tkey"
const testBogusKey = "bogusKey"
const testVal1 = "abc123"

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
	err := Set(testKey, testVal1, 0)
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
