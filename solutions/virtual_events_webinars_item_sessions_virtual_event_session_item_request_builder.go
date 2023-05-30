package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder provides operations to manage the sessions property of the microsoft.graph.virtualEvent entity.
type VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderGetQueryParameters get sessions from solutions
type VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderGetQueryParameters
}
// VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AlternativeRecording provides operations to manage the media for the solutionsRoot entity.
func (m *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) AlternativeRecording()(*VirtualEventsWebinarsItemSessionsItemAlternativeRecordingRequestBuilder) {
    return NewVirtualEventsWebinarsItemSessionsItemAlternativeRecordingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AttendanceReports provides operations to manage the attendanceReports property of the microsoft.graph.onlineMeeting entity.
func (m *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) AttendanceReports()(*VirtualEventsWebinarsItemSessionsItemAttendanceReportsRequestBuilder) {
    return NewVirtualEventsWebinarsItemSessionsItemAttendanceReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AttendeeReport provides operations to manage the media for the solutionsRoot entity.
func (m *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) AttendeeReport()(*VirtualEventsWebinarsItemSessionsItemAttendeeReportRequestBuilder) {
    return NewVirtualEventsWebinarsItemSessionsItemAttendeeReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BroadcastRecording provides operations to manage the media for the solutionsRoot entity.
func (m *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) BroadcastRecording()(*VirtualEventsWebinarsItemSessionsItemBroadcastRecordingRequestBuilder) {
    return NewVirtualEventsWebinarsItemSessionsItemBroadcastRecordingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderInternal instantiates a new VirtualEventSessionItemRequestBuilder and sets the default values.
func NewVirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) {
    m := &VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/webinars/{virtualEventWebinar%2Did}/sessions/{virtualEventSession%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewVirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder instantiates a new VirtualEventSessionItemRequestBuilder and sets the default values.
func NewVirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property sessions for solutions
func (m *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get sessions from solutions
func (m *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventSessionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateVirtualEventSessionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventSessionable), nil
}
// MeetingAttendanceReport provides operations to manage the meetingAttendanceReport property of the microsoft.graph.onlineMeeting entity.
func (m *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) MeetingAttendanceReport()(*VirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportRequestBuilder) {
    return NewVirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property sessions in solutions
func (m *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventSessionable, requestConfiguration *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventSessionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateVirtualEventSessionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventSessionable), nil
}
// Recording provides operations to manage the media for the solutionsRoot entity.
func (m *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) Recording()(*VirtualEventsWebinarsItemSessionsItemRecordingRequestBuilder) {
    return NewVirtualEventsWebinarsItemSessionsItemRecordingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Recordings provides operations to manage the recordings property of the microsoft.graph.onlineMeeting entity.
func (m *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) Recordings()(*VirtualEventsWebinarsItemSessionsItemRecordingsRequestBuilder) {
    return NewVirtualEventsWebinarsItemSessionsItemRecordingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Registration provides operations to manage the registration property of the microsoft.graph.onlineMeeting entity.
func (m *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) Registration()(*VirtualEventsWebinarsItemSessionsItemRegistrationRequestBuilder) {
    return NewVirtualEventsWebinarsItemSessionsItemRegistrationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property sessions for solutions
func (m *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation get sessions from solutions
func (m *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
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
// ToPatchRequestInformation update the navigation property sessions in solutions
func (m *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventSessionable, requestConfiguration *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Transcripts provides operations to manage the transcripts property of the microsoft.graph.onlineMeeting entity.
func (m *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) Transcripts()(*VirtualEventsWebinarsItemSessionsItemTranscriptsRequestBuilder) {
    return NewVirtualEventsWebinarsItemSessionsItemTranscriptsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// VirtualAppointment provides operations to manage the virtualAppointment property of the microsoft.graph.onlineMeeting entity.
func (m *VirtualEventsWebinarsItemSessionsVirtualEventSessionItemRequestBuilder) VirtualAppointment()(*VirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilder) {
    return NewVirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
