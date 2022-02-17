package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i48aa07ca5717eddbded5e2e8384a452f2c48c6a0ce5df5d6cc71d0ac0e99d828 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/featurerolloutpolicies/item/appliesto"
    i8cee271c00bcd456d3c1e1d15da437202f66853d03f66cb99750b1d59ab873b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/featurerolloutpolicies/item/appliesto/item"
)

// FeatureRolloutPolicyRequestBuilder builds and executes requests for operations under \directory\featureRolloutPolicies\{featureRolloutPolicy-id}
type FeatureRolloutPolicyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// FeatureRolloutPolicyRequestBuilderDeleteOptions options for Delete
type FeatureRolloutPolicyRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// FeatureRolloutPolicyRequestBuilderGetOptions options for Get
type FeatureRolloutPolicyRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *FeatureRolloutPolicyRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// FeatureRolloutPolicyRequestBuilderGetQueryParameters nullable.
type FeatureRolloutPolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// FeatureRolloutPolicyRequestBuilderPatchOptions options for Patch
type FeatureRolloutPolicyRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.FeatureRolloutPolicy;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *FeatureRolloutPolicyRequestBuilder) AppliesTo()(*i48aa07ca5717eddbded5e2e8384a452f2c48c6a0ce5df5d6cc71d0ac0e99d828.AppliesToRequestBuilder) {
    return i48aa07ca5717eddbded5e2e8384a452f2c48c6a0ce5df5d6cc71d0ac0e99d828.NewAppliesToRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppliesToById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.featureRolloutPolicies.item.appliesTo.item collection
func (m *FeatureRolloutPolicyRequestBuilder) AppliesToById(id string)(*i8cee271c00bcd456d3c1e1d15da437202f66853d03f66cb99750b1d59ab873b4.DirectoryObjectRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return i8cee271c00bcd456d3c1e1d15da437202f66853d03f66cb99750b1d59ab873b4.NewDirectoryObjectRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewFeatureRolloutPolicyRequestBuilderInternal instantiates a new FeatureRolloutPolicyRequestBuilder and sets the default values.
func NewFeatureRolloutPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*FeatureRolloutPolicyRequestBuilder) {
    m := &FeatureRolloutPolicyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory/featureRolloutPolicies/{featureRolloutPolicy_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewFeatureRolloutPolicyRequestBuilder instantiates a new FeatureRolloutPolicyRequestBuilder and sets the default values.
func NewFeatureRolloutPolicyRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*FeatureRolloutPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFeatureRolloutPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation nullable.
func (m *FeatureRolloutPolicyRequestBuilder) CreateDeleteRequestInformation(options *FeatureRolloutPolicyRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation nullable.
func (m *FeatureRolloutPolicyRequestBuilder) CreateGetRequestInformation(options *FeatureRolloutPolicyRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation nullable.
func (m *FeatureRolloutPolicyRequestBuilder) CreatePatchRequestInformation(options *FeatureRolloutPolicyRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete nullable.
func (m *FeatureRolloutPolicyRequestBuilder) Delete(options *FeatureRolloutPolicyRequestBuilderDeleteOptions)(error) {
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
// Get nullable.
func (m *FeatureRolloutPolicyRequestBuilder) Get(options *FeatureRolloutPolicyRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.FeatureRolloutPolicy, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewFeatureRolloutPolicy() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.FeatureRolloutPolicy), nil
}
// Patch nullable.
func (m *FeatureRolloutPolicyRequestBuilder) Patch(options *FeatureRolloutPolicyRequestBuilderPatchOptions)(error) {
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
