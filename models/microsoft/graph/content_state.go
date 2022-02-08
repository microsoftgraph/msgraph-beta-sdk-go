package graph
import (
    "strings"
    "errors"
)
// 
type ContentState int

const (
    REST_CONTENTSTATE ContentState = iota
    MOTION_CONTENTSTATE
    USE_CONTENTSTATE
)

func (i ContentState) String() string {
    return []string{"REST", "MOTION", "USE"}[i]
}
func ParseContentState(v string) (interface{}, error) {
    result := REST_CONTENTSTATE
    switch strings.ToUpper(v) {
        case "REST":
            result = REST_CONTENTSTATE
        case "MOTION":
            result = MOTION_CONTENTSTATE
        case "USE":
            result = USE_CONTENTSTATE
        default:
            return 0, errors.New("Unknown ContentState value: " + v)
    }
    return &result, nil
}
func SerializeContentState(values []ContentState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
