package nextversiondefinition

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i4d4e245af6721bc39b0ed5b42130efdfd5a1fc3654c3bd63cfbabf940a72ebb3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicydefinitions/item/previousversiondefinition/nextversiondefinition/definitionfile"
    i8c068a336e694856444efa8c9530bbfab994d44a29f10414147eea9838587f10 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicydefinitions/item/previousversiondefinition/nextversiondefinition/presentations"
    icd163ce32e48cec41a54cf51929022d155cff3b928a52ef12e2c4fce1bf256aa "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicydefinitions/item/previousversiondefinition/nextversiondefinition/category"
    i116532920a384b7a633bc588077a2000112ae5f7dfa6fb6cf05bdc16cb9610a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicydefinitions/item/previousversiondefinition/nextversiondefinition/presentations/item"
)

// NextVersionDefinitionRequestBuilder provides operations to manage the nextVersionDefinition property of the microsoft.graph.groupPolicyDefinition entity.
type NextVersionDefinitionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// NextVersionDefinitionRequestBuilderDeleteOptions options for Delete
type NextVersionDefinitionRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NextVersionDefinitionRequestBuilderGetOptions options for Get
type NextVersionDefinitionRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *NextVersionDefinitionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NextVersionDefinitionRequestBuilderGetQueryParameters definition of the next version of this definition
type NextVersionDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NextVersionDefinitionRequestBuilderPatchOptions options for Patch
type NextVersionDefinitionRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyDefinitionable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *NextVersionDefinitionRequestBuilder) Category()(*icd163ce32e48cec41a54cf51929022d155cff3b928a52ef12e2c4fce1bf256aa.CategoryRequestBuilder) {
    return icd163ce32e48cec41a54cf51929022d155cff3b928a52ef12e2c4fce1bf256aa.NewCategoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewNextVersionDefinitionRequestBuilderInternal instantiates a new NextVersionDefinitionRequestBuilder and sets the default values.
func NewNextVersionDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*NextVersionDefinitionRequestBuilder) {
    m := &NextVersionDefinitionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyDefinitions/{groupPolicyDefinition_id}/previousVersionDefinition/nextVersionDefinition{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewNextVersionDefinitionRequestBuilder instantiates a new NextVersionDefinitionRequestBuilder and sets the default values.
func NewNextVersionDefinitionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*NextVersionDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewNextVersionDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property nextVersionDefinition for deviceManagement
func (m *NextVersionDefinitionRequestBuilder) CreateDeleteRequestInformation(options *NextVersionDefinitionRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation definition of the next version of this definition
func (m *NextVersionDefinitionRequestBuilder) CreateGetRequestInformation(options *NextVersionDefinitionRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property nextVersionDefinition in deviceManagement
func (m *NextVersionDefinitionRequestBuilder) CreatePatchRequestInformation(options *NextVersionDefinitionRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *NextVersionDefinitionRequestBuilder) DefinitionFile()(*i4d4e245af6721bc39b0ed5b42130efdfd5a1fc3654c3bd63cfbabf940a72ebb3.DefinitionFileRequestBuilder) {
    return i4d4e245af6721bc39b0ed5b42130efdfd5a1fc3654c3bd63cfbabf940a72ebb3.NewDefinitionFileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property nextVersionDefinition for deviceManagement
func (m *NextVersionDefinitionRequestBuilder) Delete(options *NextVersionDefinitionRequestBuilderDeleteOptions)(error) {
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
// Get definition of the next version of this definition
func (m *NextVersionDefinitionRequestBuilder) Get(options *NextVersionDefinitionRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyDefinitionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateGroupPolicyDefinitionFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyDefinitionable), nil
}
// Patch update the navigation property nextVersionDefinition in deviceManagement
func (m *NextVersionDefinitionRequestBuilder) Patch(options *NextVersionDefinitionRequestBuilderPatchOptions)(error) {
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
func (m *NextVersionDefinitionRequestBuilder) Presentations()(*i8c068a336e694856444efa8c9530bbfab994d44a29f10414147eea9838587f10.PresentationsRequestBuilder) {
    return i8c068a336e694856444efa8c9530bbfab994d44a29f10414147eea9838587f10.NewPresentationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PresentationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.groupPolicyDefinitions.item.previousVersionDefinition.nextVersionDefinition.presentations.item collection
func (m *NextVersionDefinitionRequestBuilder) PresentationsById(id string)(*i116532920a384b7a633bc588077a2000112ae5f7dfa6fb6cf05bdc16cb9610a9.GroupPolicyPresentationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyPresentation_id"] = id
    }
    return i116532920a384b7a633bc588077a2000112ae5f7dfa6fb6cf05bdc16cb9610a9.NewGroupPolicyPresentationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
