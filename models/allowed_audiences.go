package models
import (
    "errors"
    "strings"
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
    var values []string
    for p := AllowedAudiences(1); p <= UNKNOWNFUTUREVALUE_ALLOWEDAUDIENCES; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"me", "family", "contacts", "groupMembers", "organization", "federatedOrganizations", "everyone", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseAllowedAudiences(v string) (any, error) {
    var result AllowedAudiences
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "me":
                result |= ME_ALLOWEDAUDIENCES
            case "family":
                result |= FAMILY_ALLOWEDAUDIENCES
            case "contacts":
                result |= CONTACTS_ALLOWEDAUDIENCES
            case "groupMembers":
                result |= GROUPMEMBERS_ALLOWEDAUDIENCES
            case "organization":
                result |= ORGANIZATION_ALLOWEDAUDIENCES
            case "federatedOrganizations":
                result |= FEDERATEDORGANIZATIONS_ALLOWEDAUDIENCES
            case "everyone":
                result |= EVERYONE_ALLOWEDAUDIENCES
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_ALLOWEDAUDIENCES
            default:
                return 0, errors.New("Unknown AllowedAudiences value: " + v)
        }
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
func (i AllowedAudiences) isMultiValue() bool {
    return true
}
