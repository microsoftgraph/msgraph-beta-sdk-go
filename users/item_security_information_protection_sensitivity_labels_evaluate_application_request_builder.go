package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationRequestBuilder provides operations to call the evaluateApplication method.
type ItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationRequestBuilderInternal instantiates a new EvaluateApplicationRequestBuilder and sets the default values.
func NewItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationRequestBuilder) {
    m := &ItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/security/informationProtection/sensitivityLabels/microsoft.graph.security.evaluateApplication";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationRequestBuilder instantiates a new EvaluateApplicationRequestBuilder and sets the default values.
func NewItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post compute the sensitivity label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set manually or explicitly by a user or service, rather than automatically based on file contents. Given contentInfo, which includes existing content metadata key-value pairs, and labelingOptions as an input, the API returns an informationProtectionAction object that contains one of more of the following: 
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/security-sensitivitylabel-evaluateapplication?view=graph-rest-1.0
func (m *ItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationRequestBuilder) Post(ctx context.Context, body ItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationPostRequestBodyable, requestConfiguration *ItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationRequestBuilderPostRequestConfiguration)(ItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, CreateItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationResponseable), nil
}
// ToPostRequestInformation compute the sensitivity label that should be applied and return the set of actions that must be taken to correctly label the information. This API is useful when a label should be set manually or explicitly by a user or service, rather than automatically based on file contents. Given contentInfo, which includes existing content metadata key-value pairs, and labelingOptions as an input, the API returns an informationProtectionAction object that contains one of more of the following: 
func (m *ItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationPostRequestBodyable, requestConfiguration *ItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
