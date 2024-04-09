package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder provides operations to manage the windowsAutopilotDeploymentProfiles property of the microsoft.graph.deviceManagement entity.
type WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderGetQueryParameters windows auto pilot deployment profiles
type WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderGetQueryParameters
}
// WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
// returns a *WindowsAutopilotDeploymentProfilesItemAssignRequestBuilder when successful
func (m *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) Assign()(*WindowsAutopilotDeploymentProfilesItemAssignRequestBuilder) {
    return NewWindowsAutopilotDeploymentProfilesItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignedDevices provides operations to manage the assignedDevices property of the microsoft.graph.windowsAutopilotDeploymentProfile entity.
// returns a *WindowsAutopilotDeploymentProfilesItemAssignedDevicesRequestBuilder when successful
func (m *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) AssignedDevices()(*WindowsAutopilotDeploymentProfilesItemAssignedDevicesRequestBuilder) {
    return NewWindowsAutopilotDeploymentProfilesItemAssignedDevicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.windowsAutopilotDeploymentProfile entity.
// returns a *WindowsAutopilotDeploymentProfilesItemAssignmentsRequestBuilder when successful
func (m *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) Assignments()(*WindowsAutopilotDeploymentProfilesItemAssignmentsRequestBuilder) {
    return NewWindowsAutopilotDeploymentProfilesItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderInternal instantiates a new WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder and sets the default values.
func NewWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) {
    m := &WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/windowsAutopilotDeploymentProfiles/{windowsAutopilotDeploymentProfile%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder instantiates a new WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder and sets the default values.
func NewWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property windowsAutopilotDeploymentProfiles for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get windows auto pilot deployment profiles
// returns a WindowsAutopilotDeploymentProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeploymentProfileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsAutopilotDeploymentProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeploymentProfileable), nil
}
// Patch update the navigation property windowsAutopilotDeploymentProfiles in deviceManagement
// returns a WindowsAutopilotDeploymentProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeploymentProfileable, requestConfiguration *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeploymentProfileable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsAutopilotDeploymentProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeploymentProfileable), nil
}
// ToDeleteRequestInformation delete navigation property windowsAutopilotDeploymentProfiles for deviceManagement
// returns a *RequestInformation when successful
func (m *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation windows auto pilot deployment profiles
// returns a *RequestInformation when successful
func (m *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property windowsAutopilotDeploymentProfiles in deviceManagement
// returns a *RequestInformation when successful
func (m *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeploymentProfileable, requestConfiguration *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder when successful
func (m *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) WithUrl(rawUrl string)(*WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) {
    return NewWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
