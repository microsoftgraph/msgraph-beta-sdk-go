package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualendpointDeviceimagesItemReuploadRequestBuilder provides operations to call the reupload method.
type VirtualendpointDeviceimagesItemReuploadRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointDeviceimagesItemReuploadRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointDeviceimagesItemReuploadRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualendpointDeviceimagesItemReuploadRequestBuilderInternal instantiates a new VirtualendpointDeviceimagesItemReuploadRequestBuilder and sets the default values.
func NewVirtualendpointDeviceimagesItemReuploadRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointDeviceimagesItemReuploadRequestBuilder) {
    m := &VirtualendpointDeviceimagesItemReuploadRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/deviceImages/{cloudPcDeviceImage%2Did}/reupload", pathParameters),
    }
    return m
}
// NewVirtualendpointDeviceimagesItemReuploadRequestBuilder instantiates a new VirtualendpointDeviceimagesItemReuploadRequestBuilder and sets the default values.
func NewVirtualendpointDeviceimagesItemReuploadRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointDeviceimagesItemReuploadRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointDeviceimagesItemReuploadRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reupload a cloudPcDeviceImage object that failed to upload.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcdeviceimage-reupload?view=graph-rest-beta
func (m *VirtualendpointDeviceimagesItemReuploadRequestBuilder) Post(ctx context.Context, requestConfiguration *VirtualendpointDeviceimagesItemReuploadRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation reupload a cloudPcDeviceImage object that failed to upload.
// returns a *RequestInformation when successful
func (m *VirtualendpointDeviceimagesItemReuploadRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointDeviceimagesItemReuploadRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualendpointDeviceimagesItemReuploadRequestBuilder when successful
func (m *VirtualendpointDeviceimagesItemReuploadRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointDeviceimagesItemReuploadRequestBuilder) {
    return NewVirtualendpointDeviceimagesItemReuploadRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
