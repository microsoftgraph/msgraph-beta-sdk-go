package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// TenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder provides operations to manage the tenantsCustomizedInformation property of the microsoft.graph.managedTenants.managedTenant entity.
type TenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderGetQueryParameters the collection of tenant level customized information across managed tenants.
type TenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderGetQueryParameters
}
// TenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderInternal instantiates a new TenantCustomizedInformationItemRequestBuilder and sets the default values.
func NewTenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder) {
    m := &TenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/tenantsCustomizedInformation/{tenantCustomizedInformation%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder instantiates a new TenantCustomizedInformationItemRequestBuilder and sets the default values.
func NewTenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property tenantsCustomizedInformation for tenantRelationships
func (m *TenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of tenant level customized information across managed tenants.
func (m *TenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property tenantsCustomizedInformation in tenantRelationships
func (m *TenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantCustomizedInformationable, requestConfiguration *TenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property tenantsCustomizedInformation for tenantRelationships
func (m *TenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the collection of tenant level customized information across managed tenants.
func (m *TenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantCustomizedInformationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateTenantCustomizedInformationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantCustomizedInformationable), nil
}
// Patch update the navigation property tenantsCustomizedInformation in tenantRelationships
func (m *TenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder) Patch(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantCustomizedInformationable, requestConfiguration *TenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderPatchRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantCustomizedInformationable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateTenantCustomizedInformationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantCustomizedInformationable), nil
}
