package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/healthmonitoring"
)

// HealthMonitoringAlertsAlertItemRequestBuilder provides operations to manage the alerts property of the microsoft.graph.healthMonitoring.healthMonitoringRoot entity.
type HealthMonitoringAlertsAlertItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// HealthMonitoringAlertsAlertItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HealthMonitoringAlertsAlertItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// HealthMonitoringAlertsAlertItemRequestBuilderGetQueryParameters get alerts from reports
type HealthMonitoringAlertsAlertItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// HealthMonitoringAlertsAlertItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HealthMonitoringAlertsAlertItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *HealthMonitoringAlertsAlertItemRequestBuilderGetQueryParameters
}
// HealthMonitoringAlertsAlertItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HealthMonitoringAlertsAlertItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewHealthMonitoringAlertsAlertItemRequestBuilderInternal instantiates a new HealthMonitoringAlertsAlertItemRequestBuilder and sets the default values.
func NewHealthMonitoringAlertsAlertItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HealthMonitoringAlertsAlertItemRequestBuilder) {
    m := &HealthMonitoringAlertsAlertItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/healthMonitoring/alerts/{alert%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewHealthMonitoringAlertsAlertItemRequestBuilder instantiates a new HealthMonitoringAlertsAlertItemRequestBuilder and sets the default values.
func NewHealthMonitoringAlertsAlertItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HealthMonitoringAlertsAlertItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewHealthMonitoringAlertsAlertItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property alerts for reports
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *HealthMonitoringAlertsAlertItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *HealthMonitoringAlertsAlertItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get alerts from reports
// returns a Alertable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *HealthMonitoringAlertsAlertItemRequestBuilder) Get(ctx context.Context, requestConfiguration *HealthMonitoringAlertsAlertItemRequestBuilderGetRequestConfiguration)(ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.Alertable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property alerts in reports
// returns a Alertable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *HealthMonitoringAlertsAlertItemRequestBuilder) Patch(ctx context.Context, body ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.Alertable, requestConfiguration *HealthMonitoringAlertsAlertItemRequestBuilderPatchRequestConfiguration)(ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.Alertable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property alerts for reports
// returns a *RequestInformation when successful
func (m *HealthMonitoringAlertsAlertItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *HealthMonitoringAlertsAlertItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get alerts from reports
// returns a *RequestInformation when successful
func (m *HealthMonitoringAlertsAlertItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *HealthMonitoringAlertsAlertItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property alerts in reports
// returns a *RequestInformation when successful
func (m *HealthMonitoringAlertsAlertItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.Alertable, requestConfiguration *HealthMonitoringAlertsAlertItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *HealthMonitoringAlertsAlertItemRequestBuilder when successful
func (m *HealthMonitoringAlertsAlertItemRequestBuilder) WithUrl(rawUrl string)(*HealthMonitoringAlertsAlertItemRequestBuilder) {
    return NewHealthMonitoringAlertsAlertItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
