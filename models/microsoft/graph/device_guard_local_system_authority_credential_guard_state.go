package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "RUNNING":
            return RUNNING_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDSTATE, nil
        case "REBOOTREQUIRED":
            return REBOOTREQUIRED_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDSTATE, nil
        case "NOTLICENSED":
            return NOTLICENSED_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDSTATE, nil
        case "NOTCONFIGURED":
            return NOTCONFIGURED_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDSTATE, nil
        case "VIRTUALIZATIONBASEDSECURITYNOTRUNNING":
            return VIRTUALIZATIONBASEDSECURITYNOTRUNNING_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDSTATE, nil
    }
    return 0, errors.New("Unknown DeviceGuardLocalSystemAuthorityCredentialGuardState value: " + v)
}
func SerializeDeviceGuardLocalSystemAuthorityCredentialGuardState(values []DeviceGuardLocalSystemAuthorityCredentialGuardState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
