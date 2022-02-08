package graph
import (
    "strings"
    "errors"
)
// 
type CloudPcDeviceImageStatusDetails int

const (
    INTERNALSERVERERROR_CLOUDPCDEVICEIMAGESTATUSDETAILS CloudPcDeviceImageStatusDetails = iota
    SOURCEIMAGENOTFOUND_CLOUDPCDEVICEIMAGESTATUSDETAILS
    OSVERSIONNOTSUPPORTED_CLOUDPCDEVICEIMAGESTATUSDETAILS
    SOURCEIMAGEINVALID_CLOUDPCDEVICEIMAGESTATUSDETAILS
    SOURCEIMAGENOTGENERALIZED_CLOUDPCDEVICEIMAGESTATUSDETAILS
    UNKNOWNFUTUREVALUE_CLOUDPCDEVICEIMAGESTATUSDETAILS
)

func (i CloudPcDeviceImageStatusDetails) String() string {
    return []string{"INTERNALSERVERERROR", "SOURCEIMAGENOTFOUND", "OSVERSIONNOTSUPPORTED", "SOURCEIMAGEINVALID", "SOURCEIMAGENOTGENERALIZED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseCloudPcDeviceImageStatusDetails(v string) (interface{}, error) {
    result := INTERNALSERVERERROR_CLOUDPCDEVICEIMAGESTATUSDETAILS
    switch strings.ToUpper(v) {
        case "INTERNALSERVERERROR":
            result = INTERNALSERVERERROR_CLOUDPCDEVICEIMAGESTATUSDETAILS
        case "SOURCEIMAGENOTFOUND":
            result = SOURCEIMAGENOTFOUND_CLOUDPCDEVICEIMAGESTATUSDETAILS
        case "OSVERSIONNOTSUPPORTED":
            result = OSVERSIONNOTSUPPORTED_CLOUDPCDEVICEIMAGESTATUSDETAILS
        case "SOURCEIMAGEINVALID":
            result = SOURCEIMAGEINVALID_CLOUDPCDEVICEIMAGESTATUSDETAILS
        case "SOURCEIMAGENOTGENERALIZED":
            result = SOURCEIMAGENOTGENERALIZED_CLOUDPCDEVICEIMAGESTATUSDETAILS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CLOUDPCDEVICEIMAGESTATUSDETAILS
        default:
            return 0, errors.New("Unknown CloudPcDeviceImageStatusDetails value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcDeviceImageStatusDetails(values []CloudPcDeviceImageStatusDetails) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
