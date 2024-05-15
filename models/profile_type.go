package models
// Profile Type
type ProfileType int

const (
    // Settings catalog profile type
    SETTINGSCATALOG_PROFILETYPE ProfileType = iota
    // Administrative Templates Profile Type
    ADMINISTRATIVETEMPLATES_PROFILETYPE
    // Imported Administrative Templates Profile Type
    IMPORTEDADMXTEMPLATES_PROFILETYPE
    // OEM Device related App Config Profile Type
    OEMAPPCONFIG_PROFILETYPE
    // Hardware Configuration Profile Type
    HARDWARECONFIG_PROFILETYPE
    // DCV1 Endpoint Protection Profile Type
    DCV1ENDPOINTPROTECTION_PROFILETYPE
    // DCV1 Device Restrictions Profile Type
    DCV1DEVICERESTRICTIONS_PROFILETYPE
)

func (i ProfileType) String() string {
    return []string{"settingsCatalog", "administrativeTemplates", "importedADMXTemplates", "oemAppConfig", "hardwareConfig", "dcV1EndpointProtection", "dcV1DeviceRestrictions"}[i]
}
func ParseProfileType(v string) (any, error) {
    result := SETTINGSCATALOG_PROFILETYPE
    switch v {
        case "settingsCatalog":
            result = SETTINGSCATALOG_PROFILETYPE
        case "administrativeTemplates":
            result = ADMINISTRATIVETEMPLATES_PROFILETYPE
        case "importedADMXTemplates":
            result = IMPORTEDADMXTEMPLATES_PROFILETYPE
        case "oemAppConfig":
            result = OEMAPPCONFIG_PROFILETYPE
        case "hardwareConfig":
            result = HARDWARECONFIG_PROFILETYPE
        case "dcV1EndpointProtection":
            result = DCV1ENDPOINTPROTECTION_PROFILETYPE
        case "dcV1DeviceRestrictions":
            result = DCV1DEVICERESTRICTIONS_PROFILETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeProfileType(values []ProfileType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ProfileType) isMultiValue() bool {
    return false
}
