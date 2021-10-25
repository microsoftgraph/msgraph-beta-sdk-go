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

type TrustFrameworkKeySetRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type TrustFrameworkKeySetRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func NewTrustFrameworkKeySetRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TrustFrameworkKeySetRequestBuilder) {
    m := &TrustFrameworkKeySetRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/trustFramework/keySets/{trustFrameworkKeySet_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewTrustFrameworkKeySetRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TrustFrameworkKeySetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTrustFrameworkKeySetRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *TrustFrameworkKeySetRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *TrustFrameworkKeySetRequestBuilder) CreateGetRequestInformation(q func (value *TrustFrameworkKeySetRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(TrustFrameworkKeySetRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *TrustFrameworkKeySetRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TrustFrameworkKeySet, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *TrustFrameworkKeySetRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *TrustFrameworkKeySetRequestBuilder) GenerateKey()(*ie04ccf263002e1d9853509d405e0878dfbf7c145b0a12ed976019d944eb6fe4d.GenerateKeyRequestBuilder) {
    return ie04ccf263002e1d9853509d405e0878dfbf7c145b0a12ed976019d944eb6fe4d.NewGenerateKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TrustFrameworkKeySetRequestBuilder) Get(q func (value *TrustFrameworkKeySetRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TrustFrameworkKeySet, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewTrustFrameworkKeySet() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TrustFrameworkKeySet), nil
}
func (m *TrustFrameworkKeySetRequestBuilder) GetActiveKey()(*ib72c6292feaaf57eee55722f788b219f51c4f92a445271c20e5af998098f6daf.GetActiveKeyRequestBuilder) {
    return ib72c6292feaaf57eee55722f788b219f51c4f92a445271c20e5af998098f6daf.NewGetActiveKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TrustFrameworkKeySetRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TrustFrameworkKeySet, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
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
