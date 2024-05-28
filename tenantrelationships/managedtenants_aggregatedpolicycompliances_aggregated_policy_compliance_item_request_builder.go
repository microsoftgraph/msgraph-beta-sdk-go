package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilder provides operations to manage the aggregatedPolicyCompliances property of the microsoft.graph.managedTenants.managedTenant entity.
type ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilderGetQueryParameters aggregate view of device compliance policies across managed tenants.
type ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilderGetQueryParameters
}
// ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilderInternal instantiates a new ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilder and sets the default values.
func NewManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilder) {
    m := &ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/aggregatedPolicyCompliances/{aggregatedPolicyCompliance%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilder instantiates a new ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilder and sets the default values.
func NewManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property aggregatedPolicyCompliances for tenantRelationships
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get aggregate view of device compliance policies across managed tenants.
// returns a AggregatedPolicyComplianceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.AggregatedPolicyComplianceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateAggregatedPolicyComplianceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.AggregatedPolicyComplianceable), nil
}
// Patch update the navigation property aggregatedPolicyCompliances in tenantRelationships
// returns a AggregatedPolicyComplianceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilder) Patch(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.AggregatedPolicyComplianceable, requestConfiguration *ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilderPatchRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.AggregatedPolicyComplianceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateAggregatedPolicyComplianceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.AggregatedPolicyComplianceable), nil
}
// ToDeleteRequestInformation delete navigation property aggregatedPolicyCompliances for tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation aggregate view of device compliance policies across managed tenants.
// returns a *RequestInformation when successful
func (m *ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property aggregatedPolicyCompliances in tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.AggregatedPolicyComplianceable, requestConfiguration *ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilder when successful
func (m *ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilder) WithUrl(rawUrl string)(*ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilder) {
    return NewManagedtenantsAggregatedpolicycompliancesAggregatedPolicyComplianceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
