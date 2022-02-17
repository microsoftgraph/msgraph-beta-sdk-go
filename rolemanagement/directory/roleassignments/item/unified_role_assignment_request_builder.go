package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1522474c74e47117b7b5d32f30b9ee09e9a4497d181ef48a790b5dd525353391 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignments/item/directoryscope"
    i3474a2e38b281f4e7797434d422c48f9318be118347d877da4600124658ab447 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignments/item/appscope"
    idbde65951a67811bb15ce6245951166d482b6787a2b87ced4038bae710df35af "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignments/item/principal"
    ife4fc797bde6a1f61ef8a96d5a81eea6e837ac644934fadf77ca8be443553675 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignments/item/roledefinition"
)

// UnifiedRoleAssignmentRequestBuilder builds and executes requests for operations under \roleManagement\directory\roleAssignments\{unifiedRoleAssignment-id}
type UnifiedRoleAssignmentRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UnifiedRoleAssignmentRequestBuilderDeleteOptions options for Delete
type UnifiedRoleAssignmentRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UnifiedRoleAssignmentRequestBuilderGetOptions options for Get
type UnifiedRoleAssignmentRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UnifiedRoleAssignmentRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UnifiedRoleAssignmentRequestBuilderGetQueryParameters resource to grant access to users or groups.
type UnifiedRoleAssignmentRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UnifiedRoleAssignmentRequestBuilderPatchOptions options for Patch
type UnifiedRoleAssignmentRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignment;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *UnifiedRoleAssignmentRequestBuilder) AppScope()(*i3474a2e38b281f4e7797434d422c48f9318be118347d877da4600124658ab447.AppScopeRequestBuilder) {
    return i3474a2e38b281f4e7797434d422c48f9318be118347d877da4600124658ab447.NewAppScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUnifiedRoleAssignmentRequestBuilderInternal instantiates a new UnifiedRoleAssignmentRequestBuilder and sets the default values.
func NewUnifiedRoleAssignmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRoleAssignmentRequestBuilder) {
    m := &UnifiedRoleAssignmentRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/directory/roleAssignments/{unifiedRoleAssignment_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUnifiedRoleAssignmentRequestBuilder instantiates a new UnifiedRoleAssignmentRequestBuilder and sets the default values.
func NewUnifiedRoleAssignmentRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRoleAssignmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUnifiedRoleAssignmentRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation resource to grant access to users or groups.
func (m *UnifiedRoleAssignmentRequestBuilder) CreateDeleteRequestInformation(options *UnifiedRoleAssignmentRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation resource to grant access to users or groups.
func (m *UnifiedRoleAssignmentRequestBuilder) CreateGetRequestInformation(options *UnifiedRoleAssignmentRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation resource to grant access to users or groups.
func (m *UnifiedRoleAssignmentRequestBuilder) CreatePatchRequestInformation(options *UnifiedRoleAssignmentRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete resource to grant access to users or groups.
func (m *UnifiedRoleAssignmentRequestBuilder) Delete(options *UnifiedRoleAssignmentRequestBuilderDeleteOptions)(error) {
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
func (m *UnifiedRoleAssignmentRequestBuilder) DirectoryScope()(*i1522474c74e47117b7b5d32f30b9ee09e9a4497d181ef48a790b5dd525353391.DirectoryScopeRequestBuilder) {
    return i1522474c74e47117b7b5d32f30b9ee09e9a4497d181ef48a790b5dd525353391.NewDirectoryScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get resource to grant access to users or groups.
func (m *UnifiedRoleAssignmentRequestBuilder) Get(options *UnifiedRoleAssignmentRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignment, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUnifiedRoleAssignment() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignment), nil
}
// Patch resource to grant access to users or groups.
func (m *UnifiedRoleAssignmentRequestBuilder) Patch(options *UnifiedRoleAssignmentRequestBuilderPatchOptions)(error) {
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
func (m *UnifiedRoleAssignmentRequestBuilder) Principal()(*idbde65951a67811bb15ce6245951166d482b6787a2b87ced4038bae710df35af.PrincipalRequestBuilder) {
    return idbde65951a67811bb15ce6245951166d482b6787a2b87ced4038bae710df35af.NewPrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UnifiedRoleAssignmentRequestBuilder) RoleDefinition()(*ife4fc797bde6a1f61ef8a96d5a81eea6e837ac644934fadf77ca8be443553675.RoleDefinitionRequestBuilder) {
    return ife4fc797bde6a1f61ef8a96d5a81eea6e837ac644934fadf77ca8be443553675.NewRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
