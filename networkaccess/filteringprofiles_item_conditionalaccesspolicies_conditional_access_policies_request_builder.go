package networkaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/networkaccess"
)

// FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder provides operations to manage the conditionalAccessPolicies property of the microsoft.graph.networkaccess.filteringProfile entity.
type FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilderGetQueryParameters a set of associated policies defined to regulate access to resources or systems based on specific conditions. Automatically expanded.
type FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilderGetQueryParameters struct {
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
// FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilderGetQueryParameters
}
// ByConditionalAccessPolicyId provides operations to manage the conditionalAccessPolicies property of the microsoft.graph.networkaccess.filteringProfile entity.
// returns a *FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder when successful
func (m *FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder) ByConditionalAccessPolicyId(conditionalAccessPolicyId string)(*FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if conditionalAccessPolicyId != "" {
        urlTplParams["conditionalAccessPolicy%2Did"] = conditionalAccessPolicyId
    }
    return NewFilteringprofilesItemConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewFilteringprofilesItemConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilderInternal instantiates a new FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder and sets the default values.
func NewFilteringprofilesItemConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder) {
    m := &FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networkAccess/filteringProfiles/{filteringProfile%2Did}/conditionalAccessPolicies{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewFilteringprofilesItemConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder instantiates a new FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder and sets the default values.
func NewFilteringprofilesItemConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilteringprofilesItemConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *FilteringprofilesItemConditionalaccesspoliciesCountRequestBuilder when successful
func (m *FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder) Count()(*FilteringprofilesItemConditionalaccesspoliciesCountRequestBuilder) {
    return NewFilteringprofilesItemConditionalaccesspoliciesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a set of associated policies defined to regulate access to resources or systems based on specific conditions. Automatically expanded.
// returns a ConditionalAccessPolicyCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilderGetRequestConfiguration)(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.ConditionalAccessPolicyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CreateConditionalAccessPolicyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.ConditionalAccessPolicyCollectionResponseable), nil
}
// ToGetRequestInformation a set of associated policies defined to regulate access to resources or systems based on specific conditions. Automatically expanded.
// returns a *RequestInformation when successful
func (m *FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder when successful
func (m *FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder) WithUrl(rawUrl string)(*FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder) {
    return NewFilteringprofilesItemConditionalaccesspoliciesConditionalAccessPoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
