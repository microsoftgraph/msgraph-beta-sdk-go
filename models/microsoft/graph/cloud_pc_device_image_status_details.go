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
    switch strings.ToUpper(v) {
        case "INTERNALSERVERERROR":
            return INTERNALSERVERERROR_CLOUDPCDEVICEIMAGESTATUSDETAILS, nil
        case "SOURCEIMAGENOTFOUND":
            return SOURCEIMAGENOTFOUND_CLOUDPCDEVICEIMAGESTATUSDETAILS, nil
        case "OSVERSIONNOTSUPPORTED":
            return OSVERSIONNOTSUPPORTED_CLOUDPCDEVICEIMAGESTATUSDETAILS, nil
        case "SOURCEIMAGEINVALID":
            return SOURCEIMAGEINVALID_CLOUDPCDEVICEIMAGESTATUSDETAILS, nil
        case "SOURCEIMAGENOTGENERALIZED":
            return SOURCEIMAGENOTGENERALIZED_CLOUDPCDEVICEIMAGESTATUSDETAILS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CLOUDPCDEVICEIMAGESTATUSDETAILS, nil
    }
    return 0, errors.New("Unknown CloudPcDeviceImageStatusDetails value: " + v)
}
func SerializeCloudPcDeviceImageStatusDetails(values []CloudPcDeviceImageStatusDetails) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
