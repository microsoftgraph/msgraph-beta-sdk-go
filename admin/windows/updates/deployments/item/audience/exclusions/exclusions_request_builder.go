package exclusions

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5c9170b81c201b5105566ee9502169d272d67357d920434087742a266367025d "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/exclusions/unenrollassets"
    i7a5ce3b3753fea9d5a62c39633c51d83d1cd2403b75643a52562b7a607be2f57 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/exclusions/enrollassetsbyid"
    ic1f0bd7dc659fdc462a2c99f69ef901e44c18a5ef5ab6d0830bbfd7789e3667d "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/exclusions/enrollassets"
    ie0c0fd2f52c8949442d7606cb6f684535237f75ae9a8505ddead286fdb640838 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/exclusions/unenrollassetsbyid"
)

// ExclusionsRequestBuilder builds and executes requests for operations under \admin\windows\updates\deployments\{deployment-id}\audience\exclusions
type ExclusionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ExclusionsRequestBuilderGetOptions options for Get
type ExclusionsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ExclusionsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ExclusionsRequestBuilderGetQueryParameters specifies the assets to exclude from the audience.
type ExclusionsRequestBuilderGetQueryParameters struct {
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
// ExclusionsRequestBuilderPostOptions options for Post
type ExclusionsRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UpdatableAsset;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewExclusionsRequestBuilderInternal instantiates a new ExclusionsRequestBuilder and sets the default values.
func NewExclusionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExclusionsRequestBuilder) {
    m := &ExclusionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/windows/updates/deployments/{deployment_id}/audience/exclusions{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewExclusionsRequestBuilder instantiates a new ExclusionsRequestBuilder and sets the default values.
func NewExclusionsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExclusionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExclusionsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation specifies the assets to exclude from the audience.
func (m *ExclusionsRequestBuilder) CreateGetRequestInformation(options *ExclusionsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation specifies the assets to exclude from the audience.
func (m *ExclusionsRequestBuilder) CreatePostRequestInformation(options *ExclusionsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ExclusionsRequestBuilder) EnrollAssets()(*ic1f0bd7dc659fdc462a2c99f69ef901e44c18a5ef5ab6d0830bbfd7789e3667d.EnrollAssetsRequestBuilder) {
    return ic1f0bd7dc659fdc462a2c99f69ef901e44c18a5ef5ab6d0830bbfd7789e3667d.NewEnrollAssetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ExclusionsRequestBuilder) EnrollAssetsById()(*i7a5ce3b3753fea9d5a62c39633c51d83d1cd2403b75643a52562b7a607be2f57.EnrollAssetsByIdRequestBuilder) {
    return i7a5ce3b3753fea9d5a62c39633c51d83d1cd2403b75643a52562b7a607be2f57.NewEnrollAssetsByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get specifies the assets to exclude from the audience.
func (m *ExclusionsRequestBuilder) Get(options *ExclusionsRequestBuilderGetOptions)(*ExclusionsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExclusionsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ExclusionsResponse), nil
}
// Post specifies the assets to exclude from the audience.
func (m *ExclusionsRequestBuilder) Post(options *ExclusionsRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UpdatableAsset, error) {
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
func (m *ExclusionsRequestBuilder) UnenrollAssets()(*i5c9170b81c201b5105566ee9502169d272d67357d920434087742a266367025d.UnenrollAssetsRequestBuilder) {
    return i5c9170b81c201b5105566ee9502169d272d67357d920434087742a266367025d.NewUnenrollAssetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ExclusionsRequestBuilder) UnenrollAssetsById()(*ie0c0fd2f52c8949442d7606cb6f684535237f75ae9a8505ddead286fdb640838.UnenrollAssetsByIdRequestBuilder) {
    return ie0c0fd2f52c8949442d7606cb6f684535237f75ae9a8505ddead286fdb640838.NewUnenrollAssetsByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
