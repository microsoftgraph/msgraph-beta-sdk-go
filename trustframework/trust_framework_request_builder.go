package trustframework

import (
    i45ac5b9b0bfa1cab1bbb052f4f9c7d4e05cfb96edb3cb1d48ad9d002664d39d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/trustframework/policies"
    i7b667425d5ba5070b20ca9a446c105ed8c4749fa25d168849c6e0258ab640270 "github.com/microsoftgraph/msgraph-beta-sdk-go/trustframework/keysets"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i130f0626fe3c071232b6f7dc633093963e558c51511332c8d4c462636204a280 "github.com/microsoftgraph/msgraph-beta-sdk-go/trustframework/keysets/item"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i76c4dfef5c53e4b89a114a49ab372d39e04290df03064db59620347dee08cbcd "github.com/microsoftgraph/msgraph-beta-sdk-go/trustframework/policies/item"
)

// TrustFrameworkRequestBuilder builds and executes requests for operations under \trustFramework
type TrustFrameworkRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TrustFrameworkRequestBuilderGetOptions options for Get
type TrustFrameworkRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TrustFrameworkRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TrustFrameworkRequestBuilderGetQueryParameters get trustFramework
type TrustFrameworkRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// TrustFrameworkRequestBuilderPatchOptions options for Patch
type TrustFrameworkRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TrustFramework;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewTrustFrameworkRequestBuilderInternal instantiates a new TrustFrameworkRequestBuilder and sets the default values.
func NewTrustFrameworkRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TrustFrameworkRequestBuilder) {
    m := &TrustFrameworkRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/trustFramework{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTrustFrameworkRequestBuilder instantiates a new TrustFrameworkRequestBuilder and sets the default values.
func NewTrustFrameworkRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TrustFrameworkRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTrustFrameworkRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get trustFramework
func (m *TrustFrameworkRequestBuilder) CreateGetRequestInformation(options *TrustFrameworkRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update trustFramework
func (m *TrustFrameworkRequestBuilder) CreatePatchRequestInformation(options *TrustFrameworkRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get get trustFramework
func (m *TrustFrameworkRequestBuilder) Get(options *TrustFrameworkRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TrustFramework, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewTrustFramework() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TrustFramework), nil
}
func (m *TrustFrameworkRequestBuilder) KeySets()(*i7b667425d5ba5070b20ca9a446c105ed8c4749fa25d168849c6e0258ab640270.KeySetsRequestBuilder) {
    return i7b667425d5ba5070b20ca9a446c105ed8c4749fa25d168849c6e0258ab640270.NewKeySetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// KeySetsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.trustFramework.keySets.item collection
func (m *TrustFrameworkRequestBuilder) KeySetsById(id string)(*i130f0626fe3c071232b6f7dc633093963e558c51511332c8d4c462636204a280.TrustFrameworkKeySetRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["trustFrameworkKeySet_id"] = id
    }
    return i130f0626fe3c071232b6f7dc633093963e558c51511332c8d4c462636204a280.NewTrustFrameworkKeySetRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update trustFramework
func (m *TrustFrameworkRequestBuilder) Patch(options *TrustFrameworkRequestBuilderPatchOptions)(error) {
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
func (m *TrustFrameworkRequestBuilder) Policies()(*i45ac5b9b0bfa1cab1bbb052f4f9c7d4e05cfb96edb3cb1d48ad9d002664d39d9.PoliciesRequestBuilder) {
    return i45ac5b9b0bfa1cab1bbb052f4f9c7d4e05cfb96edb3cb1d48ad9d002664d39d9.NewPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.trustFramework.policies.item collection
func (m *TrustFrameworkRequestBuilder) PoliciesById(id string)(*i76c4dfef5c53e4b89a114a49ab372d39e04290df03064db59620347dee08cbcd.TrustFrameworkPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["trustFrameworkPolicy_id"] = id
    }
    return i76c4dfef5c53e4b89a114a49ab372d39e04290df03064db59620347dee08cbcd.NewTrustFrameworkPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
