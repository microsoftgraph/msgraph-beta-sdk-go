package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsRequestBuilder provides operations to call the evaluateClassificationResults method.
type ItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsRequestBuilderInternal instantiates a new ItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsRequestBuilder and sets the default values.
func NewItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsRequestBuilder) {
    m := &ItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/informationProtection/policy/labels/evaluateClassificationResults", pathParameters),
    }
    return m
}
// NewItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsRequestBuilder instantiates a new ItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsRequestBuilder and sets the default values.
func NewItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post using classification results, compute the information protection label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set automatically based on classification of the file contents, rather than labeled directly by a user or service. To evaluate based on classification results, provide contentInfo, which includes existing content metadata key/value pairs, and classification results. The API returns an informationProtectionAction that contains one of more of the following:
// Deprecated: This method is obsolete. Use PostAsEvaluateClassificationResultsPostResponse instead.
// returns a ItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/informationprotectionlabel-evaluateclassificationresults?view=graph-rest-beta
func (m *ItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsRequestBuilder) Post(ctx context.Context, body ItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsPostRequestBodyable, requestConfiguration *ItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsRequestBuilderPostRequestConfiguration)(ItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsResponseable), nil
}
// PostAsEvaluateClassificationResultsPostResponse using classification results, compute the information protection label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set automatically based on classification of the file contents, rather than labeled directly by a user or service. To evaluate based on classification results, provide contentInfo, which includes existing content metadata key/value pairs, and classification results. The API returns an informationProtectionAction that contains one of more of the following:
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels
// returns a ItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/informationprotectionlabel-evaluateclassificationresults?view=graph-rest-beta
func (m *ItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsRequestBuilder) PostAsEvaluateClassificationResultsPostResponse(ctx context.Context, body ItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsPostRequestBodyable, requestConfiguration *ItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsRequestBuilderPostRequestConfiguration)(ItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsPostResponseable), nil
}
// ToPostRequestInformation using classification results, compute the information protection label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set automatically based on classification of the file contents, rather than labeled directly by a user or service. To evaluate based on classification results, provide contentInfo, which includes existing content metadata key/value pairs, and classification results. The API returns an informationProtectionAction that contains one of more of the following:
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels
// returns a *RequestInformation when successful
func (m *ItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsPostRequestBodyable, requestConfiguration *ItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsRequestBuilder when successful
func (m *ItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsRequestBuilder) WithUrl(rawUrl string)(*ItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsRequestBuilder) {
    return NewItemInformationprotectionPolicyLabelsEvaluateclassificationresultsEvaluateClassificationResultsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
