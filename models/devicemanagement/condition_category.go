package devicemanagement
type ConditionCategory int

const (
    PROVISIONFAILURES_CONDITIONCATEGORY ConditionCategory = iota
    IMAGEUPLOADFAILURES_CONDITIONCATEGORY
    AZURENETWORKCONNECTIONCHECKFAILURES_CONDITIONCATEGORY
    CLOUDPCINGRACEPERIOD_CONDITIONCATEGORY
    FRONTLINEINSUFFICIENTLICENSES_CONDITIONCATEGORY
    CLOUDPCCONNECTIONERRORS_CONDITIONCATEGORY
    CLOUDPCHOSTHEALTHCHECKFAILURES_CONDITIONCATEGORY
    CLOUDPCZONEOUTAGE_CONDITIONCATEGORY
    UNKNOWNFUTUREVALUE_CONDITIONCATEGORY
    FRONTLINEBUFFERUSAGEDURATION_CONDITIONCATEGORY
    FRONTLINEBUFFERUSAGETHRESHOLD_CONDITIONCATEGORY
)

func (i ConditionCategory) String() string {
    return []string{"provisionFailures", "imageUploadFailures", "azureNetworkConnectionCheckFailures", "cloudPcInGracePeriod", "frontlineInsufficientLicenses", "cloudPcConnectionErrors", "cloudPcHostHealthCheckFailures", "cloudPcZoneOutage", "unknownFutureValue", "frontlineBufferUsageDuration", "frontlineBufferUsageThreshold"}[i]
}
func ParseConditionCategory(v string) (any, error) {
    result := PROVISIONFAILURES_CONDITIONCATEGORY
    switch v {
        case "provisionFailures":
            result = PROVISIONFAILURES_CONDITIONCATEGORY
        case "imageUploadFailures":
            result = IMAGEUPLOADFAILURES_CONDITIONCATEGORY
        case "azureNetworkConnectionCheckFailures":
            result = AZURENETWORKCONNECTIONCHECKFAILURES_CONDITIONCATEGORY
        case "cloudPcInGracePeriod":
            result = CLOUDPCINGRACEPERIOD_CONDITIONCATEGORY
        case "frontlineInsufficientLicenses":
            result = FRONTLINEINSUFFICIENTLICENSES_CONDITIONCATEGORY
        case "cloudPcConnectionErrors":
            result = CLOUDPCCONNECTIONERRORS_CONDITIONCATEGORY
        case "cloudPcHostHealthCheckFailures":
            result = CLOUDPCHOSTHEALTHCHECKFAILURES_CONDITIONCATEGORY
        case "cloudPcZoneOutage":
            result = CLOUDPCZONEOUTAGE_CONDITIONCATEGORY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CONDITIONCATEGORY
        case "frontlineBufferUsageDuration":
            result = FRONTLINEBUFFERUSAGEDURATION_CONDITIONCATEGORY
        case "frontlineBufferUsageThreshold":
            result = FRONTLINEBUFFERUSAGETHRESHOLD_CONDITIONCATEGORY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeConditionCategory(values []ConditionCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ConditionCategory) isMultiValue() bool {
    return false
}
