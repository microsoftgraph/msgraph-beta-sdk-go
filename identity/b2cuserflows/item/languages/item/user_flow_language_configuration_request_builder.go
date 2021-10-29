package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ic56e48b961e57de85d10390f870dd3e57722e4a44a0964781543b92374e68259 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/languages/item/defaultpages"
    iecf50a78d9d6bd3f01a4f78210ea956d05c66f139964a7db9869f81ca435eaca "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/languages/item/overridespages"
    i2fb4388026c86a21e8b8d1776d2c5f2b64cc93696f3f4d055161caa5dff7e785 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/languages/item/defaultpages/item"
    ia8cb0ef728fb55863ea7e003c3f99fbd388a2e1017ae424b5f63e5e4b1660d67 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/languages/item/overridespages/item"
)

// Builds and executes requests for operations under \identity\b2cUserFlows\{b2cIdentityUserFlow-id}\languages\{userFlowLanguageConfiguration-id}
type UserFlowLanguageConfigurationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type UserFlowLanguageConfigurationRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type UserFlowLanguageConfigurationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UserFlowLanguageConfigurationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The languages supported for customization within the user flow. Language customization is not enabled by default in B2C user flows.
type UserFlowLanguageConfigurationRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type UserFlowLanguageConfigurationRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserFlowLanguageConfiguration;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new UserFlowLanguageConfigurationRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewUserFlowLanguageConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserFlowLanguageConfigurationRequestBuilder) {
    m := &UserFlowLanguageConfigurationRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/identity/b2cUserFlows/{b2cIdentityUserFlow_id}/languages/{userFlowLanguageConfiguration_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new UserFlowLanguageConfigurationRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewUserFlowLanguageConfigurationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserFlowLanguageConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserFlowLanguageConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// The languages supported for customization within the user flow. Language customization is not enabled by default in B2C user flows.
// Parameters:
//  - options : Options for the request
func (m *UserFlowLanguageConfigurationRequestBuilder) CreateDeleteRequestInformation(options *UserFlowLanguageConfigurationRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The languages supported for customization within the user flow. Language customization is not enabled by default in B2C user flows.
// Parameters:
//  - options : Options for the request
func (m *UserFlowLanguageConfigurationRequestBuilder) CreateGetRequestInformation(options *UserFlowLanguageConfigurationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The languages supported for customization within the user flow. Language customization is not enabled by default in B2C user flows.
// Parameters:
//  - options : Options for the request
func (m *UserFlowLanguageConfigurationRequestBuilder) CreatePatchRequestInformation(options *UserFlowLanguageConfigurationRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *UserFlowLanguageConfigurationRequestBuilder) DefaultPages()(*ic56e48b961e57de85d10390f870dd3e57722e4a44a0964781543b92374e68259.DefaultPagesRequestBuilder) {
    return ic56e48b961e57de85d10390f870dd3e57722e4a44a0964781543b92374e68259.NewDefaultPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.identity.b2cUserFlows.item.languages.item.defaultPages.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *UserFlowLanguageConfigurationRequestBuilder) DefaultPagesById(id string)(*i2fb4388026c86a21e8b8d1776d2c5f2b64cc93696f3f4d055161caa5dff7e785.UserFlowLanguagePageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userFlowLanguagePage_id"] = id
    }
    return i2fb4388026c86a21e8b8d1776d2c5f2b64cc93696f3f4d055161caa5dff7e785.NewUserFlowLanguagePageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The languages supported for customization within the user flow. Language customization is not enabled by default in B2C user flows.
// Parameters:
//  - options : Options for the request
func (m *UserFlowLanguageConfigurationRequestBuilder) Delete(options *UserFlowLanguageConfigurationRequestBuilderDeleteOptions)(error) {
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
// The languages supported for customization within the user flow. Language customization is not enabled by default in B2C user flows.
// Parameters:
//  - options : Options for the request
func (m *UserFlowLanguageConfigurationRequestBuilder) Get(options *UserFlowLanguageConfigurationRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserFlowLanguageConfiguration, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUserFlowLanguageConfiguration() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserFlowLanguageConfiguration), nil
}
func (m *UserFlowLanguageConfigurationRequestBuilder) OverridesPages()(*iecf50a78d9d6bd3f01a4f78210ea956d05c66f139964a7db9869f81ca435eaca.OverridesPagesRequestBuilder) {
    return iecf50a78d9d6bd3f01a4f78210ea956d05c66f139964a7db9869f81ca435eaca.NewOverridesPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.identity.b2cUserFlows.item.languages.item.overridesPages.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *UserFlowLanguageConfigurationRequestBuilder) OverridesPagesById(id string)(*ia8cb0ef728fb55863ea7e003c3f99fbd388a2e1017ae424b5f63e5e4b1660d67.UserFlowLanguagePageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userFlowLanguagePage_id"] = id
    }
    return ia8cb0ef728fb55863ea7e003c3f99fbd388a2e1017ae424b5f63e5e4b1660d67.NewUserFlowLanguagePageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The languages supported for customization within the user flow. Language customization is not enabled by default in B2C user flows.
// Parameters:
//  - options : Options for the request
func (m *UserFlowLanguageConfigurationRequestBuilder) Patch(options *UserFlowLanguageConfigurationRequestBuilderPatchOptions)(error) {
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
