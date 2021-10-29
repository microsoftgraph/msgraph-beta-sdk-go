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
    switch strings.ToUpper(v) {
        case "SUPPORTED":
            return SUPPORTED_CLOUDPCGALLERYIMAGESTATUS, nil
        case "SUPPORTEDWITHWARNING":
            return SUPPORTEDWITHWARNING_CLOUDPCGALLERYIMAGESTATUS, nil
        case "NOTSUPPORTED":
            return NOTSUPPORTED_CLOUDPCGALLERYIMAGESTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CLOUDPCGALLERYIMAGESTATUS, nil
    }
    return 0, errors.New("Unknown CloudPcGalleryImageStatus value: " + v)
}
func SerializeCloudPcGalleryImageStatus(values []CloudPcGalleryImageStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
