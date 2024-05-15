package models
// Possible values of Credential Guard settings.
type DeviceGuardLocalSystemAuthorityCredentialGuardType int

const (
    // Turns off Credential Guard remotely if configured previously without UEFI Lock.
    NOTCONFIGURED_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDTYPE DeviceGuardLocalSystemAuthorityCredentialGuardType = iota
    // Turns on Credential Guard with UEFI lock.
    ENABLEWITHUEFILOCK_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDTYPE
    // Turns on Credential Guard without UEFI lock.
    ENABLEWITHOUTUEFILOCK_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDTYPE
    // Disables Credential Guard. This is the default OS value.
    DISABLE_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDTYPE
)

func (i DeviceGuardLocalSystemAuthorityCredentialGuardType) String() string {
    return []string{"notConfigured", "enableWithUEFILock", "enableWithoutUEFILock", "disable"}[i]
}
func ParseDeviceGuardLocalSystemAuthorityCredentialGuardType(v string) (any, error) {
    result := NOTCONFIGURED_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDTYPE
        case "enableWithUEFILock":
            result = ENABLEWITHUEFILOCK_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDTYPE
        case "enableWithoutUEFILock":
            result = ENABLEWITHOUTUEFILOCK_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDTYPE
        case "disable":
            result = DISABLE_DEVICEGUARDLOCALSYSTEMAUTHORITYCREDENTIALGUARDTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDeviceGuardLocalSystemAuthorityCredentialGuardType(values []DeviceGuardLocalSystemAuthorityCredentialGuardType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DeviceGuardLocalSystemAuthorityCredentialGuardType) isMultiValue() bool {
    return false
}
