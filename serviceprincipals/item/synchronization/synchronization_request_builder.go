package synchronization

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i441115fbd107244dd6504672d846ac2adcc602e6da088866d93d998f95107c40 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/ping"
    i8d4af922a9d1b5130f294d250ea800e845868ae335b95eccdbe95097a1393573 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/templates"
    i9f2b960ccd9a3439ea358b94aa154106d5794fdd4bdd4786b63fb9896b804094 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs"
    ie4e7228c8413ba04b04a7f1fac840439946274cff81aadbd7cf5252263fe1cc4 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/acquireaccesstoken"
    ibe6466d39a6e0f7d78fef7ec330777772ce9e4942297f090c6e91a7b4f7da1a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs/item"
    idbe71426e8a7cf17f46df86ea54d1a8d22e20242b9a4e0a5c7d1e8c8405dd581 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/templates/item"
)

type SynchronizationRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type SynchronizationRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *SynchronizationRequestBuilder) AcquireAccessToken()(*ie4e7228c8413ba04b04a7f1fac840439946274cff81aadbd7cf5252263fe1cc4.AcquireAccessTokenRequestBuilder) {
    return ie4e7228c8413ba04b04a7f1fac840439946274cff81aadbd7cf5252263fe1cc4.NewAcquireAccessTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewSynchronizationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SynchronizationRequestBuilder) {
    m := &SynchronizationRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/servicePrincipals/{servicePrincipal_id}/synchronization{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewSynchronizationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SynchronizationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSynchronizationRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *SynchronizationRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SynchronizationRequestBuilder) CreateGetRequestInformation(q func (value *SynchronizationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(SynchronizationRequestBuilderGetQueryParameters)
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
func (m *SynchronizationRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Synchronization, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SynchronizationRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *SynchronizationRequestBuilder) Get(q func (value *SynchronizationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Synchronization, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSynchronization() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Synchronization), nil
}
func (m *SynchronizationRequestBuilder) Jobs()(*i9f2b960ccd9a3439ea358b94aa154106d5794fdd4bdd4786b63fb9896b804094.JobsRequestBuilder) {
    return i9f2b960ccd9a3439ea358b94aa154106d5794fdd4bdd4786b63fb9896b804094.NewJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SynchronizationRequestBuilder) JobsById(id string)(*ibe6466d39a6e0f7d78fef7ec330777772ce9e4942297f090c6e91a7b4f7da1a8.SynchronizationJobRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["synchronizationJob_id"] = id
    }
    return ibe6466d39a6e0f7d78fef7ec330777772ce9e4942297f090c6e91a7b4f7da1a8.NewSynchronizationJobRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SynchronizationRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Synchronization, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *SynchronizationRequestBuilder) Ping()(*i441115fbd107244dd6504672d846ac2adcc602e6da088866d93d998f95107c40.PingRequestBuilder) {
    return i441115fbd107244dd6504672d846ac2adcc602e6da088866d93d998f95107c40.NewPingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SynchronizationRequestBuilder) Templates()(*i8d4af922a9d1b5130f294d250ea800e845868ae335b95eccdbe95097a1393573.TemplatesRequestBuilder) {
    return i8d4af922a9d1b5130f294d250ea800e845868ae335b95eccdbe95097a1393573.NewTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SynchronizationRequestBuilder) TemplatesById(id string)(*idbe71426e8a7cf17f46df86ea54d1a8d22e20242b9a4e0a5c7d1e8c8405dd581.SynchronizationTemplateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["synchronizationTemplate_id"] = id
    }
    return idbe71426e8a7cf17f46df86ea54d1a8d22e20242b9a4e0a5c7d1e8c8405dd581.NewSynchronizationTemplateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
