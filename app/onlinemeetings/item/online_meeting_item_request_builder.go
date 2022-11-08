package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i36d7f6a1b1906c3cdb16b4ede2175f668fa5388566026ae225a46e8d0be75ba4 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/onlinemeetings/item/attendancereports"
    i5792a57d1b509cf5853eddfc75cdd3d12d7fffac6c1d1ec273e69dfc55edec5d "github.com/microsoftgraph/msgraph-beta-sdk-go/app/onlinemeetings/item/attendeereport"
    i621490855f8aac21cf8d7032977ac3e4dd93569a78b9ec995379c44e1ed33eaf "github.com/microsoftgraph/msgraph-beta-sdk-go/app/onlinemeetings/item/alternativerecording"
    i67de322803fb9ceef33e79fb4b659d5e8d0ffaa90284f5ea1260109536f446ce "github.com/microsoftgraph/msgraph-beta-sdk-go/app/onlinemeetings/item/transcripts"
    ia44b3dfc614f798ddda4e107c7b00a962e62b26be06acb8ba2c30fda4607647a "github.com/microsoftgraph/msgraph-beta-sdk-go/app/onlinemeetings/item/meetingattendancereport"
    ic4cea0d5545c1ec81a2c193eb0ee5d7a032aa549a87e4b5ea7d5e3d4250137fc "github.com/microsoftgraph/msgraph-beta-sdk-go/app/onlinemeetings/item/recording"
    if386b4892e4921e066f9d8ba80ef001577fdf97ddb75e8a48e18973728537f9e "github.com/microsoftgraph/msgraph-beta-sdk-go/app/onlinemeetings/item/virtualappointment"
    ifc43ad7bd5674ca951e2271a8b4d99b63fa47b18f4fca9e422daedf0278acaf6 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/onlinemeetings/item/registration"
    i29fda42a3d65f7aed1e8a1e9c5006e031671b1aed329c71b8d3daa638210ed5a "github.com/microsoftgraph/msgraph-beta-sdk-go/app/onlinemeetings/item/attendancereports/item"
    ide6e7b6bb6c57b727e00bfe216300dcd9aa6cd3b52ecddbd9256ee94b3816d97 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/onlinemeetings/item/transcripts/item"
)

// OnlineMeetingItemRequestBuilder provides operations to manage the onlineMeetings property of the microsoft.graph.commsApplication entity.
type OnlineMeetingItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OnlineMeetingItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlineMeetingItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OnlineMeetingItemRequestBuilderGetQueryParameters get onlineMeetings from app
type OnlineMeetingItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OnlineMeetingItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlineMeetingItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OnlineMeetingItemRequestBuilderGetQueryParameters
}
// OnlineMeetingItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlineMeetingItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AlternativeRecording provides operations to manage the media for the commsApplication entity.
func (m *OnlineMeetingItemRequestBuilder) AlternativeRecording()(*i621490855f8aac21cf8d7032977ac3e4dd93569a78b9ec995379c44e1ed33eaf.AlternativeRecordingRequestBuilder) {
    return i621490855f8aac21cf8d7032977ac3e4dd93569a78b9ec995379c44e1ed33eaf.NewAlternativeRecordingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttendanceReports provides operations to manage the attendanceReports property of the microsoft.graph.onlineMeeting entity.
func (m *OnlineMeetingItemRequestBuilder) AttendanceReports()(*i36d7f6a1b1906c3cdb16b4ede2175f668fa5388566026ae225a46e8d0be75ba4.AttendanceReportsRequestBuilder) {
    return i36d7f6a1b1906c3cdb16b4ede2175f668fa5388566026ae225a46e8d0be75ba4.NewAttendanceReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttendanceReportsById provides operations to manage the attendanceReports property of the microsoft.graph.onlineMeeting entity.
func (m *OnlineMeetingItemRequestBuilder) AttendanceReportsById(id string)(*i29fda42a3d65f7aed1e8a1e9c5006e031671b1aed329c71b8d3daa638210ed5a.MeetingAttendanceReportItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["meetingAttendanceReport%2Did"] = id
    }
    return i29fda42a3d65f7aed1e8a1e9c5006e031671b1aed329c71b8d3daa638210ed5a.NewMeetingAttendanceReportItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AttendeeReport provides operations to manage the media for the commsApplication entity.
func (m *OnlineMeetingItemRequestBuilder) AttendeeReport()(*i5792a57d1b509cf5853eddfc75cdd3d12d7fffac6c1d1ec273e69dfc55edec5d.AttendeeReportRequestBuilder) {
    return i5792a57d1b509cf5853eddfc75cdd3d12d7fffac6c1d1ec273e69dfc55edec5d.NewAttendeeReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOnlineMeetingItemRequestBuilderInternal instantiates a new OnlineMeetingItemRequestBuilder and sets the default values.
func NewOnlineMeetingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlineMeetingItemRequestBuilder) {
    m := &OnlineMeetingItemRequestBuilder{
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
// NewOnlineMeetingItemRequestBuilder instantiates a new OnlineMeetingItemRequestBuilder and sets the default values.
func NewOnlineMeetingItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlineMeetingItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnlineMeetingItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property onlineMeetings for app
func (m *OnlineMeetingItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *OnlineMeetingItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *OnlineMeetingItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *OnlineMeetingItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *OnlineMeetingItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable, requestConfiguration *OnlineMeetingItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *OnlineMeetingItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *OnlineMeetingItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *OnlineMeetingItemRequestBuilder) Get(ctx context.Context, requestConfiguration *OnlineMeetingItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable, error) {
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
func (m *OnlineMeetingItemRequestBuilder) MeetingAttendanceReport()(*ia44b3dfc614f798ddda4e107c7b00a962e62b26be06acb8ba2c30fda4607647a.MeetingAttendanceReportRequestBuilder) {
    return ia44b3dfc614f798ddda4e107c7b00a962e62b26be06acb8ba2c30fda4607647a.NewMeetingAttendanceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property onlineMeetings in app
func (m *OnlineMeetingItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable, requestConfiguration *OnlineMeetingItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable, error) {
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
func (m *OnlineMeetingItemRequestBuilder) Recording()(*ic4cea0d5545c1ec81a2c193eb0ee5d7a032aa549a87e4b5ea7d5e3d4250137fc.RecordingRequestBuilder) {
    return ic4cea0d5545c1ec81a2c193eb0ee5d7a032aa549a87e4b5ea7d5e3d4250137fc.NewRecordingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Registration provides operations to manage the registration property of the microsoft.graph.onlineMeeting entity.
func (m *OnlineMeetingItemRequestBuilder) Registration()(*ifc43ad7bd5674ca951e2271a8b4d99b63fa47b18f4fca9e422daedf0278acaf6.RegistrationRequestBuilder) {
    return ifc43ad7bd5674ca951e2271a8b4d99b63fa47b18f4fca9e422daedf0278acaf6.NewRegistrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Transcripts provides operations to manage the transcripts property of the microsoft.graph.onlineMeeting entity.
func (m *OnlineMeetingItemRequestBuilder) Transcripts()(*i67de322803fb9ceef33e79fb4b659d5e8d0ffaa90284f5ea1260109536f446ce.TranscriptsRequestBuilder) {
    return i67de322803fb9ceef33e79fb4b659d5e8d0ffaa90284f5ea1260109536f446ce.NewTranscriptsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranscriptsById provides operations to manage the transcripts property of the microsoft.graph.onlineMeeting entity.
func (m *OnlineMeetingItemRequestBuilder) TranscriptsById(id string)(*ide6e7b6bb6c57b727e00bfe216300dcd9aa6cd3b52ecddbd9256ee94b3816d97.CallTranscriptItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["callTranscript%2Did"] = id
    }
    return ide6e7b6bb6c57b727e00bfe216300dcd9aa6cd3b52ecddbd9256ee94b3816d97.NewCallTranscriptItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// VirtualAppointment provides operations to manage the virtualAppointment property of the microsoft.graph.onlineMeeting entity.
func (m *OnlineMeetingItemRequestBuilder) VirtualAppointment()(*if386b4892e4921e066f9d8ba80ef001577fdf97ddb75e8a48e18973728537f9e.VirtualAppointmentRequestBuilder) {
    return if386b4892e4921e066f9d8ba80ef001577fdf97ddb75e8a48e18973728537f9e.NewVirtualAppointmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
