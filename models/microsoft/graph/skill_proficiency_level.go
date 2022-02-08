package graph
import (
    "strings"
    "errors"
)
// 
type SkillProficiencyLevel int

const (
    ELEMENTARY_SKILLPROFICIENCYLEVEL SkillProficiencyLevel = iota
    LIMITEDWORKING_SKILLPROFICIENCYLEVEL
    GENERALPROFESSIONAL_SKILLPROFICIENCYLEVEL
    ADVANCEDPROFESSIONAL_SKILLPROFICIENCYLEVEL
    EXPERT_SKILLPROFICIENCYLEVEL
    UNKNOWNFUTUREVALUE_SKILLPROFICIENCYLEVEL
)

func (i SkillProficiencyLevel) String() string {
    return []string{"ELEMENTARY", "LIMITEDWORKING", "GENERALPROFESSIONAL", "ADVANCEDPROFESSIONAL", "EXPERT", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseSkillProficiencyLevel(v string) (interface{}, error) {
    result := ELEMENTARY_SKILLPROFICIENCYLEVEL
    switch strings.ToUpper(v) {
        case "ELEMENTARY":
            result = ELEMENTARY_SKILLPROFICIENCYLEVEL
        case "LIMITEDWORKING":
            result = LIMITEDWORKING_SKILLPROFICIENCYLEVEL
        case "GENERALPROFESSIONAL":
            result = GENERALPROFESSIONAL_SKILLPROFICIENCYLEVEL
        case "ADVANCEDPROFESSIONAL":
            result = ADVANCEDPROFESSIONAL_SKILLPROFICIENCYLEVEL
        case "EXPERT":
            result = EXPERT_SKILLPROFICIENCYLEVEL
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_SKILLPROFICIENCYLEVEL
        default:
            return 0, errors.New("Unknown SkillProficiencyLevel value: " + v)
    }
    return &result, nil
}
func SerializeSkillProficiencyLevel(values []SkillProficiencyLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
