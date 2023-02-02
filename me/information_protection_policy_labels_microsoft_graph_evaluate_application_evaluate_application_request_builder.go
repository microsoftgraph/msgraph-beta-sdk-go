package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// InformationProtectionPolicyLabelsMicrosoftGraphEvaluateApplicationEvaluateApplicationRequestBuilder provides operations to call the evaluateApplication method.
type InformationProtectionPolicyLabelsMicrosoftGraphEvaluateApplicationEvaluateApplicationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// InformationProtectionPolicyLabelsMicrosoftGraphEvaluateApplicationEvaluateApplicationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InformationProtectionPolicyLabelsMicrosoftGraphEvaluateApplicationEvaluateApplicationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewInformationProtectionPolicyLabelsMicrosoftGraphEvaluateApplicationEvaluateApplicationRequestBuilderInternal instantiates a new EvaluateApplicationRequestBuilder and sets the default values.
func NewInformationProtectionPolicyLabelsMicrosoftGraphEvaluateApplicationEvaluateApplicationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InformationProtectionPolicyLabelsMicrosoftGraphEvaluateApplicationEvaluateApplicationRequestBuilder) {
    m := &InformationProtectionPolicyLabelsMicrosoftGraphEvaluateApplicationEvaluateApplicationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/informationProtection/policy/labels/microsoft.graph.evaluateApplication";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewInformationProtectionPolicyLabelsMicrosoftGraphEvaluateApplicationEvaluateApplicationRequestBuilder instantiates a new EvaluateApplicationRequestBuilder and sets the default values.
func NewInformationProtectionPolicyLabelsMicrosoftGraphEvaluateApplicationEvaluateApplicationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InformationProtectionPolicyLabelsMicrosoftGraphEvaluateApplicationEvaluateApplicationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInformationProtectionPolicyLabelsMicrosoftGraphEvaluateApplicationEvaluateApplicationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post compute the information protection label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set manually or explicitly by a user or service, rather than automatically based on file contents.  Given contentInfo, which includes existing content metadata key/value pairs, and labelingOptions as an input, the API returns an informationProtectionAction object that contains one of more of the following: 
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/informationprotectionlabel-evaluateapplication?view=graph-rest-1.0
func (m *InformationProtectionPolicyLabelsMicrosoftGraphEvaluateApplicationEvaluateApplicationRequestBuilder) Post(ctx context.Context, body InformationProtectionPolicyLabelsMicrosoftGraphEvaluateApplicationEvaluateApplicationPostRequestBodyable, requestConfiguration *InformationProtectionPolicyLabelsMicrosoftGraphEvaluateApplicationEvaluateApplicationRequestBuilderPostRequestConfiguration)(InformationProtectionPolicyLabelsMicrosoftGraphEvaluateApplicationEvaluateApplicationResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateInformationProtectionPolicyLabelsMicrosoftGraphEvaluateApplicationEvaluateApplicationResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InformationProtectionPolicyLabelsMicrosoftGraphEvaluateApplicationEvaluateApplicationResponseable), nil
}
// ToPostRequestInformation compute the information protection label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set manually or explicitly by a user or service, rather than automatically based on file contents.  Given contentInfo, which includes existing content metadata key/value pairs, and labelingOptions as an input, the API returns an informationProtectionAction object that contains one of more of the following: 
func (m *InformationProtectionPolicyLabelsMicrosoftGraphEvaluateApplicationEvaluateApplicationRequestBuilder) ToPostRequestInformation(ctx context.Context, body InformationProtectionPolicyLabelsMicrosoftGraphEvaluateApplicationEvaluateApplicationPostRequestBodyable, requestConfiguration *InformationProtectionPolicyLabelsMicrosoftGraphEvaluateApplicationEvaluateApplicationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
