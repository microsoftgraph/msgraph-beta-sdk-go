package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemInformationProtectionDataLossPreventionPoliciesEvaluateRequestBuilder provides operations to call the evaluate method.
type ItemInformationProtectionDataLossPreventionPoliciesEvaluateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInformationProtectionDataLossPreventionPoliciesEvaluateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInformationProtectionDataLossPreventionPoliciesEvaluateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemInformationProtectionDataLossPreventionPoliciesEvaluateRequestBuilderInternal instantiates a new EvaluateRequestBuilder and sets the default values.
func NewItemInformationProtectionDataLossPreventionPoliciesEvaluateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInformationProtectionDataLossPreventionPoliciesEvaluateRequestBuilder) {
    m := &ItemInformationProtectionDataLossPreventionPoliciesEvaluateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/informationProtection/dataLossPreventionPolicies/evaluate", pathParameters),
    }
    return m
}
// NewItemInformationProtectionDataLossPreventionPoliciesEvaluateRequestBuilder instantiates a new EvaluateRequestBuilder and sets the default values.
func NewItemInformationProtectionDataLossPreventionPoliciesEvaluateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInformationProtectionDataLossPreventionPoliciesEvaluateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInformationProtectionDataLossPreventionPoliciesEvaluateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action evaluate
func (m *ItemInformationProtectionDataLossPreventionPoliciesEvaluateRequestBuilder) Post(ctx context.Context, body ItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBodyable, requestConfiguration *ItemInformationProtectionDataLossPreventionPoliciesEvaluateRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DlpEvaluatePoliciesJobResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDlpEvaluatePoliciesJobResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DlpEvaluatePoliciesJobResponseable), nil
}
// ToPostRequestInformation invoke action evaluate
func (m *ItemInformationProtectionDataLossPreventionPoliciesEvaluateRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBodyable, requestConfiguration *ItemInformationProtectionDataLossPreventionPoliciesEvaluateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemInformationProtectionDataLossPreventionPoliciesEvaluateRequestBuilder) WithUrl(rawUrl string)(*ItemInformationProtectionDataLossPreventionPoliciesEvaluateRequestBuilder) {
    return NewItemInformationProtectionDataLossPreventionPoliciesEvaluateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
