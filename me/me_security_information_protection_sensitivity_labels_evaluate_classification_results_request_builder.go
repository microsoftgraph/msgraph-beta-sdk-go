package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsRequestBuilder provides operations to call the evaluateClassificationResults method.
type MeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsRequestBuilderInternal instantiates a new EvaluateClassificationResultsRequestBuilder and sets the default values.
func NewMeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsRequestBuilder) {
    m := &MeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/security/informationProtection/sensitivityLabels/microsoft.graph.security.evaluateClassificationResults";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsRequestBuilder instantiates a new EvaluateClassificationResultsRequestBuilder and sets the default values.
func NewMeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action evaluateClassificationResults
func (m *MeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsPostRequestBodyable, requestConfiguration *MeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action evaluateClassificationResults
func (m *MeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsPostRequestBodyable, requestConfiguration *MeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponseable), nil
}
