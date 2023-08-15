package models
import (
    "errors"
)
// 
type CloudPcReportName int

const (
    REMOTECONNECTIONHISTORICALREPORTS_CLOUDPCREPORTNAME CloudPcReportName = iota
    DAILYAGGREGATEDREMOTECONNECTIONREPORTS_CLOUDPCREPORTNAME
    TOTALAGGREGATEDREMOTECONNECTIONREPORTS_CLOUDPCREPORTNAME
    SHAREDUSELICENSEUSAGEREPORT_CLOUDPCREPORTNAME
    SHAREDUSELICENSEUSAGEREALTIMEREPORT_CLOUDPCREPORTNAME
    UNKNOWNFUTUREVALUE_CLOUDPCREPORTNAME
    NOLICENSEAVAILABLECONNECTIVITYFAILUREREPORT_CLOUDPCREPORTNAME
    REMOTECONNECTIONQUALITYREPORTS_CLOUDPCREPORTNAME
    INACCESSIBLECLOUDPCREPORTS_CLOUDPCREPORTNAME
)

func (i CloudPcReportName) String() string {
    return []string{"remoteConnectionHistoricalReports", "dailyAggregatedRemoteConnectionReports", "totalAggregatedRemoteConnectionReports", "sharedUseLicenseUsageReport", "sharedUseLicenseUsageRealTimeReport", "unknownFutureValue", "noLicenseAvailableConnectivityFailureReport", "remoteConnectionQualityReports", "inaccessibleCloudPcReports"}[i]
}
func ParseCloudPcReportName(v string) (any, error) {
    result := REMOTECONNECTIONHISTORICALREPORTS_CLOUDPCREPORTNAME
    switch v {
        case "remoteConnectionHistoricalReports":
            result = REMOTECONNECTIONHISTORICALREPORTS_CLOUDPCREPORTNAME
        case "dailyAggregatedRemoteConnectionReports":
            result = DAILYAGGREGATEDREMOTECONNECTIONREPORTS_CLOUDPCREPORTNAME
        case "totalAggregatedRemoteConnectionReports":
            result = TOTALAGGREGATEDREMOTECONNECTIONREPORTS_CLOUDPCREPORTNAME
        case "sharedUseLicenseUsageReport":
            result = SHAREDUSELICENSEUSAGEREPORT_CLOUDPCREPORTNAME
        case "sharedUseLicenseUsageRealTimeReport":
            result = SHAREDUSELICENSEUSAGEREALTIMEREPORT_CLOUDPCREPORTNAME
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCREPORTNAME
        case "noLicenseAvailableConnectivityFailureReport":
            result = NOLICENSEAVAILABLECONNECTIVITYFAILUREREPORT_CLOUDPCREPORTNAME
        case "remoteConnectionQualityReports":
            result = REMOTECONNECTIONQUALITYREPORTS_CLOUDPCREPORTNAME
        case "inaccessibleCloudPcReports":
            result = INACCESSIBLECLOUDPCREPORTS_CLOUDPCREPORTNAME
        default:
            return 0, errors.New("Unknown CloudPcReportName value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcReportName(values []CloudPcReportName) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
