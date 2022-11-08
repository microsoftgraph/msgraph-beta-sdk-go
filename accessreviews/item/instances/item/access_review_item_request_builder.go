package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i427417357745b8948faf1189b247d6d7532aaf80f445dbe4ae64088c48d587ef "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/instances/item/stop"
    i582c25d0b7cb97c8a565f6cf5d70558f6ba5b94d6324ffe4eadbd370e08df083 "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/instances/item/resetdecisions"
    i65c01fa596a41558434812572afb4c05db36b5cfbea9355b7b2d68eba6548691 "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/instances/item/mydecisions"
    i7bc12143e353b4d4f5d4853894578fef285042b09d3a2bf02a3b93e71752ac83 "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/instances/item/reviewers"
    i8ab53548eb87e2b488fd6dda1768c4ad29fd4e638c24a409daf072d07ff44804 "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/instances/item/decisions"
    i961781b3050f29063663fdf5b0570296ccb6be6f13a899851a5ba688dca06f3c "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/instances/item/sendreminder"
    ic27c7952d00624c2da5fd2be5a0a7aacc4f5e7a14917639e6c7f5b726f99ce45 "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/instances/item/applydecisions"
    i701f07b747d69019caad4740f3b1724263906917267c5af3244d67f4509f2b99 "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/instances/item/reviewers/item"
    i8daa9a485197443d7a67077ab2f41041a6f8236cb9437bc2446ae5f14432ea51 "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/instances/item/mydecisions/item"
    id3fe44f01780fcde4c699a6a202fc6e147e89515315aba30b2594c57bab487ac "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/instances/item/decisions/item"
)

// AccessReviewItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.accessReview entity.
type AccessReviewItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AccessReviewItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessReviewItemRequestBuilderGetQueryParameters the collection of access reviews instances past, present and future, if this object is a recurring access review.
type AccessReviewItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AccessReviewItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessReviewItemRequestBuilderGetQueryParameters
}
// AccessReviewItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApplyDecisions provides operations to call the applyDecisions method.
func (m *AccessReviewItemRequestBuilder) ApplyDecisions()(*ic27c7952d00624c2da5fd2be5a0a7aacc4f5e7a14917639e6c7f5b726f99ce45.ApplyDecisionsRequestBuilder) {
    return ic27c7952d00624c2da5fd2be5a0a7aacc4f5e7a14917639e6c7f5b726f99ce45.NewApplyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAccessReviewItemRequestBuilderInternal instantiates a new AccessReviewItemRequestBuilder and sets the default values.
func NewAccessReviewItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewItemRequestBuilder) {
    m := &AccessReviewItemRequestBuilder{
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
// NewAccessReviewItemRequestBuilder instantiates a new AccessReviewItemRequestBuilder and sets the default values.
func NewAccessReviewItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property instances for accessReviews
func (m *AccessReviewItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *AccessReviewItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of access reviews instances past, present and future, if this object is a recurring access review.
func (m *AccessReviewItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *AccessReviewItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property instances in accessReviews
func (m *AccessReviewItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewable, requestConfiguration *AccessReviewItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Decisions provides operations to manage the decisions property of the microsoft.graph.accessReview entity.
func (m *AccessReviewItemRequestBuilder) Decisions()(*i8ab53548eb87e2b488fd6dda1768c4ad29fd4e638c24a409daf072d07ff44804.DecisionsRequestBuilder) {
    return i8ab53548eb87e2b488fd6dda1768c4ad29fd4e638c24a409daf072d07ff44804.NewDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DecisionsById provides operations to manage the decisions property of the microsoft.graph.accessReview entity.
func (m *AccessReviewItemRequestBuilder) DecisionsById(id string)(*id3fe44f01780fcde4c699a6a202fc6e147e89515315aba30b2594c57bab487ac.AccessReviewDecisionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewDecision%2Did"] = id
    }
    return id3fe44f01780fcde4c699a6a202fc6e147e89515315aba30b2594c57bab487ac.NewAccessReviewDecisionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property instances for accessReviews
func (m *AccessReviewItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AccessReviewItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *AccessReviewItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AccessReviewItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewable, error) {
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
func (m *AccessReviewItemRequestBuilder) MyDecisions()(*i65c01fa596a41558434812572afb4c05db36b5cfbea9355b7b2d68eba6548691.MyDecisionsRequestBuilder) {
    return i65c01fa596a41558434812572afb4c05db36b5cfbea9355b7b2d68eba6548691.NewMyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MyDecisionsById provides operations to manage the myDecisions property of the microsoft.graph.accessReview entity.
func (m *AccessReviewItemRequestBuilder) MyDecisionsById(id string)(*i8daa9a485197443d7a67077ab2f41041a6f8236cb9437bc2446ae5f14432ea51.AccessReviewDecisionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewDecision%2Did"] = id
    }
    return i8daa9a485197443d7a67077ab2f41041a6f8236cb9437bc2446ae5f14432ea51.NewAccessReviewDecisionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in accessReviews
func (m *AccessReviewItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewable, requestConfiguration *AccessReviewItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewable, error) {
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
func (m *AccessReviewItemRequestBuilder) ResetDecisions()(*i582c25d0b7cb97c8a565f6cf5d70558f6ba5b94d6324ffe4eadbd370e08df083.ResetDecisionsRequestBuilder) {
    return i582c25d0b7cb97c8a565f6cf5d70558f6ba5b94d6324ffe4eadbd370e08df083.NewResetDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reviewers provides operations to manage the reviewers property of the microsoft.graph.accessReview entity.
func (m *AccessReviewItemRequestBuilder) Reviewers()(*i7bc12143e353b4d4f5d4853894578fef285042b09d3a2bf02a3b93e71752ac83.ReviewersRequestBuilder) {
    return i7bc12143e353b4d4f5d4853894578fef285042b09d3a2bf02a3b93e71752ac83.NewReviewersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReviewersById provides operations to manage the reviewers property of the microsoft.graph.accessReview entity.
func (m *AccessReviewItemRequestBuilder) ReviewersById(id string)(*i701f07b747d69019caad4740f3b1724263906917267c5af3244d67f4509f2b99.AccessReviewReviewerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewReviewer%2Did"] = id
    }
    return i701f07b747d69019caad4740f3b1724263906917267c5af3244d67f4509f2b99.NewAccessReviewReviewerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SendReminder provides operations to call the sendReminder method.
func (m *AccessReviewItemRequestBuilder) SendReminder()(*i961781b3050f29063663fdf5b0570296ccb6be6f13a899851a5ba688dca06f3c.SendReminderRequestBuilder) {
    return i961781b3050f29063663fdf5b0570296ccb6be6f13a899851a5ba688dca06f3c.NewSendReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Stop provides operations to call the stop method.
func (m *AccessReviewItemRequestBuilder) Stop()(*i427417357745b8948faf1189b247d6d7532aaf80f445dbe4ae64088c48d587ef.StopRequestBuilder) {
    return i427417357745b8948faf1189b247d6d7532aaf80f445dbe4ae64088c48d587ef.NewStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
