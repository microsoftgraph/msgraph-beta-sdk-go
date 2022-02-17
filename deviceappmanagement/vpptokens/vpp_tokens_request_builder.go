package vpptokens

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i24db185bc1635f2b2be054ec817ced066302c11adf27d59692cd7b262f32698e "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/vpptokens/getlicensesforappwithbundleid"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ib956f45626ba71f8b797fc4b16070658a39a611e2d6aa111aea4eca180955477 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/vpptokens/synclicensecounts"
)

// VppTokensRequestBuilder builds and executes requests for operations under \deviceAppManagement\vppTokens
type VppTokensRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// VppTokensRequestBuilderGetOptions options for Get
type VppTokensRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *VppTokensRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// VppTokensRequestBuilderGetQueryParameters list of Vpp tokens for this organization.
type VppTokensRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// VppTokensRequestBuilderPostOptions options for Post
type VppTokensRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.VppToken;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewVppTokensRequestBuilderInternal instantiates a new VppTokensRequestBuilder and sets the default values.
func NewVppTokensRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*VppTokensRequestBuilder) {
    m := &VppTokensRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/vppTokens{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewVppTokensRequestBuilder instantiates a new VppTokensRequestBuilder and sets the default values.
func NewVppTokensRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*VppTokensRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVppTokensRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation list of Vpp tokens for this organization.
func (m *VppTokensRequestBuilder) CreateGetRequestInformation(options *VppTokensRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation list of Vpp tokens for this organization.
func (m *VppTokensRequestBuilder) CreatePostRequestInformation(options *VppTokensRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
// Get list of Vpp tokens for this organization.
func (m *VppTokensRequestBuilder) Get(options *VppTokensRequestBuilderGetOptions)(*VppTokensResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewVppTokensResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*VppTokensResponse), nil
}
// GetLicensesForAppWithBundleId builds and executes requests for operations under \deviceAppManagement\vppTokens\microsoft.graph.getLicensesForApp(bundleId='{bundleId}')
func (m *VppTokensRequestBuilder) GetLicensesForAppWithBundleId(bundleId *string)(*i24db185bc1635f2b2be054ec817ced066302c11adf27d59692cd7b262f32698e.GetLicensesForAppWithBundleIdRequestBuilder) {
    return i24db185bc1635f2b2be054ec817ced066302c11adf27d59692cd7b262f32698e.NewGetLicensesForAppWithBundleIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, bundleId);
}
// Post list of Vpp tokens for this organization.
func (m *VppTokensRequestBuilder) Post(options *VppTokensRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.VppToken, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewVppToken() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.VppToken), nil
}
func (m *VppTokensRequestBuilder) SyncLicenseCounts()(*ib956f45626ba71f8b797fc4b16070658a39a611e2d6aa111aea4eca180955477.SyncLicenseCountsRequestBuilder) {
    return ib956f45626ba71f8b797fc4b16070658a39a611e2d6aa111aea4eca180955477.NewSyncLicenseCountsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
