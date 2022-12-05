package trustframework

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilder provides operations to manage the keySets property of the microsoft.graph.trustFramework entity.
type TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilderGetQueryParameters get keySets from trustFramework
type TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilderGetQueryParameters
}
// TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilderInternal instantiates a new TrustFrameworkKeySetItemRequestBuilder and sets the default values.
func NewTrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilder) {
    m := &TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilder{
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
// NewTrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilder instantiates a new TrustFrameworkKeySetItemRequestBuilder and sets the default values.
func NewTrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property keySets for trustFramework
func (m *TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeySetable, requestConfiguration *TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// GenerateKey provides operations to call the generateKey method.
func (m *TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilder) GenerateKey()(*TrustFrameworkKeySetsItemGenerateKeyRequestBuilder) {
    return NewTrustFrameworkKeySetsItemGenerateKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get keySets from trustFramework
func (m *TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeySetable, error) {
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
func (m *TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilder) GetActiveKey()(*TrustFrameworkKeySetsItemGetActiveKeyRequestBuilder) {
    return NewTrustFrameworkKeySetsItemGetActiveKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property keySets in trustFramework
func (m *TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeySetable, requestConfiguration *TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeySetable, error) {
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
// UploadCertificate provides operations to call the uploadCertificate method.
func (m *TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilder) UploadCertificate()(*TrustFrameworkKeySetsItemUploadCertificateRequestBuilder) {
    return NewTrustFrameworkKeySetsItemUploadCertificateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UploadPkcs12 provides operations to call the uploadPkcs12 method.
func (m *TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilder) UploadPkcs12()(*TrustFrameworkKeySetsItemUploadPkcs12RequestBuilder) {
    return NewTrustFrameworkKeySetsItemUploadPkcs12RequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UploadSecret provides operations to call the uploadSecret method.
func (m *TrustFrameworkKeySetsTrustFrameworkKeySetItemRequestBuilder) UploadSecret()(*TrustFrameworkKeySetsItemUploadSecretRequestBuilder) {
    return NewTrustFrameworkKeySetsItemUploadSecretRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
