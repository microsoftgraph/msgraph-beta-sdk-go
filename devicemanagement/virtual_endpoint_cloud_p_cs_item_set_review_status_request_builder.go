package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointCloudPCsItemSetReviewStatusRequestBuilder provides operations to call the setReviewStatus method.
type VirtualEndpointCloudPCsItemSetReviewStatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointCloudPCsItemSetReviewStatusRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointCloudPCsItemSetReviewStatusRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointCloudPCsItemSetReviewStatusRequestBuilderInternal instantiates a new VirtualEndpointCloudPCsItemSetReviewStatusRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsItemSetReviewStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsItemSetReviewStatusRequestBuilder) {
    m := &VirtualEndpointCloudPCsItemSetReviewStatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/{cloudPC%2Did}/setReviewStatus", pathParameters),
    }
    return m
}
// NewVirtualEndpointCloudPCsItemSetReviewStatusRequestBuilder instantiates a new VirtualEndpointCloudPCsItemSetReviewStatusRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsItemSetReviewStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsItemSetReviewStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointCloudPCsItemSetReviewStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action setReviewStatus
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEndpointCloudPCsItemSetReviewStatusRequestBuilder) Post(ctx context.Context, body VirtualEndpointCloudPCsItemSetReviewStatusPostRequestBodyable, requestConfiguration *VirtualEndpointCloudPCsItemSetReviewStatusRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action setReviewStatus
// returns a *RequestInformation when successful
func (m *VirtualEndpointCloudPCsItemSetReviewStatusRequestBuilder) ToPostRequestInformation(ctx context.Context, body VirtualEndpointCloudPCsItemSetReviewStatusPostRequestBodyable, requestConfiguration *VirtualEndpointCloudPCsItemSetReviewStatusRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualEndpointCloudPCsItemSetReviewStatusRequestBuilder when successful
func (m *VirtualEndpointCloudPCsItemSetReviewStatusRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointCloudPCsItemSetReviewStatusRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemSetReviewStatusRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
