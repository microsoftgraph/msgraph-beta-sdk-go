package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEventsWebinarsItemRegistrationsItemSessionsVirtualEventSessionItemRequestBuilder provides operations to manage the sessions property of the microsoft.graph.virtualEventRegistration entity.
type VirtualEventsWebinarsItemRegistrationsItemSessionsVirtualEventSessionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEventsWebinarsItemRegistrationsItemSessionsVirtualEventSessionItemRequestBuilderGetQueryParameters sessions of the webinar.
type VirtualEventsWebinarsItemRegistrationsItemSessionsVirtualEventSessionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualEventsWebinarsItemRegistrationsItemSessionsVirtualEventSessionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsWebinarsItemRegistrationsItemSessionsVirtualEventSessionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEventsWebinarsItemRegistrationsItemSessionsVirtualEventSessionItemRequestBuilderGetQueryParameters
}
// AlternativeRecording provides operations to manage the media for the solutionsRoot entity.
func (m *VirtualEventsWebinarsItemRegistrationsItemSessionsVirtualEventSessionItemRequestBuilder) AlternativeRecording()(*VirtualEventsWebinarsItemRegistrationsItemSessionsItemAlternativeRecordingRequestBuilder) {
    return NewVirtualEventsWebinarsItemRegistrationsItemSessionsItemAlternativeRecordingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AttendeeReport provides operations to manage the media for the solutionsRoot entity.
func (m *VirtualEventsWebinarsItemRegistrationsItemSessionsVirtualEventSessionItemRequestBuilder) AttendeeReport()(*VirtualEventsWebinarsItemRegistrationsItemSessionsItemAttendeeReportRequestBuilder) {
    return NewVirtualEventsWebinarsItemRegistrationsItemSessionsItemAttendeeReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BroadcastRecording provides operations to manage the media for the solutionsRoot entity.
func (m *VirtualEventsWebinarsItemRegistrationsItemSessionsVirtualEventSessionItemRequestBuilder) BroadcastRecording()(*VirtualEventsWebinarsItemRegistrationsItemSessionsItemBroadcastRecordingRequestBuilder) {
    return NewVirtualEventsWebinarsItemRegistrationsItemSessionsItemBroadcastRecordingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualEventsWebinarsItemRegistrationsItemSessionsVirtualEventSessionItemRequestBuilderInternal instantiates a new VirtualEventSessionItemRequestBuilder and sets the default values.
func NewVirtualEventsWebinarsItemRegistrationsItemSessionsVirtualEventSessionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsWebinarsItemRegistrationsItemSessionsVirtualEventSessionItemRequestBuilder) {
    m := &VirtualEventsWebinarsItemRegistrationsItemSessionsVirtualEventSessionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/webinars/{virtualEventWebinar%2Did}/registrations/{virtualEventRegistration%2Did}/sessions/{virtualEventSession%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewVirtualEventsWebinarsItemRegistrationsItemSessionsVirtualEventSessionItemRequestBuilder instantiates a new VirtualEventSessionItemRequestBuilder and sets the default values.
func NewVirtualEventsWebinarsItemRegistrationsItemSessionsVirtualEventSessionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsWebinarsItemRegistrationsItemSessionsVirtualEventSessionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEventsWebinarsItemRegistrationsItemSessionsVirtualEventSessionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get sessions of the webinar.
func (m *VirtualEventsWebinarsItemRegistrationsItemSessionsVirtualEventSessionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEventsWebinarsItemRegistrationsItemSessionsVirtualEventSessionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventSessionable, error) {
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
// Recording provides operations to manage the media for the solutionsRoot entity.
func (m *VirtualEventsWebinarsItemRegistrationsItemSessionsVirtualEventSessionItemRequestBuilder) Recording()(*VirtualEventsWebinarsItemRegistrationsItemSessionsItemRecordingRequestBuilder) {
    return NewVirtualEventsWebinarsItemRegistrationsItemSessionsItemRecordingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation sessions of the webinar.
func (m *VirtualEventsWebinarsItemRegistrationsItemSessionsVirtualEventSessionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEventsWebinarsItemRegistrationsItemSessionsVirtualEventSessionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *VirtualEventsWebinarsItemRegistrationsItemSessionsVirtualEventSessionItemRequestBuilder) WithUrl(rawUrl string)(*VirtualEventsWebinarsItemRegistrationsItemSessionsVirtualEventSessionItemRequestBuilder) {
    return NewVirtualEventsWebinarsItemRegistrationsItemSessionsVirtualEventSessionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
