package models
import (
    "errors"
)
// Provides operations to manage the deviceAppManagement singleton.
type ManagedAppDataIngestionLocation int

const (
    // OneDrive for business
    ONEDRIVEFORBUSINESS_MANAGEDAPPDATAINGESTIONLOCATION ManagedAppDataIngestionLocation = iota
    // SharePoint Online
    SHAREPOINT_MANAGEDAPPDATAINGESTIONLOCATION
    // The device's camera
    CAMERA_MANAGEDAPPDATAINGESTIONLOCATION
)

func (i ManagedAppDataIngestionLocation) String() string {
    return []string{"oneDriveForBusiness", "sharePoint", "camera"}[i]
}
func ParseManagedAppDataIngestionLocation(v string) (interface{}, error) {
    result := ONEDRIVEFORBUSINESS_MANAGEDAPPDATAINGESTIONLOCATION
    switch v {
        case "oneDriveForBusiness":
            result = ONEDRIVEFORBUSINESS_MANAGEDAPPDATAINGESTIONLOCATION
        case "sharePoint":
            result = SHAREPOINT_MANAGEDAPPDATAINGESTIONLOCATION
        case "camera":
            result = CAMERA_MANAGEDAPPDATAINGESTIONLOCATION
        default:
            return 0, errors.New("Unknown ManagedAppDataIngestionLocation value: " + v)
    }
    return &result, nil
}
func SerializeManagedAppDataIngestionLocation(values []ManagedAppDataIngestionLocation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
