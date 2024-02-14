package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BookingBusinessesItemAppointmentsItemCancelRequestBuilder provides operations to call the cancel method.
type BookingBusinessesItemAppointmentsItemCancelRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BookingBusinessesItemAppointmentsItemCancelRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingBusinessesItemAppointmentsItemCancelRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBookingBusinessesItemAppointmentsItemCancelRequestBuilderInternal instantiates a new BookingBusinessesItemAppointmentsItemCancelRequestBuilder and sets the default values.
func NewBookingBusinessesItemAppointmentsItemCancelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingBusinessesItemAppointmentsItemCancelRequestBuilder) {
    m := &BookingBusinessesItemAppointmentsItemCancelRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/bookingBusinesses/{bookingBusiness%2Did}/appointments/{bookingAppointment%2Did}/cancel", pathParameters),
    }
    return m
}
// NewBookingBusinessesItemAppointmentsItemCancelRequestBuilder instantiates a new BookingBusinessesItemAppointmentsItemCancelRequestBuilder and sets the default values.
func NewBookingBusinessesItemAppointmentsItemCancelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingBusinessesItemAppointmentsItemCancelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBookingBusinessesItemAppointmentsItemCancelRequestBuilderInternal(urlParams, requestAdapter)
}
// Post cancel the specified bookingAppointment in the specified bookingBusiness, and send a message to the involved customer and staff members.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/bookingappointment-cancel?view=graph-rest-1.0
func (m *BookingBusinessesItemAppointmentsItemCancelRequestBuilder) Post(ctx context.Context, body BookingBusinessesItemAppointmentsItemCancelPostRequestBodyable, requestConfiguration *BookingBusinessesItemAppointmentsItemCancelRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation cancel the specified bookingAppointment in the specified bookingBusiness, and send a message to the involved customer and staff members.
// returns a *RequestInformation when successful
func (m *BookingBusinessesItemAppointmentsItemCancelRequestBuilder) ToPostRequestInformation(ctx context.Context, body BookingBusinessesItemAppointmentsItemCancelPostRequestBodyable, requestConfiguration *BookingBusinessesItemAppointmentsItemCancelRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BookingBusinessesItemAppointmentsItemCancelRequestBuilder when successful
func (m *BookingBusinessesItemAppointmentsItemCancelRequestBuilder) WithUrl(rawUrl string)(*BookingBusinessesItemAppointmentsItemCancelRequestBuilder) {
    return NewBookingBusinessesItemAppointmentsItemCancelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
