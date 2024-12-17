package security
type LogDataProvider int

const (
    BARRACUDA_LOGDATAPROVIDER LogDataProvider = iota
    BLUECOAT_LOGDATAPROVIDER
    CHECKPOINT_LOGDATAPROVIDER
    CISCOASA_LOGDATAPROVIDER
    CISCOIRONPORTPROXY_LOGDATAPROVIDER
    FORTIGATE_LOGDATAPROVIDER
    PALOALTO_LOGDATAPROVIDER
    SQUID_LOGDATAPROVIDER
    ZSCALER_LOGDATAPROVIDER
    MCAFEESWG_LOGDATAPROVIDER
    CISCOSCANSAFE_LOGDATAPROVIDER
    JUNIPERSRX_LOGDATAPROVIDER
    SOPHOSSG_LOGDATAPROVIDER
    WEBSENSEV75_LOGDATAPROVIDER
    WEBSENSESIEMCEF_LOGDATAPROVIDER
    MACHINEZONEMERAKI_LOGDATAPROVIDER
    SQUIDNATIVE_LOGDATAPROVIDER
    CISCOFWSM_LOGDATAPROVIDER
    MICROSOFTISAW3C_LOGDATAPROVIDER
    SONICWALL_LOGDATAPROVIDER
    SOPHOSCYBEROAM_LOGDATAPROVIDER
    CLAVISTER_LOGDATAPROVIDER
    CUSTOMPARSER_LOGDATAPROVIDER
    JUNIPERSSG_LOGDATAPROVIDER
    ZSCALERQRADAR_LOGDATAPROVIDER
    JUNIPERSRXSD_LOGDATAPROVIDER
    JUNIPERSRXWELF_LOGDATAPROVIDER
    MICROSOFTCONDITIONALAPPACCESS_LOGDATAPROVIDER
    CISCOASAFIREPOWER_LOGDATAPROVIDER
    GENERICCEF_LOGDATAPROVIDER
    GENERICLEEF_LOGDATAPROVIDER
    GENERICW3C_LOGDATAPROVIDER
    IFILTER_LOGDATAPROVIDER
    CHECKPOINTXML_LOGDATAPROVIDER
    CHECKPOINTSMARTVIEWTRACKER_LOGDATAPROVIDER
    BARRACUDANEXTGENFW_LOGDATAPROVIDER
    BARRACUDANEXTGENFWWEBLOG_LOGDATAPROVIDER
    MICROSOFTDEFENDERFORENDPOINT_LOGDATAPROVIDER
    ZSCALERCEF_LOGDATAPROVIDER
    SOPHOSXG_LOGDATAPROVIDER
    IBOSS_LOGDATAPROVIDER
    FORCEPOINT_LOGDATAPROVIDER
    FORTIOS_LOGDATAPROVIDER
    CISCOIRONPORTWSAII_LOGDATAPROVIDER
    PALOALTOLEEF_LOGDATAPROVIDER
    FORCEPOINTLEEF_LOGDATAPROVIDER
    STORMSHIELD_LOGDATAPROVIDER
    CONTENTKEEPER_LOGDATAPROVIDER
    CISCOIRONPORTWSAIII_LOGDATAPROVIDER
    CHECKPOINTCEF_LOGDATAPROVIDER
    CORRATA_LOGDATAPROVIDER
    CISCOFIREPOWERV6_LOGDATAPROVIDER
    MENLOSECURITYCEF_LOGDATAPROVIDER
    WATCHGUARDXTM_LOGDATAPROVIDER
    OPENSYSTEMSSECUREWEBGATEWAY_LOGDATAPROVIDER
    WANDERA_LOGDATAPROVIDER
    UNKNOWNFUTUREVALUE_LOGDATAPROVIDER
)

func (i LogDataProvider) String() string {
    return []string{"barracuda", "bluecoat", "checkpoint", "ciscoAsa", "ciscoIronportProxy", "fortigate", "paloAlto", "squid", "zscaler", "mcafeeSwg", "ciscoScanSafe", "juniperSrx", "sophosSg", "websenseV75", "websenseSiemCef", "machineZoneMeraki", "squidNative", "ciscoFwsm", "microsoftIsaW3C", "sonicwall", "sophosCyberoam", "clavister", "customParser", "juniperSsg", "zscalerQradar", "juniperSrxSd", "juniperSrxWelf", "microsoftConditionalAppAccess", "ciscoAsaFirepower", "genericCef", "genericLeef", "genericW3C", "iFilter", "checkpointXml", "checkpointSmartViewTracker", "barracudaNextGenFw", "barracudaNextGenFwWeblog", "microsoftDefenderForEndpoint", "zscalerCef", "sophosXg", "iboss", "forcepoint", "fortios", "ciscoIronportWsaIi", "paloAltoLeef", "forcepointLeef", "stormshield", "contentkeeper", "ciscoIronportWsaIii", "checkpointCef", "corrata", "ciscoFirepowerV6", "menloSecurityCef", "watchguardXtm", "openSystemsSecureWebGateway", "wandera", "unknownFutureValue"}[i]
}
func ParseLogDataProvider(v string) (any, error) {
    result := BARRACUDA_LOGDATAPROVIDER
    switch v {
        case "barracuda":
            result = BARRACUDA_LOGDATAPROVIDER
        case "bluecoat":
            result = BLUECOAT_LOGDATAPROVIDER
        case "checkpoint":
            result = CHECKPOINT_LOGDATAPROVIDER
        case "ciscoAsa":
            result = CISCOASA_LOGDATAPROVIDER
        case "ciscoIronportProxy":
            result = CISCOIRONPORTPROXY_LOGDATAPROVIDER
        case "fortigate":
            result = FORTIGATE_LOGDATAPROVIDER
        case "paloAlto":
            result = PALOALTO_LOGDATAPROVIDER
        case "squid":
            result = SQUID_LOGDATAPROVIDER
        case "zscaler":
            result = ZSCALER_LOGDATAPROVIDER
        case "mcafeeSwg":
            result = MCAFEESWG_LOGDATAPROVIDER
        case "ciscoScanSafe":
            result = CISCOSCANSAFE_LOGDATAPROVIDER
        case "juniperSrx":
            result = JUNIPERSRX_LOGDATAPROVIDER
        case "sophosSg":
            result = SOPHOSSG_LOGDATAPROVIDER
        case "websenseV75":
            result = WEBSENSEV75_LOGDATAPROVIDER
        case "websenseSiemCef":
            result = WEBSENSESIEMCEF_LOGDATAPROVIDER
        case "machineZoneMeraki":
            result = MACHINEZONEMERAKI_LOGDATAPROVIDER
        case "squidNative":
            result = SQUIDNATIVE_LOGDATAPROVIDER
        case "ciscoFwsm":
            result = CISCOFWSM_LOGDATAPROVIDER
        case "microsoftIsaW3C":
            result = MICROSOFTISAW3C_LOGDATAPROVIDER
        case "sonicwall":
            result = SONICWALL_LOGDATAPROVIDER
        case "sophosCyberoam":
            result = SOPHOSCYBEROAM_LOGDATAPROVIDER
        case "clavister":
            result = CLAVISTER_LOGDATAPROVIDER
        case "customParser":
            result = CUSTOMPARSER_LOGDATAPROVIDER
        case "juniperSsg":
            result = JUNIPERSSG_LOGDATAPROVIDER
        case "zscalerQradar":
            result = ZSCALERQRADAR_LOGDATAPROVIDER
        case "juniperSrxSd":
            result = JUNIPERSRXSD_LOGDATAPROVIDER
        case "juniperSrxWelf":
            result = JUNIPERSRXWELF_LOGDATAPROVIDER
        case "microsoftConditionalAppAccess":
            result = MICROSOFTCONDITIONALAPPACCESS_LOGDATAPROVIDER
        case "ciscoAsaFirepower":
            result = CISCOASAFIREPOWER_LOGDATAPROVIDER
        case "genericCef":
            result = GENERICCEF_LOGDATAPROVIDER
        case "genericLeef":
            result = GENERICLEEF_LOGDATAPROVIDER
        case "genericW3C":
            result = GENERICW3C_LOGDATAPROVIDER
        case "iFilter":
            result = IFILTER_LOGDATAPROVIDER
        case "checkpointXml":
            result = CHECKPOINTXML_LOGDATAPROVIDER
        case "checkpointSmartViewTracker":
            result = CHECKPOINTSMARTVIEWTRACKER_LOGDATAPROVIDER
        case "barracudaNextGenFw":
            result = BARRACUDANEXTGENFW_LOGDATAPROVIDER
        case "barracudaNextGenFwWeblog":
            result = BARRACUDANEXTGENFWWEBLOG_LOGDATAPROVIDER
        case "microsoftDefenderForEndpoint":
            result = MICROSOFTDEFENDERFORENDPOINT_LOGDATAPROVIDER
        case "zscalerCef":
            result = ZSCALERCEF_LOGDATAPROVIDER
        case "sophosXg":
            result = SOPHOSXG_LOGDATAPROVIDER
        case "iboss":
            result = IBOSS_LOGDATAPROVIDER
        case "forcepoint":
            result = FORCEPOINT_LOGDATAPROVIDER
        case "fortios":
            result = FORTIOS_LOGDATAPROVIDER
        case "ciscoIronportWsaIi":
            result = CISCOIRONPORTWSAII_LOGDATAPROVIDER
        case "paloAltoLeef":
            result = PALOALTOLEEF_LOGDATAPROVIDER
        case "forcepointLeef":
            result = FORCEPOINTLEEF_LOGDATAPROVIDER
        case "stormshield":
            result = STORMSHIELD_LOGDATAPROVIDER
        case "contentkeeper":
            result = CONTENTKEEPER_LOGDATAPROVIDER
        case "ciscoIronportWsaIii":
            result = CISCOIRONPORTWSAIII_LOGDATAPROVIDER
        case "checkpointCef":
            result = CHECKPOINTCEF_LOGDATAPROVIDER
        case "corrata":
            result = CORRATA_LOGDATAPROVIDER
        case "ciscoFirepowerV6":
            result = CISCOFIREPOWERV6_LOGDATAPROVIDER
        case "menloSecurityCef":
            result = MENLOSECURITYCEF_LOGDATAPROVIDER
        case "watchguardXtm":
            result = WATCHGUARDXTM_LOGDATAPROVIDER
        case "openSystemsSecureWebGateway":
            result = OPENSYSTEMSSECUREWEBGATEWAY_LOGDATAPROVIDER
        case "wandera":
            result = WANDERA_LOGDATAPROVIDER
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_LOGDATAPROVIDER
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeLogDataProvider(values []LogDataProvider) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i LogDataProvider) isMultiValue() bool {
    return false
}
