package callrecords
import (
    "errors"
)
// 
type PstnUserBlockMode int

const (
    BLOCKED_PSTNUSERBLOCKMODE PstnUserBlockMode = iota
    UNBLOCKED_PSTNUSERBLOCKMODE
    UNKNOWNFUTUREVALUE_PSTNUSERBLOCKMODE
)

func (i PstnUserBlockMode) String() string {
    return []string{"blocked", "unblocked", "unknownFutureValue"}[i]
}
func ParsePstnUserBlockMode(v string) (any, error) {
    result := BLOCKED_PSTNUSERBLOCKMODE
    switch v {
        case "blocked":
            result = BLOCKED_PSTNUSERBLOCKMODE
        case "unblocked":
            result = UNBLOCKED_PSTNUSERBLOCKMODE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PSTNUSERBLOCKMODE
        default:
            return 0, errors.New("Unknown PstnUserBlockMode value: " + v)
    }
    return &result, nil
}
func SerializePstnUserBlockMode(values []PstnUserBlockMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PstnUserBlockMode) isMultiValue() bool {
    return false
}
