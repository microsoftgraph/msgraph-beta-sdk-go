package security
type AppCategory int

const (
    SECURITY_APPCATEGORY AppCategory = iota
    COLLABORATION_APPCATEGORY
    HOSTINGSERVICES_APPCATEGORY
    ONLINEMEETINGS_APPCATEGORY
    NEWSANDENTERTAINMENT_APPCATEGORY
    ECOMMERCE_APPCATEGORY
    EDUCATION_APPCATEGORY
    CLOUDSTORAGE_APPCATEGORY
    MARKETING_APPCATEGORY
    OPERATIONSMANAGEMENT_APPCATEGORY
    HEALTH_APPCATEGORY
    ADVERTISING_APPCATEGORY
    PRODUCTIVITY_APPCATEGORY
    ACCOUNTINGANDFINANCE_APPCATEGORY
    CONTENTMANAGEMENT_APPCATEGORY
    CONTENTSHARING_APPCATEGORY
    BUSINESSMANAGEMENT_APPCATEGORY
    COMMUNICATIONS_APPCATEGORY
    DATAANALYTICS_APPCATEGORY
    BUSINESSINTELLIGENCE_APPCATEGORY
    WEBEMAIL_APPCATEGORY
    CODEHOSTING_APPCATEGORY
    WEBANALYTICS_APPCATEGORY
    SOCIALNETWORK_APPCATEGORY
    CRM_APPCATEGORY
    FORUMS_APPCATEGORY
    HUMANRESOURCEMANAGEMENT_APPCATEGORY
    TRANSPORTATIONANDTRAVEL_APPCATEGORY
    PRODUCTDESIGN_APPCATEGORY
    SALES_APPCATEGORY
    CLOUDCOMPUTINGPLATFORM_APPCATEGORY
    PROJECTMANAGEMENT_APPCATEGORY
    PERSONALINSTANTMESSAGING_APPCATEGORY
    DEVELOPMENTTOOLS_APPCATEGORY
    ITSERVICES_APPCATEGORY
    SUPPLYCHAINANDLOGISTICS_APPCATEGORY
    PROPERTYMANAGEMENT_APPCATEGORY
    CUSTOMERSUPPORT_APPCATEGORY
    INTERNETOFTHINGS_APPCATEGORY
    VENDORMANAGEMENTSYSTEMS_APPCATEGORY
    WEBSITEMONITORING_APPCATEGORY
    GENERATIVEAI_APPCATEGORY
    UNKNOWN_APPCATEGORY
    UNKNOWNFUTUREVALUE_APPCATEGORY
)

func (i AppCategory) String() string {
    return []string{"security", "collaboration", "hostingServices", "onlineMeetings", "newsAndEntertainment", "eCommerce", "education", "cloudStorage", "marketing", "operationsManagement", "health", "advertising", "productivity", "accountingAndFinance", "contentManagement", "contentSharing", "businessManagement", "communications", "dataAnalytics", "businessIntelligence", "webemail", "codeHosting", "webAnalytics", "socialNetwork", "crm", "forums", "humanResourceManagement", "transportationAndTravel", "productDesign", "sales", "cloudComputingPlatform", "projectManagement", "personalInstantMessaging", "developmentTools", "itServices", "supplyChainAndLogistics", "propertyManagement", "customerSupport", "internetOfThings", "vendorManagementSystems", "websiteMonitoring", "generativeAi", "unknown", "unknownFutureValue"}[i]
}
func ParseAppCategory(v string) (any, error) {
    result := SECURITY_APPCATEGORY
    switch v {
        case "security":
            result = SECURITY_APPCATEGORY
        case "collaboration":
            result = COLLABORATION_APPCATEGORY
        case "hostingServices":
            result = HOSTINGSERVICES_APPCATEGORY
        case "onlineMeetings":
            result = ONLINEMEETINGS_APPCATEGORY
        case "newsAndEntertainment":
            result = NEWSANDENTERTAINMENT_APPCATEGORY
        case "eCommerce":
            result = ECOMMERCE_APPCATEGORY
        case "education":
            result = EDUCATION_APPCATEGORY
        case "cloudStorage":
            result = CLOUDSTORAGE_APPCATEGORY
        case "marketing":
            result = MARKETING_APPCATEGORY
        case "operationsManagement":
            result = OPERATIONSMANAGEMENT_APPCATEGORY
        case "health":
            result = HEALTH_APPCATEGORY
        case "advertising":
            result = ADVERTISING_APPCATEGORY
        case "productivity":
            result = PRODUCTIVITY_APPCATEGORY
        case "accountingAndFinance":
            result = ACCOUNTINGANDFINANCE_APPCATEGORY
        case "contentManagement":
            result = CONTENTMANAGEMENT_APPCATEGORY
        case "contentSharing":
            result = CONTENTSHARING_APPCATEGORY
        case "businessManagement":
            result = BUSINESSMANAGEMENT_APPCATEGORY
        case "communications":
            result = COMMUNICATIONS_APPCATEGORY
        case "dataAnalytics":
            result = DATAANALYTICS_APPCATEGORY
        case "businessIntelligence":
            result = BUSINESSINTELLIGENCE_APPCATEGORY
        case "webemail":
            result = WEBEMAIL_APPCATEGORY
        case "codeHosting":
            result = CODEHOSTING_APPCATEGORY
        case "webAnalytics":
            result = WEBANALYTICS_APPCATEGORY
        case "socialNetwork":
            result = SOCIALNETWORK_APPCATEGORY
        case "crm":
            result = CRM_APPCATEGORY
        case "forums":
            result = FORUMS_APPCATEGORY
        case "humanResourceManagement":
            result = HUMANRESOURCEMANAGEMENT_APPCATEGORY
        case "transportationAndTravel":
            result = TRANSPORTATIONANDTRAVEL_APPCATEGORY
        case "productDesign":
            result = PRODUCTDESIGN_APPCATEGORY
        case "sales":
            result = SALES_APPCATEGORY
        case "cloudComputingPlatform":
            result = CLOUDCOMPUTINGPLATFORM_APPCATEGORY
        case "projectManagement":
            result = PROJECTMANAGEMENT_APPCATEGORY
        case "personalInstantMessaging":
            result = PERSONALINSTANTMESSAGING_APPCATEGORY
        case "developmentTools":
            result = DEVELOPMENTTOOLS_APPCATEGORY
        case "itServices":
            result = ITSERVICES_APPCATEGORY
        case "supplyChainAndLogistics":
            result = SUPPLYCHAINANDLOGISTICS_APPCATEGORY
        case "propertyManagement":
            result = PROPERTYMANAGEMENT_APPCATEGORY
        case "customerSupport":
            result = CUSTOMERSUPPORT_APPCATEGORY
        case "internetOfThings":
            result = INTERNETOFTHINGS_APPCATEGORY
        case "vendorManagementSystems":
            result = VENDORMANAGEMENTSYSTEMS_APPCATEGORY
        case "websiteMonitoring":
            result = WEBSITEMONITORING_APPCATEGORY
        case "generativeAi":
            result = GENERATIVEAI_APPCATEGORY
        case "unknown":
            result = UNKNOWN_APPCATEGORY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_APPCATEGORY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAppCategory(values []AppCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppCategory) isMultiValue() bool {
    return false
}
