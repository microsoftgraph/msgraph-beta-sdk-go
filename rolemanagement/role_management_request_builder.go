package rolemanagement

import (
    i10372b71511cc47bf754e90b050fc97fee639d10b03f642f5ac8321f1d6eee9e "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement"
    i45ee08ad7f18888228aa301bf50b8ba0f15edd802a1aab237f60b36937784346 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement"
    i8a93901c4a3aed17badfa3c9c54168d3e7506fa6cc5ed2d60a7be06761c00759 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ie179728a78650b7ea21a75747c7c0c71a4b3eeffeb7edcfa3f337c75f028ffa7 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/cloudpc"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type RoleManagementRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type RoleManagementRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *RoleManagementRequestBuilder) CloudPC()(*ie179728a78650b7ea21a75747c7c0c71a4b3eeffeb7edcfa3f337c75f028ffa7.CloudPCRequestBuilder) {
    return ie179728a78650b7ea21a75747c7c0c71a4b3eeffeb7edcfa3f337c75f028ffa7.NewCloudPCRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewRoleManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RoleManagementRequestBuilder) {
    m := &RoleManagementRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/roleManagement{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewRoleManagementRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RoleManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleManagementRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *RoleManagementRequestBuilder) CreateGetRequestInformation(q func (value *RoleManagementRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(RoleManagementRequestBuilderGetQueryParameters)
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
func (m *RoleManagementRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RoleManagement, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *RoleManagementRequestBuilder) DeviceManagement()(*i45ee08ad7f18888228aa301bf50b8ba0f15edd802a1aab237f60b36937784346.DeviceManagementRequestBuilder) {
    return i45ee08ad7f18888228aa301bf50b8ba0f15edd802a1aab237f60b36937784346.NewDeviceManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *RoleManagementRequestBuilder) Directory()(*i8a93901c4a3aed17badfa3c9c54168d3e7506fa6cc5ed2d60a7be06761c00759.DirectoryRequestBuilder) {
    return i8a93901c4a3aed17badfa3c9c54168d3e7506fa6cc5ed2d60a7be06761c00759.NewDirectoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *RoleManagementRequestBuilder) EntitlementManagement()(*i10372b71511cc47bf754e90b050fc97fee639d10b03f642f5ac8321f1d6eee9e.EntitlementManagementRequestBuilder) {
    return i10372b71511cc47bf754e90b050fc97fee639d10b03f642f5ac8321f1d6eee9e.NewEntitlementManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *RoleManagementRequestBuilder) Get(q func (value *RoleManagementRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RoleManagement, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewRoleManagement() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RoleManagement), nil
}
func (m *RoleManagementRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RoleManagement, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
