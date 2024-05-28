package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DomainjoinconnectorsDomainJoinConnectorsRequestBuilder provides operations to manage the domainJoinConnectors property of the microsoft.graph.deviceManagement entity.
type DomainjoinconnectorsDomainJoinConnectorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DomainjoinconnectorsDomainJoinConnectorsRequestBuilderGetQueryParameters a list of connector objects.
type DomainjoinconnectorsDomainJoinConnectorsRequestBuilderGetQueryParameters struct {
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
// DomainjoinconnectorsDomainJoinConnectorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DomainjoinconnectorsDomainJoinConnectorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DomainjoinconnectorsDomainJoinConnectorsRequestBuilderGetQueryParameters
}
// DomainjoinconnectorsDomainJoinConnectorsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DomainjoinconnectorsDomainJoinConnectorsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceManagementDomainJoinConnectorId provides operations to manage the domainJoinConnectors property of the microsoft.graph.deviceManagement entity.
// returns a *DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilder when successful
func (m *DomainjoinconnectorsDomainJoinConnectorsRequestBuilder) ByDeviceManagementDomainJoinConnectorId(deviceManagementDomainJoinConnectorId string)(*DomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceManagementDomainJoinConnectorId != "" {
        urlTplParams["deviceManagementDomainJoinConnector%2Did"] = deviceManagementDomainJoinConnectorId
    }
    return NewDomainjoinconnectorsDeviceManagementDomainJoinConnectorItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDomainjoinconnectorsDomainJoinConnectorsRequestBuilderInternal instantiates a new DomainjoinconnectorsDomainJoinConnectorsRequestBuilder and sets the default values.
func NewDomainjoinconnectorsDomainJoinConnectorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DomainjoinconnectorsDomainJoinConnectorsRequestBuilder) {
    m := &DomainjoinconnectorsDomainJoinConnectorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/domainJoinConnectors{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDomainjoinconnectorsDomainJoinConnectorsRequestBuilder instantiates a new DomainjoinconnectorsDomainJoinConnectorsRequestBuilder and sets the default values.
func NewDomainjoinconnectorsDomainJoinConnectorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DomainjoinconnectorsDomainJoinConnectorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDomainjoinconnectorsDomainJoinConnectorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DomainjoinconnectorsCountRequestBuilder when successful
func (m *DomainjoinconnectorsDomainJoinConnectorsRequestBuilder) Count()(*DomainjoinconnectorsCountRequestBuilder) {
    return NewDomainjoinconnectorsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a list of connector objects.
// returns a DeviceManagementDomainJoinConnectorCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DomainjoinconnectorsDomainJoinConnectorsRequestBuilder) Get(ctx context.Context, requestConfiguration *DomainjoinconnectorsDomainJoinConnectorsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementDomainJoinConnectorCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementDomainJoinConnectorCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementDomainJoinConnectorCollectionResponseable), nil
}
// Post create new navigation property to domainJoinConnectors for deviceManagement
// returns a DeviceManagementDomainJoinConnectorable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DomainjoinconnectorsDomainJoinConnectorsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementDomainJoinConnectorable, requestConfiguration *DomainjoinconnectorsDomainJoinConnectorsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementDomainJoinConnectorable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementDomainJoinConnectorFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementDomainJoinConnectorable), nil
}
// ToGetRequestInformation a list of connector objects.
// returns a *RequestInformation when successful
func (m *DomainjoinconnectorsDomainJoinConnectorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DomainjoinconnectorsDomainJoinConnectorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to domainJoinConnectors for deviceManagement
// returns a *RequestInformation when successful
func (m *DomainjoinconnectorsDomainJoinConnectorsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementDomainJoinConnectorable, requestConfiguration *DomainjoinconnectorsDomainJoinConnectorsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DomainjoinconnectorsDomainJoinConnectorsRequestBuilder when successful
func (m *DomainjoinconnectorsDomainJoinConnectorsRequestBuilder) WithUrl(rawUrl string)(*DomainjoinconnectorsDomainJoinConnectorsRequestBuilder) {
    return NewDomainjoinconnectorsDomainJoinConnectorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
