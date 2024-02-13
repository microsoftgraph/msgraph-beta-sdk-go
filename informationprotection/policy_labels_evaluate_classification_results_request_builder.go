package informationprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PolicyLabelsEvaluateClassificationResultsRequestBuilder provides operations to call the evaluateClassificationResults method.
type PolicyLabelsEvaluateClassificationResultsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PolicyLabelsEvaluateClassificationResultsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PolicyLabelsEvaluateClassificationResultsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPolicyLabelsEvaluateClassificationResultsRequestBuilderInternal instantiates a new PolicyLabelsEvaluateClassificationResultsRequestBuilder and sets the default values.
func NewPolicyLabelsEvaluateClassificationResultsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PolicyLabelsEvaluateClassificationResultsRequestBuilder) {
    m := &PolicyLabelsEvaluateClassificationResultsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/informationProtection/policy/labels/evaluateClassificationResults", pathParameters),
    }
    return m
}
// NewPolicyLabelsEvaluateClassificationResultsRequestBuilder instantiates a new PolicyLabelsEvaluateClassificationResultsRequestBuilder and sets the default values.
func NewPolicyLabelsEvaluateClassificationResultsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PolicyLabelsEvaluateClassificationResultsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPolicyLabelsEvaluateClassificationResultsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post using classification results, compute the information protection label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set automatically based on classification of the file contents, rather than labeled directly by a user or service. To evaluate based on classification results, provide contentInfo, which includes existing content metadata key/value pairs, and classification results. The API returns an informationProtectionAction that contains one of more of the following:
// Deprecated: This method is obsolete. Use {TypeName} instead.
// returns a PolicyLabelsEvaluateClassificationResultsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/informationprotectionlabel-evaluateclassificationresults?view=graph-rest-1.0
func (m *PolicyLabelsEvaluateClassificationResultsRequestBuilder) Post(ctx context.Context, body PolicyLabelsEvaluateClassificationResultsPostRequestBodyable, requestConfiguration *PolicyLabelsEvaluateClassificationResultsRequestBuilderPostRequestConfiguration)(PolicyLabelsEvaluateClassificationResultsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreatePolicyLabelsEvaluateClassificationResultsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(PolicyLabelsEvaluateClassificationResultsResponseable), nil
}
// PostAsEvaluateClassificationResultsPostResponse using classification results, compute the information protection label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set automatically based on classification of the file contents, rather than labeled directly by a user or service. To evaluate based on classification results, provide contentInfo, which includes existing content metadata key/value pairs, and classification results. The API returns an informationProtectionAction that contains one of more of the following:
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels
// returns a PolicyLabelsEvaluateClassificationResultsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/informationprotectionlabel-evaluateclassificationresults?view=graph-rest-1.0
func (m *PolicyLabelsEvaluateClassificationResultsRequestBuilder) PostAsEvaluateClassificationResultsPostResponse(ctx context.Context, body PolicyLabelsEvaluateClassificationResultsPostRequestBodyable, requestConfiguration *PolicyLabelsEvaluateClassificationResultsRequestBuilderPostRequestConfiguration)(PolicyLabelsEvaluateClassificationResultsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreatePolicyLabelsEvaluateClassificationResultsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(PolicyLabelsEvaluateClassificationResultsPostResponseable), nil
}
// ToPostRequestInformation using classification results, compute the information protection label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set automatically based on classification of the file contents, rather than labeled directly by a user or service. To evaluate based on classification results, provide contentInfo, which includes existing content metadata key/value pairs, and classification results. The API returns an informationProtectionAction that contains one of more of the following:
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels
// returns a *RequestInformation when successful
func (m *PolicyLabelsEvaluateClassificationResultsRequestBuilder) ToPostRequestInformation(ctx context.Context, body PolicyLabelsEvaluateClassificationResultsPostRequestBodyable, requestConfiguration *PolicyLabelsEvaluateClassificationResultsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels
// returns a *PolicyLabelsEvaluateClassificationResultsRequestBuilder when successful
func (m *PolicyLabelsEvaluateClassificationResultsRequestBuilder) WithUrl(rawUrl string)(*PolicyLabelsEvaluateClassificationResultsRequestBuilder) {
    return NewPolicyLabelsEvaluateClassificationResultsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
