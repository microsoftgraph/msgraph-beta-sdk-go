package trustframework

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// KeysetsTrustFrameworkKeySetItemRequestBuilder provides operations to manage the keySets property of the microsoft.graph.trustFramework entity.
type KeysetsTrustFrameworkKeySetItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// KeysetsTrustFrameworkKeySetItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type KeysetsTrustFrameworkKeySetItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// KeysetsTrustFrameworkKeySetItemRequestBuilderGetQueryParameters retrieve the properties and associations for a Trustframeworkkeyset.
type KeysetsTrustFrameworkKeySetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// KeysetsTrustFrameworkKeySetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type KeysetsTrustFrameworkKeySetItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *KeysetsTrustFrameworkKeySetItemRequestBuilderGetQueryParameters
}
// KeysetsTrustFrameworkKeySetItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type KeysetsTrustFrameworkKeySetItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewKeysetsTrustFrameworkKeySetItemRequestBuilderInternal instantiates a new KeysetsTrustFrameworkKeySetItemRequestBuilder and sets the default values.
func NewKeysetsTrustFrameworkKeySetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*KeysetsTrustFrameworkKeySetItemRequestBuilder) {
    m := &KeysetsTrustFrameworkKeySetItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/trustFramework/keySets/{trustFrameworkKeySet%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewKeysetsTrustFrameworkKeySetItemRequestBuilder instantiates a new KeysetsTrustFrameworkKeySetItemRequestBuilder and sets the default values.
func NewKeysetsTrustFrameworkKeySetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*KeysetsTrustFrameworkKeySetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewKeysetsTrustFrameworkKeySetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a trustFrameworkKeySet.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/trustframeworkkeyset-delete?view=graph-rest-beta
func (m *KeysetsTrustFrameworkKeySetItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *KeysetsTrustFrameworkKeySetItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// GenerateKey provides operations to call the generateKey method.
// returns a *KeysetsItemGeneratekeyGenerateKeyRequestBuilder when successful
func (m *KeysetsTrustFrameworkKeySetItemRequestBuilder) GenerateKey()(*KeysetsItemGeneratekeyGenerateKeyRequestBuilder) {
    return NewKeysetsItemGeneratekeyGenerateKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve the properties and associations for a Trustframeworkkeyset.
// returns a TrustFrameworkKeySetable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/trustframeworkkeyset-get?view=graph-rest-beta
func (m *KeysetsTrustFrameworkKeySetItemRequestBuilder) Get(ctx context.Context, requestConfiguration *KeysetsTrustFrameworkKeySetItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeySetable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTrustFrameworkKeySetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeySetable), nil
}
// GetActiveKey provides operations to call the getActiveKey method.
// returns a *KeysetsItemGetactivekeyGetActiveKeyRequestBuilder when successful
func (m *KeysetsTrustFrameworkKeySetItemRequestBuilder) GetActiveKey()(*KeysetsItemGetactivekeyGetActiveKeyRequestBuilder) {
    return NewKeysetsItemGetactivekeyGetActiveKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Keys_v2 provides operations to manage the keys_v2 property of the microsoft.graph.trustFrameworkKeySet entity.
// returns a *KeysetsItemKeys_v2RequestBuilder when successful
func (m *KeysetsTrustFrameworkKeySetItemRequestBuilder) Keys_v2()(*KeysetsItemKeys_v2RequestBuilder) {
    return NewKeysetsItemKeys_v2RequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of a trustFrameworkKeyset. This operation will replace the content of an existing keyset. Specifying the ID in the request payload is optional.
// returns a TrustFrameworkKeySetable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/trustframeworkkeyset-update?view=graph-rest-beta
func (m *KeysetsTrustFrameworkKeySetItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeySetable, requestConfiguration *KeysetsTrustFrameworkKeySetItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeySetable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTrustFrameworkKeySetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeySetable), nil
}
// ToDeleteRequestInformation delete a trustFrameworkKeySet.
// returns a *RequestInformation when successful
func (m *KeysetsTrustFrameworkKeySetItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *KeysetsTrustFrameworkKeySetItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the properties and associations for a Trustframeworkkeyset.
// returns a *RequestInformation when successful
func (m *KeysetsTrustFrameworkKeySetItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *KeysetsTrustFrameworkKeySetItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the properties of a trustFrameworkKeyset. This operation will replace the content of an existing keyset. Specifying the ID in the request payload is optional.
// returns a *RequestInformation when successful
func (m *KeysetsTrustFrameworkKeySetItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeySetable, requestConfiguration *KeysetsTrustFrameworkKeySetItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// UploadCertificate provides operations to call the uploadCertificate method.
// returns a *KeysetsItemUploadcertificateUploadCertificateRequestBuilder when successful
func (m *KeysetsTrustFrameworkKeySetItemRequestBuilder) UploadCertificate()(*KeysetsItemUploadcertificateUploadCertificateRequestBuilder) {
    return NewKeysetsItemUploadcertificateUploadCertificateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UploadPkcs12 provides operations to call the uploadPkcs12 method.
// returns a *KeysetsItemUploadpkcs12UploadPkcs12RequestBuilder when successful
func (m *KeysetsTrustFrameworkKeySetItemRequestBuilder) UploadPkcs12()(*KeysetsItemUploadpkcs12UploadPkcs12RequestBuilder) {
    return NewKeysetsItemUploadpkcs12UploadPkcs12RequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UploadSecret provides operations to call the uploadSecret method.
// returns a *KeysetsItemUploadsecretUploadSecretRequestBuilder when successful
func (m *KeysetsTrustFrameworkKeySetItemRequestBuilder) UploadSecret()(*KeysetsItemUploadsecretUploadSecretRequestBuilder) {
    return NewKeysetsItemUploadsecretUploadSecretRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *KeysetsTrustFrameworkKeySetItemRequestBuilder when successful
func (m *KeysetsTrustFrameworkKeySetItemRequestBuilder) WithUrl(rawUrl string)(*KeysetsTrustFrameworkKeySetItemRequestBuilder) {
    return NewKeysetsTrustFrameworkKeySetItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
