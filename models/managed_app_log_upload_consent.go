package models
import (
    "errors"
)
// Represents the current consent status of the associated `managedAppLogCollectionRequest`.
type ManagedAppLogUploadConsent int

const (
    // Default. Indicates app log consent state is 'Unknown'. This state is automatically assigned at request creation time and is updated when the log collection completes.
    UNKNOWN_MANAGEDAPPLOGUPLOADCONSENT ManagedAppLogUploadConsent = iota
    // The User has Declined the Log Collection Request. The Log collection and uploads will not be initiated/triggered, and the log collection request will be abandoned.
    DECLINED_MANAGEDAPPLOGUPLOADCONSENT
    // The User has Accepted the Log Collection Request. The log collection and upload will be initiated.
    ACCEPTED_MANAGEDAPPLOGUPLOADCONSENT
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_MANAGEDAPPLOGUPLOADCONSENT
)

func (i ManagedAppLogUploadConsent) String() string {
    return []string{"unknown", "declined", "accepted", "unknownFutureValue"}[i]
}
func ParseManagedAppLogUploadConsent(v string) (any, error) {
    result := UNKNOWN_MANAGEDAPPLOGUPLOADCONSENT
    switch v {
        case "unknown":
            result = UNKNOWN_MANAGEDAPPLOGUPLOADCONSENT
        case "declined":
            result = DECLINED_MANAGEDAPPLOGUPLOADCONSENT
        case "accepted":
            result = ACCEPTED_MANAGEDAPPLOGUPLOADCONSENT
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MANAGEDAPPLOGUPLOADCONSENT
        default:
            return 0, errors.New("Unknown ManagedAppLogUploadConsent value: " + v)
    }
    return &result, nil
}
func SerializeManagedAppLogUploadConsent(values []ManagedAppLogUploadConsent) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ManagedAppLogUploadConsent) isMultiValue() bool {
    return false
}
