package graph
import (
    "strings"
    "errors"
)
// 
type Modality int

const (
    UNKNOWN_MODALITY Modality = iota
    AUDIO_MODALITY
    VIDEO_MODALITY
    VIDEOBASEDSCREENSHARING_MODALITY
    DATA_MODALITY
    UNKNOWNFUTUREVALUE_MODALITY
)

func (i Modality) String() string {
    return []string{"UNKNOWN", "AUDIO", "VIDEO", "VIDEOBASEDSCREENSHARING", "DATA", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseModality(v string) (interface{}, error) {
    result := UNKNOWN_MODALITY
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_MODALITY
        case "AUDIO":
            result = AUDIO_MODALITY
        case "VIDEO":
            result = VIDEO_MODALITY
        case "VIDEOBASEDSCREENSHARING":
            result = VIDEOBASEDSCREENSHARING_MODALITY
        case "DATA":
            result = DATA_MODALITY
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_MODALITY
        default:
            return 0, errors.New("Unknown Modality value: " + v)
    }
    return &result, nil
}
func SerializeModality(values []Modality) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
