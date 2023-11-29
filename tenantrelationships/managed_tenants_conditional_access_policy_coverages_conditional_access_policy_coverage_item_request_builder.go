package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilder provides operations to manage the conditionalAccessPolicyCoverages property of the microsoft.graph.managedTenants.managedTenant entity.
type ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilderGetQueryParameters read the properties and relationships of a conditionalAccessPolicyCoverage object.
type ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilderGetQueryParameters
}
// ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilderInternal instantiates a new ConditionalAccessPolicyCoverageItemRequestBuilder and sets the default values.
func NewManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilder) {
    m := &ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/conditionalAccessPolicyCoverages/{conditionalAccessPolicyCoverage%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilder instantiates a new ConditionalAccessPolicyCoverageItemRequestBuilder and sets the default values.
func NewManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property conditionalAccessPolicyCoverages for tenantRelationships
func (m *ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read the properties and relationships of a conditionalAccessPolicyCoverage object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/managedtenants-conditionalaccesspolicycoverage-get?view=graph-rest-1.0
func (m *ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ConditionalAccessPolicyCoverageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilder) Patch(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ConditionalAccessPolicyCoverageable, requestConfiguration *ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilderPatchRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ConditionalAccessPolicyCoverageable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a conditionalAccessPolicyCoverage object.
func (m *ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ConditionalAccessPolicyCoverageable, requestConfiguration *ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilder) WithUrl(rawUrl string)(*ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilder) {
    return NewManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
