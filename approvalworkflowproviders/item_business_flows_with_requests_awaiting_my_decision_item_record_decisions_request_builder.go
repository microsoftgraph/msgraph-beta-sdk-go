package approvalworkflowproviders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsRequestBuilder provides operations to call the recordDecisions method.
type ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsRequestBuilderInternal instantiates a new RecordDecisionsRequestBuilder and sets the default values.
func NewItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsRequestBuilder) {
    m := &ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/approvalWorkflowProviders/{approvalWorkflowProvider%2Did}/businessFlowsWithRequestsAwaitingMyDecision/{businessFlow%2Did}/microsoft.graph.recordDecisions";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsRequestBuilder instantiates a new RecordDecisionsRequestBuilder and sets the default values.
func NewItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action recordDecisions
func (m *ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsPostRequestBodyable, requestConfiguration *ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action recordDecisions
func (m *ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsRequestBuilder) Post(ctx context.Context, body ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsPostRequestBodyable, requestConfiguration *ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
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
