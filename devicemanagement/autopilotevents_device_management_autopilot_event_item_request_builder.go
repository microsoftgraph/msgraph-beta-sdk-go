package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilder provides operations to manage the autopilotEvents property of the microsoft.graph.deviceManagement entity.
type AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilderGetQueryParameters the list of autopilot events for the tenant.
type AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilderGetQueryParameters
}
// AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAutopiloteventsDeviceManagementAutopilotEventItemRequestBuilderInternal instantiates a new AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilder and sets the default values.
func NewAutopiloteventsDeviceManagementAutopilotEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilder) {
    m := &AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/autopilotEvents/{deviceManagementAutopilotEvent%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAutopiloteventsDeviceManagementAutopilotEventItemRequestBuilder instantiates a new AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilder and sets the default values.
func NewAutopiloteventsDeviceManagementAutopilotEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAutopiloteventsDeviceManagementAutopilotEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property autopilotEvents for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of autopilot events for the tenant.
// returns a DeviceManagementAutopilotEventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementAutopilotEventable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementAutopilotEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementAutopilotEventable), nil
}
// Patch update the navigation property autopilotEvents in deviceManagement
// returns a DeviceManagementAutopilotEventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementAutopilotEventable, requestConfiguration *AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementAutopilotEventable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementAutopilotEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementAutopilotEventable), nil
}
// PolicyStatusDetails provides operations to manage the policyStatusDetails property of the microsoft.graph.deviceManagementAutopilotEvent entity.
// returns a *AutopiloteventsItemPolicystatusdetailsPolicyStatusDetailsRequestBuilder when successful
func (m *AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilder) PolicyStatusDetails()(*AutopiloteventsItemPolicystatusdetailsPolicyStatusDetailsRequestBuilder) {
    return NewAutopiloteventsItemPolicystatusdetailsPolicyStatusDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property autopilotEvents for deviceManagement
// returns a *RequestInformation when successful
func (m *AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of autopilot events for the tenant.
// returns a *RequestInformation when successful
func (m *AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property autopilotEvents in deviceManagement
// returns a *RequestInformation when successful
func (m *AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementAutopilotEventable, requestConfiguration *AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilder when successful
func (m *AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilder) WithUrl(rawUrl string)(*AutopiloteventsDeviceManagementAutopilotEventItemRequestBuilder) {
    return NewAutopiloteventsDeviceManagementAutopilotEventItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
