package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointReportsRequestBuilder provides operations to manage the reports property of the microsoft.graph.virtualEndpoint entity.
type VirtualEndpointReportsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointReportsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointReportsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualEndpointReportsRequestBuilderGetQueryParameters cloud PC related reports.
type VirtualEndpointReportsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualEndpointReportsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointReportsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEndpointReportsRequestBuilderGetQueryParameters
}
// VirtualEndpointReportsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointReportsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointReportsRequestBuilderInternal instantiates a new VirtualEndpointReportsRequestBuilder and sets the default values.
func NewVirtualEndpointReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointReportsRequestBuilder) {
    m := &VirtualEndpointReportsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/reports{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewVirtualEndpointReportsRequestBuilder instantiates a new VirtualEndpointReportsRequestBuilder and sets the default values.
func NewVirtualEndpointReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property reports for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEndpointReportsRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualEndpointReportsRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *VirtualEndpointReportsExportJobsRequestBuilder when successful
func (m *VirtualEndpointReportsRequestBuilder) ExportJobs()(*VirtualEndpointReportsExportJobsRequestBuilder) {
    return NewVirtualEndpointReportsExportJobsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get cloud PC related reports.
// returns a CloudPcReportsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEndpointReportsRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointReportsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReportsable, error) {
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
// returns a *VirtualEndpointReportsGetActionStatusReportsRequestBuilder when successful
func (m *VirtualEndpointReportsRequestBuilder) GetActionStatusReports()(*VirtualEndpointReportsGetActionStatusReportsRequestBuilder) {
    return NewVirtualEndpointReportsGetActionStatusReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetCloudPcPerformanceReport provides operations to call the getCloudPcPerformanceReport method.
// returns a *VirtualEndpointReportsGetCloudPcPerformanceReportRequestBuilder when successful
func (m *VirtualEndpointReportsRequestBuilder) GetCloudPcPerformanceReport()(*VirtualEndpointReportsGetCloudPcPerformanceReportRequestBuilder) {
    return NewVirtualEndpointReportsGetCloudPcPerformanceReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetCloudPcRecommendationReports provides operations to call the getCloudPcRecommendationReports method.
// returns a *VirtualEndpointReportsGetCloudPcRecommendationReportsRequestBuilder when successful
func (m *VirtualEndpointReportsRequestBuilder) GetCloudPcRecommendationReports()(*VirtualEndpointReportsGetCloudPcRecommendationReportsRequestBuilder) {
    return NewVirtualEndpointReportsGetCloudPcRecommendationReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetConnectionQualityReports provides operations to call the getConnectionQualityReports method.
// returns a *VirtualEndpointReportsGetConnectionQualityReportsRequestBuilder when successful
func (m *VirtualEndpointReportsRequestBuilder) GetConnectionQualityReports()(*VirtualEndpointReportsGetConnectionQualityReportsRequestBuilder) {
    return NewVirtualEndpointReportsGetConnectionQualityReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetDailyAggregatedRemoteConnectionReports provides operations to call the getDailyAggregatedRemoteConnectionReports method.
// returns a *VirtualEndpointReportsGetDailyAggregatedRemoteConnectionReportsRequestBuilder when successful
func (m *VirtualEndpointReportsRequestBuilder) GetDailyAggregatedRemoteConnectionReports()(*VirtualEndpointReportsGetDailyAggregatedRemoteConnectionReportsRequestBuilder) {
    return NewVirtualEndpointReportsGetDailyAggregatedRemoteConnectionReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetFrontlineReport provides operations to call the getFrontlineReport method.
// returns a *VirtualEndpointReportsGetFrontlineReportRequestBuilder when successful
func (m *VirtualEndpointReportsRequestBuilder) GetFrontlineReport()(*VirtualEndpointReportsGetFrontlineReportRequestBuilder) {
    return NewVirtualEndpointReportsGetFrontlineReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetInaccessibleCloudPcReports provides operations to call the getInaccessibleCloudPcReports method.
// returns a *VirtualEndpointReportsGetInaccessibleCloudPcReportsRequestBuilder when successful
func (m *VirtualEndpointReportsRequestBuilder) GetInaccessibleCloudPcReports()(*VirtualEndpointReportsGetInaccessibleCloudPcReportsRequestBuilder) {
    return NewVirtualEndpointReportsGetInaccessibleCloudPcReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetRawRemoteConnectionReports provides operations to call the getRawRemoteConnectionReports method.
// returns a *VirtualEndpointReportsGetRawRemoteConnectionReportsRequestBuilder when successful
func (m *VirtualEndpointReportsRequestBuilder) GetRawRemoteConnectionReports()(*VirtualEndpointReportsGetRawRemoteConnectionReportsRequestBuilder) {
    return NewVirtualEndpointReportsGetRawRemoteConnectionReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetRealTimeRemoteConnectionLatencyWithCloudPcId provides operations to call the getRealTimeRemoteConnectionLatency method.
// returns a *VirtualEndpointReportsGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder when successful
func (m *VirtualEndpointReportsRequestBuilder) GetRealTimeRemoteConnectionLatencyWithCloudPcId(cloudPcId *string)(*VirtualEndpointReportsGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder) {
    return NewVirtualEndpointReportsGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, cloudPcId)
}
// GetRealTimeRemoteConnectionStatusWithCloudPcId provides operations to call the getRealTimeRemoteConnectionStatus method.
// returns a *VirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder when successful
func (m *VirtualEndpointReportsRequestBuilder) GetRealTimeRemoteConnectionStatusWithCloudPcId(cloudPcId *string)(*VirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder) {
    return NewVirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, cloudPcId)
}
// GetRemoteConnectionHistoricalReports provides operations to call the getRemoteConnectionHistoricalReports method.
// returns a *VirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilder when successful
func (m *VirtualEndpointReportsRequestBuilder) GetRemoteConnectionHistoricalReports()(*VirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilder) {
    return NewVirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetSharedUseLicenseUsageReport provides operations to call the getSharedUseLicenseUsageReport method.
// returns a *VirtualEndpointReportsGetSharedUseLicenseUsageReportRequestBuilder when successful
func (m *VirtualEndpointReportsRequestBuilder) GetSharedUseLicenseUsageReport()(*VirtualEndpointReportsGetSharedUseLicenseUsageReportRequestBuilder) {
    return NewVirtualEndpointReportsGetSharedUseLicenseUsageReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetTotalAggregatedRemoteConnectionReports provides operations to call the getTotalAggregatedRemoteConnectionReports method.
// returns a *VirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder when successful
func (m *VirtualEndpointReportsRequestBuilder) GetTotalAggregatedRemoteConnectionReports()(*VirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder) {
    return NewVirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property reports in deviceManagement
// returns a CloudPcReportsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEndpointReportsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReportsable, requestConfiguration *VirtualEndpointReportsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReportsable, error) {
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
// RetrieveBulkActionStatusReport provides operations to call the retrieveBulkActionStatusReport method.
// returns a *VirtualEndpointReportsRetrieveBulkActionStatusReportRequestBuilder when successful
func (m *VirtualEndpointReportsRequestBuilder) RetrieveBulkActionStatusReport()(*VirtualEndpointReportsRetrieveBulkActionStatusReportRequestBuilder) {
    return NewVirtualEndpointReportsRetrieveBulkActionStatusReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RetrieveConnectionQualityReports provides operations to call the retrieveConnectionQualityReports method.
// returns a *VirtualEndpointReportsRetrieveConnectionQualityReportsRequestBuilder when successful
func (m *VirtualEndpointReportsRequestBuilder) RetrieveConnectionQualityReports()(*VirtualEndpointReportsRetrieveConnectionQualityReportsRequestBuilder) {
    return NewVirtualEndpointReportsRetrieveConnectionQualityReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RetrieveCrossRegionDisasterRecoveryReport provides operations to call the retrieveCrossRegionDisasterRecoveryReport method.
// returns a *VirtualEndpointReportsRetrieveCrossRegionDisasterRecoveryReportRequestBuilder when successful
func (m *VirtualEndpointReportsRequestBuilder) RetrieveCrossRegionDisasterRecoveryReport()(*VirtualEndpointReportsRetrieveCrossRegionDisasterRecoveryReportRequestBuilder) {
    return NewVirtualEndpointReportsRetrieveCrossRegionDisasterRecoveryReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RetrieveFrontlineReports provides operations to call the retrieveFrontlineReports method.
// returns a *VirtualEndpointReportsRetrieveFrontlineReportsRequestBuilder when successful
func (m *VirtualEndpointReportsRequestBuilder) RetrieveFrontlineReports()(*VirtualEndpointReportsRetrieveFrontlineReportsRequestBuilder) {
    return NewVirtualEndpointReportsRetrieveFrontlineReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property reports for deviceManagement
// returns a *RequestInformation when successful
func (m *VirtualEndpointReportsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointReportsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *VirtualEndpointReportsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointReportsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *VirtualEndpointReportsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReportsable, requestConfiguration *VirtualEndpointReportsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualEndpointReportsRequestBuilder when successful
func (m *VirtualEndpointReportsRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointReportsRequestBuilder) {
    return NewVirtualEndpointReportsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
