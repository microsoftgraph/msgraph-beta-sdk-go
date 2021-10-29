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
    switch strings.ToUpper(v) {
        case "INTERCEDE":
            return INTERCEDE_DEVICEMANAGEMENTDERIVEDCREDENTIALISSUER, nil
        case "ENTRUSTDATACARD":
            return ENTRUSTDATACARD_DEVICEMANAGEMENTDERIVEDCREDENTIALISSUER, nil
        case "PUREBRED":
            return PUREBRED_DEVICEMANAGEMENTDERIVEDCREDENTIALISSUER, nil
        case "XTEC":
            return XTEC_DEVICEMANAGEMENTDERIVEDCREDENTIALISSUER, nil
    }
    return 0, errors.New("Unknown DeviceManagementDerivedCredentialIssuer value: " + v)
}
func SerializeDeviceManagementDerivedCredentialIssuer(values []DeviceManagementDerivedCredentialIssuer) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
