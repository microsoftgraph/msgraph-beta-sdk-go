package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilder provides operations to manage the attendanceRecords property of the microsoft.graph.meetingAttendanceReport entity.
type CommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilderGetQueryParameters list of attendance records of an attendance report. Read-only.
type CommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilderGetQueryParameters
}
// CommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilderInternal instantiates a new AttendanceRecordItemRequestBuilder and sets the default values.
func NewCommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilder) {
    m := &CommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/onlineMeetings/{onlineMeeting%2Did}/meetingAttendanceReport/attendanceRecords/{attendanceRecord%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilder instantiates a new AttendanceRecordItemRequestBuilder and sets the default values.
func NewCommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property attendanceRecords for communications
func (m *CommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *CommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation list of attendance records of an attendance report. Read-only.
func (m *CommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *CommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property attendanceRecords in communications
func (m *CommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendanceRecordable, requestConfiguration *CommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property attendanceRecords for communications
func (m *CommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get list of attendance records of an attendance report. Read-only.
func (m *CommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendanceRecordable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAttendanceRecordFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendanceRecordable), nil
}
// Patch update the navigation property attendanceRecords in communications
func (m *CommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendanceRecordable, requestConfiguration *CommunicationsOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendanceRecordable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAttendanceRecordFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendanceRecordable), nil
}
