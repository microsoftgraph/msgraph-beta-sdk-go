package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// InformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelRequestBuilder provides operations to call the extractContentLabel method.
type InformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// InformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewInformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelRequestBuilderInternal instantiates a new MicrosoftGraphSecurityExtractContentLabelRequestBuilder and sets the default values.
func NewInformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelRequestBuilder) {
    m := &InformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/informationProtection/sensitivityLabels/microsoft.graph.security.extractContentLabel", pathParameters),
    }
    return m
}
// NewInformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelRequestBuilder instantiates a new MicrosoftGraphSecurityExtractContentLabelRequestBuilder and sets the default values.
func NewInformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelRequestBuilderInternal(urlParams, requestAdapter)
}
// Post use the metadata that exists on an already-labeled piece of information to resolve the metadata to a specific sensitivity label. The contentInfo input is resolved to informationProtectionContentLabel. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-sensitivitylabel-extractcontentlabel?view=graph-rest-1.0
func (m *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelRequestBuilder) Post(ctx context.Context, body InformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelExtractContentLabelPostRequestBodyable, requestConfiguration *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelRequestBuilderPostRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ContentLabelable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateContentLabelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ContentLabelable), nil
}
// ToPostRequestInformation use the metadata that exists on an already-labeled piece of information to resolve the metadata to a specific sensitivity label. The contentInfo input is resolved to informationProtectionContentLabel. This API is available in the following national cloud deployments.
func (m *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelRequestBuilder) ToPostRequestInformation(ctx context.Context, body InformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelExtractContentLabelPostRequestBodyable, requestConfiguration *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *InformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelRequestBuilder) WithUrl(rawUrl string)(*InformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelRequestBuilder) {
    return NewInformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
