package governanceresources

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder provides operations to manage the roleAssignmentRequests property of the microsoft.graph.governanceResource entity.
type ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetQueryParameters the collection of role assignment requests for the resource.
type ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetQueryParameters
}
// ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Cancel provides operations to call the cancel method.
// returns a *ItemRoleAssignmentRequestsItemCancelRequestBuilder when successful
func (m *ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) Cancel()(*ItemRoleAssignmentRequestsItemCancelRequestBuilder) {
    return NewItemRoleAssignmentRequestsItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderInternal instantiates a new ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder and sets the default values.
func NewItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) {
    m := &ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/governanceResources/{governanceResource%2Did}/roleAssignmentRequests/{governanceRoleAssignmentRequest%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder instantiates a new ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder and sets the default values.
func NewItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property roleAssignmentRequests for governanceResources
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection of role assignment requests for the resource.
// returns a GovernanceRoleAssignmentRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceRoleAssignmentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable), nil
}
// Patch update the navigation property roleAssignmentRequests in governanceResources
// returns a GovernanceRoleAssignmentRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable, requestConfiguration *ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceRoleAssignmentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable), nil
}
// Resource provides operations to manage the resource property of the microsoft.graph.governanceRoleAssignmentRequest entity.
// returns a *ItemRoleAssignmentRequestsItemResourceRequestBuilder when successful
func (m *ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) Resource()(*ItemRoleAssignmentRequestsItemResourceRequestBuilder) {
    return NewItemRoleAssignmentRequestsItemResourceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleDefinition provides operations to manage the roleDefinition property of the microsoft.graph.governanceRoleAssignmentRequest entity.
// returns a *ItemRoleAssignmentRequestsItemRoleDefinitionRequestBuilder when successful
func (m *ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) RoleDefinition()(*ItemRoleAssignmentRequestsItemRoleDefinitionRequestBuilder) {
    return NewItemRoleAssignmentRequestsItemRoleDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Subject provides operations to manage the subject property of the microsoft.graph.governanceRoleAssignmentRequest entity.
// returns a *ItemRoleAssignmentRequestsItemSubjectRequestBuilder when successful
func (m *ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) Subject()(*ItemRoleAssignmentRequestsItemSubjectRequestBuilder) {
    return NewItemRoleAssignmentRequestsItemSubjectRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property roleAssignmentRequests for governanceResources
// returns a *RequestInformation when successful
func (m *ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/governanceResources/{governanceResource%2Did}/roleAssignmentRequests/{governanceRoleAssignmentRequest%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the collection of role assignment requests for the resource.
// returns a *RequestInformation when successful
func (m *ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property roleAssignmentRequests in governanceResources
// returns a *RequestInformation when successful
func (m *ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable, requestConfiguration *ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/governanceResources/{governanceResource%2Did}/roleAssignmentRequests/{governanceRoleAssignmentRequest%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// UpdateRequest provides operations to call the updateRequest method.
// returns a *ItemRoleAssignmentRequestsItemUpdateRequestRequestBuilder when successful
func (m *ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) UpdateRequest()(*ItemRoleAssignmentRequestsItemUpdateRequestRequestBuilder) {
    return NewItemRoleAssignmentRequestsItemUpdateRequestRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder when successful
func (m *ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) WithUrl(rawUrl string)(*ItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder) {
    return NewItemRoleAssignmentRequestsGovernanceRoleAssignmentRequestItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
