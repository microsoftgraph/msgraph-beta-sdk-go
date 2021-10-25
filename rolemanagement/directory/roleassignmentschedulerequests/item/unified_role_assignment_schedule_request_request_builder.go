package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i372cbf778edb590ece49802b48d0cad9a9d2816d413f41e919e303ddfb6fb205 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignmentschedulerequests/item/roledefinition"
    i4ea6311926981c1b8d508a0070714d2147904536e152e876f02accc7f84a4f1d "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignmentschedulerequests/item/cancel"
    i54a5814c7827f67076072df5cc4ac9db8de62466b2f4e1b5d6a168149227146a "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignmentschedulerequests/item/appscope"
    i72d227126694ade555f595b177e45004fe908d5b68f8c78d76c71a0232cc9c84 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignmentschedulerequests/item/activatedusing"
    i9c5799826ae3170b29df9f4222bb9541adaa678ea6a00d5e8648787ab57d56a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignmentschedulerequests/item/principal"
    idd132590acfc190330a9c73516f45785b1f34dca9e2ba200f14d70dea8c209d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignmentschedulerequests/item/targetschedule"
    ife2e638b4789000c032141a5d5b523ac665e0a27d645bcbc23e304da45bb8e1b "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignmentschedulerequests/item/directoryscope"
)

type UnifiedRoleAssignmentScheduleRequestRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type UnifiedRoleAssignmentScheduleRequestRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) ActivatedUsing()(*i72d227126694ade555f595b177e45004fe908d5b68f8c78d76c71a0232cc9c84.ActivatedUsingRequestBuilder) {
    return i72d227126694ade555f595b177e45004fe908d5b68f8c78d76c71a0232cc9c84.NewActivatedUsingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) AppScope()(*i54a5814c7827f67076072df5cc4ac9db8de62466b2f4e1b5d6a168149227146a.AppScopeRequestBuilder) {
    return i54a5814c7827f67076072df5cc4ac9db8de62466b2f4e1b5d6a168149227146a.NewAppScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) Cancel()(*i4ea6311926981c1b8d508a0070714d2147904536e152e876f02accc7f84a4f1d.CancelRequestBuilder) {
    return i4ea6311926981c1b8d508a0070714d2147904536e152e876f02accc7f84a4f1d.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewUnifiedRoleAssignmentScheduleRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRoleAssignmentScheduleRequestRequestBuilder) {
    m := &UnifiedRoleAssignmentScheduleRequestRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/roleManagement/directory/roleAssignmentScheduleRequests/{unifiedRoleAssignmentScheduleRequest_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewUnifiedRoleAssignmentScheduleRequestRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRoleAssignmentScheduleRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUnifiedRoleAssignmentScheduleRequestRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) CreateGetRequestInformation(q func (value *UnifiedRoleAssignmentScheduleRequestRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(UnifiedRoleAssignmentScheduleRequestRequestBuilderGetQueryParameters)
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
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentScheduleRequest, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) DirectoryScope()(*ife2e638b4789000c032141a5d5b523ac665e0a27d645bcbc23e304da45bb8e1b.DirectoryScopeRequestBuilder) {
    return ife2e638b4789000c032141a5d5b523ac665e0a27d645bcbc23e304da45bb8e1b.NewDirectoryScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) Get(q func (value *UnifiedRoleAssignmentScheduleRequestRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentScheduleRequest, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUnifiedRoleAssignmentScheduleRequest() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentScheduleRequest), nil
}
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentScheduleRequest, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) Principal()(*i9c5799826ae3170b29df9f4222bb9541adaa678ea6a00d5e8648787ab57d56a2.PrincipalRequestBuilder) {
    return i9c5799826ae3170b29df9f4222bb9541adaa678ea6a00d5e8648787ab57d56a2.NewPrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) RoleDefinition()(*i372cbf778edb590ece49802b48d0cad9a9d2816d413f41e919e303ddfb6fb205.RoleDefinitionRequestBuilder) {
    return i372cbf778edb590ece49802b48d0cad9a9d2816d413f41e919e303ddfb6fb205.NewRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) TargetSchedule()(*idd132590acfc190330a9c73516f45785b1f34dca9e2ba200f14d70dea8c209d6.TargetScheduleRequestBuilder) {
    return idd132590acfc190330a9c73516f45785b1f34dca9e2ba200f14d70dea8c209d6.NewTargetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
