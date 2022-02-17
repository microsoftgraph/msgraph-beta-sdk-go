package decisions

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i269f88743cf372690457151a8581e751ffc2ef68e31ca0c2be26aa553ab5995b "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/decisions/item/instance/stages/item/decisions/recordalldecisions"
    ib076ef8da31438a323cb2984e4d69eba725d2020869372899ce16a032d996868 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/decisions/item/instance/stages/item/decisions/filterbycurrentuserwithon"
)

// DecisionsRequestBuilder builds and executes requests for operations under \identityGovernance\accessReviews\decisions\{accessReviewInstanceDecisionItem-id}\instance\stages\{accessReviewStage-id}\decisions
type DecisionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DecisionsRequestBuilderGetOptions options for Get
type DecisionsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DecisionsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DecisionsRequestBuilderGetQueryParameters each user reviewed in an accessReviewStage has a decision item representing if they were approved, denied, or not yet reviewed.
type DecisionsRequestBuilderGetQueryParameters struct {
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
// DecisionsRequestBuilderPostOptions options for Post
type DecisionsRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessReviewInstanceDecisionItem;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDecisionsRequestBuilderInternal instantiates a new DecisionsRequestBuilder and sets the default values.
func NewDecisionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DecisionsRequestBuilder) {
    m := &DecisionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/accessReviews/decisions/{accessReviewInstanceDecisionItem_id}/instance/stages/{accessReviewStage_id}/decisions{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDecisionsRequestBuilder instantiates a new DecisionsRequestBuilder and sets the default values.
func NewDecisionsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DecisionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDecisionsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation each user reviewed in an accessReviewStage has a decision item representing if they were approved, denied, or not yet reviewed.
func (m *DecisionsRequestBuilder) CreateGetRequestInformation(options *DecisionsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation each user reviewed in an accessReviewStage has a decision item representing if they were approved, denied, or not yet reviewed.
func (m *DecisionsRequestBuilder) CreatePostRequestInformation(options *DecisionsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// FilterByCurrentUserWithOn builds and executes requests for operations under \identityGovernance\accessReviews\decisions\{accessReviewInstanceDecisionItem-id}\instance\stages\{accessReviewStage-id}\decisions\microsoft.graph.filterByCurrentUser(on={on})
func (m *DecisionsRequestBuilder) FilterByCurrentUserWithOn(on *string)(*ib076ef8da31438a323cb2984e4d69eba725d2020869372899ce16a032d996868.FilterByCurrentUserWithOnRequestBuilder) {
    return ib076ef8da31438a323cb2984e4d69eba725d2020869372899ce16a032d996868.NewFilterByCurrentUserWithOnRequestBuilderInternal(m.pathParameters, m.requestAdapter, on);
}
// Get each user reviewed in an accessReviewStage has a decision item representing if they were approved, denied, or not yet reviewed.
func (m *DecisionsRequestBuilder) Get(options *DecisionsRequestBuilderGetOptions)(*DecisionsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDecisionsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*DecisionsResponse), nil
}
// Post each user reviewed in an accessReviewStage has a decision item representing if they were approved, denied, or not yet reviewed.
func (m *DecisionsRequestBuilder) Post(options *DecisionsRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessReviewInstanceDecisionItem, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAccessReviewInstanceDecisionItem() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessReviewInstanceDecisionItem), nil
}
func (m *DecisionsRequestBuilder) RecordAllDecisions()(*i269f88743cf372690457151a8581e751ffc2ef68e31ca0c2be26aa553ab5995b.RecordAllDecisionsRequestBuilder) {
    return i269f88743cf372690457151a8581e751ffc2ef68e31ca0c2be26aa553ab5995b.NewRecordAllDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
