package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i14f99bca7a957a331b5e7cb13fd8a56e803d7b84d2e6c38be90eb5551714c94e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/embeddedsimactivationcodepools/item/assignments"
    i1a8c6326f7b74495a2d4a1da0167deac80e12606e08ca1320fb159025e628c81 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/embeddedsimactivationcodepools/item/assign"
    ic4dac275164f8d53c7fa85048cdeb7907d2f7d42ec422190e6843e16ab03bfd3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/embeddedsimactivationcodepools/item/devicestates"
    i4a0e4526cd89dcbf6b464d731ad1079beb451ea477a00769d03119affe5e8e45 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/embeddedsimactivationcodepools/item/assignments/item"
    if044a617b4d30ce4a26c463e72b94403997b338b8aeafd10eb8d535111e4cc01 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/embeddedsimactivationcodepools/item/devicestates/item"
)

type EmbeddedSIMActivationCodePoolRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type EmbeddedSIMActivationCodePoolRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *EmbeddedSIMActivationCodePoolRequestBuilder) Assign()(*i1a8c6326f7b74495a2d4a1da0167deac80e12606e08ca1320fb159025e628c81.AssignRequestBuilder) {
    return i1a8c6326f7b74495a2d4a1da0167deac80e12606e08ca1320fb159025e628c81.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EmbeddedSIMActivationCodePoolRequestBuilder) Assignments()(*i14f99bca7a957a331b5e7cb13fd8a56e803d7b84d2e6c38be90eb5551714c94e.AssignmentsRequestBuilder) {
    return i14f99bca7a957a331b5e7cb13fd8a56e803d7b84d2e6c38be90eb5551714c94e.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EmbeddedSIMActivationCodePoolRequestBuilder) AssignmentsById(id string)(*i4a0e4526cd89dcbf6b464d731ad1079beb451ea477a00769d03119affe5e8e45.EmbeddedSIMActivationCodePoolAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["embeddedSIMActivationCodePoolAssignment_id"] = id
    }
    return i4a0e4526cd89dcbf6b464d731ad1079beb451ea477a00769d03119affe5e8e45.NewEmbeddedSIMActivationCodePoolAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewEmbeddedSIMActivationCodePoolRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EmbeddedSIMActivationCodePoolRequestBuilder) {
    m := &EmbeddedSIMActivationCodePoolRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceManagement/embeddedSIMActivationCodePools/{embeddedSIMActivationCodePool_id}{?select,expand}";
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
func NewEmbeddedSIMActivationCodePoolRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EmbeddedSIMActivationCodePoolRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEmbeddedSIMActivationCodePoolRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *EmbeddedSIMActivationCodePoolRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EmbeddedSIMActivationCodePoolRequestBuilder) CreateGetRequestInformation(q func (value *EmbeddedSIMActivationCodePoolRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(EmbeddedSIMActivationCodePoolRequestBuilderGetQueryParameters)
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
func (m *EmbeddedSIMActivationCodePoolRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EmbeddedSIMActivationCodePool, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EmbeddedSIMActivationCodePoolRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *EmbeddedSIMActivationCodePoolRequestBuilder) DeviceStates()(*ic4dac275164f8d53c7fa85048cdeb7907d2f7d42ec422190e6843e16ab03bfd3.DeviceStatesRequestBuilder) {
    return ic4dac275164f8d53c7fa85048cdeb7907d2f7d42ec422190e6843e16ab03bfd3.NewDeviceStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EmbeddedSIMActivationCodePoolRequestBuilder) DeviceStatesById(id string)(*if044a617b4d30ce4a26c463e72b94403997b338b8aeafd10eb8d535111e4cc01.EmbeddedSIMDeviceStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["embeddedSIMDeviceState_id"] = id
    }
    return if044a617b4d30ce4a26c463e72b94403997b338b8aeafd10eb8d535111e4cc01.NewEmbeddedSIMDeviceStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EmbeddedSIMActivationCodePoolRequestBuilder) Get(q func (value *EmbeddedSIMActivationCodePoolRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EmbeddedSIMActivationCodePool, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEmbeddedSIMActivationCodePool() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EmbeddedSIMActivationCodePool), nil
}
func (m *EmbeddedSIMActivationCodePoolRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EmbeddedSIMActivationCodePool, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
