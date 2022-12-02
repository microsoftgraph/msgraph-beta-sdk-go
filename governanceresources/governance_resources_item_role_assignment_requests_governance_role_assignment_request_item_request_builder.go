package governanceresources

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder provides operations to manage the roleAssignmentRequests property of the microsoft.graph.governanceResource entity.
type GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetQueryParameters the collection of role assignment requests for the resource.
type GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetQueryParameters
}
// GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Cancel provides operations to call the cancel method.
func (m *GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) Cancel()(*GovernanceResourcesItemRoleAssignmentRequestsItemCancelRequestBuilder) {
    return NewGovernanceResourcesItemRoleAssignmentRequestsItemCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderInternal instantiates a new GovernanceRoleAssignmentRequestItemRequestBuilder and sets the default values.
func NewGovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) {
    m := &GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/governanceResources/{governanceResource%2Did}/roleAssignmentRequests/{governanceRoleAssignmentRequest%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder instantiates a new GovernanceRoleAssignmentRequestItemRequestBuilder and sets the default values.
func NewGovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property roleAssignmentRequests for governanceResources
func (m *GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property roleAssignmentRequests in governanceResources
func (m *GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable, requestConfiguration *GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property roleAssignmentRequests for governanceResources
func (m *GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable, error) {
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
// Patch update the navigation property roleAssignmentRequests in governanceResources
func (m *GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable, requestConfiguration *GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable, error) {
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
func (m *GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) Resource()(*GovernanceResourcesItemRoleAssignmentRequestsItemResourceRequestBuilder) {
    return NewGovernanceResourcesItemRoleAssignmentRequestsItemResourceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinition provides operations to manage the roleDefinition property of the microsoft.graph.governanceRoleAssignmentRequest entity.
func (m *GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) RoleDefinition()(*GovernanceResourcesItemRoleAssignmentRequestsItemRoleDefinitionRequestBuilder) {
    return NewGovernanceResourcesItemRoleAssignmentRequestsItemRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Subject provides operations to manage the subject property of the microsoft.graph.governanceRoleAssignmentRequest entity.
func (m *GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) Subject()(*GovernanceResourcesItemRoleAssignmentRequestsItemSubjectRequestBuilder) {
    return NewGovernanceResourcesItemRoleAssignmentRequestsItemSubjectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateRequest provides operations to call the updateRequest method.
func (m *GovernanceResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) UpdateRequest()(*GovernanceResourcesItemRoleAssignmentRequestsItemUpdateRequestRequestBuilder) {
    return NewGovernanceResourcesItemRoleAssignmentRequestsItemUpdateRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
