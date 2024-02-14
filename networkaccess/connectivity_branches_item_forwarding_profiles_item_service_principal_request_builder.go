package networkaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ConnectivityBranchesItemForwardingProfilesItemServicePrincipalRequestBuilder provides operations to manage the servicePrincipal property of the microsoft.graph.networkaccess.forwardingProfile entity.
type ConnectivityBranchesItemForwardingProfilesItemServicePrincipalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConnectivityBranchesItemForwardingProfilesItemServicePrincipalRequestBuilderGetQueryParameters get servicePrincipal from networkAccess
type ConnectivityBranchesItemForwardingProfilesItemServicePrincipalRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ConnectivityBranchesItemForwardingProfilesItemServicePrincipalRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectivityBranchesItemForwardingProfilesItemServicePrincipalRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConnectivityBranchesItemForwardingProfilesItemServicePrincipalRequestBuilderGetQueryParameters
}
// NewConnectivityBranchesItemForwardingProfilesItemServicePrincipalRequestBuilderInternal instantiates a new ConnectivityBranchesItemForwardingProfilesItemServicePrincipalRequestBuilder and sets the default values.
func NewConnectivityBranchesItemForwardingProfilesItemServicePrincipalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectivityBranchesItemForwardingProfilesItemServicePrincipalRequestBuilder) {
    m := &ConnectivityBranchesItemForwardingProfilesItemServicePrincipalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networkAccess/connectivity/branches/{branchSite%2Did}/forwardingProfiles/{forwardingProfile%2Did}/servicePrincipal{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewConnectivityBranchesItemForwardingProfilesItemServicePrincipalRequestBuilder instantiates a new ConnectivityBranchesItemForwardingProfilesItemServicePrincipalRequestBuilder and sets the default values.
func NewConnectivityBranchesItemForwardingProfilesItemServicePrincipalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectivityBranchesItemForwardingProfilesItemServicePrincipalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConnectivityBranchesItemForwardingProfilesItemServicePrincipalRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get servicePrincipal from networkAccess
// Deprecated: The Branches API is deprecated and will stop returning data on January 20, 2024. Please use the new Remote Network API. as of 2022-06/PrivatePreview:NetworkAccess
// returns a ServicePrincipalable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConnectivityBranchesItemForwardingProfilesItemServicePrincipalRequestBuilder) Get(ctx context.Context, requestConfiguration *ConnectivityBranchesItemForwardingProfilesItemServicePrincipalRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable, error) {
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
// Deprecated: The Branches API is deprecated and will stop returning data on January 20, 2024. Please use the new Remote Network API. as of 2022-06/PrivatePreview:NetworkAccess
// returns a *RequestInformation when successful
func (m *ConnectivityBranchesItemForwardingProfilesItemServicePrincipalRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConnectivityBranchesItemForwardingProfilesItemServicePrincipalRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The Branches API is deprecated and will stop returning data on January 20, 2024. Please use the new Remote Network API. as of 2022-06/PrivatePreview:NetworkAccess
// returns a *ConnectivityBranchesItemForwardingProfilesItemServicePrincipalRequestBuilder when successful
func (m *ConnectivityBranchesItemForwardingProfilesItemServicePrincipalRequestBuilder) WithUrl(rawUrl string)(*ConnectivityBranchesItemForwardingProfilesItemServicePrincipalRequestBuilder) {
    return NewConnectivityBranchesItemForwardingProfilesItemServicePrincipalRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
