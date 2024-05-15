package industrydata
type IndustryDataRunStatus int

const (
    RUNNING_INDUSTRYDATARUNSTATUS IndustryDataRunStatus = iota
    FAILED_INDUSTRYDATARUNSTATUS
    COMPLETED_INDUSTRYDATARUNSTATUS
    COMPLETEDWITHERRORS_INDUSTRYDATARUNSTATUS
    COMPLETEDWITHWARNINGS_INDUSTRYDATARUNSTATUS
    UNKNOWNFUTUREVALUE_INDUSTRYDATARUNSTATUS
)

func (i IndustryDataRunStatus) String() string {
    return []string{"running", "failed", "completed", "completedWithErrors", "completedWithWarnings", "unknownFutureValue"}[i]
}
func ParseIndustryDataRunStatus(v string) (any, error) {
    result := RUNNING_INDUSTRYDATARUNSTATUS
    switch v {
        case "running":
            result = RUNNING_INDUSTRYDATARUNSTATUS
        case "failed":
            result = FAILED_INDUSTRYDATARUNSTATUS
        case "completed":
            result = COMPLETED_INDUSTRYDATARUNSTATUS
        case "completedWithErrors":
            result = COMPLETEDWITHERRORS_INDUSTRYDATARUNSTATUS
        case "completedWithWarnings":
            result = COMPLETEDWITHWARNINGS_INDUSTRYDATARUNSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_INDUSTRYDATARUNSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeIndustryDataRunStatus(values []IndustryDataRunStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i IndustryDataRunStatus) isMultiValue() bool {
    return false
}
