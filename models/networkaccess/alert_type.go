package networkaccess
type AlertType int

const (
    UNHEALTHYREMOTENETWORKS_ALERTTYPE AlertType = iota
    UNHEALTHYCONNECTORS_ALERTTYPE
    DEVICETOKENINCONSISTENCY_ALERTTYPE
    CROSSTENANTANOMALY_ALERTTYPE
    SUSPICIOUSPROCESS_ALERTTYPE
    THREATINTELLIGENCETRANSACTIONS_ALERTTYPE
    UNKNOWNFUTUREVALUE_ALERTTYPE
    WEBCONTENTBLOCKED_ALERTTYPE
)

func (i AlertType) String() string {
    return []string{"unhealthyRemoteNetworks", "unhealthyConnectors", "deviceTokenInconsistency", "crossTenantAnomaly", "suspiciousProcess", "threatIntelligenceTransactions", "unknownFutureValue", "webContentBlocked"}[i]
}
func ParseAlertType(v string) (any, error) {
    result := UNHEALTHYREMOTENETWORKS_ALERTTYPE
    switch v {
        case "unhealthyRemoteNetworks":
            result = UNHEALTHYREMOTENETWORKS_ALERTTYPE
        case "unhealthyConnectors":
            result = UNHEALTHYCONNECTORS_ALERTTYPE
        case "deviceTokenInconsistency":
            result = DEVICETOKENINCONSISTENCY_ALERTTYPE
        case "crossTenantAnomaly":
            result = CROSSTENANTANOMALY_ALERTTYPE
        case "suspiciousProcess":
            result = SUSPICIOUSPROCESS_ALERTTYPE
        case "threatIntelligenceTransactions":
            result = THREATINTELLIGENCETRANSACTIONS_ALERTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ALERTTYPE
        case "webContentBlocked":
            result = WEBCONTENTBLOCKED_ALERTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAlertType(values []AlertType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AlertType) isMultiValue() bool {
    return false
}
