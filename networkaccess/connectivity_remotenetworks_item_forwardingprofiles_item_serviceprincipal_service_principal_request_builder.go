package networkaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ConnectivityRemotenetworksItemForwardingprofilesItemServiceprincipalServicePrincipalRequestBuilder provides operations to manage the servicePrincipal property of the microsoft.graph.networkaccess.forwardingProfile entity.
type ConnectivityRemotenetworksItemForwardingprofilesItemServiceprincipalServicePrincipalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConnectivityRemotenetworksItemForwardingprofilesItemServiceprincipalServicePrincipalRequestBuilderGetQueryParameters get servicePrincipal from networkAccess
type ConnectivityRemotenetworksItemForwardingprofilesItemServiceprincipalServicePrincipalRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ConnectivityRemotenetworksItemForwardingprofilesItemServiceprincipalServicePrincipalRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectivityRemotenetworksItemForwardingprofilesItemServiceprincipalServicePrincipalRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConnectivityRemotenetworksItemForwardingprofilesItemServiceprincipalServicePrincipalRequestBuilderGetQueryParameters
}
// NewConnectivityRemotenetworksItemForwardingprofilesItemServiceprincipalServicePrincipalRequestBuilderInternal instantiates a new ConnectivityRemotenetworksItemForwardingprofilesItemServiceprincipalServicePrincipalRequestBuilder and sets the default values.
func NewConnectivityRemotenetworksItemForwardingprofilesItemServiceprincipalServicePrincipalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectivityRemotenetworksItemForwardingprofilesItemServiceprincipalServicePrincipalRequestBuilder) {
    m := &ConnectivityRemotenetworksItemForwardingprofilesItemServiceprincipalServicePrincipalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networkAccess/connectivity/remoteNetworks/{remoteNetwork%2Did}/forwardingProfiles/{forwardingProfile%2Did}/servicePrincipal{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewConnectivityRemotenetworksItemForwardingprofilesItemServiceprincipalServicePrincipalRequestBuilder instantiates a new ConnectivityRemotenetworksItemForwardingprofilesItemServiceprincipalServicePrincipalRequestBuilder and sets the default values.
func NewConnectivityRemotenetworksItemForwardingprofilesItemServiceprincipalServicePrincipalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectivityRemotenetworksItemForwardingprofilesItemServiceprincipalServicePrincipalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConnectivityRemotenetworksItemForwardingprofilesItemServiceprincipalServicePrincipalRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get servicePrincipal from networkAccess
// returns a ServicePrincipalable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConnectivityRemotenetworksItemForwardingprofilesItemServiceprincipalServicePrincipalRequestBuilder) Get(ctx context.Context, requestConfiguration *ConnectivityRemotenetworksItemForwardingprofilesItemServiceprincipalServicePrincipalRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServicePrincipalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable), nil
}
// ToGetRequestInformation get servicePrincipal from networkAccess
// returns a *RequestInformation when successful
func (m *ConnectivityRemotenetworksItemForwardingprofilesItemServiceprincipalServicePrincipalRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConnectivityRemotenetworksItemForwardingprofilesItemServiceprincipalServicePrincipalRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ConnectivityRemotenetworksItemForwardingprofilesItemServiceprincipalServicePrincipalRequestBuilder when successful
func (m *ConnectivityRemotenetworksItemForwardingprofilesItemServiceprincipalServicePrincipalRequestBuilder) WithUrl(rawUrl string)(*ConnectivityRemotenetworksItemForwardingprofilesItemServiceprincipalServicePrincipalRequestBuilder) {
    return NewConnectivityRemotenetworksItemForwardingprofilesItemServiceprincipalServicePrincipalRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
