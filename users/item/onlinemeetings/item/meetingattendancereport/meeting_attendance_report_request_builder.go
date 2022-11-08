package meetingattendancereport

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i7c3f73ee97b46c517c39cf5f9ef35e96c7ea82da0fa085235533f44b8bfcd31a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onlinemeetings/item/meetingattendancereport/attendancerecords"
    i0bc4a0926c42fea04621acf009dbca85914d3eb39c6864d11dbba854e7394f41 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onlinemeetings/item/meetingattendancereport/attendancerecords/item"
)

// MeetingAttendanceReportRequestBuilder provides operations to manage the meetingAttendanceReport property of the microsoft.graph.onlineMeeting entity.
type MeetingAttendanceReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeetingAttendanceReportRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeetingAttendanceReportRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MeetingAttendanceReportRequestBuilderGetQueryParameters get the meetingAttendanceReport for an onlineMeeting. Each time an online meeting ends, an attendance report will be generated for that session.
type MeetingAttendanceReportRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeetingAttendanceReportRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeetingAttendanceReportRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeetingAttendanceReportRequestBuilderGetQueryParameters
}
// MeetingAttendanceReportRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeetingAttendanceReportRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AttendanceRecords provides operations to manage the attendanceRecords property of the microsoft.graph.meetingAttendanceReport entity.
func (m *MeetingAttendanceReportRequestBuilder) AttendanceRecords()(*i7c3f73ee97b46c517c39cf5f9ef35e96c7ea82da0fa085235533f44b8bfcd31a.AttendanceRecordsRequestBuilder) {
    return i7c3f73ee97b46c517c39cf5f9ef35e96c7ea82da0fa085235533f44b8bfcd31a.NewAttendanceRecordsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttendanceRecordsById provides operations to manage the attendanceRecords property of the microsoft.graph.meetingAttendanceReport entity.
func (m *MeetingAttendanceReportRequestBuilder) AttendanceRecordsById(id string)(*i0bc4a0926c42fea04621acf009dbca85914d3eb39c6864d11dbba854e7394f41.AttendanceRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attendanceRecord%2Did"] = id
    }
    return i0bc4a0926c42fea04621acf009dbca85914d3eb39c6864d11dbba854e7394f41.NewAttendanceRecordItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewMeetingAttendanceReportRequestBuilderInternal instantiates a new MeetingAttendanceReportRequestBuilder and sets the default values.
func NewMeetingAttendanceReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeetingAttendanceReportRequestBuilder) {
    m := &MeetingAttendanceReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/onlineMeetings/{onlineMeeting%2Did}/meetingAttendanceReport{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeetingAttendanceReportRequestBuilder instantiates a new MeetingAttendanceReportRequestBuilder and sets the default values.
func NewMeetingAttendanceReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeetingAttendanceReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeetingAttendanceReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property meetingAttendanceReport for users
func (m *MeetingAttendanceReportRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *MeetingAttendanceReportRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get the meetingAttendanceReport for an onlineMeeting. Each time an online meeting ends, an attendance report will be generated for that session.
func (m *MeetingAttendanceReportRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MeetingAttendanceReportRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property meetingAttendanceReport in users
func (m *MeetingAttendanceReportRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingAttendanceReportable, requestConfiguration *MeetingAttendanceReportRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property meetingAttendanceReport for users
func (m *MeetingAttendanceReportRequestBuilder) Delete(ctx context.Context, requestConfiguration *MeetingAttendanceReportRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get the meetingAttendanceReport for an onlineMeeting. Each time an online meeting ends, an attendance report will be generated for that session.
func (m *MeetingAttendanceReportRequestBuilder) Get(ctx context.Context, requestConfiguration *MeetingAttendanceReportRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingAttendanceReportable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMeetingAttendanceReportFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingAttendanceReportable), nil
}
// Patch update the navigation property meetingAttendanceReport in users
func (m *MeetingAttendanceReportRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingAttendanceReportable, requestConfiguration *MeetingAttendanceReportRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingAttendanceReportable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMeetingAttendanceReportFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingAttendanceReportable), nil
}
