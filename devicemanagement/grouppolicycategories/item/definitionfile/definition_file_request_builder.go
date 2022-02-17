package definitionfile

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ibf2781b63a27300f1ed5dadf7ac3afbfbb2c5184ee96653d78e83cfd8fe25952 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicycategories/item/definitionfile/ref"
)

// DefinitionFileRequestBuilder builds and executes requests for operations under \deviceManagement\groupPolicyCategories\{groupPolicyCategory-id}\definitionFile
type DefinitionFileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DefinitionFileRequestBuilderGetOptions options for Get
type DefinitionFileRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DefinitionFileRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DefinitionFileRequestBuilderGetQueryParameters the id of the definition file the category came from
type DefinitionFileRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NewDefinitionFileRequestBuilderInternal instantiates a new DefinitionFileRequestBuilder and sets the default values.
func NewDefinitionFileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DefinitionFileRequestBuilder) {
    m := &DefinitionFileRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyCategories/{groupPolicyCategory_id}/definitionFile{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDefinitionFileRequestBuilder instantiates a new DefinitionFileRequestBuilder and sets the default values.
func NewDefinitionFileRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DefinitionFileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDefinitionFileRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the id of the definition file the category came from
func (m *DefinitionFileRequestBuilder) CreateGetRequestInformation(options *DefinitionFileRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the id of the definition file the category came from
func (m *DefinitionFileRequestBuilder) Get(options *DefinitionFileRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyDefinitionFile, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewGroupPolicyDefinitionFile() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyDefinitionFile), nil
}
func (m *DefinitionFileRequestBuilder) Ref()(*ibf2781b63a27300f1ed5dadf7ac3afbfbb2c5184ee96653d78e83cfd8fe25952.RefRequestBuilder) {
    return ibf2781b63a27300f1ed5dadf7ac3afbfbb2c5184ee96653d78e83cfd8fe25952.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
