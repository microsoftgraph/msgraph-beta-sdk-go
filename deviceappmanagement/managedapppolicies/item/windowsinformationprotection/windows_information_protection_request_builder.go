package windowsinformationprotection

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    id6b454652f426ba2d2dbcd4b7e28483e23cc3e4171a1cc413c6fde8042d806d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedapppolicies/item/windowsinformationprotection/assign"
)

// WindowsInformationProtectionRequestBuilder builds and executes requests for operations under \deviceAppManagement\managedAppPolicies\{managedAppPolicy-id}\microsoft.graph.windowsInformationProtection
type WindowsInformationProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Assign the assign property
func (m *WindowsInformationProtectionRequestBuilder) Assign()(*id6b454652f426ba2d2dbcd4b7e28483e23cc3e4171a1cc413c6fde8042d806d7.AssignRequestBuilder) {
    return id6b454652f426ba2d2dbcd4b7e28483e23cc3e4171a1cc413c6fde8042d806d7.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewWindowsInformationProtectionRequestBuilderInternal instantiates a new WindowsInformationProtectionRequestBuilder and sets the default values.
func NewWindowsInformationProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsInformationProtectionRequestBuilder) {
    m := &WindowsInformationProtectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/managedAppPolicies/{managedAppPolicy%2Did}/microsoft.graph.windowsInformationProtection";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWindowsInformationProtectionRequestBuilder instantiates a new WindowsInformationProtectionRequestBuilder and sets the default values.
func NewWindowsInformationProtectionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsInformationProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsInformationProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
