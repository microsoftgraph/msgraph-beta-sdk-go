package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilder provides operations to manage the delegatedAdminCustomers property of the microsoft.graph.tenantRelationship entity.
type TenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilderGetQueryParameters the customer who has a delegated admin relationship with a Microsoft partner.
type TenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilderGetQueryParameters
}
// TenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilderInternal instantiates a new DelegatedAdminCustomerItemRequestBuilder and sets the default values.
func NewTenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilder) {
    m := &TenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/delegatedAdminCustomers/{delegatedAdminCustomer%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilder instantiates a new DelegatedAdminCustomerItemRequestBuilder and sets the default values.
func NewTenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property delegatedAdminCustomers for tenantRelationships
func (m *TenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *TenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the customer who has a delegated admin relationship with a Microsoft partner.
func (m *TenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property delegatedAdminCustomers in tenantRelationships
func (m *TenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminCustomerable, requestConfiguration *TenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property delegatedAdminCustomers for tenantRelationships
func (m *TenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the customer who has a delegated admin relationship with a Microsoft partner.
func (m *TenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminCustomerable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDelegatedAdminCustomerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminCustomerable), nil
}
// Patch update the navigation property delegatedAdminCustomers in tenantRelationships
func (m *TenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminCustomerable, requestConfiguration *TenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminCustomerable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDelegatedAdminCustomerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminCustomerable), nil
}
// ServiceManagementDetails provides operations to manage the serviceManagementDetails property of the microsoft.graph.delegatedAdminCustomer entity.
func (m *TenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilder) ServiceManagementDetails()(*TenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsRequestBuilder) {
    return NewTenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServiceManagementDetailsById provides operations to manage the serviceManagementDetails property of the microsoft.graph.delegatedAdminCustomer entity.
func (m *TenantRelationshipsDelegatedAdminCustomersDelegatedAdminCustomerItemRequestBuilder) ServiceManagementDetailsById(id string)(*TenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["delegatedAdminServiceManagementDetail%2Did"] = id
    }
    return NewTenantRelationshipsDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
