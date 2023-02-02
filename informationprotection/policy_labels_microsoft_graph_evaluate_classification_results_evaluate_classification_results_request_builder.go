package informationprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PolicyLabelsMicrosoftGraphEvaluateClassificationResultsEvaluateClassificationResultsRequestBuilder provides operations to call the evaluateClassificationResults method.
type PolicyLabelsMicrosoftGraphEvaluateClassificationResultsEvaluateClassificationResultsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PolicyLabelsMicrosoftGraphEvaluateClassificationResultsEvaluateClassificationResultsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PolicyLabelsMicrosoftGraphEvaluateClassificationResultsEvaluateClassificationResultsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPolicyLabelsMicrosoftGraphEvaluateClassificationResultsEvaluateClassificationResultsRequestBuilderInternal instantiates a new EvaluateClassificationResultsRequestBuilder and sets the default values.
func NewPolicyLabelsMicrosoftGraphEvaluateClassificationResultsEvaluateClassificationResultsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PolicyLabelsMicrosoftGraphEvaluateClassificationResultsEvaluateClassificationResultsRequestBuilder) {
    m := &PolicyLabelsMicrosoftGraphEvaluateClassificationResultsEvaluateClassificationResultsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/informationProtection/policy/labels/microsoft.graph.evaluateClassificationResults";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewPolicyLabelsMicrosoftGraphEvaluateClassificationResultsEvaluateClassificationResultsRequestBuilder instantiates a new EvaluateClassificationResultsRequestBuilder and sets the default values.
func NewPolicyLabelsMicrosoftGraphEvaluateClassificationResultsEvaluateClassificationResultsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PolicyLabelsMicrosoftGraphEvaluateClassificationResultsEvaluateClassificationResultsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPolicyLabelsMicrosoftGraphEvaluateClassificationResultsEvaluateClassificationResultsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post using classification results, compute the information protection label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set automatically based on classification of the file contents, rather than labeled directly by a user or service.  To evaluate based on classification results, provide contentInfo, which includes existing content metadata key/value pairs, and classification results. The API returns an informationProtectionAction that contains one of more of the following: 
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/informationprotectionlabel-evaluateclassificationresults?view=graph-rest-1.0
func (m *PolicyLabelsMicrosoftGraphEvaluateClassificationResultsEvaluateClassificationResultsRequestBuilder) Post(ctx context.Context, body PolicyLabelsMicrosoftGraphEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBodyable, requestConfiguration *PolicyLabelsMicrosoftGraphEvaluateClassificationResultsEvaluateClassificationResultsRequestBuilderPostRequestConfiguration)(PolicyLabelsMicrosoftGraphEvaluateClassificationResultsEvaluateClassificationResultsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreatePolicyLabelsMicrosoftGraphEvaluateClassificationResultsEvaluateClassificationResultsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(PolicyLabelsMicrosoftGraphEvaluateClassificationResultsEvaluateClassificationResultsResponseable), nil
}
// ToPostRequestInformation using classification results, compute the information protection label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set automatically based on classification of the file contents, rather than labeled directly by a user or service.  To evaluate based on classification results, provide contentInfo, which includes existing content metadata key/value pairs, and classification results. The API returns an informationProtectionAction that contains one of more of the following: 
func (m *PolicyLabelsMicrosoftGraphEvaluateClassificationResultsEvaluateClassificationResultsRequestBuilder) ToPostRequestInformation(ctx context.Context, body PolicyLabelsMicrosoftGraphEvaluateClassificationResultsEvaluateClassificationResultsPostRequestBodyable, requestConfiguration *PolicyLabelsMicrosoftGraphEvaluateClassificationResultsEvaluateClassificationResultsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
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
