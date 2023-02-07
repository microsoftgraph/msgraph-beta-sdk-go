package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemInformationProtectionPolicyLabelsMicrosoftGraphEvaluateRemovalRequestBuilder provides operations to call the evaluateRemoval method.
type ItemInformationProtectionPolicyLabelsMicrosoftGraphEvaluateRemovalRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemInformationProtectionPolicyLabelsMicrosoftGraphEvaluateRemovalRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInformationProtectionPolicyLabelsMicrosoftGraphEvaluateRemovalRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemInformationProtectionPolicyLabelsMicrosoftGraphEvaluateRemovalRequestBuilderInternal instantiates a new MicrosoftGraphEvaluateRemovalRequestBuilder and sets the default values.
func NewItemInformationProtectionPolicyLabelsMicrosoftGraphEvaluateRemovalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInformationProtectionPolicyLabelsMicrosoftGraphEvaluateRemovalRequestBuilder) {
    m := &ItemInformationProtectionPolicyLabelsMicrosoftGraphEvaluateRemovalRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/informationProtection/policy/labels/microsoft.graph.evaluateRemoval";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemInformationProtectionPolicyLabelsMicrosoftGraphEvaluateRemovalRequestBuilder instantiates a new MicrosoftGraphEvaluateRemovalRequestBuilder and sets the default values.
func NewItemInformationProtectionPolicyLabelsMicrosoftGraphEvaluateRemovalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInformationProtectionPolicyLabelsMicrosoftGraphEvaluateRemovalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInformationProtectionPolicyLabelsMicrosoftGraphEvaluateRemovalRequestBuilderInternal(urlParams, requestAdapter)
}
// Post indicate to the consuming application what actions it should take to remove the label information. Given contentInfo as an input, which includes existing content metadata key/value pairs, the API returns an informationProtectionAction that contains some combination of one of more of the following: 
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/informationprotectionlabel-evaluateremoval?view=graph-rest-1.0
func (m *ItemInformationProtectionPolicyLabelsMicrosoftGraphEvaluateRemovalRequestBuilder) Post(ctx context.Context, body ItemInformationProtectionPolicyLabelsMicrosoftGraphEvaluateRemovalEvaluateRemovalPostRequestBodyable, requestConfiguration *ItemInformationProtectionPolicyLabelsMicrosoftGraphEvaluateRemovalRequestBuilderPostRequestConfiguration)(ItemInformationProtectionPolicyLabelsMicrosoftGraphEvaluateRemovalEvaluateRemovalResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemInformationProtectionPolicyLabelsMicrosoftGraphEvaluateRemovalEvaluateRemovalResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemInformationProtectionPolicyLabelsMicrosoftGraphEvaluateRemovalEvaluateRemovalResponseable), nil
}
// ToPostRequestInformation indicate to the consuming application what actions it should take to remove the label information. Given contentInfo as an input, which includes existing content metadata key/value pairs, the API returns an informationProtectionAction that contains some combination of one of more of the following: 
func (m *ItemInformationProtectionPolicyLabelsMicrosoftGraphEvaluateRemovalRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemInformationProtectionPolicyLabelsMicrosoftGraphEvaluateRemovalEvaluateRemovalPostRequestBodyable, requestConfiguration *ItemInformationProtectionPolicyLabelsMicrosoftGraphEvaluateRemovalRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
