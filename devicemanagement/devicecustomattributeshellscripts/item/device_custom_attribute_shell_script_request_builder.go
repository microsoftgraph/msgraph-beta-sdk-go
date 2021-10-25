package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i367a45f04f002f42b81ce6533585aeeba4211ee6e776fb0ddb4cedf64aa10d6f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/runsummary"
    i494c760a2f26845bb6a83e1609802c363f94efb574cc06272e4db09ffdd5950e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/groupassignments"
    i5860a2467c9825730f7e5fe8b8587f837d06a6bce92dbb17df1197b8c3e63d0c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/assign"
    i9caa21903d914933c04aab2668b67aaf07ff8c88530bbc0f796266bff7b47bb2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/assignments"
    ia6cd2ea0f3d71e372a080d28d01095fd6857f4d1b6250e215940436d6784cbba "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates"
    ib4c60324ec4c204995c1740d3c32783fe43242863d0216859b5414b0a63bd4a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates"
    i076188fa8b5444fe1eee63d44522c5efaa3b78288e012e2ab0e0bf58bc49a392 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/groupassignments/item"
    i2d5a54257a3a154494e70602191b25e474f0eba5ac59afc7bbe7d9947af31558 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item"
    i868d4ff15f075e109c23674f69f3c9fe368b53fdc85733d5cad6217558ffb536 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item"
    ife64bae8f85f9d9ffa4f9b0e1f47986455765630e26bfebb51ad049410aab2c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/assignments/item"
)

type DeviceCustomAttributeShellScriptRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type DeviceCustomAttributeShellScriptRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *DeviceCustomAttributeShellScriptRequestBuilder) Assign()(*i5860a2467c9825730f7e5fe8b8587f837d06a6bce92dbb17df1197b8c3e63d0c.AssignRequestBuilder) {
    return i5860a2467c9825730f7e5fe8b8587f837d06a6bce92dbb17df1197b8c3e63d0c.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceCustomAttributeShellScriptRequestBuilder) Assignments()(*i9caa21903d914933c04aab2668b67aaf07ff8c88530bbc0f796266bff7b47bb2.AssignmentsRequestBuilder) {
    return i9caa21903d914933c04aab2668b67aaf07ff8c88530bbc0f796266bff7b47bb2.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceCustomAttributeShellScriptRequestBuilder) AssignmentsById(id string)(*ife64bae8f85f9d9ffa4f9b0e1f47986455765630e26bfebb51ad049410aab2c6.DeviceManagementScriptAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptAssignment_id"] = id
    }
    return ife64bae8f85f9d9ffa4f9b0e1f47986455765630e26bfebb51ad049410aab2c6.NewDeviceManagementScriptAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewDeviceCustomAttributeShellScriptRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceCustomAttributeShellScriptRequestBuilder) {
    m := &DeviceCustomAttributeShellScriptRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceManagement/deviceCustomAttributeShellScripts/{deviceCustomAttributeShellScript_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewDeviceCustomAttributeShellScriptRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceCustomAttributeShellScriptRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceCustomAttributeShellScriptRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DeviceCustomAttributeShellScriptRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DeviceCustomAttributeShellScriptRequestBuilder) CreateGetRequestInformation(q func (value *DeviceCustomAttributeShellScriptRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(DeviceCustomAttributeShellScriptRequestBuilderGetQueryParameters)
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
func (m *DeviceCustomAttributeShellScriptRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceCustomAttributeShellScript, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DeviceCustomAttributeShellScriptRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *DeviceCustomAttributeShellScriptRequestBuilder) DeviceRunStates()(*ib4c60324ec4c204995c1740d3c32783fe43242863d0216859b5414b0a63bd4a6.DeviceRunStatesRequestBuilder) {
    return ib4c60324ec4c204995c1740d3c32783fe43242863d0216859b5414b0a63bd4a6.NewDeviceRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceCustomAttributeShellScriptRequestBuilder) DeviceRunStatesById(id string)(*i868d4ff15f075e109c23674f69f3c9fe368b53fdc85733d5cad6217558ffb536.DeviceManagementScriptDeviceStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptDeviceState_id"] = id
    }
    return i868d4ff15f075e109c23674f69f3c9fe368b53fdc85733d5cad6217558ffb536.NewDeviceManagementScriptDeviceStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceCustomAttributeShellScriptRequestBuilder) Get(q func (value *DeviceCustomAttributeShellScriptRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceCustomAttributeShellScript, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceCustomAttributeShellScript() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceCustomAttributeShellScript), nil
}
func (m *DeviceCustomAttributeShellScriptRequestBuilder) GroupAssignments()(*i494c760a2f26845bb6a83e1609802c363f94efb574cc06272e4db09ffdd5950e.GroupAssignmentsRequestBuilder) {
    return i494c760a2f26845bb6a83e1609802c363f94efb574cc06272e4db09ffdd5950e.NewGroupAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceCustomAttributeShellScriptRequestBuilder) GroupAssignmentsById(id string)(*i076188fa8b5444fe1eee63d44522c5efaa3b78288e012e2ab0e0bf58bc49a392.DeviceManagementScriptGroupAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptGroupAssignment_id"] = id
    }
    return i076188fa8b5444fe1eee63d44522c5efaa3b78288e012e2ab0e0bf58bc49a392.NewDeviceManagementScriptGroupAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceCustomAttributeShellScriptRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceCustomAttributeShellScript, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *DeviceCustomAttributeShellScriptRequestBuilder) RunSummary()(*i367a45f04f002f42b81ce6533585aeeba4211ee6e776fb0ddb4cedf64aa10d6f.RunSummaryRequestBuilder) {
    return i367a45f04f002f42b81ce6533585aeeba4211ee6e776fb0ddb4cedf64aa10d6f.NewRunSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceCustomAttributeShellScriptRequestBuilder) UserRunStates()(*ia6cd2ea0f3d71e372a080d28d01095fd6857f4d1b6250e215940436d6784cbba.UserRunStatesRequestBuilder) {
    return ia6cd2ea0f3d71e372a080d28d01095fd6857f4d1b6250e215940436d6784cbba.NewUserRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceCustomAttributeShellScriptRequestBuilder) UserRunStatesById(id string)(*i2d5a54257a3a154494e70602191b25e474f0eba5ac59afc7bbe7d9947af31558.DeviceManagementScriptUserStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptUserState_id"] = id
    }
    return i2d5a54257a3a154494e70602191b25e474f0eba5ac59afc7bbe7d9947af31558.NewDeviceManagementScriptUserStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
