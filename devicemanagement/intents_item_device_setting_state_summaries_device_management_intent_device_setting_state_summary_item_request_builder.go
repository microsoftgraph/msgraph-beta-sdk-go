package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder provides operations to manage the deviceSettingStateSummaries property of the microsoft.graph.deviceManagementIntent entity.
type IntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderGetQueryParameters collection of settings and their states and counts of devices that belong to corresponding state for all settings within the intent
type IntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderGetQueryParameters
}
// IntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderInternal instantiates a new DeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder and sets the default values.
func NewIntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder) {
    m := &IntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/intents/{deviceManagementIntent%2Did}/deviceSettingStateSummaries/{deviceManagementIntentDeviceSettingStateSummary%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewIntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder instantiates a new DeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder and sets the default values.
func NewIntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceSettingStateSummaries for deviceManagement
func (m *IntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get collection of settings and their states and counts of devices that belong to corresponding state for all settings within the intent
func (m *IntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentDeviceSettingStateSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *IntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentDeviceSettingStateSummaryable, requestConfiguration *IntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentDeviceSettingStateSummaryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *IntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *IntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation collection of settings and their states and counts of devices that belong to corresponding state for all settings within the intent
func (m *IntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *IntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementIntentDeviceSettingStateSummaryable, requestConfiguration *IntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *IntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder) WithUrl(rawUrl string)(*IntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder) {
    return NewIntentsItemDeviceSettingStateSummariesDeviceManagementIntentDeviceSettingStateSummaryItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
