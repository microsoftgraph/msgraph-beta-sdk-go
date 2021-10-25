package virtualendpoint

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i1529c4b52ff4b5b9e3d21861baf5f70ff01e22c7ce170e59bf3ce66277506553 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/cloudpcs"
    i1696d9b0fb89b84debc07dd2077aff8e4f8ff2b028c14e5a9b3d5139c682a3f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/deviceimages"
    i23a63a02c83d28d72ce351404f7169106e8b8114b7d9d0ca4052db2a624e6273 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/usersettings"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5b2233e79dc3147871fe57df3812d50c932074cfced41065bb12f77e2edec21e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/auditevents"
    i641f590b75b18bc875786f11f8c01028ef7a86284da618ffc6636c5f70071b7c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/geteffectivepermissions"
    ic615011cfc05f9e79dc505191df45c19871dc563dfe317f14f5eb7e21000fc23 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/onpremisesconnections"
    if4100d044fd5ea8af7d00890548b33d7f96329a5a6d63f806c3f7705cafc3c7b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/provisioningpolicies"
    i471da8905b96f5687f46d1c5dc0de1420e0c35b35374f400b7618e2937b4bf34 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/onpremisesconnections/item"
    i57a7497f8ee90067ee6f4aa2a9b6074c547470907dd198f9de96f123682ddda0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/auditevents/item"
    i6c86691f21aa3210591c58f278b27fee3b5fb3a866b020940f913e3f9a20a060 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/provisioningpolicies/item"
    i940a1af81d34e1d5a5cfef93acb216fcbeba9c4d02700ce3000fd8f0588fd162 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/usersettings/item"
    ia3bcb5abfa0df15f63e2a6fa1f06bcbd43b05fc488a36754660b8c694246fe0b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/deviceimages/item"
    ic14d63ed61f1b60b506fcd81a146444ab96691619c2ac32f274b0793b622f23f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/cloudpcs/item"
)

type VirtualEndpointRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type VirtualEndpointRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *VirtualEndpointRequestBuilder) AuditEvents()(*i5b2233e79dc3147871fe57df3812d50c932074cfced41065bb12f77e2edec21e.AuditEventsRequestBuilder) {
    return i5b2233e79dc3147871fe57df3812d50c932074cfced41065bb12f77e2edec21e.NewAuditEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *VirtualEndpointRequestBuilder) AuditEventsById(id string)(*i57a7497f8ee90067ee6f4aa2a9b6074c547470907dd198f9de96f123682ddda0.CloudPcAuditEventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcAuditEvent_id"] = id
    }
    return i57a7497f8ee90067ee6f4aa2a9b6074c547470907dd198f9de96f123682ddda0.NewCloudPcAuditEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *VirtualEndpointRequestBuilder) CloudPCs()(*i1529c4b52ff4b5b9e3d21861baf5f70ff01e22c7ce170e59bf3ce66277506553.CloudPCsRequestBuilder) {
    return i1529c4b52ff4b5b9e3d21861baf5f70ff01e22c7ce170e59bf3ce66277506553.NewCloudPCsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *VirtualEndpointRequestBuilder) CloudPCsById(id string)(*ic14d63ed61f1b60b506fcd81a146444ab96691619c2ac32f274b0793b622f23f.CloudPCRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPC_id"] = id
    }
    return ic14d63ed61f1b60b506fcd81a146444ab96691619c2ac32f274b0793b622f23f.NewCloudPCRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewVirtualEndpointRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*VirtualEndpointRequestBuilder) {
    m := &VirtualEndpointRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceManagement/virtualEndpoint{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewVirtualEndpointRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*VirtualEndpointRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *VirtualEndpointRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *VirtualEndpointRequestBuilder) CreateGetRequestInformation(q func (value *VirtualEndpointRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(VirtualEndpointRequestBuilderGetQueryParameters)
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
func (m *VirtualEndpointRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.VirtualEndpoint, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *VirtualEndpointRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *VirtualEndpointRequestBuilder) DeviceImages()(*i1696d9b0fb89b84debc07dd2077aff8e4f8ff2b028c14e5a9b3d5139c682a3f3.DeviceImagesRequestBuilder) {
    return i1696d9b0fb89b84debc07dd2077aff8e4f8ff2b028c14e5a9b3d5139c682a3f3.NewDeviceImagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *VirtualEndpointRequestBuilder) DeviceImagesById(id string)(*ia3bcb5abfa0df15f63e2a6fa1f06bcbd43b05fc488a36754660b8c694246fe0b.CloudPcDeviceImageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcDeviceImage_id"] = id
    }
    return ia3bcb5abfa0df15f63e2a6fa1f06bcbd43b05fc488a36754660b8c694246fe0b.NewCloudPcDeviceImageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *VirtualEndpointRequestBuilder) Get(q func (value *VirtualEndpointRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.VirtualEndpoint, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewVirtualEndpoint() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.VirtualEndpoint), nil
}
func (m *VirtualEndpointRequestBuilder) GetEffectivePermissions()(*i641f590b75b18bc875786f11f8c01028ef7a86284da618ffc6636c5f70071b7c.GetEffectivePermissionsRequestBuilder) {
    return i641f590b75b18bc875786f11f8c01028ef7a86284da618ffc6636c5f70071b7c.NewGetEffectivePermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *VirtualEndpointRequestBuilder) OnPremisesConnections()(*ic615011cfc05f9e79dc505191df45c19871dc563dfe317f14f5eb7e21000fc23.OnPremisesConnectionsRequestBuilder) {
    return ic615011cfc05f9e79dc505191df45c19871dc563dfe317f14f5eb7e21000fc23.NewOnPremisesConnectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *VirtualEndpointRequestBuilder) OnPremisesConnectionsById(id string)(*i471da8905b96f5687f46d1c5dc0de1420e0c35b35374f400b7618e2937b4bf34.CloudPcOnPremisesConnectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcOnPremisesConnection_id"] = id
    }
    return i471da8905b96f5687f46d1c5dc0de1420e0c35b35374f400b7618e2937b4bf34.NewCloudPcOnPremisesConnectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *VirtualEndpointRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.VirtualEndpoint, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *VirtualEndpointRequestBuilder) ProvisioningPolicies()(*if4100d044fd5ea8af7d00890548b33d7f96329a5a6d63f806c3f7705cafc3c7b.ProvisioningPoliciesRequestBuilder) {
    return if4100d044fd5ea8af7d00890548b33d7f96329a5a6d63f806c3f7705cafc3c7b.NewProvisioningPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *VirtualEndpointRequestBuilder) ProvisioningPoliciesById(id string)(*i6c86691f21aa3210591c58f278b27fee3b5fb3a866b020940f913e3f9a20a060.CloudPcProvisioningPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcProvisioningPolicy_id"] = id
    }
    return i6c86691f21aa3210591c58f278b27fee3b5fb3a866b020940f913e3f9a20a060.NewCloudPcProvisioningPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *VirtualEndpointRequestBuilder) UserSettings()(*i23a63a02c83d28d72ce351404f7169106e8b8114b7d9d0ca4052db2a624e6273.UserSettingsRequestBuilder) {
    return i23a63a02c83d28d72ce351404f7169106e8b8114b7d9d0ca4052db2a624e6273.NewUserSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *VirtualEndpointRequestBuilder) UserSettingsById(id string)(*i940a1af81d34e1d5a5cfef93acb216fcbeba9c4d02700ce3000fd8f0588fd162.CloudPcUserSettingRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcUserSetting_id"] = id
    }
    return i940a1af81d34e1d5a5cfef93acb216fcbeba9c4d02700ce3000fd8f0588fd162.NewCloudPcUserSettingRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
