package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i28844d8a5d325bb0845a1778d6df0388d975703f04633f34ffc2f4073b871389 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/onlinemeetings/item/registration"
    i3a2bfcc71796b9365b1b580ed5a11589b09828c31bbc6758b6afbd6ee2ee53a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/onlinemeetings/item/attendeereport"
    i973f10e60f8f0226efcd773b4162b5082a7c2927993ac09fbb3f58bd0fb4fa4d "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/onlinemeetings/item/attendancereports"
    ia176566273a5afb80848389bc0868b65818d8d6d3ab8bc727ddcd5d4015393f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/onlinemeetings/item/recording"
    ibc5e0e3096be23e911a3fdca31d7d21b8c3cc6a11a0e00e3e861fdc763ebedfd "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/onlinemeetings/item/virtualappointment"
    ibf215dcedc685523e33219b91fd02247c51376f20991fc707be93a6e0640d955 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/onlinemeetings/item/transcripts"
    ie765355a32ec6cb4ba90da216a62ec4fb6b7b232eb3780cefe410dc49551bbe5 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/onlinemeetings/item/meetingattendancereport"
    ifa2556b906140d07529e0b82dd914940679909ab68a2bc38a29559eceb0566e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/onlinemeetings/item/alternativerecording"
    i1df49adb94e4624ec86cfc9e723c409892b4e3c93067331f6ac6280b4cd1c473 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/onlinemeetings/item/transcripts/item"
    ib99f82b2686be9ab3c0709927e690dd3df4c41d545906785c734827457ee1d56 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/onlinemeetings/item/attendancereports/item"
)

// OnlineMeetingItemRequestBuilder provides operations to manage the onlineMeetings property of the microsoft.graph.cloudCommunications entity.
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
// OnlineMeetingItemRequestBuilderGetQueryParameters get onlineMeetings from communications
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
// AlternativeRecording the alternativeRecording property
func (m *OnlineMeetingItemRequestBuilder) AlternativeRecording()(*ifa2556b906140d07529e0b82dd914940679909ab68a2bc38a29559eceb0566e9.AlternativeRecordingRequestBuilder) {
    return ifa2556b906140d07529e0b82dd914940679909ab68a2bc38a29559eceb0566e9.NewAlternativeRecordingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttendanceReports the attendanceReports property
func (m *OnlineMeetingItemRequestBuilder) AttendanceReports()(*i973f10e60f8f0226efcd773b4162b5082a7c2927993ac09fbb3f58bd0fb4fa4d.AttendanceReportsRequestBuilder) {
    return i973f10e60f8f0226efcd773b4162b5082a7c2927993ac09fbb3f58bd0fb4fa4d.NewAttendanceReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttendanceReportsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.communications.onlineMeetings.item.attendanceReports.item collection
func (m *OnlineMeetingItemRequestBuilder) AttendanceReportsById(id string)(*ib99f82b2686be9ab3c0709927e690dd3df4c41d545906785c734827457ee1d56.MeetingAttendanceReportItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["meetingAttendanceReport%2Did"] = id
    }
    return ib99f82b2686be9ab3c0709927e690dd3df4c41d545906785c734827457ee1d56.NewMeetingAttendanceReportItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AttendeeReport the attendeeReport property
func (m *OnlineMeetingItemRequestBuilder) AttendeeReport()(*i3a2bfcc71796b9365b1b580ed5a11589b09828c31bbc6758b6afbd6ee2ee53a9.AttendeeReportRequestBuilder) {
    return i3a2bfcc71796b9365b1b580ed5a11589b09828c31bbc6758b6afbd6ee2ee53a9.NewAttendeeReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOnlineMeetingItemRequestBuilderInternal instantiates a new OnlineMeetingItemRequestBuilder and sets the default values.
func NewOnlineMeetingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlineMeetingItemRequestBuilder) {
    m := &OnlineMeetingItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/onlineMeetings/{onlineMeeting%2Did}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property onlineMeetings for communications
func (m *OnlineMeetingItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property onlineMeetings for communications
func (m *OnlineMeetingItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *OnlineMeetingItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get onlineMeetings from communications
func (m *OnlineMeetingItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get onlineMeetings from communications
func (m *OnlineMeetingItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *OnlineMeetingItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property onlineMeetings in communications
func (m *OnlineMeetingItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property onlineMeetings in communications
func (m *OnlineMeetingItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable, requestConfiguration *OnlineMeetingItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property onlineMeetings for communications
func (m *OnlineMeetingItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property onlineMeetings for communications
func (m *OnlineMeetingItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *OnlineMeetingItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get onlineMeetings from communications
func (m *OnlineMeetingItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get onlineMeetings from communications
func (m *OnlineMeetingItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *OnlineMeetingItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnlineMeetingFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable), nil
}
// MeetingAttendanceReport the meetingAttendanceReport property
func (m *OnlineMeetingItemRequestBuilder) MeetingAttendanceReport()(*ie765355a32ec6cb4ba90da216a62ec4fb6b7b232eb3780cefe410dc49551bbe5.MeetingAttendanceReportRequestBuilder) {
    return ie765355a32ec6cb4ba90da216a62ec4fb6b7b232eb3780cefe410dc49551bbe5.NewMeetingAttendanceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property onlineMeetings in communications
func (m *OnlineMeetingItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property onlineMeetings in communications
func (m *OnlineMeetingItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable, requestConfiguration *OnlineMeetingItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Recording the recording property
func (m *OnlineMeetingItemRequestBuilder) Recording()(*ia176566273a5afb80848389bc0868b65818d8d6d3ab8bc727ddcd5d4015393f9.RecordingRequestBuilder) {
    return ia176566273a5afb80848389bc0868b65818d8d6d3ab8bc727ddcd5d4015393f9.NewRecordingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Registration the registration property
func (m *OnlineMeetingItemRequestBuilder) Registration()(*i28844d8a5d325bb0845a1778d6df0388d975703f04633f34ffc2f4073b871389.RegistrationRequestBuilder) {
    return i28844d8a5d325bb0845a1778d6df0388d975703f04633f34ffc2f4073b871389.NewRegistrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Transcripts the transcripts property
func (m *OnlineMeetingItemRequestBuilder) Transcripts()(*ibf215dcedc685523e33219b91fd02247c51376f20991fc707be93a6e0640d955.TranscriptsRequestBuilder) {
    return ibf215dcedc685523e33219b91fd02247c51376f20991fc707be93a6e0640d955.NewTranscriptsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranscriptsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.communications.onlineMeetings.item.transcripts.item collection
func (m *OnlineMeetingItemRequestBuilder) TranscriptsById(id string)(*i1df49adb94e4624ec86cfc9e723c409892b4e3c93067331f6ac6280b4cd1c473.CallTranscriptItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["callTranscript%2Did"] = id
    }
    return i1df49adb94e4624ec86cfc9e723c409892b4e3c93067331f6ac6280b4cd1c473.NewCallTranscriptItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// VirtualAppointment the virtualAppointment property
func (m *OnlineMeetingItemRequestBuilder) VirtualAppointment()(*ibc5e0e3096be23e911a3fdca31d7d21b8c3cc6a11a0e00e3e861fdc763ebedfd.VirtualAppointmentRequestBuilder) {
    return ibc5e0e3096be23e911a3fdca31d7d21b8c3cc6a11a0e00e3e861fdc763ebedfd.NewVirtualAppointmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
