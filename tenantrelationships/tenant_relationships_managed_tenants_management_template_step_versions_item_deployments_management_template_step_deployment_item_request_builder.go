package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilder provides operations to manage the deployments property of the microsoft.graph.managedTenants.managementTemplateStepVersion entity.
type TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilderGetQueryParameters get deployments from tenantRelationships
type TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilderGetQueryParameters
}
// TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ChangeDeploymentStatus provides operations to call the changeDeploymentStatus method.
func (m *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilder) ChangeDeploymentStatus()(*TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemChangeDeploymentStatusRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemChangeDeploymentStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilderInternal instantiates a new ManagementTemplateStepDeploymentItemRequestBuilder and sets the default values.
func NewTenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilder) {
    m := &TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/managementTemplateStepVersions/{managementTemplateStepVersion%2Did}/deployments/{managementTemplateStepDeployment%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilder instantiates a new ManagementTemplateStepDeploymentItemRequestBuilder and sets the default values.
func NewTenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property deployments for tenantRelationships
func (m *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get deployments from tenantRelationships
func (m *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property deployments in tenantRelationships
func (m *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepDeploymentable, requestConfiguration *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property deployments for tenantRelationships
func (m *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get deployments from tenantRelationships
func (m *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepDeploymentable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagementTemplateStepDeploymentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepDeploymentable), nil
}
// Patch update the navigation property deployments in tenantRelationships
func (m *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilder) Patch(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepDeploymentable, requestConfiguration *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilderPatchRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepDeploymentable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagementTemplateStepDeploymentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepDeploymentable), nil
}
// TemplateStepVersion provides operations to manage the templateStepVersion property of the microsoft.graph.managedTenants.managementTemplateStepDeployment entity.
func (m *TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilder) TemplateStepVersion()(*TenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemTemplateStepVersionRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemTemplateStepVersionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
