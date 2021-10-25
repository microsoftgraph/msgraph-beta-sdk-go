package graph
import (
    "strings"
    "errors"
)
type EducationSynchronizationProfileState int

const (
    DELETING_EDUCATIONSYNCHRONIZATIONPROFILESTATE EducationSynchronizationProfileState = iota
    DELETIONFAILED_EDUCATIONSYNCHRONIZATIONPROFILESTATE
    PROVISIONINGFAILED_EDUCATIONSYNCHRONIZATIONPROFILESTATE
    PROVISIONED_EDUCATIONSYNCHRONIZATIONPROFILESTATE
    PROVISIONING_EDUCATIONSYNCHRONIZATIONPROFILESTATE
    UNKNOWNFUTUREVALUE_EDUCATIONSYNCHRONIZATIONPROFILESTATE
)

func (i EducationSynchronizationProfileState) String() string {
    return []string{"DELETING", "DELETIONFAILED", "PROVISIONINGFAILED", "PROVISIONED", "PROVISIONING", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseEducationSynchronizationProfileState(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "DELETING":
            return DELETING_EDUCATIONSYNCHRONIZATIONPROFILESTATE, nil
        case "DELETIONFAILED":
            return DELETIONFAILED_EDUCATIONSYNCHRONIZATIONPROFILESTATE, nil
        case "PROVISIONINGFAILED":
            return PROVISIONINGFAILED_EDUCATIONSYNCHRONIZATIONPROFILESTATE, nil
        case "PROVISIONED":
            return PROVISIONED_EDUCATIONSYNCHRONIZATIONPROFILESTATE, nil
        case "PROVISIONING":
            return PROVISIONING_EDUCATIONSYNCHRONIZATIONPROFILESTATE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_EDUCATIONSYNCHRONIZATIONPROFILESTATE, nil
    }
    return 0, errors.New("Unknown EducationSynchronizationProfileState value: " + v)
}
func SerializeEducationSynchronizationProfileState(values []EducationSynchronizationProfileState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
