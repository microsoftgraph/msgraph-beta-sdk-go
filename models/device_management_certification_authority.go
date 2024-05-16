package models
// Device Management Certification Authority Types.
type DeviceManagementCertificationAuthority int

const (
    // Not configured.
    NOTCONFIGURED_DEVICEMANAGEMENTCERTIFICATIONAUTHORITY DeviceManagementCertificationAuthority = iota
    // Microsoft Certification Authority type.
    MICROSOFT_DEVICEMANAGEMENTCERTIFICATIONAUTHORITY
    // DigiCert Certification Authority type.
    DIGICERT_DEVICEMANAGEMENTCERTIFICATIONAUTHORITY
)

func (i DeviceManagementCertificationAuthority) String() string {
    return []string{"notConfigured", "microsoft", "digiCert"}[i]
}
func ParseDeviceManagementCertificationAuthority(v string) (any, error) {
    result := NOTCONFIGURED_DEVICEMANAGEMENTCERTIFICATIONAUTHORITY
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_DEVICEMANAGEMENTCERTIFICATIONAUTHORITY
        case "microsoft":
            result = MICROSOFT_DEVICEMANAGEMENTCERTIFICATIONAUTHORITY
        case "digiCert":
            result = DIGICERT_DEVICEMANAGEMENTCERTIFICATIONAUTHORITY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDeviceManagementCertificationAuthority(values []DeviceManagementCertificationAuthority) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DeviceManagementCertificationAuthority) isMultiValue() bool {
    return false
}
