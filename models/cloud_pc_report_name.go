package models
type CloudPcReportName int

const (
    REMOTECONNECTIONHISTORICALREPORTS_CLOUDPCREPORTNAME CloudPcReportName = iota
    DAILYAGGREGATEDREMOTECONNECTIONREPORTS_CLOUDPCREPORTNAME
    TOTALAGGREGATEDREMOTECONNECTIONREPORTS_CLOUDPCREPORTNAME
    UNKNOWNFUTUREVALUE_CLOUDPCREPORTNAME
    NOLICENSEAVAILABLECONNECTIVITYFAILUREREPORT_CLOUDPCREPORTNAME
    FRONTLINELICENSEUSAGEREPORT_CLOUDPCREPORTNAME
    FRONTLINELICENSEUSAGEREALTIMEREPORT_CLOUDPCREPORTNAME
    REMOTECONNECTIONQUALITYREPORTS_CLOUDPCREPORTNAME
    INACCESSIBLECLOUDPCREPORTS_CLOUDPCREPORTNAME
    ACTIONSTATUSREPORT_CLOUDPCREPORTNAME
    RAWREMOTECONNECTIONREPORTS_CLOUDPCREPORTNAME
    CLOUDPCUSAGECATEGORYREPORTS_CLOUDPCREPORTNAME
    CROSSREGIONDISASTERRECOVERYREPORT_CLOUDPCREPORTNAME
    PERFORMANCETRENDREPORT_CLOUDPCREPORTNAME
    INACCESSIBLECLOUDPCTRENDREPORT_CLOUDPCREPORTNAME
    REGIONALCONNECTIONQUALITYTRENDREPORT_CLOUDPCREPORTNAME
    REGIONALCONNECTIONQUALITYINSIGHTSREPORT_CLOUDPCREPORTNAME
    REMOTECONNECTIONQUALITYREPORT_CLOUDPCREPORTNAME
    FRONTLINELICENSEHOURLYUSAGEREPORT_CLOUDPCREPORTNAME
    FRONTLINEREALTIMEUSERCONNECTIONSREPORT_CLOUDPCREPORTNAME
    BULKACTIONSTATUSREPORT_CLOUDPCREPORTNAME
    TROUBLESHOOTDETAILSREPORT_CLOUDPCREPORTNAME
    TROUBLESHOOTTRENDCOUNTREPORT_CLOUDPCREPORTNAME
    TROUBLESHOOTREGIONALREPORT_CLOUDPCREPORTNAME
    TROUBLESHOOTISSUECOUNTREPORT_CLOUDPCREPORTNAME
)

func (i CloudPcReportName) String() string {
    return []string{"remoteConnectionHistoricalReports", "dailyAggregatedRemoteConnectionReports", "totalAggregatedRemoteConnectionReports", "unknownFutureValue", "noLicenseAvailableConnectivityFailureReport", "frontlineLicenseUsageReport", "frontlineLicenseUsageRealTimeReport", "remoteConnectionQualityReports", "inaccessibleCloudPcReports", "actionStatusReport", "rawRemoteConnectionReports", "cloudPcUsageCategoryReports", "crossRegionDisasterRecoveryReport", "performanceTrendReport", "inaccessibleCloudPcTrendReport", "regionalConnectionQualityTrendReport", "regionalConnectionQualityInsightsReport", "remoteConnectionQualityReport", "frontlineLicenseHourlyUsageReport", "frontlineRealtimeUserConnectionsReport", "bulkActionStatusReport", "troubleshootDetailsReport", "troubleshootTrendCountReport", "troubleshootRegionalReport", "troubleshootIssueCountReport"}[i]
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
        case "actionStatusReport":
            result = ACTIONSTATUSREPORT_CLOUDPCREPORTNAME
        case "rawRemoteConnectionReports":
            result = RAWREMOTECONNECTIONREPORTS_CLOUDPCREPORTNAME
        case "cloudPcUsageCategoryReports":
            result = CLOUDPCUSAGECATEGORYREPORTS_CLOUDPCREPORTNAME
        case "crossRegionDisasterRecoveryReport":
            result = CROSSREGIONDISASTERRECOVERYREPORT_CLOUDPCREPORTNAME
        case "performanceTrendReport":
            result = PERFORMANCETRENDREPORT_CLOUDPCREPORTNAME
        case "inaccessibleCloudPcTrendReport":
            result = INACCESSIBLECLOUDPCTRENDREPORT_CLOUDPCREPORTNAME
        case "regionalConnectionQualityTrendReport":
            result = REGIONALCONNECTIONQUALITYTRENDREPORT_CLOUDPCREPORTNAME
        case "regionalConnectionQualityInsightsReport":
            result = REGIONALCONNECTIONQUALITYINSIGHTSREPORT_CLOUDPCREPORTNAME
        case "remoteConnectionQualityReport":
            result = REMOTECONNECTIONQUALITYREPORT_CLOUDPCREPORTNAME
        case "frontlineLicenseHourlyUsageReport":
            result = FRONTLINELICENSEHOURLYUSAGEREPORT_CLOUDPCREPORTNAME
        case "frontlineRealtimeUserConnectionsReport":
            result = FRONTLINEREALTIMEUSERCONNECTIONSREPORT_CLOUDPCREPORTNAME
        case "bulkActionStatusReport":
            result = BULKACTIONSTATUSREPORT_CLOUDPCREPORTNAME
        case "troubleshootDetailsReport":
            result = TROUBLESHOOTDETAILSREPORT_CLOUDPCREPORTNAME
        case "troubleshootTrendCountReport":
            result = TROUBLESHOOTTRENDCOUNTREPORT_CLOUDPCREPORTNAME
        case "troubleshootRegionalReport":
            result = TROUBLESHOOTREGIONALREPORT_CLOUDPCREPORTNAME
        case "troubleshootIssueCountReport":
            result = TROUBLESHOOTISSUECOUNTREPORT_CLOUDPCREPORTNAME
        default:
            return nil, nil
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
