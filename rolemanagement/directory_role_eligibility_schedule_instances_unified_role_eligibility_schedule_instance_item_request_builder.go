package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder provides operations to manage the roleEligibilityScheduleInstances property of the microsoft.graph.rbacApplication entity.
type DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderGetQueryParameters read the properties and relationships of an unifiedRoleEligibilityScheduleInstance object.
type DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderGetQueryParameters
}
// DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppScope provides operations to manage the appScope property of the microsoft.graph.unifiedRoleScheduleInstanceBase entity.
// returns a *DirectoryRoleEligibilityScheduleInstancesItemAppScopeRequestBuilder when successful
func (m *DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) AppScope()(*DirectoryRoleEligibilityScheduleInstancesItemAppScopeRequestBuilder) {
    return NewDirectoryRoleEligibilityScheduleInstancesItemAppScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderInternal instantiates a new DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder and sets the default values.
func NewDirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) {
    m := &DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory/roleEligibilityScheduleInstances/{unifiedRoleEligibilityScheduleInstance%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder instantiates a new DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder and sets the default values.
func NewDirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property roleEligibilityScheduleInstances for roleManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DirectoryScope provides operations to manage the directoryScope property of the microsoft.graph.unifiedRoleScheduleInstanceBase entity.
// returns a *DirectoryRoleEligibilityScheduleInstancesItemDirectoryScopeRequestBuilder when successful
func (m *DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) DirectoryScope()(*DirectoryRoleEligibilityScheduleInstancesItemDirectoryScopeRequestBuilder) {
    return NewDirectoryRoleEligibilityScheduleInstancesItemDirectoryScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read the properties and relationships of an unifiedRoleEligibilityScheduleInstance object.
// returns a UnifiedRoleEligibilityScheduleInstanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedroleeligibilityscheduleinstance-get?view=graph-rest-1.0
func (m *DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleInstanceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleEligibilityScheduleInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleInstanceable), nil
}
// Patch update the navigation property roleEligibilityScheduleInstances in roleManagement
// returns a UnifiedRoleEligibilityScheduleInstanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleInstanceable, requestConfiguration *DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleInstanceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleEligibilityScheduleInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleInstanceable), nil
}
// Principal provides operations to manage the principal property of the microsoft.graph.unifiedRoleScheduleInstanceBase entity.
// returns a *DirectoryRoleEligibilityScheduleInstancesItemPrincipalRequestBuilder when successful
func (m *DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) Principal()(*DirectoryRoleEligibilityScheduleInstancesItemPrincipalRequestBuilder) {
    return NewDirectoryRoleEligibilityScheduleInstancesItemPrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleDefinition provides operations to manage the roleDefinition property of the microsoft.graph.unifiedRoleScheduleInstanceBase entity.
// returns a *DirectoryRoleEligibilityScheduleInstancesItemRoleDefinitionRequestBuilder when successful
func (m *DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) RoleDefinition()(*DirectoryRoleEligibilityScheduleInstancesItemRoleDefinitionRequestBuilder) {
    return NewDirectoryRoleEligibilityScheduleInstancesItemRoleDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property roleEligibilityScheduleInstances for roleManagement
// returns a *RequestInformation when successful
func (m *DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an unifiedRoleEligibilityScheduleInstance object.
// returns a *RequestInformation when successful
func (m *DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property roleEligibilityScheduleInstances in roleManagement
// returns a *RequestInformation when successful
func (m *DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleInstanceable, requestConfiguration *DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder when successful
func (m *DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) WithUrl(rawUrl string)(*DirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) {
    return NewDirectoryRoleEligibilityScheduleInstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
