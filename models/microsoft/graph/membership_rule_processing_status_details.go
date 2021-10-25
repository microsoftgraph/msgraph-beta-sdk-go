package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NOTSTARTED":
            return NOTSTARTED_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS, nil
        case "RUNNING":
            return RUNNING_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS, nil
        case "FAILED":
            return FAILED_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS, nil
        case "SUCCEEDED":
            return SUCCEEDED_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS, nil
        case "UNSUPPORTEDFUTUREVALUE":
            return UNSUPPORTEDFUTUREVALUE_MEMBERSHIPRULEPROCESSINGSTATUSDETAILS, nil
    }
    return 0, errors.New("Unknown MembershipRuleProcessingStatusDetails value: " + v)
}
func SerializeMembershipRuleProcessingStatusDetails(values []MembershipRuleProcessingStatusDetails) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
