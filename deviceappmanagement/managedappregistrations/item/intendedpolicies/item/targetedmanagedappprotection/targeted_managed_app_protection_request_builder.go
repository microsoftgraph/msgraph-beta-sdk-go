package targetedmanagedappprotection

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i2037a96acc6d6ceda8d3fb74c43e1657b95e46707ff3a175fc7cab027d1de1e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item/intendedpolicies/item/targetedmanagedappprotection/assign"
    i978ea4fd48451609ff8d62534ea158a9c36bce809adadb1bc0117689a528940c "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item/intendedpolicies/item/targetedmanagedappprotection/targetapps"
)

// TargetedManagedAppProtectionRequestBuilder builds and executes requests for operations under \deviceAppManagement\managedAppRegistrations\{managedAppRegistration-id}\intendedPolicies\{managedAppPolicy-id}\microsoft.graph.targetedManagedAppProtection
type TargetedManagedAppProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Assign the assign property
func (m *TargetedManagedAppProtectionRequestBuilder) Assign()(*i2037a96acc6d6ceda8d3fb74c43e1657b95e46707ff3a175fc7cab027d1de1e4.AssignRequestBuilder) {
    return i2037a96acc6d6ceda8d3fb74c43e1657b95e46707ff3a175fc7cab027d1de1e4.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTargetedManagedAppProtectionRequestBuilderInternal instantiates a new TargetedManagedAppProtectionRequestBuilder and sets the default values.
func NewTargetedManagedAppProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TargetedManagedAppProtectionRequestBuilder) {
    m := &TargetedManagedAppProtectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/managedAppRegistrations/{managedAppRegistration%2Did}/intendedPolicies/{managedAppPolicy%2Did}/microsoft.graph.targetedManagedAppProtection";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTargetedManagedAppProtectionRequestBuilder instantiates a new TargetedManagedAppProtectionRequestBuilder and sets the default values.
func NewTargetedManagedAppProtectionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TargetedManagedAppProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTargetedManagedAppProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
// TargetApps the targetApps property
func (m *TargetedManagedAppProtectionRequestBuilder) TargetApps()(*i978ea4fd48451609ff8d62534ea158a9c36bce809adadb1bc0117689a528940c.TargetAppsRequestBuilder) {
    return i978ea4fd48451609ff8d62534ea158a9c36bce809adadb1bc0117689a528940c.NewTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
