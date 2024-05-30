package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilder provides operations to manage the conditionalAccessPolicyCoverages property of the microsoft.graph.managedTenants.managedTenant entity.
type ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilderGetQueryParameters get a list of the conditionalAccessPolicyCoverage objects and their properties. Use this operation to list Microsoft Entra Conditional Access policy coverage for all tenants that are being managed by the multi-tenant management platform.
type ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilderGetQueryParameters struct {
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
// ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilderGetQueryParameters
}
// ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByConditionalAccessPolicyCoverageId provides operations to manage the conditionalAccessPolicyCoverages property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilder when successful
func (m *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilder) ByConditionalAccessPolicyCoverageId(conditionalAccessPolicyCoverageId string)(*ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if conditionalAccessPolicyCoverageId != "" {
        urlTplParams["conditionalAccessPolicyCoverage%2Did"] = conditionalAccessPolicyCoverageId
    }
    return NewManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilderInternal instantiates a new ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilder and sets the default values.
func NewManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilder) {
    m := &ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/conditionalAccessPolicyCoverages{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilder instantiates a new ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilder and sets the default values.
func NewManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ManagedtenantsConditionalaccesspolicycoveragesCountRequestBuilder when successful
func (m *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilder) Count()(*ManagedtenantsConditionalaccesspolicycoveragesCountRequestBuilder) {
    return NewManagedtenantsConditionalaccesspolicycoveragesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the conditionalAccessPolicyCoverage objects and their properties. Use this operation to list Microsoft Entra Conditional Access policy coverage for all tenants that are being managed by the multi-tenant management platform.
// returns a ConditionalAccessPolicyCoverageCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/managedtenants-managedtenant-list-conditionalaccesspolicycoverages?view=graph-rest-beta
func (m *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ConditionalAccessPolicyCoverageCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateConditionalAccessPolicyCoverageCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ConditionalAccessPolicyCoverageCollectionResponseable), nil
}
// Post create new navigation property to conditionalAccessPolicyCoverages for tenantRelationships
// returns a ConditionalAccessPolicyCoverageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilder) Post(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ConditionalAccessPolicyCoverageable, requestConfiguration *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilderPostRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ConditionalAccessPolicyCoverageable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateConditionalAccessPolicyCoverageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ConditionalAccessPolicyCoverageable), nil
}
// ToGetRequestInformation get a list of the conditionalAccessPolicyCoverage objects and their properties. Use this operation to list Microsoft Entra Conditional Access policy coverage for all tenants that are being managed by the multi-tenant management platform.
// returns a *RequestInformation when successful
func (m *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to conditionalAccessPolicyCoverages for tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilder) ToPostRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ConditionalAccessPolicyCoverageable, requestConfiguration *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilder when successful
func (m *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilder) WithUrl(rawUrl string)(*ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilder) {
    return NewManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
