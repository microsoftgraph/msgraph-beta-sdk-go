package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilder provides operations to manage the alertIncidents property of the microsoft.graph.unifiedRoleManagementAlert entity.
type RolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilderGetQueryParameters represents the incidents of this type of alert that have been triggered in Privileged Identity Management (PIM) for Microsoft Entra roles in the tenant. Supports $expand.
type RolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// RolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilderGetQueryParameters
}
// RolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUnifiedRoleManagementAlertIncidentId provides operations to manage the alertIncidents property of the microsoft.graph.unifiedRoleManagementAlert entity.
// returns a *RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder when successful
func (m *RolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilder) ByUnifiedRoleManagementAlertIncidentId(unifiedRoleManagementAlertIncidentId string)(*RolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if unifiedRoleManagementAlertIncidentId != "" {
        urlTplParams["unifiedRoleManagementAlertIncident%2Did"] = unifiedRoleManagementAlertIncidentId
    }
    return NewRolemanagementalertsAlertsItemAlertincidentsUnifiedRoleManagementAlertIncidentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewRolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilderInternal instantiates a new RolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilder and sets the default values.
func NewRolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilder) {
    m := &RolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/roleManagementAlerts/alerts/{unifiedRoleManagementAlert%2Did}/alertIncidents{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewRolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilder instantiates a new RolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilder and sets the default values.
func NewRolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *RolemanagementalertsAlertsItemAlertincidentsCountRequestBuilder when successful
func (m *RolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilder) Count()(*RolemanagementalertsAlertsItemAlertincidentsCountRequestBuilder) {
    return NewRolemanagementalertsAlertsItemAlertincidentsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get represents the incidents of this type of alert that have been triggered in Privileged Identity Management (PIM) for Microsoft Entra roles in the tenant. Supports $expand.
// returns a UnifiedRoleManagementAlertIncidentCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilder) Get(ctx context.Context, requestConfiguration *RolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertIncidentCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleManagementAlertIncidentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertIncidentCollectionResponseable), nil
}
// Post create new navigation property to alertIncidents for identityGovernance
// returns a UnifiedRoleManagementAlertIncidentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertIncidentable, requestConfiguration *RolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertIncidentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleManagementAlertIncidentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertIncidentable), nil
}
// ToGetRequestInformation represents the incidents of this type of alert that have been triggered in Privileged Identity Management (PIM) for Microsoft Entra roles in the tenant. Supports $expand.
// returns a *RequestInformation when successful
func (m *RolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to alertIncidents for identityGovernance
// returns a *RequestInformation when successful
func (m *RolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertIncidentable, requestConfiguration *RolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *RolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilder when successful
func (m *RolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilder) WithUrl(rawUrl string)(*RolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilder) {
    return NewRolemanagementalertsAlertsItemAlertincidentsAlertIncidentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
