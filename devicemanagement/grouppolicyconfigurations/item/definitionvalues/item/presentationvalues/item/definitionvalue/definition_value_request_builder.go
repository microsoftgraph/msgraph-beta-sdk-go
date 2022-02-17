package definitionvalue

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    idf760bbe698f0e7ec33f0f98c22c00961039c99efc61075cdf539d6dc777e0c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyconfigurations/item/definitionvalues/item/presentationvalues/item/definitionvalue/ref"
)

// DefinitionValueRequestBuilder builds and executes requests for operations under \deviceManagement\groupPolicyConfigurations\{groupPolicyConfiguration-id}\definitionValues\{groupPolicyDefinitionValue-id}\presentationValues\{groupPolicyPresentationValue-id}\definitionValue
type DefinitionValueRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DefinitionValueRequestBuilderGetOptions options for Get
type DefinitionValueRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DefinitionValueRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DefinitionValueRequestBuilderGetQueryParameters the group policy definition value associated with the presentation value.
type DefinitionValueRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NewDefinitionValueRequestBuilderInternal instantiates a new DefinitionValueRequestBuilder and sets the default values.
func NewDefinitionValueRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DefinitionValueRequestBuilder) {
    m := &DefinitionValueRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyConfigurations/{groupPolicyConfiguration_id}/definitionValues/{groupPolicyDefinitionValue_id}/presentationValues/{groupPolicyPresentationValue_id}/definitionValue{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDefinitionValueRequestBuilder instantiates a new DefinitionValueRequestBuilder and sets the default values.
func NewDefinitionValueRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DefinitionValueRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDefinitionValueRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the group policy definition value associated with the presentation value.
func (m *DefinitionValueRequestBuilder) CreateGetRequestInformation(options *DefinitionValueRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the group policy definition value associated with the presentation value.
func (m *DefinitionValueRequestBuilder) Get(options *DefinitionValueRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyDefinitionValue, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewGroupPolicyDefinitionValue() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyDefinitionValue), nil
}
func (m *DefinitionValueRequestBuilder) Ref()(*idf760bbe698f0e7ec33f0f98c22c00961039c99efc61075cdf539d6dc777e0c2.RefRequestBuilder) {
    return idf760bbe698f0e7ec33f0f98c22c00961039c99efc61075cdf539d6dc777e0c2.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
