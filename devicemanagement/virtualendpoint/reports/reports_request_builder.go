package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i20cc57880f960babffb3a64a31f837f82335cd20188a8bbaea1d5511fa20d9b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/reports/getrealtimeremoteconnectionlatencywithcloudpcid"
    i6c63539b2ea836874719dad8cf88a5233e911e2374f162838afedeb7a9351a84 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/reports/getremoteconnectionhistoricalreports"
    i95b7c8af0745b2a1da9f540d5a777c99385bcebf02b4080501f89dbd40107ba3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/reports/getdailyaggregatedremoteconnectionreports"
    ie187a9a51e2529ea4c837ede8a8e6f91dfdb9cfe062f4562b9c42cd6c2d7385f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/reports/gettotalaggregatedremoteconnectionreports"
    ie41435632a5520fecdaee112e828e6671b343b789c3194f59d9f0f509c06ef75 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/reports/exportjobs"
    if50a57cae97cf86969f33dc48050fe5aeee7579c3225b782f7c05a91663bb493 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/reports/getrealtimeremoteconnectionstatuswithcloudpcid"
    ie2ea01ca6f1a0dce0c2cb6a13a16a43889baea0dedb63d5685be60340e932e58 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/reports/exportjobs/item"
)

// ReportsRequestBuilder provides operations to manage the reports property of the microsoft.graph.virtualEndpoint entity.
type ReportsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ReportsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ReportsRequestBuilderGetQueryParameters get reports from deviceManagement
type ReportsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ReportsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ReportsRequestBuilderGetQueryParameters
}
// ReportsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReportsRequestBuilderInternal instantiates a new ReportsRequestBuilder and sets the default values.
func NewReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsRequestBuilder) {
    m := &ReportsRequestBuilder{
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
// NewReportsRequestBuilder instantiates a new ReportsRequestBuilder and sets the default values.
func NewReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property reports for deviceManagement
func (m *ReportsRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property reports for deviceManagement
func (m *ReportsRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *ReportsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get reports from deviceManagement
func (m *ReportsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get reports from deviceManagement
func (m *ReportsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ReportsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property reports in deviceManagement
func (m *ReportsRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReportsable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property reports in deviceManagement
func (m *ReportsRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReportsable, requestConfiguration *ReportsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property reports for deviceManagement
func (m *ReportsRequestBuilder) Delete(ctx context.Context, requestConfiguration *ReportsRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
// ExportJobs the exportJobs property
func (m *ReportsRequestBuilder) ExportJobs()(*ie41435632a5520fecdaee112e828e6671b343b789c3194f59d9f0f509c06ef75.ExportJobsRequestBuilder) {
    return ie41435632a5520fecdaee112e828e6671b343b789c3194f59d9f0f509c06ef75.NewExportJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportJobsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.virtualEndpoint.reports.exportJobs.item collection
func (m *ReportsRequestBuilder) ExportJobsById(id string)(*ie2ea01ca6f1a0dce0c2cb6a13a16a43889baea0dedb63d5685be60340e932e58.CloudPcExportJobItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcExportJob%2Did"] = id
    }
    return ie2ea01ca6f1a0dce0c2cb6a13a16a43889baea0dedb63d5685be60340e932e58.NewCloudPcExportJobItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get reports from deviceManagement
func (m *ReportsRequestBuilder) Get(ctx context.Context, requestConfiguration *ReportsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReportsable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
// GetDailyAggregatedRemoteConnectionReports the getDailyAggregatedRemoteConnectionReports property
func (m *ReportsRequestBuilder) GetDailyAggregatedRemoteConnectionReports()(*i95b7c8af0745b2a1da9f540d5a777c99385bcebf02b4080501f89dbd40107ba3.GetDailyAggregatedRemoteConnectionReportsRequestBuilder) {
    return i95b7c8af0745b2a1da9f540d5a777c99385bcebf02b4080501f89dbd40107ba3.NewGetDailyAggregatedRemoteConnectionReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetRealTimeRemoteConnectionLatencyWithCloudPcId provides operations to call the getRealTimeRemoteConnectionLatency method.
func (m *ReportsRequestBuilder) GetRealTimeRemoteConnectionLatencyWithCloudPcId(cloudPcId *string)(*i20cc57880f960babffb3a64a31f837f82335cd20188a8bbaea1d5511fa20d9b6.GetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder) {
    return i20cc57880f960babffb3a64a31f837f82335cd20188a8bbaea1d5511fa20d9b6.NewGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, cloudPcId);
}
// GetRealTimeRemoteConnectionStatusWithCloudPcId provides operations to call the getRealTimeRemoteConnectionStatus method.
func (m *ReportsRequestBuilder) GetRealTimeRemoteConnectionStatusWithCloudPcId(cloudPcId *string)(*if50a57cae97cf86969f33dc48050fe5aeee7579c3225b782f7c05a91663bb493.GetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder) {
    return if50a57cae97cf86969f33dc48050fe5aeee7579c3225b782f7c05a91663bb493.NewGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, cloudPcId);
}
// GetRemoteConnectionHistoricalReports the getRemoteConnectionHistoricalReports property
func (m *ReportsRequestBuilder) GetRemoteConnectionHistoricalReports()(*i6c63539b2ea836874719dad8cf88a5233e911e2374f162838afedeb7a9351a84.GetRemoteConnectionHistoricalReportsRequestBuilder) {
    return i6c63539b2ea836874719dad8cf88a5233e911e2374f162838afedeb7a9351a84.NewGetRemoteConnectionHistoricalReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetTotalAggregatedRemoteConnectionReports the getTotalAggregatedRemoteConnectionReports property
func (m *ReportsRequestBuilder) GetTotalAggregatedRemoteConnectionReports()(*ie187a9a51e2529ea4c837ede8a8e6f91dfdb9cfe062f4562b9c42cd6c2d7385f.GetTotalAggregatedRemoteConnectionReportsRequestBuilder) {
    return ie187a9a51e2529ea4c837ede8a8e6f91dfdb9cfe062f4562b9c42cd6c2d7385f.NewGetTotalAggregatedRemoteConnectionReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property reports in deviceManagement
func (m *ReportsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReportsable, requestConfiguration *ReportsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReportsable, error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
