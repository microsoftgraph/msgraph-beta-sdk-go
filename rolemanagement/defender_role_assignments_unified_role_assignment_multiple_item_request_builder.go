package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder provides operations to manage the roleAssignments property of the microsoft.graph.rbacApplicationMultiple entity.
type DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderGetQueryParameters get roleAssignments from roleManagement
type DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderGetQueryParameters
}
// DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppScopes provides operations to manage the appScopes property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
// returns a *DefenderRoleAssignmentsItemAppScopesRequestBuilder when successful
func (m *DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) AppScopes()(*DefenderRoleAssignmentsItemAppScopesRequestBuilder) {
    return NewDefenderRoleAssignmentsItemAppScopesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderInternal instantiates a new DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder and sets the default values.
func NewDefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) {
    m := &DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/defender/roleAssignments/{unifiedRoleAssignmentMultiple%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder instantiates a new DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder and sets the default values.
func NewDefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property roleAssignments for roleManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DirectoryScopes provides operations to manage the directoryScopes property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
// returns a *DefenderRoleAssignmentsItemDirectoryScopesRequestBuilder when successful
func (m *DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) DirectoryScopes()(*DefenderRoleAssignmentsItemDirectoryScopesRequestBuilder) {
    return NewDefenderRoleAssignmentsItemDirectoryScopesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get roleAssignments from roleManagement
// returns a UnifiedRoleAssignmentMultipleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentMultipleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleAssignmentMultipleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentMultipleable), nil
}
// Patch update the navigation property roleAssignments in roleManagement
// returns a UnifiedRoleAssignmentMultipleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentMultipleable, requestConfiguration *DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentMultipleable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleAssignmentMultipleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentMultipleable), nil
}
// Principals provides operations to manage the principals property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
// returns a *DefenderRoleAssignmentsItemPrincipalsRequestBuilder when successful
func (m *DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) Principals()(*DefenderRoleAssignmentsItemPrincipalsRequestBuilder) {
    return NewDefenderRoleAssignmentsItemPrincipalsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleDefinition provides operations to manage the roleDefinition property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
// returns a *DefenderRoleAssignmentsItemRoleDefinitionRequestBuilder when successful
func (m *DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) RoleDefinition()(*DefenderRoleAssignmentsItemRoleDefinitionRequestBuilder) {
    return NewDefenderRoleAssignmentsItemRoleDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property roleAssignments for roleManagement
// returns a *RequestInformation when successful
func (m *DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get roleAssignments from roleManagement
// returns a *RequestInformation when successful
func (m *DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property roleAssignments in roleManagement
// returns a *RequestInformation when successful
func (m *DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentMultipleable, requestConfiguration *DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder when successful
func (m *DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) WithUrl(rawUrl string)(*DefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder) {
    return NewDefenderRoleAssignmentsUnifiedRoleAssignmentMultipleItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
