package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OperationApprovalRequestsItemCancelApprovalRequestBuilder provides operations to call the cancelApproval method.
type OperationApprovalRequestsItemCancelApprovalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OperationApprovalRequestsItemCancelApprovalRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OperationApprovalRequestsItemCancelApprovalRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOperationApprovalRequestsItemCancelApprovalRequestBuilderInternal instantiates a new OperationApprovalRequestsItemCancelApprovalRequestBuilder and sets the default values.
func NewOperationApprovalRequestsItemCancelApprovalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationApprovalRequestsItemCancelApprovalRequestBuilder) {
    m := &OperationApprovalRequestsItemCancelApprovalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/operationApprovalRequests/{operationApprovalRequest%2Did}/cancelApproval", pathParameters),
    }
    return m
}
// NewOperationApprovalRequestsItemCancelApprovalRequestBuilder instantiates a new OperationApprovalRequestsItemCancelApprovalRequestBuilder and sets the default values.
func NewOperationApprovalRequestsItemCancelApprovalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationApprovalRequestsItemCancelApprovalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOperationApprovalRequestsItemCancelApprovalRequestBuilderInternal(urlParams, requestAdapter)
}
// Post cancels an already approved instance of an operationApprovalRequest.
// Deprecated: This method is obsolete. Use {TypeName} instead.
// returns a OperationApprovalRequestsItemCancelApprovalResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationApprovalRequestsItemCancelApprovalRequestBuilder) Post(ctx context.Context, body OperationApprovalRequestsItemCancelApprovalPostRequestBodyable, requestConfiguration *OperationApprovalRequestsItemCancelApprovalRequestBuilderPostRequestConfiguration)(OperationApprovalRequestsItemCancelApprovalResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOperationApprovalRequestsItemCancelApprovalResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OperationApprovalRequestsItemCancelApprovalResponseable), nil
}
// PostAsCancelApprovalPostResponse cancels an already approved instance of an operationApprovalRequest.
// returns a OperationApprovalRequestsItemCancelApprovalPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationApprovalRequestsItemCancelApprovalRequestBuilder) PostAsCancelApprovalPostResponse(ctx context.Context, body OperationApprovalRequestsItemCancelApprovalPostRequestBodyable, requestConfiguration *OperationApprovalRequestsItemCancelApprovalRequestBuilderPostRequestConfiguration)(OperationApprovalRequestsItemCancelApprovalPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOperationApprovalRequestsItemCancelApprovalPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OperationApprovalRequestsItemCancelApprovalPostResponseable), nil
}
// ToPostRequestInformation cancels an already approved instance of an operationApprovalRequest.
// returns a *RequestInformation when successful
func (m *OperationApprovalRequestsItemCancelApprovalRequestBuilder) ToPostRequestInformation(ctx context.Context, body OperationApprovalRequestsItemCancelApprovalPostRequestBodyable, requestConfiguration *OperationApprovalRequestsItemCancelApprovalRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *OperationApprovalRequestsItemCancelApprovalRequestBuilder when successful
func (m *OperationApprovalRequestsItemCancelApprovalRequestBuilder) WithUrl(rawUrl string)(*OperationApprovalRequestsItemCancelApprovalRequestBuilder) {
    return NewOperationApprovalRequestsItemCancelApprovalRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
