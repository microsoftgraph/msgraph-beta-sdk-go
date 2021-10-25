package rolescopetags

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4edbe639b32662e436fc1a2df91f2c1f4f9b92866c08a7d5338db735db338052 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/roleassignments/item/rolescopetags/getrolescopetagsbyid"
    i59b421f7e62c770e57ae33b84588a1af6e28a2f7b0dacd4d0f5e6d5590f30f65 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/roleassignments/item/rolescopetags/hascustomrolescopetag"
    ie1c6af2bb765fdf2e7854c03f501167604144eaec4781e2148d60afb3225dd4f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/roleassignments/item/rolescopetags/ref"
)

type RoleScopeTagsRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type RoleScopeTagsRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Count *bool;
    Expand []string;
    Filter *string;
    Orderby []string;
    Search *string;
    Select_escpaped []string;
    Skip *int32;
    Top *int32;
}
func NewRoleScopeTagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RoleScopeTagsRequestBuilder) {
    m := &RoleScopeTagsRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceManagement/roleAssignments/{deviceAndAppManagementRoleAssignment_id}/roleScopeTags{?top,skip,search,filter,count,orderby,select,expand}";
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
func NewRoleScopeTagsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RoleScopeTagsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleScopeTagsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *RoleScopeTagsRequestBuilder) CreateGetRequestInformation(q func (value *RoleScopeTagsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(RoleScopeTagsRequestBuilderGetQueryParameters)
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
func (m *RoleScopeTagsRequestBuilder) Get(q func (value *RoleScopeTagsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*RoleScopeTagsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRoleScopeTagsResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*RoleScopeTagsResponse), nil
}
func (m *RoleScopeTagsRequestBuilder) GetRoleScopeTagsById()(*i4edbe639b32662e436fc1a2df91f2c1f4f9b92866c08a7d5338db735db338052.GetRoleScopeTagsByIdRequestBuilder) {
    return i4edbe639b32662e436fc1a2df91f2c1f4f9b92866c08a7d5338db735db338052.NewGetRoleScopeTagsByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *RoleScopeTagsRequestBuilder) HasCustomRoleScopeTag()(*i59b421f7e62c770e57ae33b84588a1af6e28a2f7b0dacd4d0f5e6d5590f30f65.HasCustomRoleScopeTagRequestBuilder) {
    return i59b421f7e62c770e57ae33b84588a1af6e28a2f7b0dacd4d0f5e6d5590f30f65.NewHasCustomRoleScopeTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *RoleScopeTagsRequestBuilder) Ref()(*ie1c6af2bb765fdf2e7854c03f501167604144eaec4781e2148d60afb3225dd4f.RefRequestBuilder) {
    return ie1c6af2bb765fdf2e7854c03f501167604144eaec4781e2148d60afb3225dd4f.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
