package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i15947f922198c0bab892f4a7c084ebe5f7cbb43f5394d31b937d5157fbe0edba "github.com/microsoftgraph/msgraph-beta-sdk-go/permissiongrants/item/checkmemberobjects"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i751dbbe15e72bbd1122914748b96ee1a20fca8119351121e2c491922693257d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/permissiongrants/item/getmembergroups"
    ibf0a4d81ee4b8a7e4e3ed9c9fb6d0882fdbab481f7b848c7db240c388dfbec2f "github.com/microsoftgraph/msgraph-beta-sdk-go/permissiongrants/item/getmemberobjects"
    ie01778f41a26871b3f51e401f57da7e144998794d08141d92ee03a581180728e "github.com/microsoftgraph/msgraph-beta-sdk-go/permissiongrants/item/checkmembergroups"
    if1160871f4538f1c2881002e619cb11a8c41ca9ef807c8c5f7a0f39c74706b6d "github.com/microsoftgraph/msgraph-beta-sdk-go/permissiongrants/item/restore"
)

type ResourceSpecificPermissionGrantRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ResourceSpecificPermissionGrantRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *ResourceSpecificPermissionGrantRequestBuilder) CheckMemberGroups()(*ie01778f41a26871b3f51e401f57da7e144998794d08141d92ee03a581180728e.CheckMemberGroupsRequestBuilder) {
    return ie01778f41a26871b3f51e401f57da7e144998794d08141d92ee03a581180728e.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceSpecificPermissionGrantRequestBuilder) CheckMemberObjects()(*i15947f922198c0bab892f4a7c084ebe5f7cbb43f5394d31b937d5157fbe0edba.CheckMemberObjectsRequestBuilder) {
    return i15947f922198c0bab892f4a7c084ebe5f7cbb43f5394d31b937d5157fbe0edba.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewResourceSpecificPermissionGrantRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ResourceSpecificPermissionGrantRequestBuilder) {
    m := &ResourceSpecificPermissionGrantRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/permissionGrants/{resourceSpecificPermissionGrant_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewResourceSpecificPermissionGrantRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ResourceSpecificPermissionGrantRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewResourceSpecificPermissionGrantRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ResourceSpecificPermissionGrantRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ResourceSpecificPermissionGrantRequestBuilder) CreateGetRequestInformation(q func (value *ResourceSpecificPermissionGrantRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ResourceSpecificPermissionGrantRequestBuilderGetQueryParameters)
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
func (m *ResourceSpecificPermissionGrantRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ResourceSpecificPermissionGrant, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ResourceSpecificPermissionGrantRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ResourceSpecificPermissionGrantRequestBuilder) Get(q func (value *ResourceSpecificPermissionGrantRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ResourceSpecificPermissionGrant, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewResourceSpecificPermissionGrant() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ResourceSpecificPermissionGrant), nil
}
func (m *ResourceSpecificPermissionGrantRequestBuilder) GetMemberGroups()(*i751dbbe15e72bbd1122914748b96ee1a20fca8119351121e2c491922693257d9.GetMemberGroupsRequestBuilder) {
    return i751dbbe15e72bbd1122914748b96ee1a20fca8119351121e2c491922693257d9.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceSpecificPermissionGrantRequestBuilder) GetMemberObjects()(*ibf0a4d81ee4b8a7e4e3ed9c9fb6d0882fdbab481f7b848c7db240c388dfbec2f.GetMemberObjectsRequestBuilder) {
    return ibf0a4d81ee4b8a7e4e3ed9c9fb6d0882fdbab481f7b848c7db240c388dfbec2f.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceSpecificPermissionGrantRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ResourceSpecificPermissionGrant, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ResourceSpecificPermissionGrantRequestBuilder) Restore()(*if1160871f4538f1c2881002e619cb11a8c41ca9ef807c8c5f7a0f39c74706b6d.RestoreRequestBuilder) {
    return if1160871f4538f1c2881002e619cb11a8c41ca9ef807c8c5f7a0f39c74706b6d.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
