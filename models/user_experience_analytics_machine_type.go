package models
// Indicates if machine is physical or virtual. Possible values are: physical or virtual
type UserExperienceAnalyticsMachineType int

const (
    // Indicates that the type is unknown.
    UNKNOWN_USEREXPERIENCEANALYTICSMACHINETYPE UserExperienceAnalyticsMachineType = iota
    // Indicates that the Machine is physical.
    PHYSICAL_USEREXPERIENCEANALYTICSMACHINETYPE
    // Indicates that the machine is virtual.
    VIRTUAL_USEREXPERIENCEANALYTICSMACHINETYPE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_USEREXPERIENCEANALYTICSMACHINETYPE
)

func (i UserExperienceAnalyticsMachineType) String() string {
    return []string{"unknown", "physical", "virtual", "unknownFutureValue"}[i]
}
func ParseUserExperienceAnalyticsMachineType(v string) (any, error) {
    result := UNKNOWN_USEREXPERIENCEANALYTICSMACHINETYPE
    switch v {
        case "unknown":
            result = UNKNOWN_USEREXPERIENCEANALYTICSMACHINETYPE
        case "physical":
            result = PHYSICAL_USEREXPERIENCEANALYTICSMACHINETYPE
        case "virtual":
            result = VIRTUAL_USEREXPERIENCEANALYTICSMACHINETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_USEREXPERIENCEANALYTICSMACHINETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeUserExperienceAnalyticsMachineType(values []UserExperienceAnalyticsMachineType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i UserExperienceAnalyticsMachineType) isMultiValue() bool {
    return false
}
