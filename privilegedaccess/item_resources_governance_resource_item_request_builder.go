package privilegedaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemResourcesGovernanceResourceItemRequestBuilder provides operations to manage the resources property of the microsoft.graph.privilegedAccess entity.
type ItemResourcesGovernanceResourceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemResourcesGovernanceResourceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemResourcesGovernanceResourceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemResourcesGovernanceResourceItemRequestBuilderGetQueryParameters a collection of resources for the provider.
type ItemResourcesGovernanceResourceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemResourcesGovernanceResourceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemResourcesGovernanceResourceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemResourcesGovernanceResourceItemRequestBuilderGetQueryParameters
}
// ItemResourcesGovernanceResourceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemResourcesGovernanceResourceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemResourcesGovernanceResourceItemRequestBuilderInternal instantiates a new ItemResourcesGovernanceResourceItemRequestBuilder and sets the default values.
func NewItemResourcesGovernanceResourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemResourcesGovernanceResourceItemRequestBuilder) {
    m := &ItemResourcesGovernanceResourceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/privilegedAccess/{privilegedAccess%2Did}/resources/{governanceResource%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemResourcesGovernanceResourceItemRequestBuilder instantiates a new ItemResourcesGovernanceResourceItemRequestBuilder and sets the default values.
func NewItemResourcesGovernanceResourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemResourcesGovernanceResourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemResourcesGovernanceResourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property resources for privilegedAccess
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemResourcesGovernanceResourceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemResourcesGovernanceResourceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a collection of resources for the provider.
// returns a GovernanceResourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemResourcesGovernanceResourceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemResourcesGovernanceResourceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable), nil
}
// Parent provides operations to manage the parent property of the microsoft.graph.governanceResource entity.
// returns a *ItemResourcesItemParentRequestBuilder when successful
func (m *ItemResourcesGovernanceResourceItemRequestBuilder) Parent()(*ItemResourcesItemParentRequestBuilder) {
    return NewItemResourcesItemParentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property resources in privilegedAccess
// returns a GovernanceResourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemResourcesGovernanceResourceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable, requestConfiguration *ItemResourcesGovernanceResourceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable), nil
}
// RoleAssignmentRequests provides operations to manage the roleAssignmentRequests property of the microsoft.graph.governanceResource entity.
// returns a *ItemResourcesItemRoleAssignmentRequestsRequestBuilder when successful
func (m *ItemResourcesGovernanceResourceItemRequestBuilder) RoleAssignmentRequests()(*ItemResourcesItemRoleAssignmentRequestsRequestBuilder) {
    return NewItemResourcesItemRoleAssignmentRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleAssignments provides operations to manage the roleAssignments property of the microsoft.graph.governanceResource entity.
// returns a *ItemResourcesItemRoleAssignmentsRequestBuilder when successful
func (m *ItemResourcesGovernanceResourceItemRequestBuilder) RoleAssignments()(*ItemResourcesItemRoleAssignmentsRequestBuilder) {
    return NewItemResourcesItemRoleAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleDefinitions provides operations to manage the roleDefinitions property of the microsoft.graph.governanceResource entity.
// returns a *ItemResourcesItemRoleDefinitionsRequestBuilder when successful
func (m *ItemResourcesGovernanceResourceItemRequestBuilder) RoleDefinitions()(*ItemResourcesItemRoleDefinitionsRequestBuilder) {
    return NewItemResourcesItemRoleDefinitionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleSettings provides operations to manage the roleSettings property of the microsoft.graph.governanceResource entity.
// returns a *ItemResourcesItemRoleSettingsRequestBuilder when successful
func (m *ItemResourcesGovernanceResourceItemRequestBuilder) RoleSettings()(*ItemResourcesItemRoleSettingsRequestBuilder) {
    return NewItemResourcesItemRoleSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property resources for privilegedAccess
// returns a *RequestInformation when successful
func (m *ItemResourcesGovernanceResourceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemResourcesGovernanceResourceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/privilegedAccess/{privilegedAccess%2Did}/resources/{governanceResource%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a collection of resources for the provider.
// returns a *RequestInformation when successful
func (m *ItemResourcesGovernanceResourceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemResourcesGovernanceResourceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property resources in privilegedAccess
// returns a *RequestInformation when successful
func (m *ItemResourcesGovernanceResourceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceResourceable, requestConfiguration *ItemResourcesGovernanceResourceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/privilegedAccess/{privilegedAccess%2Did}/resources/{governanceResource%2Did}", m.BaseRequestBuilder.PathParameters)
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemResourcesGovernanceResourceItemRequestBuilder when successful
func (m *ItemResourcesGovernanceResourceItemRequestBuilder) WithUrl(rawUrl string)(*ItemResourcesGovernanceResourceItemRequestBuilder) {
    return NewItemResourcesGovernanceResourceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
