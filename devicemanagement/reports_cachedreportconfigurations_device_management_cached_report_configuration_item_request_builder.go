package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilder provides operations to manage the cachedReportConfigurations property of the microsoft.graph.deviceManagementReports entity.
type ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilderGetQueryParameters entity representing the configuration of a cached report.
type ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilderGetQueryParameters
}
// ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilderInternal instantiates a new ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilder and sets the default values.
func NewReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilder) {
    m := &ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/reports/cachedReportConfigurations/{deviceManagementCachedReportConfiguration%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilder instantiates a new ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilder and sets the default values.
func NewReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property cachedReportConfigurations for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get entity representing the configuration of a cached report.
// returns a DeviceManagementCachedReportConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementCachedReportConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementCachedReportConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementCachedReportConfigurationable), nil
}
// Patch update the navigation property cachedReportConfigurations in deviceManagement
// returns a DeviceManagementCachedReportConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementCachedReportConfigurationable, requestConfiguration *ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementCachedReportConfigurationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementCachedReportConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementCachedReportConfigurationable), nil
}
// ToDeleteRequestInformation delete navigation property cachedReportConfigurations for deviceManagement
// returns a *RequestInformation when successful
func (m *ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation entity representing the configuration of a cached report.
// returns a *RequestInformation when successful
func (m *ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property cachedReportConfigurations in deviceManagement
// returns a *RequestInformation when successful
func (m *ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementCachedReportConfigurationable, requestConfiguration *ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilder when successful
func (m *ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilder) WithUrl(rawUrl string)(*ReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilder) {
    return NewReportsCachedreportconfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
