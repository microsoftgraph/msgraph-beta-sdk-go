package app

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OnlinemeetingsItemMeetingattendancereportAttendancerecordsCountRequestBuilder provides operations to count the resources in the collection.
type OnlinemeetingsItemMeetingattendancereportAttendancerecordsCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OnlinemeetingsItemMeetingattendancereportAttendancerecordsCountRequestBuilderGetQueryParameters get the number of the resource
type OnlinemeetingsItemMeetingattendancereportAttendancerecordsCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// OnlinemeetingsItemMeetingattendancereportAttendancerecordsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlinemeetingsItemMeetingattendancereportAttendancerecordsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OnlinemeetingsItemMeetingattendancereportAttendancerecordsCountRequestBuilderGetQueryParameters
}
// NewOnlinemeetingsItemMeetingattendancereportAttendancerecordsCountRequestBuilderInternal instantiates a new OnlinemeetingsItemMeetingattendancereportAttendancerecordsCountRequestBuilder and sets the default values.
func NewOnlinemeetingsItemMeetingattendancereportAttendancerecordsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlinemeetingsItemMeetingattendancereportAttendancerecordsCountRequestBuilder) {
    m := &OnlinemeetingsItemMeetingattendancereportAttendancerecordsCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/app/onlineMeetings/{onlineMeeting%2Did}/meetingAttendanceReport/attendanceRecords/$count{?%24filter,%24search}", pathParameters),
    }
    return m
}
// NewOnlinemeetingsItemMeetingattendancereportAttendancerecordsCountRequestBuilder instantiates a new OnlinemeetingsItemMeetingattendancereportAttendancerecordsCountRequestBuilder and sets the default values.
func NewOnlinemeetingsItemMeetingattendancereportAttendancerecordsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlinemeetingsItemMeetingattendancereportAttendancerecordsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnlinemeetingsItemMeetingattendancereportAttendancerecordsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
// returns a *int32 when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlinemeetingsItemMeetingattendancereportAttendancerecordsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *OnlinemeetingsItemMeetingattendancereportAttendancerecordsCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
// returns a *RequestInformation when successful
func (m *OnlinemeetingsItemMeetingattendancereportAttendancerecordsCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OnlinemeetingsItemMeetingattendancereportAttendancerecordsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *OnlinemeetingsItemMeetingattendancereportAttendancerecordsCountRequestBuilder when successful
func (m *OnlinemeetingsItemMeetingattendancereportAttendancerecordsCountRequestBuilder) WithUrl(rawUrl string)(*OnlinemeetingsItemMeetingattendancereportAttendancerecordsCountRequestBuilder) {
    return NewOnlinemeetingsItemMeetingattendancereportAttendancerecordsCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
