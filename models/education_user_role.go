package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the educationRoot singleton.
type EducationUserRole int

const (
    STUDENT_EDUCATIONUSERROLE EducationUserRole = iota
    TEACHER_EDUCATIONUSERROLE
    NONE_EDUCATIONUSERROLE
    UNKNOWNFUTUREVALUE_EDUCATIONUSERROLE
    FACULTY_EDUCATIONUSERROLE
)

func (i EducationUserRole) String() string {
    return []string{"STUDENT", "TEACHER", "NONE", "UNKNOWNFUTUREVALUE", "FACULTY"}[i]
}
func ParseEducationUserRole(v string) (interface{}, error) {
    result := STUDENT_EDUCATIONUSERROLE
    switch strings.ToUpper(v) {
        case "STUDENT":
            result = STUDENT_EDUCATIONUSERROLE
        case "TEACHER":
            result = TEACHER_EDUCATIONUSERROLE
        case "NONE":
            result = NONE_EDUCATIONUSERROLE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_EDUCATIONUSERROLE
        case "FACULTY":
            result = FACULTY_EDUCATIONUSERROLE
        default:
            return 0, errors.New("Unknown EducationUserRole value: " + v)
    }
    return &result, nil
}
func SerializeEducationUserRole(values []EducationUserRole) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}