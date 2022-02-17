package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i574815dbbb1e9c51a965e25539eed078ec9ab55919033c1c32dc721a519d9f21 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2xuserflows/item/userattributeassignments"
    i7dcf98ba71a64ad6e1a7f1d9cbe306c7616a2b4fbec2dbc655c658192d35a36b "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2xuserflows/item/languages"
    i961eb1bb9c95ad9a1c53819fc521c8a4e4e3dd72967d787f44f08b9b8c796ad2 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2xuserflows/item/identityproviders"
    icd6b8fe789e0ddefec348d941d6505a9be0443397887095b870d0a0a621d081a "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2xuserflows/item/userflowidentityproviders"
    i45862c76d997425b3529626dabb7258b739baf239b9d5edccd7c60df97c16804 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2xuserflows/item/languages/item"
    i8706143552eeb09a05c35bc53efe01f6fe2a821a3d0568b7c706289ea9919d1d "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2xuserflows/item/userattributeassignments/item"
)

// B2xIdentityUserFlowRequestBuilder builds and executes requests for operations under \identity\b2xUserFlows\{b2xIdentityUserFlow-id}
type B2xIdentityUserFlowRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// B2xIdentityUserFlowRequestBuilderDeleteOptions options for Delete
type B2xIdentityUserFlowRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// B2xIdentityUserFlowRequestBuilderGetOptions options for Get
type B2xIdentityUserFlowRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *B2xIdentityUserFlowRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// B2xIdentityUserFlowRequestBuilderGetQueryParameters represents entry point for B2X/self-service sign-up identity userflows.
type B2xIdentityUserFlowRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// B2xIdentityUserFlowRequestBuilderPatchOptions options for Patch
type B2xIdentityUserFlowRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.B2xIdentityUserFlow;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewB2xIdentityUserFlowRequestBuilderInternal instantiates a new B2xIdentityUserFlowRequestBuilder and sets the default values.
func NewB2xIdentityUserFlowRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*B2xIdentityUserFlowRequestBuilder) {
    m := &B2xIdentityUserFlowRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewB2xIdentityUserFlowRequestBuilder instantiates a new B2xIdentityUserFlowRequestBuilder and sets the default values.
func NewB2xIdentityUserFlowRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*B2xIdentityUserFlowRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2xIdentityUserFlowRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation represents entry point for B2X/self-service sign-up identity userflows.
func (m *B2xIdentityUserFlowRequestBuilder) CreateDeleteRequestInformation(options *B2xIdentityUserFlowRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation represents entry point for B2X/self-service sign-up identity userflows.
func (m *B2xIdentityUserFlowRequestBuilder) CreateGetRequestInformation(options *B2xIdentityUserFlowRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation represents entry point for B2X/self-service sign-up identity userflows.
func (m *B2xIdentityUserFlowRequestBuilder) CreatePatchRequestInformation(options *B2xIdentityUserFlowRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete represents entry point for B2X/self-service sign-up identity userflows.
func (m *B2xIdentityUserFlowRequestBuilder) Delete(options *B2xIdentityUserFlowRequestBuilderDeleteOptions)(error) {
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
// Get represents entry point for B2X/self-service sign-up identity userflows.
func (m *B2xIdentityUserFlowRequestBuilder) Get(options *B2xIdentityUserFlowRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.B2xIdentityUserFlow, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewB2xIdentityUserFlow() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.B2xIdentityUserFlow), nil
}
func (m *B2xIdentityUserFlowRequestBuilder) IdentityProviders()(*i961eb1bb9c95ad9a1c53819fc521c8a4e4e3dd72967d787f44f08b9b8c796ad2.IdentityProvidersRequestBuilder) {
    return i961eb1bb9c95ad9a1c53819fc521c8a4e4e3dd72967d787f44f08b9b8c796ad2.NewIdentityProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *B2xIdentityUserFlowRequestBuilder) Languages()(*i7dcf98ba71a64ad6e1a7f1d9cbe306c7616a2b4fbec2dbc655c658192d35a36b.LanguagesRequestBuilder) {
    return i7dcf98ba71a64ad6e1a7f1d9cbe306c7616a2b4fbec2dbc655c658192d35a36b.NewLanguagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LanguagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.b2xUserFlows.item.languages.item collection
func (m *B2xIdentityUserFlowRequestBuilder) LanguagesById(id string)(*i45862c76d997425b3529626dabb7258b739baf239b9d5edccd7c60df97c16804.UserFlowLanguageConfigurationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userFlowLanguageConfiguration_id"] = id
    }
    return i45862c76d997425b3529626dabb7258b739baf239b9d5edccd7c60df97c16804.NewUserFlowLanguageConfigurationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch represents entry point for B2X/self-service sign-up identity userflows.
func (m *B2xIdentityUserFlowRequestBuilder) Patch(options *B2xIdentityUserFlowRequestBuilderPatchOptions)(error) {
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
func (m *B2xIdentityUserFlowRequestBuilder) UserAttributeAssignments()(*i574815dbbb1e9c51a965e25539eed078ec9ab55919033c1c32dc721a519d9f21.UserAttributeAssignmentsRequestBuilder) {
    return i574815dbbb1e9c51a965e25539eed078ec9ab55919033c1c32dc721a519d9f21.NewUserAttributeAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserAttributeAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.b2xUserFlows.item.userAttributeAssignments.item collection
func (m *B2xIdentityUserFlowRequestBuilder) UserAttributeAssignmentsById(id string)(*i8706143552eeb09a05c35bc53efe01f6fe2a821a3d0568b7c706289ea9919d1d.IdentityUserFlowAttributeAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityUserFlowAttributeAssignment_id"] = id
    }
    return i8706143552eeb09a05c35bc53efe01f6fe2a821a3d0568b7c706289ea9919d1d.NewIdentityUserFlowAttributeAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *B2xIdentityUserFlowRequestBuilder) UserFlowIdentityProviders()(*icd6b8fe789e0ddefec348d941d6505a9be0443397887095b870d0a0a621d081a.UserFlowIdentityProvidersRequestBuilder) {
    return icd6b8fe789e0ddefec348d941d6505a9be0443397887095b870d0a0a621d081a.NewUserFlowIdentityProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
