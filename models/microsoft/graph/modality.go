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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_MODALITY, nil
        case "AUDIO":
            return AUDIO_MODALITY, nil
        case "VIDEO":
            return VIDEO_MODALITY, nil
        case "VIDEOBASEDSCREENSHARING":
            return VIDEOBASEDSCREENSHARING_MODALITY, nil
        case "DATA":
            return DATA_MODALITY, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_MODALITY, nil
    }
    return 0, errors.New("Unknown Modality value: " + v)
}
func SerializeModality(values []Modality) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
