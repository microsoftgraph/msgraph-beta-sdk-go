package graph
import (
    "strings"
    "errors"
)
type EncryptionReadinessState int

const (
    NOTREADY_ENCRYPTIONREADINESSSTATE EncryptionReadinessState = iota
    READY_ENCRYPTIONREADINESSSTATE
)

func (i EncryptionReadinessState) String() string {
    return []string{"NOTREADY", "READY"}[i]
}
func ParseEncryptionReadinessState(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NOTREADY":
            return NOTREADY_ENCRYPTIONREADINESSSTATE, nil
        case "READY":
            return READY_ENCRYPTIONREADINESSSTATE, nil
    }
    return 0, errors.New("Unknown EncryptionReadinessState value: " + v)
}
func SerializeEncryptionReadinessState(values []EncryptionReadinessState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
