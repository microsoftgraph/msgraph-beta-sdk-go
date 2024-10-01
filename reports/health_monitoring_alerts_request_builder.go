package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/healthmonitoring"
)

// HealthMonitoringAlertsRequestBuilder provides operations to manage the alerts property of the microsoft.graph.healthMonitoring.healthMonitoringRoot entity.
type HealthMonitoringAlertsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// HealthMonitoringAlertsRequestBuilderGetQueryParameters get alerts from reports
type HealthMonitoringAlertsRequestBuilderGetQueryParameters struct {
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
// HealthMonitoringAlertsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HealthMonitoringAlertsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *HealthMonitoringAlertsRequestBuilderGetQueryParameters
}
// HealthMonitoringAlertsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HealthMonitoringAlertsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAlertId provides operations to manage the alerts property of the microsoft.graph.healthMonitoring.healthMonitoringRoot entity.
// returns a *HealthMonitoringAlertsAlertItemRequestBuilder when successful
func (m *HealthMonitoringAlertsRequestBuilder) ByAlertId(alertId string)(*HealthMonitoringAlertsAlertItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if alertId != "" {
        urlTplParams["alert%2Did"] = alertId
    }
    return NewHealthMonitoringAlertsAlertItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewHealthMonitoringAlertsRequestBuilderInternal instantiates a new HealthMonitoringAlertsRequestBuilder and sets the default values.
func NewHealthMonitoringAlertsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HealthMonitoringAlertsRequestBuilder) {
    m := &HealthMonitoringAlertsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/healthMonitoring/alerts{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewHealthMonitoringAlertsRequestBuilder instantiates a new HealthMonitoringAlertsRequestBuilder and sets the default values.
func NewHealthMonitoringAlertsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HealthMonitoringAlertsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewHealthMonitoringAlertsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *HealthMonitoringAlertsCountRequestBuilder when successful
func (m *HealthMonitoringAlertsRequestBuilder) Count()(*HealthMonitoringAlertsCountRequestBuilder) {
    return NewHealthMonitoringAlertsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get alerts from reports
// returns a AlertCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *HealthMonitoringAlertsRequestBuilder) Get(ctx context.Context, requestConfiguration *HealthMonitoringAlertsRequestBuilderGetRequestConfiguration)(ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.AlertCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.CreateAlertCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.AlertCollectionResponseable), nil
}
// Post create new navigation property to alerts for reports
// returns a Alertable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *HealthMonitoringAlertsRequestBuilder) Post(ctx context.Context, body ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.Alertable, requestConfiguration *HealthMonitoringAlertsRequestBuilderPostRequestConfiguration)(ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.Alertable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.CreateAlertFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.Alertable), nil
}
// ToGetRequestInformation get alerts from reports
// returns a *RequestInformation when successful
func (m *HealthMonitoringAlertsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *HealthMonitoringAlertsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to alerts for reports
// returns a *RequestInformation when successful
func (m *HealthMonitoringAlertsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.Alertable, requestConfiguration *HealthMonitoringAlertsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *HealthMonitoringAlertsRequestBuilder when successful
func (m *HealthMonitoringAlertsRequestBuilder) WithUrl(rawUrl string)(*HealthMonitoringAlertsRequestBuilder) {
    return NewHealthMonitoringAlertsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
