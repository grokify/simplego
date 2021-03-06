package severity

import (
	"testing"
)

var severityTests = []struct {
	execSeverity string
	itemSeverity string
	include      bool
	errorIsNil   bool
}{
	{SeverityDisabled, SeverityCritical, false, true},
	{SeverityError, SeverityDisabled, false, true},
	{SeverityError, SeverityWarning, false, true},
	{SeverityError, SeverityCritical, true, true},
	{SeverityDebug, SeverityWarning, true, true},
	{SeverityDebug, SeverityCritical, true, true},
	{SeverityCritical, SeverityError, false, true},
	{SeverityDebug, SeverityDisabled, false, true},
	{"foo", "bar", false, false},
}

func TestSeverity(t *testing.T) {
	n := len(severityTests)
	for i, tt := range severityTests {
		tryIncl, err := SeverityInclude(tt.execSeverity, tt.itemSeverity)
		if err != nil && tt.errorIsNil {
			t.Errorf("[%d/%d] severity.SeverityInclude(\"%s\",\"%s\") error [%v]", i+1, n, tt.execSeverity, tt.itemSeverity, err.Error())
		}
		if tryIncl != tt.include {
			t.Errorf("[%d/%d] severity.SeverityInclude(\"%s\",\"%s\") want [%v] got [%v]", i+1, n, tt.execSeverity, tt.itemSeverity, tt.include, tryIncl)
		}
	}
}
