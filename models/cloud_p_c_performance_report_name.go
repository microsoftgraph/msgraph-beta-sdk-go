package models
import (
    "errors"
)
// 
type CloudPCPerformanceReportName int

const (
    PERFORMANCETRENDREPORT_CLOUDPCPERFORMANCEREPORTNAME CloudPCPerformanceReportName = iota
    UNKNOWNFUTUREVALUE_CLOUDPCPERFORMANCEREPORTNAME
)

func (i CloudPCPerformanceReportName) String() string {
    return []string{"performanceTrendReport", "unknownFutureValue"}[i]
}
func ParseCloudPCPerformanceReportName(v string) (any, error) {
    result := PERFORMANCETRENDREPORT_CLOUDPCPERFORMANCEREPORTNAME
    switch v {
        case "performanceTrendReport":
            result = PERFORMANCETRENDREPORT_CLOUDPCPERFORMANCEREPORTNAME
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCPERFORMANCEREPORTNAME
        default:
            return 0, errors.New("Unknown CloudPCPerformanceReportName value: " + v)
    }
    return &result, nil
}
func SerializeCloudPCPerformanceReportName(values []CloudPCPerformanceReportName) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudPCPerformanceReportName) isMultiValue() bool {
    return false
}
