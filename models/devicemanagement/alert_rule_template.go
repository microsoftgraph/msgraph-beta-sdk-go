package devicemanagement
type AlertRuleTemplate int

const (
    CLOUDPCPROVISIONSCENARIO_ALERTRULETEMPLATE AlertRuleTemplate = iota
    CLOUDPCIMAGEUPLOADSCENARIO_ALERTRULETEMPLATE
    CLOUDPCONPREMISENETWORKCONNECTIONCHECKSCENARIO_ALERTRULETEMPLATE
    UNKNOWNFUTUREVALUE_ALERTRULETEMPLATE
    CLOUDPCINGRACEPERIODSCENARIO_ALERTRULETEMPLATE
    CLOUDPCFRONTLINEINSUFFICIENTLICENSESSCENARIO_ALERTRULETEMPLATE
    CLOUDPCINACCESSIBLESCENARIO_ALERTRULETEMPLATE
    CLOUDPCFRONTLINECONCURRENCYSCENARIO_ALERTRULETEMPLATE
)

func (i AlertRuleTemplate) String() string {
    return []string{"cloudPcProvisionScenario", "cloudPcImageUploadScenario", "cloudPcOnPremiseNetworkConnectionCheckScenario", "unknownFutureValue", "cloudPcInGracePeriodScenario", "cloudPcFrontlineInsufficientLicensesScenario", "cloudPcInaccessibleScenario", "cloudPcFrontlineConcurrencyScenario"}[i]
}
func ParseAlertRuleTemplate(v string) (any, error) {
    result := CLOUDPCPROVISIONSCENARIO_ALERTRULETEMPLATE
    switch v {
        case "cloudPcProvisionScenario":
            result = CLOUDPCPROVISIONSCENARIO_ALERTRULETEMPLATE
        case "cloudPcImageUploadScenario":
            result = CLOUDPCIMAGEUPLOADSCENARIO_ALERTRULETEMPLATE
        case "cloudPcOnPremiseNetworkConnectionCheckScenario":
            result = CLOUDPCONPREMISENETWORKCONNECTIONCHECKSCENARIO_ALERTRULETEMPLATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ALERTRULETEMPLATE
        case "cloudPcInGracePeriodScenario":
            result = CLOUDPCINGRACEPERIODSCENARIO_ALERTRULETEMPLATE
        case "cloudPcFrontlineInsufficientLicensesScenario":
            result = CLOUDPCFRONTLINEINSUFFICIENTLICENSESSCENARIO_ALERTRULETEMPLATE
        case "cloudPcInaccessibleScenario":
            result = CLOUDPCINACCESSIBLESCENARIO_ALERTRULETEMPLATE
        case "cloudPcFrontlineConcurrencyScenario":
            result = CLOUDPCFRONTLINECONCURRENCYSCENARIO_ALERTRULETEMPLATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAlertRuleTemplate(values []AlertRuleTemplate) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AlertRuleTemplate) isMultiValue() bool {
    return false
}
