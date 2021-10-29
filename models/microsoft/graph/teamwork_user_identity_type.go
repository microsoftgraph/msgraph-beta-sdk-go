package graph
import (
    "strings"
    "errors"
)
// 
type TeamworkUserIdentityType int

const (
    AADUSER_TEAMWORKUSERIDENTITYTYPE TeamworkUserIdentityType = iota
    ONPREMISEAADUSER_TEAMWORKUSERIDENTITYTYPE
    ANONYMOUSGUEST_TEAMWORKUSERIDENTITYTYPE
    FEDERATEDUSER_TEAMWORKUSERIDENTITYTYPE
    PERSONALMICROSOFTACCOUNTUSER_TEAMWORKUSERIDENTITYTYPE
    SKYPEUSER_TEAMWORKUSERIDENTITYTYPE
    PHONEUSER_TEAMWORKUSERIDENTITYTYPE
    UNKNOWNFUTUREVALUE_TEAMWORKUSERIDENTITYTYPE
)

func (i TeamworkUserIdentityType) String() string {
    return []string{"AADUSER", "ONPREMISEAADUSER", "ANONYMOUSGUEST", "FEDERATEDUSER", "PERSONALMICROSOFTACCOUNTUSER", "SKYPEUSER", "PHONEUSER", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTeamworkUserIdentityType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "AADUSER":
            return AADUSER_TEAMWORKUSERIDENTITYTYPE, nil
        case "ONPREMISEAADUSER":
            return ONPREMISEAADUSER_TEAMWORKUSERIDENTITYTYPE, nil
        case "ANONYMOUSGUEST":
            return ANONYMOUSGUEST_TEAMWORKUSERIDENTITYTYPE, nil
        case "FEDERATEDUSER":
            return FEDERATEDUSER_TEAMWORKUSERIDENTITYTYPE, nil
        case "PERSONALMICROSOFTACCOUNTUSER":
            return PERSONALMICROSOFTACCOUNTUSER_TEAMWORKUSERIDENTITYTYPE, nil
        case "SKYPEUSER":
            return SKYPEUSER_TEAMWORKUSERIDENTITYTYPE, nil
        case "PHONEUSER":
            return PHONEUSER_TEAMWORKUSERIDENTITYTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_TEAMWORKUSERIDENTITYTYPE, nil
    }
    return 0, errors.New("Unknown TeamworkUserIdentityType value: " + v)
}
func SerializeTeamworkUserIdentityType(values []TeamworkUserIdentityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
