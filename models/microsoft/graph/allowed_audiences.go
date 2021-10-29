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
    switch strings.ToUpper(v) {
        case "ME":
            return ME_ALLOWEDAUDIENCES, nil
        case "FAMILY":
            return FAMILY_ALLOWEDAUDIENCES, nil
        case "CONTACTS":
            return CONTACTS_ALLOWEDAUDIENCES, nil
        case "GROUPMEMBERS":
            return GROUPMEMBERS_ALLOWEDAUDIENCES, nil
        case "ORGANIZATION":
            return ORGANIZATION_ALLOWEDAUDIENCES, nil
        case "FEDERATEDORGANIZATIONS":
            return FEDERATEDORGANIZATIONS_ALLOWEDAUDIENCES, nil
        case "EVERYONE":
            return EVERYONE_ALLOWEDAUDIENCES, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ALLOWEDAUDIENCES, nil
    }
    return 0, errors.New("Unknown AllowedAudiences value: " + v)
}
func SerializeAllowedAudiences(values []AllowedAudiences) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
