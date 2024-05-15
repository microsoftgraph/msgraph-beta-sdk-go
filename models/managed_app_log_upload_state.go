package models
// Represents the current status of the associated `managedAppLogCollectionRequest`.
type ManagedAppLogUploadState int

const (
    // Default. The log upload request has been successfully created, and is pending delivery to the Mobile Application Management (MAM) application.
    PENDING_MANAGEDAPPLOGUPLOADSTATE ManagedAppLogUploadState = iota
    // One or more log upload components have uploaded their logs.
    INPROGRESS_MANAGEDAPPLOGUPLOADSTATE
    // All log upload successfully components have uploaded their logs.
    COMPLETED_MANAGEDAPPLOGUPLOADSTATE
    // The log upload request was declined by the user on the device.
    DECLINEDBYUSER_MANAGEDAPPLOGUPLOADSTATE
    // The log upload request was not acknowledged by the user within the allowed time window.
    TIMEDOUT_MANAGEDAPPLOGUPLOADSTATE
    // The log upload request encountered an error.
    FAILED_MANAGEDAPPLOGUPLOADSTATE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_MANAGEDAPPLOGUPLOADSTATE
)

func (i ManagedAppLogUploadState) String() string {
    return []string{"pending", "inProgress", "completed", "declinedByUser", "timedOut", "failed", "unknownFutureValue"}[i]
}
func ParseManagedAppLogUploadState(v string) (any, error) {
    result := PENDING_MANAGEDAPPLOGUPLOADSTATE
    switch v {
        case "pending":
            result = PENDING_MANAGEDAPPLOGUPLOADSTATE
        case "inProgress":
            result = INPROGRESS_MANAGEDAPPLOGUPLOADSTATE
        case "completed":
            result = COMPLETED_MANAGEDAPPLOGUPLOADSTATE
        case "declinedByUser":
            result = DECLINEDBYUSER_MANAGEDAPPLOGUPLOADSTATE
        case "timedOut":
            result = TIMEDOUT_MANAGEDAPPLOGUPLOADSTATE
        case "failed":
            result = FAILED_MANAGEDAPPLOGUPLOADSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MANAGEDAPPLOGUPLOADSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeManagedAppLogUploadState(values []ManagedAppLogUploadState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ManagedAppLogUploadState) isMultiValue() bool {
    return false
}
