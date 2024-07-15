package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEventsWebinarsItemRegistrationsWithEmailCancelRequestBuilder provides operations to call the cancel method.
type VirtualEventsWebinarsItemRegistrationsWithEmailCancelRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEventsWebinarsItemRegistrationsWithEmailCancelRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsWebinarsItemRegistrationsWithEmailCancelRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEventsWebinarsItemRegistrationsWithEmailCancelRequestBuilderInternal instantiates a new VirtualEventsWebinarsItemRegistrationsWithEmailCancelRequestBuilder and sets the default values.
func NewVirtualEventsWebinarsItemRegistrationsWithEmailCancelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsWebinarsItemRegistrationsWithEmailCancelRequestBuilder) {
    m := &VirtualEventsWebinarsItemRegistrationsWithEmailCancelRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/webinars/{virtualEventWebinar%2Did}/registrations(email='{email}')/cancel", pathParameters),
    }
    return m
}
// NewVirtualEventsWebinarsItemRegistrationsWithEmailCancelRequestBuilder instantiates a new VirtualEventsWebinarsItemRegistrationsWithEmailCancelRequestBuilder and sets the default values.
func NewVirtualEventsWebinarsItemRegistrationsWithEmailCancelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsWebinarsItemRegistrationsWithEmailCancelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEventsWebinarsItemRegistrationsWithEmailCancelRequestBuilderInternal(urlParams, requestAdapter)
}
// Post cancel a registrant's registration record for a webinar. 
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEventsWebinarsItemRegistrationsWithEmailCancelRequestBuilder) Post(ctx context.Context, requestConfiguration *VirtualEventsWebinarsItemRegistrationsWithEmailCancelRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation cancel a registrant's registration record for a webinar. 
// returns a *RequestInformation when successful
func (m *VirtualEventsWebinarsItemRegistrationsWithEmailCancelRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *VirtualEventsWebinarsItemRegistrationsWithEmailCancelRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualEventsWebinarsItemRegistrationsWithEmailCancelRequestBuilder when successful
func (m *VirtualEventsWebinarsItemRegistrationsWithEmailCancelRequestBuilder) WithUrl(rawUrl string)(*VirtualEventsWebinarsItemRegistrationsWithEmailCancelRequestBuilder) {
    return NewVirtualEventsWebinarsItemRegistrationsWithEmailCancelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
