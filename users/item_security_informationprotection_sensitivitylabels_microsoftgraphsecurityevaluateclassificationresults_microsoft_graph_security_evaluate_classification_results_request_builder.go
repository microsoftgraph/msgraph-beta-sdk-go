package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder provides operations to call the evaluateClassificationResults method.
type ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilderInternal instantiates a new ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder and sets the default values.
func NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder) {
    m := &ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/security/informationProtection/sensitivityLabels/microsoft.graph.security.evaluateClassificationResults", pathParameters),
    }
    return m
}
// NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder instantiates a new ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder and sets the default values.
func NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post use the classification results to compute the sensitivity label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set automatically based on classification of the file contents, rather than labeled directly by a user or service.  To evaluate based on classification results, provide the contentInfo, which includes existing content metadata key-value pairs, and classification results. The API returns an informationProtectionAction that contains one of more of the following:
// Deprecated: This method is obsolete. Use PostAsEvaluateClassificationResultsPostResponse instead.
// returns a ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-sensitivitylabel-evaluateclassificationresults?view=graph-rest-beta
func (m *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder) Post(ctx context.Context, body ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostRequestBodyable, requestConfiguration *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilderPostRequestConfiguration)(ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsResponseable), nil
}
// PostAsEvaluateClassificationResultsPostResponse use the classification results to compute the sensitivity label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set automatically based on classification of the file contents, rather than labeled directly by a user or service.  To evaluate based on classification results, provide the contentInfo, which includes existing content metadata key-value pairs, and classification results. The API returns an informationProtectionAction that contains one of more of the following:
// returns a ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-sensitivitylabel-evaluateclassificationresults?view=graph-rest-beta
func (m *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder) PostAsEvaluateClassificationResultsPostResponse(ctx context.Context, body ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostRequestBodyable, requestConfiguration *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilderPostRequestConfiguration)(ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostResponseable), nil
}
// ToPostRequestInformation use the classification results to compute the sensitivity label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set automatically based on classification of the file contents, rather than labeled directly by a user or service.  To evaluate based on classification results, provide the contentInfo, which includes existing content metadata key-value pairs, and classification results. The API returns an informationProtectionAction that contains one of more of the following:
// returns a *RequestInformation when successful
func (m *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsEvaluateClassificationResultsPostRequestBodyable, requestConfiguration *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder when successful
func (m *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder) WithUrl(rawUrl string)(*ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder) {
    return NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateclassificationresultsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
