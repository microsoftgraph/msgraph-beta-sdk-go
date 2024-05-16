package models
// Locations which can be used to bring data into organization documents
type ManagedAppDataIngestionLocation int

const (
    // OneDrive for business
    ONEDRIVEFORBUSINESS_MANAGEDAPPDATAINGESTIONLOCATION ManagedAppDataIngestionLocation = iota
    // SharePoint Online
    SHAREPOINT_MANAGEDAPPDATAINGESTIONLOCATION
    // The device's camera
    CAMERA_MANAGEDAPPDATAINGESTIONLOCATION
    // The device's photo library
    PHOTOLIBRARY_MANAGEDAPPDATAINGESTIONLOCATION
)

func (i ManagedAppDataIngestionLocation) String() string {
    return []string{"oneDriveForBusiness", "sharePoint", "camera", "photoLibrary"}[i]
}
func ParseManagedAppDataIngestionLocation(v string) (any, error) {
    result := ONEDRIVEFORBUSINESS_MANAGEDAPPDATAINGESTIONLOCATION
    switch v {
        case "oneDriveForBusiness":
            result = ONEDRIVEFORBUSINESS_MANAGEDAPPDATAINGESTIONLOCATION
        case "sharePoint":
            result = SHAREPOINT_MANAGEDAPPDATAINGESTIONLOCATION
        case "camera":
            result = CAMERA_MANAGEDAPPDATAINGESTIONLOCATION
        case "photoLibrary":
            result = PHOTOLIBRARY_MANAGEDAPPDATAINGESTIONLOCATION
        default:
            return nil, nil
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
func (i ManagedAppDataIngestionLocation) isMultiValue() bool {
    return false
}
