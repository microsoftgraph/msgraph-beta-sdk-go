package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i4bf83ccc76f94e75a1334b2b27a8a4d5f954c67c2ac09f3211ab9279a2c0f379 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/recommendations/item/impactedresources/item/reactivate"
    i7c01e9ec847c8cee15c0ce69a7ffc881ef74d2d10434a70d532cb3f20c36890f "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/recommendations/item/impactedresources/item/postpone"
    i8bd3a602819bab93ef30ba13423316d36438881a750ab8b4323f2e30935f8ca2 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/recommendations/item/impactedresources/item/complete"
    i91abf1f45533fe91ddeadb12b2f158073abbf92e451bc97c91cd0e45bb1421a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/recommendations/item/impactedresources/item/dismiss"
)

// RecommendationResourceRequestBuilder builds and executes requests for operations under \directory\recommendations\{recommendation-id}\impactedResources\{recommendationResource-id}
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
func (m *RecommendationResourceRequestBuilder) Complete()(*i8bd3a602819bab93ef30ba13423316d36438881a750ab8b4323f2e30935f8ca2.CompleteRequestBuilder) {
    return i8bd3a602819bab93ef30ba13423316d36438881a750ab8b4323f2e30935f8ca2.NewCompleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewRecommendationResourceRequestBuilderInternal instantiates a new RecommendationResourceRequestBuilder and sets the default values.
func NewRecommendationResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RecommendationResourceRequestBuilder) {
    m := &RecommendationResourceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory/recommendations/{recommendation_id}/impactedResources/{recommendationResource_id}{?select,expand}";
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
func (m *RecommendationResourceRequestBuilder) Dismiss()(*i91abf1f45533fe91ddeadb12b2f158073abbf92e451bc97c91cd0e45bb1421a1.DismissRequestBuilder) {
    return i91abf1f45533fe91ddeadb12b2f158073abbf92e451bc97c91cd0e45bb1421a1.NewDismissRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *RecommendationResourceRequestBuilder) Postpone()(*i7c01e9ec847c8cee15c0ce69a7ffc881ef74d2d10434a70d532cb3f20c36890f.PostponeRequestBuilder) {
    return i7c01e9ec847c8cee15c0ce69a7ffc881ef74d2d10434a70d532cb3f20c36890f.NewPostponeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *RecommendationResourceRequestBuilder) Reactivate()(*i4bf83ccc76f94e75a1334b2b27a8a4d5f954c67c2ac09f3211ab9279a2c0f379.ReactivateRequestBuilder) {
    return i4bf83ccc76f94e75a1334b2b27a8a4d5f954c67c2ac09f3211ab9279a2c0f379.NewReactivateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
