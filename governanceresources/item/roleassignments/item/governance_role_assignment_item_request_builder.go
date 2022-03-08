package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i3abc3d9218e91a5fafadab38ade57dead754f7030e54ff7bc0cbad14ddab0a7b "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceresources/item/roleassignments/item/subject"
    icaba2f6b01e964573246ad4168789811e2c6315b16dfedbc695d1ef6560520a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceresources/item/roleassignments/item/linkedeligibleroleassignment"
    iecec41b98ea15ad4fe92aa166baac5bd2638a6c802ca82a30f2929034826fd0f "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceresources/item/roleassignments/item/resource"
    iee4dbea39593613c78ec24e53cdf8d2924237ab40a8e7e05550aae33b0bb3a16 "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceresources/item/roleassignments/item/roledefinition"
)

// GovernanceRoleAssignmentItemRequestBuilder provides operations to manage the roleAssignments property of the microsoft.graph.governanceResource entity.
type GovernanceRoleAssignmentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GovernanceRoleAssignmentItemRequestBuilderDeleteOptions options for Delete
type GovernanceRoleAssignmentItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GovernanceRoleAssignmentItemRequestBuilderGetOptions options for Get
type GovernanceRoleAssignmentItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *GovernanceRoleAssignmentItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GovernanceRoleAssignmentItemRequestBuilderGetQueryParameters the collection of role assignments for the resource.
type GovernanceRoleAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// GovernanceRoleAssignmentItemRequestBuilderPatchOptions options for Patch
type GovernanceRoleAssignmentItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceRoleAssignmentable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGovernanceRoleAssignmentItemRequestBuilderInternal instantiates a new GovernanceRoleAssignmentItemRequestBuilder and sets the default values.
func NewGovernanceRoleAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GovernanceRoleAssignmentItemRequestBuilder) {
    m := &GovernanceRoleAssignmentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/governanceResources/{governanceResource_id}/roleAssignments/{governanceRoleAssignment_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGovernanceRoleAssignmentItemRequestBuilder instantiates a new GovernanceRoleAssignmentItemRequestBuilder and sets the default values.
func NewGovernanceRoleAssignmentItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GovernanceRoleAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGovernanceRoleAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property roleAssignments for governanceResources
func (m *GovernanceRoleAssignmentItemRequestBuilder) CreateDeleteRequestInformation(options *GovernanceRoleAssignmentItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of role assignments for the resource.
func (m *GovernanceRoleAssignmentItemRequestBuilder) CreateGetRequestInformation(options *GovernanceRoleAssignmentItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property roleAssignments in governanceResources
func (m *GovernanceRoleAssignmentItemRequestBuilder) CreatePatchRequestInformation(options *GovernanceRoleAssignmentItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property roleAssignments for governanceResources
func (m *GovernanceRoleAssignmentItemRequestBuilder) Delete(options *GovernanceRoleAssignmentItemRequestBuilderDeleteOptions)(error) {
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
// Get the collection of role assignments for the resource.
func (m *GovernanceRoleAssignmentItemRequestBuilder) Get(options *GovernanceRoleAssignmentItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceRoleAssignmentable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateGovernanceRoleAssignmentFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceRoleAssignmentable), nil
}
func (m *GovernanceRoleAssignmentItemRequestBuilder) LinkedEligibleRoleAssignment()(*icaba2f6b01e964573246ad4168789811e2c6315b16dfedbc695d1ef6560520a3.LinkedEligibleRoleAssignmentRequestBuilder) {
    return icaba2f6b01e964573246ad4168789811e2c6315b16dfedbc695d1ef6560520a3.NewLinkedEligibleRoleAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property roleAssignments in governanceResources
func (m *GovernanceRoleAssignmentItemRequestBuilder) Patch(options *GovernanceRoleAssignmentItemRequestBuilderPatchOptions)(error) {
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
func (m *GovernanceRoleAssignmentItemRequestBuilder) Resource()(*iecec41b98ea15ad4fe92aa166baac5bd2638a6c802ca82a30f2929034826fd0f.ResourceRequestBuilder) {
    return iecec41b98ea15ad4fe92aa166baac5bd2638a6c802ca82a30f2929034826fd0f.NewResourceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GovernanceRoleAssignmentItemRequestBuilder) RoleDefinition()(*iee4dbea39593613c78ec24e53cdf8d2924237ab40a8e7e05550aae33b0bb3a16.RoleDefinitionRequestBuilder) {
    return iee4dbea39593613c78ec24e53cdf8d2924237ab40a8e7e05550aae33b0bb3a16.NewRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GovernanceRoleAssignmentItemRequestBuilder) Subject()(*i3abc3d9218e91a5fafadab38ade57dead754f7030e54ff7bc0cbad14ddab0a7b.SubjectRequestBuilder) {
    return i3abc3d9218e91a5fafadab38ade57dead754f7030e54ff7bc0cbad14ddab0a7b.NewSubjectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
