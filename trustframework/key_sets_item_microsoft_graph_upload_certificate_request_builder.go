package trustframework

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// KeySetsItemMicrosoftGraphUploadCertificateRequestBuilder provides operations to call the uploadCertificate method.
type KeySetsItemMicrosoftGraphUploadCertificateRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// KeySetsItemMicrosoftGraphUploadCertificateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type KeySetsItemMicrosoftGraphUploadCertificateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewKeySetsItemMicrosoftGraphUploadCertificateRequestBuilderInternal instantiates a new MicrosoftGraphUploadCertificateRequestBuilder and sets the default values.
func NewKeySetsItemMicrosoftGraphUploadCertificateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*KeySetsItemMicrosoftGraphUploadCertificateRequestBuilder) {
    m := &KeySetsItemMicrosoftGraphUploadCertificateRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/trustFramework/keySets/{trustFrameworkKeySet%2Did}/microsoft.graph.uploadCertificate";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewKeySetsItemMicrosoftGraphUploadCertificateRequestBuilder instantiates a new MicrosoftGraphUploadCertificateRequestBuilder and sets the default values.
func NewKeySetsItemMicrosoftGraphUploadCertificateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*KeySetsItemMicrosoftGraphUploadCertificateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewKeySetsItemMicrosoftGraphUploadCertificateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post upload a certificate to a trustFrameworkKeyset. The input is a base-64 encoded value of the certificate contents. This method returns trustFrameworkKey.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/trustframeworkkeyset-uploadcertificate?view=graph-rest-1.0
func (m *KeySetsItemMicrosoftGraphUploadCertificateRequestBuilder) Post(ctx context.Context, body KeySetsItemMicrosoftGraphUploadCertificateUploadCertificatePostRequestBodyable, requestConfiguration *KeySetsItemMicrosoftGraphUploadCertificateRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeyable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTrustFrameworkKeyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeyable), nil
}
// ToPostRequestInformation upload a certificate to a trustFrameworkKeyset. The input is a base-64 encoded value of the certificate contents. This method returns trustFrameworkKey.
func (m *KeySetsItemMicrosoftGraphUploadCertificateRequestBuilder) ToPostRequestInformation(ctx context.Context, body KeySetsItemMicrosoftGraphUploadCertificateUploadCertificatePostRequestBodyable, requestConfiguration *KeySetsItemMicrosoftGraphUploadCertificateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
