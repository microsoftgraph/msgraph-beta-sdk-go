package models
import (
    "errors"
)
type CloudPcResizeValidationCode int

const (
    SUCCESS_CLOUDPCRESIZEVALIDATIONCODE CloudPcResizeValidationCode = iota
    CLOUDPCNOTFOUND_CLOUDPCRESIZEVALIDATIONCODE
    OPERATIONCONFLICT_CLOUDPCRESIZEVALIDATIONCODE
    OPERATIONNOTSUPPORTED_CLOUDPCRESIZEVALIDATIONCODE
    TARGETLICENSEHASASSIGNED_CLOUDPCRESIZEVALIDATIONCODE
    INTERNALSERVERERROR_CLOUDPCRESIZEVALIDATIONCODE
    UNKNOWNFUTUREVALUE_CLOUDPCRESIZEVALIDATIONCODE
)

func (i CloudPcResizeValidationCode) String() string {
    return []string{"success", "cloudPcNotFound", "operationConflict", "operationNotSupported", "targetLicenseHasAssigned", "internalServerError", "unknownFutureValue"}[i]
}
func ParseCloudPcResizeValidationCode(v string) (any, error) {
    result := SUCCESS_CLOUDPCRESIZEVALIDATIONCODE
    switch v {
        case "success":
            result = SUCCESS_CLOUDPCRESIZEVALIDATIONCODE
        case "cloudPcNotFound":
            result = CLOUDPCNOTFOUND_CLOUDPCRESIZEVALIDATIONCODE
        case "operationConflict":
            result = OPERATIONCONFLICT_CLOUDPCRESIZEVALIDATIONCODE
        case "operationNotSupported":
            result = OPERATIONNOTSUPPORTED_CLOUDPCRESIZEVALIDATIONCODE
        case "targetLicenseHasAssigned":
            result = TARGETLICENSEHASASSIGNED_CLOUDPCRESIZEVALIDATIONCODE
        case "internalServerError":
            result = INTERNALSERVERERROR_CLOUDPCRESIZEVALIDATIONCODE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCRESIZEVALIDATIONCODE
        default:
            return 0, errors.New("Unknown CloudPcResizeValidationCode value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcResizeValidationCode(values []CloudPcResizeValidationCode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudPcResizeValidationCode) isMultiValue() bool {
    return false
}
