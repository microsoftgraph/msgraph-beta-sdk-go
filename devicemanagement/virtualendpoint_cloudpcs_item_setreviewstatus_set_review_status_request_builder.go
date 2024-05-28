package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualendpointCloudpcsItemSetreviewstatusSetReviewStatusRequestBuilder provides operations to call the setReviewStatus method.
type VirtualendpointCloudpcsItemSetreviewstatusSetReviewStatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointCloudpcsItemSetreviewstatusSetReviewStatusRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointCloudpcsItemSetreviewstatusSetReviewStatusRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualendpointCloudpcsItemSetreviewstatusSetReviewStatusRequestBuilderInternal instantiates a new VirtualendpointCloudpcsItemSetreviewstatusSetReviewStatusRequestBuilder and sets the default values.
func NewVirtualendpointCloudpcsItemSetreviewstatusSetReviewStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointCloudpcsItemSetreviewstatusSetReviewStatusRequestBuilder) {
    m := &VirtualendpointCloudpcsItemSetreviewstatusSetReviewStatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/{cloudPC%2Did}/setReviewStatus", pathParameters),
    }
    return m
}
// NewVirtualendpointCloudpcsItemSetreviewstatusSetReviewStatusRequestBuilder instantiates a new VirtualendpointCloudpcsItemSetreviewstatusSetReviewStatusRequestBuilder and sets the default values.
func NewVirtualendpointCloudpcsItemSetreviewstatusSetReviewStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointCloudpcsItemSetreviewstatusSetReviewStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointCloudpcsItemSetreviewstatusSetReviewStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action setReviewStatus
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualendpointCloudpcsItemSetreviewstatusSetReviewStatusRequestBuilder) Post(ctx context.Context, body VirtualendpointCloudpcsItemSetreviewstatusSetReviewStatusPostRequestBodyable, requestConfiguration *VirtualendpointCloudpcsItemSetreviewstatusSetReviewStatusRequestBuilderPostRequestConfiguration)(error) {
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
func (m *VirtualendpointCloudpcsItemSetreviewstatusSetReviewStatusRequestBuilder) ToPostRequestInformation(ctx context.Context, body VirtualendpointCloudpcsItemSetreviewstatusSetReviewStatusPostRequestBodyable, requestConfiguration *VirtualendpointCloudpcsItemSetreviewstatusSetReviewStatusRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualendpointCloudpcsItemSetreviewstatusSetReviewStatusRequestBuilder when successful
func (m *VirtualendpointCloudpcsItemSetreviewstatusSetReviewStatusRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointCloudpcsItemSetreviewstatusSetReviewStatusRequestBuilder) {
    return NewVirtualendpointCloudpcsItemSetreviewstatusSetReviewStatusRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
