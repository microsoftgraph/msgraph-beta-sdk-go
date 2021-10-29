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
    switch strings.ToUpper(v) {
        case "NOTENCRYPTED":
            return NOTENCRYPTED_ENCRYPTIONSTATE, nil
        case "ENCRYPTED":
            return ENCRYPTED_ENCRYPTIONSTATE, nil
    }
    return 0, errors.New("Unknown EncryptionState value: " + v)
}
func SerializeEncryptionState(values []EncryptionState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
