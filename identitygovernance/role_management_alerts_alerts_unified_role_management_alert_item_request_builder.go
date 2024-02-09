package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilder provides operations to manage the alerts property of the microsoft.graph.roleManagementAlert entity.
type RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilderGetQueryParameters represents the alert entity.
type RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilderGetQueryParameters
}
// RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AlertConfiguration provides operations to manage the alertConfiguration property of the microsoft.graph.unifiedRoleManagementAlert entity.
// returns a *RoleManagementAlertsAlertsItemAlertConfigurationRequestBuilder when successful
func (m *RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilder) AlertConfiguration()(*RoleManagementAlertsAlertsItemAlertConfigurationRequestBuilder) {
    return NewRoleManagementAlertsAlertsItemAlertConfigurationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AlertDefinition provides operations to manage the alertDefinition property of the microsoft.graph.unifiedRoleManagementAlert entity.
// returns a *RoleManagementAlertsAlertsItemAlertDefinitionRequestBuilder when successful
func (m *RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilder) AlertDefinition()(*RoleManagementAlertsAlertsItemAlertDefinitionRequestBuilder) {
    return NewRoleManagementAlertsAlertsItemAlertDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AlertIncidents provides operations to manage the alertIncidents property of the microsoft.graph.unifiedRoleManagementAlert entity.
// returns a *RoleManagementAlertsAlertsItemAlertIncidentsRequestBuilder when successful
func (m *RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilder) AlertIncidents()(*RoleManagementAlertsAlertsItemAlertIncidentsRequestBuilder) {
    return NewRoleManagementAlertsAlertsItemAlertIncidentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewRoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilderInternal instantiates a new RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilder and sets the default values.
func NewRoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilder) {
    m := &RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/roleManagementAlerts/alerts/{unifiedRoleManagementAlert%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewRoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilder instantiates a new RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilder and sets the default values.
func NewRoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property alerts for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get represents the alert entity.
// returns a UnifiedRoleManagementAlertable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleManagementAlertFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertable), nil
}
// Patch update the navigation property alerts in identityGovernance
// returns a UnifiedRoleManagementAlertable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertable, requestConfiguration *RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleManagementAlertFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertable), nil
}
// Refresh provides operations to call the refresh method.
// returns a *RoleManagementAlertsAlertsItemRefreshRequestBuilder when successful
func (m *RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilder) Refresh()(*RoleManagementAlertsAlertsItemRefreshRequestBuilder) {
    return NewRoleManagementAlertsAlertsItemRefreshRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property alerts for identityGovernance
// returns a *RequestInformation when successful
func (m *RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/identityGovernance/roleManagementAlerts/alerts/{unifiedRoleManagementAlert%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation represents the alert entity.
// returns a *RequestInformation when successful
func (m *RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property alerts in identityGovernance
// returns a *RequestInformation when successful
func (m *RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertable, requestConfiguration *RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/identityGovernance/roleManagementAlerts/alerts/{unifiedRoleManagementAlert%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilder when successful
func (m *RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilder) WithUrl(rawUrl string)(*RoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilder) {
    return NewRoleManagementAlertsAlertsUnifiedRoleManagementAlertItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
