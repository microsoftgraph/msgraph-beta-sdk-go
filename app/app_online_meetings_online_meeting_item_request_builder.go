package app

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AppOnlineMeetingsOnlineMeetingItemRequestBuilder provides operations to manage the onlineMeetings property of the microsoft.graph.commsApplication entity.
type AppOnlineMeetingsOnlineMeetingItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AppOnlineMeetingsOnlineMeetingItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppOnlineMeetingsOnlineMeetingItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppOnlineMeetingsOnlineMeetingItemRequestBuilderGetQueryParameters get onlineMeetings from app
type AppOnlineMeetingsOnlineMeetingItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AppOnlineMeetingsOnlineMeetingItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppOnlineMeetingsOnlineMeetingItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AppOnlineMeetingsOnlineMeetingItemRequestBuilderGetQueryParameters
}
// AppOnlineMeetingsOnlineMeetingItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppOnlineMeetingsOnlineMeetingItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AlternativeRecording provides operations to manage the media for the commsApplication entity.
func (m *AppOnlineMeetingsOnlineMeetingItemRequestBuilder) AlternativeRecording()(*AppOnlineMeetingsItemAlternativeRecordingRequestBuilder) {
    return NewAppOnlineMeetingsItemAlternativeRecordingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttendanceReports provides operations to manage the attendanceReports property of the microsoft.graph.onlineMeeting entity.
func (m *AppOnlineMeetingsOnlineMeetingItemRequestBuilder) AttendanceReports()(*AppOnlineMeetingsItemAttendanceReportsRequestBuilder) {
    return NewAppOnlineMeetingsItemAttendanceReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttendanceReportsById provides operations to manage the attendanceReports property of the microsoft.graph.onlineMeeting entity.
func (m *AppOnlineMeetingsOnlineMeetingItemRequestBuilder) AttendanceReportsById(id string)(*AppOnlineMeetingsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["meetingAttendanceReport%2Did"] = id
    }
    return NewAppOnlineMeetingsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AttendeeReport provides operations to manage the media for the commsApplication entity.
func (m *AppOnlineMeetingsOnlineMeetingItemRequestBuilder) AttendeeReport()(*AppOnlineMeetingsItemAttendeeReportRequestBuilder) {
    return NewAppOnlineMeetingsItemAttendeeReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAppOnlineMeetingsOnlineMeetingItemRequestBuilderInternal instantiates a new OnlineMeetingItemRequestBuilder and sets the default values.
func NewAppOnlineMeetingsOnlineMeetingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppOnlineMeetingsOnlineMeetingItemRequestBuilder) {
    m := &AppOnlineMeetingsOnlineMeetingItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/app/onlineMeetings/{onlineMeeting%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAppOnlineMeetingsOnlineMeetingItemRequestBuilder instantiates a new OnlineMeetingItemRequestBuilder and sets the default values.
func NewAppOnlineMeetingsOnlineMeetingItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppOnlineMeetingsOnlineMeetingItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppOnlineMeetingsOnlineMeetingItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property onlineMeetings for app
func (m *AppOnlineMeetingsOnlineMeetingItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *AppOnlineMeetingsOnlineMeetingItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get onlineMeetings from app
func (m *AppOnlineMeetingsOnlineMeetingItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *AppOnlineMeetingsOnlineMeetingItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property onlineMeetings in app
func (m *AppOnlineMeetingsOnlineMeetingItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable, requestConfiguration *AppOnlineMeetingsOnlineMeetingItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property onlineMeetings for app
func (m *AppOnlineMeetingsOnlineMeetingItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AppOnlineMeetingsOnlineMeetingItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get onlineMeetings from app
func (m *AppOnlineMeetingsOnlineMeetingItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AppOnlineMeetingsOnlineMeetingItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnlineMeetingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable), nil
}
// MeetingAttendanceReport provides operations to manage the meetingAttendanceReport property of the microsoft.graph.onlineMeeting entity.
func (m *AppOnlineMeetingsOnlineMeetingItemRequestBuilder) MeetingAttendanceReport()(*AppOnlineMeetingsItemMeetingAttendanceReportRequestBuilder) {
    return NewAppOnlineMeetingsItemMeetingAttendanceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property onlineMeetings in app
func (m *AppOnlineMeetingsOnlineMeetingItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable, requestConfiguration *AppOnlineMeetingsOnlineMeetingItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnlineMeetingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable), nil
}
// Recording provides operations to manage the media for the commsApplication entity.
func (m *AppOnlineMeetingsOnlineMeetingItemRequestBuilder) Recording()(*AppOnlineMeetingsItemRecordingRequestBuilder) {
    return NewAppOnlineMeetingsItemRecordingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Registration provides operations to manage the registration property of the microsoft.graph.onlineMeeting entity.
func (m *AppOnlineMeetingsOnlineMeetingItemRequestBuilder) Registration()(*AppOnlineMeetingsItemRegistrationRequestBuilder) {
    return NewAppOnlineMeetingsItemRegistrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Transcripts provides operations to manage the transcripts property of the microsoft.graph.onlineMeeting entity.
func (m *AppOnlineMeetingsOnlineMeetingItemRequestBuilder) Transcripts()(*AppOnlineMeetingsItemTranscriptsRequestBuilder) {
    return NewAppOnlineMeetingsItemTranscriptsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranscriptsById provides operations to manage the transcripts property of the microsoft.graph.onlineMeeting entity.
func (m *AppOnlineMeetingsOnlineMeetingItemRequestBuilder) TranscriptsById(id string)(*AppOnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["callTranscript%2Did"] = id
    }
    return NewAppOnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// VirtualAppointment provides operations to manage the virtualAppointment property of the microsoft.graph.onlineMeeting entity.
func (m *AppOnlineMeetingsOnlineMeetingItemRequestBuilder) VirtualAppointment()(*AppOnlineMeetingsItemVirtualAppointmentRequestBuilder) {
    return NewAppOnlineMeetingsItemVirtualAppointmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
