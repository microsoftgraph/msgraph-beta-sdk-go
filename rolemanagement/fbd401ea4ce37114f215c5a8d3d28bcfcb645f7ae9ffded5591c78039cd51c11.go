package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder provides operations to manage the roleEligibilitySchedules property of the microsoft.graph.rbacApplication entity.
type EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderGetQueryParameters get roleEligibilitySchedules from roleManagement
type EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderGetQueryParameters
}
// EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppScope provides operations to manage the appScope property of the microsoft.graph.unifiedRoleScheduleBase entity.
// returns a *EnterpriseAppsItemRoleEligibilitySchedulesItemAppScopeRequestBuilder when successful
func (m *EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) AppScope()(*EnterpriseAppsItemRoleEligibilitySchedulesItemAppScopeRequestBuilder) {
    return NewEnterpriseAppsItemRoleEligibilitySchedulesItemAppScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderInternal instantiates a new EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder and sets the default values.
func NewEnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) {
    m := &EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/enterpriseApps/{rbacApplication%2Did}/roleEligibilitySchedules/{unifiedRoleEligibilitySchedule%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder instantiates a new EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder and sets the default values.
func NewEnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property roleEligibilitySchedules for roleManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DirectoryScope provides operations to manage the directoryScope property of the microsoft.graph.unifiedRoleScheduleBase entity.
// returns a *EnterpriseAppsItemRoleEligibilitySchedulesItemDirectoryScopeRequestBuilder when successful
func (m *EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) DirectoryScope()(*EnterpriseAppsItemRoleEligibilitySchedulesItemDirectoryScopeRequestBuilder) {
    return NewEnterpriseAppsItemRoleEligibilitySchedulesItemDirectoryScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get roleEligibilitySchedules from roleManagement
// returns a UnifiedRoleEligibilityScheduleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleEligibilityScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleable), nil
}
// Patch update the navigation property roleEligibilitySchedules in roleManagement
// returns a UnifiedRoleEligibilityScheduleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleable, requestConfiguration *EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleEligibilityScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleable), nil
}
// Principal provides operations to manage the principal property of the microsoft.graph.unifiedRoleScheduleBase entity.
// returns a *EnterpriseAppsItemRoleEligibilitySchedulesItemPrincipalRequestBuilder when successful
func (m *EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) Principal()(*EnterpriseAppsItemRoleEligibilitySchedulesItemPrincipalRequestBuilder) {
    return NewEnterpriseAppsItemRoleEligibilitySchedulesItemPrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleDefinition provides operations to manage the roleDefinition property of the microsoft.graph.unifiedRoleScheduleBase entity.
// returns a *EnterpriseAppsItemRoleEligibilitySchedulesItemRoleDefinitionRequestBuilder when successful
func (m *EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) RoleDefinition()(*EnterpriseAppsItemRoleEligibilitySchedulesItemRoleDefinitionRequestBuilder) {
    return NewEnterpriseAppsItemRoleEligibilitySchedulesItemRoleDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property roleEligibilitySchedules for roleManagement
// returns a *RequestInformation when successful
func (m *EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get roleEligibilitySchedules from roleManagement
// returns a *RequestInformation when successful
func (m *EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property roleEligibilitySchedules in roleManagement
// returns a *RequestInformation when successful
func (m *EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleEligibilityScheduleable, requestConfiguration *EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder when successful
func (m *EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) WithUrl(rawUrl string)(*EnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) {
    return NewEnterpriseAppsItemRoleEligibilitySchedulesUnifiedRoleEligibilityScheduleItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
