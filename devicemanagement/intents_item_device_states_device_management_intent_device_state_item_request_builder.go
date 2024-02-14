package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilder provides operations to manage the deviceStates property of the microsoft.graph.deviceManagementIntent entity.
type IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilderGetQueryParameters collection of states of all devices that the intent is applied to
type IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilderGetQueryParameters
}
// IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilderInternal instantiates a new IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilder and sets the default values.
func NewIntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilder) {
    m := &IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/intents/{deviceManagementIntent%2Did}/deviceStates/{deviceManagementIntentDeviceState%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewIntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilder instantiates a new IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilder and sets the default values.
func NewIntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceStates for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get collection of states of all devices that the intent is applied to
// returns a DeviceManagementIntentDeviceStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentDeviceStateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementIntentDeviceStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentDeviceStateable), nil
}
// Patch update the navigation property deviceStates in deviceManagement
// returns a DeviceManagementIntentDeviceStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentDeviceStateable, requestConfiguration *IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentDeviceStateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementIntentDeviceStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentDeviceStateable), nil
}
// ToDeleteRequestInformation delete navigation property deviceStates for deviceManagement
// returns a *RequestInformation when successful
func (m *IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/deviceManagement/intents/{deviceManagementIntent%2Did}/deviceStates/{deviceManagementIntentDeviceState%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation collection of states of all devices that the intent is applied to
// returns a *RequestInformation when successful
func (m *IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property deviceStates in deviceManagement
// returns a *RequestInformation when successful
func (m *IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentDeviceStateable, requestConfiguration *IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/deviceManagement/intents/{deviceManagementIntent%2Did}/deviceStates/{deviceManagementIntentDeviceState%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilder when successful
func (m *IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilder) WithUrl(rawUrl string)(*IntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilder) {
    return NewIntentsItemDeviceStatesDeviceManagementIntentDeviceStateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
