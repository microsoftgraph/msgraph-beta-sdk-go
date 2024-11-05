package models
type CloudPCFrontlineReportType int

const (
    NOLICENSEAVAILABLECONNECTIVITYFAILUREREPORT_CLOUDPCFRONTLINEREPORTTYPE CloudPCFrontlineReportType = iota
    LICENSEUSAGEREPORT_CLOUDPCFRONTLINEREPORTTYPE
    LICENSEUSAGEREALTIMEREPORT_CLOUDPCFRONTLINEREPORTTYPE
    LICENSEHOURLYUSAGEREPORT_CLOUDPCFRONTLINEREPORTTYPE
    CONNECTEDUSERREALTIMEREPORT_CLOUDPCFRONTLINEREPORTTYPE
    UNKNOWNFUTUREVALUE_CLOUDPCFRONTLINEREPORTTYPE
)

func (i CloudPCFrontlineReportType) String() string {
    return []string{"noLicenseAvailableConnectivityFailureReport", "licenseUsageReport", "licenseUsageRealTimeReport", "licenseHourlyUsageReport", "connectedUserRealtimeReport", "unknownFutureValue"}[i]
}
func ParseCloudPCFrontlineReportType(v string) (any, error) {
    result := NOLICENSEAVAILABLECONNECTIVITYFAILUREREPORT_CLOUDPCFRONTLINEREPORTTYPE
    switch v {
        case "noLicenseAvailableConnectivityFailureReport":
            result = NOLICENSEAVAILABLECONNECTIVITYFAILUREREPORT_CLOUDPCFRONTLINEREPORTTYPE
        case "licenseUsageReport":
            result = LICENSEUSAGEREPORT_CLOUDPCFRONTLINEREPORTTYPE
        case "licenseUsageRealTimeReport":
            result = LICENSEUSAGEREALTIMEREPORT_CLOUDPCFRONTLINEREPORTTYPE
        case "licenseHourlyUsageReport":
            result = LICENSEHOURLYUSAGEREPORT_CLOUDPCFRONTLINEREPORTTYPE
        case "connectedUserRealtimeReport":
            result = CONNECTEDUSERREALTIMEREPORT_CLOUDPCFRONTLINEREPORTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCFRONTLINEREPORTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCloudPCFrontlineReportType(values []CloudPCFrontlineReportType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudPCFrontlineReportType) isMultiValue() bool {
    return false
}
