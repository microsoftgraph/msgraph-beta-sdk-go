package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityGovernanceAccessReviewsDecisionsItemInstanceStopRequestBuilder provides operations to call the stop method.
type IdentityGovernanceAccessReviewsDecisionsItemInstanceStopRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceAccessReviewsDecisionsItemInstanceStopRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceAccessReviewsDecisionsItemInstanceStopRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIdentityGovernanceAccessReviewsDecisionsItemInstanceStopRequestBuilderInternal instantiates a new StopRequestBuilder and sets the default values.
func NewIdentityGovernanceAccessReviewsDecisionsItemInstanceStopRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceAccessReviewsDecisionsItemInstanceStopRequestBuilder) {
    m := &IdentityGovernanceAccessReviewsDecisionsItemInstanceStopRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/accessReviews/decisions/{accessReviewInstanceDecisionItem%2Did}/instance/microsoft.graph.stop";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceAccessReviewsDecisionsItemInstanceStopRequestBuilder instantiates a new StopRequestBuilder and sets the default values.
func NewIdentityGovernanceAccessReviewsDecisionsItemInstanceStopRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceAccessReviewsDecisionsItemInstanceStopRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceAccessReviewsDecisionsItemInstanceStopRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation stop a currently active accessReviewInstance. After the access review instance stops, the instance status will be `Completed`, the reviewers can no longer give input, and the access review decisions can be applied. Stopping an instance will not effect future instances. To prevent a recurring access review from starting future instances, update the schedule definition to change its scheduled end date.
func (m *IdentityGovernanceAccessReviewsDecisionsItemInstanceStopRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceAccessReviewsDecisionsItemInstanceStopRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post stop a currently active accessReviewInstance. After the access review instance stops, the instance status will be `Completed`, the reviewers can no longer give input, and the access review decisions can be applied. Stopping an instance will not effect future instances. To prevent a recurring access review from starting future instances, update the schedule definition to change its scheduled end date.
func (m *IdentityGovernanceAccessReviewsDecisionsItemInstanceStopRequestBuilder) Post(ctx context.Context, requestConfiguration *IdentityGovernanceAccessReviewsDecisionsItemInstanceStopRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, requestConfiguration);
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
