package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/healthmonitoring"
)

// HealthMonitoringRequestBuilder provides operations to manage the healthMonitoring property of the microsoft.graph.reportRoot entity.
type HealthMonitoringRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// HealthMonitoringRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HealthMonitoringRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// HealthMonitoringRequestBuilderGetQueryParameters reports for Microsoft Entra Health Monitoring.
type HealthMonitoringRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// HealthMonitoringRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HealthMonitoringRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *HealthMonitoringRequestBuilderGetQueryParameters
}
// HealthMonitoringRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HealthMonitoringRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AlertConfigurations provides operations to manage the alertConfigurations property of the microsoft.graph.healthMonitoring.healthMonitoringRoot entity.
// returns a *HealthMonitoringAlertConfigurationsRequestBuilder when successful
func (m *HealthMonitoringRequestBuilder) AlertConfigurations()(*HealthMonitoringAlertConfigurationsRequestBuilder) {
    return NewHealthMonitoringAlertConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Alerts provides operations to manage the alerts property of the microsoft.graph.healthMonitoring.healthMonitoringRoot entity.
// returns a *HealthMonitoringAlertsRequestBuilder when successful
func (m *HealthMonitoringRequestBuilder) Alerts()(*HealthMonitoringAlertsRequestBuilder) {
    return NewHealthMonitoringAlertsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewHealthMonitoringRequestBuilderInternal instantiates a new HealthMonitoringRequestBuilder and sets the default values.
func NewHealthMonitoringRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HealthMonitoringRequestBuilder) {
    m := &HealthMonitoringRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/healthMonitoring{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewHealthMonitoringRequestBuilder instantiates a new HealthMonitoringRequestBuilder and sets the default values.
func NewHealthMonitoringRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HealthMonitoringRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewHealthMonitoringRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property healthMonitoring for reports
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *HealthMonitoringRequestBuilder) Delete(ctx context.Context, requestConfiguration *HealthMonitoringRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get reports for Microsoft Entra Health Monitoring.
// returns a HealthMonitoringRootable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *HealthMonitoringRequestBuilder) Get(ctx context.Context, requestConfiguration *HealthMonitoringRequestBuilderGetRequestConfiguration)(ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.HealthMonitoringRootable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.CreateHealthMonitoringRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.HealthMonitoringRootable), nil
}
// Patch update the navigation property healthMonitoring in reports
// returns a HealthMonitoringRootable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *HealthMonitoringRequestBuilder) Patch(ctx context.Context, body ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.HealthMonitoringRootable, requestConfiguration *HealthMonitoringRequestBuilderPatchRequestConfiguration)(ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.HealthMonitoringRootable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.CreateHealthMonitoringRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.HealthMonitoringRootable), nil
}
// ToDeleteRequestInformation delete navigation property healthMonitoring for reports
// returns a *RequestInformation when successful
func (m *HealthMonitoringRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *HealthMonitoringRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation reports for Microsoft Entra Health Monitoring.
// returns a *RequestInformation when successful
func (m *HealthMonitoringRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *HealthMonitoringRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property healthMonitoring in reports
// returns a *RequestInformation when successful
func (m *HealthMonitoringRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.HealthMonitoringRootable, requestConfiguration *HealthMonitoringRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *HealthMonitoringRequestBuilder when successful
func (m *HealthMonitoringRequestBuilder) WithUrl(rawUrl string)(*HealthMonitoringRequestBuilder) {
    return NewHealthMonitoringRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
