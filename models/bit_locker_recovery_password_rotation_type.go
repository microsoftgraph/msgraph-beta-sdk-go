package models
// BitLocker recovery password rotation type
type BitLockerRecoveryPasswordRotationType int

const (
    // Not configured
    NOTCONFIGURED_BITLOCKERRECOVERYPASSWORDROTATIONTYPE BitLockerRecoveryPasswordRotationType = iota
    // Recovery password rotation off
    DISABLED_BITLOCKERRECOVERYPASSWORDROTATIONTYPE
    // Recovery password rotation on for Azure AD joined devices
    ENABLEDFORAZUREAD_BITLOCKERRECOVERYPASSWORDROTATIONTYPE
    // Recovery password rotation on for both Azure AD joined and hybrid joined devices
    ENABLEDFORAZUREADANDHYBRID_BITLOCKERRECOVERYPASSWORDROTATIONTYPE
)

func (i BitLockerRecoveryPasswordRotationType) String() string {
    return []string{"notConfigured", "disabled", "enabledForAzureAd", "enabledForAzureAdAndHybrid"}[i]
}
func ParseBitLockerRecoveryPasswordRotationType(v string) (any, error) {
    result := NOTCONFIGURED_BITLOCKERRECOVERYPASSWORDROTATIONTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_BITLOCKERRECOVERYPASSWORDROTATIONTYPE
        case "disabled":
            result = DISABLED_BITLOCKERRECOVERYPASSWORDROTATIONTYPE
        case "enabledForAzureAd":
            result = ENABLEDFORAZUREAD_BITLOCKERRECOVERYPASSWORDROTATIONTYPE
        case "enabledForAzureAdAndHybrid":
            result = ENABLEDFORAZUREADANDHYBRID_BITLOCKERRECOVERYPASSWORDROTATIONTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeBitLockerRecoveryPasswordRotationType(values []BitLockerRecoveryPasswordRotationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i BitLockerRecoveryPasswordRotationType) isMultiValue() bool {
    return false
}
