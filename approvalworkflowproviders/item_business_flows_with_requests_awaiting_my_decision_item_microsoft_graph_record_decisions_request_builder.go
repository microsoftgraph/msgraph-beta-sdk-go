package approvalworkflowproviders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRequestBuilder provides operations to call the recordDecisions method.
type ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRequestBuilderInternal instantiates a new MicrosoftGraphRecordDecisionsRequestBuilder and sets the default values.
func NewItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRequestBuilder) {
    m := &ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRequestBuilder{
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
// NewItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRequestBuilder instantiates a new MicrosoftGraphRecordDecisionsRequestBuilder and sets the default values.
func NewItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action recordDecisions
func (m *ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRequestBuilder) Post(ctx context.Context, body ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsPostRequestBodyable, requestConfiguration *ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRequestBuilderPostRequestConfiguration)(error) {
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
func (m *ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsPostRequestBodyable, requestConfiguration *ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
