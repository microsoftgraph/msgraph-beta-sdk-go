package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder provides operations to call the evaluateRemoval method.
type InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilderInternal instantiates a new InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder and sets the default values.
func NewInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) {
    m := &InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/informationProtection/sensitivityLabels/microsoft.graph.security.evaluateRemoval", pathParameters),
    }
    return m
}
// NewInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder instantiates a new InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder and sets the default values.
func NewInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilderInternal(urlParams, requestAdapter)
}
// Post indicate to the consuming application what actions it should take to remove the label information. Given contentInfo as an input, which includes existing content metadata key-value pairs, the API returns an informationProtectionAction that contains some combination of one or more of the following: 
// Deprecated: This method is obsolete. Use PostAsEvaluateRemovalPostResponse instead.
// returns a InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-sensitivitylabel-evaluateremoval?view=graph-rest-beta
func (m *InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) Post(ctx context.Context, body InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostRequestBodyable, requestConfiguration *InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration)(InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalResponseable), nil
}
// PostAsEvaluateRemovalPostResponse indicate to the consuming application what actions it should take to remove the label information. Given contentInfo as an input, which includes existing content metadata key-value pairs, the API returns an informationProtectionAction that contains some combination of one or more of the following: 
// returns a InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-sensitivitylabel-evaluateremoval?view=graph-rest-beta
func (m *InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) PostAsEvaluateRemovalPostResponse(ctx context.Context, body InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostRequestBodyable, requestConfiguration *InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration)(InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostResponseable), nil
}
// ToPostRequestInformation indicate to the consuming application what actions it should take to remove the label information. Given contentInfo as an input, which includes existing content metadata key-value pairs, the API returns an informationProtectionAction that contains some combination of one or more of the following: 
// returns a *RequestInformation when successful
func (m *InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) ToPostRequestInformation(ctx context.Context, body InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalEvaluateRemovalPostRequestBodyable, requestConfiguration *InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder when successful
func (m *InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) WithUrl(rawUrl string)(*InformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) {
    return NewInformationprotectionSensitivitylabelsMicrosoftgraphsecurityevaluateremovalMicrosoftGraphSecurityEvaluateRemovalRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
