package networkaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/networkaccess"
)

// ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilder provides operations to manage the links property of the microsoft.graph.networkaccess.remoteNetworkConnectivityConfiguration entity.
type ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilderGetQueryParameters list of connectivity configurations for deviceLink objects.
type ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilderGetQueryParameters
}
// ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilderInternal instantiates a new ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilder and sets the default values.
func NewConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilder) {
    m := &ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networkAccess/connectivity/remoteNetworks/{remoteNetwork%2Did}/connectivityConfiguration/links/{connectivityConfigurationLink%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilder instantiates a new ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilder and sets the default values.
func NewConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property links for networkAccess
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get list of connectivity configurations for deviceLink objects.
// returns a ConnectivityConfigurationLinkable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilderGetRequestConfiguration)(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.ConnectivityConfigurationLinkable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property links in networkAccess
// returns a ConnectivityConfigurationLinkable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilder) Patch(ctx context.Context, body i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.ConnectivityConfigurationLinkable, requestConfiguration *ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilderPatchRequestConfiguration)(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.ConnectivityConfigurationLinkable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property links for networkAccess
// returns a *RequestInformation when successful
func (m *ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation list of connectivity configurations for deviceLink objects.
// returns a *RequestInformation when successful
func (m *ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property links in networkAccess
// returns a *RequestInformation when successful
func (m *ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.ConnectivityConfigurationLinkable, requestConfiguration *ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilder when successful
func (m *ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilder) WithUrl(rawUrl string)(*ConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilder) {
    return NewConnectivityRemotenetworksItemConnectivityconfigurationLinksConnectivityConfigurationLinkItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
