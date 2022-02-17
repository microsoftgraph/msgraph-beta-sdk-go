package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i58db9e4793251cf9ae32fa068af461b00d7ceae629ebdc2aa7054c7827578fb3 "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceresources/item/rolesettings"
    i8994915f6b921c39a8137a306b12e45f26896a0146d2c3163937128e4f0b5053 "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceresources/item/parent"
    iccb81886f13078be67906d67b607f4f7bed31ff7dbd649ca2ce2c93d6b838ee0 "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceresources/item/roleassignmentrequests"
    ie4dc11adf5ace6ad678d0f22990accf9d7e8cd8c8fe6080254ad1e674cff8621 "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceresources/item/roledefinitions"
    ie79a6e359646d3f119e90ecd4fe15c896747add74360c87f3375f7297d6b2061 "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceresources/item/roleassignments"
    i45fabc4315060a7a2958038a6d8acadc4ae3b133060b3d51a24591bef93cbd24 "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceresources/item/rolesettings/item"
    ib874dfaa0f52241f63218e4bd5551aae0a845a5ffdc8fe31ba0ba8d76ef68368 "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceresources/item/roleassignmentrequests/item"
    iebcf8addd43958225cc668f85c54db37378718447b94561d98640bf82caf0a25 "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceresources/item/roleassignments/item"
    if64c28382662f3e40b50d01a6d6cccdf239d73815a84e10346df6028d2331d87 "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceresources/item/roledefinitions/item"
)

// GovernanceResourceRequestBuilder builds and executes requests for operations under \governanceResources\{governanceResource-id}
type GovernanceResourceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GovernanceResourceRequestBuilderDeleteOptions options for Delete
type GovernanceResourceRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GovernanceResourceRequestBuilderGetOptions options for Get
type GovernanceResourceRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *GovernanceResourceRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GovernanceResourceRequestBuilderGetQueryParameters get entity from governanceResources by key
type GovernanceResourceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// GovernanceResourceRequestBuilderPatchOptions options for Patch
type GovernanceResourceRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceResource;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGovernanceResourceRequestBuilderInternal instantiates a new GovernanceResourceRequestBuilder and sets the default values.
func NewGovernanceResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GovernanceResourceRequestBuilder) {
    m := &GovernanceResourceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/governanceResources/{governanceResource_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGovernanceResourceRequestBuilder instantiates a new GovernanceResourceRequestBuilder and sets the default values.
func NewGovernanceResourceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GovernanceResourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGovernanceResourceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from governanceResources
func (m *GovernanceResourceRequestBuilder) CreateDeleteRequestInformation(options *GovernanceResourceRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from governanceResources by key
func (m *GovernanceResourceRequestBuilder) CreateGetRequestInformation(options *GovernanceResourceRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in governanceResources
func (m *GovernanceResourceRequestBuilder) CreatePatchRequestInformation(options *GovernanceResourceRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from governanceResources
func (m *GovernanceResourceRequestBuilder) Delete(options *GovernanceResourceRequestBuilderDeleteOptions)(error) {
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
// Get get entity from governanceResources by key
func (m *GovernanceResourceRequestBuilder) Get(options *GovernanceResourceRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceResource, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewGovernanceResource() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceResource), nil
}
func (m *GovernanceResourceRequestBuilder) Parent()(*i8994915f6b921c39a8137a306b12e45f26896a0146d2c3163937128e4f0b5053.ParentRequestBuilder) {
    return i8994915f6b921c39a8137a306b12e45f26896a0146d2c3163937128e4f0b5053.NewParentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in governanceResources
func (m *GovernanceResourceRequestBuilder) Patch(options *GovernanceResourceRequestBuilderPatchOptions)(error) {
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
func (m *GovernanceResourceRequestBuilder) RoleAssignmentRequests()(*iccb81886f13078be67906d67b607f4f7bed31ff7dbd649ca2ce2c93d6b838ee0.RoleAssignmentRequestsRequestBuilder) {
    return iccb81886f13078be67906d67b607f4f7bed31ff7dbd649ca2ce2c93d6b838ee0.NewRoleAssignmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.governanceResources.item.roleAssignmentRequests.item collection
func (m *GovernanceResourceRequestBuilder) RoleAssignmentRequestsById(id string)(*ib874dfaa0f52241f63218e4bd5551aae0a845a5ffdc8fe31ba0ba8d76ef68368.GovernanceRoleAssignmentRequestRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleAssignmentRequest_id"] = id
    }
    return ib874dfaa0f52241f63218e4bd5551aae0a845a5ffdc8fe31ba0ba8d76ef68368.NewGovernanceRoleAssignmentRequestRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GovernanceResourceRequestBuilder) RoleAssignments()(*ie79a6e359646d3f119e90ecd4fe15c896747add74360c87f3375f7297d6b2061.RoleAssignmentsRequestBuilder) {
    return ie79a6e359646d3f119e90ecd4fe15c896747add74360c87f3375f7297d6b2061.NewRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.governanceResources.item.roleAssignments.item collection
func (m *GovernanceResourceRequestBuilder) RoleAssignmentsById(id string)(*iebcf8addd43958225cc668f85c54db37378718447b94561d98640bf82caf0a25.GovernanceRoleAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleAssignment_id"] = id
    }
    return iebcf8addd43958225cc668f85c54db37378718447b94561d98640bf82caf0a25.NewGovernanceRoleAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GovernanceResourceRequestBuilder) RoleDefinitions()(*ie4dc11adf5ace6ad678d0f22990accf9d7e8cd8c8fe6080254ad1e674cff8621.RoleDefinitionsRequestBuilder) {
    return ie4dc11adf5ace6ad678d0f22990accf9d7e8cd8c8fe6080254ad1e674cff8621.NewRoleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinitionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.governanceResources.item.roleDefinitions.item collection
func (m *GovernanceResourceRequestBuilder) RoleDefinitionsById(id string)(*if64c28382662f3e40b50d01a6d6cccdf239d73815a84e10346df6028d2331d87.GovernanceRoleDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleDefinition_id"] = id
    }
    return if64c28382662f3e40b50d01a6d6cccdf239d73815a84e10346df6028d2331d87.NewGovernanceRoleDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *GovernanceResourceRequestBuilder) RoleSettings()(*i58db9e4793251cf9ae32fa068af461b00d7ceae629ebdc2aa7054c7827578fb3.RoleSettingsRequestBuilder) {
    return i58db9e4793251cf9ae32fa068af461b00d7ceae629ebdc2aa7054c7827578fb3.NewRoleSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleSettingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.governanceResources.item.roleSettings.item collection
func (m *GovernanceResourceRequestBuilder) RoleSettingsById(id string)(*i45fabc4315060a7a2958038a6d8acadc4ae3b133060b3d51a24591bef93cbd24.GovernanceRoleSettingRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleSetting_id"] = id
    }
    return i45fabc4315060a7a2958038a6d8acadc4ae3b133060b3d51a24591bef93cbd24.NewGovernanceRoleSettingRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
