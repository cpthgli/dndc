package notificationForMac

import (
	"testing"
	"time"
)

var StartWaitTime = 10 * time.Second
var BetweenWaitTime = 5 * time.Second

func TestEnable(t *testing.T) {
	time.Sleep(StartWaitTime)
	err := Enable()
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(BetweenWaitTime)
	flag, err := IsEnable()
	t.Logf("%v, %T", flag, flag)
	if err != nil {
		t.Fatal(err)
	}
	if !flag {
		t.Fail()
	}
}
func TestDisable(t *testing.T) {
	time.Sleep(StartWaitTime)
	err := Disable()
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(BetweenWaitTime)
	flag, err := IsEnable()
	t.Logf("%v, %T", flag, flag)
	if err != nil {
		t.Fatal(err)
	}
	if flag {
		t.Fail()
	}
}
func TestToggle(t *testing.T) {
	time.Sleep(StartWaitTime)
	before, err := IsEnable()
	t.Logf("%v, %T", before, before)
	if err != nil {
		t.Fatal(err)
	}
	err = Toggle()
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(BetweenWaitTime)
	after, err := IsEnable()
	t.Logf("%v, %T", after, after)
	if err != nil {
		t.Fatal(err)
	}
	if before == after {
		t.Fail()
	}
}
func TestGetNotificationCenterUI(t *testing.T) {
	result, err := GetNotificationCenterUI()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v, %T", string(result), result)
}
func TestIsEnable(t *testing.T) {
	result, err := IsEnable()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v, %T", result, result)
}
