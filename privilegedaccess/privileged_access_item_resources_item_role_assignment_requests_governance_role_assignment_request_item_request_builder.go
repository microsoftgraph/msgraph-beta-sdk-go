package privilegedaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder provides operations to manage the roleAssignmentRequests property of the microsoft.graph.governanceResource entity.
type PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetQueryParameters the collection of role assignment requests for the resource.
type PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetQueryParameters
}
// PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Cancel provides operations to call the cancel method.
func (m *PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) Cancel()(*PrivilegedAccessItemResourcesItemRoleAssignmentRequestsItemCancelRequestBuilder) {
    return NewPrivilegedAccessItemResourcesItemRoleAssignmentRequestsItemCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewPrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderInternal instantiates a new GovernanceRoleAssignmentRequestItemRequestBuilder and sets the default values.
func NewPrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) {
    m := &PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privilegedAccess/{privilegedAccess%2Did}/resources/{governanceResource%2Did}/roleAssignmentRequests/{governanceRoleAssignmentRequest%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder instantiates a new GovernanceRoleAssignmentRequestItemRequestBuilder and sets the default values.
func NewPrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property roleAssignmentRequests for privilegedAccess
func (m *PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of role assignment requests for the resource.
func (m *PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property roleAssignmentRequests in privilegedAccess
func (m *PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable, requestConfiguration *PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property roleAssignmentRequests for privilegedAccess
func (m *PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the collection of role assignment requests for the resource.
func (m *PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceRoleAssignmentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable), nil
}
// Patch update the navigation property roleAssignmentRequests in privilegedAccess
func (m *PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable, requestConfiguration *PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceRoleAssignmentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable), nil
}
// Resource provides operations to manage the resource property of the microsoft.graph.governanceRoleAssignmentRequest entity.
func (m *PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) Resource()(*PrivilegedAccessItemResourcesItemRoleAssignmentRequestsItemResourceRequestBuilder) {
    return NewPrivilegedAccessItemResourcesItemRoleAssignmentRequestsItemResourceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinition provides operations to manage the roleDefinition property of the microsoft.graph.governanceRoleAssignmentRequest entity.
func (m *PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) RoleDefinition()(*PrivilegedAccessItemResourcesItemRoleAssignmentRequestsItemRoleDefinitionRequestBuilder) {
    return NewPrivilegedAccessItemResourcesItemRoleAssignmentRequestsItemRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Subject provides operations to manage the subject property of the microsoft.graph.governanceRoleAssignmentRequest entity.
func (m *PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) Subject()(*PrivilegedAccessItemResourcesItemRoleAssignmentRequestsItemSubjectRequestBuilder) {
    return NewPrivilegedAccessItemResourcesItemRoleAssignmentRequestsItemSubjectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateRequest provides operations to call the updateRequest method.
func (m *PrivilegedAccessItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) UpdateRequest()(*PrivilegedAccessItemResourcesItemRoleAssignmentRequestsItemUpdateRequestRequestBuilder) {
    return NewPrivilegedAccessItemResourcesItemRoleAssignmentRequestsItemUpdateRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
