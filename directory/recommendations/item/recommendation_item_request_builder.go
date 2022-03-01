package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i4018c02f3fbf5218282cea136e679104019a4c34f3c1902f574244076c8fb432 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/recommendations/item/complete"
    i85047cb55fafd88055bcdeb243251139d17240509426cf420901b907aa3f73da "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/recommendations/item/reactivate"
    i9159a081c2a0d0abecb435cb8c52d12d3cb4c5cc3be944e4400678fa5eaf97f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/recommendations/item/impactedresources"
    ib1bc9d8bdb20f875e5c28e8e4d42895f9148cf2ba25fb010132a667a01115017 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/recommendations/item/postpone"
    ib872c308280e73f616388d32a30c0860fa7b0ee0f18b5fafe036548766c8a60a "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/recommendations/item/dismiss"
    i15d57ca0f29006194bb50f1f3633607a5d2e31816ad92ef8c6d5faf215b3732b "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/recommendations/item/impactedresources/item"
)

// RecommendationItemRequestBuilder builds and executes requests for operations under \directory\recommendations\{recommendation-id}
type RecommendationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RecommendationItemRequestBuilderDeleteOptions options for Delete
type RecommendationItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RecommendationItemRequestBuilderGetOptions options for Get
type RecommendationItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RecommendationItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RecommendationItemRequestBuilderGetQueryParameters get recommendations from directory
type RecommendationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// RecommendationItemRequestBuilderPatchOptions options for Patch
type RecommendationItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recommendation;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *RecommendationItemRequestBuilder) Complete()(*i4018c02f3fbf5218282cea136e679104019a4c34f3c1902f574244076c8fb432.CompleteRequestBuilder) {
    return i4018c02f3fbf5218282cea136e679104019a4c34f3c1902f574244076c8fb432.NewCompleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewRecommendationItemRequestBuilderInternal instantiates a new RecommendationItemRequestBuilder and sets the default values.
func NewRecommendationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RecommendationItemRequestBuilder) {
    m := &RecommendationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory/recommendations/{recommendation_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRecommendationItemRequestBuilder instantiates a new RecommendationItemRequestBuilder and sets the default values.
func NewRecommendationItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RecommendationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRecommendationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property recommendations for directory
func (m *RecommendationItemRequestBuilder) CreateDeleteRequestInformation(options *RecommendationItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get recommendations from directory
func (m *RecommendationItemRequestBuilder) CreateGetRequestInformation(options *RecommendationItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property recommendations in directory
func (m *RecommendationItemRequestBuilder) CreatePatchRequestInformation(options *RecommendationItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property recommendations for directory
func (m *RecommendationItemRequestBuilder) Delete(options *RecommendationItemRequestBuilderDeleteOptions)(error) {
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
func (m *RecommendationItemRequestBuilder) Dismiss()(*ib872c308280e73f616388d32a30c0860fa7b0ee0f18b5fafe036548766c8a60a.DismissRequestBuilder) {
    return ib872c308280e73f616388d32a30c0860fa7b0ee0f18b5fafe036548766c8a60a.NewDismissRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get recommendations from directory
func (m *RecommendationItemRequestBuilder) Get(options *RecommendationItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recommendation, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewRecommendation() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Recommendation), nil
}
func (m *RecommendationItemRequestBuilder) ImpactedResources()(*i9159a081c2a0d0abecb435cb8c52d12d3cb4c5cc3be944e4400678fa5eaf97f6.ImpactedResourcesRequestBuilder) {
    return i9159a081c2a0d0abecb435cb8c52d12d3cb4c5cc3be944e4400678fa5eaf97f6.NewImpactedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ImpactedResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.recommendations.item.impactedResources.item collection
func (m *RecommendationItemRequestBuilder) ImpactedResourcesById(id string)(*i15d57ca0f29006194bb50f1f3633607a5d2e31816ad92ef8c6d5faf215b3732b.RecommendationResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["recommendationResource_id"] = id
    }
    return i15d57ca0f29006194bb50f1f3633607a5d2e31816ad92ef8c6d5faf215b3732b.NewRecommendationResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property recommendations in directory
func (m *RecommendationItemRequestBuilder) Patch(options *RecommendationItemRequestBuilderPatchOptions)(error) {
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
func (m *RecommendationItemRequestBuilder) Postpone()(*ib1bc9d8bdb20f875e5c28e8e4d42895f9148cf2ba25fb010132a667a01115017.PostponeRequestBuilder) {
    return ib1bc9d8bdb20f875e5c28e8e4d42895f9148cf2ba25fb010132a667a01115017.NewPostponeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *RecommendationItemRequestBuilder) Reactivate()(*i85047cb55fafd88055bcdeb243251139d17240509426cf420901b907aa3f73da.ReactivateRequestBuilder) {
    return i85047cb55fafd88055bcdeb243251139d17240509426cf420901b907aa3f73da.NewReactivateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
