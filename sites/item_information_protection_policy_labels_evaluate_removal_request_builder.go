package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder provides operations to call the evaluateRemoval method.
type ItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilderInternal instantiates a new EvaluateRemovalRequestBuilder and sets the default values.
func NewItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder) {
    m := &ItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/informationProtection/policy/labels/evaluateRemoval", pathParameters),
    }
    return m
}
// NewItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder instantiates a new EvaluateRemovalRequestBuilder and sets the default values.
func NewItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilderInternal(urlParams, requestAdapter)
}
// Post indicate to the consuming application what actions it should take to remove the label information. Given contentInfo as an input, which includes existing content metadata key/value pairs, the API returns an informationProtectionAction that contains some combination of one of more of the following: 
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels on 2021-02-15 and will be removed 2022-08-15
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/informationprotectionlabel-evaluateremoval?view=graph-rest-1.0
func (m *ItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder) Post(ctx context.Context, body ItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBodyable, requestConfiguration *ItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilderPostRequestConfiguration)(ItemInformationProtectionPolicyLabelsEvaluateRemovalResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemInformationProtectionPolicyLabelsEvaluateRemovalResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemInformationProtectionPolicyLabelsEvaluateRemovalResponseable), nil
}
// ToPostRequestInformation indicate to the consuming application what actions it should take to remove the label information. Given contentInfo as an input, which includes existing content metadata key/value pairs, the API returns an informationProtectionAction that contains some combination of one of more of the following: 
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels on 2021-02-15 and will be removed 2022-08-15
func (m *ItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBodyable, requestConfiguration *ItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels on 2021-02-15 and will be removed 2022-08-15
func (m *ItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder) WithUrl(rawUrl string)(*ItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder) {
    return NewItemInformationProtectionPolicyLabelsEvaluateRemovalRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
