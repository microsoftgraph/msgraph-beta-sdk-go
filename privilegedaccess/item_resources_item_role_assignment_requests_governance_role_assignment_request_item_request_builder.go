package privilegedaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder provides operations to manage the roleAssignmentRequests property of the microsoft.graph.governanceResource entity.
type ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetQueryParameters the collection of role assignment requests for the resource.
type ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetQueryParameters
}
// ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderInternal instantiates a new GovernanceRoleAssignmentRequestItemRequestBuilder and sets the default values.
func NewItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) {
    m := &ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privilegedAccess/{privilegedAccess%2Did}/resources/{governanceResource%2Did}/roleAssignmentRequests/{governanceRoleAssignmentRequest%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder instantiates a new GovernanceRoleAssignmentRequestItemRequestBuilder and sets the default values.
func NewItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property roleAssignmentRequests for privilegedAccess
func (m *ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection of role assignment requests for the resource.
func (m *ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceRoleAssignmentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable), nil
}
// MicrosoftGraphCancel provides operations to call the cancel method.
func (m *ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) MicrosoftGraphCancel()(*ItemResourcesItemRoleAssignmentRequestsItemMicrosoftGraphCancelRequestBuilder) {
    return NewItemResourcesItemRoleAssignmentRequestsItemMicrosoftGraphCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUpdateRequest provides operations to call the updateRequest method.
func (m *ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) MicrosoftGraphUpdateRequest()(*ItemResourcesItemRoleAssignmentRequestsItemMicrosoftGraphUpdateRequestRequestBuilder) {
    return NewItemResourcesItemRoleAssignmentRequestsItemMicrosoftGraphUpdateRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property roleAssignmentRequests in privilegedAccess
func (m *ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable, requestConfiguration *ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceRoleAssignmentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable), nil
}
// Resource provides operations to manage the resource property of the microsoft.graph.governanceRoleAssignmentRequest entity.
func (m *ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) Resource()(*ItemResourcesItemRoleAssignmentRequestsItemResourceRequestBuilder) {
    return NewItemResourcesItemRoleAssignmentRequestsItemResourceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RoleDefinition provides operations to manage the roleDefinition property of the microsoft.graph.governanceRoleAssignmentRequest entity.
func (m *ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) RoleDefinition()(*ItemResourcesItemRoleAssignmentRequestsItemRoleDefinitionRequestBuilder) {
    return NewItemResourcesItemRoleAssignmentRequestsItemRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Subject provides operations to manage the subject property of the microsoft.graph.governanceRoleAssignmentRequest entity.
func (m *ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) Subject()(*ItemResourcesItemRoleAssignmentRequestsItemSubjectRequestBuilder) {
    return NewItemResourcesItemRoleAssignmentRequestsItemSubjectRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property roleAssignmentRequests for privilegedAccess
func (m *ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation the collection of role assignment requests for the resource.
func (m *ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property roleAssignmentRequests in privilegedAccess
func (m *ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable, requestConfiguration *ItemResourcesItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
