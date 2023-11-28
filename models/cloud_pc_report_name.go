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
    FRONTLINELICENSEUSAGEREPORT_CLOUDPCREPORTNAME
    FRONTLINELICENSEUSAGEREALTIMEREPORT_CLOUDPCREPORTNAME
    REMOTECONNECTIONQUALITYREPORTS_CLOUDPCREPORTNAME
    INACCESSIBLECLOUDPCREPORTS_CLOUDPCREPORTNAME
    RAWREMOTECONNECTIONREPORTS_CLOUDPCREPORTNAME
    CLOUDPCUSAGECATEGORYREPORTS_CLOUDPCREPORTNAME
)

func (i CloudPcReportName) String() string {
    return []string{"remoteConnectionHistoricalReports", "dailyAggregatedRemoteConnectionReports", "totalAggregatedRemoteConnectionReports", "sharedUseLicenseUsageReport", "sharedUseLicenseUsageRealTimeReport", "unknownFutureValue", "noLicenseAvailableConnectivityFailureReport", "frontlineLicenseUsageReport", "frontlineLicenseUsageRealTimeReport", "remoteConnectionQualityReports", "inaccessibleCloudPcReports", "rawRemoteConnectionReports", "cloudPcUsageCategoryReports"}[i]
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
        case "frontlineLicenseUsageReport":
            result = FRONTLINELICENSEUSAGEREPORT_CLOUDPCREPORTNAME
        case "frontlineLicenseUsageRealTimeReport":
            result = FRONTLINELICENSEUSAGEREALTIMEREPORT_CLOUDPCREPORTNAME
        case "remoteConnectionQualityReports":
            result = REMOTECONNECTIONQUALITYREPORTS_CLOUDPCREPORTNAME
        case "inaccessibleCloudPcReports":
            result = INACCESSIBLECLOUDPCREPORTS_CLOUDPCREPORTNAME
        case "rawRemoteConnectionReports":
            result = RAWREMOTECONNECTIONREPORTS_CLOUDPCREPORTNAME
        case "cloudPcUsageCategoryReports":
            result = CLOUDPCUSAGECATEGORYREPORTS_CLOUDPCREPORTNAME
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
func (i CloudPcReportName) isMultiValue() bool {
    return false
}
