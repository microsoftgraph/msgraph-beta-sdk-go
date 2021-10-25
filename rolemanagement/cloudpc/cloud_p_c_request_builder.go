package cloudpc

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i21e54f9ee2d47268e20afd63152400c9b4c1f5fccc344e57fac33bd5bafc27e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/cloudpc/roleassignments"
    i31977c399a73e8dceaa0f68ce2147b8b69cdb6d448081c8df74e04127a6f5b7e "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/cloudpc/roledefinitions"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    id0c942d6747644ac6d906984e3b3d0046ec64b19ed0e042c591f0eda0c0f8957 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/cloudpc/resourcenamespaces"
    i5122da689a05212ccc762ebf1a452cf0621ee3e3ee2ed9c79b323f7d87ab82ba "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/cloudpc/roleassignments/item"
    i6596aa8a9eb296ce0e4c8af117114f5b675a99379d789079c98d15d2b4b41cb0 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/cloudpc/roledefinitions/item"
    if2fe6d33056a86a8c02bd09d388e1ad97378adc4d03f5d4afe21fd7135b39314 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/cloudpc/resourcenamespaces/item"
)

type CloudPCRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type CloudPCRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func NewCloudPCRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CloudPCRequestBuilder) {
    m := &CloudPCRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/roleManagement/cloudPC{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewCloudPCRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CloudPCRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudPCRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *CloudPCRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *CloudPCRequestBuilder) CreateGetRequestInformation(q func (value *CloudPCRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(CloudPCRequestBuilderGetQueryParameters)
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
func (m *CloudPCRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RbacApplicationMultiple, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *CloudPCRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *CloudPCRequestBuilder) Get(q func (value *CloudPCRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RbacApplicationMultiple, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewRbacApplicationMultiple() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RbacApplicationMultiple), nil
}
func (m *CloudPCRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RbacApplicationMultiple, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *CloudPCRequestBuilder) ResourceNamespaces()(*id0c942d6747644ac6d906984e3b3d0046ec64b19ed0e042c591f0eda0c0f8957.ResourceNamespacesRequestBuilder) {
    return id0c942d6747644ac6d906984e3b3d0046ec64b19ed0e042c591f0eda0c0f8957.NewResourceNamespacesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CloudPCRequestBuilder) ResourceNamespacesById(id string)(*if2fe6d33056a86a8c02bd09d388e1ad97378adc4d03f5d4afe21fd7135b39314.UnifiedRbacResourceNamespaceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRbacResourceNamespace_id"] = id
    }
    return if2fe6d33056a86a8c02bd09d388e1ad97378adc4d03f5d4afe21fd7135b39314.NewUnifiedRbacResourceNamespaceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CloudPCRequestBuilder) RoleAssignments()(*i21e54f9ee2d47268e20afd63152400c9b4c1f5fccc344e57fac33bd5bafc27e9.RoleAssignmentsRequestBuilder) {
    return i21e54f9ee2d47268e20afd63152400c9b4c1f5fccc344e57fac33bd5bafc27e9.NewRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CloudPCRequestBuilder) RoleAssignmentsById(id string)(*i5122da689a05212ccc762ebf1a452cf0621ee3e3ee2ed9c79b323f7d87ab82ba.UnifiedRoleAssignmentMultipleRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignmentMultiple_id"] = id
    }
    return i5122da689a05212ccc762ebf1a452cf0621ee3e3ee2ed9c79b323f7d87ab82ba.NewUnifiedRoleAssignmentMultipleRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CloudPCRequestBuilder) RoleDefinitions()(*i31977c399a73e8dceaa0f68ce2147b8b69cdb6d448081c8df74e04127a6f5b7e.RoleDefinitionsRequestBuilder) {
    return i31977c399a73e8dceaa0f68ce2147b8b69cdb6d448081c8df74e04127a6f5b7e.NewRoleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CloudPCRequestBuilder) RoleDefinitionsById(id string)(*i6596aa8a9eb296ce0e4c8af117114f5b675a99379d789079c98d15d2b4b41cb0.UnifiedRoleDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleDefinition_id"] = id
    }
    return i6596aa8a9eb296ce0e4c8af117114f5b675a99379d789079c98d15d2b4b41cb0.NewUnifiedRoleDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
