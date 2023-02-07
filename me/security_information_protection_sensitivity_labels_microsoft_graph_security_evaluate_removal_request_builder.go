package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder provides operations to call the evaluateRemoval method.
type SecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilderInternal instantiates a new MicrosoftGraphSecurityEvaluateRemovalRequestBuilder and sets the default values.
func NewSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) {
    m := &SecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/security/informationProtection/sensitivityLabels/microsoft.graph.security.evaluateRemoval";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder instantiates a new MicrosoftGraphSecurityEvaluateRemovalRequestBuilder and sets the default values.
func NewSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilderInternal(urlParams, requestAdapter)
}
// Post indicate to the consuming application what actions it should take to remove the label information. Given contentInfo as an input, which includes existing content metadata key-value pairs, the API returns an informationProtectionAction that contains some combination of one or more of the following: 
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/security-sensitivitylabel-evaluateremoval?view=graph-rest-1.0
func (m *SecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) Post(ctx context.Context, body SecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalPostRequestBodyable, requestConfiguration *SecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration)(SecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(SecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalResponseable), nil
}
// ToPostRequestInformation indicate to the consuming application what actions it should take to remove the label information. Given contentInfo as an input, which includes existing content metadata key-value pairs, the API returns an informationProtectionAction that contains some combination of one or more of the following: 
func (m *SecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) ToPostRequestInformation(ctx context.Context, body SecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalPostRequestBodyable, requestConfiguration *SecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
