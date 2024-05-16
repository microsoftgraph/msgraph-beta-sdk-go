package security
type HuntingRuleErrorCode int

const (
    QUERYEXECUTIONFAILED_HUNTINGRULEERRORCODE HuntingRuleErrorCode = iota
    QUERYEXECUTIONTHROTTLING_HUNTINGRULEERRORCODE
    QUERYEXCEEDEDRESULTSIZE_HUNTINGRULEERRORCODE
    QUERYLIMITSEXCEEDED_HUNTINGRULEERRORCODE
    QUERYTIMEOUT_HUNTINGRULEERRORCODE
    ALERTCREATIONFAILED_HUNTINGRULEERRORCODE
    ALERTREPORTNOTFOUND_HUNTINGRULEERRORCODE
    PARTIALROWSFAILED_HUNTINGRULEERRORCODE
    UNKNOWNFUTUREVALUE_HUNTINGRULEERRORCODE
    NOIMPACTEDENTITY_HUNTINGRULEERRORCODE
)

func (i HuntingRuleErrorCode) String() string {
    return []string{"queryExecutionFailed", "queryExecutionThrottling", "queryExceededResultSize", "queryLimitsExceeded", "queryTimeout", "alertCreationFailed", "alertReportNotFound", "partialRowsFailed", "unknownFutureValue", "noImpactedEntity"}[i]
}
func ParseHuntingRuleErrorCode(v string) (any, error) {
    result := QUERYEXECUTIONFAILED_HUNTINGRULEERRORCODE
    switch v {
        case "queryExecutionFailed":
            result = QUERYEXECUTIONFAILED_HUNTINGRULEERRORCODE
        case "queryExecutionThrottling":
            result = QUERYEXECUTIONTHROTTLING_HUNTINGRULEERRORCODE
        case "queryExceededResultSize":
            result = QUERYEXCEEDEDRESULTSIZE_HUNTINGRULEERRORCODE
        case "queryLimitsExceeded":
            result = QUERYLIMITSEXCEEDED_HUNTINGRULEERRORCODE
        case "queryTimeout":
            result = QUERYTIMEOUT_HUNTINGRULEERRORCODE
        case "alertCreationFailed":
            result = ALERTCREATIONFAILED_HUNTINGRULEERRORCODE
        case "alertReportNotFound":
            result = ALERTREPORTNOTFOUND_HUNTINGRULEERRORCODE
        case "partialRowsFailed":
            result = PARTIALROWSFAILED_HUNTINGRULEERRORCODE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_HUNTINGRULEERRORCODE
        case "noImpactedEntity":
            result = NOIMPACTEDENTITY_HUNTINGRULEERRORCODE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHuntingRuleErrorCode(values []HuntingRuleErrorCode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HuntingRuleErrorCode) isMultiValue() bool {
    return false
}
