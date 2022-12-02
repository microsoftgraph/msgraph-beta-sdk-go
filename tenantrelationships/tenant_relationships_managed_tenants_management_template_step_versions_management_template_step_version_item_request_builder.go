package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder provides operations to manage the managementTemplateStepVersions property of the microsoft.graph.managedTenants.managedTenant entity.
type TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderGetQueryParameters get managementTemplateStepVersions from tenantRelationships
type TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderGetQueryParameters
}
// TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AcceptedFor provides operations to manage the acceptedFor property of the microsoft.graph.managedTenants.managementTemplateStepVersion entity.
func (m *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder) AcceptedFor()(*TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemAcceptedForRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemAcceptedForRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderInternal instantiates a new ManagementTemplateStepVersionItemRequestBuilder and sets the default values.
func NewTenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder) {
    m := &TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/managementTemplateStepVersions/{managementTemplateStepVersion%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder instantiates a new ManagementTemplateStepVersionItemRequestBuilder and sets the default values.
func NewTenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property managementTemplateStepVersions for tenantRelationships
func (m *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get managementTemplateStepVersions from tenantRelationships
func (m *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property managementTemplateStepVersions in tenantRelationships
func (m *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepVersionable, requestConfiguration *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property managementTemplateStepVersions for tenantRelationships
func (m *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Deployments provides operations to manage the deployments property of the microsoft.graph.managedTenants.managementTemplateStepVersion entity.
func (m *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder) Deployments()(*TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeploymentsById provides operations to manage the deployments property of the microsoft.graph.managedTenants.managementTemplateStepVersion entity.
func (m *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder) DeploymentsById(id string)(*TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementTemplateStepDeployment%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get managementTemplateStepVersions from tenantRelationships
func (m *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepVersionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagementTemplateStepVersionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepVersionable), nil
}
// Patch update the navigation property managementTemplateStepVersions in tenantRelationships
func (m *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder) Patch(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepVersionable, requestConfiguration *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderPatchRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepVersionable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagementTemplateStepVersionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepVersionable), nil
}
// TemplateStep provides operations to manage the templateStep property of the microsoft.graph.managedTenants.managementTemplateStepVersion entity.
func (m *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder) TemplateStep()(*TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemTemplateStepRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemTemplateStepRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
