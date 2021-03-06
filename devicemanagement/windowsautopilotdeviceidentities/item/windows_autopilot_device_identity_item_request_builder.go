package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0603021b313096a6726ff00fcc4dfd8a8b9c2fc89775400414461f4d36fd7d04 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeviceidentities/item/assignresourceaccounttodevice"
    i08637003c24136cfb79d3be758b142dd1115777f5410f7f104df5f6983cc4367 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeviceidentities/item/deploymentprofile"
    i3f74d7cb02f41203c9cd7d155fac63204acbe2e5d81f2402d7b639e0c9a7b5f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeviceidentities/item/assignusertodevice"
    i7336d162e1b3abcfc9696b9aee7f016a3c7fa2dbb98e24cf3154a91e3365d661 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeviceidentities/item/unassignresourceaccountfromdevice"
    i83461ac3f25e4e8a3194b7af8d988758e07ed71d4aa5b3f61ded8939dc620c98 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeviceidentities/item/updatedeviceproperties"
    id6142e63ae5bd7453e3630d6c28f631a110e460306d06d64699f247d9bb218f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeviceidentities/item/intendeddeploymentprofile"
    iea2a6744b4e629234021d6f6c8e9fa076a3ab8a7876d7c57abc69043cf535a5b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeviceidentities/item/unassignuserfromdevice"
)

// WindowsAutopilotDeviceIdentityItemRequestBuilder provides operations to manage the windowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
type WindowsAutopilotDeviceIdentityItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// WindowsAutopilotDeviceIdentityItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsAutopilotDeviceIdentityItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsAutopilotDeviceIdentityItemRequestBuilderGetQueryParameters the Windows autopilot device identities contained collection.
type WindowsAutopilotDeviceIdentityItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsAutopilotDeviceIdentityItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsAutopilotDeviceIdentityItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsAutopilotDeviceIdentityItemRequestBuilderGetQueryParameters
}
// WindowsAutopilotDeviceIdentityItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsAutopilotDeviceIdentityItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssignResourceAccountToDevice the assignResourceAccountToDevice property
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) AssignResourceAccountToDevice()(*i0603021b313096a6726ff00fcc4dfd8a8b9c2fc89775400414461f4d36fd7d04.AssignResourceAccountToDeviceRequestBuilder) {
    return i0603021b313096a6726ff00fcc4dfd8a8b9c2fc89775400414461f4d36fd7d04.NewAssignResourceAccountToDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignUserToDevice the assignUserToDevice property
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) AssignUserToDevice()(*i3f74d7cb02f41203c9cd7d155fac63204acbe2e5d81f2402d7b639e0c9a7b5f6.AssignUserToDeviceRequestBuilder) {
    return i3f74d7cb02f41203c9cd7d155fac63204acbe2e5d81f2402d7b639e0c9a7b5f6.NewAssignUserToDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewWindowsAutopilotDeviceIdentityItemRequestBuilderInternal instantiates a new WindowsAutopilotDeviceIdentityItemRequestBuilder and sets the default values.
func NewWindowsAutopilotDeviceIdentityItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsAutopilotDeviceIdentityItemRequestBuilder) {
    m := &WindowsAutopilotDeviceIdentityItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/windowsAutopilotDeviceIdentities/{windowsAutopilotDeviceIdentity%2Did}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property windowsAutopilotDeviceIdentities for deviceManagement
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property windowsAutopilotDeviceIdentities for deviceManagement
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *WindowsAutopilotDeviceIdentityItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the Windows autopilot device identities contained collection.
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the Windows autopilot device identities contained collection.
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *WindowsAutopilotDeviceIdentityItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property windowsAutopilotDeviceIdentities in deviceManagement
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property windowsAutopilotDeviceIdentities in deviceManagement
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable, requestConfiguration *WindowsAutopilotDeviceIdentityItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property windowsAutopilotDeviceIdentities for deviceManagement
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property windowsAutopilotDeviceIdentities for deviceManagement
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *WindowsAutopilotDeviceIdentityItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeploymentProfile the deploymentProfile property
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) DeploymentProfile()(*i08637003c24136cfb79d3be758b142dd1115777f5410f7f104df5f6983cc4367.DeploymentProfileRequestBuilder) {
    return i08637003c24136cfb79d3be758b142dd1115777f5410f7f104df5f6983cc4367.NewDeploymentProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the Windows autopilot device identities contained collection.
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the Windows autopilot device identities contained collection.
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *WindowsAutopilotDeviceIdentityItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsAutopilotDeviceIdentityFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable), nil
}
// IntendedDeploymentProfile the intendedDeploymentProfile property
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) IntendedDeploymentProfile()(*id6142e63ae5bd7453e3630d6c28f631a110e460306d06d64699f247d9bb218f3.IntendedDeploymentProfileRequestBuilder) {
    return id6142e63ae5bd7453e3630d6c28f631a110e460306d06d64699f247d9bb218f3.NewIntendedDeploymentProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property windowsAutopilotDeviceIdentities in deviceManagement
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property windowsAutopilotDeviceIdentities in deviceManagement
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable, requestConfiguration *WindowsAutopilotDeviceIdentityItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// UnassignResourceAccountFromDevice the unassignResourceAccountFromDevice property
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) UnassignResourceAccountFromDevice()(*i7336d162e1b3abcfc9696b9aee7f016a3c7fa2dbb98e24cf3154a91e3365d661.UnassignResourceAccountFromDeviceRequestBuilder) {
    return i7336d162e1b3abcfc9696b9aee7f016a3c7fa2dbb98e24cf3154a91e3365d661.NewUnassignResourceAccountFromDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnassignUserFromDevice the unassignUserFromDevice property
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) UnassignUserFromDevice()(*iea2a6744b4e629234021d6f6c8e9fa076a3ab8a7876d7c57abc69043cf535a5b.UnassignUserFromDeviceRequestBuilder) {
    return iea2a6744b4e629234021d6f6c8e9fa076a3ab8a7876d7c57abc69043cf535a5b.NewUnassignUserFromDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateDeviceProperties the updateDeviceProperties property
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) UpdateDeviceProperties()(*i83461ac3f25e4e8a3194b7af8d988758e07ed71d4aa5b3f61ded8939dc620c98.UpdateDevicePropertiesRequestBuilder) {
    return i83461ac3f25e4e8a3194b7af8d988758e07ed71d4aa5b3f61ded8939dc620c98.NewUpdateDevicePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
