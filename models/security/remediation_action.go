package security
import (
    "errors"
)
type RemediationAction int

const (
    MOVETOJUNK_REMEDIATIONACTION RemediationAction = iota
    MOVETOINBOX_REMEDIATIONACTION
    HARDDELETE_REMEDIATIONACTION
    SOFTDELETE_REMEDIATIONACTION
    MOVETODELETEDITEMS_REMEDIATIONACTION
    UNKNOWNFUTUREVALUE_REMEDIATIONACTION
)

func (i RemediationAction) String() string {
    return []string{"moveToJunk", "moveToInbox", "hardDelete", "softDelete", "moveToDeletedItems", "unknownFutureValue"}[i]
}
func ParseRemediationAction(v string) (any, error) {
    result := MOVETOJUNK_REMEDIATIONACTION
    switch v {
        case "moveToJunk":
            result = MOVETOJUNK_REMEDIATIONACTION
        case "moveToInbox":
            result = MOVETOINBOX_REMEDIATIONACTION
        case "hardDelete":
            result = HARDDELETE_REMEDIATIONACTION
        case "softDelete":
            result = SOFTDELETE_REMEDIATIONACTION
        case "moveToDeletedItems":
            result = MOVETODELETEDITEMS_REMEDIATIONACTION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_REMEDIATIONACTION
        default:
            return 0, errors.New("Unknown RemediationAction value: " + v)
    }
    return &result, nil
}
func SerializeRemediationAction(values []RemediationAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RemediationAction) isMultiValue() bool {
    return false
}
