package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualendpointReportsRequestBuilder provides operations to manage the reports property of the microsoft.graph.virtualEndpoint entity.
type VirtualendpointReportsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointReportsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointReportsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualendpointReportsRequestBuilderGetQueryParameters cloud PC related reports.
type VirtualendpointReportsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualendpointReportsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointReportsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualendpointReportsRequestBuilderGetQueryParameters
}
// VirtualendpointReportsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointReportsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualendpointReportsRequestBuilderInternal instantiates a new VirtualendpointReportsRequestBuilder and sets the default values.
func NewVirtualendpointReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointReportsRequestBuilder) {
    m := &VirtualendpointReportsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/reports{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewVirtualendpointReportsRequestBuilder instantiates a new VirtualendpointReportsRequestBuilder and sets the default values.
func NewVirtualendpointReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property reports for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualendpointReportsRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualendpointReportsRequestBuilderDeleteRequestConfiguration)(error) {
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
// ExportJobs provides operations to manage the exportJobs property of the microsoft.graph.cloudPcReports entity.
// returns a *VirtualendpointReportsExportjobsExportJobsRequestBuilder when successful
func (m *VirtualendpointReportsRequestBuilder) ExportJobs()(*VirtualendpointReportsExportjobsExportJobsRequestBuilder) {
    return NewVirtualendpointReportsExportjobsExportJobsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get cloud PC related reports.
// returns a CloudPcReportsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualendpointReportsRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualendpointReportsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReportsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcReportsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReportsable), nil
}
// GetActionStatusReports provides operations to call the getActionStatusReports method.
// returns a *VirtualendpointReportsGetactionstatusreportsGetActionStatusReportsRequestBuilder when successful
func (m *VirtualendpointReportsRequestBuilder) GetActionStatusReports()(*VirtualendpointReportsGetactionstatusreportsGetActionStatusReportsRequestBuilder) {
    return NewVirtualendpointReportsGetactionstatusreportsGetActionStatusReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetCloudPcPerformanceReport provides operations to call the getCloudPcPerformanceReport method.
// returns a *VirtualendpointReportsGetcloudpcperformancereportGetCloudPcPerformanceReportRequestBuilder when successful
func (m *VirtualendpointReportsRequestBuilder) GetCloudPcPerformanceReport()(*VirtualendpointReportsGetcloudpcperformancereportGetCloudPcPerformanceReportRequestBuilder) {
    return NewVirtualendpointReportsGetcloudpcperformancereportGetCloudPcPerformanceReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetCloudPcRecommendationReports provides operations to call the getCloudPcRecommendationReports method.
// returns a *VirtualendpointReportsGetcloudpcrecommendationreportsGetCloudPcRecommendationReportsRequestBuilder when successful
func (m *VirtualendpointReportsRequestBuilder) GetCloudPcRecommendationReports()(*VirtualendpointReportsGetcloudpcrecommendationreportsGetCloudPcRecommendationReportsRequestBuilder) {
    return NewVirtualendpointReportsGetcloudpcrecommendationreportsGetCloudPcRecommendationReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetConnectionQualityReports provides operations to call the getConnectionQualityReports method.
// returns a *VirtualendpointReportsGetconnectionqualityreportsGetConnectionQualityReportsRequestBuilder when successful
func (m *VirtualendpointReportsRequestBuilder) GetConnectionQualityReports()(*VirtualendpointReportsGetconnectionqualityreportsGetConnectionQualityReportsRequestBuilder) {
    return NewVirtualendpointReportsGetconnectionqualityreportsGetConnectionQualityReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetDailyAggregatedRemoteConnectionReports provides operations to call the getDailyAggregatedRemoteConnectionReports method.
// returns a *VirtualendpointReportsGetdailyaggregatedremoteconnectionreportsGetDailyAggregatedRemoteConnectionReportsRequestBuilder when successful
func (m *VirtualendpointReportsRequestBuilder) GetDailyAggregatedRemoteConnectionReports()(*VirtualendpointReportsGetdailyaggregatedremoteconnectionreportsGetDailyAggregatedRemoteConnectionReportsRequestBuilder) {
    return NewVirtualendpointReportsGetdailyaggregatedremoteconnectionreportsGetDailyAggregatedRemoteConnectionReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetFrontlineReport provides operations to call the getFrontlineReport method.
// returns a *VirtualendpointReportsGetfrontlinereportGetFrontlineReportRequestBuilder when successful
func (m *VirtualendpointReportsRequestBuilder) GetFrontlineReport()(*VirtualendpointReportsGetfrontlinereportGetFrontlineReportRequestBuilder) {
    return NewVirtualendpointReportsGetfrontlinereportGetFrontlineReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetInaccessibleCloudPcReports provides operations to call the getInaccessibleCloudPcReports method.
// returns a *VirtualendpointReportsGetinaccessiblecloudpcreportsGetInaccessibleCloudPcReportsRequestBuilder when successful
func (m *VirtualendpointReportsRequestBuilder) GetInaccessibleCloudPcReports()(*VirtualendpointReportsGetinaccessiblecloudpcreportsGetInaccessibleCloudPcReportsRequestBuilder) {
    return NewVirtualendpointReportsGetinaccessiblecloudpcreportsGetInaccessibleCloudPcReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetRawRemoteConnectionReports provides operations to call the getRawRemoteConnectionReports method.
// returns a *VirtualendpointReportsGetrawremoteconnectionreportsGetRawRemoteConnectionReportsRequestBuilder when successful
func (m *VirtualendpointReportsRequestBuilder) GetRawRemoteConnectionReports()(*VirtualendpointReportsGetrawremoteconnectionreportsGetRawRemoteConnectionReportsRequestBuilder) {
    return NewVirtualendpointReportsGetrawremoteconnectionreportsGetRawRemoteConnectionReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetRealTimeRemoteConnectionLatencyWithCloudPcId provides operations to call the getRealTimeRemoteConnectionLatency method.
// returns a *VirtualendpointReportsGetrealtimeremoteconnectionlatencywithcloudpcidGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder when successful
func (m *VirtualendpointReportsRequestBuilder) GetRealTimeRemoteConnectionLatencyWithCloudPcId(cloudPcId *string)(*VirtualendpointReportsGetrealtimeremoteconnectionlatencywithcloudpcidGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder) {
    return NewVirtualendpointReportsGetrealtimeremoteconnectionlatencywithcloudpcidGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, cloudPcId)
}
// GetRealTimeRemoteConnectionStatusWithCloudPcId provides operations to call the getRealTimeRemoteConnectionStatus method.
// returns a *VirtualendpointReportsGetrealtimeremoteconnectionstatuswithcloudpcidGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder when successful
func (m *VirtualendpointReportsRequestBuilder) GetRealTimeRemoteConnectionStatusWithCloudPcId(cloudPcId *string)(*VirtualendpointReportsGetrealtimeremoteconnectionstatuswithcloudpcidGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder) {
    return NewVirtualendpointReportsGetrealtimeremoteconnectionstatuswithcloudpcidGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, cloudPcId)
}
// GetRemoteConnectionHistoricalReports provides operations to call the getRemoteConnectionHistoricalReports method.
// returns a *VirtualendpointReportsGetremoteconnectionhistoricalreportsGetRemoteConnectionHistoricalReportsRequestBuilder when successful
func (m *VirtualendpointReportsRequestBuilder) GetRemoteConnectionHistoricalReports()(*VirtualendpointReportsGetremoteconnectionhistoricalreportsGetRemoteConnectionHistoricalReportsRequestBuilder) {
    return NewVirtualendpointReportsGetremoteconnectionhistoricalreportsGetRemoteConnectionHistoricalReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetSharedUseLicenseUsageReport provides operations to call the getSharedUseLicenseUsageReport method.
// returns a *VirtualendpointReportsGetshareduselicenseusagereportGetSharedUseLicenseUsageReportRequestBuilder when successful
func (m *VirtualendpointReportsRequestBuilder) GetSharedUseLicenseUsageReport()(*VirtualendpointReportsGetshareduselicenseusagereportGetSharedUseLicenseUsageReportRequestBuilder) {
    return NewVirtualendpointReportsGetshareduselicenseusagereportGetSharedUseLicenseUsageReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetTotalAggregatedRemoteConnectionReports provides operations to call the getTotalAggregatedRemoteConnectionReports method.
// returns a *VirtualendpointReportsGettotalaggregatedremoteconnectionreportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder when successful
func (m *VirtualendpointReportsRequestBuilder) GetTotalAggregatedRemoteConnectionReports()(*VirtualendpointReportsGettotalaggregatedremoteconnectionreportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder) {
    return NewVirtualendpointReportsGettotalaggregatedremoteconnectionreportsGetTotalAggregatedRemoteConnectionReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property reports in deviceManagement
// returns a CloudPcReportsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualendpointReportsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReportsable, requestConfiguration *VirtualendpointReportsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReportsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcReportsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReportsable), nil
}
// RetrieveCrossRegionDisasterRecoveryReport provides operations to call the retrieveCrossRegionDisasterRecoveryReport method.
// returns a *VirtualendpointReportsRetrievecrossregiondisasterrecoveryreportRetrieveCrossRegionDisasterRecoveryReportRequestBuilder when successful
func (m *VirtualendpointReportsRequestBuilder) RetrieveCrossRegionDisasterRecoveryReport()(*VirtualendpointReportsRetrievecrossregiondisasterrecoveryreportRetrieveCrossRegionDisasterRecoveryReportRequestBuilder) {
    return NewVirtualendpointReportsRetrievecrossregiondisasterrecoveryreportRetrieveCrossRegionDisasterRecoveryReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property reports for deviceManagement
// returns a *RequestInformation when successful
func (m *VirtualendpointReportsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointReportsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation cloud PC related reports.
// returns a *RequestInformation when successful
func (m *VirtualendpointReportsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointReportsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property reports in deviceManagement
// returns a *RequestInformation when successful
func (m *VirtualendpointReportsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReportsable, requestConfiguration *VirtualendpointReportsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualendpointReportsRequestBuilder when successful
func (m *VirtualendpointReportsRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointReportsRequestBuilder) {
    return NewVirtualendpointReportsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
