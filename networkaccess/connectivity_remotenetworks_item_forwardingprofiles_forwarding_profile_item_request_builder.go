package networkaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/networkaccess"
)

// ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilder provides operations to manage the forwardingProfiles property of the microsoft.graph.networkaccess.remoteNetwork entity.
type ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilderGetQueryParameters each forwarding profile associated with a remote network is specified. Supports $expand and $select.
type ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilderGetQueryParameters
}
// ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilderInternal instantiates a new ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilder and sets the default values.
func NewConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilder) {
    m := &ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networkAccess/connectivity/remoteNetworks/{remoteNetwork%2Did}/forwardingProfiles/{forwardingProfile%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilder instantiates a new ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilder and sets the default values.
func NewConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property forwardingProfiles for networkAccess
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get each forwarding profile associated with a remote network is specified. Supports $expand and $select.
// returns a ForwardingProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilderGetRequestConfiguration)(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.ForwardingProfileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CreateForwardingProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.ForwardingProfileable), nil
}
// Patch update the navigation property forwardingProfiles in networkAccess
// returns a ForwardingProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilder) Patch(ctx context.Context, body i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.ForwardingProfileable, requestConfiguration *ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilderPatchRequestConfiguration)(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.ForwardingProfileable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CreateForwardingProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.ForwardingProfileable), nil
}
// Policies provides operations to manage the policies property of the microsoft.graph.networkaccess.profile entity.
// returns a *ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesRequestBuilder when successful
func (m *ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilder) Policies()(*ConnectivityRemotenetworksItemForwardingprofilesItemPoliciesRequestBuilder) {
    return NewConnectivityRemotenetworksItemForwardingprofilesItemPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServicePrincipal provides operations to manage the servicePrincipal property of the microsoft.graph.networkaccess.forwardingProfile entity.
// returns a *ConnectivityRemotenetworksItemForwardingprofilesItemServiceprincipalServicePrincipalRequestBuilder when successful
func (m *ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilder) ServicePrincipal()(*ConnectivityRemotenetworksItemForwardingprofilesItemServiceprincipalServicePrincipalRequestBuilder) {
    return NewConnectivityRemotenetworksItemForwardingprofilesItemServiceprincipalServicePrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property forwardingProfiles for networkAccess
// returns a *RequestInformation when successful
func (m *ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation each forwarding profile associated with a remote network is specified. Supports $expand and $select.
// returns a *RequestInformation when successful
func (m *ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property forwardingProfiles in networkAccess
// returns a *RequestInformation when successful
func (m *ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.ForwardingProfileable, requestConfiguration *ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilder when successful
func (m *ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilder) WithUrl(rawUrl string)(*ConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilder) {
    return NewConnectivityRemotenetworksItemForwardingprofilesForwardingProfileItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
