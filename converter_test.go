package converter

import "testing"

func TestConvertIso2To3(t *testing.T) {
	isoCode3Digit := ConvertIso2To3("TR")

	if isoCode3Digit != "TUR" {
		t.Error("convert error")
	}

}

func TestConvertIso3To2(t *testing.T) {
	isoCode2Digit := ConvertIso3To2("TUR")

	if isoCode2Digit != "TR" {
		t.Error("convert error")
	}
}
