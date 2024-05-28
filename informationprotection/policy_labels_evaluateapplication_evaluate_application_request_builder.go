package informationprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PolicyLabelsEvaluateapplicationEvaluateApplicationRequestBuilder provides operations to call the evaluateApplication method.
type PolicyLabelsEvaluateapplicationEvaluateApplicationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PolicyLabelsEvaluateapplicationEvaluateApplicationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PolicyLabelsEvaluateapplicationEvaluateApplicationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPolicyLabelsEvaluateapplicationEvaluateApplicationRequestBuilderInternal instantiates a new PolicyLabelsEvaluateapplicationEvaluateApplicationRequestBuilder and sets the default values.
func NewPolicyLabelsEvaluateapplicationEvaluateApplicationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PolicyLabelsEvaluateapplicationEvaluateApplicationRequestBuilder) {
    m := &PolicyLabelsEvaluateapplicationEvaluateApplicationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/informationProtection/policy/labels/evaluateApplication", pathParameters),
    }
    return m
}
// NewPolicyLabelsEvaluateapplicationEvaluateApplicationRequestBuilder instantiates a new PolicyLabelsEvaluateapplicationEvaluateApplicationRequestBuilder and sets the default values.
func NewPolicyLabelsEvaluateapplicationEvaluateApplicationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PolicyLabelsEvaluateapplicationEvaluateApplicationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPolicyLabelsEvaluateapplicationEvaluateApplicationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post compute the information protection label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set manually or explicitly by a user or service, rather than automatically based on file contents.  Given contentInfo, which includes existing content metadata key/value pairs, and labelingOptions as an input, the API returns an informationProtectionAction object that contains one of more of the following: 
// Deprecated: This method is obsolete. Use PostAsEvaluateApplicationPostResponse instead.
// returns a PolicyLabelsEvaluateapplicationEvaluateApplicationResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/informationprotectionlabel-evaluateapplication?view=graph-rest-beta
func (m *PolicyLabelsEvaluateapplicationEvaluateApplicationRequestBuilder) Post(ctx context.Context, body PolicyLabelsEvaluateapplicationEvaluateApplicationPostRequestBodyable, requestConfiguration *PolicyLabelsEvaluateapplicationEvaluateApplicationRequestBuilderPostRequestConfiguration)(PolicyLabelsEvaluateapplicationEvaluateApplicationResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreatePolicyLabelsEvaluateapplicationEvaluateApplicationResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(PolicyLabelsEvaluateapplicationEvaluateApplicationResponseable), nil
}
// PostAsEvaluateApplicationPostResponse compute the information protection label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set manually or explicitly by a user or service, rather than automatically based on file contents.  Given contentInfo, which includes existing content metadata key/value pairs, and labelingOptions as an input, the API returns an informationProtectionAction object that contains one of more of the following: 
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels
// returns a PolicyLabelsEvaluateapplicationEvaluateApplicationPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/informationprotectionlabel-evaluateapplication?view=graph-rest-beta
func (m *PolicyLabelsEvaluateapplicationEvaluateApplicationRequestBuilder) PostAsEvaluateApplicationPostResponse(ctx context.Context, body PolicyLabelsEvaluateapplicationEvaluateApplicationPostRequestBodyable, requestConfiguration *PolicyLabelsEvaluateapplicationEvaluateApplicationRequestBuilderPostRequestConfiguration)(PolicyLabelsEvaluateapplicationEvaluateApplicationPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreatePolicyLabelsEvaluateapplicationEvaluateApplicationPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(PolicyLabelsEvaluateapplicationEvaluateApplicationPostResponseable), nil
}
// ToPostRequestInformation compute the information protection label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set manually or explicitly by a user or service, rather than automatically based on file contents.  Given contentInfo, which includes existing content metadata key/value pairs, and labelingOptions as an input, the API returns an informationProtectionAction object that contains one of more of the following: 
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels
// returns a *RequestInformation when successful
func (m *PolicyLabelsEvaluateapplicationEvaluateApplicationRequestBuilder) ToPostRequestInformation(ctx context.Context, body PolicyLabelsEvaluateapplicationEvaluateApplicationPostRequestBodyable, requestConfiguration *PolicyLabelsEvaluateapplicationEvaluateApplicationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PolicyLabelsEvaluateapplicationEvaluateApplicationRequestBuilder when successful
func (m *PolicyLabelsEvaluateapplicationEvaluateApplicationRequestBuilder) WithUrl(rawUrl string)(*PolicyLabelsEvaluateapplicationEvaluateApplicationRequestBuilder) {
    return NewPolicyLabelsEvaluateapplicationEvaluateApplicationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
