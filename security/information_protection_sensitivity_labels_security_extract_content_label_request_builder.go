package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// InformationProtectionSensitivityLabelsSecurityExtractContentLabelRequestBuilder provides operations to call the extractContentLabel method.
type InformationProtectionSensitivityLabelsSecurityExtractContentLabelRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// InformationProtectionSensitivityLabelsSecurityExtractContentLabelRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InformationProtectionSensitivityLabelsSecurityExtractContentLabelRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewInformationProtectionSensitivityLabelsSecurityExtractContentLabelRequestBuilderInternal instantiates a new SecurityExtractContentLabelRequestBuilder and sets the default values.
func NewInformationProtectionSensitivityLabelsSecurityExtractContentLabelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InformationProtectionSensitivityLabelsSecurityExtractContentLabelRequestBuilder) {
    m := &InformationProtectionSensitivityLabelsSecurityExtractContentLabelRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/informationProtection/sensitivityLabels/security.extractContentLabel";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewInformationProtectionSensitivityLabelsSecurityExtractContentLabelRequestBuilder instantiates a new SecurityExtractContentLabelRequestBuilder and sets the default values.
func NewInformationProtectionSensitivityLabelsSecurityExtractContentLabelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InformationProtectionSensitivityLabelsSecurityExtractContentLabelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInformationProtectionSensitivityLabelsSecurityExtractContentLabelRequestBuilderInternal(urlParams, requestAdapter)
}
// Post use the metadata that exists on an already-labeled piece of information to resolve the metadata to a specific sensitivity label. The contentInfo input is resolved to informationProtectionContentLabel.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/security-sensitivitylabel-extractcontentlabel?view=graph-rest-1.0
func (m *InformationProtectionSensitivityLabelsSecurityExtractContentLabelRequestBuilder) Post(ctx context.Context, body InformationProtectionSensitivityLabelsSecurityExtractContentLabelExtractContentLabelPostRequestBodyable, requestConfiguration *InformationProtectionSensitivityLabelsSecurityExtractContentLabelRequestBuilderPostRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ContentLabelable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateContentLabelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ContentLabelable), nil
}
// ToPostRequestInformation use the metadata that exists on an already-labeled piece of information to resolve the metadata to a specific sensitivity label. The contentInfo input is resolved to informationProtectionContentLabel.
func (m *InformationProtectionSensitivityLabelsSecurityExtractContentLabelRequestBuilder) ToPostRequestInformation(ctx context.Context, body InformationProtectionSensitivityLabelsSecurityExtractContentLabelExtractContentLabelPostRequestBodyable, requestConfiguration *InformationProtectionSensitivityLabelsSecurityExtractContentLabelRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
