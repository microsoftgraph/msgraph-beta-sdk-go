package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i16903a7f82aef1b9172b084c2f688434ec7df2fe7382cd313d3acfe5d99975b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/assignments"
    i1c6baf2c64fdbaacd3e4b6e2fc919aa7f1ee2113534c7b651c6f577f13cdcbd9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/groupassignments"
    i7f208efc5d45cbfede2e60d9bbf7b141234bf2bb3b1aa6abb8ccde6a36185ad5 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates"
    i801a613b15d2cbc0073244db83bf53fd7b41404a5087e44c919ad86255b3948b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/runsummary"
    ibc0bdc7d62407627f799f1887a2a83100a7295907f2203c9b7be0ceb05727fa8 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/assign"
    ic9701e6869c3e935d0c0f54c628c153ce833fd193972ac4ccd97a4090f51262c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates"
    i3afdf7675877765716ac7dd1dadedd54dc760d42e9a7d4d3f623971af61f9ddd "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item"
    i4e99eee29f4d27b3c3ed9b39a732efaaadc4e736ad51030857642293b548d3cb "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item"
    i60b1d872fe3c79385687ac815dc9a345c65cf3a1d7475fcd16834aa195d0ef49 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/assignments/item"
    ib0522360ed077a6c16fe77b5b9aca7adca496e97d4d37f10bf42394220f94da8 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/groupassignments/item"
)

type DeviceShellScriptRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type DeviceShellScriptRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *DeviceShellScriptRequestBuilder) Assign()(*ibc0bdc7d62407627f799f1887a2a83100a7295907f2203c9b7be0ceb05727fa8.AssignRequestBuilder) {
    return ibc0bdc7d62407627f799f1887a2a83100a7295907f2203c9b7be0ceb05727fa8.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceShellScriptRequestBuilder) Assignments()(*i16903a7f82aef1b9172b084c2f688434ec7df2fe7382cd313d3acfe5d99975b1.AssignmentsRequestBuilder) {
    return i16903a7f82aef1b9172b084c2f688434ec7df2fe7382cd313d3acfe5d99975b1.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceShellScriptRequestBuilder) AssignmentsById(id string)(*i60b1d872fe3c79385687ac815dc9a345c65cf3a1d7475fcd16834aa195d0ef49.DeviceManagementScriptAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["deviceManagementScriptAssignment_id"] = id
    }
    return i60b1d872fe3c79385687ac815dc9a345c65cf3a1d7475fcd16834aa195d0ef49.NewDeviceManagementScriptAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewDeviceShellScriptRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceShellScriptRequestBuilder) {
    m := &DeviceShellScriptRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceManagement/deviceShellScripts/{deviceShellScript_id}{?select,expand}";
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
func NewDeviceShellScriptRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceShellScriptRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceShellScriptRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DeviceShellScriptRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DeviceShellScriptRequestBuilder) CreateGetRequestInformation(q func (value *DeviceShellScriptRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(DeviceShellScriptRequestBuilderGetQueryParameters)
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
func (m *DeviceShellScriptRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceShellScript, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DeviceShellScriptRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *DeviceShellScriptRequestBuilder) DeviceRunStates()(*ic9701e6869c3e935d0c0f54c628c153ce833fd193972ac4ccd97a4090f51262c.DeviceRunStatesRequestBuilder) {
    return ic9701e6869c3e935d0c0f54c628c153ce833fd193972ac4ccd97a4090f51262c.NewDeviceRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceShellScriptRequestBuilder) DeviceRunStatesById(id string)(*i4e99eee29f4d27b3c3ed9b39a732efaaadc4e736ad51030857642293b548d3cb.DeviceManagementScriptDeviceStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["deviceManagementScriptDeviceState_id"] = id
    }
    return i4e99eee29f4d27b3c3ed9b39a732efaaadc4e736ad51030857642293b548d3cb.NewDeviceManagementScriptDeviceStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceShellScriptRequestBuilder) Get(q func (value *DeviceShellScriptRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceShellScript, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceShellScript() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceShellScript), nil
}
func (m *DeviceShellScriptRequestBuilder) GroupAssignments()(*i1c6baf2c64fdbaacd3e4b6e2fc919aa7f1ee2113534c7b651c6f577f13cdcbd9.GroupAssignmentsRequestBuilder) {
    return i1c6baf2c64fdbaacd3e4b6e2fc919aa7f1ee2113534c7b651c6f577f13cdcbd9.NewGroupAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceShellScriptRequestBuilder) GroupAssignmentsById(id string)(*ib0522360ed077a6c16fe77b5b9aca7adca496e97d4d37f10bf42394220f94da8.DeviceManagementScriptGroupAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["deviceManagementScriptGroupAssignment_id"] = id
    }
    return ib0522360ed077a6c16fe77b5b9aca7adca496e97d4d37f10bf42394220f94da8.NewDeviceManagementScriptGroupAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceShellScriptRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceShellScript, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *DeviceShellScriptRequestBuilder) RunSummary()(*i801a613b15d2cbc0073244db83bf53fd7b41404a5087e44c919ad86255b3948b.RunSummaryRequestBuilder) {
    return i801a613b15d2cbc0073244db83bf53fd7b41404a5087e44c919ad86255b3948b.NewRunSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceShellScriptRequestBuilder) UserRunStates()(*i7f208efc5d45cbfede2e60d9bbf7b141234bf2bb3b1aa6abb8ccde6a36185ad5.UserRunStatesRequestBuilder) {
    return i7f208efc5d45cbfede2e60d9bbf7b141234bf2bb3b1aa6abb8ccde6a36185ad5.NewUserRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceShellScriptRequestBuilder) UserRunStatesById(id string)(*i3afdf7675877765716ac7dd1dadedd54dc760d42e9a7d4d3f623971af61f9ddd.DeviceManagementScriptUserStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["deviceManagementScriptUserState_id"] = id
    }
    return i3afdf7675877765716ac7dd1dadedd54dc760d42e9a7d4d3f623971af61f9ddd.NewDeviceManagementScriptUserStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
