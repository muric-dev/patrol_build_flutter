package build_constants

import "testing"

func TestPlatformConstants(t *testing.T) {
	// GIVEN platform constants
	// WHEN validating their values
	if PlatformAndroid != "android" {
		t.Errorf("expected PlatformAndroid to be 'android', got %q", PlatformAndroid)
	}
	if PlatformIOS != "ios" {
		t.Errorf("expected PlatformIOS to be 'ios', got %q", PlatformIOS)
	}
	if PlatformBoth != "both" {
		t.Errorf("expected PlatformBoth to be 'both', got %q", PlatformBoth)
	}
	// THEN values should match expected platform strings
}
