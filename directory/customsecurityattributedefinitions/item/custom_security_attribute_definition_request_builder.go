package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    iabb2c9c8d603422583bcb8f8258070d32341f88769a528bbf7389076f4dc84ab "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/customsecurityattributedefinitions/item/allowedvalues"
    i1fe7ff93f58439008a626f169f7eeff4e9a77d0871cc58d3033500a4d65a4a2c "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/customsecurityattributedefinitions/item/allowedvalues/item"
)

// CustomSecurityAttributeDefinitionRequestBuilder builds and executes requests for operations under \directory\customSecurityAttributeDefinitions\{customSecurityAttributeDefinition-id}
type CustomSecurityAttributeDefinitionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CustomSecurityAttributeDefinitionRequestBuilderDeleteOptions options for Delete
type CustomSecurityAttributeDefinitionRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CustomSecurityAttributeDefinitionRequestBuilderGetOptions options for Get
type CustomSecurityAttributeDefinitionRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CustomSecurityAttributeDefinitionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CustomSecurityAttributeDefinitionRequestBuilderGetQueryParameters schema of a custom security attributes (key-value pairs).
type CustomSecurityAttributeDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// CustomSecurityAttributeDefinitionRequestBuilderPatchOptions options for Patch
type CustomSecurityAttributeDefinitionRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CustomSecurityAttributeDefinition;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *CustomSecurityAttributeDefinitionRequestBuilder) AllowedValues()(*iabb2c9c8d603422583bcb8f8258070d32341f88769a528bbf7389076f4dc84ab.AllowedValuesRequestBuilder) {
    return iabb2c9c8d603422583bcb8f8258070d32341f88769a528bbf7389076f4dc84ab.NewAllowedValuesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AllowedValuesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.customSecurityAttributeDefinitions.item.allowedValues.item collection
func (m *CustomSecurityAttributeDefinitionRequestBuilder) AllowedValuesById(id string)(*i1fe7ff93f58439008a626f169f7eeff4e9a77d0871cc58d3033500a4d65a4a2c.AllowedValueRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["allowedValue_id"] = id
    }
    return i1fe7ff93f58439008a626f169f7eeff4e9a77d0871cc58d3033500a4d65a4a2c.NewAllowedValueRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewCustomSecurityAttributeDefinitionRequestBuilderInternal instantiates a new CustomSecurityAttributeDefinitionRequestBuilder and sets the default values.
func NewCustomSecurityAttributeDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CustomSecurityAttributeDefinitionRequestBuilder) {
    m := &CustomSecurityAttributeDefinitionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory/customSecurityAttributeDefinitions/{customSecurityAttributeDefinition_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCustomSecurityAttributeDefinitionRequestBuilder instantiates a new CustomSecurityAttributeDefinitionRequestBuilder and sets the default values.
func NewCustomSecurityAttributeDefinitionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CustomSecurityAttributeDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCustomSecurityAttributeDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation schema of a custom security attributes (key-value pairs).
func (m *CustomSecurityAttributeDefinitionRequestBuilder) CreateDeleteRequestInformation(options *CustomSecurityAttributeDefinitionRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation schema of a custom security attributes (key-value pairs).
func (m *CustomSecurityAttributeDefinitionRequestBuilder) CreateGetRequestInformation(options *CustomSecurityAttributeDefinitionRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation schema of a custom security attributes (key-value pairs).
func (m *CustomSecurityAttributeDefinitionRequestBuilder) CreatePatchRequestInformation(options *CustomSecurityAttributeDefinitionRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete schema of a custom security attributes (key-value pairs).
func (m *CustomSecurityAttributeDefinitionRequestBuilder) Delete(options *CustomSecurityAttributeDefinitionRequestBuilderDeleteOptions)(error) {
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
// Get schema of a custom security attributes (key-value pairs).
func (m *CustomSecurityAttributeDefinitionRequestBuilder) Get(options *CustomSecurityAttributeDefinitionRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CustomSecurityAttributeDefinition, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewCustomSecurityAttributeDefinition() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CustomSecurityAttributeDefinition), nil
}
// Patch schema of a custom security attributes (key-value pairs).
func (m *CustomSecurityAttributeDefinitionRequestBuilder) Patch(options *CustomSecurityAttributeDefinitionRequestBuilderPatchOptions)(error) {
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
