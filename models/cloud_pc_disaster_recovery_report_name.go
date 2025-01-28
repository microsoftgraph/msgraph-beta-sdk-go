package models
type CloudPcDisasterRecoveryReportName int

const (
    CROSSREGIONDISASTERRECOVERYREPORT_CLOUDPCDISASTERRECOVERYREPORTNAME CloudPcDisasterRecoveryReportName = iota
    DISASTERRECOVERYREPORT_CLOUDPCDISASTERRECOVERYREPORTNAME
    UNKNOWNFUTUREVALUE_CLOUDPCDISASTERRECOVERYREPORTNAME
)

func (i CloudPcDisasterRecoveryReportName) String() string {
    return []string{"crossRegionDisasterRecoveryReport", "disasterRecoveryReport", "unknownFutureValue"}[i]
}
func ParseCloudPcDisasterRecoveryReportName(v string) (any, error) {
    result := CROSSREGIONDISASTERRECOVERYREPORT_CLOUDPCDISASTERRECOVERYREPORTNAME
    switch v {
        case "crossRegionDisasterRecoveryReport":
            result = CROSSREGIONDISASTERRECOVERYREPORT_CLOUDPCDISASTERRECOVERYREPORTNAME
        case "disasterRecoveryReport":
            result = DISASTERRECOVERYREPORT_CLOUDPCDISASTERRECOVERYREPORTNAME
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCDISASTERRECOVERYREPORTNAME
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCloudPcDisasterRecoveryReportName(values []CloudPcDisasterRecoveryReportName) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudPcDisasterRecoveryReportName) isMultiValue() bool {
    return false
}
