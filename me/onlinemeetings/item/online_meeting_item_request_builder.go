package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2a0e0d48f1a133f80557e834fceb9001d321597582147fa92c8d557b1e60c722 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onlinemeetings/item/recording"
    i47c2284118cc15348aa203315c0a4637fa39b84c65e76ecdc39fc175c32c90b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onlinemeetings/item/meetingattendancereport"
    i520e7c5dd1b2e99f3cf76ee8e21c2e881be13a393c94622086fa442c45f27e8f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onlinemeetings/item/attendancereports"
    i60f747f04da18ee740c2d7cec8aa333e174de51a71c041da85eebd852e1da259 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onlinemeetings/item/registration"
    i637e5f11ab6a0348359da991cd3af4494645f065cc80ca54794bd4a94e21eab5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onlinemeetings/item/transcripts"
    i9751454805a6089ded23da943214fdbe9028e1de749d44a87219eee6d363460e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onlinemeetings/item/alternativerecording"
    i98d7ddf881a02ae894ac70424b059f80f8d245996309997a347c38739b78710b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onlinemeetings/item/attendeereport"
    ibb397ff063cfeb8d66586f5575db66957e02eb2a920735e9720ca3df8f70f83d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onlinemeetings/item/virtualappointment"
    ib6386ca31a5eed83c6e42da470d03db28a79e495504bfb7719dca5c299d78f4f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onlinemeetings/item/transcripts/item"
    id7c487e4aca56b8a93bfe6b7467589181f3783dae6bf5fc07a9339b86fdc342e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onlinemeetings/item/attendancereports/item"
)

// OnlineMeetingItemRequestBuilder provides operations to manage the onlineMeetings property of the microsoft.graph.user entity.
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
// OnlineMeetingItemRequestBuilderGetQueryParameters get onlineMeetings from me
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
// AlternativeRecording provides operations to manage the media for the user entity.
func (m *OnlineMeetingItemRequestBuilder) AlternativeRecording()(*i9751454805a6089ded23da943214fdbe9028e1de749d44a87219eee6d363460e.AlternativeRecordingRequestBuilder) {
    return i9751454805a6089ded23da943214fdbe9028e1de749d44a87219eee6d363460e.NewAlternativeRecordingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttendanceReports provides operations to manage the attendanceReports property of the microsoft.graph.onlineMeeting entity.
func (m *OnlineMeetingItemRequestBuilder) AttendanceReports()(*i520e7c5dd1b2e99f3cf76ee8e21c2e881be13a393c94622086fa442c45f27e8f.AttendanceReportsRequestBuilder) {
    return i520e7c5dd1b2e99f3cf76ee8e21c2e881be13a393c94622086fa442c45f27e8f.NewAttendanceReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttendanceReportsById provides operations to manage the attendanceReports property of the microsoft.graph.onlineMeeting entity.
func (m *OnlineMeetingItemRequestBuilder) AttendanceReportsById(id string)(*id7c487e4aca56b8a93bfe6b7467589181f3783dae6bf5fc07a9339b86fdc342e.MeetingAttendanceReportItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["meetingAttendanceReport%2Did"] = id
    }
    return id7c487e4aca56b8a93bfe6b7467589181f3783dae6bf5fc07a9339b86fdc342e.NewMeetingAttendanceReportItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AttendeeReport provides operations to manage the media for the user entity.
func (m *OnlineMeetingItemRequestBuilder) AttendeeReport()(*i98d7ddf881a02ae894ac70424b059f80f8d245996309997a347c38739b78710b.AttendeeReportRequestBuilder) {
    return i98d7ddf881a02ae894ac70424b059f80f8d245996309997a347c38739b78710b.NewAttendeeReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOnlineMeetingItemRequestBuilderInternal instantiates a new OnlineMeetingItemRequestBuilder and sets the default values.
func NewOnlineMeetingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlineMeetingItemRequestBuilder) {
    m := &OnlineMeetingItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/onlineMeetings/{onlineMeeting%2Did}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property onlineMeetings for me
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
// CreateGetRequestInformation get onlineMeetings from me
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
// CreatePatchRequestInformation update the navigation property onlineMeetings in me
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
// Delete delete navigation property onlineMeetings for me
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
// Get get onlineMeetings from me
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
func (m *OnlineMeetingItemRequestBuilder) MeetingAttendanceReport()(*i47c2284118cc15348aa203315c0a4637fa39b84c65e76ecdc39fc175c32c90b9.MeetingAttendanceReportRequestBuilder) {
    return i47c2284118cc15348aa203315c0a4637fa39b84c65e76ecdc39fc175c32c90b9.NewMeetingAttendanceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property onlineMeetings in me
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
// Recording provides operations to manage the media for the user entity.
func (m *OnlineMeetingItemRequestBuilder) Recording()(*i2a0e0d48f1a133f80557e834fceb9001d321597582147fa92c8d557b1e60c722.RecordingRequestBuilder) {
    return i2a0e0d48f1a133f80557e834fceb9001d321597582147fa92c8d557b1e60c722.NewRecordingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Registration provides operations to manage the registration property of the microsoft.graph.onlineMeeting entity.
func (m *OnlineMeetingItemRequestBuilder) Registration()(*i60f747f04da18ee740c2d7cec8aa333e174de51a71c041da85eebd852e1da259.RegistrationRequestBuilder) {
    return i60f747f04da18ee740c2d7cec8aa333e174de51a71c041da85eebd852e1da259.NewRegistrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Transcripts provides operations to manage the transcripts property of the microsoft.graph.onlineMeeting entity.
func (m *OnlineMeetingItemRequestBuilder) Transcripts()(*i637e5f11ab6a0348359da991cd3af4494645f065cc80ca54794bd4a94e21eab5.TranscriptsRequestBuilder) {
    return i637e5f11ab6a0348359da991cd3af4494645f065cc80ca54794bd4a94e21eab5.NewTranscriptsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranscriptsById provides operations to manage the transcripts property of the microsoft.graph.onlineMeeting entity.
func (m *OnlineMeetingItemRequestBuilder) TranscriptsById(id string)(*ib6386ca31a5eed83c6e42da470d03db28a79e495504bfb7719dca5c299d78f4f.CallTranscriptItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["callTranscript%2Did"] = id
    }
    return ib6386ca31a5eed83c6e42da470d03db28a79e495504bfb7719dca5c299d78f4f.NewCallTranscriptItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// VirtualAppointment provides operations to manage the virtualAppointment property of the microsoft.graph.onlineMeeting entity.
func (m *OnlineMeetingItemRequestBuilder) VirtualAppointment()(*ibb397ff063cfeb8d66586f5575db66957e02eb2a920735e9720ca3df8f70f83d.VirtualAppointmentRequestBuilder) {
    return ibb397ff063cfeb8d66586f5575db66957e02eb2a920735e9720ca3df8f70f83d.NewVirtualAppointmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
