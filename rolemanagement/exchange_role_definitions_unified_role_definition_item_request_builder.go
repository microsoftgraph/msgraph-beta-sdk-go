package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder provides operations to manage the roleDefinitions property of the microsoft.graph.unifiedRbacApplication entity.
type ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderGetQueryParameters the roles allowed by RBAC providers and the permissions assigned to the roles.
type ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderGetQueryParameters
}
// ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeId provides operations to call the assignedPrincipals method.
// returns a *ExchangeRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilder when successful
func (m *ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) AssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeId()(*ExchangeRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilder) {
    return NewExchangeRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderInternal instantiates a new ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder and sets the default values.
func NewExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) {
    m := &ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/exchange/roleDefinitions/{unifiedRoleDefinition%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder instantiates a new ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder and sets the default values.
func NewExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property roleDefinitions for roleManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the roles allowed by RBAC providers and the permissions assigned to the roles.
// returns a UnifiedRoleDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable), nil
}
// InheritsPermissionsFrom provides operations to manage the inheritsPermissionsFrom property of the microsoft.graph.unifiedRoleDefinition entity.
// returns a *ExchangeRoleDefinitionsItemInheritsPermissionsFromRequestBuilder when successful
func (m *ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) InheritsPermissionsFrom()(*ExchangeRoleDefinitionsItemInheritsPermissionsFromRequestBuilder) {
    return NewExchangeRoleDefinitionsItemInheritsPermissionsFromRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property roleDefinitions in roleManagement
// returns a UnifiedRoleDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable, requestConfiguration *ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable), nil
}
// ToDeleteRequestInformation delete navigation property roleDefinitions for roleManagement
// returns a *RequestInformation when successful
func (m *ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the roles allowed by RBAC providers and the permissions assigned to the roles.
// returns a *RequestInformation when successful
func (m *ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property roleDefinitions in roleManagement
// returns a *RequestInformation when successful
func (m *ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleDefinitionable, requestConfiguration *ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder when successful
func (m *ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) WithUrl(rawUrl string)(*ExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder) {
    return NewExchangeRoleDefinitionsUnifiedRoleDefinitionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
