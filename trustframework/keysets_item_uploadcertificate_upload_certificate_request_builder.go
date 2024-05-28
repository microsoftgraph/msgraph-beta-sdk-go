package trustframework

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// KeysetsItemUploadcertificateUploadCertificateRequestBuilder provides operations to call the uploadCertificate method.
type KeysetsItemUploadcertificateUploadCertificateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// KeysetsItemUploadcertificateUploadCertificateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type KeysetsItemUploadcertificateUploadCertificateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewKeysetsItemUploadcertificateUploadCertificateRequestBuilderInternal instantiates a new KeysetsItemUploadcertificateUploadCertificateRequestBuilder and sets the default values.
func NewKeysetsItemUploadcertificateUploadCertificateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*KeysetsItemUploadcertificateUploadCertificateRequestBuilder) {
    m := &KeysetsItemUploadcertificateUploadCertificateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/trustFramework/keySets/{trustFrameworkKeySet%2Did}/uploadCertificate", pathParameters),
    }
    return m
}
// NewKeysetsItemUploadcertificateUploadCertificateRequestBuilder instantiates a new KeysetsItemUploadcertificateUploadCertificateRequestBuilder and sets the default values.
func NewKeysetsItemUploadcertificateUploadCertificateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*KeysetsItemUploadcertificateUploadCertificateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewKeysetsItemUploadcertificateUploadCertificateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post upload a certificate to a trustFrameworkKeyset. The input is a base-64 encoded value of the certificate contents. This method returns trustFrameworkKey.
// returns a TrustFrameworkKeyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/trustframeworkkeyset-uploadcertificate?view=graph-rest-beta
func (m *KeysetsItemUploadcertificateUploadCertificateRequestBuilder) Post(ctx context.Context, body KeysetsItemUploadcertificateUploadCertificatePostRequestBodyable, requestConfiguration *KeysetsItemUploadcertificateUploadCertificateRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeyable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTrustFrameworkKeyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeyable), nil
}
// ToPostRequestInformation upload a certificate to a trustFrameworkKeyset. The input is a base-64 encoded value of the certificate contents. This method returns trustFrameworkKey.
// returns a *RequestInformation when successful
func (m *KeysetsItemUploadcertificateUploadCertificateRequestBuilder) ToPostRequestInformation(ctx context.Context, body KeysetsItemUploadcertificateUploadCertificatePostRequestBodyable, requestConfiguration *KeysetsItemUploadcertificateUploadCertificateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *KeysetsItemUploadcertificateUploadCertificateRequestBuilder when successful
func (m *KeysetsItemUploadcertificateUploadCertificateRequestBuilder) WithUrl(rawUrl string)(*KeysetsItemUploadcertificateUploadCertificateRequestBuilder) {
    return NewKeysetsItemUploadcertificateUploadCertificateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
