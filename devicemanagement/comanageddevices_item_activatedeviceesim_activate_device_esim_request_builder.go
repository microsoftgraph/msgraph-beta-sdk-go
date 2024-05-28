package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComanageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilder provides operations to call the activateDeviceEsim method.
type ComanageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ComanageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewComanageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilderInternal instantiates a new ComanageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilder and sets the default values.
func NewComanageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilder) {
    m := &ComanageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}/activateDeviceEsim", pathParameters),
    }
    return m
}
// NewComanageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilder instantiates a new ComanageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilder and sets the default values.
func NewComanageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilderInternal(urlParams, requestAdapter)
}
// Post activate eSIM on the device.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilder) Post(ctx context.Context, body ComanageddevicesItemActivatedeviceesimActivateDeviceEsimPostRequestBodyable, requestConfiguration *ComanageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation activate eSIM on the device.
// returns a *RequestInformation when successful
func (m *ComanageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilder) ToPostRequestInformation(ctx context.Context, body ComanageddevicesItemActivatedeviceesimActivateDeviceEsimPostRequestBodyable, requestConfiguration *ComanageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ComanageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilder when successful
func (m *ComanageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilder) WithUrl(rawUrl string)(*ComanageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilder) {
    return NewComanageddevicesItemActivatedeviceesimActivateDeviceEsimRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
