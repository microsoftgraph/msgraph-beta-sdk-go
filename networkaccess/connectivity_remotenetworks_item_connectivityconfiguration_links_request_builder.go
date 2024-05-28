package networkaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/networkaccess"
)

// ConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilder provides operations to manage the links property of the microsoft.graph.networkaccess.remoteNetworkConnectivityConfiguration entity.
type ConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilderGetQueryParameters list of connectivity configurations for deviceLink objects.
type ConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilderGetQueryParameters struct {
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
// ConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilderGetQueryParameters
}
// ConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByConnectivityConfigurationLinkId provides operations to manage the links property of the microsoft.graph.networkaccess.remoteNetworkConnectivityConfiguration entity.
// returns a *ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilder when successful
func (m *ConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilder) ByConnectivityConfigurationLinkId(connectivityConfigurationLinkId string)(*ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if connectivityConfigurationLinkId != "" {
        urlTplParams["connectivityConfigurationLink%2Did"] = connectivityConfigurationLinkId
    }
    return NewConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilderInternal instantiates a new ConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilder and sets the default values.
func NewConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilder) {
    m := &ConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networkAccess/connectivity/remoteNetworks/{remoteNetwork%2Did}/connectivityConfiguration/links{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilder instantiates a new ConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilder and sets the default values.
func NewConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ConnectivityRemotenetworksItemConnectivityconfigurationLinksCountRequestBuilder when successful
func (m *ConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilder) Count()(*ConnectivityRemotenetworksItemConnectivityconfigurationLinksCountRequestBuilder) {
    return NewConnectivityRemotenetworksItemConnectivityconfigurationLinksCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list of connectivity configurations for deviceLink objects.
// returns a ConnectivityConfigurationLinkCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilder) Get(ctx context.Context, requestConfiguration *ConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilderGetRequestConfiguration)(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.ConnectivityConfigurationLinkCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CreateConnectivityConfigurationLinkCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.ConnectivityConfigurationLinkCollectionResponseable), nil
}
// Post create new navigation property to links for networkAccess
// returns a ConnectivityConfigurationLinkable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilder) Post(ctx context.Context, body i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.ConnectivityConfigurationLinkable, requestConfiguration *ConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilderPostRequestConfiguration)(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.ConnectivityConfigurationLinkable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CreateConnectivityConfigurationLinkFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.ConnectivityConfigurationLinkable), nil
}
// ToGetRequestInformation list of connectivity configurations for deviceLink objects.
// returns a *RequestInformation when successful
func (m *ConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to links for networkAccess
// returns a *RequestInformation when successful
func (m *ConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilder) ToPostRequestInformation(ctx context.Context, body i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.ConnectivityConfigurationLinkable, requestConfiguration *ConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilder when successful
func (m *ConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilder) WithUrl(rawUrl string)(*ConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilder) {
    return NewConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
