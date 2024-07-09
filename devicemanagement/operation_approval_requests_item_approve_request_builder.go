package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OperationApprovalRequestsItemApproveRequestBuilder provides operations to call the approve method.
type OperationApprovalRequestsItemApproveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OperationApprovalRequestsItemApproveRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OperationApprovalRequestsItemApproveRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOperationApprovalRequestsItemApproveRequestBuilderInternal instantiates a new OperationApprovalRequestsItemApproveRequestBuilder and sets the default values.
func NewOperationApprovalRequestsItemApproveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationApprovalRequestsItemApproveRequestBuilder) {
    m := &OperationApprovalRequestsItemApproveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/operationApprovalRequests/{operationApprovalRequest%2Did}/approve", pathParameters),
    }
    return m
}
// NewOperationApprovalRequestsItemApproveRequestBuilder instantiates a new OperationApprovalRequestsItemApproveRequestBuilder and sets the default values.
func NewOperationApprovalRequestsItemApproveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationApprovalRequestsItemApproveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOperationApprovalRequestsItemApproveRequestBuilderInternal(urlParams, requestAdapter)
}
// Post approves the requested instance of an operationApprovalRequest.
// Deprecated: This method is obsolete. Use PostAsApprovePostResponse instead.
// returns a OperationApprovalRequestsItemApproveResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationApprovalRequestsItemApproveRequestBuilder) Post(ctx context.Context, body OperationApprovalRequestsItemApprovePostRequestBodyable, requestConfiguration *OperationApprovalRequestsItemApproveRequestBuilderPostRequestConfiguration)(OperationApprovalRequestsItemApproveResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOperationApprovalRequestsItemApproveResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OperationApprovalRequestsItemApproveResponseable), nil
}
// PostAsApprovePostResponse approves the requested instance of an operationApprovalRequest.
// returns a OperationApprovalRequestsItemApprovePostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationApprovalRequestsItemApproveRequestBuilder) PostAsApprovePostResponse(ctx context.Context, body OperationApprovalRequestsItemApprovePostRequestBodyable, requestConfiguration *OperationApprovalRequestsItemApproveRequestBuilderPostRequestConfiguration)(OperationApprovalRequestsItemApprovePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOperationApprovalRequestsItemApprovePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OperationApprovalRequestsItemApprovePostResponseable), nil
}
// ToPostRequestInformation approves the requested instance of an operationApprovalRequest.
// returns a *RequestInformation when successful
func (m *OperationApprovalRequestsItemApproveRequestBuilder) ToPostRequestInformation(ctx context.Context, body OperationApprovalRequestsItemApprovePostRequestBodyable, requestConfiguration *OperationApprovalRequestsItemApproveRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *OperationApprovalRequestsItemApproveRequestBuilder when successful
func (m *OperationApprovalRequestsItemApproveRequestBuilder) WithUrl(rawUrl string)(*OperationApprovalRequestsItemApproveRequestBuilder) {
    return NewOperationApprovalRequestsItemApproveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
