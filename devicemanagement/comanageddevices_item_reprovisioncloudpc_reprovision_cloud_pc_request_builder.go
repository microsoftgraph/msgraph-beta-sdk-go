package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComanageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilder provides operations to call the reprovisionCloudPc method.
type ComanageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ComanageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewComanageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilderInternal instantiates a new ComanageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilder and sets the default values.
func NewComanageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilder) {
    m := &ComanageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}/reprovisionCloudPc", pathParameters),
    }
    return m
}
// NewComanageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilder instantiates a new ComanageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilder and sets the default values.
func NewComanageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reprovision a Cloud PC with an Intune managed device ID.
// Deprecated: The reprovisionCloudPc API is deprecated and will stop returning on Sep 30, 2023. Please use reprovision instead as of 2023-07/reprovisionCloudPc
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/manageddevice-reprovisioncloudpc?view=graph-rest-beta
func (m *ComanageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilder) Post(ctx context.Context, requestConfiguration *ComanageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation reprovision a Cloud PC with an Intune managed device ID.
// Deprecated: The reprovisionCloudPc API is deprecated and will stop returning on Sep 30, 2023. Please use reprovision instead as of 2023-07/reprovisionCloudPc
// returns a *RequestInformation when successful
func (m *ComanageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ComanageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The reprovisionCloudPc API is deprecated and will stop returning on Sep 30, 2023. Please use reprovision instead as of 2023-07/reprovisionCloudPc
// returns a *ComanageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilder when successful
func (m *ComanageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilder) WithUrl(rawUrl string)(*ComanageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilder) {
    return NewComanageddevicesItemReprovisioncloudpcReprovisionCloudPcRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
