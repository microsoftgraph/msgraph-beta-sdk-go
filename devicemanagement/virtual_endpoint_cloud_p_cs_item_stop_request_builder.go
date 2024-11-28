package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointCloudPCsItemStopRequestBuilder provides operations to call the stop method.
type VirtualEndpointCloudPCsItemStopRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointCloudPCsItemStopRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointCloudPCsItemStopRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointCloudPCsItemStopRequestBuilderInternal instantiates a new VirtualEndpointCloudPCsItemStopRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsItemStopRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsItemStopRequestBuilder) {
    m := &VirtualEndpointCloudPCsItemStopRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/{cloudPC%2Did}/stop", pathParameters),
    }
    return m
}
// NewVirtualEndpointCloudPCsItemStopRequestBuilder instantiates a new VirtualEndpointCloudPCsItemStopRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsItemStopRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsItemStopRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointCloudPCsItemStopRequestBuilderInternal(urlParams, requestAdapter)
}
// Post stop a specific Cloud PC. Currently, only Windows 365 Frontline Cloud PCs are supported.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-stop?view=graph-rest-beta
func (m *VirtualEndpointCloudPCsItemStopRequestBuilder) Post(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsItemStopRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation stop a specific Cloud PC. Currently, only Windows 365 Frontline Cloud PCs are supported.
// returns a *RequestInformation when successful
func (m *VirtualEndpointCloudPCsItemStopRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsItemStopRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualEndpointCloudPCsItemStopRequestBuilder when successful
func (m *VirtualEndpointCloudPCsItemStopRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointCloudPCsItemStopRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemStopRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
