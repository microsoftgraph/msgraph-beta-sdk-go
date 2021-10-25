package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "ONEDRIVEFORBUSINESS":
            return ONEDRIVEFORBUSINESS_MANAGEDAPPDATAINGESTIONLOCATION, nil
        case "SHAREPOINT":
            return SHAREPOINT_MANAGEDAPPDATAINGESTIONLOCATION, nil
        case "CAMERA":
            return CAMERA_MANAGEDAPPDATAINGESTIONLOCATION, nil
    }
    return 0, errors.New("Unknown ManagedAppDataIngestionLocation value: " + v)
}
func SerializeManagedAppDataIngestionLocation(values []ManagedAppDataIngestionLocation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
