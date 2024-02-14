package models
import (
    "errors"
)
type CloudPcDeviceImageStatusDetails int

const (
    INTERNALSERVERERROR_CLOUDPCDEVICEIMAGESTATUSDETAILS CloudPcDeviceImageStatusDetails = iota
    SOURCEIMAGENOTFOUND_CLOUDPCDEVICEIMAGESTATUSDETAILS
    OSVERSIONNOTSUPPORTED_CLOUDPCDEVICEIMAGESTATUSDETAILS
    SOURCEIMAGEINVALID_CLOUDPCDEVICEIMAGESTATUSDETAILS
    SOURCEIMAGENOTGENERALIZED_CLOUDPCDEVICEIMAGESTATUSDETAILS
    UNKNOWNFUTUREVALUE_CLOUDPCDEVICEIMAGESTATUSDETAILS
    VMALREADYAZUREADJOINED_CLOUDPCDEVICEIMAGESTATUSDETAILS
    PAIDSOURCEIMAGENOTSUPPORT_CLOUDPCDEVICEIMAGESTATUSDETAILS
    SOURCEIMAGENOTSUPPORTCUSTOMIZEVMNAME_CLOUDPCDEVICEIMAGESTATUSDETAILS
    SOURCEIMAGESIZEEXCEEDSLIMITATION_CLOUDPCDEVICEIMAGESTATUSDETAILS
)

func (i CloudPcDeviceImageStatusDetails) String() string {
    return []string{"internalServerError", "sourceImageNotFound", "osVersionNotSupported", "sourceImageInvalid", "sourceImageNotGeneralized", "unknownFutureValue", "vmAlreadyAzureAdjoined", "paidSourceImageNotSupport", "sourceImageNotSupportCustomizeVMName", "sourceImageSizeExceedsLimitation"}[i]
}
func ParseCloudPcDeviceImageStatusDetails(v string) (any, error) {
    result := INTERNALSERVERERROR_CLOUDPCDEVICEIMAGESTATUSDETAILS
    switch v {
        case "internalServerError":
            result = INTERNALSERVERERROR_CLOUDPCDEVICEIMAGESTATUSDETAILS
        case "sourceImageNotFound":
            result = SOURCEIMAGENOTFOUND_CLOUDPCDEVICEIMAGESTATUSDETAILS
        case "osVersionNotSupported":
            result = OSVERSIONNOTSUPPORTED_CLOUDPCDEVICEIMAGESTATUSDETAILS
        case "sourceImageInvalid":
            result = SOURCEIMAGEINVALID_CLOUDPCDEVICEIMAGESTATUSDETAILS
        case "sourceImageNotGeneralized":
            result = SOURCEIMAGENOTGENERALIZED_CLOUDPCDEVICEIMAGESTATUSDETAILS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCDEVICEIMAGESTATUSDETAILS
        case "vmAlreadyAzureAdjoined":
            result = VMALREADYAZUREADJOINED_CLOUDPCDEVICEIMAGESTATUSDETAILS
        case "paidSourceImageNotSupport":
            result = PAIDSOURCEIMAGENOTSUPPORT_CLOUDPCDEVICEIMAGESTATUSDETAILS
        case "sourceImageNotSupportCustomizeVMName":
            result = SOURCEIMAGENOTSUPPORTCUSTOMIZEVMNAME_CLOUDPCDEVICEIMAGESTATUSDETAILS
        case "sourceImageSizeExceedsLimitation":
            result = SOURCEIMAGESIZEEXCEEDSLIMITATION_CLOUDPCDEVICEIMAGESTATUSDETAILS
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
func (i CloudPcDeviceImageStatusDetails) isMultiValue() bool {
    return false
}
