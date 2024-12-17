package models
type RoadmapItemDeliveryStage int

const (
    PRIVATEPREVIEW_ROADMAPITEMDELIVERYSTAGE RoadmapItemDeliveryStage = iota
    PUBLICPREVIEW_ROADMAPITEMDELIVERYSTAGE
    GA_ROADMAPITEMDELIVERYSTAGE
    UNKNOWNFUTUREVALUE_ROADMAPITEMDELIVERYSTAGE
)

func (i RoadmapItemDeliveryStage) String() string {
    return []string{"privatePreview", "publicPreview", "ga", "unknownFutureValue"}[i]
}
func ParseRoadmapItemDeliveryStage(v string) (any, error) {
    result := PRIVATEPREVIEW_ROADMAPITEMDELIVERYSTAGE
    switch v {
        case "privatePreview":
            result = PRIVATEPREVIEW_ROADMAPITEMDELIVERYSTAGE
        case "publicPreview":
            result = PUBLICPREVIEW_ROADMAPITEMDELIVERYSTAGE
        case "ga":
            result = GA_ROADMAPITEMDELIVERYSTAGE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ROADMAPITEMDELIVERYSTAGE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRoadmapItemDeliveryStage(values []RoadmapItemDeliveryStage) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RoadmapItemDeliveryStage) isMultiValue() bool {
    return false
}
