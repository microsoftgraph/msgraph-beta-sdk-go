package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OperationapprovalrequestsItemCancelapprovalCancelApprovalRequestBuilder provides operations to call the cancelApproval method.
type OperationapprovalrequestsItemCancelapprovalCancelApprovalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OperationapprovalrequestsItemCancelapprovalCancelApprovalRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OperationapprovalrequestsItemCancelapprovalCancelApprovalRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOperationapprovalrequestsItemCancelapprovalCancelApprovalRequestBuilderInternal instantiates a new OperationapprovalrequestsItemCancelapprovalCancelApprovalRequestBuilder and sets the default values.
func NewOperationapprovalrequestsItemCancelapprovalCancelApprovalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationapprovalrequestsItemCancelapprovalCancelApprovalRequestBuilder) {
    m := &OperationapprovalrequestsItemCancelapprovalCancelApprovalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/operationApprovalRequests/{operationApprovalRequest%2Did}/cancelApproval", pathParameters),
    }
    return m
}
// NewOperationapprovalrequestsItemCancelapprovalCancelApprovalRequestBuilder instantiates a new OperationapprovalrequestsItemCancelapprovalCancelApprovalRequestBuilder and sets the default values.
func NewOperationapprovalrequestsItemCancelapprovalCancelApprovalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationapprovalrequestsItemCancelapprovalCancelApprovalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOperationapprovalrequestsItemCancelapprovalCancelApprovalRequestBuilderInternal(urlParams, requestAdapter)
}
// Post cancels an already approved instance of an operationApprovalRequest.
// Deprecated: This method is obsolete. Use PostAsCancelApprovalPostResponse instead.
// returns a OperationapprovalrequestsItemCancelapprovalCancelApprovalResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationapprovalrequestsItemCancelapprovalCancelApprovalRequestBuilder) Post(ctx context.Context, body OperationapprovalrequestsItemCancelapprovalCancelApprovalPostRequestBodyable, requestConfiguration *OperationapprovalrequestsItemCancelapprovalCancelApprovalRequestBuilderPostRequestConfiguration)(OperationapprovalrequestsItemCancelapprovalCancelApprovalResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOperationapprovalrequestsItemCancelapprovalCancelApprovalResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OperationapprovalrequestsItemCancelapprovalCancelApprovalResponseable), nil
}
// PostAsCancelApprovalPostResponse cancels an already approved instance of an operationApprovalRequest.
// returns a OperationapprovalrequestsItemCancelapprovalCancelApprovalPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationapprovalrequestsItemCancelapprovalCancelApprovalRequestBuilder) PostAsCancelApprovalPostResponse(ctx context.Context, body OperationapprovalrequestsItemCancelapprovalCancelApprovalPostRequestBodyable, requestConfiguration *OperationapprovalrequestsItemCancelapprovalCancelApprovalRequestBuilderPostRequestConfiguration)(OperationapprovalrequestsItemCancelapprovalCancelApprovalPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOperationapprovalrequestsItemCancelapprovalCancelApprovalPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OperationapprovalrequestsItemCancelapprovalCancelApprovalPostResponseable), nil
}
// ToPostRequestInformation cancels an already approved instance of an operationApprovalRequest.
// returns a *RequestInformation when successful
func (m *OperationapprovalrequestsItemCancelapprovalCancelApprovalRequestBuilder) ToPostRequestInformation(ctx context.Context, body OperationapprovalrequestsItemCancelapprovalCancelApprovalPostRequestBodyable, requestConfiguration *OperationapprovalrequestsItemCancelapprovalCancelApprovalRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *OperationapprovalrequestsItemCancelapprovalCancelApprovalRequestBuilder when successful
func (m *OperationapprovalrequestsItemCancelapprovalCancelApprovalRequestBuilder) WithUrl(rawUrl string)(*OperationapprovalrequestsItemCancelapprovalCancelApprovalRequestBuilder) {
    return NewOperationapprovalrequestsItemCancelapprovalCancelApprovalRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
