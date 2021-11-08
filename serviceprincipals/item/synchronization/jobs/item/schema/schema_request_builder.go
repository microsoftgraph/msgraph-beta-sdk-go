package schema

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0660671695b3f2242dd3113221c5e3f72342e0ff0ffa9d380502e223eb7c75f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs/item/schema/filteroperators"
    i2bb0f986999c75c0ca4bdcdc1c396d94ceed8cc193cfa0837b2802ca7f2c8fab "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs/item/schema/functions"
    if23bb949ea15062726d8be4224904c7bd42d0e0e3dd45612ea423af36b5ad737 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs/item/schema/parseexpression"
    if7fa57f25aa6ec684df0c748c3058c91a1b7076617c054b00c61cd67da946cfb "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs/item/schema/directories"
    i3b5a871a284736b01431c7efcb2c26c841b70a3fec1db50432ea66085f8a95dc "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs/item/schema/directories/item"
)

// Builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}\synchronization\jobs\{synchronizationJob-id}\schema
type SchemaRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type SchemaRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
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
// The synchronization schema configured for the job.
type SchemaRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
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
// Instantiates a new SchemaRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewSchemaRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SchemaRequestBuilder) {
    m := &SchemaRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal_id}/synchronization/jobs/{synchronizationJob_id}/schema{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new SchemaRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewSchemaRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SchemaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSchemaRequestBuilderInternal(urlParams, requestAdapter)
}
// The synchronization schema configured for the job.
// Parameters:
//  - options : Options for the request
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
// The synchronization schema configured for the job.
// Parameters:
//  - options : Options for the request
func (m *SchemaRequestBuilder) CreateGetRequestInformation(options *SchemaRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The synchronization schema configured for the job.
// Parameters:
//  - options : Options for the request
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
// The synchronization schema configured for the job.
// Parameters:
//  - options : Options for the request
func (m *SchemaRequestBuilder) Delete(options *SchemaRequestBuilderDeleteOptions)(error) {
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
func (m *SchemaRequestBuilder) Directories()(*if7fa57f25aa6ec684df0c748c3058c91a1b7076617c054b00c61cd67da946cfb.DirectoriesRequestBuilder) {
    return if7fa57f25aa6ec684df0c748c3058c91a1b7076617c054b00c61cd67da946cfb.NewDirectoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.servicePrincipals.item.synchronization.jobs.item.schema.directories.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *SchemaRequestBuilder) DirectoriesById(id string)(*i3b5a871a284736b01431c7efcb2c26c841b70a3fec1db50432ea66085f8a95dc.DirectoryDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryDefinition_id"] = id
    }
    return i3b5a871a284736b01431c7efcb2c26c841b70a3fec1db50432ea66085f8a95dc.NewDirectoryDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}\synchronization\jobs\{synchronizationJob-id}\schema\microsoft.graph.filterOperators()
func (m *SchemaRequestBuilder) FilterOperators()(*i0660671695b3f2242dd3113221c5e3f72342e0ff0ffa9d380502e223eb7c75f7.FilterOperatorsRequestBuilder) {
    return i0660671695b3f2242dd3113221c5e3f72342e0ff0ffa9d380502e223eb7c75f7.NewFilterOperatorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}\synchronization\jobs\{synchronizationJob-id}\schema\microsoft.graph.functions()
func (m *SchemaRequestBuilder) Functions()(*i2bb0f986999c75c0ca4bdcdc1c396d94ceed8cc193cfa0837b2802ca7f2c8fab.FunctionsRequestBuilder) {
    return i2bb0f986999c75c0ca4bdcdc1c396d94ceed8cc193cfa0837b2802ca7f2c8fab.NewFunctionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The synchronization schema configured for the job.
// Parameters:
//  - options : Options for the request
func (m *SchemaRequestBuilder) Get(options *SchemaRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SynchronizationSchema, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSynchronizationSchema() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SynchronizationSchema), nil
}
func (m *SchemaRequestBuilder) ParseExpression()(*if23bb949ea15062726d8be4224904c7bd42d0e0e3dd45612ea423af36b5ad737.ParseExpressionRequestBuilder) {
    return if23bb949ea15062726d8be4224904c7bd42d0e0e3dd45612ea423af36b5ad737.NewParseExpressionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The synchronization schema configured for the job.
// Parameters:
//  - options : Options for the request
func (m *SchemaRequestBuilder) Patch(options *SchemaRequestBuilderPatchOptions)(error) {
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
