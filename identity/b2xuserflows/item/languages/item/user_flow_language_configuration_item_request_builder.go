package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i38edbc2c15228f0a4386c06fe1cd74f5b61b0532c060a337d0d143cee71f76dc "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2xuserflows/item/languages/item/overridespages"
    id8b4aa20110754f7a6332241da005226519c2fdc8bae4a32149df9d988ded399 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2xuserflows/item/languages/item/defaultpages"
    id97df1ec1b4e9135f6abf467520dc7657ad6787d1073dfe17aa2b9434dc66f74 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2xuserflows/item/languages/item/overridespages/item"
    if7f959c11e06f172ec88344202e39c37c579835cbcd855889fc28d546c7c6126 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2xuserflows/item/languages/item/defaultpages/item"
)

// UserFlowLanguageConfigurationItemRequestBuilder provides operations to manage the languages property of the microsoft.graph.b2xIdentityUserFlow entity.
type UserFlowLanguageConfigurationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UserFlowLanguageConfigurationItemRequestBuilderDeleteOptions options for Delete
type UserFlowLanguageConfigurationItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserFlowLanguageConfigurationItemRequestBuilderGetOptions options for Get
type UserFlowLanguageConfigurationItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UserFlowLanguageConfigurationItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserFlowLanguageConfigurationItemRequestBuilderGetQueryParameters the languages supported for customization within the user flow. Language customization is enabled by default in self-service sign-up user flow. You cannot create custom languages in self-service sign-up user flows.
type UserFlowLanguageConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UserFlowLanguageConfigurationItemRequestBuilderPatchOptions options for Patch
type UserFlowLanguageConfigurationItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserFlowLanguageConfigurationable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUserFlowLanguageConfigurationItemRequestBuilderInternal instantiates a new UserFlowLanguageConfigurationItemRequestBuilder and sets the default values.
func NewUserFlowLanguageConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserFlowLanguageConfigurationItemRequestBuilder) {
    m := &UserFlowLanguageConfigurationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow_id}/languages/{userFlowLanguageConfiguration_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserFlowLanguageConfigurationItemRequestBuilder instantiates a new UserFlowLanguageConfigurationItemRequestBuilder and sets the default values.
func NewUserFlowLanguageConfigurationItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserFlowLanguageConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserFlowLanguageConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property languages for identity
func (m *UserFlowLanguageConfigurationItemRequestBuilder) CreateDeleteRequestInformation(options *UserFlowLanguageConfigurationItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the languages supported for customization within the user flow. Language customization is enabled by default in self-service sign-up user flow. You cannot create custom languages in self-service sign-up user flows.
func (m *UserFlowLanguageConfigurationItemRequestBuilder) CreateGetRequestInformation(options *UserFlowLanguageConfigurationItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property languages in identity
func (m *UserFlowLanguageConfigurationItemRequestBuilder) CreatePatchRequestInformation(options *UserFlowLanguageConfigurationItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *UserFlowLanguageConfigurationItemRequestBuilder) DefaultPages()(*id8b4aa20110754f7a6332241da005226519c2fdc8bae4a32149df9d988ded399.DefaultPagesRequestBuilder) {
    return id8b4aa20110754f7a6332241da005226519c2fdc8bae4a32149df9d988ded399.NewDefaultPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DefaultPagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.b2xUserFlows.item.languages.item.defaultPages.item collection
func (m *UserFlowLanguageConfigurationItemRequestBuilder) DefaultPagesById(id string)(*if7f959c11e06f172ec88344202e39c37c579835cbcd855889fc28d546c7c6126.UserFlowLanguagePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userFlowLanguagePage_id"] = id
    }
    return if7f959c11e06f172ec88344202e39c37c579835cbcd855889fc28d546c7c6126.NewUserFlowLanguagePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property languages for identity
func (m *UserFlowLanguageConfigurationItemRequestBuilder) Delete(options *UserFlowLanguageConfigurationItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the languages supported for customization within the user flow. Language customization is enabled by default in self-service sign-up user flow. You cannot create custom languages in self-service sign-up user flows.
func (m *UserFlowLanguageConfigurationItemRequestBuilder) Get(options *UserFlowLanguageConfigurationItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserFlowLanguageConfigurationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateUserFlowLanguageConfigurationFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserFlowLanguageConfigurationable), nil
}
func (m *UserFlowLanguageConfigurationItemRequestBuilder) OverridesPages()(*i38edbc2c15228f0a4386c06fe1cd74f5b61b0532c060a337d0d143cee71f76dc.OverridesPagesRequestBuilder) {
    return i38edbc2c15228f0a4386c06fe1cd74f5b61b0532c060a337d0d143cee71f76dc.NewOverridesPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OverridesPagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.b2xUserFlows.item.languages.item.overridesPages.item collection
func (m *UserFlowLanguageConfigurationItemRequestBuilder) OverridesPagesById(id string)(*id97df1ec1b4e9135f6abf467520dc7657ad6787d1073dfe17aa2b9434dc66f74.UserFlowLanguagePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userFlowLanguagePage_id"] = id
    }
    return id97df1ec1b4e9135f6abf467520dc7657ad6787d1073dfe17aa2b9434dc66f74.NewUserFlowLanguagePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property languages in identity
func (m *UserFlowLanguageConfigurationItemRequestBuilder) Patch(options *UserFlowLanguageConfigurationItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
