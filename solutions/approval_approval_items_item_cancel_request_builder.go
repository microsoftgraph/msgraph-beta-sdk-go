package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ApprovalApprovalItemsItemCancelRequestBuilder provides operations to call the cancel method.
type ApprovalApprovalItemsItemCancelRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ApprovalApprovalItemsItemCancelRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApprovalApprovalItemsItemCancelRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewApprovalApprovalItemsItemCancelRequestBuilderInternal instantiates a new ApprovalApprovalItemsItemCancelRequestBuilder and sets the default values.
func NewApprovalApprovalItemsItemCancelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApprovalApprovalItemsItemCancelRequestBuilder) {
    m := &ApprovalApprovalItemsItemCancelRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/approval/approvalItems/{approvalItem%2Did}/cancel", pathParameters),
    }
    return m
}
// NewApprovalApprovalItemsItemCancelRequestBuilder instantiates a new ApprovalApprovalItemsItemCancelRequestBuilder and sets the default values.
func NewApprovalApprovalItemsItemCancelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApprovalApprovalItemsItemCancelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApprovalApprovalItemsItemCancelRequestBuilderInternal(urlParams, requestAdapter)
}
// Post cancel the approval item. The owner of the approval can trigger this endpoint.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/approvalitem-cancel?view=graph-rest-beta
func (m *ApprovalApprovalItemsItemCancelRequestBuilder) Post(ctx context.Context, requestConfiguration *ApprovalApprovalItemsItemCancelRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation cancel the approval item. The owner of the approval can trigger this endpoint.
// returns a *RequestInformation when successful
func (m *ApprovalApprovalItemsItemCancelRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ApprovalApprovalItemsItemCancelRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ApprovalApprovalItemsItemCancelRequestBuilder when successful
func (m *ApprovalApprovalItemsItemCancelRequestBuilder) WithUrl(rawUrl string)(*ApprovalApprovalItemsItemCancelRequestBuilder) {
    return NewApprovalApprovalItemsItemCancelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
