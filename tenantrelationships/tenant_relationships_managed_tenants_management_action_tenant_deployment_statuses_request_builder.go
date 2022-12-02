package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilder provides operations to manage the managementActionTenantDeploymentStatuses property of the microsoft.graph.managedTenants.managedTenant entity.
type TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilderGetQueryParameters get a list of the managementActionTenantDeploymentStatus objects and their properties.
type TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilderGetQueryParameters struct {
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
// TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilderGetQueryParameters
}
// TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ChangeDeploymentStatus provides operations to call the changeDeploymentStatus method.
func (m *TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilder) ChangeDeploymentStatus()(*TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilderInternal instantiates a new ManagementActionTenantDeploymentStatusesRequestBuilder and sets the default values.
func NewTenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilder) {
    m := &TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/managementActionTenantDeploymentStatuses{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilder instantiates a new ManagementActionTenantDeploymentStatusesRequestBuilder and sets the default values.
func NewTenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilder) Count()(*TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesCountRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get a list of the managementActionTenantDeploymentStatus objects and their properties.
func (m *TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to managementActionTenantDeploymentStatuses for tenantRelationships
func (m *TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilder) CreatePostRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementActionTenantDeploymentStatusable, requestConfiguration *TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get get a list of the managementActionTenantDeploymentStatus objects and their properties.
func (m *TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilder) Get(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementActionTenantDeploymentStatusCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagementActionTenantDeploymentStatusCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementActionTenantDeploymentStatusCollectionResponseable), nil
}
// Post create new navigation property to managementActionTenantDeploymentStatuses for tenantRelationships
func (m *TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilder) Post(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementActionTenantDeploymentStatusable, requestConfiguration *TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilderPostRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementActionTenantDeploymentStatusable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagementActionTenantDeploymentStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementActionTenantDeploymentStatusable), nil
}
