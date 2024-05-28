package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder provides operations to call the getVirtualAppointmentJoinWebUrl method.
type OnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilderInternal instantiates a new OnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder and sets the default values.
func NewOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder) {
    m := &OnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/onlineMeetings/{onlineMeeting%2Did}/getVirtualAppointmentJoinWebUrl()", pathParameters),
    }
    return m
}
// NewOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder instantiates a new OnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder and sets the default values.
func NewOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a join web URL for a Teams Virtual Appointment. This web URL includes enhanced business-to-customer experiences such as mobile browser join and virtual lobby rooms. With Teams Premium, you can configure a custom lobby room experience for attendees by adding your company logo and access the Virtual Appointments usage report for organizational analytics.
// Deprecated: This method is obsolete. Use GetAsGetVirtualAppointmentJoinWebUrlGetResponse instead.
// returns a OnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualappointment-getvirtualappointmentjoinweburl?view=graph-rest-beta
func (m *OnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder) Get(ctx context.Context, requestConfiguration *OnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilderGetRequestConfiguration)(OnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlResponseable), nil
}
// GetAsGetVirtualAppointmentJoinWebUrlGetResponse get a join web URL for a Teams Virtual Appointment. This web URL includes enhanced business-to-customer experiences such as mobile browser join and virtual lobby rooms. With Teams Premium, you can configure a custom lobby room experience for attendees by adding your company logo and access the Virtual Appointments usage report for organizational analytics.
// returns a OnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualappointment-getvirtualappointmentjoinweburl?view=graph-rest-beta
func (m *OnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder) GetAsGetVirtualAppointmentJoinWebUrlGetResponse(ctx context.Context, requestConfiguration *OnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilderGetRequestConfiguration)(OnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlGetResponseable), nil
}
// ToGetRequestInformation get a join web URL for a Teams Virtual Appointment. This web URL includes enhanced business-to-customer experiences such as mobile browser join and virtual lobby rooms. With Teams Premium, you can configure a custom lobby room experience for attendees by adding your company logo and access the Virtual Appointments usage report for organizational analytics.
// returns a *RequestInformation when successful
func (m *OnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *OnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder when successful
func (m *OnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder) WithUrl(rawUrl string)(*OnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder) {
    return NewOnlinemeetingsItemGetvirtualappointmentjoinweburlGetVirtualAppointmentJoinWebUrlRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
