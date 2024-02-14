package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComanagedDevicesItemDeprovisionRequestBuilder provides operations to call the deprovision method.
type ComanagedDevicesItemDeprovisionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ComanagedDevicesItemDeprovisionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanagedDevicesItemDeprovisionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewComanagedDevicesItemDeprovisionRequestBuilderInternal instantiates a new ComanagedDevicesItemDeprovisionRequestBuilder and sets the default values.
func NewComanagedDevicesItemDeprovisionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesItemDeprovisionRequestBuilder) {
    m := &ComanagedDevicesItemDeprovisionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}/deprovision", pathParameters),
    }
    return m
}
// NewComanagedDevicesItemDeprovisionRequestBuilder instantiates a new ComanagedDevicesItemDeprovisionRequestBuilder and sets the default values.
func NewComanagedDevicesItemDeprovisionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesItemDeprovisionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanagedDevicesItemDeprovisionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action deprovision
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanagedDevicesItemDeprovisionRequestBuilder) Post(ctx context.Context, body ComanagedDevicesItemDeprovisionPostRequestBodyable, requestConfiguration *ComanagedDevicesItemDeprovisionRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action deprovision
// returns a *RequestInformation when successful
func (m *ComanagedDevicesItemDeprovisionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ComanagedDevicesItemDeprovisionPostRequestBodyable, requestConfiguration *ComanagedDevicesItemDeprovisionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ComanagedDevicesItemDeprovisionRequestBuilder when successful
func (m *ComanagedDevicesItemDeprovisionRequestBuilder) WithUrl(rawUrl string)(*ComanagedDevicesItemDeprovisionRequestBuilder) {
    return NewComanagedDevicesItemDeprovisionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
