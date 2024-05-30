package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationMicrosoftGraphSecurityEvaluateApplicationRequestBuilder provides operations to call the evaluateApplication method.
type ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationMicrosoftGraphSecurityEvaluateApplicationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationMicrosoftGraphSecurityEvaluateApplicationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationMicrosoftGraphSecurityEvaluateApplicationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationMicrosoftGraphSecurityEvaluateApplicationRequestBuilderInternal instantiates a new ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationMicrosoftGraphSecurityEvaluateApplicationRequestBuilder and sets the default values.
func NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationMicrosoftGraphSecurityEvaluateApplicationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationMicrosoftGraphSecurityEvaluateApplicationRequestBuilder) {
    m := &ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationMicrosoftGraphSecurityEvaluateApplicationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/security/informationProtection/sensitivityLabels/microsoft.graph.security.evaluateApplication", pathParameters),
    }
    return m
}
// NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationMicrosoftGraphSecurityEvaluateApplicationRequestBuilder instantiates a new ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationMicrosoftGraphSecurityEvaluateApplicationRequestBuilder and sets the default values.
func NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationMicrosoftGraphSecurityEvaluateApplicationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationMicrosoftGraphSecurityEvaluateApplicationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationMicrosoftGraphSecurityEvaluateApplicationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post compute the sensitivity label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set manually or explicitly by a user or service, rather than automatically based on file contents. Given contentInfo, which includes existing content metadata key-value pairs, and labelingOptions as an input, the API returns an informationProtectionAction object that contains one of more of the following: 
// Deprecated: This method is obsolete. Use PostAsEvaluateApplicationPostResponse instead.
// returns a ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationEvaluateApplicationResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-sensitivitylabel-evaluateapplication?view=graph-rest-beta
func (m *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationMicrosoftGraphSecurityEvaluateApplicationRequestBuilder) Post(ctx context.Context, body ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationEvaluateApplicationPostRequestBodyable, requestConfiguration *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationMicrosoftGraphSecurityEvaluateApplicationRequestBuilderPostRequestConfiguration)(ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationEvaluateApplicationResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationEvaluateApplicationResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationEvaluateApplicationResponseable), nil
}
// PostAsEvaluateApplicationPostResponse compute the sensitivity label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set manually or explicitly by a user or service, rather than automatically based on file contents. Given contentInfo, which includes existing content metadata key-value pairs, and labelingOptions as an input, the API returns an informationProtectionAction object that contains one of more of the following: 
// returns a ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationEvaluateApplicationPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-sensitivitylabel-evaluateapplication?view=graph-rest-beta
func (m *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationMicrosoftGraphSecurityEvaluateApplicationRequestBuilder) PostAsEvaluateApplicationPostResponse(ctx context.Context, body ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationEvaluateApplicationPostRequestBodyable, requestConfiguration *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationMicrosoftGraphSecurityEvaluateApplicationRequestBuilderPostRequestConfiguration)(ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationEvaluateApplicationPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationEvaluateApplicationPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationEvaluateApplicationPostResponseable), nil
}
// ToPostRequestInformation compute the sensitivity label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set manually or explicitly by a user or service, rather than automatically based on file contents. Given contentInfo, which includes existing content metadata key-value pairs, and labelingOptions as an input, the API returns an informationProtectionAction object that contains one of more of the following: 
// returns a *RequestInformation when successful
func (m *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationMicrosoftGraphSecurityEvaluateApplicationRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationEvaluateApplicationPostRequestBodyable, requestConfiguration *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationMicrosoftGraphSecurityEvaluateApplicationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationMicrosoftGraphSecurityEvaluateApplicationRequestBuilder when successful
func (m *ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationMicrosoftGraphSecurityEvaluateApplicationRequestBuilder) WithUrl(rawUrl string)(*ItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationMicrosoftGraphSecurityEvaluateApplicationRequestBuilder) {
    return NewItemSecurityInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateapplicationMicrosoftGraphSecurityEvaluateApplicationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
