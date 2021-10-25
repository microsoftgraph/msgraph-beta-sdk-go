package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i06231952f18c3df2ed3c4845512be6748846cddc42430ea22f0110bc844935bb "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates"
    i6e0733508da7ff0356e2ca316d29bfbd6cda93ef5db9a39817c2f424b503dec6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/assignments"
    i9d58942b600e533779502c28f94ba3cb6ca7a519ae97eef6c53e20d6be174561 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates"
    if33d8d67c34ee54fe873013b66ebc2b6e3dc6e392d8ab242161653be770f5d4b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/assign"
    if54f4ed088e5af17a948f9194c1f6e7aa23d72ff5ebeb5ae40c4ac2c1510ee2a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/runsummary"
    ifbbe46308813fcfd21a07bcac76e413b36786c8c9c9bcea463584d058d80cdaa "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/groupassignments"
    i3e402295bee0361812ae67ce9531f7061435900ac0c50c895a12d292d4bc0ec5 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/groupassignments/item"
    i47cd5e12b88aa45ee308ba79574764b3fc2f141f8bbf112dc9af34ef38267657 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item"
    i84a45e8ade4a4e0a4b92c01abedac4601754d6c53b15136864f2af744f8c2b02 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/assignments/item"
    icf8e74a74bc8538b6c9fa9f2bec4d9d624de373f52d2137852ef78d9754c2779 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item"
)

type DeviceManagementScriptRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type DeviceManagementScriptRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *DeviceManagementScriptRequestBuilder) Assign()(*if33d8d67c34ee54fe873013b66ebc2b6e3dc6e392d8ab242161653be770f5d4b.AssignRequestBuilder) {
    return if33d8d67c34ee54fe873013b66ebc2b6e3dc6e392d8ab242161653be770f5d4b.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceManagementScriptRequestBuilder) Assignments()(*i6e0733508da7ff0356e2ca316d29bfbd6cda93ef5db9a39817c2f424b503dec6.AssignmentsRequestBuilder) {
    return i6e0733508da7ff0356e2ca316d29bfbd6cda93ef5db9a39817c2f424b503dec6.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceManagementScriptRequestBuilder) AssignmentsById(id string)(*i84a45e8ade4a4e0a4b92c01abedac4601754d6c53b15136864f2af744f8c2b02.DeviceManagementScriptAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptAssignment_id"] = id
    }
    return i84a45e8ade4a4e0a4b92c01abedac4601754d6c53b15136864f2af744f8c2b02.NewDeviceManagementScriptAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewDeviceManagementScriptRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementScriptRequestBuilder) {
    m := &DeviceManagementScriptRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceManagement/deviceManagementScripts/{deviceManagementScript_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewDeviceManagementScriptRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementScriptRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementScriptRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DeviceManagementScriptRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DeviceManagementScriptRequestBuilder) CreateGetRequestInformation(q func (value *DeviceManagementScriptRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(DeviceManagementScriptRequestBuilderGetQueryParameters)
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
func (m *DeviceManagementScriptRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementScript, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DeviceManagementScriptRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *DeviceManagementScriptRequestBuilder) DeviceRunStates()(*i06231952f18c3df2ed3c4845512be6748846cddc42430ea22f0110bc844935bb.DeviceRunStatesRequestBuilder) {
    return i06231952f18c3df2ed3c4845512be6748846cddc42430ea22f0110bc844935bb.NewDeviceRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceManagementScriptRequestBuilder) DeviceRunStatesById(id string)(*i47cd5e12b88aa45ee308ba79574764b3fc2f141f8bbf112dc9af34ef38267657.DeviceManagementScriptDeviceStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptDeviceState_id"] = id
    }
    return i47cd5e12b88aa45ee308ba79574764b3fc2f141f8bbf112dc9af34ef38267657.NewDeviceManagementScriptDeviceStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementScriptRequestBuilder) Get(q func (value *DeviceManagementScriptRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementScript, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceManagementScript() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementScript), nil
}
func (m *DeviceManagementScriptRequestBuilder) GroupAssignments()(*ifbbe46308813fcfd21a07bcac76e413b36786c8c9c9bcea463584d058d80cdaa.GroupAssignmentsRequestBuilder) {
    return ifbbe46308813fcfd21a07bcac76e413b36786c8c9c9bcea463584d058d80cdaa.NewGroupAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceManagementScriptRequestBuilder) GroupAssignmentsById(id string)(*i3e402295bee0361812ae67ce9531f7061435900ac0c50c895a12d292d4bc0ec5.DeviceManagementScriptGroupAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptGroupAssignment_id"] = id
    }
    return i3e402295bee0361812ae67ce9531f7061435900ac0c50c895a12d292d4bc0ec5.NewDeviceManagementScriptGroupAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementScriptRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementScript, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *DeviceManagementScriptRequestBuilder) RunSummary()(*if54f4ed088e5af17a948f9194c1f6e7aa23d72ff5ebeb5ae40c4ac2c1510ee2a.RunSummaryRequestBuilder) {
    return if54f4ed088e5af17a948f9194c1f6e7aa23d72ff5ebeb5ae40c4ac2c1510ee2a.NewRunSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceManagementScriptRequestBuilder) UserRunStates()(*i9d58942b600e533779502c28f94ba3cb6ca7a519ae97eef6c53e20d6be174561.UserRunStatesRequestBuilder) {
    return i9d58942b600e533779502c28f94ba3cb6ca7a519ae97eef6c53e20d6be174561.NewUserRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceManagementScriptRequestBuilder) UserRunStatesById(id string)(*icf8e74a74bc8538b6c9fa9f2bec4d9d624de373f52d2137852ef78d9754c2779.DeviceManagementScriptUserStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptUserState_id"] = id
    }
    return icf8e74a74bc8538b6c9fa9f2bec4d9d624de373f52d2137852ef78d9754c2779.NewDeviceManagementScriptUserStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
