package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder provides operations to call the evaluateApplication method.
type InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilderInternal instantiates a new InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder and sets the default values.
func NewInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder) {
    m := &InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/informationProtection/sensitivityLabels/microsoft.graph.security.evaluateApplication", pathParameters),
    }
    return m
}
// NewInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder instantiates a new InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder and sets the default values.
func NewInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post compute the sensitivity label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set manually or explicitly by a user or service, rather than automatically based on file contents. Given contentInfo, which includes existing content metadata key-value pairs, and labelingOptions as an input, the API returns an informationProtectionAction object that contains one of more of the following: 
// Deprecated: This method is obsolete. Use {TypeName} instead.
// returns a InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-sensitivitylabel-evaluateapplication?view=graph-rest-1.0
func (m *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder) Post(ctx context.Context, body InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostRequestBodyable, requestConfiguration *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilderPostRequestConfiguration)(InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationResponseable), nil
}
// PostAsEvaluateApplicationPostResponse compute the sensitivity label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set manually or explicitly by a user or service, rather than automatically based on file contents. Given contentInfo, which includes existing content metadata key-value pairs, and labelingOptions as an input, the API returns an informationProtectionAction object that contains one of more of the following: 
// returns a InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-sensitivitylabel-evaluateapplication?view=graph-rest-1.0
func (m *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder) PostAsEvaluateApplicationPostResponse(ctx context.Context, body InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostRequestBodyable, requestConfiguration *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilderPostRequestConfiguration)(InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostResponseable), nil
}
// ToPostRequestInformation compute the sensitivity label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set manually or explicitly by a user or service, rather than automatically based on file contents. Given contentInfo, which includes existing content metadata key-value pairs, and labelingOptions as an input, the API returns an informationProtectionAction object that contains one of more of the following: 
// returns a *RequestInformation when successful
func (m *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder) ToPostRequestInformation(ctx context.Context, body InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationEvaluateApplicationPostRequestBodyable, requestConfiguration *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder when successful
func (m *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder) WithUrl(rawUrl string)(*InformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder) {
    return NewInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
