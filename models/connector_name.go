package models
// Connectors name for connector status
type ConnectorName int

const (
    // Indicates the expiration date/time for the Apple MDM Push Certificate.
    APPLEPUSHNOTIFICATIONSERVICEEXPIRATIONDATETIME_CONNECTORNAME ConnectorName = iota
    // Indicates the expiration date/time for Vpp Token.
    VPPTOKENEXPIRATIONDATETIME_CONNECTORNAME
    // Indicate the last sync data/time that the Vpp Token performed a sync.
    VPPTOKENLASTSYNCDATETIME_CONNECTORNAME
    // Indicate the last sync date/time that the Windows Autopilot performed a sync.
    WINDOWSAUTOPILOTLASTSYNCDATETIME_CONNECTORNAME
    // Indicates the last sync date/time that the Windows Store for Business performed a sync.
    WINDOWSSTOREFORBUSINESSLASTSYNCDATETIME_CONNECTORNAME
    // Indicates the last sync date/time that the JAMF connector performed a sync.
    JAMFLASTSYNCDATETIME_CONNECTORNAME
    // Indicates the last sync date/time that the NDES connector performed a sync.
    NDESCONNECTORLASTCONNECTIONDATETIME_CONNECTORNAME
    // Indicates the expiration date/time for the Apple Enrollment Program token.
    APPLEDEPEXPIRATIONDATETIME_CONNECTORNAME
    // Indicates the last sync date/time that the Apple Enrollment Program token performed a sync.
    APPLEDEPLASTSYNCDATETIME_CONNECTORNAME
    // Indicates the last sync date/time that the Exchange ActiveSync connector performed a sync.
    ONPREMCONNECTORLASTSYNCDATETIME_CONNECTORNAME
    // Indicates the last sync date/time that the Google Play App performed a sync.
    GOOGLEPLAYAPPLASTSYNCDATETIME_CONNECTORNAME
    // Indicates the last modified date / time that the Google Play connector was updated.
    GOOGLEPLAYCONNECTORLASTMODIFIEDDATETIME_CONNECTORNAME
    // Indicates the last heartbeat date/time that the Windows Defender ATP connector was contacted.
    WINDOWSDEFENDERATPCONNECTORLASTHEARTBEATDATETIME_CONNECTORNAME
    // Indicates the last heartbeat date/time that the Mobile Threat Defence connector was contacted.
    MOBILETHREATDEFENCECONNECTORLASTHEARTBEATDATETIME_CONNECTORNAME
    // Indicates the last sync date/time that the Chrombook Last Directory performed a sync.
    CHROMEBOOKLASTDIRECTORYSYNCDATETIME_CONNECTORNAME
    // Future use
    FUTUREVALUE_CONNECTORNAME
)

func (i ConnectorName) String() string {
    return []string{"applePushNotificationServiceExpirationDateTime", "vppTokenExpirationDateTime", "vppTokenLastSyncDateTime", "windowsAutopilotLastSyncDateTime", "windowsStoreForBusinessLastSyncDateTime", "jamfLastSyncDateTime", "ndesConnectorLastConnectionDateTime", "appleDepExpirationDateTime", "appleDepLastSyncDateTime", "onPremConnectorLastSyncDateTime", "googlePlayAppLastSyncDateTime", "googlePlayConnectorLastModifiedDateTime", "windowsDefenderATPConnectorLastHeartbeatDateTime", "mobileThreatDefenceConnectorLastHeartbeatDateTime", "chromebookLastDirectorySyncDateTime", "futureValue"}[i]
}
func ParseConnectorName(v string) (any, error) {
    result := APPLEPUSHNOTIFICATIONSERVICEEXPIRATIONDATETIME_CONNECTORNAME
    switch v {
        case "applePushNotificationServiceExpirationDateTime":
            result = APPLEPUSHNOTIFICATIONSERVICEEXPIRATIONDATETIME_CONNECTORNAME
        case "vppTokenExpirationDateTime":
            result = VPPTOKENEXPIRATIONDATETIME_CONNECTORNAME
        case "vppTokenLastSyncDateTime":
            result = VPPTOKENLASTSYNCDATETIME_CONNECTORNAME
        case "windowsAutopilotLastSyncDateTime":
            result = WINDOWSAUTOPILOTLASTSYNCDATETIME_CONNECTORNAME
        case "windowsStoreForBusinessLastSyncDateTime":
            result = WINDOWSSTOREFORBUSINESSLASTSYNCDATETIME_CONNECTORNAME
        case "jamfLastSyncDateTime":
            result = JAMFLASTSYNCDATETIME_CONNECTORNAME
        case "ndesConnectorLastConnectionDateTime":
            result = NDESCONNECTORLASTCONNECTIONDATETIME_CONNECTORNAME
        case "appleDepExpirationDateTime":
            result = APPLEDEPEXPIRATIONDATETIME_CONNECTORNAME
        case "appleDepLastSyncDateTime":
            result = APPLEDEPLASTSYNCDATETIME_CONNECTORNAME
        case "onPremConnectorLastSyncDateTime":
            result = ONPREMCONNECTORLASTSYNCDATETIME_CONNECTORNAME
        case "googlePlayAppLastSyncDateTime":
            result = GOOGLEPLAYAPPLASTSYNCDATETIME_CONNECTORNAME
        case "googlePlayConnectorLastModifiedDateTime":
            result = GOOGLEPLAYCONNECTORLASTMODIFIEDDATETIME_CONNECTORNAME
        case "windowsDefenderATPConnectorLastHeartbeatDateTime":
            result = WINDOWSDEFENDERATPCONNECTORLASTHEARTBEATDATETIME_CONNECTORNAME
        case "mobileThreatDefenceConnectorLastHeartbeatDateTime":
            result = MOBILETHREATDEFENCECONNECTORLASTHEARTBEATDATETIME_CONNECTORNAME
        case "chromebookLastDirectorySyncDateTime":
            result = CHROMEBOOKLASTDIRECTORYSYNCDATETIME_CONNECTORNAME
        case "futureValue":
            result = FUTUREVALUE_CONNECTORNAME
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeConnectorName(values []ConnectorName) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ConnectorName) isMultiValue() bool {
    return false
}
