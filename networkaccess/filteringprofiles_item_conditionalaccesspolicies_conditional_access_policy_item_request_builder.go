package networkaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/networkaccess"
)

// FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder provides operations to manage the conditionalAccessPolicies property of the microsoft.graph.networkaccess.filteringProfile entity.
type FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderGetQueryParameters a set of associated policies defined to regulate access to resources or systems based on specific conditions. Automatically expanded.
type FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderGetQueryParameters
}
// NewFilteringprofilesItemConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderInternal instantiates a new FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder and sets the default values.
func NewFilteringprofilesItemConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder) {
    m := &FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networkAccess/filteringProfiles/{filteringProfile%2Did}/conditionalAccessPolicies/{conditionalAccessPolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFilteringprofilesItemConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder instantiates a new FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder and sets the default values.
func NewFilteringprofilesItemConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilteringprofilesItemConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get a set of associated policies defined to regulate access to resources or systems based on specific conditions. Automatically expanded.
// returns a ConditionalAccessPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderGetRequestConfiguration)(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.ConditionalAccessPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CreateConditionalAccessPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.ConditionalAccessPolicyable), nil
}
// ToGetRequestInformation a set of associated policies defined to regulate access to resources or systems based on specific conditions. Automatically expanded.
// returns a *RequestInformation when successful
func (m *FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder when successful
func (m *FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder) WithUrl(rawUrl string)(*FilteringprofilesItemConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder) {
    return NewFilteringprofilesItemConditionalaccesspoliciesConditionalAccessPolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
