package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder provides operations to call the evaluateClassificationResults method.
type InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilderInternal instantiates a new InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder and sets the default values.
func NewInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder) {
    m := &InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/informationProtection/sensitivityLabels/microsoft.graph.security.evaluateClassificationResults", pathParameters),
    }
    return m
}
// NewInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder instantiates a new InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder and sets the default values.
func NewInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post use the classification results to compute the sensitivity label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set automatically based on classification of the file contents, rather than labeled directly by a user or service.  To evaluate based on classification results, provide the contentInfo, which includes existing content metadata key-value pairs, and classification results. The API returns an informationProtectionAction that contains one of more of the following:
// Deprecated: This method is obsolete. Use PostAsEvaluateClassificationResultsPostResponse instead.
// returns a InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-sensitivitylabel-evaluateclassificationresults?view=graph-rest-beta
func (m *InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder) Post(ctx context.Context, body InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostRequestBodyable, requestConfiguration *InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilderPostRequestConfiguration)(InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsResponseable), nil
}
// PostAsEvaluateClassificationResultsPostResponse use the classification results to compute the sensitivity label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set automatically based on classification of the file contents, rather than labeled directly by a user or service.  To evaluate based on classification results, provide the contentInfo, which includes existing content metadata key-value pairs, and classification results. The API returns an informationProtectionAction that contains one of more of the following:
// returns a InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-sensitivitylabel-evaluateclassificationresults?view=graph-rest-beta
func (m *InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder) PostAsEvaluateClassificationResultsPostResponse(ctx context.Context, body InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostRequestBodyable, requestConfiguration *InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilderPostRequestConfiguration)(InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostResponseable), nil
}
// ToPostRequestInformation use the classification results to compute the sensitivity label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set automatically based on classification of the file contents, rather than labeled directly by a user or service.  To evaluate based on classification results, provide the contentInfo, which includes existing content metadata key-value pairs, and classification results. The API returns an informationProtectionAction that contains one of more of the following:
// returns a *RequestInformation when successful
func (m *InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder) ToPostRequestInformation(ctx context.Context, body InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostRequestBodyable, requestConfiguration *InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder when successful
func (m *InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder) WithUrl(rawUrl string)(*InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder) {
    return NewInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
