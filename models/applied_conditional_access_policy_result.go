package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the auditLogRoot singleton.
type AppliedConditionalAccessPolicyResult int

const (
    SUCCESS_APPLIEDCONDITIONALACCESSPOLICYRESULT AppliedConditionalAccessPolicyResult = iota
    FAILURE_APPLIEDCONDITIONALACCESSPOLICYRESULT
    NOTAPPLIED_APPLIEDCONDITIONALACCESSPOLICYRESULT
    NOTENABLED_APPLIEDCONDITIONALACCESSPOLICYRESULT
    UNKNOWN_APPLIEDCONDITIONALACCESSPOLICYRESULT
    UNKNOWNFUTUREVALUE_APPLIEDCONDITIONALACCESSPOLICYRESULT
    REPORTONLYSUCCESS_APPLIEDCONDITIONALACCESSPOLICYRESULT
    REPORTONLYFAILURE_APPLIEDCONDITIONALACCESSPOLICYRESULT
    REPORTONLYNOTAPPLIED_APPLIEDCONDITIONALACCESSPOLICYRESULT
    REPORTONLYINTERRUPTED_APPLIEDCONDITIONALACCESSPOLICYRESULT
)

func (i AppliedConditionalAccessPolicyResult) String() string {
    return []string{"SUCCESS", "FAILURE", "NOTAPPLIED", "NOTENABLED", "UNKNOWN", "UNKNOWNFUTUREVALUE", "REPORTONLYSUCCESS", "REPORTONLYFAILURE", "REPORTONLYNOTAPPLIED", "REPORTONLYINTERRUPTED"}[i]
}
func ParseAppliedConditionalAccessPolicyResult(v string) (interface{}, error) {
    result := SUCCESS_APPLIEDCONDITIONALACCESSPOLICYRESULT
    switch strings.ToUpper(v) {
        case "SUCCESS":
            result = SUCCESS_APPLIEDCONDITIONALACCESSPOLICYRESULT
        case "FAILURE":
            result = FAILURE_APPLIEDCONDITIONALACCESSPOLICYRESULT
        case "NOTAPPLIED":
            result = NOTAPPLIED_APPLIEDCONDITIONALACCESSPOLICYRESULT
        case "NOTENABLED":
            result = NOTENABLED_APPLIEDCONDITIONALACCESSPOLICYRESULT
        case "UNKNOWN":
            result = UNKNOWN_APPLIEDCONDITIONALACCESSPOLICYRESULT
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_APPLIEDCONDITIONALACCESSPOLICYRESULT
        case "REPORTONLYSUCCESS":
            result = REPORTONLYSUCCESS_APPLIEDCONDITIONALACCESSPOLICYRESULT
        case "REPORTONLYFAILURE":
            result = REPORTONLYFAILURE_APPLIEDCONDITIONALACCESSPOLICYRESULT
        case "REPORTONLYNOTAPPLIED":
            result = REPORTONLYNOTAPPLIED_APPLIEDCONDITIONALACCESSPOLICYRESULT
        case "REPORTONLYINTERRUPTED":
            result = REPORTONLYINTERRUPTED_APPLIEDCONDITIONALACCESSPOLICYRESULT
        default:
            return 0, errors.New("Unknown AppliedConditionalAccessPolicyResult value: " + v)
    }
    return &result, nil
}
func SerializeAppliedConditionalAccessPolicyResult(values []AppliedConditionalAccessPolicyResult) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}