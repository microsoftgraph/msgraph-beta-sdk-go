package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilder provides operations to manage the insights property of the microsoft.graph.accessReviewInstanceDecisionItem entity.
type MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilderGetQueryParameters insights are recommendations to reviewers on whether to approve or deny a decision. There can be multiple insights associated with an accessReviewInstanceDecisionItem.
type MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
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
// MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilderGetQueryParameters
}
// MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilderInternal instantiates a new InsightsRequestBuilder and sets the default values.
func NewMePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilder) {
    m := &MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/pendingAccessReviewInstances/{accessReviewInstance%2Did}/stages/{accessReviewStage%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance/decisions/{accessReviewInstanceDecisionItem%2Did1}/insights{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilder instantiates a new InsightsRequestBuilder and sets the default values.
func NewMePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilder) Count()(*MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsCountRequestBuilder) {
    return NewMePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation insights are recommendations to reviewers on whether to approve or deny a decision. There can be multiple insights associated with an accessReviewInstanceDecisionItem.
func (m *MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePostRequestInformation create new navigation property to insights for me
func (m *MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceInsightable, requestConfiguration *MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get insights are recommendations to reviewers on whether to approve or deny a decision. There can be multiple insights associated with an accessReviewInstanceDecisionItem.
func (m *MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilder) Get(ctx context.Context, requestConfiguration *MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceInsightCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceInsightCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceInsightCollectionResponseable), nil
}
// Post create new navigation property to insights for me
func (m *MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceInsightable, requestConfiguration *MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceInsightable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceInsightFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceInsightable), nil
}
