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

// ApprovalWorkflowProviderRequestBuilder builds and executes requests for operations under \approvalWorkflowProviders\{approvalWorkflowProvider-id}
type ApprovalWorkflowProviderRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ApprovalWorkflowProviderRequestBuilderDeleteOptions options for Delete
type ApprovalWorkflowProviderRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ApprovalWorkflowProviderRequestBuilderGetOptions options for Get
type ApprovalWorkflowProviderRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ApprovalWorkflowProviderRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ApprovalWorkflowProviderRequestBuilderGetQueryParameters get entity from approvalWorkflowProviders by key
type ApprovalWorkflowProviderRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ApprovalWorkflowProviderRequestBuilderPatchOptions options for Patch
type ApprovalWorkflowProviderRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ApprovalWorkflowProvider;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ApprovalWorkflowProviderRequestBuilder) BusinessFlows()(*i00b6268726a46fcf8e15e0ca3c5f8889dfdd53aaf48fe4ce2411243820496cf2.BusinessFlowsRequestBuilder) {
    return i00b6268726a46fcf8e15e0ca3c5f8889dfdd53aaf48fe4ce2411243820496cf2.NewBusinessFlowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BusinessFlowsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.approvalWorkflowProviders.item.businessFlows.item collection
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
// BusinessFlowsWithRequestsAwaitingMyDecisionById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.approvalWorkflowProviders.item.businessFlowsWithRequestsAwaitingMyDecision.item collection
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
// NewApprovalWorkflowProviderRequestBuilderInternal instantiates a new ApprovalWorkflowProviderRequestBuilder and sets the default values.
func NewApprovalWorkflowProviderRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ApprovalWorkflowProviderRequestBuilder) {
    m := &ApprovalWorkflowProviderRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/approvalWorkflowProviders/{approvalWorkflowProvider_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewApprovalWorkflowProviderRequestBuilder instantiates a new ApprovalWorkflowProviderRequestBuilder and sets the default values.
func NewApprovalWorkflowProviderRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ApprovalWorkflowProviderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApprovalWorkflowProviderRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from approvalWorkflowProviders
func (m *ApprovalWorkflowProviderRequestBuilder) CreateDeleteRequestInformation(options *ApprovalWorkflowProviderRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get entity from approvalWorkflowProviders by key
func (m *ApprovalWorkflowProviderRequestBuilder) CreateGetRequestInformation(options *ApprovalWorkflowProviderRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update entity in approvalWorkflowProviders
func (m *ApprovalWorkflowProviderRequestBuilder) CreatePatchRequestInformation(options *ApprovalWorkflowProviderRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete entity from approvalWorkflowProviders
func (m *ApprovalWorkflowProviderRequestBuilder) Delete(options *ApprovalWorkflowProviderRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get get entity from approvalWorkflowProviders by key
func (m *ApprovalWorkflowProviderRequestBuilder) Get(options *ApprovalWorkflowProviderRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ApprovalWorkflowProvider, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewApprovalWorkflowProvider() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ApprovalWorkflowProvider), nil
}
// Patch update entity in approvalWorkflowProviders
func (m *ApprovalWorkflowProviderRequestBuilder) Patch(options *ApprovalWorkflowProviderRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ApprovalWorkflowProviderRequestBuilder) PolicyTemplates()(*i31436781316332b768fad45035d32f5befd814379c0ccbd0c58fb23e0e08e107.PolicyTemplatesRequestBuilder) {
    return i31436781316332b768fad45035d32f5befd814379c0ccbd0c58fb23e0e08e107.NewPolicyTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PolicyTemplatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.approvalWorkflowProviders.item.policyTemplates.item collection
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
