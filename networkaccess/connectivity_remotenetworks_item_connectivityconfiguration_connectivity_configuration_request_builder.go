package networkaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/networkaccess"
)

// ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilder provides operations to manage the connectivityConfiguration property of the microsoft.graph.networkaccess.remoteNetwork entity.
type ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilderGetQueryParameters specifies the connectivity details of all device links associated with a remote network.
type ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilderGetQueryParameters
}
// ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilderInternal instantiates a new ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilder and sets the default values.
func NewConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilder) {
    m := &ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networkAccess/connectivity/remoteNetworks/{remoteNetwork%2Did}/connectivityConfiguration{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilder instantiates a new ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilder and sets the default values.
func NewConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property connectivityConfiguration for networkAccess
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilder) Delete(ctx context.Context, requestConfiguration *ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get specifies the connectivity details of all device links associated with a remote network.
// returns a RemoteNetworkConnectivityConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilder) Get(ctx context.Context, requestConfiguration *ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilderGetRequestConfiguration)(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.RemoteNetworkConnectivityConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CreateRemoteNetworkConnectivityConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.RemoteNetworkConnectivityConfigurationable), nil
}
// Links provides operations to manage the links property of the microsoft.graph.networkaccess.remoteNetworkConnectivityConfiguration entity.
// returns a *ConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilder when successful
func (m *ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilder) Links()(*ConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilder) {
    return NewConnectivityRemotenetworksItemConnectivityconfigurationLinksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property connectivityConfiguration in networkAccess
// returns a RemoteNetworkConnectivityConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilder) Patch(ctx context.Context, body i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.RemoteNetworkConnectivityConfigurationable, requestConfiguration *ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilderPatchRequestConfiguration)(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.RemoteNetworkConnectivityConfigurationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CreateRemoteNetworkConnectivityConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.RemoteNetworkConnectivityConfigurationable), nil
}
// ToDeleteRequestInformation delete navigation property connectivityConfiguration for networkAccess
// returns a *RequestInformation when successful
func (m *ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation specifies the connectivity details of all device links associated with a remote network.
// returns a *RequestInformation when successful
func (m *ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property connectivityConfiguration in networkAccess
// returns a *RequestInformation when successful
func (m *ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.RemoteNetworkConnectivityConfigurationable, requestConfiguration *ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilder when successful
func (m *ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilder) WithUrl(rawUrl string)(*ConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilder) {
    return NewConnectivityRemotenetworksItemConnectivityconfigurationConnectivityConfigurationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
