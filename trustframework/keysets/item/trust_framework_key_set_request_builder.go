package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i4df38ca1b862714120e25419c3ed380c7c772471f65cbf3e0d620c5a0d60f6bf "github.com/microsoftgraph/msgraph-beta-sdk-go/trustframework/keysets/item/uploadpkcs12"
    i78db104952796dbfc3b27052328de3db351b4c77d4b5dffcd1d52a18b26c3220 "github.com/microsoftgraph/msgraph-beta-sdk-go/trustframework/keysets/item/uploadsecret"
    i9d1297a38f5879454a214ecb884be348ff387c1b4b4287a28598568142a017af "github.com/microsoftgraph/msgraph-beta-sdk-go/trustframework/keysets/item/uploadcertificate"
    ib72c6292feaaf57eee55722f788b219f51c4f92a445271c20e5af998098f6daf "github.com/microsoftgraph/msgraph-beta-sdk-go/trustframework/keysets/item/getactivekey"
    ie04ccf263002e1d9853509d405e0878dfbf7c145b0a12ed976019d944eb6fe4d "github.com/microsoftgraph/msgraph-beta-sdk-go/trustframework/keysets/item/generatekey"
)

// TrustFrameworkKeySetRequestBuilder builds and executes requests for operations under \trustFramework\keySets\{trustFrameworkKeySet-id}
type TrustFrameworkKeySetRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TrustFrameworkKeySetRequestBuilderDeleteOptions options for Delete
type TrustFrameworkKeySetRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TrustFrameworkKeySetRequestBuilderGetOptions options for Get
type TrustFrameworkKeySetRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TrustFrameworkKeySetRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TrustFrameworkKeySetRequestBuilderGetQueryParameters get keySets from trustFramework
type TrustFrameworkKeySetRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// TrustFrameworkKeySetRequestBuilderPatchOptions options for Patch
type TrustFrameworkKeySetRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TrustFrameworkKeySet;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewTrustFrameworkKeySetRequestBuilderInternal instantiates a new TrustFrameworkKeySetRequestBuilder and sets the default values.
func NewTrustFrameworkKeySetRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TrustFrameworkKeySetRequestBuilder) {
    m := &TrustFrameworkKeySetRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/trustFramework/keySets/{trustFrameworkKeySet_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTrustFrameworkKeySetRequestBuilder instantiates a new TrustFrameworkKeySetRequestBuilder and sets the default values.
func NewTrustFrameworkKeySetRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TrustFrameworkKeySetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTrustFrameworkKeySetRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property keySets for trustFramework
func (m *TrustFrameworkKeySetRequestBuilder) CreateDeleteRequestInformation(options *TrustFrameworkKeySetRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get keySets from trustFramework
func (m *TrustFrameworkKeySetRequestBuilder) CreateGetRequestInformation(options *TrustFrameworkKeySetRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property keySets in trustFramework
func (m *TrustFrameworkKeySetRequestBuilder) CreatePatchRequestInformation(options *TrustFrameworkKeySetRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property keySets for trustFramework
func (m *TrustFrameworkKeySetRequestBuilder) Delete(options *TrustFrameworkKeySetRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *TrustFrameworkKeySetRequestBuilder) GenerateKey()(*ie04ccf263002e1d9853509d405e0878dfbf7c145b0a12ed976019d944eb6fe4d.GenerateKeyRequestBuilder) {
    return ie04ccf263002e1d9853509d405e0878dfbf7c145b0a12ed976019d944eb6fe4d.NewGenerateKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get keySets from trustFramework
func (m *TrustFrameworkKeySetRequestBuilder) Get(options *TrustFrameworkKeySetRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TrustFrameworkKeySet, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewTrustFrameworkKeySet() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TrustFrameworkKeySet), nil
}
// GetActiveKey builds and executes requests for operations under \trustFramework\keySets\{trustFrameworkKeySet-id}\microsoft.graph.getActiveKey()
func (m *TrustFrameworkKeySetRequestBuilder) GetActiveKey()(*ib72c6292feaaf57eee55722f788b219f51c4f92a445271c20e5af998098f6daf.GetActiveKeyRequestBuilder) {
    return ib72c6292feaaf57eee55722f788b219f51c4f92a445271c20e5af998098f6daf.NewGetActiveKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property keySets in trustFramework
func (m *TrustFrameworkKeySetRequestBuilder) Patch(options *TrustFrameworkKeySetRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *TrustFrameworkKeySetRequestBuilder) UploadCertificate()(*i9d1297a38f5879454a214ecb884be348ff387c1b4b4287a28598568142a017af.UploadCertificateRequestBuilder) {
    return i9d1297a38f5879454a214ecb884be348ff387c1b4b4287a28598568142a017af.NewUploadCertificateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TrustFrameworkKeySetRequestBuilder) UploadPkcs12()(*i4df38ca1b862714120e25419c3ed380c7c772471f65cbf3e0d620c5a0d60f6bf.UploadPkcs12RequestBuilder) {
    return i4df38ca1b862714120e25419c3ed380c7c772471f65cbf3e0d620c5a0d60f6bf.NewUploadPkcs12RequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TrustFrameworkKeySetRequestBuilder) UploadSecret()(*i78db104952796dbfc3b27052328de3db351b4c77d4b5dffcd1d52a18b26c3220.UploadSecretRequestBuilder) {
    return i78db104952796dbfc3b27052328de3db351b4c77d4b5dffcd1d52a18b26c3220.NewUploadSecretRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
