package graph
import (
    "strings"
    "errors"
)
// 
type AllowedAudiences int

const (
    ME_ALLOWEDAUDIENCES AllowedAudiences = iota
    FAMILY_ALLOWEDAUDIENCES
    CONTACTS_ALLOWEDAUDIENCES
    GROUPMEMBERS_ALLOWEDAUDIENCES
    ORGANIZATION_ALLOWEDAUDIENCES
    FEDERATEDORGANIZATIONS_ALLOWEDAUDIENCES
    EVERYONE_ALLOWEDAUDIENCES
    UNKNOWNFUTUREVALUE_ALLOWEDAUDIENCES
)

func (i AllowedAudiences) String() string {
    return []string{"ME", "FAMILY", "CONTACTS", "GROUPMEMBERS", "ORGANIZATION", "FEDERATEDORGANIZATIONS", "EVERYONE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAllowedAudiences(v string) (interface{}, error) {
    result := ME_ALLOWEDAUDIENCES
    switch strings.ToUpper(v) {
        case "ME":
            result = ME_ALLOWEDAUDIENCES
        case "FAMILY":
            result = FAMILY_ALLOWEDAUDIENCES
        case "CONTACTS":
            result = CONTACTS_ALLOWEDAUDIENCES
        case "GROUPMEMBERS":
            result = GROUPMEMBERS_ALLOWEDAUDIENCES
        case "ORGANIZATION":
            result = ORGANIZATION_ALLOWEDAUDIENCES
        case "FEDERATEDORGANIZATIONS":
            result = FEDERATEDORGANIZATIONS_ALLOWEDAUDIENCES
        case "EVERYONE":
            result = EVERYONE_ALLOWEDAUDIENCES
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ALLOWEDAUDIENCES
        default:
            return 0, errors.New("Unknown AllowedAudiences value: " + v)
    }
    return &result, nil
}
func SerializeAllowedAudiences(values []AllowedAudiences) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
