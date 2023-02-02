package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsMicrosoftGraphFilterByCurrentUserWithOnFilterByCurrentUserWithOnRequestBuilder provides operations to call the filterByCurrentUser method.
type PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsMicrosoftGraphFilterByCurrentUserWithOnFilterByCurrentUserWithOnRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsMicrosoftGraphFilterByCurrentUserWithOnFilterByCurrentUserWithOnRequestBuilderGetQueryParameters invoke function filterByCurrentUser
type PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsMicrosoftGraphFilterByCurrentUserWithOnFilterByCurrentUserWithOnRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsMicrosoftGraphFilterByCurrentUserWithOnFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsMicrosoftGraphFilterByCurrentUserWithOnFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsMicrosoftGraphFilterByCurrentUserWithOnFilterByCurrentUserWithOnRequestBuilderGetQueryParameters
}
// NewPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsMicrosoftGraphFilterByCurrentUserWithOnFilterByCurrentUserWithOnRequestBuilderInternal instantiates a new FilterByCurrentUserWithOnRequestBuilder and sets the default values.
func NewPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsMicrosoftGraphFilterByCurrentUserWithOnFilterByCurrentUserWithOnRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, on *string)(*PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsMicrosoftGraphFilterByCurrentUserWithOnFilterByCurrentUserWithOnRequestBuilder) {
    m := &PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsMicrosoftGraphFilterByCurrentUserWithOnFilterByCurrentUserWithOnRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/pendingAccessReviewInstances/{accessReviewInstance%2Did}/stages/{accessReviewStage%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance/decisions/microsoft.graph.filterByCurrentUser(on='{on}'){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if on != nil {
        urlTplParams["on"] = *on
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsMicrosoftGraphFilterByCurrentUserWithOnFilterByCurrentUserWithOnRequestBuilder instantiates a new FilterByCurrentUserWithOnRequestBuilder and sets the default values.
func NewPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsMicrosoftGraphFilterByCurrentUserWithOnFilterByCurrentUserWithOnRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsMicrosoftGraphFilterByCurrentUserWithOnFilterByCurrentUserWithOnRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsMicrosoftGraphFilterByCurrentUserWithOnFilterByCurrentUserWithOnRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function filterByCurrentUser
func (m *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsMicrosoftGraphFilterByCurrentUserWithOnFilterByCurrentUserWithOnRequestBuilder) Get(ctx context.Context, requestConfiguration *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsMicrosoftGraphFilterByCurrentUserWithOnFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration)(PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsMicrosoftGraphFilterByCurrentUserWithOnFilterByCurrentUserWithOnResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreatePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsMicrosoftGraphFilterByCurrentUserWithOnFilterByCurrentUserWithOnResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsMicrosoftGraphFilterByCurrentUserWithOnFilterByCurrentUserWithOnResponseable), nil
}
// ToGetRequestInformation invoke function filterByCurrentUser
func (m *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsMicrosoftGraphFilterByCurrentUserWithOnFilterByCurrentUserWithOnRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsMicrosoftGraphFilterByCurrentUserWithOnFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
