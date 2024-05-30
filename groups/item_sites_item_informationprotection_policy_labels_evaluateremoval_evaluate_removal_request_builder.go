package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder provides operations to call the evaluateRemoval method.
type ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilderInternal instantiates a new ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder and sets the default values.
func NewItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder) {
    m := &ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/informationProtection/policy/labels/evaluateRemoval", pathParameters),
    }
    return m
}
// NewItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder instantiates a new ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder and sets the default values.
func NewItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilderInternal(urlParams, requestAdapter)
}
// Post indicate to the consuming application what actions it should take to remove the label information. Given contentInfo as an input, which includes existing content metadata key/value pairs, the API returns an informationProtectionAction that contains some combination of one of more of the following: 
// Deprecated: This method is obsolete. Use PostAsEvaluateRemovalPostResponse instead.
// returns a ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/informationprotectionlabel-evaluateremoval?view=graph-rest-beta
func (m *ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder) Post(ctx context.Context, body ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalPostRequestBodyable, requestConfiguration *ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilderPostRequestConfiguration)(ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalResponseable), nil
}
// PostAsEvaluateRemovalPostResponse indicate to the consuming application what actions it should take to remove the label information. Given contentInfo as an input, which includes existing content metadata key/value pairs, the API returns an informationProtectionAction that contains some combination of one of more of the following: 
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels
// returns a ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/informationprotectionlabel-evaluateremoval?view=graph-rest-beta
func (m *ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder) PostAsEvaluateRemovalPostResponse(ctx context.Context, body ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalPostRequestBodyable, requestConfiguration *ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilderPostRequestConfiguration)(ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalPostResponseable), nil
}
// ToPostRequestInformation indicate to the consuming application what actions it should take to remove the label information. Given contentInfo as an input, which includes existing content metadata key/value pairs, the API returns an informationProtectionAction that contains some combination of one of more of the following: 
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels
// returns a *RequestInformation when successful
func (m *ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalPostRequestBodyable, requestConfiguration *ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder when successful
func (m *ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder) {
    return NewItemSitesItemInformationprotectionPolicyLabelsEvaluateremovalEvaluateRemovalRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
