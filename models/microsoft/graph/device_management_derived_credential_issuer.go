package graph
import (
    "strings"
    "errors"
)
// 
type DeviceManagementDerivedCredentialIssuer int

const (
    INTERCEDE_DEVICEMANAGEMENTDERIVEDCREDENTIALISSUER DeviceManagementDerivedCredentialIssuer = iota
    ENTRUSTDATACARD_DEVICEMANAGEMENTDERIVEDCREDENTIALISSUER
    PUREBRED_DEVICEMANAGEMENTDERIVEDCREDENTIALISSUER
    XTEC_DEVICEMANAGEMENTDERIVEDCREDENTIALISSUER
)

func (i DeviceManagementDerivedCredentialIssuer) String() string {
    return []string{"INTERCEDE", "ENTRUSTDATACARD", "PUREBRED", "XTEC"}[i]
}
func ParseDeviceManagementDerivedCredentialIssuer(v string) (interface{}, error) {
    result := INTERCEDE_DEVICEMANAGEMENTDERIVEDCREDENTIALISSUER
    switch strings.ToUpper(v) {
        case "INTERCEDE":
            result = INTERCEDE_DEVICEMANAGEMENTDERIVEDCREDENTIALISSUER
        case "ENTRUSTDATACARD":
            result = ENTRUSTDATACARD_DEVICEMANAGEMENTDERIVEDCREDENTIALISSUER
        case "PUREBRED":
            result = PUREBRED_DEVICEMANAGEMENTDERIVEDCREDENTIALISSUER
        case "XTEC":
            result = XTEC_DEVICEMANAGEMENTDERIVEDCREDENTIALISSUER
        default:
            return 0, errors.New("Unknown DeviceManagementDerivedCredentialIssuer value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementDerivedCredentialIssuer(values []DeviceManagementDerivedCredentialIssuer) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
