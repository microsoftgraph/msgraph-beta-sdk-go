package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityRequestBuilder provides operations to call the getStaffAvailability method.
type BookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityRequestBuilderInternal instantiates a new BookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityRequestBuilder and sets the default values.
func NewBookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityRequestBuilder) {
    m := &BookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/bookingBusinesses/{bookingBusiness%2Did}/getStaffAvailability", pathParameters),
    }
    return m
}
// NewBookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityRequestBuilder instantiates a new BookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityRequestBuilder and sets the default values.
func NewBookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityRequestBuilderInternal(urlParams, requestAdapter)
}
// Post get the availability information of staff members of a Microsoft Bookings calendar.
// Deprecated: This method is obsolete. Use PostAsGetStaffAvailabilityPostResponse instead.
// returns a BookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/bookingbusiness-getstaffavailability?view=graph-rest-beta
func (m *BookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityRequestBuilder) Post(ctx context.Context, body BookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityPostRequestBodyable, requestConfiguration *BookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityRequestBuilderPostRequestConfiguration)(BookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateBookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(BookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityResponseable), nil
}
// PostAsGetStaffAvailabilityPostResponse get the availability information of staff members of a Microsoft Bookings calendar.
// returns a BookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/bookingbusiness-getstaffavailability?view=graph-rest-beta
func (m *BookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityRequestBuilder) PostAsGetStaffAvailabilityPostResponse(ctx context.Context, body BookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityPostRequestBodyable, requestConfiguration *BookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityRequestBuilderPostRequestConfiguration)(BookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateBookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(BookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityPostResponseable), nil
}
// ToPostRequestInformation get the availability information of staff members of a Microsoft Bookings calendar.
// returns a *RequestInformation when successful
func (m *BookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityRequestBuilder) ToPostRequestInformation(ctx context.Context, body BookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityPostRequestBodyable, requestConfiguration *BookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *BookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityRequestBuilder when successful
func (m *BookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityRequestBuilder) WithUrl(rawUrl string)(*BookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityRequestBuilder) {
    return NewBookingbusinessesItemGetstaffavailabilityGetStaffAvailabilityRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
