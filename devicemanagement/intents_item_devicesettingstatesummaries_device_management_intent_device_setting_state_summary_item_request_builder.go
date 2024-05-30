package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder provides operations to manage the deviceSettingStateSummaries property of the microsoft.graph.deviceManagementIntent entity.
type IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderGetQueryParameters collection of settings and their states and counts of devices that belong to corresponding state for all settings within the intent
type IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderGetQueryParameters
}
// IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderInternal instantiates a new IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder and sets the default values.
func NewIntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder) {
    m := &IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/intents/{deviceManagementIntent%2Did}/deviceSettingStateSummaries/{deviceManagementIntentDeviceSettingStateSummary%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewIntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder instantiates a new IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder and sets the default values.
func NewIntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceSettingStateSummaries for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get collection of settings and their states and counts of devices that belong to corresponding state for all settings within the intent
// returns a DeviceManagementIntentDeviceSettingStateSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentDeviceSettingStateSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementIntentDeviceSettingStateSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentDeviceSettingStateSummaryable), nil
}
// Patch update the navigation property deviceSettingStateSummaries in deviceManagement
// returns a DeviceManagementIntentDeviceSettingStateSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentDeviceSettingStateSummaryable, requestConfiguration *IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentDeviceSettingStateSummaryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementIntentDeviceSettingStateSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentDeviceSettingStateSummaryable), nil
}
// ToDeleteRequestInformation delete navigation property deviceSettingStateSummaries for deviceManagement
// returns a *RequestInformation when successful
func (m *IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation collection of settings and their states and counts of devices that belong to corresponding state for all settings within the intent
// returns a *RequestInformation when successful
func (m *IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property deviceSettingStateSummaries in deviceManagement
// returns a *RequestInformation when successful
func (m *IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentDeviceSettingStateSummaryable, requestConfiguration *IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder when successful
func (m *IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder) WithUrl(rawUrl string)(*IntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder) {
    return NewIntentsItemDevicesettingstatesummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
