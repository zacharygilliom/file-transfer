package system

import "testing"

func TestGetLatestDate(t *testing.T) {
	greaterDate := PhotosDateTime{1901, 1, 1}
	badReturn := PhotosDateTime{1900, 1, 1}
	result := GetLatestDate(badReturn, greaterDate)
	if badReturn == result {
		t.Fatalf(`GetLatestDate({1901, 1, 1}) = %q, want match for %#q`, result, greaterDate)
	}

}
