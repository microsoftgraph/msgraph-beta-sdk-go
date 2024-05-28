package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilder provides operations to manage the deviceConfigurationConflictSummary property of the microsoft.graph.deviceManagement entity.
type DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilderGetQueryParameters summary of policies in conflict state for this account.
type DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilderGetQueryParameters struct {
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
// DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilderGetQueryParameters
}
// DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceConfigurationConflictSummaryId provides operations to manage the deviceConfigurationConflictSummary property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryItemRequestBuilder when successful
func (m *DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilder) ByDeviceConfigurationConflictSummaryId(deviceConfigurationConflictSummaryId string)(*DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceConfigurationConflictSummaryId != "" {
        urlTplParams["deviceConfigurationConflictSummary%2Did"] = deviceConfigurationConflictSummaryId
    }
    return NewDeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilderInternal instantiates a new DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilder and sets the default values.
func NewDeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilder) {
    m := &DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceConfigurationConflictSummary{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilder instantiates a new DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilder and sets the default values.
func NewDeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DeviceconfigurationconflictsummaryCountRequestBuilder when successful
func (m *DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilder) Count()(*DeviceconfigurationconflictsummaryCountRequestBuilder) {
    return NewDeviceconfigurationconflictsummaryCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get summary of policies in conflict state for this account.
// returns a DeviceConfigurationConflictSummaryCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationConflictSummaryCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceConfigurationConflictSummaryCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationConflictSummaryCollectionResponseable), nil
}
// Post create new navigation property to deviceConfigurationConflictSummary for deviceManagement
// returns a DeviceConfigurationConflictSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationConflictSummaryable, requestConfiguration *DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationConflictSummaryable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceConfigurationConflictSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationConflictSummaryable), nil
}
// ToGetRequestInformation summary of policies in conflict state for this account.
// returns a *RequestInformation when successful
func (m *DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to deviceConfigurationConflictSummary for deviceManagement
// returns a *RequestInformation when successful
func (m *DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationConflictSummaryable, requestConfiguration *DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilder when successful
func (m *DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilder) WithUrl(rawUrl string)(*DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilder) {
    return NewDeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
