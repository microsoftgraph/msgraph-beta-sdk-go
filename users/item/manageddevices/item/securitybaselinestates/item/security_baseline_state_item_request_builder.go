package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i70aafda266ed48512d40299c2aeb1764c358fa5ac74245325656d31ccd11c080 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/securitybaselinestates/item/settingstates"
    id053ea3d5bc635d6ccfc94bdc721142a88d882ac2c76384142b42f7c8eb43657 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/securitybaselinestates/item/settingstates/item"
)

// SecurityBaselineStateItemRequestBuilder builds and executes requests for operations under \users\{user-id}\managedDevices\{managedDevice-id}\securityBaselineStates\{securityBaselineState-id}
type SecurityBaselineStateItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SecurityBaselineStateItemRequestBuilderDeleteOptions options for Delete
type SecurityBaselineStateItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SecurityBaselineStateItemRequestBuilderGetOptions options for Get
type SecurityBaselineStateItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SecurityBaselineStateItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SecurityBaselineStateItemRequestBuilderGetQueryParameters security baseline states for this device.
type SecurityBaselineStateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SecurityBaselineStateItemRequestBuilderPatchOptions options for Patch
type SecurityBaselineStateItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SecurityBaselineState;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewSecurityBaselineStateItemRequestBuilderInternal instantiates a new SecurityBaselineStateItemRequestBuilder and sets the default values.
func NewSecurityBaselineStateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SecurityBaselineStateItemRequestBuilder) {
    m := &SecurityBaselineStateItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/managedDevices/{managedDevice_id}/securityBaselineStates/{securityBaselineState_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSecurityBaselineStateItemRequestBuilder instantiates a new SecurityBaselineStateItemRequestBuilder and sets the default values.
func NewSecurityBaselineStateItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SecurityBaselineStateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSecurityBaselineStateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation security baseline states for this device.
func (m *SecurityBaselineStateItemRequestBuilder) CreateDeleteRequestInformation(options *SecurityBaselineStateItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation security baseline states for this device.
func (m *SecurityBaselineStateItemRequestBuilder) CreateGetRequestInformation(options *SecurityBaselineStateItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation security baseline states for this device.
func (m *SecurityBaselineStateItemRequestBuilder) CreatePatchRequestInformation(options *SecurityBaselineStateItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete security baseline states for this device.
func (m *SecurityBaselineStateItemRequestBuilder) Delete(options *SecurityBaselineStateItemRequestBuilderDeleteOptions)(error) {
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
// Get security baseline states for this device.
func (m *SecurityBaselineStateItemRequestBuilder) Get(options *SecurityBaselineStateItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SecurityBaselineState, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSecurityBaselineState() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SecurityBaselineState), nil
}
// Patch security baseline states for this device.
func (m *SecurityBaselineStateItemRequestBuilder) Patch(options *SecurityBaselineStateItemRequestBuilderPatchOptions)(error) {
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
func (m *SecurityBaselineStateItemRequestBuilder) SettingStates()(*i70aafda266ed48512d40299c2aeb1764c358fa5ac74245325656d31ccd11c080.SettingStatesRequestBuilder) {
    return i70aafda266ed48512d40299c2aeb1764c358fa5ac74245325656d31ccd11c080.NewSettingStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.managedDevices.item.securityBaselineStates.item.settingStates.item collection
func (m *SecurityBaselineStateItemRequestBuilder) SettingStatesById(id string)(*id053ea3d5bc635d6ccfc94bdc721142a88d882ac2c76384142b42f7c8eb43657.SecurityBaselineSettingStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["securityBaselineSettingState_id"] = id
    }
    return id053ea3d5bc635d6ccfc94bdc721142a88d882ac2c76384142b42f7c8eb43657.NewSecurityBaselineSettingStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
