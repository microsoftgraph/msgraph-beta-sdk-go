package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilder provides operations to manage the groupLifecyclePolicies property of the microsoft.graph.group entity.
type ItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilderGetQueryParameters retrieves a list of groupLifecyclePolicy objects to which a group belongs.
type ItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilderGetQueryParameters struct {
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
// ItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilderGetQueryParameters
}
// ItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByGroupLifecyclePolicyId provides operations to manage the groupLifecyclePolicies property of the microsoft.graph.group entity.
// returns a *ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilder when successful
func (m *ItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilder) ByGroupLifecyclePolicyId(groupLifecyclePolicyId string)(*ItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if groupLifecyclePolicyId != "" {
        urlTplParams["groupLifecyclePolicy%2Did"] = groupLifecyclePolicyId
    }
    return NewItemGrouplifecyclepoliciesGroupLifecyclePolicyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilderInternal instantiates a new ItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilder and sets the default values.
func NewItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilder) {
    m := &ItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/groupLifecyclePolicies{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilder instantiates a new ItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilder and sets the default values.
func NewItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemGrouplifecyclepoliciesCountRequestBuilder when successful
func (m *ItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilder) Count()(*ItemGrouplifecyclepoliciesCountRequestBuilder) {
    return NewItemGrouplifecyclepoliciesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieves a list of groupLifecyclePolicy objects to which a group belongs.
// returns a GroupLifecyclePolicyCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/group-list-grouplifecyclepolicies?view=graph-rest-beta
func (m *ItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupLifecyclePolicyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupLifecyclePolicyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupLifecyclePolicyCollectionResponseable), nil
}
// Post create new navigation property to groupLifecyclePolicies for groups
// returns a GroupLifecyclePolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupLifecyclePolicyable, requestConfiguration *ItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupLifecyclePolicyable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupLifecyclePolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupLifecyclePolicyable), nil
}
// RenewGroup provides operations to call the renewGroup method.
// returns a *ItemGrouplifecyclepoliciesRenewgroupRenewGroupRequestBuilder when successful
func (m *ItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilder) RenewGroup()(*ItemGrouplifecyclepoliciesRenewgroupRenewGroupRequestBuilder) {
    return NewItemGrouplifecyclepoliciesRenewgroupRenewGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation retrieves a list of groupLifecyclePolicy objects to which a group belongs.
// returns a *RequestInformation when successful
func (m *ItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to groupLifecyclePolicies for groups
// returns a *RequestInformation when successful
func (m *ItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupLifecyclePolicyable, requestConfiguration *ItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilder when successful
func (m *ItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilder) WithUrl(rawUrl string)(*ItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilder) {
    return NewItemGrouplifecyclepoliciesGroupLifecyclePoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
