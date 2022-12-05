package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilder provides operations to manage the serviceManagementDetails property of the microsoft.graph.delegatedAdminCustomer entity.
type TenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilderGetQueryParameters contains the management details of a service in the customer tenant that's managed by delegated administration.
type TenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilderGetQueryParameters
}
// TenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilderInternal instantiates a new DelegatedAdminServiceManagementDetailItemRequestBuilder and sets the default values.
func NewTenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilder) {
    m := &TenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/delegatedAdminCustomers/{delegatedAdminCustomer%2Did}/serviceManagementDetails/{delegatedAdminServiceManagementDetail%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilder instantiates a new DelegatedAdminServiceManagementDetailItemRequestBuilder and sets the default values.
func NewTenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property serviceManagementDetails for tenantRelationships
func (m *TenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *TenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation contains the management details of a service in the customer tenant that's managed by delegated administration.
func (m *TenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property serviceManagementDetails in tenantRelationships
func (m *TenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminServiceManagementDetailable, requestConfiguration *TenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property serviceManagementDetails for tenantRelationships
func (m *TenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get contains the management details of a service in the customer tenant that's managed by delegated administration.
func (m *TenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminServiceManagementDetailable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDelegatedAdminServiceManagementDetailFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminServiceManagementDetailable), nil
}
// Patch update the navigation property serviceManagementDetails in tenantRelationships
func (m *TenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminServiceManagementDetailable, requestConfiguration *TenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminServiceManagementDetailable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDelegatedAdminServiceManagementDetailFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminServiceManagementDetailable), nil
}
