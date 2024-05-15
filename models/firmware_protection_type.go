package models
// A list of possible Firmware protection type for a device. Firmware protection is a set of features that helps to ensure attackers can't get your device to start with untrusted or malicious firmware. Firmware protection type is determined by report sent from Microsoft Azure Attestation service. Only Windows 11 devices will have values "systemGuardSecureLaunch" or "firmwareAttackSurfaceReduction" or "disabled". Windows 10 devices will have value "notApplicable".
type FirmwareProtectionType int

const (
    // Indicates that the device is not a Windows 11 device.
    NOTAPPLICABLE_FIRMWAREPROTECTIONTYPE FirmwareProtectionType = iota
    // Indicates that System Guard Secure Launch is enabled for Firmware protection.
    SYSTEMGUARDSECURELAUNCH_FIRMWAREPROTECTIONTYPE
    // Indicates that Firmware Attack Surface Reduction is enabled for Firmware protection. This is only applicable to Surface devices.
    FIRMWAREATTACKSURFACEREDUCTION_FIRMWAREPROTECTIONTYPE
    // Indicates that the device has Firmware protection disabled.
    DISABLED_FIRMWAREPROTECTIONTYPE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_FIRMWAREPROTECTIONTYPE
)

func (i FirmwareProtectionType) String() string {
    return []string{"notApplicable", "systemGuardSecureLaunch", "firmwareAttackSurfaceReduction", "disabled", "unknownFutureValue"}[i]
}
func ParseFirmwareProtectionType(v string) (any, error) {
    result := NOTAPPLICABLE_FIRMWAREPROTECTIONTYPE
    switch v {
        case "notApplicable":
            result = NOTAPPLICABLE_FIRMWAREPROTECTIONTYPE
        case "systemGuardSecureLaunch":
            result = SYSTEMGUARDSECURELAUNCH_FIRMWAREPROTECTIONTYPE
        case "firmwareAttackSurfaceReduction":
            result = FIRMWAREATTACKSURFACEREDUCTION_FIRMWAREPROTECTIONTYPE
        case "disabled":
            result = DISABLED_FIRMWAREPROTECTIONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_FIRMWAREPROTECTIONTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeFirmwareProtectionType(values []FirmwareProtectionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i FirmwareProtectionType) isMultiValue() bool {
    return false
}
