package models
import (
    "errors"
)
// Encryption state
type EncryptionState int

const (
    // Not encrypted
    NOTENCRYPTED_ENCRYPTIONSTATE EncryptionState = iota
    // Encrypted
    ENCRYPTED_ENCRYPTIONSTATE
)

func (i EncryptionState) String() string {
    return []string{"notEncrypted", "encrypted"}[i]
}
func ParseEncryptionState(v string) (any, error) {
    result := NOTENCRYPTED_ENCRYPTIONSTATE
    switch v {
        case "notEncrypted":
            result = NOTENCRYPTED_ENCRYPTIONSTATE
        case "encrypted":
            result = ENCRYPTED_ENCRYPTIONSTATE
        default:
            return 0, errors.New("Unknown EncryptionState value: " + v)
    }
    return &result, nil
}
func SerializeEncryptionState(values []EncryptionState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i EncryptionState) isMultiValue() bool {
    return false
}
