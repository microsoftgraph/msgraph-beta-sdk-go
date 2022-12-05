package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder provides operations to manage the delegatedAdminRelationships property of the microsoft.graph.tenantRelationship entity.
type TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderGetQueryParameters the details of the delegated administrative privileges that a Microsoft partner has in a customer tenant.
type TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderGetQueryParameters
}
// TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessAssignments provides operations to manage the accessAssignments property of the microsoft.graph.delegatedAdminRelationship entity.
func (m *TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) AccessAssignments()(*TenantRelationshipsDelegatedAdminRelationshipsItemAccessAssignmentsRequestBuilder) {
    return NewTenantRelationshipsDelegatedAdminRelationshipsItemAccessAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessAssignmentsById provides operations to manage the accessAssignments property of the microsoft.graph.delegatedAdminRelationship entity.
func (m *TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) AccessAssignmentsById(id string)(*TenantRelationshipsDelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["delegatedAdminAccessAssignment%2Did"] = id
    }
    return NewTenantRelationshipsDelegatedAdminRelationshipsItemAccessAssignmentsDelegatedAdminAccessAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderInternal instantiates a new DelegatedAdminRelationshipItemRequestBuilder and sets the default values.
func NewTenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) {
    m := &TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/delegatedAdminRelationships/{delegatedAdminRelationship%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder instantiates a new DelegatedAdminRelationshipItemRequestBuilder and sets the default values.
func NewTenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property delegatedAdminRelationships for tenantRelationships
func (m *TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the details of the delegated administrative privileges that a Microsoft partner has in a customer tenant.
func (m *TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property delegatedAdminRelationships in tenantRelationships
func (m *TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminRelationshipable, requestConfiguration *TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property delegatedAdminRelationships for tenantRelationships
func (m *TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the details of the delegated administrative privileges that a Microsoft partner has in a customer tenant.
func (m *TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminRelationshipable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDelegatedAdminRelationshipFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminRelationshipable), nil
}
// Operations provides operations to manage the operations property of the microsoft.graph.delegatedAdminRelationship entity.
func (m *TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) Operations()(*TenantRelationshipsDelegatedAdminRelationshipsItemOperationsRequestBuilder) {
    return NewTenantRelationshipsDelegatedAdminRelationshipsItemOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.delegatedAdminRelationship entity.
func (m *TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) OperationsById(id string)(*TenantRelationshipsDelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["delegatedAdminRelationshipOperation%2Did"] = id
    }
    return NewTenantRelationshipsDelegatedAdminRelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property delegatedAdminRelationships in tenantRelationships
func (m *TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminRelationshipable, requestConfiguration *TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminRelationshipable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDelegatedAdminRelationshipFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminRelationshipable), nil
}
// Requests provides operations to manage the requests property of the microsoft.graph.delegatedAdminRelationship entity.
func (m *TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) Requests()(*TenantRelationshipsDelegatedAdminRelationshipsItemRequestsRequestBuilder) {
    return NewTenantRelationshipsDelegatedAdminRelationshipsItemRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RequestsById provides operations to manage the requests property of the microsoft.graph.delegatedAdminRelationship entity.
func (m *TenantRelationshipsDelegatedAdminRelationshipsDelegatedAdminRelationshipItemRequestBuilder) RequestsById(id string)(*TenantRelationshipsDelegatedAdminRelationshipsItemRequestsDelegatedAdminRelationshipRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["delegatedAdminRelationshipRequest%2Did"] = id
    }
    return NewTenantRelationshipsDelegatedAdminRelationshipsItemRequestsDelegatedAdminRelationshipRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
