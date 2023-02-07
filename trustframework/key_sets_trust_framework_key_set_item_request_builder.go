package trustframework

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// KeySetsTrustFrameworkKeySetItemRequestBuilder provides operations to manage the keySets property of the microsoft.graph.trustFramework entity.
type KeySetsTrustFrameworkKeySetItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// KeySetsTrustFrameworkKeySetItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type KeySetsTrustFrameworkKeySetItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// KeySetsTrustFrameworkKeySetItemRequestBuilderGetQueryParameters get keySets from trustFramework
type KeySetsTrustFrameworkKeySetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// KeySetsTrustFrameworkKeySetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type KeySetsTrustFrameworkKeySetItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *KeySetsTrustFrameworkKeySetItemRequestBuilderGetQueryParameters
}
// KeySetsTrustFrameworkKeySetItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type KeySetsTrustFrameworkKeySetItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewKeySetsTrustFrameworkKeySetItemRequestBuilderInternal instantiates a new TrustFrameworkKeySetItemRequestBuilder and sets the default values.
func NewKeySetsTrustFrameworkKeySetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*KeySetsTrustFrameworkKeySetItemRequestBuilder) {
    m := &KeySetsTrustFrameworkKeySetItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/trustFramework/keySets/{trustFrameworkKeySet%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewKeySetsTrustFrameworkKeySetItemRequestBuilder instantiates a new TrustFrameworkKeySetItemRequestBuilder and sets the default values.
func NewKeySetsTrustFrameworkKeySetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*KeySetsTrustFrameworkKeySetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewKeySetsTrustFrameworkKeySetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property keySets for trustFramework
func (m *KeySetsTrustFrameworkKeySetItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *KeySetsTrustFrameworkKeySetItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get keySets from trustFramework
func (m *KeySetsTrustFrameworkKeySetItemRequestBuilder) Get(ctx context.Context, requestConfiguration *KeySetsTrustFrameworkKeySetItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeySetable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTrustFrameworkKeySetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeySetable), nil
}
// MicrosoftGraphGenerateKey provides operations to call the generateKey method.
func (m *KeySetsTrustFrameworkKeySetItemRequestBuilder) MicrosoftGraphGenerateKey()(*KeySetsItemMicrosoftGraphGenerateKeyRequestBuilder) {
    return NewKeySetsItemMicrosoftGraphGenerateKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetActiveKey provides operations to call the getActiveKey method.
func (m *KeySetsTrustFrameworkKeySetItemRequestBuilder) MicrosoftGraphGetActiveKey()(*KeySetsItemMicrosoftGraphGetActiveKeyRequestBuilder) {
    return NewKeySetsItemMicrosoftGraphGetActiveKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUploadCertificate provides operations to call the uploadCertificate method.
func (m *KeySetsTrustFrameworkKeySetItemRequestBuilder) MicrosoftGraphUploadCertificate()(*KeySetsItemMicrosoftGraphUploadCertificateRequestBuilder) {
    return NewKeySetsItemMicrosoftGraphUploadCertificateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUploadPkcs12 provides operations to call the uploadPkcs12 method.
func (m *KeySetsTrustFrameworkKeySetItemRequestBuilder) MicrosoftGraphUploadPkcs12()(*KeySetsItemMicrosoftGraphUploadPkcs12RequestBuilder) {
    return NewKeySetsItemMicrosoftGraphUploadPkcs12RequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUploadSecret provides operations to call the uploadSecret method.
func (m *KeySetsTrustFrameworkKeySetItemRequestBuilder) MicrosoftGraphUploadSecret()(*KeySetsItemMicrosoftGraphUploadSecretRequestBuilder) {
    return NewKeySetsItemMicrosoftGraphUploadSecretRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property keySets in trustFramework
func (m *KeySetsTrustFrameworkKeySetItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeySetable, requestConfiguration *KeySetsTrustFrameworkKeySetItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeySetable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTrustFrameworkKeySetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeySetable), nil
}
// ToDeleteRequestInformation delete navigation property keySets for trustFramework
func (m *KeySetsTrustFrameworkKeySetItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *KeySetsTrustFrameworkKeySetItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation get keySets from trustFramework
func (m *KeySetsTrustFrameworkKeySetItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *KeySetsTrustFrameworkKeySetItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property keySets in trustFramework
func (m *KeySetsTrustFrameworkKeySetItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeySetable, requestConfiguration *KeySetsTrustFrameworkKeySetItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
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
