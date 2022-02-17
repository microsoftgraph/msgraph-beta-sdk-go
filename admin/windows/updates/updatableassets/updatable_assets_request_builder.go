package updatableassets

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1ad9e4f03fd41b4ff3469626c2221aacd3c3b097757de111b29a11d0e40e5af7 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/updatableassets/enrollassetsbyid"
    i4dfa7ffb4b8c911e44699c4105daaccddf21f9cbf6d740540587c4fac3cb8746 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/updatableassets/unenrollassets"
    i5e033ef8b15c49341c073fc21c659e5769ead616f1066db183a743d44638c27e "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/updatableassets/enrollassets"
    i9b7b56068ae594d878d9164fec1e047566feb9a550ec5e453c176a1f0a5f2796 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/updatableassets/unenrollassetsbyid"
)

// UpdatableAssetsRequestBuilder builds and executes requests for operations under \admin\windows\updates\updatableAssets
type UpdatableAssetsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UpdatableAssetsRequestBuilderGetOptions options for Get
type UpdatableAssetsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UpdatableAssetsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UpdatableAssetsRequestBuilderGetQueryParameters assets registered with the deployment service that can receive updates. Read-only.
type UpdatableAssetsRequestBuilderGetQueryParameters struct {
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
// UpdatableAssetsRequestBuilderPostOptions options for Post
type UpdatableAssetsRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UpdatableAsset;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUpdatableAssetsRequestBuilderInternal instantiates a new UpdatableAssetsRequestBuilder and sets the default values.
func NewUpdatableAssetsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UpdatableAssetsRequestBuilder) {
    m := &UpdatableAssetsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/windows/updates/updatableAssets{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUpdatableAssetsRequestBuilder instantiates a new UpdatableAssetsRequestBuilder and sets the default values.
func NewUpdatableAssetsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UpdatableAssetsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUpdatableAssetsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation assets registered with the deployment service that can receive updates. Read-only.
func (m *UpdatableAssetsRequestBuilder) CreateGetRequestInformation(options *UpdatableAssetsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation assets registered with the deployment service that can receive updates. Read-only.
func (m *UpdatableAssetsRequestBuilder) CreatePostRequestInformation(options *UpdatableAssetsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *UpdatableAssetsRequestBuilder) EnrollAssets()(*i5e033ef8b15c49341c073fc21c659e5769ead616f1066db183a743d44638c27e.EnrollAssetsRequestBuilder) {
    return i5e033ef8b15c49341c073fc21c659e5769ead616f1066db183a743d44638c27e.NewEnrollAssetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UpdatableAssetsRequestBuilder) EnrollAssetsById()(*i1ad9e4f03fd41b4ff3469626c2221aacd3c3b097757de111b29a11d0e40e5af7.EnrollAssetsByIdRequestBuilder) {
    return i1ad9e4f03fd41b4ff3469626c2221aacd3c3b097757de111b29a11d0e40e5af7.NewEnrollAssetsByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get assets registered with the deployment service that can receive updates. Read-only.
func (m *UpdatableAssetsRequestBuilder) Get(options *UpdatableAssetsRequestBuilderGetOptions)(*UpdatableAssetsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUpdatableAssetsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*UpdatableAssetsResponse), nil
}
// Post assets registered with the deployment service that can receive updates. Read-only.
func (m *UpdatableAssetsRequestBuilder) Post(options *UpdatableAssetsRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UpdatableAsset, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUpdatableAsset() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UpdatableAsset), nil
}
func (m *UpdatableAssetsRequestBuilder) UnenrollAssets()(*i4dfa7ffb4b8c911e44699c4105daaccddf21f9cbf6d740540587c4fac3cb8746.UnenrollAssetsRequestBuilder) {
    return i4dfa7ffb4b8c911e44699c4105daaccddf21f9cbf6d740540587c4fac3cb8746.NewUnenrollAssetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UpdatableAssetsRequestBuilder) UnenrollAssetsById()(*i9b7b56068ae594d878d9164fec1e047566feb9a550ec5e453c176a1f0a5f2796.UnenrollAssetsByIdRequestBuilder) {
    return i9b7b56068ae594d878d9164fec1e047566feb9a550ec5e453c176a1f0a5f2796.NewUnenrollAssetsByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
