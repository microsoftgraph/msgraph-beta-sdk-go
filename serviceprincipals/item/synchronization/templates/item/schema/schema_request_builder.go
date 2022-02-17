package schema

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i2875f990bc810a0e416397fd14a01a4725fa7f0a20108aba18cb3e18e820df4a "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/templates/item/schema/directories"
    i36fdc7569d3f0809e9eefdd0c8c16c5e0395c922bdaace8196665e44a09adb27 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/templates/item/schema/functions"
    i9303411a20178a126012bdf802c85819c1aa86d05b037afdb540ae9ec7518178 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/templates/item/schema/filteroperators"
    idd8cf83633c367895d6f418067b0f38663f134d6ecc5ab7c35bc0987d1b4ac9e "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/templates/item/schema/parseexpression"
    i7fa3917741c65999dc9b62767c91ce9eefb3d6af40f071ce0f23f64af751b350 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/templates/item/schema/directories/item"
)

// SchemaRequestBuilder builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}\synchronization\templates\{synchronizationTemplate-id}\schema
type SchemaRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SchemaRequestBuilderDeleteOptions options for Delete
type SchemaRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SchemaRequestBuilderGetOptions options for Get
type SchemaRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SchemaRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SchemaRequestBuilderGetQueryParameters default synchronization schema for the jobs based on this template.
type SchemaRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SchemaRequestBuilderPatchOptions options for Patch
type SchemaRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SynchronizationSchema;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewSchemaRequestBuilderInternal instantiates a new SchemaRequestBuilder and sets the default values.
func NewSchemaRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SchemaRequestBuilder) {
    m := &SchemaRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal_id}/synchronization/templates/{synchronizationTemplate_id}/schema{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSchemaRequestBuilder instantiates a new SchemaRequestBuilder and sets the default values.
func NewSchemaRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SchemaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchemaRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation default synchronization schema for the jobs based on this template.
func (m *SchemaRequestBuilder) CreateDeleteRequestInformation(options *SchemaRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation default synchronization schema for the jobs based on this template.
func (m *SchemaRequestBuilder) CreateGetRequestInformation(options *SchemaRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation default synchronization schema for the jobs based on this template.
func (m *SchemaRequestBuilder) CreatePatchRequestInformation(options *SchemaRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete default synchronization schema for the jobs based on this template.
func (m *SchemaRequestBuilder) Delete(options *SchemaRequestBuilderDeleteOptions)(error) {
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
func (m *SchemaRequestBuilder) Directories()(*i2875f990bc810a0e416397fd14a01a4725fa7f0a20108aba18cb3e18e820df4a.DirectoriesRequestBuilder) {
    return i2875f990bc810a0e416397fd14a01a4725fa7f0a20108aba18cb3e18e820df4a.NewDirectoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectoriesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.servicePrincipals.item.synchronization.templates.item.schema.directories.item collection
func (m *SchemaRequestBuilder) DirectoriesById(id string)(*i7fa3917741c65999dc9b62767c91ce9eefb3d6af40f071ce0f23f64af751b350.DirectoryDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryDefinition_id"] = id
    }
    return i7fa3917741c65999dc9b62767c91ce9eefb3d6af40f071ce0f23f64af751b350.NewDirectoryDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// FilterOperators builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}\synchronization\templates\{synchronizationTemplate-id}\schema\microsoft.graph.filterOperators()
func (m *SchemaRequestBuilder) FilterOperators()(*i9303411a20178a126012bdf802c85819c1aa86d05b037afdb540ae9ec7518178.FilterOperatorsRequestBuilder) {
    return i9303411a20178a126012bdf802c85819c1aa86d05b037afdb540ae9ec7518178.NewFilterOperatorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Functions builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}\synchronization\templates\{synchronizationTemplate-id}\schema\microsoft.graph.functions()
func (m *SchemaRequestBuilder) Functions()(*i36fdc7569d3f0809e9eefdd0c8c16c5e0395c922bdaace8196665e44a09adb27.FunctionsRequestBuilder) {
    return i36fdc7569d3f0809e9eefdd0c8c16c5e0395c922bdaace8196665e44a09adb27.NewFunctionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get default synchronization schema for the jobs based on this template.
func (m *SchemaRequestBuilder) Get(options *SchemaRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SynchronizationSchema, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSynchronizationSchema() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SynchronizationSchema), nil
}
func (m *SchemaRequestBuilder) ParseExpression()(*idd8cf83633c367895d6f418067b0f38663f134d6ecc5ab7c35bc0987d1b4ac9e.ParseExpressionRequestBuilder) {
    return idd8cf83633c367895d6f418067b0f38663f134d6ecc5ab7c35bc0987d1b4ac9e.NewParseExpressionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch default synchronization schema for the jobs based on this template.
func (m *SchemaRequestBuilder) Patch(options *SchemaRequestBuilderPatchOptions)(error) {
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
