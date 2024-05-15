package models
type ClaimConditionUserType int

const (
    ANY_CLAIMCONDITIONUSERTYPE ClaimConditionUserType = iota
    MEMBERS_CLAIMCONDITIONUSERTYPE
    ALLGUESTS_CLAIMCONDITIONUSERTYPE
    AADGUESTS_CLAIMCONDITIONUSERTYPE
    EXTERNALGUESTS_CLAIMCONDITIONUSERTYPE
    UNKNOWNFUTUREVALUE_CLAIMCONDITIONUSERTYPE
)

func (i ClaimConditionUserType) String() string {
    return []string{"any", "members", "allGuests", "aadGuests", "externalGuests", "unknownFutureValue"}[i]
}
func ParseClaimConditionUserType(v string) (any, error) {
    result := ANY_CLAIMCONDITIONUSERTYPE
    switch v {
        case "any":
            result = ANY_CLAIMCONDITIONUSERTYPE
        case "members":
            result = MEMBERS_CLAIMCONDITIONUSERTYPE
        case "allGuests":
            result = ALLGUESTS_CLAIMCONDITIONUSERTYPE
        case "aadGuests":
            result = AADGUESTS_CLAIMCONDITIONUSERTYPE
        case "externalGuests":
            result = EXTERNALGUESTS_CLAIMCONDITIONUSERTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLAIMCONDITIONUSERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeClaimConditionUserType(values []ClaimConditionUserType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ClaimConditionUserType) isMultiValue() bool {
    return false
}
