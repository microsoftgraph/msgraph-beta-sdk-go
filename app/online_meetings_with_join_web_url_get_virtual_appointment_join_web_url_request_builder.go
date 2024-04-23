package app

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder provides operations to call the getVirtualAppointmentJoinWebUrl method.
type OnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilderInternal instantiates a new OnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder and sets the default values.
func NewOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder) {
    m := &OnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/app/onlineMeetings(joinWebUrl='{joinWebUrl}')/getVirtualAppointmentJoinWebUrl()", pathParameters),
    }
    return m
}
// NewOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder instantiates a new OnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder and sets the default values.
func NewOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a join web URL for a Teams Virtual Appointment. This web URL includes enhanced business-to-customer experiences such as mobile browser join and virtual lobby rooms. With Teams Premium, you can configure a custom lobby room experience for attendees by adding your company logo and access the Virtual Appointments usage report for organizational analytics.
// Deprecated: This method is obsolete. Use GetAsGetVirtualAppointmentJoinWebUrlGetResponse instead.
// returns a OnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualappointment-getvirtualappointmentjoinweburl?view=graph-rest-beta
func (m *OnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder) Get(ctx context.Context, requestConfiguration *OnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilderGetRequestConfiguration)(OnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlResponseable), nil
}
// GetAsGetVirtualAppointmentJoinWebUrlGetResponse get a join web URL for a Teams Virtual Appointment. This web URL includes enhanced business-to-customer experiences such as mobile browser join and virtual lobby rooms. With Teams Premium, you can configure a custom lobby room experience for attendees by adding your company logo and access the Virtual Appointments usage report for organizational analytics.
// returns a OnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualappointment-getvirtualappointmentjoinweburl?view=graph-rest-beta
func (m *OnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder) GetAsGetVirtualAppointmentJoinWebUrlGetResponse(ctx context.Context, requestConfiguration *OnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilderGetRequestConfiguration)(OnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlGetResponseable), nil
}
// ToGetRequestInformation get a join web URL for a Teams Virtual Appointment. This web URL includes enhanced business-to-customer experiences such as mobile browser join and virtual lobby rooms. With Teams Premium, you can configure a custom lobby room experience for attendees by adding your company logo and access the Virtual Appointments usage report for organizational analytics.
// returns a *RequestInformation when successful
func (m *OnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *OnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder when successful
func (m *OnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder) WithUrl(rawUrl string)(*OnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder) {
    return NewOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
