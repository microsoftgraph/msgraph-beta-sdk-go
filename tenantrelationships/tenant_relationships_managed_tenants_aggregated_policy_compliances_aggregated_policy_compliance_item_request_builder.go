package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// TenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilder provides operations to manage the aggregatedPolicyCompliances property of the microsoft.graph.managedTenants.managedTenant entity.
type TenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilderGetQueryParameters aggregate view of device compliance policies across managed tenants.
type TenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilderGetQueryParameters
}
// TenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilderInternal instantiates a new AggregatedPolicyComplianceItemRequestBuilder and sets the default values.
func NewTenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilder) {
    m := &TenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/aggregatedPolicyCompliances/{aggregatedPolicyCompliance%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilder instantiates a new AggregatedPolicyComplianceItemRequestBuilder and sets the default values.
func NewTenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property aggregatedPolicyCompliances for tenantRelationships
func (m *TenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation aggregate view of device compliance policies across managed tenants.
func (m *TenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property aggregatedPolicyCompliances in tenantRelationships
func (m *TenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.AggregatedPolicyComplianceable, requestConfiguration *TenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property aggregatedPolicyCompliances for tenantRelationships
func (m *TenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get aggregate view of device compliance policies across managed tenants.
func (m *TenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.AggregatedPolicyComplianceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateAggregatedPolicyComplianceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.AggregatedPolicyComplianceable), nil
}
// Patch update the navigation property aggregatedPolicyCompliances in tenantRelationships
func (m *TenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilder) Patch(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.AggregatedPolicyComplianceable, requestConfiguration *TenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilderPatchRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.AggregatedPolicyComplianceable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateAggregatedPolicyComplianceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.AggregatedPolicyComplianceable), nil
}
