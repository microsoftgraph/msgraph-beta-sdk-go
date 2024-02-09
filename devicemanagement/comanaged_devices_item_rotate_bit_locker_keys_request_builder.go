package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComanagedDevicesItemRotateBitLockerKeysRequestBuilder provides operations to call the rotateBitLockerKeys method.
type ComanagedDevicesItemRotateBitLockerKeysRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ComanagedDevicesItemRotateBitLockerKeysRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanagedDevicesItemRotateBitLockerKeysRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewComanagedDevicesItemRotateBitLockerKeysRequestBuilderInternal instantiates a new ComanagedDevicesItemRotateBitLockerKeysRequestBuilder and sets the default values.
func NewComanagedDevicesItemRotateBitLockerKeysRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesItemRotateBitLockerKeysRequestBuilder) {
    m := &ComanagedDevicesItemRotateBitLockerKeysRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}/rotateBitLockerKeys", pathParameters),
    }
    return m
}
// NewComanagedDevicesItemRotateBitLockerKeysRequestBuilder instantiates a new ComanagedDevicesItemRotateBitLockerKeysRequestBuilder and sets the default values.
func NewComanagedDevicesItemRotateBitLockerKeysRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesItemRotateBitLockerKeysRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanagedDevicesItemRotateBitLockerKeysRequestBuilderInternal(urlParams, requestAdapter)
}
// Post rotate BitLockerKeys
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanagedDevicesItemRotateBitLockerKeysRequestBuilder) Post(ctx context.Context, requestConfiguration *ComanagedDevicesItemRotateBitLockerKeysRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation rotate BitLockerKeys
// returns a *RequestInformation when successful
func (m *ComanagedDevicesItemRotateBitLockerKeysRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ComanagedDevicesItemRotateBitLockerKeysRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ComanagedDevicesItemRotateBitLockerKeysRequestBuilder when successful
func (m *ComanagedDevicesItemRotateBitLockerKeysRequestBuilder) WithUrl(rawUrl string)(*ComanagedDevicesItemRotateBitLockerKeysRequestBuilder) {
    return NewComanagedDevicesItemRotateBitLockerKeysRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
