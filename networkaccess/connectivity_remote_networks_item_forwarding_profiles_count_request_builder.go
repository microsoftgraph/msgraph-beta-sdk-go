package networkaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ConnectivityRemoteNetworksItemForwardingProfilesCountRequestBuilder provides operations to count the resources in the collection.
type ConnectivityRemoteNetworksItemForwardingProfilesCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConnectivityRemoteNetworksItemForwardingProfilesCountRequestBuilderGetQueryParameters get the number of the resource
type ConnectivityRemoteNetworksItemForwardingProfilesCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// ConnectivityRemoteNetworksItemForwardingProfilesCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectivityRemoteNetworksItemForwardingProfilesCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConnectivityRemoteNetworksItemForwardingProfilesCountRequestBuilderGetQueryParameters
}
// NewConnectivityRemoteNetworksItemForwardingProfilesCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewConnectivityRemoteNetworksItemForwardingProfilesCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectivityRemoteNetworksItemForwardingProfilesCountRequestBuilder) {
    m := &ConnectivityRemoteNetworksItemForwardingProfilesCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networkAccess/connectivity/remoteNetworks/{remoteNetwork%2Did}/forwardingProfiles/$count{?%24search,%24filter}", pathParameters),
    }
    return m
}
// NewConnectivityRemoteNetworksItemForwardingProfilesCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewConnectivityRemoteNetworksItemForwardingProfilesCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectivityRemoteNetworksItemForwardingProfilesCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConnectivityRemoteNetworksItemForwardingProfilesCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
func (m *ConnectivityRemoteNetworksItemForwardingProfilesCountRequestBuilder) Get(ctx context.Context, requestConfiguration *ConnectivityRemoteNetworksItemForwardingProfilesCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
func (m *ConnectivityRemoteNetworksItemForwardingProfilesCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConnectivityRemoteNetworksItemForwardingProfilesCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ConnectivityRemoteNetworksItemForwardingProfilesCountRequestBuilder) WithUrl(rawUrl string)(*ConnectivityRemoteNetworksItemForwardingProfilesCountRequestBuilder) {
    return NewConnectivityRemoteNetworksItemForwardingProfilesCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
