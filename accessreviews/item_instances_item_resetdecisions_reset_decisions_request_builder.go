package accessreviews

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemInstancesItemResetdecisionsResetDecisionsRequestBuilder provides operations to call the resetDecisions method.
type ItemInstancesItemResetdecisionsResetDecisionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInstancesItemResetdecisionsResetDecisionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInstancesItemResetdecisionsResetDecisionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemInstancesItemResetdecisionsResetDecisionsRequestBuilderInternal instantiates a new ItemInstancesItemResetdecisionsResetDecisionsRequestBuilder and sets the default values.
func NewItemInstancesItemResetdecisionsResetDecisionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInstancesItemResetdecisionsResetDecisionsRequestBuilder) {
    m := &ItemInstancesItemResetdecisionsResetDecisionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/accessReviews/{accessReview%2Did}/instances/{accessReview%2Did1}/resetDecisions", pathParameters),
    }
    return m
}
// NewItemInstancesItemResetdecisionsResetDecisionsRequestBuilder instantiates a new ItemInstancesItemResetdecisionsResetDecisionsRequestBuilder and sets the default values.
func NewItemInstancesItemResetdecisionsResetDecisionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInstancesItemResetdecisionsResetDecisionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInstancesItemResetdecisionsResetDecisionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post in the Microsoft Entra access reviews feature, reset the decisions of a currently active accessReview.  The target object can be either a one-time access review, or an instance of a recurring access review.  Previous decisions are no longer recorded, but reviewers can continue to update decisions.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreview-reset?view=graph-rest-beta
func (m *ItemInstancesItemResetdecisionsResetDecisionsRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemInstancesItemResetdecisionsResetDecisionsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation in the Microsoft Entra access reviews feature, reset the decisions of a currently active accessReview.  The target object can be either a one-time access review, or an instance of a recurring access review.  Previous decisions are no longer recorded, but reviewers can continue to update decisions.
// returns a *RequestInformation when successful
func (m *ItemInstancesItemResetdecisionsResetDecisionsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemInstancesItemResetdecisionsResetDecisionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemInstancesItemResetdecisionsResetDecisionsRequestBuilder when successful
func (m *ItemInstancesItemResetdecisionsResetDecisionsRequestBuilder) WithUrl(rawUrl string)(*ItemInstancesItemResetdecisionsResetDecisionsRequestBuilder) {
    return NewItemInstancesItemResetdecisionsResetDecisionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
