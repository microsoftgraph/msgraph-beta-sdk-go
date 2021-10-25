package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i00b6268726a46fcf8e15e0ca3c5f8889dfdd53aaf48fe4ce2411243820496cf2 "github.com/microsoftgraph/msgraph-beta-sdk-go/approvalworkflowproviders/item/businessflows"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i30206120795d8fda84ba9e547a109b67da276500cec13099585adedbf956e532 "github.com/microsoftgraph/msgraph-beta-sdk-go/approvalworkflowproviders/item/businessflowswithrequestsawaitingmydecision"
    i31436781316332b768fad45035d32f5befd814379c0ccbd0c58fb23e0e08e107 "github.com/microsoftgraph/msgraph-beta-sdk-go/approvalworkflowproviders/item/policytemplates"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    iad8cec4aa318577e66fff21478bdf279aabc1597d98a1031afe46ae7262c9daa "github.com/microsoftgraph/msgraph-beta-sdk-go/approvalworkflowproviders/item/policytemplates/item"
    icdd8feb64f3d3f11cf1c42a81a0ec83a900ebf29181eae132e78625db81a1fa7 "github.com/microsoftgraph/msgraph-beta-sdk-go/approvalworkflowproviders/item/businessflows/item"
    id3e3e89c17dc1203ba2dba0acce89454c2eacc6063d8b15b83854c381c75fc28 "github.com/microsoftgraph/msgraph-beta-sdk-go/approvalworkflowproviders/item/businessflowswithrequestsawaitingmydecision/item"
)

type ApprovalWorkflowProviderRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ApprovalWorkflowProviderRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *ApprovalWorkflowProviderRequestBuilder) BusinessFlows()(*i00b6268726a46fcf8e15e0ca3c5f8889dfdd53aaf48fe4ce2411243820496cf2.BusinessFlowsRequestBuilder) {
    return i00b6268726a46fcf8e15e0ca3c5f8889dfdd53aaf48fe4ce2411243820496cf2.NewBusinessFlowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApprovalWorkflowProviderRequestBuilder) BusinessFlowsById(id string)(*icdd8feb64f3d3f11cf1c42a81a0ec83a900ebf29181eae132e78625db81a1fa7.BusinessFlowRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["businessFlow_id"] = id
    }
    return icdd8feb64f3d3f11cf1c42a81a0ec83a900ebf29181eae132e78625db81a1fa7.NewBusinessFlowRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ApprovalWorkflowProviderRequestBuilder) BusinessFlowsWithRequestsAwaitingMyDecision()(*i30206120795d8fda84ba9e547a109b67da276500cec13099585adedbf956e532.BusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder) {
    return i30206120795d8fda84ba9e547a109b67da276500cec13099585adedbf956e532.NewBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApprovalWorkflowProviderRequestBuilder) BusinessFlowsWithRequestsAwaitingMyDecisionById(id string)(*id3e3e89c17dc1203ba2dba0acce89454c2eacc6063d8b15b83854c381c75fc28.BusinessFlowRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["businessFlow_id"] = id
    }
    return id3e3e89c17dc1203ba2dba0acce89454c2eacc6063d8b15b83854c381c75fc28.NewBusinessFlowRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewApprovalWorkflowProviderRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ApprovalWorkflowProviderRequestBuilder) {
    m := &ApprovalWorkflowProviderRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/approvalWorkflowProviders/{approvalWorkflowProvider_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewApprovalWorkflowProviderRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ApprovalWorkflowProviderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApprovalWorkflowProviderRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ApprovalWorkflowProviderRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ApprovalWorkflowProviderRequestBuilder) CreateGetRequestInformation(q func (value *ApprovalWorkflowProviderRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ApprovalWorkflowProviderRequestBuilderGetQueryParameters)
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
func (m *ApprovalWorkflowProviderRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ApprovalWorkflowProvider, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ApprovalWorkflowProviderRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ApprovalWorkflowProviderRequestBuilder) Get(q func (value *ApprovalWorkflowProviderRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ApprovalWorkflowProvider, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewApprovalWorkflowProvider() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ApprovalWorkflowProvider), nil
}
func (m *ApprovalWorkflowProviderRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ApprovalWorkflowProvider, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ApprovalWorkflowProviderRequestBuilder) PolicyTemplates()(*i31436781316332b768fad45035d32f5befd814379c0ccbd0c58fb23e0e08e107.PolicyTemplatesRequestBuilder) {
    return i31436781316332b768fad45035d32f5befd814379c0ccbd0c58fb23e0e08e107.NewPolicyTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApprovalWorkflowProviderRequestBuilder) PolicyTemplatesById(id string)(*iad8cec4aa318577e66fff21478bdf279aabc1597d98a1031afe46ae7262c9daa.GovernancePolicyTemplateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governancePolicyTemplate_id"] = id
    }
    return iad8cec4aa318577e66fff21478bdf279aabc1597d98a1031afe46ae7262c9daa.NewGovernancePolicyTemplateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
