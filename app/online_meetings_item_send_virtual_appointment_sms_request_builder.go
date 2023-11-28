package app

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OnlineMeetingsItemSendVirtualAppointmentSmsRequestBuilder provides operations to call the sendVirtualAppointmentSms method.
type OnlineMeetingsItemSendVirtualAppointmentSmsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OnlineMeetingsItemSendVirtualAppointmentSmsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlineMeetingsItemSendVirtualAppointmentSmsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOnlineMeetingsItemSendVirtualAppointmentSmsRequestBuilderInternal instantiates a new SendVirtualAppointmentSmsRequestBuilder and sets the default values.
func NewOnlineMeetingsItemSendVirtualAppointmentSmsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlineMeetingsItemSendVirtualAppointmentSmsRequestBuilder) {
    m := &OnlineMeetingsItemSendVirtualAppointmentSmsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/app/onlineMeetings/{onlineMeeting%2Did}/sendVirtualAppointmentSms", pathParameters),
    }
    return m
}
// NewOnlineMeetingsItemSendVirtualAppointmentSmsRequestBuilder instantiates a new SendVirtualAppointmentSmsRequestBuilder and sets the default values.
func NewOnlineMeetingsItemSendVirtualAppointmentSmsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlineMeetingsItemSendVirtualAppointmentSmsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnlineMeetingsItemSendVirtualAppointmentSmsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action sendVirtualAppointmentSms
func (m *OnlineMeetingsItemSendVirtualAppointmentSmsRequestBuilder) Post(ctx context.Context, body OnlineMeetingsItemSendVirtualAppointmentSmsPostRequestBodyable, requestConfiguration *OnlineMeetingsItemSendVirtualAppointmentSmsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation invoke action sendVirtualAppointmentSms
func (m *OnlineMeetingsItemSendVirtualAppointmentSmsRequestBuilder) ToPostRequestInformation(ctx context.Context, body OnlineMeetingsItemSendVirtualAppointmentSmsPostRequestBodyable, requestConfiguration *OnlineMeetingsItemSendVirtualAppointmentSmsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *OnlineMeetingsItemSendVirtualAppointmentSmsRequestBuilder) WithUrl(rawUrl string)(*OnlineMeetingsItemSendVirtualAppointmentSmsRequestBuilder) {
    return NewOnlineMeetingsItemSendVirtualAppointmentSmsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
