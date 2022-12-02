package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder provides operations to manage the decisions property of the microsoft.graph.accessReviewInstance entity.
type UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetQueryParameters each user reviewed in an accessReviewInstance has a decision item representing if they were approved, denied, or not yet reviewed.
type UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetQueryParameters
}
// UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderInternal instantiates a new AccessReviewInstanceDecisionItemItemRequestBuilder and sets the default values.
func NewUsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) {
    m := &UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/pendingAccessReviewInstances/{accessReviewInstance%2Did}/stages/{accessReviewStage%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance/decisions/{accessReviewInstanceDecisionItem%2Did1}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder instantiates a new AccessReviewInstanceDecisionItemItemRequestBuilder and sets the default values.
func NewUsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property decisions for users
func (m *UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation each user reviewed in an accessReviewInstance has a decision item representing if they were approved, denied, or not yet reviewed.
func (m *UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property decisions in users
func (m *UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable, requestConfiguration *UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property decisions for users
func (m *UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get each user reviewed in an accessReviewInstance has a decision item representing if they were approved, denied, or not yet reviewed.
func (m *UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewInstanceDecisionItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable), nil
}
// Insights provides operations to manage the insights property of the microsoft.graph.accessReviewInstanceDecisionItem entity.
func (m *UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) Insights()(*UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilder) {
    return NewUsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InsightsById provides operations to manage the insights property of the microsoft.graph.accessReviewInstanceDecisionItem entity.
func (m *UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) InsightsById(id string)(*UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceInsight%2Did"] = id
    }
    return NewUsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsItemInsightsGovernanceInsightItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property decisions in users
func (m *UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable, requestConfiguration *UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewInstanceDecisionItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable), nil
}
