package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilder provides operations to manage the attendanceReports property of the microsoft.graph.onlineMeetingBase entity.
type VirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilderGetQueryParameters the attendance reports of an online meeting. Read-only.
type VirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilderGetQueryParameters struct {
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
// VirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilderGetQueryParameters
}
// VirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByMeetingAttendanceReportId provides operations to manage the attendanceReports property of the microsoft.graph.onlineMeetingBase entity.
// returns a *VirtualEventsTownhallsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilder when successful
func (m *VirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilder) ByMeetingAttendanceReportId(meetingAttendanceReportId string)(*VirtualEventsTownhallsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if meetingAttendanceReportId != "" {
        urlTplParams["meetingAttendanceReport%2Did"] = meetingAttendanceReportId
    }
    return NewVirtualEventsTownhallsItemSessionsItemAttendanceReportsMeetingAttendanceReportItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilderInternal instantiates a new VirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilder and sets the default values.
func NewVirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilder) {
    m := &VirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/townhalls/{virtualEventTownhall%2Did}/sessions/{virtualEventSession%2Did}/attendanceReports{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewVirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilder instantiates a new VirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilder and sets the default values.
func NewVirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *VirtualEventsTownhallsItemSessionsItemAttendanceReportsCountRequestBuilder when successful
func (m *VirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilder) Count()(*VirtualEventsTownhallsItemSessionsItemAttendanceReportsCountRequestBuilder) {
    return NewVirtualEventsTownhallsItemSessionsItemAttendanceReportsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the attendance reports of an online meeting. Read-only.
// returns a MeetingAttendanceReportCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingAttendanceReportCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMeetingAttendanceReportCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingAttendanceReportCollectionResponseable), nil
}
// Post create new navigation property to attendanceReports for solutions
// returns a MeetingAttendanceReportable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingAttendanceReportable, requestConfiguration *VirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingAttendanceReportable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMeetingAttendanceReportFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingAttendanceReportable), nil
}
// ToGetRequestInformation the attendance reports of an online meeting. Read-only.
// returns a *RequestInformation when successful
func (m *VirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to attendanceReports for solutions
// returns a *RequestInformation when successful
func (m *VirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingAttendanceReportable, requestConfiguration *VirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/solutions/virtualEvents/townhalls/{virtualEventTownhall%2Did}/sessions/{virtualEventSession%2Did}/attendanceReports", m.BaseRequestBuilder.PathParameters)
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilder when successful
func (m *VirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilder) WithUrl(rawUrl string)(*VirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilder) {
    return NewVirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
