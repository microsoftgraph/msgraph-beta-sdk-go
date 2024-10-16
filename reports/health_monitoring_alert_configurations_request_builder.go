package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/healthmonitoring"
)

// HealthMonitoringAlertConfigurationsRequestBuilder provides operations to manage the alertConfigurations property of the microsoft.graph.healthMonitoring.healthMonitoringRoot entity.
type HealthMonitoringAlertConfigurationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// HealthMonitoringAlertConfigurationsRequestBuilderGetQueryParameters get a list of the Microsoft Entra health monitoring alertConfiguration objects and their properties.
type HealthMonitoringAlertConfigurationsRequestBuilderGetQueryParameters struct {
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
// HealthMonitoringAlertConfigurationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HealthMonitoringAlertConfigurationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *HealthMonitoringAlertConfigurationsRequestBuilderGetQueryParameters
}
// HealthMonitoringAlertConfigurationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HealthMonitoringAlertConfigurationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAlertConfigurationId provides operations to manage the alertConfigurations property of the microsoft.graph.healthMonitoring.healthMonitoringRoot entity.
// returns a *HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilder when successful
func (m *HealthMonitoringAlertConfigurationsRequestBuilder) ByAlertConfigurationId(alertConfigurationId string)(*HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if alertConfigurationId != "" {
        urlTplParams["alertConfiguration%2Did"] = alertConfigurationId
    }
    return NewHealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewHealthMonitoringAlertConfigurationsRequestBuilderInternal instantiates a new HealthMonitoringAlertConfigurationsRequestBuilder and sets the default values.
func NewHealthMonitoringAlertConfigurationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HealthMonitoringAlertConfigurationsRequestBuilder) {
    m := &HealthMonitoringAlertConfigurationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/healthMonitoring/alertConfigurations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewHealthMonitoringAlertConfigurationsRequestBuilder instantiates a new HealthMonitoringAlertConfigurationsRequestBuilder and sets the default values.
func NewHealthMonitoringAlertConfigurationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HealthMonitoringAlertConfigurationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewHealthMonitoringAlertConfigurationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *HealthMonitoringAlertConfigurationsCountRequestBuilder when successful
func (m *HealthMonitoringAlertConfigurationsRequestBuilder) Count()(*HealthMonitoringAlertConfigurationsCountRequestBuilder) {
    return NewHealthMonitoringAlertConfigurationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the Microsoft Entra health monitoring alertConfiguration objects and their properties.
// returns a AlertConfigurationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/healthmonitoring-healthmonitoringroot-list-alertconfigurations?view=graph-rest-beta
func (m *HealthMonitoringAlertConfigurationsRequestBuilder) Get(ctx context.Context, requestConfiguration *HealthMonitoringAlertConfigurationsRequestBuilderGetRequestConfiguration)(ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.AlertConfigurationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.CreateAlertConfigurationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.AlertConfigurationCollectionResponseable), nil
}
// Post create new navigation property to alertConfigurations for reports
// returns a AlertConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *HealthMonitoringAlertConfigurationsRequestBuilder) Post(ctx context.Context, body ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.AlertConfigurationable, requestConfiguration *HealthMonitoringAlertConfigurationsRequestBuilderPostRequestConfiguration)(ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.AlertConfigurationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.CreateAlertConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.AlertConfigurationable), nil
}
// ToGetRequestInformation get a list of the Microsoft Entra health monitoring alertConfiguration objects and their properties.
// returns a *RequestInformation when successful
func (m *HealthMonitoringAlertConfigurationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *HealthMonitoringAlertConfigurationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to alertConfigurations for reports
// returns a *RequestInformation when successful
func (m *HealthMonitoringAlertConfigurationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.AlertConfigurationable, requestConfiguration *HealthMonitoringAlertConfigurationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *HealthMonitoringAlertConfigurationsRequestBuilder when successful
func (m *HealthMonitoringAlertConfigurationsRequestBuilder) WithUrl(rawUrl string)(*HealthMonitoringAlertConfigurationsRequestBuilder) {
    return NewHealthMonitoringAlertConfigurationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
