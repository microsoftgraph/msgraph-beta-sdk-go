package item

import (
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
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TrustFrameworkKeySetItemRequestBuilderDeleteOptions options for Delete
type TrustFrameworkKeySetItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// TrustFrameworkKeySetItemRequestBuilderGetOptions options for Get
type TrustFrameworkKeySetItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *TrustFrameworkKeySetItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// TrustFrameworkKeySetItemRequestBuilderGetQueryParameters get keySets from trustFramework
type TrustFrameworkKeySetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// TrustFrameworkKeySetItemRequestBuilderPatchOptions options for Patch
type TrustFrameworkKeySetItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeySetable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewTrustFrameworkKeySetItemRequestBuilderInternal instantiates a new TrustFrameworkKeySetItemRequestBuilder and sets the default values.
func NewTrustFrameworkKeySetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TrustFrameworkKeySetItemRequestBuilder) {
    m := &TrustFrameworkKeySetItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/trustFramework/keySets/{trustFrameworkKeySet_id}{?select,expand}";
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
func (m *TrustFrameworkKeySetItemRequestBuilder) CreateDeleteRequestInformation(options *TrustFrameworkKeySetItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get keySets from trustFramework
func (m *TrustFrameworkKeySetItemRequestBuilder) CreateGetRequestInformation(options *TrustFrameworkKeySetItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property keySets in trustFramework
func (m *TrustFrameworkKeySetItemRequestBuilder) CreatePatchRequestInformation(options *TrustFrameworkKeySetItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property keySets for trustFramework
func (m *TrustFrameworkKeySetItemRequestBuilder) Delete(options *TrustFrameworkKeySetItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
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
func (m *TrustFrameworkKeySetItemRequestBuilder) Get(options *TrustFrameworkKeySetItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeySetable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTrustFrameworkKeySetFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeySetable), nil
}
// GetActiveKey provides operations to call the getActiveKey method.
func (m *TrustFrameworkKeySetItemRequestBuilder) GetActiveKey()(*ib72c6292feaaf57eee55722f788b219f51c4f92a445271c20e5af998098f6daf.GetActiveKeyRequestBuilder) {
    return ib72c6292feaaf57eee55722f788b219f51c4f92a445271c20e5af998098f6daf.NewGetActiveKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property keySets in trustFramework
func (m *TrustFrameworkKeySetItemRequestBuilder) Patch(options *TrustFrameworkKeySetItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
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
