// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package networkaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/networkaccess"
)

// ConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilder provides operations to manage the policies property of the microsoft.graph.networkaccess.profile entity.
type ConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilderGetQueryParameters the traffic forwarding policies associated with this profile.
type ConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilderGetQueryParameters struct {
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
// ConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilderGetQueryParameters
}
// ConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByPolicyLinkId provides operations to manage the policies property of the microsoft.graph.networkaccess.profile entity.
// Deprecated: The Branches API is deprecated and will stop returning data on March 20, 2024. Please use the new Remote Network API. as of 2022-06/PrivatePreview:NetworkAccess on 2023-12-19 and will be removed 2024-03-20
// returns a *ConnectivityBranchesItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilder when successful
func (m *ConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilder) ByPolicyLinkId(policyLinkId string)(*ConnectivityBranchesItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if policyLinkId != "" {
        urlTplParams["policyLink%2Did"] = policyLinkId
    }
    return NewConnectivityBranchesItemForwardingProfilesItemPoliciesPolicyLinkItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilderInternal instantiates a new ConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilder and sets the default values.
func NewConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilder) {
    m := &ConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networkAccess/connectivity/branches/{branchSite%2Did}/forwardingProfiles/{forwardingProfile%2Did}/policies{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilder instantiates a new ConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilder and sets the default values.
func NewConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ConnectivityBranchesItemForwardingProfilesItemPoliciesCountRequestBuilder when successful
func (m *ConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilder) Count()(*ConnectivityBranchesItemForwardingProfilesItemPoliciesCountRequestBuilder) {
    return NewConnectivityBranchesItemForwardingProfilesItemPoliciesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the traffic forwarding policies associated with this profile.
// Deprecated: The Branches API is deprecated and will stop returning data on March 20, 2024. Please use the new Remote Network API. as of 2022-06/PrivatePreview:NetworkAccess on 2023-12-19 and will be removed 2024-03-20
// returns a PolicyLinkCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *ConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilderGetRequestConfiguration)(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.PolicyLinkCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CreatePolicyLinkCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.PolicyLinkCollectionResponseable), nil
}
// Post create new navigation property to policies for networkAccess
// Deprecated: The Branches API is deprecated and will stop returning data on March 20, 2024. Please use the new Remote Network API. as of 2022-06/PrivatePreview:NetworkAccess on 2023-12-19 and will be removed 2024-03-20
// returns a PolicyLinkable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilder) Post(ctx context.Context, body i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.PolicyLinkable, requestConfiguration *ConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilderPostRequestConfiguration)(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.PolicyLinkable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation the traffic forwarding policies associated with this profile.
// Deprecated: The Branches API is deprecated and will stop returning data on March 20, 2024. Please use the new Remote Network API. as of 2022-06/PrivatePreview:NetworkAccess on 2023-12-19 and will be removed 2024-03-20
// returns a *RequestInformation when successful
func (m *ConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to policies for networkAccess
// Deprecated: The Branches API is deprecated and will stop returning data on March 20, 2024. Please use the new Remote Network API. as of 2022-06/PrivatePreview:NetworkAccess on 2023-12-19 and will be removed 2024-03-20
// returns a *RequestInformation when successful
func (m *ConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilder) ToPostRequestInformation(ctx context.Context, body i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.PolicyLinkable, requestConfiguration *ConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The Branches API is deprecated and will stop returning data on March 20, 2024. Please use the new Remote Network API. as of 2022-06/PrivatePreview:NetworkAccess on 2023-12-19 and will be removed 2024-03-20
// returns a *ConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilder when successful
func (m *ConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilder) WithUrl(rawUrl string)(*ConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilder) {
    return NewConnectivityBranchesItemForwardingProfilesItemPoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
