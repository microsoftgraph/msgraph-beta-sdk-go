package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i67aaa727d73c164087ea2ce778b06cfa6a2637659c79b79f9aa2a76313e2af81 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/decisions/item/insights"
    ifb44db4feb17e6296767ce4346bc33d57927ca8f62aa8ceadf20daf20778e8cc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/decisions/item/instance"
    ie44e1255b972f859503eb871261f069d2ff1e5a0b64d5c67e3a2e405706b2c36 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/decisions/item/insights/item"
)

// AccessReviewInstanceDecisionItemRequestBuilder builds and executes requests for operations under \me\pendingAccessReviewInstances\{accessReviewInstance-id}\decisions\{accessReviewInstanceDecisionItem-id}
type AccessReviewInstanceDecisionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AccessReviewInstanceDecisionItemRequestBuilderDeleteOptions options for Delete
type AccessReviewInstanceDecisionItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessReviewInstanceDecisionItemRequestBuilderGetOptions options for Get
type AccessReviewInstanceDecisionItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AccessReviewInstanceDecisionItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessReviewInstanceDecisionItemRequestBuilderGetQueryParameters each principal reviewed in an accessReviewInstance has a decision item representing if they were approved, denied, or not yet reviewed.
type AccessReviewInstanceDecisionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AccessReviewInstanceDecisionItemRequestBuilderPatchOptions options for Patch
type AccessReviewInstanceDecisionItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessReviewInstanceDecisionItem;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewAccessReviewInstanceDecisionItemRequestBuilderInternal instantiates a new AccessReviewInstanceDecisionItemRequestBuilder and sets the default values.
func NewAccessReviewInstanceDecisionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessReviewInstanceDecisionItemRequestBuilder) {
    m := &AccessReviewInstanceDecisionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/pendingAccessReviewInstances/{accessReviewInstance_id}/decisions/{accessReviewInstanceDecisionItem_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessReviewInstanceDecisionItemRequestBuilder instantiates a new AccessReviewInstanceDecisionItemRequestBuilder and sets the default values.
func NewAccessReviewInstanceDecisionItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessReviewInstanceDecisionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewInstanceDecisionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation each principal reviewed in an accessReviewInstance has a decision item representing if they were approved, denied, or not yet reviewed.
func (m *AccessReviewInstanceDecisionItemRequestBuilder) CreateDeleteRequestInformation(options *AccessReviewInstanceDecisionItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation each principal reviewed in an accessReviewInstance has a decision item representing if they were approved, denied, or not yet reviewed.
func (m *AccessReviewInstanceDecisionItemRequestBuilder) CreateGetRequestInformation(options *AccessReviewInstanceDecisionItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation each principal reviewed in an accessReviewInstance has a decision item representing if they were approved, denied, or not yet reviewed.
func (m *AccessReviewInstanceDecisionItemRequestBuilder) CreatePatchRequestInformation(options *AccessReviewInstanceDecisionItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete each principal reviewed in an accessReviewInstance has a decision item representing if they were approved, denied, or not yet reviewed.
func (m *AccessReviewInstanceDecisionItemRequestBuilder) Delete(options *AccessReviewInstanceDecisionItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get each principal reviewed in an accessReviewInstance has a decision item representing if they were approved, denied, or not yet reviewed.
func (m *AccessReviewInstanceDecisionItemRequestBuilder) Get(options *AccessReviewInstanceDecisionItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessReviewInstanceDecisionItem, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAccessReviewInstanceDecisionItem() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessReviewInstanceDecisionItem), nil
}
func (m *AccessReviewInstanceDecisionItemRequestBuilder) Insights()(*i67aaa727d73c164087ea2ce778b06cfa6a2637659c79b79f9aa2a76313e2af81.InsightsRequestBuilder) {
    return i67aaa727d73c164087ea2ce778b06cfa6a2637659c79b79f9aa2a76313e2af81.NewInsightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InsightsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.pendingAccessReviewInstances.item.decisions.item.insights.item collection
func (m *AccessReviewInstanceDecisionItemRequestBuilder) InsightsById(id string)(*ie44e1255b972f859503eb871261f069d2ff1e5a0b64d5c67e3a2e405706b2c36.GovernanceInsightRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceInsight_id"] = id
    }
    return ie44e1255b972f859503eb871261f069d2ff1e5a0b64d5c67e3a2e405706b2c36.NewGovernanceInsightRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessReviewInstanceDecisionItemRequestBuilder) Instance()(*ifb44db4feb17e6296767ce4346bc33d57927ca8f62aa8ceadf20daf20778e8cc.InstanceRequestBuilder) {
    return ifb44db4feb17e6296767ce4346bc33d57927ca8f62aa8ceadf20daf20778e8cc.NewInstanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch each principal reviewed in an accessReviewInstance has a decision item representing if they were approved, denied, or not yet reviewed.
func (m *AccessReviewInstanceDecisionItemRequestBuilder) Patch(options *AccessReviewInstanceDecisionItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
