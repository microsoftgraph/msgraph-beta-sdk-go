package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualeventsWebinarsItemRegistrationswithemailCancelRequestBuilder provides operations to call the cancel method.
type VirtualeventsWebinarsItemRegistrationswithemailCancelRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualeventsWebinarsItemRegistrationswithemailCancelRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualeventsWebinarsItemRegistrationswithemailCancelRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualeventsWebinarsItemRegistrationswithemailCancelRequestBuilderInternal instantiates a new VirtualeventsWebinarsItemRegistrationswithemailCancelRequestBuilder and sets the default values.
func NewVirtualeventsWebinarsItemRegistrationswithemailCancelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualeventsWebinarsItemRegistrationswithemailCancelRequestBuilder) {
    m := &VirtualeventsWebinarsItemRegistrationswithemailCancelRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/webinars/{virtualEventWebinar%2Did}/registrations(email='{email}')/cancel", pathParameters),
    }
    return m
}
// NewVirtualeventsWebinarsItemRegistrationswithemailCancelRequestBuilder instantiates a new VirtualeventsWebinarsItemRegistrationswithemailCancelRequestBuilder and sets the default values.
func NewVirtualeventsWebinarsItemRegistrationswithemailCancelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualeventsWebinarsItemRegistrationswithemailCancelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualeventsWebinarsItemRegistrationswithemailCancelRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action cancel
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualeventsWebinarsItemRegistrationswithemailCancelRequestBuilder) Post(ctx context.Context, requestConfiguration *VirtualeventsWebinarsItemRegistrationswithemailCancelRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action cancel
// returns a *RequestInformation when successful
func (m *VirtualeventsWebinarsItemRegistrationswithemailCancelRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *VirtualeventsWebinarsItemRegistrationswithemailCancelRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualeventsWebinarsItemRegistrationswithemailCancelRequestBuilder when successful
func (m *VirtualeventsWebinarsItemRegistrationswithemailCancelRequestBuilder) WithUrl(rawUrl string)(*VirtualeventsWebinarsItemRegistrationswithemailCancelRequestBuilder) {
    return NewVirtualeventsWebinarsItemRegistrationswithemailCancelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
