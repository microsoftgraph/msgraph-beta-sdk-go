package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilder provides operations to manage the conditionalAccessPolicyCoverages property of the microsoft.graph.managedTenants.managedTenant entity.
type ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilderGetQueryParameters read the properties and relationships of a conditionalAccessPolicyCoverage object.
type ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilderGetQueryParameters
}
// ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilderInternal instantiates a new ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilder and sets the default values.
func NewManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilder) {
    m := &ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/conditionalAccessPolicyCoverages/{conditionalAccessPolicyCoverage%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilder instantiates a new ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilder and sets the default values.
func NewManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property conditionalAccessPolicyCoverages for tenantRelationships
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a conditionalAccessPolicyCoverage object.
// returns a ConditionalAccessPolicyCoverageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/managedtenants-conditionalaccesspolicycoverage-get?view=graph-rest-beta
func (m *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ConditionalAccessPolicyCoverageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property conditionalAccessPolicyCoverages in tenantRelationships
// returns a ConditionalAccessPolicyCoverageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilder) Patch(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ConditionalAccessPolicyCoverageable, requestConfiguration *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilderPatchRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ConditionalAccessPolicyCoverageable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property conditionalAccessPolicyCoverages for tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a conditionalAccessPolicyCoverage object.
// returns a *RequestInformation when successful
func (m *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property conditionalAccessPolicyCoverages in tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ConditionalAccessPolicyCoverageable, requestConfiguration *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilder when successful
func (m *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilder) WithUrl(rawUrl string)(*ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilder) {
    return NewManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoverageItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
