package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilder provides operations to manage the assignedDevices property of the microsoft.graph.windowsAutopilotDeploymentProfile entity.
type WindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilderGetQueryParameters the list of assigned devices for the profile.
type WindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// WindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilderGetQueryParameters
}
// WindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByWindowsAutopilotDeviceIdentityId provides operations to manage the assignedDevices property of the microsoft.graph.windowsAutopilotDeploymentProfile entity.
// returns a *WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilder when successful
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilder) ByWindowsAutopilotDeviceIdentityId(windowsAutopilotDeviceIdentityId string)(*WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if windowsAutopilotDeviceIdentityId != "" {
        urlTplParams["windowsAutopilotDeviceIdentity%2Did"] = windowsAutopilotDeviceIdentityId
    }
    return NewWindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilderInternal instantiates a new WindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilder and sets the default values.
func NewWindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilder) {
    m := &WindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/windowsAutopilotDeploymentProfiles/{windowsAutopilotDeploymentProfile%2Did}/assignedDevices{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewWindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilder instantiates a new WindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilder and sets the default values.
func NewWindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *WindowsautopilotdeploymentprofilesItemAssigneddevicesCountRequestBuilder when successful
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilder) Count()(*WindowsautopilotdeploymentprofilesItemAssigneddevicesCountRequestBuilder) {
    return NewWindowsautopilotdeploymentprofilesItemAssigneddevicesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of assigned devices for the profile.
// returns a WindowsAutopilotDeviceIdentityCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsAutopilotDeviceIdentityCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityCollectionResponseable), nil
}
// Post create new navigation property to assignedDevices for deviceManagement
// returns a WindowsAutopilotDeviceIdentityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable, requestConfiguration *WindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsAutopilotDeviceIdentityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable), nil
}
// ToGetRequestInformation the list of assigned devices for the profile.
// returns a *RequestInformation when successful
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation create new navigation property to assignedDevices for deviceManagement
// returns a *RequestInformation when successful
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable, requestConfiguration *WindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilder when successful
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilder) WithUrl(rawUrl string)(*WindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilder) {
    return NewWindowsautopilotdeploymentprofilesItemAssigneddevicesAssignedDevicesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
