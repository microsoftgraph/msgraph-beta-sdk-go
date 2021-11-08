package targetedmanagedappprotection

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i6f58b8dfa7321456be9f892045a6baaee49fd84762f55ddf6e1144f26aa29cc7 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedapppolicies/item/targetedmanagedappprotection/assign"
    i9932d162e17756ab096dc5ac32882861e26bc2f82c724b4292070852f49b8935 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedapppolicies/item/targetedmanagedappprotection/targetapps"
)

// Builds and executes requests for operations under \deviceAppManagement\managedAppPolicies\{managedAppPolicy-id}\microsoft.graph.targetedManagedAppProtection
type TargetedManagedAppProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *TargetedManagedAppProtectionRequestBuilder) Assign()(*i6f58b8dfa7321456be9f892045a6baaee49fd84762f55ddf6e1144f26aa29cc7.AssignRequestBuilder) {
    return i6f58b8dfa7321456be9f892045a6baaee49fd84762f55ddf6e1144f26aa29cc7.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new TargetedManagedAppProtectionRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewTargetedManagedAppProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TargetedManagedAppProtectionRequestBuilder) {
    m := &TargetedManagedAppProtectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/managedAppPolicies/{managedAppPolicy_id}/microsoft.graph.targetedManagedAppProtection";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new TargetedManagedAppProtectionRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewTargetedManagedAppProtectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TargetedManagedAppProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTargetedManagedAppProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *TargetedManagedAppProtectionRequestBuilder) TargetApps()(*i9932d162e17756ab096dc5ac32882861e26bc2f82c724b4292070852f49b8935.TargetAppsRequestBuilder) {
    return i9932d162e17756ab096dc5ac32882861e26bc2f82c724b4292070852f49b8935.NewTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
