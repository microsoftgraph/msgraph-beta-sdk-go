package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsCountRequestBuilder provides operations to count the resources in the collection.
type ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsCountRequestBuilderGetQueryParameters get the number of the resource
type ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsCountRequestBuilderGetQueryParameters
}
// NewItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsCountRequestBuilder) {
    m := &ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/pendingAccessReviewInstances/{accessReviewInstance%2Did}/stages/{accessReviewStage%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance/decisions/$count{?%24search,%24filter}", pathParameters),
    }
    return m
}
// NewItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
func (m *ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
func (m *ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "text/plain")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
