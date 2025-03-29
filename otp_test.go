package main

import (
	"context"
	"testing"
	"time"
)


func TestRetentionMap_VerifyOTP(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)


	rm := NewRetentionMap(ctx, 1*time.Second)

	otp := rm.NewOTP()

	if ok := rm.VerifyOTP(otp.Key); !ok{
		t.Error("failed to verify otp key that exists")
	}
	if ok := rm.VerifyOTP(otp.Key); ok{
		t.Error("Reusing a OTP should not succeed")
	}


	cancel()
}

func TestOTP_Retention(t *testing.T) {


	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	rm := NewRetentionMap(ctx, 1*time.Second)

	rm.NewOTP()
	rm.NewOTP()

	time.Sleep(2 * time.Second)

	otp := rm.NewOTP()

	if len(rm) != 1 {
		t.Error("Failed to clean up")
	}

	if rm[otp.Key] != otp {
		t.Error("The key should still be in place")
	}
	cancel()
}
