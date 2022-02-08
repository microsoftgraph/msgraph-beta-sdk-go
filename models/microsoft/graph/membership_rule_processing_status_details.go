package graph
import (
    "strings"
    "errors"
)
// 
type MembershipRuleProcessingStatusDetails int

const (
    NOTSTARTED_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS MembershipRuleProcessingStatusDetails = iota
    RUNNING_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
    FAILED_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
    SUCCEEDED_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
    UNSUPPORTEDFUTUREVALUE_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
)

func (i MembershipRuleProcessingStatusDetails) String() string {
    return []string{"NOTSTARTED", "RUNNING", "FAILED", "SUCCEEDED", "UNSUPPORTEDFUTUREVALUE"}[i]
}
func ParseMembershipRuleProcessingStatusDetails(v string) (interface{}, error) {
    result := NOTSTARTED_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
    switch strings.ToUpper(v) {
        case "NOTSTARTED":
            result = NOTSTARTED_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
        case "RUNNING":
            result = RUNNING_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
        case "FAILED":
            result = FAILED_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
        case "SUCCEEDED":
            result = SUCCEEDED_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
        case "UNSUPPORTEDFUTUREVALUE":
            result = UNSUPPORTEDFUTUREVALUE_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS
        default:
            return 0, errors.New("Unknown MembershipRuleProcessingStatusDetails value: " + v)
    }
    return &result, nil
}
func SerializeMembershipRuleProcessingStatusDetails(values []MembershipRuleProcessingStatusDetails) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
