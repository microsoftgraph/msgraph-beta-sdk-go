package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilder provides operations to manage the alertConfigurations property of the microsoft.graph.roleManagementAlert entity.
type RolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilderGetQueryParameters get a list of the alert configurations. The alert configurations are a collection of following types that are derived from the unifiedRoleManagementAlertConfiguration object:
type RolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilderGetQueryParameters struct {
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
// RolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilderGetQueryParameters
}
// RolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUnifiedRoleManagementAlertConfigurationId provides operations to manage the alertConfigurations property of the microsoft.graph.roleManagementAlert entity.
// returns a *RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder when successful
func (m *RolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilder) ByUnifiedRoleManagementAlertConfigurationId(unifiedRoleManagementAlertConfigurationId string)(*RolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if unifiedRoleManagementAlertConfigurationId != "" {
        urlTplParams["unifiedRoleManagementAlertConfiguration%2Did"] = unifiedRoleManagementAlertConfigurationId
    }
    return NewRolemanagementalertsAlertconfigurationsUnifiedRoleManagementAlertConfigurationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewRolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilderInternal instantiates a new RolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilder and sets the default values.
func NewRolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilder) {
    m := &RolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/roleManagementAlerts/alertConfigurations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewRolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilder instantiates a new RolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilder and sets the default values.
func NewRolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *RolemanagementalertsAlertconfigurationsCountRequestBuilder when successful
func (m *RolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilder) Count()(*RolemanagementalertsAlertconfigurationsCountRequestBuilder) {
    return NewRolemanagementalertsAlertconfigurationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the alert configurations. The alert configurations are a collection of following types that are derived from the unifiedRoleManagementAlertConfiguration object:
// returns a UnifiedRoleManagementAlertConfigurationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/rolemanagementalert-list-alertconfigurations?view=graph-rest-beta
func (m *RolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilder) Get(ctx context.Context, requestConfiguration *RolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertConfigurationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleManagementAlertConfigurationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertConfigurationCollectionResponseable), nil
}
// Post create new navigation property to alertConfigurations for identityGovernance
// returns a UnifiedRoleManagementAlertConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertConfigurationable, requestConfiguration *RolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertConfigurationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleManagementAlertConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertConfigurationable), nil
}
// ToGetRequestInformation get a list of the alert configurations. The alert configurations are a collection of following types that are derived from the unifiedRoleManagementAlertConfiguration object:
// returns a *RequestInformation when successful
func (m *RolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to alertConfigurations for identityGovernance
// returns a *RequestInformation when successful
func (m *RolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementAlertConfigurationable, requestConfiguration *RolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilder when successful
func (m *RolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilder) WithUrl(rawUrl string)(*RolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilder) {
    return NewRolemanagementalertsAlertconfigurationsAlertConfigurationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
