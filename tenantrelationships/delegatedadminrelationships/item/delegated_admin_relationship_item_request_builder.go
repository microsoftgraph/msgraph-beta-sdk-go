package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i4fa8edb0c045897c1eedba2d597f16b1fbe1eba86c010bbb3f99622813ad295e "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/delegatedadminrelationships/item/requests"
    ia0161922114d3b8c7cd781d7f9c4106c03db7e9ef8ea72513b4c54b5e9354ba5 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/delegatedadminrelationships/item/operations"
    iac4781d78e0362023abac407b248abb8996da024ec94a9413fc9a3d80755d0ea "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/delegatedadminrelationships/item/accessassignments"
    i0580a758f5bfbd03e74223f96f6b1bdb17dd6ffec4a223b1f1fabb8a627a42b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/delegatedadminrelationships/item/requests/item"
    i5f3e9a785873705d5204dd08bc17b79ce358f2fd632fbe5eb537c2a5d7fc7ef6 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/delegatedadminrelationships/item/operations/item"
    ic1a2347ea4abb46ab983c9d6ee916f3a9af60bf8f4caf7afae2666a77def3567 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/delegatedadminrelationships/item/accessassignments/item"
)

// DelegatedAdminRelationshipItemRequestBuilder provides operations to manage the delegatedAdminRelationships property of the microsoft.graph.tenantRelationship entity.
type DelegatedAdminRelationshipItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DelegatedAdminRelationshipItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedAdminRelationshipItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DelegatedAdminRelationshipItemRequestBuilderGetQueryParameters the details of the delegated administrative privileges that a Microsoft partner has in a customer tenant.
type DelegatedAdminRelationshipItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DelegatedAdminRelationshipItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedAdminRelationshipItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DelegatedAdminRelationshipItemRequestBuilderGetQueryParameters
}
// DelegatedAdminRelationshipItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedAdminRelationshipItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessAssignments the accessAssignments property
func (m *DelegatedAdminRelationshipItemRequestBuilder) AccessAssignments()(*iac4781d78e0362023abac407b248abb8996da024ec94a9413fc9a3d80755d0ea.AccessAssignmentsRequestBuilder) {
    return iac4781d78e0362023abac407b248abb8996da024ec94a9413fc9a3d80755d0ea.NewAccessAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.delegatedAdminRelationships.item.accessAssignments.item collection
func (m *DelegatedAdminRelationshipItemRequestBuilder) AccessAssignmentsById(id string)(*ic1a2347ea4abb46ab983c9d6ee916f3a9af60bf8f4caf7afae2666a77def3567.DelegatedAdminAccessAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["delegatedAdminAccessAssignment%2Did"] = id
    }
    return ic1a2347ea4abb46ab983c9d6ee916f3a9af60bf8f4caf7afae2666a77def3567.NewDelegatedAdminAccessAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDelegatedAdminRelationshipItemRequestBuilderInternal instantiates a new DelegatedAdminRelationshipItemRequestBuilder and sets the default values.
func NewDelegatedAdminRelationshipItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DelegatedAdminRelationshipItemRequestBuilder) {
    m := &DelegatedAdminRelationshipItemRequestBuilder{
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
// NewDelegatedAdminRelationshipItemRequestBuilder instantiates a new DelegatedAdminRelationshipItemRequestBuilder and sets the default values.
func NewDelegatedAdminRelationshipItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DelegatedAdminRelationshipItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDelegatedAdminRelationshipItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property delegatedAdminRelationships for tenantRelationships
func (m *DelegatedAdminRelationshipItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property delegatedAdminRelationships for tenantRelationships
func (m *DelegatedAdminRelationshipItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *DelegatedAdminRelationshipItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DelegatedAdminRelationshipItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the details of the delegated administrative privileges that a Microsoft partner has in a customer tenant.
func (m *DelegatedAdminRelationshipItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DelegatedAdminRelationshipItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DelegatedAdminRelationshipItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminRelationshipable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property delegatedAdminRelationships in tenantRelationships
func (m *DelegatedAdminRelationshipItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminRelationshipable, requestConfiguration *DelegatedAdminRelationshipItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property delegatedAdminRelationships for tenantRelationships
func (m *DelegatedAdminRelationshipItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DelegatedAdminRelationshipItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
func (m *DelegatedAdminRelationshipItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DelegatedAdminRelationshipItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminRelationshipable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
// Operations the operations property
func (m *DelegatedAdminRelationshipItemRequestBuilder) Operations()(*ia0161922114d3b8c7cd781d7f9c4106c03db7e9ef8ea72513b4c54b5e9354ba5.OperationsRequestBuilder) {
    return ia0161922114d3b8c7cd781d7f9c4106c03db7e9ef8ea72513b4c54b5e9354ba5.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.delegatedAdminRelationships.item.operations.item collection
func (m *DelegatedAdminRelationshipItemRequestBuilder) OperationsById(id string)(*i5f3e9a785873705d5204dd08bc17b79ce358f2fd632fbe5eb537c2a5d7fc7ef6.DelegatedAdminRelationshipOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["delegatedAdminRelationshipOperation%2Did"] = id
    }
    return i5f3e9a785873705d5204dd08bc17b79ce358f2fd632fbe5eb537c2a5d7fc7ef6.NewDelegatedAdminRelationshipOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property delegatedAdminRelationships in tenantRelationships
func (m *DelegatedAdminRelationshipItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminRelationshipable, requestConfiguration *DelegatedAdminRelationshipItemRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// Requests the requests property
func (m *DelegatedAdminRelationshipItemRequestBuilder) Requests()(*i4fa8edb0c045897c1eedba2d597f16b1fbe1eba86c010bbb3f99622813ad295e.RequestsRequestBuilder) {
    return i4fa8edb0c045897c1eedba2d597f16b1fbe1eba86c010bbb3f99622813ad295e.NewRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.delegatedAdminRelationships.item.requests.item collection
func (m *DelegatedAdminRelationshipItemRequestBuilder) RequestsById(id string)(*i0580a758f5bfbd03e74223f96f6b1bdb17dd6ffec4a223b1f1fabb8a627a42b4.DelegatedAdminRelationshipRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["delegatedAdminRelationshipRequest%2Did"] = id
    }
    return i0580a758f5bfbd03e74223f96f6b1bdb17dd6ffec4a223b1f1fabb8a627a42b4.NewDelegatedAdminRelationshipRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
