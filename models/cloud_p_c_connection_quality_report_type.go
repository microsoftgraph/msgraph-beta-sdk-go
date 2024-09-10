package models
type CloudPCConnectionQualityReportType int

const (
    REMOTECONNECTIONQUALITYREPORT_CLOUDPCCONNECTIONQUALITYREPORTTYPE CloudPCConnectionQualityReportType = iota
    REGIONALCONNECTIONQUALITYTRENDREPORT_CLOUDPCCONNECTIONQUALITYREPORTTYPE
    REGIONALCONNECTIONQUALITYINSIGHTSREPORT_CLOUDPCCONNECTIONQUALITYREPORTTYPE
    UNKNOWNFUTUREVALUE_CLOUDPCCONNECTIONQUALITYREPORTTYPE
)

func (i CloudPCConnectionQualityReportType) String() string {
    return []string{"remoteConnectionQualityReport", "regionalConnectionQualityTrendReport", "regionalConnectionQualityInsightsReport", "unknownFutureValue"}[i]
}
func ParseCloudPCConnectionQualityReportType(v string) (any, error) {
    result := REMOTECONNECTIONQUALITYREPORT_CLOUDPCCONNECTIONQUALITYREPORTTYPE
    switch v {
        case "remoteConnectionQualityReport":
            result = REMOTECONNECTIONQUALITYREPORT_CLOUDPCCONNECTIONQUALITYREPORTTYPE
        case "regionalConnectionQualityTrendReport":
            result = REGIONALCONNECTIONQUALITYTRENDREPORT_CLOUDPCCONNECTIONQUALITYREPORTTYPE
        case "regionalConnectionQualityInsightsReport":
            result = REGIONALCONNECTIONQUALITYINSIGHTSREPORT_CLOUDPCCONNECTIONQUALITYREPORTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCCONNECTIONQUALITYREPORTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCloudPCConnectionQualityReportType(values []CloudPCConnectionQualityReportType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudPCConnectionQualityReportType) isMultiValue() bool {
    return false
}
