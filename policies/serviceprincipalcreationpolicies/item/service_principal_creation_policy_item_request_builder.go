package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i458957d039676a87049f925f9a0ac58100c522d4c5494921661924e3d7de9536 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/serviceprincipalcreationpolicies/item/excludes"
    ibb33c8122164f9a21b54efbd3817c55a46e13f7c672ae3565862893e0225a426 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/serviceprincipalcreationpolicies/item/includes"
    i8b8ea18691791149f6486e79113d56f56a6921846e0823954a25159321c50bdc "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/serviceprincipalcreationpolicies/item/includes/item"
    if3d69c085cc31553bffa6238747fd2f166e883710ce3cb7738b69051017fde73 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/serviceprincipalcreationpolicies/item/excludes/item"
)

// ServicePrincipalCreationPolicyItemRequestBuilder builds and executes requests for operations under \policies\servicePrincipalCreationPolicies\{servicePrincipalCreationPolicy-id}
type ServicePrincipalCreationPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ServicePrincipalCreationPolicyItemRequestBuilderDeleteOptions options for Delete
type ServicePrincipalCreationPolicyItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ServicePrincipalCreationPolicyItemRequestBuilderGetOptions options for Get
type ServicePrincipalCreationPolicyItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ServicePrincipalCreationPolicyItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ServicePrincipalCreationPolicyItemRequestBuilderGetQueryParameters get servicePrincipalCreationPolicies from policies
type ServicePrincipalCreationPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ServicePrincipalCreationPolicyItemRequestBuilderPatchOptions options for Patch
type ServicePrincipalCreationPolicyItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServicePrincipalCreationPolicy;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewServicePrincipalCreationPolicyItemRequestBuilderInternal instantiates a new ServicePrincipalCreationPolicyItemRequestBuilder and sets the default values.
func NewServicePrincipalCreationPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ServicePrincipalCreationPolicyItemRequestBuilder) {
    m := &ServicePrincipalCreationPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/servicePrincipalCreationPolicies/{servicePrincipalCreationPolicy_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewServicePrincipalCreationPolicyItemRequestBuilder instantiates a new ServicePrincipalCreationPolicyItemRequestBuilder and sets the default values.
func NewServicePrincipalCreationPolicyItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ServicePrincipalCreationPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicePrincipalCreationPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property servicePrincipalCreationPolicies for policies
func (m *ServicePrincipalCreationPolicyItemRequestBuilder) CreateDeleteRequestInformation(options *ServicePrincipalCreationPolicyItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get servicePrincipalCreationPolicies from policies
func (m *ServicePrincipalCreationPolicyItemRequestBuilder) CreateGetRequestInformation(options *ServicePrincipalCreationPolicyItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property servicePrincipalCreationPolicies in policies
func (m *ServicePrincipalCreationPolicyItemRequestBuilder) CreatePatchRequestInformation(options *ServicePrincipalCreationPolicyItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property servicePrincipalCreationPolicies for policies
func (m *ServicePrincipalCreationPolicyItemRequestBuilder) Delete(options *ServicePrincipalCreationPolicyItemRequestBuilderDeleteOptions)(error) {
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
func (m *ServicePrincipalCreationPolicyItemRequestBuilder) Excludes()(*i458957d039676a87049f925f9a0ac58100c522d4c5494921661924e3d7de9536.ExcludesRequestBuilder) {
    return i458957d039676a87049f925f9a0ac58100c522d4c5494921661924e3d7de9536.NewExcludesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExcludesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.policies.servicePrincipalCreationPolicies.item.excludes.item collection
func (m *ServicePrincipalCreationPolicyItemRequestBuilder) ExcludesById(id string)(*if3d69c085cc31553bffa6238747fd2f166e883710ce3cb7738b69051017fde73.ServicePrincipalCreationConditionSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["servicePrincipalCreationConditionSet_id"] = id
    }
    return if3d69c085cc31553bffa6238747fd2f166e883710ce3cb7738b69051017fde73.NewServicePrincipalCreationConditionSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get servicePrincipalCreationPolicies from policies
func (m *ServicePrincipalCreationPolicyItemRequestBuilder) Get(options *ServicePrincipalCreationPolicyItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServicePrincipalCreationPolicy, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewServicePrincipalCreationPolicy() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServicePrincipalCreationPolicy), nil
}
func (m *ServicePrincipalCreationPolicyItemRequestBuilder) Includes()(*ibb33c8122164f9a21b54efbd3817c55a46e13f7c672ae3565862893e0225a426.IncludesRequestBuilder) {
    return ibb33c8122164f9a21b54efbd3817c55a46e13f7c672ae3565862893e0225a426.NewIncludesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncludesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.policies.servicePrincipalCreationPolicies.item.includes.item collection
func (m *ServicePrincipalCreationPolicyItemRequestBuilder) IncludesById(id string)(*i8b8ea18691791149f6486e79113d56f56a6921846e0823954a25159321c50bdc.ServicePrincipalCreationConditionSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["servicePrincipalCreationConditionSet_id"] = id
    }
    return i8b8ea18691791149f6486e79113d56f56a6921846e0823954a25159321c50bdc.NewServicePrincipalCreationConditionSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property servicePrincipalCreationPolicies in policies
func (m *ServicePrincipalCreationPolicyItemRequestBuilder) Patch(options *ServicePrincipalCreationPolicyItemRequestBuilderPatchOptions)(error) {
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
