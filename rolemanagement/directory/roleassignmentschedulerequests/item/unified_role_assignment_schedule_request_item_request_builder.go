package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i372cbf778edb590ece49802b48d0cad9a9d2816d413f41e919e303ddfb6fb205 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignmentschedulerequests/item/roledefinition"
    i4ea6311926981c1b8d508a0070714d2147904536e152e876f02accc7f84a4f1d "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignmentschedulerequests/item/cancel"
    i54a5814c7827f67076072df5cc4ac9db8de62466b2f4e1b5d6a168149227146a "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignmentschedulerequests/item/appscope"
    i72d227126694ade555f595b177e45004fe908d5b68f8c78d76c71a0232cc9c84 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignmentschedulerequests/item/activatedusing"
    i9c5799826ae3170b29df9f4222bb9541adaa678ea6a00d5e8648787ab57d56a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignmentschedulerequests/item/principal"
    idd132590acfc190330a9c73516f45785b1f34dca9e2ba200f14d70dea8c209d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignmentschedulerequests/item/targetschedule"
    ife2e638b4789000c032141a5d5b523ac665e0a27d645bcbc23e304da45bb8e1b "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignmentschedulerequests/item/directoryscope"
)

// UnifiedRoleAssignmentScheduleRequestItemRequestBuilder provides operations to manage the roleAssignmentScheduleRequests property of the microsoft.graph.rbacApplication entity.
type UnifiedRoleAssignmentScheduleRequestItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteOptions options for Delete
type UnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetOptions options for Get
type UnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetQueryParameters get roleAssignmentScheduleRequests from roleManagement
type UnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchOptions options for Patch
type UnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentScheduleRequestable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) ActivatedUsing()(*i72d227126694ade555f595b177e45004fe908d5b68f8c78d76c71a0232cc9c84.ActivatedUsingRequestBuilder) {
    return i72d227126694ade555f595b177e45004fe908d5b68f8c78d76c71a0232cc9c84.NewActivatedUsingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) AppScope()(*i54a5814c7827f67076072df5cc4ac9db8de62466b2f4e1b5d6a168149227146a.AppScopeRequestBuilder) {
    return i54a5814c7827f67076072df5cc4ac9db8de62466b2f4e1b5d6a168149227146a.NewAppScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Cancel()(*i4ea6311926981c1b8d508a0070714d2147904536e152e876f02accc7f84a4f1d.CancelRequestBuilder) {
    return i4ea6311926981c1b8d508a0070714d2147904536e152e876f02accc7f84a4f1d.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal instantiates a new UnifiedRoleAssignmentScheduleRequestItemRequestBuilder and sets the default values.
func NewUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) {
    m := &UnifiedRoleAssignmentScheduleRequestItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/directory/roleAssignmentScheduleRequests/{unifiedRoleAssignmentScheduleRequest_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUnifiedRoleAssignmentScheduleRequestItemRequestBuilder instantiates a new UnifiedRoleAssignmentScheduleRequestItemRequestBuilder and sets the default values.
func NewUnifiedRoleAssignmentScheduleRequestItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property roleAssignmentScheduleRequests for roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) CreateDeleteRequestInformation(options *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get roleAssignmentScheduleRequests from roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) CreateGetRequestInformation(options *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property roleAssignmentScheduleRequests in roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) CreatePatchRequestInformation(options *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property roleAssignmentScheduleRequests for roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Delete(options *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) DirectoryScope()(*ife2e638b4789000c032141a5d5b523ac665e0a27d645bcbc23e304da45bb8e1b.DirectoryScopeRequestBuilder) {
    return ife2e638b4789000c032141a5d5b523ac665e0a27d645bcbc23e304da45bb8e1b.NewDirectoryScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get roleAssignmentScheduleRequests from roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Get(options *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentScheduleRequestable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateUnifiedRoleAssignmentScheduleRequestFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentScheduleRequestable), nil
}
// Patch update the navigation property roleAssignmentScheduleRequests in roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Patch(options *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Principal()(*i9c5799826ae3170b29df9f4222bb9541adaa678ea6a00d5e8648787ab57d56a2.PrincipalRequestBuilder) {
    return i9c5799826ae3170b29df9f4222bb9541adaa678ea6a00d5e8648787ab57d56a2.NewPrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) RoleDefinition()(*i372cbf778edb590ece49802b48d0cad9a9d2816d413f41e919e303ddfb6fb205.RoleDefinitionRequestBuilder) {
    return i372cbf778edb590ece49802b48d0cad9a9d2816d413f41e919e303ddfb6fb205.NewRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) TargetSchedule()(*idd132590acfc190330a9c73516f45785b1f34dca9e2ba200f14d70dea8c209d6.TargetScheduleRequestBuilder) {
    return idd132590acfc190330a9c73516f45785b1f34dca9e2ba200f14d70dea8c209d6.NewTargetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
