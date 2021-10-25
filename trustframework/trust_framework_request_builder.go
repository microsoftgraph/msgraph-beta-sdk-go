package trustframework

import (
    i45ac5b9b0bfa1cab1bbb052f4f9c7d4e05cfb96edb3cb1d48ad9d002664d39d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/trustframework/policies"
    i7b667425d5ba5070b20ca9a446c105ed8c4749fa25d168849c6e0258ab640270 "github.com/microsoftgraph/msgraph-beta-sdk-go/trustframework/keysets"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i130f0626fe3c071232b6f7dc633093963e558c51511332c8d4c462636204a280 "github.com/microsoftgraph/msgraph-beta-sdk-go/trustframework/keysets/item"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i76c4dfef5c53e4b89a114a49ab372d39e04290df03064db59620347dee08cbcd "github.com/microsoftgraph/msgraph-beta-sdk-go/trustframework/policies/item"
)

type TrustFrameworkRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type TrustFrameworkRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func NewTrustFrameworkRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TrustFrameworkRequestBuilder) {
    m := &TrustFrameworkRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/trustFramework{?select,expand}";
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
func NewTrustFrameworkRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TrustFrameworkRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTrustFrameworkRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *TrustFrameworkRequestBuilder) CreateGetRequestInformation(q func (value *TrustFrameworkRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(TrustFrameworkRequestBuilderGetQueryParameters)
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
func (m *TrustFrameworkRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TrustFramework, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *TrustFrameworkRequestBuilder) Get(q func (value *TrustFrameworkRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TrustFramework, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewTrustFramework() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TrustFramework), nil
}
func (m *TrustFrameworkRequestBuilder) KeySets()(*i7b667425d5ba5070b20ca9a446c105ed8c4749fa25d168849c6e0258ab640270.KeySetsRequestBuilder) {
    return i7b667425d5ba5070b20ca9a446c105ed8c4749fa25d168849c6e0258ab640270.NewKeySetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TrustFrameworkRequestBuilder) KeySetsById(id string)(*i130f0626fe3c071232b6f7dc633093963e558c51511332c8d4c462636204a280.TrustFrameworkKeySetRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["trustFrameworkKeySet_id"] = id
    }
    return i130f0626fe3c071232b6f7dc633093963e558c51511332c8d4c462636204a280.NewTrustFrameworkKeySetRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TrustFrameworkRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TrustFramework, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *TrustFrameworkRequestBuilder) Policies()(*i45ac5b9b0bfa1cab1bbb052f4f9c7d4e05cfb96edb3cb1d48ad9d002664d39d9.PoliciesRequestBuilder) {
    return i45ac5b9b0bfa1cab1bbb052f4f9c7d4e05cfb96edb3cb1d48ad9d002664d39d9.NewPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TrustFrameworkRequestBuilder) PoliciesById(id string)(*i76c4dfef5c53e4b89a114a49ab372d39e04290df03064db59620347dee08cbcd.TrustFrameworkPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["trustFrameworkPolicy_id"] = id
    }
    return i76c4dfef5c53e4b89a114a49ab372d39e04290df03064db59620347dee08cbcd.NewTrustFrameworkPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
