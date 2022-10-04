package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i4df38ca1b862714120e25419c3ed380c7c772471f65cbf3e0d620c5a0d60f6bf "github.com/microsoftgraph/msgraph-beta-sdk-go/trustframework/keysets/item/uploadpkcs12"
    i78db104952796dbfc3b27052328de3db351b4c77d4b5dffcd1d52a18b26c3220 "github.com/microsoftgraph/msgraph-beta-sdk-go/trustframework/keysets/item/uploadsecret"
    i9d1297a38f5879454a214ecb884be348ff387c1b4b4287a28598568142a017af "github.com/microsoftgraph/msgraph-beta-sdk-go/trustframework/keysets/item/uploadcertificate"
    ib72c6292feaaf57eee55722f788b219f51c4f92a445271c20e5af998098f6daf "github.com/microsoftgraph/msgraph-beta-sdk-go/trustframework/keysets/item/getactivekey"
    ie04ccf263002e1d9853509d405e0878dfbf7c145b0a12ed976019d944eb6fe4d "github.com/microsoftgraph/msgraph-beta-sdk-go/trustframework/keysets/item/generatekey"
)

// TrustFrameworkKeySetItemRequestBuilder provides operations to manage the keySets property of the microsoft.graph.trustFramework entity.
type TrustFrameworkKeySetItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TrustFrameworkKeySetItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TrustFrameworkKeySetItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TrustFrameworkKeySetItemRequestBuilderGetQueryParameters get keySets from trustFramework
type TrustFrameworkKeySetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TrustFrameworkKeySetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TrustFrameworkKeySetItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TrustFrameworkKeySetItemRequestBuilderGetQueryParameters
}
// TrustFrameworkKeySetItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TrustFrameworkKeySetItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTrustFrameworkKeySetItemRequestBuilderInternal instantiates a new TrustFrameworkKeySetItemRequestBuilder and sets the default values.
func NewTrustFrameworkKeySetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TrustFrameworkKeySetItemRequestBuilder) {
    m := &TrustFrameworkKeySetItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/trustFramework/keySets/{trustFrameworkKeySet%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTrustFrameworkKeySetItemRequestBuilder instantiates a new TrustFrameworkKeySetItemRequestBuilder and sets the default values.
func NewTrustFrameworkKeySetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TrustFrameworkKeySetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTrustFrameworkKeySetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property keySets for trustFramework
func (m *TrustFrameworkKeySetItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *TrustFrameworkKeySetItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get keySets from trustFramework
func (m *TrustFrameworkKeySetItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TrustFrameworkKeySetItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property keySets in trustFramework
func (m *TrustFrameworkKeySetItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeySetable, requestConfiguration *TrustFrameworkKeySetItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property keySets for trustFramework
func (m *TrustFrameworkKeySetItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TrustFrameworkKeySetItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// GenerateKey the generateKey property
func (m *TrustFrameworkKeySetItemRequestBuilder) GenerateKey()(*ie04ccf263002e1d9853509d405e0878dfbf7c145b0a12ed976019d944eb6fe4d.GenerateKeyRequestBuilder) {
    return ie04ccf263002e1d9853509d405e0878dfbf7c145b0a12ed976019d944eb6fe4d.NewGenerateKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get keySets from trustFramework
func (m *TrustFrameworkKeySetItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TrustFrameworkKeySetItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeySetable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTrustFrameworkKeySetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeySetable), nil
}
// GetActiveKey provides operations to call the getActiveKey method.
func (m *TrustFrameworkKeySetItemRequestBuilder) GetActiveKey()(*ib72c6292feaaf57eee55722f788b219f51c4f92a445271c20e5af998098f6daf.GetActiveKeyRequestBuilder) {
    return ib72c6292feaaf57eee55722f788b219f51c4f92a445271c20e5af998098f6daf.NewGetActiveKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property keySets in trustFramework
func (m *TrustFrameworkKeySetItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeySetable, requestConfiguration *TrustFrameworkKeySetItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeySetable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTrustFrameworkKeySetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeySetable), nil
}
// UploadCertificate the uploadCertificate property
func (m *TrustFrameworkKeySetItemRequestBuilder) UploadCertificate()(*i9d1297a38f5879454a214ecb884be348ff387c1b4b4287a28598568142a017af.UploadCertificateRequestBuilder) {
    return i9d1297a38f5879454a214ecb884be348ff387c1b4b4287a28598568142a017af.NewUploadCertificateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UploadPkcs12 the uploadPkcs12 property
func (m *TrustFrameworkKeySetItemRequestBuilder) UploadPkcs12()(*i4df38ca1b862714120e25419c3ed380c7c772471f65cbf3e0d620c5a0d60f6bf.UploadPkcs12RequestBuilder) {
    return i4df38ca1b862714120e25419c3ed380c7c772471f65cbf3e0d620c5a0d60f6bf.NewUploadPkcs12RequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UploadSecret the uploadSecret property
func (m *TrustFrameworkKeySetItemRequestBuilder) UploadSecret()(*i78db104952796dbfc3b27052328de3db351b4c77d4b5dffcd1d52a18b26c3220.UploadSecretRequestBuilder) {
    return i78db104952796dbfc3b27052328de3db351b4c77d4b5dffcd1d52a18b26c3220.NewUploadSecretRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
