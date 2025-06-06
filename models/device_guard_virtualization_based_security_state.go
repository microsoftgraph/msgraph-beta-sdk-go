// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
type DeviceGuardVirtualizationBasedSecurityState int

const (
    // Running
    RUNNING_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE DeviceGuardVirtualizationBasedSecurityState = iota
    // Root required
    REBOOTREQUIRED_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
    // 64 bit architecture required
    REQUIRE64BITARCHITECTURE_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
    // Not licensed
    NOTLICENSED_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
    // Not configured
    NOTCONFIGURED_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
    // System does not meet hardware requirements
    DOESNOTMEETHARDWAREREQUIREMENTS_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
    // Other. Event logs in microsoft-Windows-DeviceGuard have more details.
    OTHER_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
)

func (i DeviceGuardVirtualizationBasedSecurityState) String() string {
    return []string{"running", "rebootRequired", "require64BitArchitecture", "notLicensed", "notConfigured", "doesNotMeetHardwareRequirements", "other"}[i]
}
func ParseDeviceGuardVirtualizationBasedSecurityState(v string) (any, error) {
    result := RUNNING_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
    switch v {
        case "running":
            result = RUNNING_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
        case "rebootRequired":
            result = REBOOTREQUIRED_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
        case "require64BitArchitecture":
            result = REQUIRE64BITARCHITECTURE_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
        case "notLicensed":
            result = NOTLICENSED_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
        case "notConfigured":
            result = NOTCONFIGURED_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
        case "doesNotMeetHardwareRequirements":
            result = DOESNOTMEETHARDWAREREQUIREMENTS_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
        case "other":
            result = OTHER_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYSTATE
        default:
            return nil, nil
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
func (i DeviceGuardVirtualizationBasedSecurityState) isMultiValue() bool {
    return false
}
