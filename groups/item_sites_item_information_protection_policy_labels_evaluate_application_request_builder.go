package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilder provides operations to call the evaluateApplication method.
type ItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilderInternal instantiates a new EvaluateApplicationRequestBuilder and sets the default values.
func NewItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilder) {
    m := &ItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/informationProtection/policy/labels/evaluateApplication", pathParameters),
    }
    return m
}
// NewItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilder instantiates a new EvaluateApplicationRequestBuilder and sets the default values.
func NewItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post compute the information protection label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set manually or explicitly by a user or service, rather than automatically based on file contents.  Given contentInfo, which includes existing content metadata key/value pairs, and labelingOptions as an input, the API returns an informationProtectionAction object that contains one of more of the following:  This API is available in the following national cloud deployments.
// Deprecated: This method is obsolete. Use PostAsEvaluateApplicationPostResponse instead.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/informationprotectionlabel-evaluateapplication?view=graph-rest-1.0
func (m *ItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilder) Post(ctx context.Context, body ItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBodyable, requestConfiguration *ItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilderPostRequestConfiguration)(ItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationResponseable), nil
}
// PostAsEvaluateApplicationPostResponse compute the information protection label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set manually or explicitly by a user or service, rather than automatically based on file contents.  Given contentInfo, which includes existing content metadata key/value pairs, and labelingOptions as an input, the API returns an informationProtectionAction object that contains one of more of the following:  This API is available in the following national cloud deployments.
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels on 2021-02-15 and will be removed 2022-08-15
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/informationprotectionlabel-evaluateapplication?view=graph-rest-1.0
func (m *ItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilder) PostAsEvaluateApplicationPostResponse(ctx context.Context, body ItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBodyable, requestConfiguration *ItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilderPostRequestConfiguration)(ItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationPostResponseable), nil
}
// ToPostRequestInformation compute the information protection label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set manually or explicitly by a user or service, rather than automatically based on file contents.  Given contentInfo, which includes existing content metadata key/value pairs, and labelingOptions as an input, the API returns an informationProtectionAction object that contains one of more of the following:  This API is available in the following national cloud deployments.
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels on 2021-02-15 and will be removed 2022-08-15
func (m *ItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBodyable, requestConfiguration *ItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilder) {
    return NewItemSitesItemInformationProtectionPolicyLabelsEvaluateApplicationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
