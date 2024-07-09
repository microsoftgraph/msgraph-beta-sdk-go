package app

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OnlineMeetingsOnlineMeetingItemRequestBuilder provides operations to manage the onlineMeetings property of the microsoft.graph.commsApplication entity.
type OnlineMeetingsOnlineMeetingItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OnlineMeetingsOnlineMeetingItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlineMeetingsOnlineMeetingItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OnlineMeetingsOnlineMeetingItemRequestBuilderGetQueryParameters get onlineMeetings from app
type OnlineMeetingsOnlineMeetingItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OnlineMeetingsOnlineMeetingItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlineMeetingsOnlineMeetingItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OnlineMeetingsOnlineMeetingItemRequestBuilderGetQueryParameters
}
// OnlineMeetingsOnlineMeetingItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlineMeetingsOnlineMeetingItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AlternativeRecording provides operations to manage the media for the commsApplication entity.
// returns a *OnlineMeetingsItemAlternativeRecordingRequestBuilder when successful
func (m *OnlineMeetingsOnlineMeetingItemRequestBuilder) AlternativeRecording()(*OnlineMeetingsItemAlternativeRecordingRequestBuilder) {
    return NewOnlineMeetingsItemAlternativeRecordingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AttendanceReports provides operations to manage the attendanceReports property of the microsoft.graph.onlineMeetingBase entity.
// returns a *OnlineMeetingsItemAttendanceReportsRequestBuilder when successful
func (m *OnlineMeetingsOnlineMeetingItemRequestBuilder) AttendanceReports()(*OnlineMeetingsItemAttendanceReportsRequestBuilder) {
    return NewOnlineMeetingsItemAttendanceReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AttendeeReport provides operations to manage the media for the commsApplication entity.
// returns a *OnlineMeetingsItemAttendeeReportRequestBuilder when successful
func (m *OnlineMeetingsOnlineMeetingItemRequestBuilder) AttendeeReport()(*OnlineMeetingsItemAttendeeReportRequestBuilder) {
    return NewOnlineMeetingsItemAttendeeReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BroadcastRecording provides operations to manage the media for the commsApplication entity.
// returns a *OnlineMeetingsItemBroadcastRecordingRequestBuilder when successful
func (m *OnlineMeetingsOnlineMeetingItemRequestBuilder) BroadcastRecording()(*OnlineMeetingsItemBroadcastRecordingRequestBuilder) {
    return NewOnlineMeetingsItemBroadcastRecordingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewOnlineMeetingsOnlineMeetingItemRequestBuilderInternal instantiates a new OnlineMeetingsOnlineMeetingItemRequestBuilder and sets the default values.
func NewOnlineMeetingsOnlineMeetingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlineMeetingsOnlineMeetingItemRequestBuilder) {
    m := &OnlineMeetingsOnlineMeetingItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/app/onlineMeetings/{onlineMeeting%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewOnlineMeetingsOnlineMeetingItemRequestBuilder instantiates a new OnlineMeetingsOnlineMeetingItemRequestBuilder and sets the default values.
func NewOnlineMeetingsOnlineMeetingItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlineMeetingsOnlineMeetingItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnlineMeetingsOnlineMeetingItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property onlineMeetings for app
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlineMeetingsOnlineMeetingItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *OnlineMeetingsOnlineMeetingItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get onlineMeetings from app
// returns a OnlineMeetingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlineMeetingsOnlineMeetingItemRequestBuilder) Get(ctx context.Context, requestConfiguration *OnlineMeetingsOnlineMeetingItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnlineMeetingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable), nil
}
// GetVirtualAppointmentJoinWebUrl provides operations to call the getVirtualAppointmentJoinWebUrl method.
// returns a *OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilder when successful
func (m *OnlineMeetingsOnlineMeetingItemRequestBuilder) GetVirtualAppointmentJoinWebUrl()(*OnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilder) {
    return NewOnlineMeetingsItemGetVirtualAppointmentJoinWebUrlRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MeetingAttendanceReport provides operations to manage the meetingAttendanceReport property of the microsoft.graph.onlineMeeting entity.
// returns a *OnlineMeetingsItemMeetingAttendanceReportRequestBuilder when successful
func (m *OnlineMeetingsOnlineMeetingItemRequestBuilder) MeetingAttendanceReport()(*OnlineMeetingsItemMeetingAttendanceReportRequestBuilder) {
    return NewOnlineMeetingsItemMeetingAttendanceReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property onlineMeetings in app
// returns a OnlineMeetingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlineMeetingsOnlineMeetingItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable, requestConfiguration *OnlineMeetingsOnlineMeetingItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnlineMeetingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable), nil
}
// Recording provides operations to manage the media for the commsApplication entity.
// returns a *OnlineMeetingsItemRecordingRequestBuilder when successful
func (m *OnlineMeetingsOnlineMeetingItemRequestBuilder) Recording()(*OnlineMeetingsItemRecordingRequestBuilder) {
    return NewOnlineMeetingsItemRecordingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Recordings provides operations to manage the recordings property of the microsoft.graph.onlineMeeting entity.
// returns a *OnlineMeetingsItemRecordingsRequestBuilder when successful
func (m *OnlineMeetingsOnlineMeetingItemRequestBuilder) Recordings()(*OnlineMeetingsItemRecordingsRequestBuilder) {
    return NewOnlineMeetingsItemRecordingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Registration provides operations to manage the registration property of the microsoft.graph.onlineMeeting entity.
// returns a *OnlineMeetingsItemRegistrationRequestBuilder when successful
func (m *OnlineMeetingsOnlineMeetingItemRequestBuilder) Registration()(*OnlineMeetingsItemRegistrationRequestBuilder) {
    return NewOnlineMeetingsItemRegistrationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SendVirtualAppointmentReminderSms provides operations to call the sendVirtualAppointmentReminderSms method.
// returns a *OnlineMeetingsItemSendVirtualAppointmentReminderSmsRequestBuilder when successful
func (m *OnlineMeetingsOnlineMeetingItemRequestBuilder) SendVirtualAppointmentReminderSms()(*OnlineMeetingsItemSendVirtualAppointmentReminderSmsRequestBuilder) {
    return NewOnlineMeetingsItemSendVirtualAppointmentReminderSmsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SendVirtualAppointmentSms provides operations to call the sendVirtualAppointmentSms method.
// returns a *OnlineMeetingsItemSendVirtualAppointmentSmsRequestBuilder when successful
func (m *OnlineMeetingsOnlineMeetingItemRequestBuilder) SendVirtualAppointmentSms()(*OnlineMeetingsItemSendVirtualAppointmentSmsRequestBuilder) {
    return NewOnlineMeetingsItemSendVirtualAppointmentSmsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property onlineMeetings for app
// returns a *RequestInformation when successful
func (m *OnlineMeetingsOnlineMeetingItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *OnlineMeetingsOnlineMeetingItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get onlineMeetings from app
// returns a *RequestInformation when successful
func (m *OnlineMeetingsOnlineMeetingItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OnlineMeetingsOnlineMeetingItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property onlineMeetings in app
// returns a *RequestInformation when successful
func (m *OnlineMeetingsOnlineMeetingItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable, requestConfiguration *OnlineMeetingsOnlineMeetingItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Transcripts provides operations to manage the transcripts property of the microsoft.graph.onlineMeeting entity.
// returns a *OnlineMeetingsItemTranscriptsRequestBuilder when successful
func (m *OnlineMeetingsOnlineMeetingItemRequestBuilder) Transcripts()(*OnlineMeetingsItemTranscriptsRequestBuilder) {
    return NewOnlineMeetingsItemTranscriptsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *OnlineMeetingsOnlineMeetingItemRequestBuilder when successful
func (m *OnlineMeetingsOnlineMeetingItemRequestBuilder) WithUrl(rawUrl string)(*OnlineMeetingsOnlineMeetingItemRequestBuilder) {
    return NewOnlineMeetingsOnlineMeetingItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
