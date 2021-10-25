package synchronization

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i11872257330c9dfde50ea9af96d17c64c29f6df124435f8ebc6077cdc9a623f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/templates"
    i11ba372ce609325ea47abb7e9a6878ff77197345ccc0b718db49c8302beae7f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/ping"
    i24aca1a09e1a9371d3ed55ec89c2a9aa93cd22ef9b2809e21ef705ec431f9373 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/jobs"
    ib7cb71c543f69663ca4aa8935b15c63566ebfd615a46160e1d5340cba65e4050 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/acquireaccesstoken"
    i81be635c95f0b471e65265af6e757fe50614612e56f6d562a96c5a5bb6d49b5a "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/templates/item"
    ie696cb911da1ddd94a64281ef404310fd8f0821d842d07dc9443fb4922f0fed2 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/jobs/item"
)

type SynchronizationRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type SynchronizationRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *SynchronizationRequestBuilder) AcquireAccessToken()(*ib7cb71c543f69663ca4aa8935b15c63566ebfd615a46160e1d5340cba65e4050.AcquireAccessTokenRequestBuilder) {
    return ib7cb71c543f69663ca4aa8935b15c63566ebfd615a46160e1d5340cba65e4050.NewAcquireAccessTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewSynchronizationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SynchronizationRequestBuilder) {
    m := &SynchronizationRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/applications/{application_id}/synchronization{?select,expand}";
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
func (m *SynchronizationRequestBuilder) Jobs()(*i24aca1a09e1a9371d3ed55ec89c2a9aa93cd22ef9b2809e21ef705ec431f9373.JobsRequestBuilder) {
    return i24aca1a09e1a9371d3ed55ec89c2a9aa93cd22ef9b2809e21ef705ec431f9373.NewJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SynchronizationRequestBuilder) JobsById(id string)(*ie696cb911da1ddd94a64281ef404310fd8f0821d842d07dc9443fb4922f0fed2.SynchronizationJobRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["synchronizationJob_id"] = id
    }
    return ie696cb911da1ddd94a64281ef404310fd8f0821d842d07dc9443fb4922f0fed2.NewSynchronizationJobRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *SynchronizationRequestBuilder) Ping()(*i11ba372ce609325ea47abb7e9a6878ff77197345ccc0b718db49c8302beae7f0.PingRequestBuilder) {
    return i11ba372ce609325ea47abb7e9a6878ff77197345ccc0b718db49c8302beae7f0.NewPingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SynchronizationRequestBuilder) Templates()(*i11872257330c9dfde50ea9af96d17c64c29f6df124435f8ebc6077cdc9a623f1.TemplatesRequestBuilder) {
    return i11872257330c9dfde50ea9af96d17c64c29f6df124435f8ebc6077cdc9a623f1.NewTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SynchronizationRequestBuilder) TemplatesById(id string)(*i81be635c95f0b471e65265af6e757fe50614612e56f6d562a96c5a5bb6d49b5a.SynchronizationTemplateRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["synchronizationTemplate_id"] = id
    }
    return i81be635c95f0b471e65265af6e757fe50614612e56f6d562a96c5a5bb6d49b5a.NewSynchronizationTemplateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
