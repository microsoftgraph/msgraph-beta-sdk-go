package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
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
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TrustFrameworkKeySetItemRequestBuilderDeleteOptions options for Delete
type TrustFrameworkKeySetItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TrustFrameworkKeySetItemRequestBuilderGetOptions options for Get
type TrustFrameworkKeySetItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TrustFrameworkKeySetItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
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
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TrustFrameworkKeySetable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewTrustFrameworkKeySetItemRequestBuilderInternal instantiates a new TrustFrameworkKeySetItemRequestBuilder and sets the default values.
func NewTrustFrameworkKeySetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TrustFrameworkKeySetItemRequestBuilder) {
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
func NewTrustFrameworkKeySetItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TrustFrameworkKeySetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTrustFrameworkKeySetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property keySets for trustFramework
func (m *TrustFrameworkKeySetItemRequestBuilder) CreateDeleteRequestInformation(options *TrustFrameworkKeySetItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *TrustFrameworkKeySetItemRequestBuilder) CreateGetRequestInformation(options *TrustFrameworkKeySetItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *TrustFrameworkKeySetItemRequestBuilder) CreatePatchRequestInformation(options *TrustFrameworkKeySetItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *TrustFrameworkKeySetItemRequestBuilder) Delete(options *TrustFrameworkKeySetItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *TrustFrameworkKeySetItemRequestBuilder) GenerateKey()(*ie04ccf263002e1d9853509d405e0878dfbf7c145b0a12ed976019d944eb6fe4d.GenerateKeyRequestBuilder) {
    return ie04ccf263002e1d9853509d405e0878dfbf7c145b0a12ed976019d944eb6fe4d.NewGenerateKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get keySets from trustFramework
func (m *TrustFrameworkKeySetItemRequestBuilder) Get(options *TrustFrameworkKeySetItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TrustFrameworkKeySetable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateTrustFrameworkKeySetFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TrustFrameworkKeySetable), nil
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
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *TrustFrameworkKeySetItemRequestBuilder) UploadCertificate()(*i9d1297a38f5879454a214ecb884be348ff387c1b4b4287a28598568142a017af.UploadCertificateRequestBuilder) {
    return i9d1297a38f5879454a214ecb884be348ff387c1b4b4287a28598568142a017af.NewUploadCertificateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TrustFrameworkKeySetItemRequestBuilder) UploadPkcs12()(*i4df38ca1b862714120e25419c3ed380c7c772471f65cbf3e0d620c5a0d60f6bf.UploadPkcs12RequestBuilder) {
    return i4df38ca1b862714120e25419c3ed380c7c772471f65cbf3e0d620c5a0d60f6bf.NewUploadPkcs12RequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TrustFrameworkKeySetItemRequestBuilder) UploadSecret()(*i78db104952796dbfc3b27052328de3db351b4c77d4b5dffcd1d52a18b26c3220.UploadSecretRequestBuilder) {
    return i78db104952796dbfc3b27052328de3db351b4c77d4b5dffcd1d52a18b26c3220.NewUploadSecretRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
