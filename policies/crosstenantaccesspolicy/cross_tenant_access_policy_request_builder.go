package crosstenantaccesspolicy

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i12ec00c0cda085cf3d19e747ae27322079cbb18e5b556d17b1a8a5d083c42b4f "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/crosstenantaccesspolicy/partners"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i9012de60632d1258f2b5642de666977d54e3da254d584ad456e64c40f801e738 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/crosstenantaccesspolicy/default_escaped"
    i671445f12e1144f0707e713947da70fbe18307579db947c865515610f6068a76 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/crosstenantaccesspolicy/partners/item"
)

// CrossTenantAccessPolicyRequestBuilder provides operations to manage the crossTenantAccessPolicy property of the microsoft.graph.policyRoot entity.
type CrossTenantAccessPolicyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CrossTenantAccessPolicyRequestBuilderDeleteOptions options for Delete
type CrossTenantAccessPolicyRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CrossTenantAccessPolicyRequestBuilderGetOptions options for Get
type CrossTenantAccessPolicyRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CrossTenantAccessPolicyRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CrossTenantAccessPolicyRequestBuilderGetQueryParameters the custom rules that define an access scenario when interacting with external Azure AD tenants.
type CrossTenantAccessPolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// CrossTenantAccessPolicyRequestBuilderPatchOptions options for Patch
type CrossTenantAccessPolicyRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CrossTenantAccessPolicyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewCrossTenantAccessPolicyRequestBuilderInternal instantiates a new CrossTenantAccessPolicyRequestBuilder and sets the default values.
func NewCrossTenantAccessPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CrossTenantAccessPolicyRequestBuilder) {
    m := &CrossTenantAccessPolicyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/crossTenantAccessPolicy{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCrossTenantAccessPolicyRequestBuilder instantiates a new CrossTenantAccessPolicyRequestBuilder and sets the default values.
func NewCrossTenantAccessPolicyRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CrossTenantAccessPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCrossTenantAccessPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property crossTenantAccessPolicy for policies
func (m *CrossTenantAccessPolicyRequestBuilder) CreateDeleteRequestInformation(options *CrossTenantAccessPolicyRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the custom rules that define an access scenario when interacting with external Azure AD tenants.
func (m *CrossTenantAccessPolicyRequestBuilder) CreateGetRequestInformation(options *CrossTenantAccessPolicyRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property crossTenantAccessPolicy in policies
func (m *CrossTenantAccessPolicyRequestBuilder) CreatePatchRequestInformation(options *CrossTenantAccessPolicyRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *CrossTenantAccessPolicyRequestBuilder) Default_escaped()(*i9012de60632d1258f2b5642de666977d54e3da254d584ad456e64c40f801e738.DefaultRequestBuilder) {
    return i9012de60632d1258f2b5642de666977d54e3da254d584ad456e64c40f801e738.NewDefaultRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property crossTenantAccessPolicy for policies
func (m *CrossTenantAccessPolicyRequestBuilder) Delete(options *CrossTenantAccessPolicyRequestBuilderDeleteOptions)(error) {
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
// Get the custom rules that define an access scenario when interacting with external Azure AD tenants.
func (m *CrossTenantAccessPolicyRequestBuilder) Get(options *CrossTenantAccessPolicyRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CrossTenantAccessPolicyable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateCrossTenantAccessPolicyFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CrossTenantAccessPolicyable), nil
}
func (m *CrossTenantAccessPolicyRequestBuilder) Partners()(*i12ec00c0cda085cf3d19e747ae27322079cbb18e5b556d17b1a8a5d083c42b4f.PartnersRequestBuilder) {
    return i12ec00c0cda085cf3d19e747ae27322079cbb18e5b556d17b1a8a5d083c42b4f.NewPartnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PartnersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.policies.crossTenantAccessPolicy.partners.item collection
func (m *CrossTenantAccessPolicyRequestBuilder) PartnersById(id string)(*i671445f12e1144f0707e713947da70fbe18307579db947c865515610f6068a76.CrossTenantAccessPolicyConfigurationPartnerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["crossTenantAccessPolicyConfigurationPartner_tenantId"] = id
    }
    return i671445f12e1144f0707e713947da70fbe18307579db947c865515610f6068a76.NewCrossTenantAccessPolicyConfigurationPartnerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property crossTenantAccessPolicy in policies
func (m *CrossTenantAccessPolicyRequestBuilder) Patch(options *CrossTenantAccessPolicyRequestBuilderPatchOptions)(error) {
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
