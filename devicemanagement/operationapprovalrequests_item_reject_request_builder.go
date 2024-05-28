package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OperationapprovalrequestsItemRejectRequestBuilder provides operations to call the reject method.
type OperationapprovalrequestsItemRejectRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OperationapprovalrequestsItemRejectRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OperationapprovalrequestsItemRejectRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOperationapprovalrequestsItemRejectRequestBuilderInternal instantiates a new OperationapprovalrequestsItemRejectRequestBuilder and sets the default values.
func NewOperationapprovalrequestsItemRejectRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationapprovalrequestsItemRejectRequestBuilder) {
    m := &OperationapprovalrequestsItemRejectRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/operationApprovalRequests/{operationApprovalRequest%2Did}/reject", pathParameters),
    }
    return m
}
// NewOperationapprovalrequestsItemRejectRequestBuilder instantiates a new OperationapprovalrequestsItemRejectRequestBuilder and sets the default values.
func NewOperationapprovalrequestsItemRejectRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationapprovalrequestsItemRejectRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOperationapprovalrequestsItemRejectRequestBuilderInternal(urlParams, requestAdapter)
}
// Post rejects the requested instance of an operationApprovalRequest.
// Deprecated: This method is obsolete. Use PostAsRejectPostResponse instead.
// returns a OperationapprovalrequestsItemRejectResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationapprovalrequestsItemRejectRequestBuilder) Post(ctx context.Context, body OperationapprovalrequestsItemRejectPostRequestBodyable, requestConfiguration *OperationapprovalrequestsItemRejectRequestBuilderPostRequestConfiguration)(OperationapprovalrequestsItemRejectResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOperationapprovalrequestsItemRejectResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OperationapprovalrequestsItemRejectResponseable), nil
}
// PostAsRejectPostResponse rejects the requested instance of an operationApprovalRequest.
// returns a OperationapprovalrequestsItemRejectPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationapprovalrequestsItemRejectRequestBuilder) PostAsRejectPostResponse(ctx context.Context, body OperationapprovalrequestsItemRejectPostRequestBodyable, requestConfiguration *OperationapprovalrequestsItemRejectRequestBuilderPostRequestConfiguration)(OperationapprovalrequestsItemRejectPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOperationapprovalrequestsItemRejectPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OperationapprovalrequestsItemRejectPostResponseable), nil
}
// ToPostRequestInformation rejects the requested instance of an operationApprovalRequest.
// returns a *RequestInformation when successful
func (m *OperationapprovalrequestsItemRejectRequestBuilder) ToPostRequestInformation(ctx context.Context, body OperationapprovalrequestsItemRejectPostRequestBodyable, requestConfiguration *OperationapprovalrequestsItemRejectRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *OperationapprovalrequestsItemRejectRequestBuilder when successful
func (m *OperationapprovalrequestsItemRejectRequestBuilder) WithUrl(rawUrl string)(*OperationapprovalrequestsItemRejectRequestBuilder) {
    return NewOperationapprovalrequestsItemRejectRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
