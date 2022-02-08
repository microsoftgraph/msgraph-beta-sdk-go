package graph
import (
    "strings"
    "errors"
)
// 
type CloudPcGalleryImageStatus int

const (
    SUPPORTED_CLOUDPCGALLERYIMAGESTATUS CloudPcGalleryImageStatus = iota
    SUPPORTEDWITHWARNING_CLOUDPCGALLERYIMAGESTATUS
    NOTSUPPORTED_CLOUDPCGALLERYIMAGESTATUS
    UNKNOWNFUTUREVALUE_CLOUDPCGALLERYIMAGESTATUS
)

func (i CloudPcGalleryImageStatus) String() string {
    return []string{"SUPPORTED", "SUPPORTEDWITHWARNING", "NOTSUPPORTED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseCloudPcGalleryImageStatus(v string) (interface{}, error) {
    result := SUPPORTED_CLOUDPCGALLERYIMAGESTATUS
    switch strings.ToUpper(v) {
        case "SUPPORTED":
            result = SUPPORTED_CLOUDPCGALLERYIMAGESTATUS
        case "SUPPORTEDWITHWARNING":
            result = SUPPORTEDWITHWARNING_CLOUDPCGALLERYIMAGESTATUS
        case "NOTSUPPORTED":
            result = NOTSUPPORTED_CLOUDPCGALLERYIMAGESTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CLOUDPCGALLERYIMAGESTATUS
        default:
            return 0, errors.New("Unknown CloudPcGalleryImageStatus value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcGalleryImageStatus(values []CloudPcGalleryImageStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
