package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i158b2f1ed97ea7c14fab9f68b26bf476d3708afeb068765960eaaf699af1b5f8 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeploymentprofiles/item/assigneddevices/item/updatedeviceproperties"
    i6305748a7c20ec648805e3bbbb3193245a7f1126da5d98e5b9b3c62a480a077f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeploymentprofiles/item/assigneddevices/item/assignusertodevice"
    i6d6560c766eb2ce774731049fc529d3b519832e9f2f53682fbfe7d35df7c1c3c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeploymentprofiles/item/assigneddevices/item/assignresourceaccounttodevice"
    i83c5f499364f220f25e381f8075396532f9c977c78fde505fd17f38a1411ece0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeploymentprofiles/item/assigneddevices/item/unassignuserfromdevice"
    ib73822b720982991e43913cc2d8c66067380da4aa5c84e88e3f2465bcd4f0a23 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeploymentprofiles/item/assigneddevices/item/deploymentprofile"
    ica77650b1b796d69d225beec92c6b9e45081ed70aef73755c178ffbf8b628d4f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeploymentprofiles/item/assigneddevices/item/unassignresourceaccountfromdevice"
    ifaa60ce9423a2f2f03b8b75f54d288b9a3382a7e5693b4c28468c616d6306ded "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeploymentprofiles/item/assigneddevices/item/intendeddeploymentprofile"
)

// WindowsAutopilotDeviceIdentityItemRequestBuilder provides operations to manage the assignedDevices property of the microsoft.graph.windowsAutopilotDeploymentProfile entity.
type WindowsAutopilotDeviceIdentityItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// WindowsAutopilotDeviceIdentityItemRequestBuilderDeleteOptions options for Delete
type WindowsAutopilotDeviceIdentityItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// WindowsAutopilotDeviceIdentityItemRequestBuilderGetOptions options for Get
type WindowsAutopilotDeviceIdentityItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsAutopilotDeviceIdentityItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// WindowsAutopilotDeviceIdentityItemRequestBuilderGetQueryParameters the list of assigned devices for the profile.
type WindowsAutopilotDeviceIdentityItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsAutopilotDeviceIdentityItemRequestBuilderPatchOptions options for Patch
type WindowsAutopilotDeviceIdentityItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// AssignResourceAccountToDevice the assignResourceAccountToDevice property
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) AssignResourceAccountToDevice()(*i6d6560c766eb2ce774731049fc529d3b519832e9f2f53682fbfe7d35df7c1c3c.AssignResourceAccountToDeviceRequestBuilder) {
    return i6d6560c766eb2ce774731049fc529d3b519832e9f2f53682fbfe7d35df7c1c3c.NewAssignResourceAccountToDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignUserToDevice the assignUserToDevice property
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) AssignUserToDevice()(*i6305748a7c20ec648805e3bbbb3193245a7f1126da5d98e5b9b3c62a480a077f.AssignUserToDeviceRequestBuilder) {
    return i6305748a7c20ec648805e3bbbb3193245a7f1126da5d98e5b9b3c62a480a077f.NewAssignUserToDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewWindowsAutopilotDeviceIdentityItemRequestBuilderInternal instantiates a new WindowsAutopilotDeviceIdentityItemRequestBuilder and sets the default values.
func NewWindowsAutopilotDeviceIdentityItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsAutopilotDeviceIdentityItemRequestBuilder) {
    m := &WindowsAutopilotDeviceIdentityItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/windowsAutopilotDeploymentProfiles/{windowsAutopilotDeploymentProfile%2Did}/assignedDevices/{windowsAutopilotDeviceIdentity%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWindowsAutopilotDeviceIdentityItemRequestBuilder instantiates a new WindowsAutopilotDeviceIdentityItemRequestBuilder and sets the default values.
func NewWindowsAutopilotDeviceIdentityItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsAutopilotDeviceIdentityItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsAutopilotDeviceIdentityItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property assignedDevices for deviceManagement
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) CreateDeleteRequestInformation(options *WindowsAutopilotDeviceIdentityItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the list of assigned devices for the profile.
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) CreateGetRequestInformation(options *WindowsAutopilotDeviceIdentityItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property assignedDevices in deviceManagement
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) CreatePatchRequestInformation(options *WindowsAutopilotDeviceIdentityItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property assignedDevices for deviceManagement
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) Delete(options *WindowsAutopilotDeviceIdentityItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeploymentProfile the deploymentProfile property
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) DeploymentProfile()(*ib73822b720982991e43913cc2d8c66067380da4aa5c84e88e3f2465bcd4f0a23.DeploymentProfileRequestBuilder) {
    return ib73822b720982991e43913cc2d8c66067380da4aa5c84e88e3f2465bcd4f0a23.NewDeploymentProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the list of assigned devices for the profile.
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) Get(options *WindowsAutopilotDeviceIdentityItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsAutopilotDeviceIdentityFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable), nil
}
// IntendedDeploymentProfile the intendedDeploymentProfile property
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) IntendedDeploymentProfile()(*ifaa60ce9423a2f2f03b8b75f54d288b9a3382a7e5693b4c28468c616d6306ded.IntendedDeploymentProfileRequestBuilder) {
    return ifaa60ce9423a2f2f03b8b75f54d288b9a3382a7e5693b4c28468c616d6306ded.NewIntendedDeploymentProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property assignedDevices in deviceManagement
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) Patch(options *WindowsAutopilotDeviceIdentityItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// UnassignResourceAccountFromDevice the unassignResourceAccountFromDevice property
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) UnassignResourceAccountFromDevice()(*ica77650b1b796d69d225beec92c6b9e45081ed70aef73755c178ffbf8b628d4f.UnassignResourceAccountFromDeviceRequestBuilder) {
    return ica77650b1b796d69d225beec92c6b9e45081ed70aef73755c178ffbf8b628d4f.NewUnassignResourceAccountFromDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnassignUserFromDevice the unassignUserFromDevice property
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) UnassignUserFromDevice()(*i83c5f499364f220f25e381f8075396532f9c977c78fde505fd17f38a1411ece0.UnassignUserFromDeviceRequestBuilder) {
    return i83c5f499364f220f25e381f8075396532f9c977c78fde505fd17f38a1411ece0.NewUnassignUserFromDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateDeviceProperties the updateDeviceProperties property
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) UpdateDeviceProperties()(*i158b2f1ed97ea7c14fab9f68b26bf476d3708afeb068765960eaaf699af1b5f8.UpdateDevicePropertiesRequestBuilder) {
    return i158b2f1ed97ea7c14fab9f68b26bf476d3708afeb068765960eaaf699af1b5f8.NewUpdateDevicePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
