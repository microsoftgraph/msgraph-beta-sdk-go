package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i077f9bb18a3b8a45bbe869035f3d5b79f48a4e8c97eeec6eee3ae05847933d3c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicycategories/item/definitions"
    i821765f6697e8edae076c3f190855ac0dea6daad75cfdcb0ceb7c24fc78d5252 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicycategories/item/parent"
    ib08bb2c1c958799079296de4709174f5e45ba516b4aba985d1cb9bc3f3277f8f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicycategories/item/definitionfile"
    id2be9225a42cb4e206ea5df101d0b19a8153d5b9b239d1b508a6a626f022511d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicycategories/item/children"
)

// Builds and executes requests for operations under \deviceManagement\groupPolicyCategories\{groupPolicyCategory-id}
type GroupPolicyCategoryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type GroupPolicyCategoryRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type GroupPolicyCategoryRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *GroupPolicyCategoryRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The available group policy categories for this account.
type GroupPolicyCategoryRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type GroupPolicyCategoryRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyCategory;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *GroupPolicyCategoryRequestBuilder) Children()(*id2be9225a42cb4e206ea5df101d0b19a8153d5b9b239d1b508a6a626f022511d.ChildrenRequestBuilder) {
    return id2be9225a42cb4e206ea5df101d0b19a8153d5b9b239d1b508a6a626f022511d.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new GroupPolicyCategoryRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewGroupPolicyCategoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupPolicyCategoryRequestBuilder) {
    m := &GroupPolicyCategoryRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceManagement/groupPolicyCategories/{groupPolicyCategory_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new GroupPolicyCategoryRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewGroupPolicyCategoryRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupPolicyCategoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyCategoryRequestBuilderInternal(urlParams, requestAdapter)
}
// The available group policy categories for this account.
// Parameters:
//  - options : Options for the request
func (m *GroupPolicyCategoryRequestBuilder) CreateDeleteRequestInformation(options *GroupPolicyCategoryRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The available group policy categories for this account.
// Parameters:
//  - options : Options for the request
func (m *GroupPolicyCategoryRequestBuilder) CreateGetRequestInformation(options *GroupPolicyCategoryRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// The available group policy categories for this account.
// Parameters:
//  - options : Options for the request
func (m *GroupPolicyCategoryRequestBuilder) CreatePatchRequestInformation(options *GroupPolicyCategoryRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *GroupPolicyCategoryRequestBuilder) DefinitionFile()(*ib08bb2c1c958799079296de4709174f5e45ba516b4aba985d1cb9bc3f3277f8f.DefinitionFileRequestBuilder) {
    return ib08bb2c1c958799079296de4709174f5e45ba516b4aba985d1cb9bc3f3277f8f.NewDefinitionFileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupPolicyCategoryRequestBuilder) Definitions()(*i077f9bb18a3b8a45bbe869035f3d5b79f48a4e8c97eeec6eee3ae05847933d3c.DefinitionsRequestBuilder) {
    return i077f9bb18a3b8a45bbe869035f3d5b79f48a4e8c97eeec6eee3ae05847933d3c.NewDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The available group policy categories for this account.
// Parameters:
//  - options : Options for the request
func (m *GroupPolicyCategoryRequestBuilder) Delete(options *GroupPolicyCategoryRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// The available group policy categories for this account.
// Parameters:
//  - options : Options for the request
func (m *GroupPolicyCategoryRequestBuilder) Get(options *GroupPolicyCategoryRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyCategory, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewGroupPolicyCategory() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyCategory), nil
}
func (m *GroupPolicyCategoryRequestBuilder) Parent()(*i821765f6697e8edae076c3f190855ac0dea6daad75cfdcb0ceb7c24fc78d5252.ParentRequestBuilder) {
    return i821765f6697e8edae076c3f190855ac0dea6daad75cfdcb0ceb7c24fc78d5252.NewParentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The available group policy categories for this account.
// Parameters:
//  - options : Options for the request
func (m *GroupPolicyCategoryRequestBuilder) Patch(options *GroupPolicyCategoryRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
