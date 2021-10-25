package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i233cfcd85fda77bc7b8ed84ae3eab86225b54c43cf6b9995aec4e929ecae0d5f "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroleassignmentrequests/item/roledefinition"
    i3c09d8808940e57cffe2a4b43ec6c03afed85c4d1b8c2a6b6f69170f1707c06e "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroleassignmentrequests/item/cancel"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i781dea20c6497af5c00396ab4ab2df03f2c4e8031835bbd6dffd7f1eb7bc6bbd "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroleassignmentrequests/item/updaterequest"
    i7a382583186c7a5d8b72c51d185f4589248347a161c9a5a7a7d8d885ab571893 "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroleassignmentrequests/item/resource"
    ia72138d764aac9bd3a352e3057f1578dd8230b4045684066c5c0d36a2be3badb "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroleassignmentrequests/item/subject"
)

type GovernanceRoleAssignmentRequestRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type GovernanceRoleAssignmentRequestRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *GovernanceRoleAssignmentRequestRequestBuilder) Cancel()(*i3c09d8808940e57cffe2a4b43ec6c03afed85c4d1b8c2a6b6f69170f1707c06e.CancelRequestBuilder) {
    return i3c09d8808940e57cffe2a4b43ec6c03afed85c4d1b8c2a6b6f69170f1707c06e.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewGovernanceRoleAssignmentRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GovernanceRoleAssignmentRequestRequestBuilder) {
    m := &GovernanceRoleAssignmentRequestRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/governanceRoleAssignmentRequests/{governanceRoleAssignmentRequest_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewGovernanceRoleAssignmentRequestRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GovernanceRoleAssignmentRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGovernanceRoleAssignmentRequestRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *GovernanceRoleAssignmentRequestRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *GovernanceRoleAssignmentRequestRequestBuilder) CreateGetRequestInformation(q func (value *GovernanceRoleAssignmentRequestRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(GovernanceRoleAssignmentRequestRequestBuilderGetQueryParameters)
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
func (m *GovernanceRoleAssignmentRequestRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceRoleAssignmentRequest, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *GovernanceRoleAssignmentRequestRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *GovernanceRoleAssignmentRequestRequestBuilder) Get(q func (value *GovernanceRoleAssignmentRequestRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceRoleAssignmentRequest, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewGovernanceRoleAssignmentRequest() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceRoleAssignmentRequest), nil
}
func (m *GovernanceRoleAssignmentRequestRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceRoleAssignmentRequest, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *GovernanceRoleAssignmentRequestRequestBuilder) Resource()(*i7a382583186c7a5d8b72c51d185f4589248347a161c9a5a7a7d8d885ab571893.ResourceRequestBuilder) {
    return i7a382583186c7a5d8b72c51d185f4589248347a161c9a5a7a7d8d885ab571893.NewResourceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GovernanceRoleAssignmentRequestRequestBuilder) RoleDefinition()(*i233cfcd85fda77bc7b8ed84ae3eab86225b54c43cf6b9995aec4e929ecae0d5f.RoleDefinitionRequestBuilder) {
    return i233cfcd85fda77bc7b8ed84ae3eab86225b54c43cf6b9995aec4e929ecae0d5f.NewRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GovernanceRoleAssignmentRequestRequestBuilder) Subject()(*ia72138d764aac9bd3a352e3057f1578dd8230b4045684066c5c0d36a2be3badb.SubjectRequestBuilder) {
    return ia72138d764aac9bd3a352e3057f1578dd8230b4045684066c5c0d36a2be3badb.NewSubjectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GovernanceRoleAssignmentRequestRequestBuilder) UpdateRequest()(*i781dea20c6497af5c00396ab4ab2df03f2c4e8031835bbd6dffd7f1eb7bc6bbd.UpdateRequestRequestBuilder) {
    return i781dea20c6497af5c00396ab4ab2df03f2c4e8031835bbd6dffd7f1eb7bc6bbd.NewUpdateRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
