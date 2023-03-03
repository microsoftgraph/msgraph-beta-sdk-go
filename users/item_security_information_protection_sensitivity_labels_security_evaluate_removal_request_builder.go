package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilder provides operations to call the evaluateRemoval method.
type ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilderInternal instantiates a new SecurityEvaluateRemovalRequestBuilder and sets the default values.
func NewItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilder) {
    m := &ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/security/informationProtection/sensitivityLabels/security.evaluateRemoval";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilder instantiates a new SecurityEvaluateRemovalRequestBuilder and sets the default values.
func NewItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilderInternal(urlParams, requestAdapter)
}
// Post indicate to the consuming application what actions it should take to remove the label information. Given contentInfo as an input, which includes existing content metadata key-value pairs, the API returns an informationProtectionAction that contains some combination of one or more of the following: 
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/security-sensitivitylabel-evaluateremoval?view=graph-rest-1.0
func (m *ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilder) Post(ctx context.Context, body ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalEvaluateRemovalPostRequestBodyable, requestConfiguration *ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration)(ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalEvaluateRemovalResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalEvaluateRemovalResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalEvaluateRemovalResponseable), nil
}
// ToPostRequestInformation indicate to the consuming application what actions it should take to remove the label information. Given contentInfo as an input, which includes existing content metadata key-value pairs, the API returns an informationProtectionAction that contains some combination of one or more of the following: 
func (m *ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalEvaluateRemovalPostRequestBodyable, requestConfiguration *ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
