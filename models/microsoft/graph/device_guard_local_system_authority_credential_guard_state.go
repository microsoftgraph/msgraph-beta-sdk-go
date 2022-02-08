package graph
import (
    "strings"
    "errors"
)
// 
type DeviceGuardLocalSystemAuthorityCredentialGuardState int

const (
    RUNNING_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDSTATE DeviceGuardLocalSystemAuthorityCredentialGuardState = iota
    REBOOTREQUIRED_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDSTATE
    NOTLICENSED_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDSTATE
    NOTCONFIGURED_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDSTATE
    VIRTUALIZATIONBASEDSECURITYNOTRUNNING_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDSTATE
)

func (i DeviceGuardLocalSystemAuthorityCredentialGuardState) String() string {
    return []string{"RUNNING", "REBOOTREQUIRED", "NOTLICENSED", "NOTCONFIGURED", "VIRTUALIZATIONBASEDSECURITYNOTRUNNING"}[i]
}
func ParseDeviceGuardLocalSystemAuthorityCredentialGuardState(v string) (interface{}, error) {
    result := RUNNING_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDSTATE
    switch strings.ToUpper(v) {
        case "RUNNING":
            result = RUNNING_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDSTATE
        case "REBOOTREQUIRED":
            result = REBOOTREQUIRED_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDSTATE
        case "NOTLICENSED":
            result = NOTLICENSED_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDSTATE
        case "NOTCONFIGURED":
            result = NOTCONFIGURED_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDSTATE
        case "VIRTUALIZATIONBASEDSECURITYNOTRUNNING":
            result = VIRTUALIZATIONBASEDSECURITYNOTRUNNING_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDSTATE
        default:
            return 0, errors.New("Unknown DeviceGuardLocalSystemAuthorityCredentialGuardState value: " + v)
    }
    return &result, nil
}
func SerializeDeviceGuardLocalSystemAuthorityCredentialGuardState(values []DeviceGuardLocalSystemAuthorityCredentialGuardState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
