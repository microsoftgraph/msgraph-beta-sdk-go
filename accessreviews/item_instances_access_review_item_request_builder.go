package accessreviews

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemInstancesAccessReviewItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.accessReview entity.
type ItemInstancesAccessReviewItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemInstancesAccessReviewItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInstancesAccessReviewItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemInstancesAccessReviewItemRequestBuilderGetQueryParameters the collection of access reviews instances past, present and future, if this object is a recurring access review.
type ItemInstancesAccessReviewItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemInstancesAccessReviewItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInstancesAccessReviewItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemInstancesAccessReviewItemRequestBuilderGetQueryParameters
}
// ItemInstancesAccessReviewItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInstancesAccessReviewItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApplyDecisions provides operations to call the applyDecisions method.
func (m *ItemInstancesAccessReviewItemRequestBuilder) ApplyDecisions()(*ItemInstancesItemApplyDecisionsRequestBuilder) {
    return NewItemInstancesItemApplyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewItemInstancesAccessReviewItemRequestBuilderInternal instantiates a new AccessReviewItemRequestBuilder and sets the default values.
func NewItemInstancesAccessReviewItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInstancesAccessReviewItemRequestBuilder) {
    m := &ItemInstancesAccessReviewItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/accessReviews/{accessReview%2Did}/instances/{accessReview%2Did1}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemInstancesAccessReviewItemRequestBuilder instantiates a new AccessReviewItemRequestBuilder and sets the default values.
func NewItemInstancesAccessReviewItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInstancesAccessReviewItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInstancesAccessReviewItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property instances for accessReviews
func (m *ItemInstancesAccessReviewItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemInstancesAccessReviewItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the collection of access reviews instances past, present and future, if this object is a recurring access review.
func (m *ItemInstancesAccessReviewItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ItemInstancesAccessReviewItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property instances in accessReviews
func (m *ItemInstancesAccessReviewItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewable, requestConfiguration *ItemInstancesAccessReviewItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Decisions provides operations to manage the decisions property of the microsoft.graph.accessReview entity.
func (m *ItemInstancesAccessReviewItemRequestBuilder) Decisions()(*ItemInstancesItemDecisionsRequestBuilder) {
    return NewItemInstancesItemDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DecisionsById provides operations to manage the decisions property of the microsoft.graph.accessReview entity.
func (m *ItemInstancesAccessReviewItemRequestBuilder) DecisionsById(id string)(*ItemInstancesItemDecisionsAccessReviewDecisionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewDecision%2Did"] = id
    }
    return NewItemInstancesItemDecisionsAccessReviewDecisionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property instances for accessReviews
func (m *ItemInstancesAccessReviewItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemInstancesAccessReviewItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the collection of access reviews instances past, present and future, if this object is a recurring access review.
func (m *ItemInstancesAccessReviewItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemInstancesAccessReviewItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewable), nil
}
// MyDecisions provides operations to manage the myDecisions property of the microsoft.graph.accessReview entity.
func (m *ItemInstancesAccessReviewItemRequestBuilder) MyDecisions()(*ItemInstancesItemMyDecisionsRequestBuilder) {
    return NewItemInstancesItemMyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MyDecisionsById provides operations to manage the myDecisions property of the microsoft.graph.accessReview entity.
func (m *ItemInstancesAccessReviewItemRequestBuilder) MyDecisionsById(id string)(*ItemInstancesItemMyDecisionsAccessReviewDecisionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewDecision%2Did"] = id
    }
    return NewItemInstancesItemMyDecisionsAccessReviewDecisionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in accessReviews
func (m *ItemInstancesAccessReviewItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewable, requestConfiguration *ItemInstancesAccessReviewItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewable), nil
}
// ResetDecisions provides operations to call the resetDecisions method.
func (m *ItemInstancesAccessReviewItemRequestBuilder) ResetDecisions()(*ItemInstancesItemResetDecisionsRequestBuilder) {
    return NewItemInstancesItemResetDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reviewers provides operations to manage the reviewers property of the microsoft.graph.accessReview entity.
func (m *ItemInstancesAccessReviewItemRequestBuilder) Reviewers()(*ItemInstancesItemReviewersRequestBuilder) {
    return NewItemInstancesItemReviewersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReviewersById provides operations to manage the reviewers property of the microsoft.graph.accessReview entity.
func (m *ItemInstancesAccessReviewItemRequestBuilder) ReviewersById(id string)(*ItemInstancesItemReviewersAccessReviewReviewerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewReviewer%2Did"] = id
    }
    return NewItemInstancesItemReviewersAccessReviewReviewerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SendReminder provides operations to call the sendReminder method.
func (m *ItemInstancesAccessReviewItemRequestBuilder) SendReminder()(*ItemInstancesItemSendReminderRequestBuilder) {
    return NewItemInstancesItemSendReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Stop provides operations to call the stop method.
func (m *ItemInstancesAccessReviewItemRequestBuilder) Stop()(*ItemInstancesItemStopRequestBuilder) {
    return NewItemInstancesItemStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
