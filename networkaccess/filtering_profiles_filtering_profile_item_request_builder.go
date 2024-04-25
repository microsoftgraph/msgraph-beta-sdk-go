package networkaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/networkaccess"
)

// FilteringProfilesFilteringProfileItemRequestBuilder provides operations to manage the filteringProfiles property of the microsoft.graph.networkaccess.networkAccessRoot entity.
type FilteringProfilesFilteringProfileItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilteringProfilesFilteringProfileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilteringProfilesFilteringProfileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FilteringProfilesFilteringProfileItemRequestBuilderGetQueryParameters a filtering profile associates network access policies with Microsoft Entra ID Conditional Access policies, so that access policies can be applied to users and groups.
type FilteringProfilesFilteringProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FilteringProfilesFilteringProfileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilteringProfilesFilteringProfileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilteringProfilesFilteringProfileItemRequestBuilderGetQueryParameters
}
// FilteringProfilesFilteringProfileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilteringProfilesFilteringProfileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ConditionalAccessPolicies provides operations to manage the conditionalAccessPolicies property of the microsoft.graph.networkaccess.filteringProfile entity.
// returns a *FilteringProfilesItemConditionalAccessPoliciesRequestBuilder when successful
func (m *FilteringProfilesFilteringProfileItemRequestBuilder) ConditionalAccessPolicies()(*FilteringProfilesItemConditionalAccessPoliciesRequestBuilder) {
    return NewFilteringProfilesItemConditionalAccessPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewFilteringProfilesFilteringProfileItemRequestBuilderInternal instantiates a new FilteringProfilesFilteringProfileItemRequestBuilder and sets the default values.
func NewFilteringProfilesFilteringProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilteringProfilesFilteringProfileItemRequestBuilder) {
    m := &FilteringProfilesFilteringProfileItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networkAccess/filteringProfiles/{filteringProfile%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFilteringProfilesFilteringProfileItemRequestBuilder instantiates a new FilteringProfilesFilteringProfileItemRequestBuilder and sets the default values.
func NewFilteringProfilesFilteringProfileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilteringProfilesFilteringProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilteringProfilesFilteringProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property filteringProfiles for networkAccess
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilteringProfilesFilteringProfileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *FilteringProfilesFilteringProfileItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a filtering profile associates network access policies with Microsoft Entra ID Conditional Access policies, so that access policies can be applied to users and groups.
// returns a FilteringProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilteringProfilesFilteringProfileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FilteringProfilesFilteringProfileItemRequestBuilderGetRequestConfiguration)(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.FilteringProfileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CreateFilteringProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.FilteringProfileable), nil
}
// Patch update the navigation property filteringProfiles in networkAccess
// returns a FilteringProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilteringProfilesFilteringProfileItemRequestBuilder) Patch(ctx context.Context, body i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.FilteringProfileable, requestConfiguration *FilteringProfilesFilteringProfileItemRequestBuilderPatchRequestConfiguration)(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.FilteringProfileable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CreateFilteringProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.FilteringProfileable), nil
}
// Policies provides operations to manage the policies property of the microsoft.graph.networkaccess.profile entity.
// returns a *FilteringProfilesItemPoliciesRequestBuilder when successful
func (m *FilteringProfilesFilteringProfileItemRequestBuilder) Policies()(*FilteringProfilesItemPoliciesRequestBuilder) {
    return NewFilteringProfilesItemPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property filteringProfiles for networkAccess
// returns a *RequestInformation when successful
func (m *FilteringProfilesFilteringProfileItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FilteringProfilesFilteringProfileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a filtering profile associates network access policies with Microsoft Entra ID Conditional Access policies, so that access policies can be applied to users and groups.
// returns a *RequestInformation when successful
func (m *FilteringProfilesFilteringProfileItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilteringProfilesFilteringProfileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property filteringProfiles in networkAccess
// returns a *RequestInformation when successful
func (m *FilteringProfilesFilteringProfileItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.FilteringProfileable, requestConfiguration *FilteringProfilesFilteringProfileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilteringProfilesFilteringProfileItemRequestBuilder when successful
func (m *FilteringProfilesFilteringProfileItemRequestBuilder) WithUrl(rawUrl string)(*FilteringProfilesFilteringProfileItemRequestBuilder) {
    return NewFilteringProfilesFilteringProfileItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
