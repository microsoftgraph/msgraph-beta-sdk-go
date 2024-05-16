package models
// A list of possible states for application status on an individual device. When devices contact the Intune service and find targeted application enforcement intent, the status of the enforcement is recorded and becomes accessible in the Graph API. Since the application status is identified during device interaction with the Intune service, status records do not immediately appear upon application group assignment; it is created only after the assignment is evaluated in the service and devices start receiving the policy during check-ins.
type ResultantAppState int

const (
    // The application is not applicable.
    NOTAPPLICABLE_RESULTANTAPPSTATE ResultantAppState = iota
    // The application is installed with no errors.
    INSTALLED_RESULTANTAPPSTATE
    // The application failed to install.
    FAILED_RESULTANTAPPSTATE
    // The application is not installed.
    NOTINSTALLED_RESULTANTAPPSTATE
    // The application failed to uninstall.
    UNINSTALLFAILED_RESULTANTAPPSTATE
    // The installation of the application is in progress.
    PENDINGINSTALL_RESULTANTAPPSTATE
    // The status of the application is unknown.
    UNKNOWN_RESULTANTAPPSTATE
)

func (i ResultantAppState) String() string {
    return []string{"notApplicable", "installed", "failed", "notInstalled", "uninstallFailed", "pendingInstall", "unknown"}[i]
}
func ParseResultantAppState(v string) (any, error) {
    result := NOTAPPLICABLE_RESULTANTAPPSTATE
    switch v {
        case "notApplicable":
            result = NOTAPPLICABLE_RESULTANTAPPSTATE
        case "installed":
            result = INSTALLED_RESULTANTAPPSTATE
        case "failed":
            result = FAILED_RESULTANTAPPSTATE
        case "notInstalled":
            result = NOTINSTALLED_RESULTANTAPPSTATE
        case "uninstallFailed":
            result = UNINSTALLFAILED_RESULTANTAPPSTATE
        case "pendingInstall":
            result = PENDINGINSTALL_RESULTANTAPPSTATE
        case "unknown":
            result = UNKNOWN_RESULTANTAPPSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeResultantAppState(values []ResultantAppState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ResultantAppState) isMultiValue() bool {
    return false
}
