package graph
import (
    "strings"
    "errors"
)
// 
type EncryptionReadinessState int

const (
    NOTREADY_ENCRYPTIONREADINESSSTATE EncryptionReadinessState = iota
    READY_ENCRYPTIONREADINESSSTATE
)

func (i EncryptionReadinessState) String() string {
    return []string{"NOTREADY", "READY"}[i]
}
func ParseEncryptionReadinessState(v string) (interface{}, error) {
    result := NOTREADY_ENCRYPTIONREADINESSSTATE
    switch strings.ToUpper(v) {
        case "NOTREADY":
            result = NOTREADY_ENCRYPTIONREADINESSSTATE
        case "READY":
            result = READY_ENCRYPTIONREADINESSSTATE
        default:
            return 0, errors.New("Unknown EncryptionReadinessState value: " + v)
    }
    return &result, nil
}
func SerializeEncryptionReadinessState(values []EncryptionReadinessState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
