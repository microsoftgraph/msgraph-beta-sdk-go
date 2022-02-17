package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i02454f4537287b6ab338217acb9968bd694788945e44d84da9b93503fb407a4e "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/impactedresources/item/dismiss"
    i6a908001799100452fb9272829db83866d7952f0456ff214df6a4849088023c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/impactedresources/item/complete"
    i9a2ce2cc07814668197ce7e2daba99556437bce33769c58c3fa88ff13ad02bdd "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/impactedresources/item/reactivate"
    ie32cb720402ff9dd0b35e1e591941d01f5c513d30bea76e934259b963be0b443 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/impactedresources/item/postpone"
)

// RecommendationResourceRequestBuilder builds and executes requests for operations under \directory\impactedResources\{recommendationResource-id}
type RecommendationResourceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RecommendationResourceRequestBuilderDeleteOptions options for Delete
type RecommendationResourceRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RecommendationResourceRequestBuilderGetOptions options for Get
type RecommendationResourceRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RecommendationResourceRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RecommendationResourceRequestBuilderGetQueryParameters get impactedResources from directory
type RecommendationResourceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// RecommendationResourceRequestBuilderPatchOptions options for Patch
type RecommendationResourceRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RecommendationResource;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *RecommendationResourceRequestBuilder) Complete()(*i6a908001799100452fb9272829db83866d7952f0456ff214df6a4849088023c6.CompleteRequestBuilder) {
    return i6a908001799100452fb9272829db83866d7952f0456ff214df6a4849088023c6.NewCompleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewRecommendationResourceRequestBuilderInternal instantiates a new RecommendationResourceRequestBuilder and sets the default values.
func NewRecommendationResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RecommendationResourceRequestBuilder) {
    m := &RecommendationResourceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory/impactedResources/{recommendationResource_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRecommendationResourceRequestBuilder instantiates a new RecommendationResourceRequestBuilder and sets the default values.
func NewRecommendationResourceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RecommendationResourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRecommendationResourceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property impactedResources for directory
func (m *RecommendationResourceRequestBuilder) CreateDeleteRequestInformation(options *RecommendationResourceRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get impactedResources from directory
func (m *RecommendationResourceRequestBuilder) CreateGetRequestInformation(options *RecommendationResourceRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property impactedResources in directory
func (m *RecommendationResourceRequestBuilder) CreatePatchRequestInformation(options *RecommendationResourceRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property impactedResources for directory
func (m *RecommendationResourceRequestBuilder) Delete(options *RecommendationResourceRequestBuilderDeleteOptions)(error) {
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
func (m *RecommendationResourceRequestBuilder) Dismiss()(*i02454f4537287b6ab338217acb9968bd694788945e44d84da9b93503fb407a4e.DismissRequestBuilder) {
    return i02454f4537287b6ab338217acb9968bd694788945e44d84da9b93503fb407a4e.NewDismissRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get impactedResources from directory
func (m *RecommendationResourceRequestBuilder) Get(options *RecommendationResourceRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RecommendationResource, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewRecommendationResource() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RecommendationResource), nil
}
// Patch update the navigation property impactedResources in directory
func (m *RecommendationResourceRequestBuilder) Patch(options *RecommendationResourceRequestBuilderPatchOptions)(error) {
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
func (m *RecommendationResourceRequestBuilder) Postpone()(*ie32cb720402ff9dd0b35e1e591941d01f5c513d30bea76e934259b963be0b443.PostponeRequestBuilder) {
    return ie32cb720402ff9dd0b35e1e591941d01f5c513d30bea76e934259b963be0b443.NewPostponeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *RecommendationResourceRequestBuilder) Reactivate()(*i9a2ce2cc07814668197ce7e2daba99556437bce33769c58c3fa88ff13ad02bdd.ReactivateRequestBuilder) {
    return i9a2ce2cc07814668197ce7e2daba99556437bce33769c58c3fa88ff13ad02bdd.NewReactivateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
