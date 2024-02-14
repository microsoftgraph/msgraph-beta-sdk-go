package industrydata
import (
    "errors"
)
type IndustryDataActivityStatus int

const (
    INPROGRESS_INDUSTRYDATAACTIVITYSTATUS IndustryDataActivityStatus = iota
    SKIPPED_INDUSTRYDATAACTIVITYSTATUS
    FAILED_INDUSTRYDATAACTIVITYSTATUS
    COMPLETED_INDUSTRYDATAACTIVITYSTATUS
    COMPLETEDWITHERRORS_INDUSTRYDATAACTIVITYSTATUS
    COMPLETEDWITHWARNINGS_INDUSTRYDATAACTIVITYSTATUS
    UNKNOWNFUTUREVALUE_INDUSTRYDATAACTIVITYSTATUS
)

func (i IndustryDataActivityStatus) String() string {
    return []string{"inProgress", "skipped", "failed", "completed", "completedWithErrors", "completedWithWarnings", "unknownFutureValue"}[i]
}
func ParseIndustryDataActivityStatus(v string) (any, error) {
    result := INPROGRESS_INDUSTRYDATAACTIVITYSTATUS
    switch v {
        case "inProgress":
            result = INPROGRESS_INDUSTRYDATAACTIVITYSTATUS
        case "skipped":
            result = SKIPPED_INDUSTRYDATAACTIVITYSTATUS
        case "failed":
            result = FAILED_INDUSTRYDATAACTIVITYSTATUS
        case "completed":
            result = COMPLETED_INDUSTRYDATAACTIVITYSTATUS
        case "completedWithErrors":
            result = COMPLETEDWITHERRORS_INDUSTRYDATAACTIVITYSTATUS
        case "completedWithWarnings":
            result = COMPLETEDWITHWARNINGS_INDUSTRYDATAACTIVITYSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_INDUSTRYDATAACTIVITYSTATUS
        default:
            return 0, errors.New("Unknown IndustryDataActivityStatus value: " + v)
    }
    return &result, nil
}
func SerializeIndustryDataActivityStatus(values []IndustryDataActivityStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i IndustryDataActivityStatus) isMultiValue() bool {
    return false
}
