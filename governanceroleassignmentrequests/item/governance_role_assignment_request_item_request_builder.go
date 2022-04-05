package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i233cfcd85fda77bc7b8ed84ae3eab86225b54c43cf6b9995aec4e929ecae0d5f "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroleassignmentrequests/item/roledefinition"
    i3c09d8808940e57cffe2a4b43ec6c03afed85c4d1b8c2a6b6f69170f1707c06e "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroleassignmentrequests/item/cancel"
    i781dea20c6497af5c00396ab4ab2df03f2c4e8031835bbd6dffd7f1eb7bc6bbd "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroleassignmentrequests/item/updaterequest"
    i7a382583186c7a5d8b72c51d185f4589248347a161c9a5a7a7d8d885ab571893 "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroleassignmentrequests/item/resource"
    ia72138d764aac9bd3a352e3057f1578dd8230b4045684066c5c0d36a2be3badb "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroleassignmentrequests/item/subject"
)

// GovernanceRoleAssignmentRequestItemRequestBuilder provides operations to manage the collection of governanceRoleAssignmentRequest entities.
type GovernanceRoleAssignmentRequestItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GovernanceRoleAssignmentRequestItemRequestBuilderDeleteOptions options for Delete
type GovernanceRoleAssignmentRequestItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// GovernanceRoleAssignmentRequestItemRequestBuilderGetOptions options for Get
type GovernanceRoleAssignmentRequestItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *GovernanceRoleAssignmentRequestItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// GovernanceRoleAssignmentRequestItemRequestBuilderGetQueryParameters get entity from governanceRoleAssignmentRequests by key
type GovernanceRoleAssignmentRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// GovernanceRoleAssignmentRequestItemRequestBuilderPatchOptions options for Patch
type GovernanceRoleAssignmentRequestItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// Cancel the cancel property
func (m *GovernanceRoleAssignmentRequestItemRequestBuilder) Cancel()(*i3c09d8808940e57cffe2a4b43ec6c03afed85c4d1b8c2a6b6f69170f1707c06e.CancelRequestBuilder) {
    return i3c09d8808940e57cffe2a4b43ec6c03afed85c4d1b8c2a6b6f69170f1707c06e.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGovernanceRoleAssignmentRequestItemRequestBuilderInternal instantiates a new GovernanceRoleAssignmentRequestItemRequestBuilder and sets the default values.
func NewGovernanceRoleAssignmentRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GovernanceRoleAssignmentRequestItemRequestBuilder) {
    m := &GovernanceRoleAssignmentRequestItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/governanceRoleAssignmentRequests/{governanceRoleAssignmentRequest_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGovernanceRoleAssignmentRequestItemRequestBuilder instantiates a new GovernanceRoleAssignmentRequestItemRequestBuilder and sets the default values.
func NewGovernanceRoleAssignmentRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GovernanceRoleAssignmentRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGovernanceRoleAssignmentRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from governanceRoleAssignmentRequests
func (m *GovernanceRoleAssignmentRequestItemRequestBuilder) CreateDeleteRequestInformation(options *GovernanceRoleAssignmentRequestItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get entity from governanceRoleAssignmentRequests by key
func (m *GovernanceRoleAssignmentRequestItemRequestBuilder) CreateGetRequestInformation(options *GovernanceRoleAssignmentRequestItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update entity in governanceRoleAssignmentRequests
func (m *GovernanceRoleAssignmentRequestItemRequestBuilder) CreatePatchRequestInformation(options *GovernanceRoleAssignmentRequestItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete entity from governanceRoleAssignmentRequests
func (m *GovernanceRoleAssignmentRequestItemRequestBuilder) Delete(options *GovernanceRoleAssignmentRequestItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get entity from governanceRoleAssignmentRequests by key
func (m *GovernanceRoleAssignmentRequestItemRequestBuilder) Get(options *GovernanceRoleAssignmentRequestItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceRoleAssignmentRequestFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable), nil
}
// Patch update entity in governanceRoleAssignmentRequests
func (m *GovernanceRoleAssignmentRequestItemRequestBuilder) Patch(options *GovernanceRoleAssignmentRequestItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Resource the resource property
func (m *GovernanceRoleAssignmentRequestItemRequestBuilder) Resource()(*i7a382583186c7a5d8b72c51d185f4589248347a161c9a5a7a7d8d885ab571893.ResourceRequestBuilder) {
    return i7a382583186c7a5d8b72c51d185f4589248347a161c9a5a7a7d8d885ab571893.NewResourceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinition the roleDefinition property
func (m *GovernanceRoleAssignmentRequestItemRequestBuilder) RoleDefinition()(*i233cfcd85fda77bc7b8ed84ae3eab86225b54c43cf6b9995aec4e929ecae0d5f.RoleDefinitionRequestBuilder) {
    return i233cfcd85fda77bc7b8ed84ae3eab86225b54c43cf6b9995aec4e929ecae0d5f.NewRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Subject the subject property
func (m *GovernanceRoleAssignmentRequestItemRequestBuilder) Subject()(*ia72138d764aac9bd3a352e3057f1578dd8230b4045684066c5c0d36a2be3badb.SubjectRequestBuilder) {
    return ia72138d764aac9bd3a352e3057f1578dd8230b4045684066c5c0d36a2be3badb.NewSubjectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateRequest the updateRequest property
func (m *GovernanceRoleAssignmentRequestItemRequestBuilder) UpdateRequest()(*i781dea20c6497af5c00396ab4ab2df03f2c4e8031835bbd6dffd7f1eb7bc6bbd.UpdateRequestRequestBuilder) {
    return i781dea20c6497af5c00396ab4ab2df03f2c4e8031835bbd6dffd7f1eb7bc6bbd.NewUpdateRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
