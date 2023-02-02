package approvalworkflowproviders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsRequestBuilder provides operations to call the recordDecisions method.
type ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsRequestBuilderInternal instantiates a new RecordDecisionsRequestBuilder and sets the default values.
func NewItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsRequestBuilder) {
    m := &ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/approvalWorkflowProviders/{approvalWorkflowProvider%2Did}/businessFlowsWithRequestsAwaitingMyDecision/{businessFlow%2Did}/microsoft.graph.recordDecisions";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsRequestBuilder instantiates a new RecordDecisionsRequestBuilder and sets the default values.
func NewItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action recordDecisions
func (m *ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsRequestBuilder) Post(ctx context.Context, body ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsPostRequestBodyable, requestConfiguration *ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action recordDecisions
func (m *ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsPostRequestBodyable, requestConfiguration *ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
