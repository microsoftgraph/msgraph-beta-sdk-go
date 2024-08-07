package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilder provides operations to manage the policyStatusDetails property of the microsoft.graph.deviceManagementAutopilotEvent entity.
type AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilderGetQueryParameters policy and application status details for this device.
type AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilderGetQueryParameters
}
// AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilderInternal instantiates a new AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilder and sets the default values.
func NewAutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilder) {
    m := &AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/autopilotEvents/{deviceManagementAutopilotEvent%2Did}/policyStatusDetails/{deviceManagementAutopilotPolicyStatusDetail%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilder instantiates a new AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilder and sets the default values.
func NewAutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property policyStatusDetails for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get policy and application status details for this device.
// returns a DeviceManagementAutopilotPolicyStatusDetailable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementAutopilotPolicyStatusDetailable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementAutopilotPolicyStatusDetailFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementAutopilotPolicyStatusDetailable), nil
}
// Patch update the navigation property policyStatusDetails in deviceManagement
// returns a DeviceManagementAutopilotPolicyStatusDetailable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementAutopilotPolicyStatusDetailable, requestConfiguration *AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementAutopilotPolicyStatusDetailable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementAutopilotPolicyStatusDetailFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementAutopilotPolicyStatusDetailable), nil
}
// ToDeleteRequestInformation delete navigation property policyStatusDetails for deviceManagement
// returns a *RequestInformation when successful
func (m *AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation policy and application status details for this device.
// returns a *RequestInformation when successful
func (m *AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property policyStatusDetails in deviceManagement
// returns a *RequestInformation when successful
func (m *AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementAutopilotPolicyStatusDetailable, requestConfiguration *AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilder when successful
func (m *AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilder) WithUrl(rawUrl string)(*AutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilder) {
    return NewAutopilotEventsItemPolicyStatusDetailsDeviceManagementAutopilotPolicyStatusDetailItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
