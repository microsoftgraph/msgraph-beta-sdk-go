package models
import (
    "errors"
)
type MfaFailureReasonCode int

const (
    MFAINCOMPLETE_MFAFAILUREREASONCODE MfaFailureReasonCode = iota
    MFADENIED_MFAFAILUREREASONCODE
    SYSTEMFAILURE_MFAFAILUREREASONCODE
    BADREQUEST_MFAFAILUREREASONCODE
    OTHER_MFAFAILUREREASONCODE
    UNKNOWNFUTUREVALUE_MFAFAILUREREASONCODE
)

func (i MfaFailureReasonCode) String() string {
    return []string{"mfaIncomplete", "mfaDenied", "systemFailure", "badRequest", "other", "unknownFutureValue"}[i]
}
func ParseMfaFailureReasonCode(v string) (any, error) {
    result := MFAINCOMPLETE_MFAFAILUREREASONCODE
    switch v {
        case "mfaIncomplete":
            result = MFAINCOMPLETE_MFAFAILUREREASONCODE
        case "mfaDenied":
            result = MFADENIED_MFAFAILUREREASONCODE
        case "systemFailure":
            result = SYSTEMFAILURE_MFAFAILUREREASONCODE
        case "badRequest":
            result = BADREQUEST_MFAFAILUREREASONCODE
        case "other":
            result = OTHER_MFAFAILUREREASONCODE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MFAFAILUREREASONCODE
        default:
            return 0, errors.New("Unknown MfaFailureReasonCode value: " + v)
    }
    return &result, nil
}
func SerializeMfaFailureReasonCode(values []MfaFailureReasonCode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MfaFailureReasonCode) isMultiValue() bool {
    return false
}
