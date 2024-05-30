package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComanageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilder provides operations to call the initiateDeviceAttestation method.
type ComanageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ComanageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewComanageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilderInternal instantiates a new ComanageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilder and sets the default values.
func NewComanageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilder) {
    m := &ComanageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}/initiateDeviceAttestation", pathParameters),
    }
    return m
}
// NewComanageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilder instantiates a new ComanageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilder and sets the default values.
func NewComanageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post perform Device Attestation
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilder) Post(ctx context.Context, requestConfiguration *ComanageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation perform Device Attestation
// returns a *RequestInformation when successful
func (m *ComanageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ComanageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ComanageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilder when successful
func (m *ComanageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilder) WithUrl(rawUrl string)(*ComanageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilder) {
    return NewComanageddevicesItemInitiatedeviceattestationInitiateDeviceAttestationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
