package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the compliance singleton.
type DeviceGuardVirtualizationBasedSecurityState int

const (
    RUNNING_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE DeviceGuardVirtualizationBasedSecurityState = iota
    REBOOTREQUIRED_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
    REQUIRE64BITARCHITECTURE_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
    NOTLICENSED_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
    NOTCONFIGURED_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
    DOESNOTMEETHARDWAREREQUIREMENTS_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
    OTHER_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
)

func (i DeviceGuardVirtualizationBasedSecurityState) String() string {
    return []string{"RUNNING", "REBOOTREQUIRED", "REQUIRE64BITARCHITECTURE", "NOTLICENSED", "NOTCONFIGURED", "DOESNOTMEETHARDWAREREQUIREMENTS", "OTHER"}[i]
}
func ParseDeviceGuardVirtualizationBasedSecurityState(v string) (interface{}, error) {
    result := RUNNING_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
    switch strings.ToUpper(v) {
        case "RUNNING":
            result = RUNNING_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
        case "REBOOTREQUIRED":
            result = REBOOTREQUIRED_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
        case "REQUIRE64BITARCHITECTURE":
            result = REQUIRE64BITARCHITECTURE_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
        case "NOTLICENSED":
            result = NOTLICENSED_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
        case "NOTCONFIGURED":
            result = NOTCONFIGURED_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
        case "DOESNOTMEETHARDWAREREQUIREMENTS":
            result = DOESNOTMEETHARDWAREREQUIREMENTS_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
        case "OTHER":
            result = OTHER_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
        default:
            return 0, errors.New("Unknown DeviceGuardVirtualizationBasedSecurityState value: " + v)
    }
    return &result, nil
}
func SerializeDeviceGuardVirtualizationBasedSecurityState(values []DeviceGuardVirtualizationBasedSecurityState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
