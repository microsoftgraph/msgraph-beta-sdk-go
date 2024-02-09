package models
import (
    "errors"
)
type CloudPCInaccessibleReportName int

const (
    INACCESSIBLECLOUDPCREPORTS_CLOUDPCINACCESSIBLEREPORTNAME CloudPCInaccessibleReportName = iota
    INACCESSIBLECLOUDPCTRENDREPORT_CLOUDPCINACCESSIBLEREPORTNAME
    UNKNOWNFUTUREVALUE_CLOUDPCINACCESSIBLEREPORTNAME
)

func (i CloudPCInaccessibleReportName) String() string {
    return []string{"inaccessibleCloudPcReports", "inaccessibleCloudPcTrendReport", "unknownFutureValue"}[i]
}
func ParseCloudPCInaccessibleReportName(v string) (any, error) {
    result := INACCESSIBLECLOUDPCREPORTS_CLOUDPCINACCESSIBLEREPORTNAME
    switch v {
        case "inaccessibleCloudPcReports":
            result = INACCESSIBLECLOUDPCREPORTS_CLOUDPCINACCESSIBLEREPORTNAME
        case "inaccessibleCloudPcTrendReport":
            result = INACCESSIBLECLOUDPCTRENDREPORT_CLOUDPCINACCESSIBLEREPORTNAME
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCINACCESSIBLEREPORTNAME
        default:
            return 0, errors.New("Unknown CloudPCInaccessibleReportName value: " + v)
    }
    return &result, nil
}
func SerializeCloudPCInaccessibleReportName(values []CloudPCInaccessibleReportName) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudPCInaccessibleReportName) isMultiValue() bool {
    return false
}
