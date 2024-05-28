package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilder provides operations to manage the deviceAppPerformances property of the microsoft.graph.managedTenants.managedTenant entity.
type ManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilderGetQueryParameters get deviceAppPerformances from tenantRelationships
type ManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilderGetQueryParameters struct {
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
// ManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilderGetQueryParameters
}
// ManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceAppPerformanceId provides operations to manage the deviceAppPerformances property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsDeviceappperformancesDeviceAppPerformanceItemRequestBuilder when successful
func (m *ManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilder) ByDeviceAppPerformanceId(deviceAppPerformanceId string)(*ManagedtenantsDeviceappperformancesDeviceAppPerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceAppPerformanceId != "" {
        urlTplParams["deviceAppPerformance%2Did"] = deviceAppPerformanceId
    }
    return NewManagedtenantsDeviceappperformancesDeviceAppPerformanceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilderInternal instantiates a new ManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilder and sets the default values.
func NewManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilder) {
    m := &ManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/deviceAppPerformances{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilder instantiates a new ManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilder and sets the default values.
func NewManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ManagedtenantsDeviceappperformancesCountRequestBuilder when successful
func (m *ManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilder) Count()(*ManagedtenantsDeviceappperformancesCountRequestBuilder) {
    return NewManagedtenantsDeviceappperformancesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get deviceAppPerformances from tenantRelationships
// returns a DeviceAppPerformanceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.DeviceAppPerformanceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateDeviceAppPerformanceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.DeviceAppPerformanceCollectionResponseable), nil
}
// Post create new navigation property to deviceAppPerformances for tenantRelationships
// returns a DeviceAppPerformanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilder) Post(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.DeviceAppPerformanceable, requestConfiguration *ManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilderPostRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.DeviceAppPerformanceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateDeviceAppPerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.DeviceAppPerformanceable), nil
}
// ToGetRequestInformation get deviceAppPerformances from tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to deviceAppPerformances for tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilder) ToPostRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.DeviceAppPerformanceable, requestConfiguration *ManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilder when successful
func (m *ManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilder) WithUrl(rawUrl string)(*ManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilder) {
    return NewManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
