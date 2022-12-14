package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointReportsRequestBuilder provides operations to manage the reports property of the microsoft.graph.virtualEndpoint entity.
type VirtualEndpointReportsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
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
// NewVirtualEndpointReportsRequestBuilderInternal instantiates a new ReportsRequestBuilder and sets the default values.
func NewVirtualEndpointReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointReportsRequestBuilder) {
    m := &VirtualEndpointReportsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/virtualEndpoint/reports{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewVirtualEndpointReportsRequestBuilder instantiates a new ReportsRequestBuilder and sets the default values.
func NewVirtualEndpointReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property reports for deviceManagement
func (m *VirtualEndpointReportsRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointReportsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation cloud PC related reports.
func (m *VirtualEndpointReportsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointReportsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property reports in deviceManagement
func (m *VirtualEndpointReportsRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReportsable, requestConfiguration *VirtualEndpointReportsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property reports for deviceManagement
func (m *VirtualEndpointReportsRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualEndpointReportsRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ExportJobs provides operations to manage the exportJobs property of the microsoft.graph.cloudPcReports entity.
func (m *VirtualEndpointReportsRequestBuilder) ExportJobs()(*VirtualEndpointReportsExportJobsRequestBuilder) {
    return NewVirtualEndpointReportsExportJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportJobsById provides operations to manage the exportJobs property of the microsoft.graph.cloudPcReports entity.
func (m *VirtualEndpointReportsRequestBuilder) ExportJobsById(id string)(*VirtualEndpointReportsExportJobsCloudPcExportJobItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcExportJob%2Did"] = id
    }
    return NewVirtualEndpointReportsExportJobsCloudPcExportJobItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get cloud PC related reports.
func (m *VirtualEndpointReportsRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointReportsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReportsable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcReportsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReportsable), nil
}
// GetDailyAggregatedRemoteConnectionReports provides operations to call the getDailyAggregatedRemoteConnectionReports method.
func (m *VirtualEndpointReportsRequestBuilder) GetDailyAggregatedRemoteConnectionReports()(*VirtualEndpointReportsGetDailyAggregatedRemoteConnectionReportsRequestBuilder) {
    return NewVirtualEndpointReportsGetDailyAggregatedRemoteConnectionReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetRealTimeRemoteConnectionLatencyWithCloudPcId provides operations to call the getRealTimeRemoteConnectionLatency method.
func (m *VirtualEndpointReportsRequestBuilder) GetRealTimeRemoteConnectionLatencyWithCloudPcId(cloudPcId *string)(*VirtualEndpointReportsGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder) {
    return NewVirtualEndpointReportsGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, cloudPcId);
}
// GetRealTimeRemoteConnectionStatusWithCloudPcId provides operations to call the getRealTimeRemoteConnectionStatus method.
func (m *VirtualEndpointReportsRequestBuilder) GetRealTimeRemoteConnectionStatusWithCloudPcId(cloudPcId *string)(*VirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder) {
    return NewVirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, cloudPcId);
}
// GetRemoteConnectionHistoricalReports provides operations to call the getRemoteConnectionHistoricalReports method.
func (m *VirtualEndpointReportsRequestBuilder) GetRemoteConnectionHistoricalReports()(*VirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilder) {
    return NewVirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetSharedUseLicenseUsageReport provides operations to call the getSharedUseLicenseUsageReport method.
func (m *VirtualEndpointReportsRequestBuilder) GetSharedUseLicenseUsageReport()(*VirtualEndpointReportsGetSharedUseLicenseUsageReportRequestBuilder) {
    return NewVirtualEndpointReportsGetSharedUseLicenseUsageReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetTotalAggregatedRemoteConnectionReports provides operations to call the getTotalAggregatedRemoteConnectionReports method.
func (m *VirtualEndpointReportsRequestBuilder) GetTotalAggregatedRemoteConnectionReports()(*VirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder) {
    return NewVirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property reports in deviceManagement
func (m *VirtualEndpointReportsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReportsable, requestConfiguration *VirtualEndpointReportsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReportsable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcReportsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReportsable), nil
}
