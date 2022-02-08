package graph
import (
    "strings"
    "errors"
)
// 
type EncryptionState int

const (
    NOTENCRYPTED_ENCRYPTIONSTATE EncryptionState = iota
    ENCRYPTED_ENCRYPTIONSTATE
)

func (i EncryptionState) String() string {
    return []string{"NOTENCRYPTED", "ENCRYPTED"}[i]
}
func ParseEncryptionState(v string) (interface{}, error) {
    result := NOTENCRYPTED_ENCRYPTIONSTATE
    switch strings.ToUpper(v) {
        case "NOTENCRYPTED":
            result = NOTENCRYPTED_ENCRYPTIONSTATE
        case "ENCRYPTED":
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
