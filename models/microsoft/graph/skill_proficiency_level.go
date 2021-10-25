package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "ELEMENTARY":
            return ELEMENTARY_SKILLPROFICIENCYLEVEL, nil
        case "LIMITEDWORKING":
            return LIMITEDWORKING_SKILLPROFICIENCYLEVEL, nil
        case "GENERALPROFESSIONAL":
            return GENERALPROFESSIONAL_SKILLPROFICIENCYLEVEL, nil
        case "ADVANCEDPROFESSIONAL":
            return ADVANCEDPROFESSIONAL_SKILLPROFICIENCYLEVEL, nil
        case "EXPERT":
            return EXPERT_SKILLPROFICIENCYLEVEL, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_SKILLPROFICIENCYLEVEL, nil
    }
    return 0, errors.New("Unknown SkillProficiencyLevel value: " + v)
}
func SerializeSkillProficiencyLevel(values []SkillProficiencyLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
