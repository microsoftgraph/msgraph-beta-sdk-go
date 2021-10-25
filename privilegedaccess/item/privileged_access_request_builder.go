package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i13776d5c18dfe5cb7926395f0d7b87e53ddf81ba000513408fbd99e65db8c5fb "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/roleassignmentrequests"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i6517906c52d298f87f2e9d9ac7e9044b06b89fb46f521c3da9b9d1f087671c44 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/roleassignments"
    i83aba682e003297032480b3ecf836e1584c1518b5430625638ccefc15ac6bc8b "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/rolesettings"
    i9a0b2cd1cee2ff8b6db6a31960b35b297ea8a82ac1c9d7147b11db2225e14166 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/resources"
    ie0282f7e2e54b40fe4176464d126008c5f839e06d60480f482df4629ee34909c "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/roledefinitions"
    i055f2c9d21e7d744d266e629aae4492646c4497616eefa08e3bdc13deae816bc "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/roleassignmentrequests/item"
    i669c13d28c7fa4ec2373a44ba75602285da1019797e39c430f2ede5ee0850356 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/roleassignments/item"
    i93a22192c0fdf52ece01982a56d4c7f6efcce6d81281c6a74a3a9d90200baf9d "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/roledefinitions/item"
    id87892c36f8f2494356ab0b890f05c975ca897b7061b4af81b617506bdd996cf "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/rolesettings/item"
    ifef34655781055c0461c93a60eb94832b88d5e7127511a01d91517b0dbb2056f "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/resources/item"
)

type PrivilegedAccessRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type PrivilegedAccessRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func NewPrivilegedAccessRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrivilegedAccessRequestBuilder) {
    m := &PrivilegedAccessRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/privilegedAccess/{privilegedAccess_id}{?select,expand}";
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
func NewPrivilegedAccessRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrivilegedAccessRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedAccessRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *PrivilegedAccessRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *PrivilegedAccessRequestBuilder) CreateGetRequestInformation(q func (value *PrivilegedAccessRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(PrivilegedAccessRequestBuilderGetQueryParameters)
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
func (m *PrivilegedAccessRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedAccess, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *PrivilegedAccessRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *PrivilegedAccessRequestBuilder) Get(q func (value *PrivilegedAccessRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedAccess, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPrivilegedAccess() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedAccess), nil
}
func (m *PrivilegedAccessRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedAccess, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *PrivilegedAccessRequestBuilder) Resources()(*i9a0b2cd1cee2ff8b6db6a31960b35b297ea8a82ac1c9d7147b11db2225e14166.ResourcesRequestBuilder) {
    return i9a0b2cd1cee2ff8b6db6a31960b35b297ea8a82ac1c9d7147b11db2225e14166.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrivilegedAccessRequestBuilder) ResourcesById(id string)(*ifef34655781055c0461c93a60eb94832b88d5e7127511a01d91517b0dbb2056f.GovernanceResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["governanceResource_id"] = id
    }
    return ifef34655781055c0461c93a60eb94832b88d5e7127511a01d91517b0dbb2056f.NewGovernanceResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrivilegedAccessRequestBuilder) RoleAssignmentRequests()(*i13776d5c18dfe5cb7926395f0d7b87e53ddf81ba000513408fbd99e65db8c5fb.RoleAssignmentRequestsRequestBuilder) {
    return i13776d5c18dfe5cb7926395f0d7b87e53ddf81ba000513408fbd99e65db8c5fb.NewRoleAssignmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrivilegedAccessRequestBuilder) RoleAssignmentRequestsById(id string)(*i055f2c9d21e7d744d266e629aae4492646c4497616eefa08e3bdc13deae816bc.GovernanceRoleAssignmentRequestRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["governanceRoleAssignmentRequest_id"] = id
    }
    return i055f2c9d21e7d744d266e629aae4492646c4497616eefa08e3bdc13deae816bc.NewGovernanceRoleAssignmentRequestRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrivilegedAccessRequestBuilder) RoleAssignments()(*i6517906c52d298f87f2e9d9ac7e9044b06b89fb46f521c3da9b9d1f087671c44.RoleAssignmentsRequestBuilder) {
    return i6517906c52d298f87f2e9d9ac7e9044b06b89fb46f521c3da9b9d1f087671c44.NewRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrivilegedAccessRequestBuilder) RoleAssignmentsById(id string)(*i669c13d28c7fa4ec2373a44ba75602285da1019797e39c430f2ede5ee0850356.GovernanceRoleAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["governanceRoleAssignment_id"] = id
    }
    return i669c13d28c7fa4ec2373a44ba75602285da1019797e39c430f2ede5ee0850356.NewGovernanceRoleAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrivilegedAccessRequestBuilder) RoleDefinitions()(*ie0282f7e2e54b40fe4176464d126008c5f839e06d60480f482df4629ee34909c.RoleDefinitionsRequestBuilder) {
    return ie0282f7e2e54b40fe4176464d126008c5f839e06d60480f482df4629ee34909c.NewRoleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrivilegedAccessRequestBuilder) RoleDefinitionsById(id string)(*i93a22192c0fdf52ece01982a56d4c7f6efcce6d81281c6a74a3a9d90200baf9d.GovernanceRoleDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["governanceRoleDefinition_id"] = id
    }
    return i93a22192c0fdf52ece01982a56d4c7f6efcce6d81281c6a74a3a9d90200baf9d.NewGovernanceRoleDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrivilegedAccessRequestBuilder) RoleSettings()(*i83aba682e003297032480b3ecf836e1584c1518b5430625638ccefc15ac6bc8b.RoleSettingsRequestBuilder) {
    return i83aba682e003297032480b3ecf836e1584c1518b5430625638ccefc15ac6bc8b.NewRoleSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrivilegedAccessRequestBuilder) RoleSettingsById(id string)(*id87892c36f8f2494356ab0b890f05c975ca897b7061b4af81b617506bdd996cf.GovernanceRoleSettingRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["governanceRoleSetting_id"] = id
    }
    return id87892c36f8f2494356ab0b890f05c975ca897b7061b4af81b617506bdd996cf.NewGovernanceRoleSettingRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
