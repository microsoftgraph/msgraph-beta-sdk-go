package models
// BitLockerRecoveryInformationType types
type BitLockerRecoveryInformationType int

const (
    // Store recovery passwords and key packages.
    PASSWORDANDKEY_BITLOCKERRECOVERYINFORMATIONTYPE BitLockerRecoveryInformationType = iota
    // Store recovery passwords only.
    PASSWORDONLY_BITLOCKERRECOVERYINFORMATIONTYPE
)

func (i BitLockerRecoveryInformationType) String() string {
    return []string{"passwordAndKey", "passwordOnly"}[i]
}
func ParseBitLockerRecoveryInformationType(v string) (any, error) {
    result := PASSWORDANDKEY_BITLOCKERRECOVERYINFORMATIONTYPE
    switch v {
        case "passwordAndKey":
            result = PASSWORDANDKEY_BITLOCKERRECOVERYINFORMATIONTYPE
        case "passwordOnly":
            result = PASSWORDONLY_BITLOCKERRECOVERYINFORMATIONTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeBitLockerRecoveryInformationType(values []BitLockerRecoveryInformationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i BitLockerRecoveryInformationType) isMultiValue() bool {
    return false
}
