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

// UnifiedRoleAssignmentScheduleRequestRequestBuilder builds and executes requests for operations under \roleManagement\directory\roleAssignmentScheduleRequests\{unifiedRoleAssignmentScheduleRequest-id}
type UnifiedRoleAssignmentScheduleRequestRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UnifiedRoleAssignmentScheduleRequestRequestBuilderDeleteOptions options for Delete
type UnifiedRoleAssignmentScheduleRequestRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UnifiedRoleAssignmentScheduleRequestRequestBuilderGetOptions options for Get
type UnifiedRoleAssignmentScheduleRequestRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UnifiedRoleAssignmentScheduleRequestRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UnifiedRoleAssignmentScheduleRequestRequestBuilderGetQueryParameters get roleAssignmentScheduleRequests from roleManagement
type UnifiedRoleAssignmentScheduleRequestRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UnifiedRoleAssignmentScheduleRequestRequestBuilderPatchOptions options for Patch
type UnifiedRoleAssignmentScheduleRequestRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentScheduleRequest;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
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
// NewUnifiedRoleAssignmentScheduleRequestRequestBuilderInternal instantiates a new UnifiedRoleAssignmentScheduleRequestRequestBuilder and sets the default values.
func NewUnifiedRoleAssignmentScheduleRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRoleAssignmentScheduleRequestRequestBuilder) {
    m := &UnifiedRoleAssignmentScheduleRequestRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/directory/roleAssignmentScheduleRequests/{unifiedRoleAssignmentScheduleRequest_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUnifiedRoleAssignmentScheduleRequestRequestBuilder instantiates a new UnifiedRoleAssignmentScheduleRequestRequestBuilder and sets the default values.
func NewUnifiedRoleAssignmentScheduleRequestRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRoleAssignmentScheduleRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUnifiedRoleAssignmentScheduleRequestRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property roleAssignmentScheduleRequests for roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) CreateDeleteRequestInformation(options *UnifiedRoleAssignmentScheduleRequestRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) CreateGetRequestInformation(options *UnifiedRoleAssignmentScheduleRequestRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) CreatePatchRequestInformation(options *UnifiedRoleAssignmentScheduleRequestRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) Delete(options *UnifiedRoleAssignmentScheduleRequestRequestBuilderDeleteOptions)(error) {
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
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) DirectoryScope()(*ife2e638b4789000c032141a5d5b523ac665e0a27d645bcbc23e304da45bb8e1b.DirectoryScopeRequestBuilder) {
    return ife2e638b4789000c032141a5d5b523ac665e0a27d645bcbc23e304da45bb8e1b.NewDirectoryScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get roleAssignmentScheduleRequests from roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) Get(options *UnifiedRoleAssignmentScheduleRequestRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentScheduleRequest, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUnifiedRoleAssignmentScheduleRequest() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentScheduleRequest), nil
}
// Patch update the navigation property roleAssignmentScheduleRequests in roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) Patch(options *UnifiedRoleAssignmentScheduleRequestRequestBuilderPatchOptions)(error) {
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
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) Principal()(*i9c5799826ae3170b29df9f4222bb9541adaa678ea6a00d5e8648787ab57d56a2.PrincipalRequestBuilder) {
    return i9c5799826ae3170b29df9f4222bb9541adaa678ea6a00d5e8648787ab57d56a2.NewPrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) RoleDefinition()(*i372cbf778edb590ece49802b48d0cad9a9d2816d413f41e919e303ddfb6fb205.RoleDefinitionRequestBuilder) {
    return i372cbf778edb590ece49802b48d0cad9a9d2816d413f41e919e303ddfb6fb205.NewRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UnifiedRoleAssignmentScheduleRequestRequestBuilder) TargetSchedule()(*idd132590acfc190330a9c73516f45785b1f34dca9e2ba200f14d70dea8c209d6.TargetScheduleRequestBuilder) {
    return idd132590acfc190330a9c73516f45785b1f34dca9e2ba200f14d70dea8c209d6.NewTargetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
