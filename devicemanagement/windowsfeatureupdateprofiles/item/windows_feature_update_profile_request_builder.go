package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i17ae4c396051acc3d8969fe0d510a15f08f74b0af97608557d2125c16de1bb7d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsfeatureupdateprofiles/item/assign"
    i75eee03eb366df128ee3931df150a0db3df6df2080e6ad94a30da4ff8b885d0f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsfeatureupdateprofiles/item/deviceupdatestates"
    ib9f19a8a1ed79d6335dbd9524123a1a111e94345212d4f98f256d4d77fb64865 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsfeatureupdateprofiles/item/assignments"
    i9fc8cb618d1d2cb2bd52c185cd35fcf810ea303a5692d3f159788228e6d3514f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsfeatureupdateprofiles/item/assignments/item"
    ic56df2819be27ad9c364885a459ad3cb4b53f09c834db6d19a528131b7e27fa4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsfeatureupdateprofiles/item/deviceupdatestates/item"
)

type WindowsFeatureUpdateProfileRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type WindowsFeatureUpdateProfileRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *WindowsFeatureUpdateProfileRequestBuilder) Assign()(*i17ae4c396051acc3d8969fe0d510a15f08f74b0af97608557d2125c16de1bb7d.AssignRequestBuilder) {
    return i17ae4c396051acc3d8969fe0d510a15f08f74b0af97608557d2125c16de1bb7d.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WindowsFeatureUpdateProfileRequestBuilder) Assignments()(*ib9f19a8a1ed79d6335dbd9524123a1a111e94345212d4f98f256d4d77fb64865.AssignmentsRequestBuilder) {
    return ib9f19a8a1ed79d6335dbd9524123a1a111e94345212d4f98f256d4d77fb64865.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WindowsFeatureUpdateProfileRequestBuilder) AssignmentsById(id string)(*i9fc8cb618d1d2cb2bd52c185cd35fcf810ea303a5692d3f159788228e6d3514f.WindowsFeatureUpdateProfileAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["windowsFeatureUpdateProfileAssignment_id"] = id
    }
    return i9fc8cb618d1d2cb2bd52c185cd35fcf810ea303a5692d3f159788228e6d3514f.NewWindowsFeatureUpdateProfileAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewWindowsFeatureUpdateProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsFeatureUpdateProfileRequestBuilder) {
    m := &WindowsFeatureUpdateProfileRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceManagement/windowsFeatureUpdateProfiles/{windowsFeatureUpdateProfile_id}{?select,expand}";
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
func NewWindowsFeatureUpdateProfileRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsFeatureUpdateProfileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsFeatureUpdateProfileRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *WindowsFeatureUpdateProfileRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *WindowsFeatureUpdateProfileRequestBuilder) CreateGetRequestInformation(q func (value *WindowsFeatureUpdateProfileRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(WindowsFeatureUpdateProfileRequestBuilderGetQueryParameters)
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
func (m *WindowsFeatureUpdateProfileRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsFeatureUpdateProfile, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *WindowsFeatureUpdateProfileRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *WindowsFeatureUpdateProfileRequestBuilder) DeviceUpdateStates()(*i75eee03eb366df128ee3931df150a0db3df6df2080e6ad94a30da4ff8b885d0f.DeviceUpdateStatesRequestBuilder) {
    return i75eee03eb366df128ee3931df150a0db3df6df2080e6ad94a30da4ff8b885d0f.NewDeviceUpdateStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WindowsFeatureUpdateProfileRequestBuilder) DeviceUpdateStatesById(id string)(*ic56df2819be27ad9c364885a459ad3cb4b53f09c834db6d19a528131b7e27fa4.WindowsUpdateStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["windowsUpdateState_id"] = id
    }
    return ic56df2819be27ad9c364885a459ad3cb4b53f09c834db6d19a528131b7e27fa4.NewWindowsUpdateStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WindowsFeatureUpdateProfileRequestBuilder) Get(q func (value *WindowsFeatureUpdateProfileRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsFeatureUpdateProfile, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWindowsFeatureUpdateProfile() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsFeatureUpdateProfile), nil
}
func (m *WindowsFeatureUpdateProfileRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsFeatureUpdateProfile, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
