package trustframework

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// KeysetsItemUploadpkcs12UploadPkcs12RequestBuilder provides operations to call the uploadPkcs12 method.
type KeysetsItemUploadpkcs12UploadPkcs12RequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// KeysetsItemUploadpkcs12UploadPkcs12RequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type KeysetsItemUploadpkcs12UploadPkcs12RequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewKeysetsItemUploadpkcs12UploadPkcs12RequestBuilderInternal instantiates a new KeysetsItemUploadpkcs12UploadPkcs12RequestBuilder and sets the default values.
func NewKeysetsItemUploadpkcs12UploadPkcs12RequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*KeysetsItemUploadpkcs12UploadPkcs12RequestBuilder) {
    m := &KeysetsItemUploadpkcs12UploadPkcs12RequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/trustFramework/keySets/{trustFrameworkKeySet%2Did}/uploadPkcs12", pathParameters),
    }
    return m
}
// NewKeysetsItemUploadpkcs12UploadPkcs12RequestBuilder instantiates a new KeysetsItemUploadpkcs12UploadPkcs12RequestBuilder and sets the default values.
func NewKeysetsItemUploadpkcs12UploadPkcs12RequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*KeysetsItemUploadpkcs12UploadPkcs12RequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewKeysetsItemUploadpkcs12UploadPkcs12RequestBuilderInternal(urlParams, requestAdapter)
}
// Post upload a PKCS12 format key (PFX) to a trustFrameworkKeyset. The input is a base-64 encoded value of the Pfx certificate contents. This method returns trustFrameworkKey.
// returns a TrustFrameworkKeyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/trustframeworkkeyset-uploadpkcs12?view=graph-rest-beta
func (m *KeysetsItemUploadpkcs12UploadPkcs12RequestBuilder) Post(ctx context.Context, body KeysetsItemUploadpkcs12UploadPkcs12PostRequestBodyable, requestConfiguration *KeysetsItemUploadpkcs12UploadPkcs12RequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeyable, error) {
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
// ToPostRequestInformation upload a PKCS12 format key (PFX) to a trustFrameworkKeyset. The input is a base-64 encoded value of the Pfx certificate contents. This method returns trustFrameworkKey.
// returns a *RequestInformation when successful
func (m *KeysetsItemUploadpkcs12UploadPkcs12RequestBuilder) ToPostRequestInformation(ctx context.Context, body KeysetsItemUploadpkcs12UploadPkcs12PostRequestBodyable, requestConfiguration *KeysetsItemUploadpkcs12UploadPkcs12RequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *KeysetsItemUploadpkcs12UploadPkcs12RequestBuilder when successful
func (m *KeysetsItemUploadpkcs12UploadPkcs12RequestBuilder) WithUrl(rawUrl string)(*KeysetsItemUploadpkcs12UploadPkcs12RequestBuilder) {
    return NewKeysetsItemUploadpkcs12UploadPkcs12RequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
