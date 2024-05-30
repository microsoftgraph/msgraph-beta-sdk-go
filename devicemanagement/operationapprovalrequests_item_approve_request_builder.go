package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OperationapprovalrequestsItemApproveRequestBuilder provides operations to call the approve method.
type OperationapprovalrequestsItemApproveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OperationapprovalrequestsItemApproveRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OperationapprovalrequestsItemApproveRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOperationapprovalrequestsItemApproveRequestBuilderInternal instantiates a new OperationapprovalrequestsItemApproveRequestBuilder and sets the default values.
func NewOperationapprovalrequestsItemApproveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationapprovalrequestsItemApproveRequestBuilder) {
    m := &OperationapprovalrequestsItemApproveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/operationApprovalRequests/{operationApprovalRequest%2Did}/approve", pathParameters),
    }
    return m
}
// NewOperationapprovalrequestsItemApproveRequestBuilder instantiates a new OperationapprovalrequestsItemApproveRequestBuilder and sets the default values.
func NewOperationapprovalrequestsItemApproveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationapprovalrequestsItemApproveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOperationapprovalrequestsItemApproveRequestBuilderInternal(urlParams, requestAdapter)
}
// Post approves the requested instance of an operationApprovalRequest.
// Deprecated: This method is obsolete. Use PostAsApprovePostResponse instead.
// returns a OperationapprovalrequestsItemApproveResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationapprovalrequestsItemApproveRequestBuilder) Post(ctx context.Context, body OperationapprovalrequestsItemApprovePostRequestBodyable, requestConfiguration *OperationapprovalrequestsItemApproveRequestBuilderPostRequestConfiguration)(OperationapprovalrequestsItemApproveResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOperationapprovalrequestsItemApproveResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OperationapprovalrequestsItemApproveResponseable), nil
}
// PostAsApprovePostResponse approves the requested instance of an operationApprovalRequest.
// returns a OperationapprovalrequestsItemApprovePostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationapprovalrequestsItemApproveRequestBuilder) PostAsApprovePostResponse(ctx context.Context, body OperationapprovalrequestsItemApprovePostRequestBodyable, requestConfiguration *OperationapprovalrequestsItemApproveRequestBuilderPostRequestConfiguration)(OperationapprovalrequestsItemApprovePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOperationapprovalrequestsItemApprovePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OperationapprovalrequestsItemApprovePostResponseable), nil
}
// ToPostRequestInformation approves the requested instance of an operationApprovalRequest.
// returns a *RequestInformation when successful
func (m *OperationapprovalrequestsItemApproveRequestBuilder) ToPostRequestInformation(ctx context.Context, body OperationapprovalrequestsItemApprovePostRequestBodyable, requestConfiguration *OperationapprovalrequestsItemApproveRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *OperationapprovalrequestsItemApproveRequestBuilder when successful
func (m *OperationapprovalrequestsItemApproveRequestBuilder) WithUrl(rawUrl string)(*OperationapprovalrequestsItemApproveRequestBuilder) {
    return NewOperationapprovalrequestsItemApproveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
