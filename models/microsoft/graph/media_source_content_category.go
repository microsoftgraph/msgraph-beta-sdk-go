package graph
import (
    "strings"
    "errors"
)
// 
type MediaSourceContentCategory int

const (
    MEETING_MEDIASOURCECONTENTCATEGORY MediaSourceContentCategory = iota
    LIVESTREAM_MEDIASOURCECONTENTCATEGORY
    PRESENTATION_MEDIASOURCECONTENTCATEGORY
    SCREENRECORDING_MEDIASOURCECONTENTCATEGORY
    UNKNOWNFUTUREVALUE_MEDIASOURCECONTENTCATEGORY
)

func (i MediaSourceContentCategory) String() string {
    return []string{"MEETING", "LIVESTREAM", "PRESENTATION", "SCREENRECORDING", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseMediaSourceContentCategory(v string) (interface{}, error) {
    result := MEETING_MEDIASOURCECONTENTCATEGORY
    switch strings.ToUpper(v) {
        case "MEETING":
            result = MEETING_MEDIASOURCECONTENTCATEGORY
        case "LIVESTREAM":
            result = LIVESTREAM_MEDIASOURCECONTENTCATEGORY
        case "PRESENTATION":
            result = PRESENTATION_MEDIASOURCECONTENTCATEGORY
        case "SCREENRECORDING":
            result = SCREENRECORDING_MEDIASOURCECONTENTCATEGORY
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_MEDIASOURCECONTENTCATEGORY
        default:
            return 0, errors.New("Unknown MediaSourceContentCategory value: " + v)
    }
    return &result, nil
}
func SerializeMediaSourceContentCategory(values []MediaSourceContentCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
