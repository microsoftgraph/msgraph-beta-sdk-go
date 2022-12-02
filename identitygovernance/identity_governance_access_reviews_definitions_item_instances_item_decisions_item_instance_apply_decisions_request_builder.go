package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityGovernanceAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceApplyDecisionsRequestBuilder provides operations to call the applyDecisions method.
type IdentityGovernanceAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceApplyDecisionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceApplyDecisionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceApplyDecisionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIdentityGovernanceAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceApplyDecisionsRequestBuilderInternal instantiates a new ApplyDecisionsRequestBuilder and sets the default values.
func NewIdentityGovernanceAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceApplyDecisionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceApplyDecisionsRequestBuilder) {
    m := &IdentityGovernanceAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceApplyDecisionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance/microsoft.graph.applyDecisions";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceApplyDecisionsRequestBuilder instantiates a new ApplyDecisionsRequestBuilder and sets the default values.
func NewIdentityGovernanceAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceApplyDecisionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceApplyDecisionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceApplyDecisionsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation apply review decisions on an accessReviewInstance if the decisions were not applied automatically because the autoApplyDecisionsEnabled property is `false` in the review's accessReviewScheduleSettings. The status of the accessReviewInstance must be `Completed` to call this method.
func (m *IdentityGovernanceAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceApplyDecisionsRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceApplyDecisionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post apply review decisions on an accessReviewInstance if the decisions were not applied automatically because the autoApplyDecisionsEnabled property is `false` in the review's accessReviewScheduleSettings. The status of the accessReviewInstance must be `Completed` to call this method.
func (m *IdentityGovernanceAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceApplyDecisionsRequestBuilder) Post(ctx context.Context, requestConfiguration *IdentityGovernanceAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceApplyDecisionsRequestBuilderPostRequestConfiguration)(error) {
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
