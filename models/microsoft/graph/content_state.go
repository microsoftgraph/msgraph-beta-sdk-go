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
    switch strings.ToUpper(v) {
        case "REST":
            return REST_CONTENTSTATE, nil
        case "MOTION":
            return MOTION_CONTENTSTATE, nil
        case "USE":
            return USE_CONTENTSTATE, nil
    }
    return 0, errors.New("Unknown ContentState value: " + v)
}
func SerializeContentState(values []ContentState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
