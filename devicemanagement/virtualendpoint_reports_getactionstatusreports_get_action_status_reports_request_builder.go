package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualendpointReportsGetactionstatusreportsGetActionStatusReportsRequestBuilder provides operations to call the getActionStatusReports method.
type VirtualendpointReportsGetactionstatusreportsGetActionStatusReportsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointReportsGetactionstatusreportsGetActionStatusReportsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointReportsGetactionstatusreportsGetActionStatusReportsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualendpointReportsGetactionstatusreportsGetActionStatusReportsRequestBuilderInternal instantiates a new VirtualendpointReportsGetactionstatusreportsGetActionStatusReportsRequestBuilder and sets the default values.
func NewVirtualendpointReportsGetactionstatusreportsGetActionStatusReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointReportsGetactionstatusreportsGetActionStatusReportsRequestBuilder) {
    m := &VirtualendpointReportsGetactionstatusreportsGetActionStatusReportsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/reports/getActionStatusReports", pathParameters),
    }
    return m
}
// NewVirtualendpointReportsGetactionstatusreportsGetActionStatusReportsRequestBuilder instantiates a new VirtualendpointReportsGetactionstatusreportsGetActionStatusReportsRequestBuilder and sets the default values.
func NewVirtualendpointReportsGetactionstatusreportsGetActionStatusReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointReportsGetactionstatusreportsGetActionStatusReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointReportsGetactionstatusreportsGetActionStatusReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post get the remote action status reports, including data such as the Cloud PC ID, Cloud PC device display name, initiating user's principal name, device owner's user principal name, action taken, and action state.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcreports-getactionstatusreports?view=graph-rest-beta
func (m *VirtualendpointReportsGetactionstatusreportsGetActionStatusReportsRequestBuilder) Post(ctx context.Context, body VirtualendpointReportsGetactionstatusreportsGetActionStatusReportsPostRequestBodyable, requestConfiguration *VirtualendpointReportsGetactionstatusreportsGetActionStatusReportsRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToPostRequestInformation get the remote action status reports, including data such as the Cloud PC ID, Cloud PC device display name, initiating user's principal name, device owner's user principal name, action taken, and action state.
// returns a *RequestInformation when successful
func (m *VirtualendpointReportsGetactionstatusreportsGetActionStatusReportsRequestBuilder) ToPostRequestInformation(ctx context.Context, body VirtualendpointReportsGetactionstatusreportsGetActionStatusReportsPostRequestBodyable, requestConfiguration *VirtualendpointReportsGetactionstatusreportsGetActionStatusReportsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualendpointReportsGetactionstatusreportsGetActionStatusReportsRequestBuilder when successful
func (m *VirtualendpointReportsGetactionstatusreportsGetActionStatusReportsRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointReportsGetactionstatusreportsGetActionStatusReportsRequestBuilder) {
    return NewVirtualendpointReportsGetactionstatusreportsGetActionStatusReportsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
