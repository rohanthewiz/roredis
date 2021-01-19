package roredis

import (
	"testing"
	"time"
)

const testValsDuration = 20 * time.Second
const testKey = "testKey"
const testBogusKey = "bogusKey"
const testVal1 = "abc123"
const testKeyDB1 = "testKeyDB1"
const testVal1DB1 = "abc123DB1"

// A Redis instance is required for the tests here
var testCfg = RedisCfg{
	Host: "localhost",
	DB:   0,
}

// func TestInitRedis(t *testing.T) {
// 	InitRedis(testCfg)
// 	if RClient == nil {
// 		t.Fatal("Failed to initialize Redis client")
// 	}
// }

func TestPing(t *testing.T) {
	rc := InitRedis(testCfg)
	ret := Ping(rc)
	t.Log("Redis Ping returned", ret)
}

func TestSet(t *testing.T) {
	rc := InitRedis(testCfg)
	err := Set(rc, testKey, testVal1, testValsDuration)
	if err != nil {
		t.Error("Set failed", err)
	}
}

func TestGetExistent(t *testing.T) {
	rc := InitRedis(testCfg)
	TestSet(t)
	ret, err := Get(rc, testKey)
	if err != nil {
		t.Error("Get failed", err)
		return
	}
	if ret != testVal1 {
		t.Error("Get: returned value did not match Set value")
	}
}

func TestGetNonExistent(t *testing.T) {
	rc := InitRedis(testCfg)
	_, err := Get(rc, testBogusKey)
	if err != nil {
		t.Log("Get failed for non-existent key", err)
		return
	}
}

func TestDel(t *testing.T) {
	rc := InitRedis(testCfg)
	err := Del(rc, testKey)
	if err != nil {
		t.Error("Delete failed", err)
	}
}

// ---- Test Scan ---
func TestScanKeys(t *testing.T) {
	rc := InitRedis(testCfg)

	err := Set(rc, "roKey1", "roVal1", testValsDuration)
	if err != nil {
		t.Error("Set failed", "roKey1", err)
	}
	err = Set(rc, "roKey2", "roVal2", testValsDuration)
	if err != nil {
		t.Error("Set failed", "roKey2", err)
	}

	keys, err := Scan(rc, "roKey*")
	if err != nil {
		t.Error("Scan failed")
	} else {
		t.Log(len(keys), "keys found")
	}
}

// ---- Test Second DB ----

// A Redis instance is required for the tests here
var testCfgDB1 = RedisCfg{
	Host: "localhost",
	DB:   1,
}

func TestInitRedisDB1(t *testing.T) {
	rc := InitRedis(testCfgDB1)
	if rc == nil {
		t.Fatal("Failed to initialize Redis client")
	}
}

func TestPingDB1(t *testing.T) {
	rc := InitRedis(testCfgDB1)
	ret := Ping(rc)
	t.Log("Redis Ping returned", ret)
}

func TestSetDB1(t *testing.T) {
	rc := InitRedis(testCfgDB1)
	err := Set(rc, testKeyDB1, testVal1DB1, testValsDuration)
	if err != nil {
		t.Error("Set failed", err)
	}
}

func TestGetExistentDB1(t *testing.T) {
	rc := InitRedis(testCfgDB1)
	TestSetDB1(t)
	ret, err := Get(rc, testKeyDB1)
	if err != nil {
		t.Error("Get failed", err)
		return
	}
	if ret != testVal1DB1 {
		t.Error("Get: returned value did not match Set value")
	}
}

func TestGetNonExistentDB1(t *testing.T) {
	rc := InitRedis(testCfgDB1)
	_, err := Get(rc, testKey) // We should not see the key from DB0
	if err != nil {
		t.Log("Get failed for non-existent key", testKey, err)
		return
	}
}
