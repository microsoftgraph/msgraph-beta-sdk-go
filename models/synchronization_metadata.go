package models
import (
    "errors"
)
// 
type SynchronizationMetadata int

const (
    GALLERYAPPLICATIONIDENTIFIER_SYNCHRONIZATIONMETADATA SynchronizationMetadata = iota
    GALLERYAPPLICATIONKEY_SYNCHRONIZATIONMETADATA
    ISOAUTHENABLED_SYNCHRONIZATIONMETADATA
    ISSYNCHRONIZATIONAGENTASSIGNMENTREQUIRED_SYNCHRONIZATIONMETADATA
    ISSYNCHRONIZATIONAGENTREQUIRED_SYNCHRONIZATIONMETADATA
    ISSYNCHRONIZATIONINPREVIEW_SYNCHRONIZATIONMETADATA
    OAUTHSETTINGS_SYNCHRONIZATIONMETADATA
    SYNCHRONIZATIONLEARNMOREIBIZAFWLINK_SYNCHRONIZATIONMETADATA
    CONFIGURATIONFIELDS_SYNCHRONIZATIONMETADATA
)

func (i SynchronizationMetadata) String() string {
    return []string{"galleryApplicationIdentifier", "galleryApplicationKey", "isOAuthEnabled", "IsSynchronizationAgentAssignmentRequired", "isSynchronizationAgentRequired", "isSynchronizationInPreview", "oAuthSettings", "synchronizationLearnMoreIbizaFwLink", "configurationFields"}[i]
}
func ParseSynchronizationMetadata(v string) (any, error) {
    result := GALLERYAPPLICATIONIDENTIFIER_SYNCHRONIZATIONMETADATA
    switch v {
        case "galleryApplicationIdentifier":
            result = GALLERYAPPLICATIONIDENTIFIER_SYNCHRONIZATIONMETADATA
        case "galleryApplicationKey":
            result = GALLERYAPPLICATIONKEY_SYNCHRONIZATIONMETADATA
        case "isOAuthEnabled":
            result = ISOAUTHENABLED_SYNCHRONIZATIONMETADATA
        case "IsSynchronizationAgentAssignmentRequired":
            result = ISSYNCHRONIZATIONAGENTASSIGNMENTREQUIRED_SYNCHRONIZATIONMETADATA
        case "isSynchronizationAgentRequired":
            result = ISSYNCHRONIZATIONAGENTREQUIRED_SYNCHRONIZATIONMETADATA
        case "isSynchronizationInPreview":
            result = ISSYNCHRONIZATIONINPREVIEW_SYNCHRONIZATIONMETADATA
        case "oAuthSettings":
            result = OAUTHSETTINGS_SYNCHRONIZATIONMETADATA
        case "synchronizationLearnMoreIbizaFwLink":
            result = SYNCHRONIZATIONLEARNMOREIBIZAFWLINK_SYNCHRONIZATIONMETADATA
        case "configurationFields":
            result = CONFIGURATIONFIELDS_SYNCHRONIZATIONMETADATA
        default:
            return 0, errors.New("Unknown SynchronizationMetadata value: " + v)
    }
    return &result, nil
}
func SerializeSynchronizationMetadata(values []SynchronizationMetadata) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SynchronizationMetadata) isMultiValue() bool {
    return false
}
