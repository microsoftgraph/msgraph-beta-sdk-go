package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder provides operations to call the evaluateRemoval method.
type InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilderInternal instantiates a new MicrosoftGraphSecurityEvaluateRemovalRequestBuilder and sets the default values.
func NewInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) {
    m := &InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/informationProtection/sensitivityLabels/microsoft.graph.security.evaluateRemoval", pathParameters),
    }
    return m
}
// NewInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder instantiates a new MicrosoftGraphSecurityEvaluateRemovalRequestBuilder and sets the default values.
func NewInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilderInternal(urlParams, requestAdapter)
}
// Post indicate to the consuming application what actions it should take to remove the label information. Given contentInfo as an input, which includes existing content metadata key-value pairs, the API returns an informationProtectionAction that contains some combination of one or more of the following:  This API is available in the following national cloud deployments.
// Deprecated: This method is obsolete. Use PostAsEvaluateRemovalPostResponse instead.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-sensitivitylabel-evaluateremoval?view=graph-rest-1.0
func (m *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) Post(ctx context.Context, body InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalPostRequestBodyable, requestConfiguration *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration)(InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalResponseable), nil
}
// PostAsEvaluateRemovalPostResponse indicate to the consuming application what actions it should take to remove the label information. Given contentInfo as an input, which includes existing content metadata key-value pairs, the API returns an informationProtectionAction that contains some combination of one or more of the following:  This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-sensitivitylabel-evaluateremoval?view=graph-rest-1.0
func (m *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) PostAsEvaluateRemovalPostResponse(ctx context.Context, body InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalPostRequestBodyable, requestConfiguration *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration)(InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalPostResponseable), nil
}
// ToPostRequestInformation indicate to the consuming application what actions it should take to remove the label information. Given contentInfo as an input, which includes existing content metadata key-value pairs, the API returns an informationProtectionAction that contains some combination of one or more of the following:  This API is available in the following national cloud deployments.
func (m *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) ToPostRequestInformation(ctx context.Context, body InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalEvaluateRemovalPostRequestBodyable, requestConfiguration *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) WithUrl(rawUrl string)(*InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) {
    return NewInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
