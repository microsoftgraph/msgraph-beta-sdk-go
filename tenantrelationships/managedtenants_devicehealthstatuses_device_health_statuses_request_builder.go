package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilder provides operations to manage the deviceHealthStatuses property of the microsoft.graph.managedTenants.managedTenant entity.
type ManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilderGetQueryParameters get deviceHealthStatuses from tenantRelationships
type ManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilderGetQueryParameters struct {
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
// ManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilderGetQueryParameters
}
// ManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceHealthStatusId provides operations to manage the deviceHealthStatuses property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsDevicehealthstatusesDeviceHealthStatusItemRequestBuilder when successful
func (m *ManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilder) ByDeviceHealthStatusId(deviceHealthStatusId string)(*ManagedtenantsDevicehealthstatusesDeviceHealthStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceHealthStatusId != "" {
        urlTplParams["deviceHealthStatus%2Did"] = deviceHealthStatusId
    }
    return NewManagedtenantsDevicehealthstatusesDeviceHealthStatusItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilderInternal instantiates a new ManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilder and sets the default values.
func NewManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilder) {
    m := &ManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/deviceHealthStatuses{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilder instantiates a new ManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilder and sets the default values.
func NewManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ManagedtenantsDevicehealthstatusesCountRequestBuilder when successful
func (m *ManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilder) Count()(*ManagedtenantsDevicehealthstatusesCountRequestBuilder) {
    return NewManagedtenantsDevicehealthstatusesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get deviceHealthStatuses from tenantRelationships
// returns a DeviceHealthStatusCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.DeviceHealthStatusCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateDeviceHealthStatusCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.DeviceHealthStatusCollectionResponseable), nil
}
// Post create new navigation property to deviceHealthStatuses for tenantRelationships
// returns a DeviceHealthStatusable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilder) Post(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.DeviceHealthStatusable, requestConfiguration *ManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilderPostRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.DeviceHealthStatusable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateDeviceHealthStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.DeviceHealthStatusable), nil
}
// ToGetRequestInformation get deviceHealthStatuses from tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to deviceHealthStatuses for tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilder) ToPostRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.DeviceHealthStatusable, requestConfiguration *ManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilder when successful
func (m *ManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilder) WithUrl(rawUrl string)(*ManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilder) {
    return NewManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
