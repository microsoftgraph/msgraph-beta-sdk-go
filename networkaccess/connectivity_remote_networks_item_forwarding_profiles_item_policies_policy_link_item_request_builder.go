package networkaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/networkaccess"
)

// ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilder provides operations to manage the policies property of the microsoft.graph.networkaccess.profile entity.
type ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilderGetQueryParameters traffic forwarding policies associated with this profile.
type ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilderGetQueryParameters
}
// ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilderInternal instantiates a new ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilder and sets the default values.
func NewConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilder) {
    m := &ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networkAccess/connectivity/remoteNetworks/{remoteNetwork%2Did}/forwardingProfiles/{forwardingProfile%2Did}/policies/{policyLink%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilder instantiates a new ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilder and sets the default values.
func NewConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property policies for networkAccess
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get traffic forwarding policies associated with this profile.
// returns a PolicyLinkable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilderGetRequestConfiguration)(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.PolicyLinkable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CreatePolicyLinkFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.PolicyLinkable), nil
}
// Patch update the navigation property policies in networkAccess
// returns a PolicyLinkable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilder) Patch(ctx context.Context, body i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.PolicyLinkable, requestConfiguration *ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilderPatchRequestConfiguration)(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.PolicyLinkable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CreatePolicyLinkFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.PolicyLinkable), nil
}
// Policy provides operations to manage the policy property of the microsoft.graph.networkaccess.policyLink entity.
// returns a *ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesItemPolicyRequestBuilder when successful
func (m *ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilder) Policy()(*ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesItemPolicyRequestBuilder) {
    return NewConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesItemPolicyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property policies for networkAccess
// returns a *RequestInformation when successful
func (m *ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/networkAccess/connectivity/remoteNetworks/{remoteNetwork%2Did}/forwardingProfiles/{forwardingProfile%2Did}/policies/{policyLink%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation traffic forwarding policies associated with this profile.
// returns a *RequestInformation when successful
func (m *ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property policies in networkAccess
// returns a *RequestInformation when successful
func (m *ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.PolicyLinkable, requestConfiguration *ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/networkAccess/connectivity/remoteNetworks/{remoteNetwork%2Did}/forwardingProfiles/{forwardingProfile%2Did}/policies/{policyLink%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilder when successful
func (m *ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilder) WithUrl(rawUrl string)(*ConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilder) {
    return NewConnectivityRemoteNetworksItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
