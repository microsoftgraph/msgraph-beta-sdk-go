package graph
import (
    "strings"
    "errors"
)
// 
type ManagedAppDataIngestionLocation int

const (
    ONEDRIVEFORBUSINESS_MANAGEDAPPDATAINGESTIONLOCATION ManagedAppDataIngestionLocation = iota
    SHAREPOINT_MANAGEDAPPDATAINGESTIONLOCATION
    CAMERA_MANAGEDAPPDATAINGESTIONLOCATION
)

func (i ManagedAppDataIngestionLocation) String() string {
    return []string{"ONEDRIVEFORBUSINESS", "SHAREPOINT", "CAMERA"}[i]
}
func ParseManagedAppDataIngestionLocation(v string) (interface{}, error) {
    result := ONEDRIVEFORBUSINESS_MANAGEDAPPDATAINGESTIONLOCATION
    switch strings.ToUpper(v) {
        case "ONEDRIVEFORBUSINESS":
            result = ONEDRIVEFORBUSINESS_MANAGEDAPPDATAINGESTIONLOCATION
        case "SHAREPOINT":
            result = SHAREPOINT_MANAGEDAPPDATAINGESTIONLOCATION
        case "CAMERA":
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
