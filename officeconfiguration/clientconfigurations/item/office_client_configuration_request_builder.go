package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1068f00153d896756e5c75882c5dbee1a31ff4ad6c6cd1299946c4acf1e1932d "github.com/microsoftgraph/msgraph-beta-sdk-go/officeconfiguration/clientconfigurations/item/policypayload"
    i6ee98b33165e3fc15ddb585c37c29eb4d3887d300742016bd30e69c1dff7f0b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/officeconfiguration/clientconfigurations/item/userpreferencepayload"
    ib708431220be37864eb91ad6ad04c27754da18439883881ba9f54b62ae1d2206 "github.com/microsoftgraph/msgraph-beta-sdk-go/officeconfiguration/clientconfigurations/item/assign"
    ie1103875ba3eb7fa000dd34e6685f4073add80b4c464e341669e479f36621a4c "github.com/microsoftgraph/msgraph-beta-sdk-go/officeconfiguration/clientconfigurations/item/assignments"
    ie58d03e4a8a1ff62a18e2ec91d18b2410cce187a18440bbde6c2754055d5ee34 "github.com/microsoftgraph/msgraph-beta-sdk-go/officeconfiguration/clientconfigurations/item/assignments/item"
)

type OfficeClientConfigurationRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type OfficeClientConfigurationRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *OfficeClientConfigurationRequestBuilder) Assign()(*ib708431220be37864eb91ad6ad04c27754da18439883881ba9f54b62ae1d2206.AssignRequestBuilder) {
    return ib708431220be37864eb91ad6ad04c27754da18439883881ba9f54b62ae1d2206.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OfficeClientConfigurationRequestBuilder) Assignments()(*ie1103875ba3eb7fa000dd34e6685f4073add80b4c464e341669e479f36621a4c.AssignmentsRequestBuilder) {
    return ie1103875ba3eb7fa000dd34e6685f4073add80b4c464e341669e479f36621a4c.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OfficeClientConfigurationRequestBuilder) AssignmentsById(id string)(*ie58d03e4a8a1ff62a18e2ec91d18b2410cce187a18440bbde6c2754055d5ee34.OfficeClientConfigurationAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["officeClientConfigurationAssignment_id"] = id
    }
    return ie58d03e4a8a1ff62a18e2ec91d18b2410cce187a18440bbde6c2754055d5ee34.NewOfficeClientConfigurationAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewOfficeClientConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OfficeClientConfigurationRequestBuilder) {
    m := &OfficeClientConfigurationRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/officeConfiguration/clientConfigurations/{officeClientConfiguration_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewOfficeClientConfigurationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OfficeClientConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOfficeClientConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *OfficeClientConfigurationRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *OfficeClientConfigurationRequestBuilder) CreateGetRequestInformation(q func (value *OfficeClientConfigurationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(OfficeClientConfigurationRequestBuilderGetQueryParameters)
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
func (m *OfficeClientConfigurationRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OfficeClientConfiguration, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *OfficeClientConfigurationRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *OfficeClientConfigurationRequestBuilder) Get(q func (value *OfficeClientConfigurationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OfficeClientConfiguration, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOfficeClientConfiguration() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OfficeClientConfiguration), nil
}
func (m *OfficeClientConfigurationRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OfficeClientConfiguration, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *OfficeClientConfigurationRequestBuilder) PolicyPayload()(*i1068f00153d896756e5c75882c5dbee1a31ff4ad6c6cd1299946c4acf1e1932d.PolicyPayloadRequestBuilder) {
    return i1068f00153d896756e5c75882c5dbee1a31ff4ad6c6cd1299946c4acf1e1932d.NewPolicyPayloadRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OfficeClientConfigurationRequestBuilder) UserPreferencePayload()(*i6ee98b33165e3fc15ddb585c37c29eb4d3887d300742016bd30e69c1dff7f0b8.UserPreferencePayloadRequestBuilder) {
    return i6ee98b33165e3fc15ddb585c37c29eb4d3887d300742016bd30e69c1dff7f0b8.NewUserPreferencePayloadRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
