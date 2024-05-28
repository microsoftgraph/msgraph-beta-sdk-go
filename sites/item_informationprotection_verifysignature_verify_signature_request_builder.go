package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemInformationprotectionVerifysignatureVerifySignatureRequestBuilder provides operations to call the verifySignature method.
type ItemInformationprotectionVerifysignatureVerifySignatureRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInformationprotectionVerifysignatureVerifySignatureRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInformationprotectionVerifysignatureVerifySignatureRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemInformationprotectionVerifysignatureVerifySignatureRequestBuilderInternal instantiates a new ItemInformationprotectionVerifysignatureVerifySignatureRequestBuilder and sets the default values.
func NewItemInformationprotectionVerifysignatureVerifySignatureRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInformationprotectionVerifysignatureVerifySignatureRequestBuilder) {
    m := &ItemInformationprotectionVerifysignatureVerifySignatureRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/informationProtection/verifySignature", pathParameters),
    }
    return m
}
// NewItemInformationprotectionVerifysignatureVerifySignatureRequestBuilder instantiates a new ItemInformationprotectionVerifysignatureVerifySignatureRequestBuilder and sets the default values.
func NewItemInformationprotectionVerifysignatureVerifySignatureRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInformationprotectionVerifysignatureVerifySignatureRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInformationprotectionVerifysignatureVerifySignatureRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action verifySignature
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels
// returns a VerificationResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemInformationprotectionVerifysignatureVerifySignatureRequestBuilder) Post(ctx context.Context, body ItemInformationprotectionVerifysignatureVerifySignaturePostRequestBodyable, requestConfiguration *ItemInformationprotectionVerifysignatureVerifySignatureRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VerificationResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateVerificationResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VerificationResultable), nil
}
// ToPostRequestInformation invoke action verifySignature
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels
// returns a *RequestInformation when successful
func (m *ItemInformationprotectionVerifysignatureVerifySignatureRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemInformationprotectionVerifysignatureVerifySignaturePostRequestBodyable, requestConfiguration *ItemInformationprotectionVerifysignatureVerifySignatureRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemInformationprotectionVerifysignatureVerifySignatureRequestBuilder when successful
func (m *ItemInformationprotectionVerifysignatureVerifySignatureRequestBuilder) WithUrl(rawUrl string)(*ItemInformationprotectionVerifysignatureVerifySignatureRequestBuilder) {
    return NewItemInformationprotectionVerifysignatureVerifySignatureRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
