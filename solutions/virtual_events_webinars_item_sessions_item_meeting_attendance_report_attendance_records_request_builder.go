package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsRequestBuilder provides operations to manage the attendanceRecords property of the microsoft.graph.meetingAttendanceReport entity.
type VirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsRequestBuilderGetQueryParameters get a list of attendanceRecord objects and their properties.
type VirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// VirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsRequestBuilderGetQueryParameters
}
// VirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAttendanceRecordIdString provides operations to manage the attendanceRecords property of the microsoft.graph.meetingAttendanceReport entity.
func (m *VirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsRequestBuilder) ByAttendanceRecordIdString(attendanceRecordId string)(*VirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if attendanceRecordId != "" {
        urlTplParams["attendanceRecord%2Did"] = attendanceRecordId
    }
    return NewVirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsRequestBuilderInternal instantiates a new AttendanceRecordsRequestBuilder and sets the default values.
func NewVirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsRequestBuilder) {
    m := &VirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/webinars/{virtualEventWebinar%2Did}/sessions/{virtualEventSession%2Did}/meetingAttendanceReport/attendanceRecords{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewVirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsRequestBuilder instantiates a new AttendanceRecordsRequestBuilder and sets the default values.
func NewVirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *VirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsRequestBuilder) Count()(*VirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsCountRequestBuilder) {
    return NewVirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of attendanceRecord objects and their properties.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/attendancerecord-list?view=graph-rest-1.0
func (m *VirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendanceRecordCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAttendanceRecordCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendanceRecordCollectionResponseable), nil
}
// Post create new navigation property to attendanceRecords for solutions
func (m *VirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendanceRecordable, requestConfiguration *VirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendanceRecordable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAttendanceRecordFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendanceRecordable), nil
}
// ToGetRequestInformation get a list of attendanceRecord objects and their properties.
func (m *VirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to attendanceRecords for solutions
func (m *VirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendanceRecordable, requestConfiguration *VirtualEventsWebinarsItemSessionsItemMeetingAttendanceReportAttendanceRecordsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
