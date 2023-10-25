package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComanagedDevicesItemReprovisionCloudPcRequestBuilder provides operations to call the reprovisionCloudPc method.
type ComanagedDevicesItemReprovisionCloudPcRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ComanagedDevicesItemReprovisionCloudPcRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanagedDevicesItemReprovisionCloudPcRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewComanagedDevicesItemReprovisionCloudPcRequestBuilderInternal instantiates a new ReprovisionCloudPcRequestBuilder and sets the default values.
func NewComanagedDevicesItemReprovisionCloudPcRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesItemReprovisionCloudPcRequestBuilder) {
    m := &ComanagedDevicesItemReprovisionCloudPcRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}/reprovisionCloudPc", pathParameters),
    }
    return m
}
// NewComanagedDevicesItemReprovisionCloudPcRequestBuilder instantiates a new ReprovisionCloudPcRequestBuilder and sets the default values.
func NewComanagedDevicesItemReprovisionCloudPcRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesItemReprovisionCloudPcRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanagedDevicesItemReprovisionCloudPcRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reprovision a Cloud PC with an Intune managed device ID. This API is available in the following national cloud deployments.
// Deprecated: The reprovisionCloudPc API is deprecated and will stop returning on Sep 30, 2023. Please use reprovision instead as of 2023-07/reprovisionCloudPc on 2023-07-17 and will be removed 2023-09-30
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/manageddevice-reprovisioncloudpc?view=graph-rest-1.0
func (m *ComanagedDevicesItemReprovisionCloudPcRequestBuilder) Post(ctx context.Context, requestConfiguration *ComanagedDevicesItemReprovisionCloudPcRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation reprovision a Cloud PC with an Intune managed device ID. This API is available in the following national cloud deployments.
// Deprecated: The reprovisionCloudPc API is deprecated and will stop returning on Sep 30, 2023. Please use reprovision instead as of 2023-07/reprovisionCloudPc on 2023-07-17 and will be removed 2023-09-30
func (m *ComanagedDevicesItemReprovisionCloudPcRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ComanagedDevicesItemReprovisionCloudPcRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The reprovisionCloudPc API is deprecated and will stop returning on Sep 30, 2023. Please use reprovision instead as of 2023-07/reprovisionCloudPc on 2023-07-17 and will be removed 2023-09-30
func (m *ComanagedDevicesItemReprovisionCloudPcRequestBuilder) WithUrl(rawUrl string)(*ComanagedDevicesItemReprovisionCloudPcRequestBuilder) {
    return NewComanagedDevicesItemReprovisionCloudPcRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
