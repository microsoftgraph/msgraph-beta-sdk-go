package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i04b2dcef167a2cae4f7fe917f68d75103bd1143f7e59881482a2a8dce2921570 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/authorizationpolicy/item/defaultuserroleoverrides"
    if552963b64d01abf5402660068a63a336053e969fbe13841bbe4817b65852c68 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/authorizationpolicy/item/defaultuserroleoverrides/item"
)

// AuthorizationPolicyRequestBuilder builds and executes requests for operations under \policies\authorizationPolicy\{authorizationPolicy-id}
type AuthorizationPolicyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AuthorizationPolicyRequestBuilderDeleteOptions options for Delete
type AuthorizationPolicyRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AuthorizationPolicyRequestBuilderGetOptions options for Get
type AuthorizationPolicyRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AuthorizationPolicyRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AuthorizationPolicyRequestBuilderGetQueryParameters the policy that controls Azure AD authorization settings.
type AuthorizationPolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AuthorizationPolicyRequestBuilderPatchOptions options for Patch
type AuthorizationPolicyRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AuthorizationPolicy;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewAuthorizationPolicyRequestBuilderInternal instantiates a new AuthorizationPolicyRequestBuilder and sets the default values.
func NewAuthorizationPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AuthorizationPolicyRequestBuilder) {
    m := &AuthorizationPolicyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/authorizationPolicy/{authorizationPolicy_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAuthorizationPolicyRequestBuilder instantiates a new AuthorizationPolicyRequestBuilder and sets the default values.
func NewAuthorizationPolicyRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AuthorizationPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthorizationPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the policy that controls Azure AD authorization settings.
func (m *AuthorizationPolicyRequestBuilder) CreateDeleteRequestInformation(options *AuthorizationPolicyRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the policy that controls Azure AD authorization settings.
func (m *AuthorizationPolicyRequestBuilder) CreateGetRequestInformation(options *AuthorizationPolicyRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the policy that controls Azure AD authorization settings.
func (m *AuthorizationPolicyRequestBuilder) CreatePatchRequestInformation(options *AuthorizationPolicyRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *AuthorizationPolicyRequestBuilder) DefaultUserRoleOverrides()(*i04b2dcef167a2cae4f7fe917f68d75103bd1143f7e59881482a2a8dce2921570.DefaultUserRoleOverridesRequestBuilder) {
    return i04b2dcef167a2cae4f7fe917f68d75103bd1143f7e59881482a2a8dce2921570.NewDefaultUserRoleOverridesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DefaultUserRoleOverridesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.policies.authorizationPolicy.item.defaultUserRoleOverrides.item collection
func (m *AuthorizationPolicyRequestBuilder) DefaultUserRoleOverridesById(id string)(*if552963b64d01abf5402660068a63a336053e969fbe13841bbe4817b65852c68.DefaultUserRoleOverrideRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["defaultUserRoleOverride_id"] = id
    }
    return if552963b64d01abf5402660068a63a336053e969fbe13841bbe4817b65852c68.NewDefaultUserRoleOverrideRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete the policy that controls Azure AD authorization settings.
func (m *AuthorizationPolicyRequestBuilder) Delete(options *AuthorizationPolicyRequestBuilderDeleteOptions)(error) {
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
// Get the policy that controls Azure AD authorization settings.
func (m *AuthorizationPolicyRequestBuilder) Get(options *AuthorizationPolicyRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AuthorizationPolicy, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAuthorizationPolicy() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AuthorizationPolicy), nil
}
// Patch the policy that controls Azure AD authorization settings.
func (m *AuthorizationPolicyRequestBuilder) Patch(options *AuthorizationPolicyRequestBuilderPatchOptions)(error) {
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
