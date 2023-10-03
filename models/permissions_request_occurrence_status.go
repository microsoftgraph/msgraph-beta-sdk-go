package models
import (
    "errors"
)
// 
type PermissionsRequestOccurrenceStatus int

const (
    GRANTINGFAILED_PERMISSIONSREQUESTOCCURRENCESTATUS PermissionsRequestOccurrenceStatus = iota
    GRANTED_PERMISSIONSREQUESTOCCURRENCESTATUS
    GRANTING_PERMISSIONSREQUESTOCCURRENCESTATUS
    REVOKED_PERMISSIONSREQUESTOCCURRENCESTATUS
    REVOKING_PERMISSIONSREQUESTOCCURRENCESTATUS
    REVOKINGFAILED_PERMISSIONSREQUESTOCCURRENCESTATUS
    UNKNOWNFUTUREVALUE_PERMISSIONSREQUESTOCCURRENCESTATUS
)

func (i PermissionsRequestOccurrenceStatus) String() string {
    return []string{"grantingFailed", "granted", "granting", "revoked", "revoking", "revokingFailed", "unknownFutureValue"}[i]
}
func ParsePermissionsRequestOccurrenceStatus(v string) (any, error) {
    result := GRANTINGFAILED_PERMISSIONSREQUESTOCCURRENCESTATUS
    switch v {
        case "grantingFailed":
            result = GRANTINGFAILED_PERMISSIONSREQUESTOCCURRENCESTATUS
        case "granted":
            result = GRANTED_PERMISSIONSREQUESTOCCURRENCESTATUS
        case "granting":
            result = GRANTING_PERMISSIONSREQUESTOCCURRENCESTATUS
        case "revoked":
            result = REVOKED_PERMISSIONSREQUESTOCCURRENCESTATUS
        case "revoking":
            result = REVOKING_PERMISSIONSREQUESTOCCURRENCESTATUS
        case "revokingFailed":
            result = REVOKINGFAILED_PERMISSIONSREQUESTOCCURRENCESTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PERMISSIONSREQUESTOCCURRENCESTATUS
        default:
            return 0, errors.New("Unknown PermissionsRequestOccurrenceStatus value: " + v)
    }
    return &result, nil
}
func SerializePermissionsRequestOccurrenceStatus(values []PermissionsRequestOccurrenceStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PermissionsRequestOccurrenceStatus) isMultiValue() bool {
    return false
}
