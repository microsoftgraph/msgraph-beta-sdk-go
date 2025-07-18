// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/healthmonitoring"
)

// HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilder provides operations to manage the alertConfigurations property of the microsoft.graph.healthMonitoring.healthMonitoringRoot entity.
type HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilderGetQueryParameters read the properties and relationships of a Microsoft Entra health monitoring alertConfiguration object. The returned alertConfiguration object contains the settings for the distribution groups where alert notifications are to be sent.
type HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilderGetQueryParameters
}
// HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewHealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilderInternal instantiates a new HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilder and sets the default values.
func NewHealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilder) {
    m := &HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/healthMonitoring/alertConfigurations/{alertConfiguration%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewHealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilder instantiates a new HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilder and sets the default values.
func NewHealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewHealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property alertConfigurations for reports
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a Microsoft Entra health monitoring alertConfiguration object. The returned alertConfiguration object contains the settings for the distribution groups where alert notifications are to be sent.
// returns a AlertConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/healthmonitoring-alertconfiguration-get?view=graph-rest-beta
func (m *HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilderGetRequestConfiguration)(ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.AlertConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the properties of a Microsoft Entra health monitoring alertConfiguration object. You can use alertConfiguration settings to specify the distribution groups where alert notifications are to be sent. This API doesn't currently support group validation.
// returns a AlertConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/healthmonitoring-alertconfiguration-update?view=graph-rest-beta
func (m *HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilder) Patch(ctx context.Context, body ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.AlertConfigurationable, requestConfiguration *HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilderPatchRequestConfiguration)(ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.AlertConfigurationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property alertConfigurations for reports
// returns a *RequestInformation when successful
func (m *HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a Microsoft Entra health monitoring alertConfiguration object. The returned alertConfiguration object contains the settings for the distribution groups where alert notifications are to be sent.
// returns a *RequestInformation when successful
func (m *HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a Microsoft Entra health monitoring alertConfiguration object. You can use alertConfiguration settings to specify the distribution groups where alert notifications are to be sent. This API doesn't currently support group validation.
// returns a *RequestInformation when successful
func (m *HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ib1b876b25b5a2e04c31b7cc3c9de83e8af19c080f01f39f172f92708d7dc7450.AlertConfigurationable, requestConfiguration *HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilder when successful
func (m *HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilder) WithUrl(rawUrl string)(*HealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilder) {
    return NewHealthMonitoringAlertConfigurationsAlertConfigurationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
