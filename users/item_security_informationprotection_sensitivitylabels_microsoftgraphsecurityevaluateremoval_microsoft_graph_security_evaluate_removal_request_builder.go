package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder provides operations to call the evaluateRemoval method.
type ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilderInternal instantiates a new ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder and sets the default values.
func NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) {
    m := &ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/security/informationProtection/sensitivityLabels/microsoft.graph.security.evaluateRemoval", pathParameters),
    }
    return m
}
// NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder instantiates a new ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder and sets the default values.
func NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilderInternal(urlParams, requestAdapter)
}
// Post indicate to the consuming application what actions it should take to remove the label information. Given contentInfo as an input, which includes existing content metadata key-value pairs, the API returns an informationProtectionAction that contains some combination of one or more of the following: 
// Deprecated: This method is obsolete. Use PostAsEvaluateRemovalPostResponse instead.
// returns a ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-sensitivitylabel-evaluateremoval?view=graph-rest-beta
func (m *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) Post(ctx context.Context, body ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostRequestBodyable, requestConfiguration *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration)(ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalResponseable), nil
}
// PostAsEvaluateRemovalPostResponse indicate to the consuming application what actions it should take to remove the label information. Given contentInfo as an input, which includes existing content metadata key-value pairs, the API returns an informationProtectionAction that contains some combination of one or more of the following: 
// returns a ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-sensitivitylabel-evaluateremoval?view=graph-rest-beta
func (m *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) PostAsEvaluateRemovalPostResponse(ctx context.Context, body ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostRequestBodyable, requestConfiguration *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration)(ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostResponseable), nil
}
// ToPostRequestInformation indicate to the consuming application what actions it should take to remove the label information. Given contentInfo as an input, which includes existing content metadata key-value pairs, the API returns an informationProtectionAction that contains some combination of one or more of the following: 
// returns a *RequestInformation when successful
func (m *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostRequestBodyable, requestConfiguration *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder when successful
func (m *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) WithUrl(rawUrl string)(*ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) {
    return NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
