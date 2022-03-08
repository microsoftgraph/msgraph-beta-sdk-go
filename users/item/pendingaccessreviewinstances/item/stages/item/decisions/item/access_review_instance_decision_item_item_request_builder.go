package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i81796c9b3bb388e047ef193315c701449a8c5498b978a31cdb67ac4a042c2efc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/stages/item/decisions/item/instance"
    i9726a71f15a35a2afae0754bc167fdf83fa607f49db40224a5399e00e09867a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/stages/item/decisions/item/insights"
    i81b099f2c33c2f064f697d68267f5f2e473d630189d67e90d2fad0a391d3fcb1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/stages/item/decisions/item/insights/item"
)

// AccessReviewInstanceDecisionItemItemRequestBuilder provides operations to manage the decisions property of the microsoft.graph.accessReviewStage entity.
type AccessReviewInstanceDecisionItemItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AccessReviewInstanceDecisionItemItemRequestBuilderDeleteOptions options for Delete
type AccessReviewInstanceDecisionItemItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessReviewInstanceDecisionItemItemRequestBuilderGetOptions options for Get
type AccessReviewInstanceDecisionItemItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AccessReviewInstanceDecisionItemItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessReviewInstanceDecisionItemItemRequestBuilderGetQueryParameters each user reviewed in an accessReviewStage has a decision item representing if they were approved, denied, or not yet reviewed.
type AccessReviewInstanceDecisionItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AccessReviewInstanceDecisionItemItemRequestBuilderPatchOptions options for Patch
type AccessReviewInstanceDecisionItemItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessReviewInstanceDecisionItemable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewAccessReviewInstanceDecisionItemItemRequestBuilderInternal instantiates a new AccessReviewInstanceDecisionItemItemRequestBuilder and sets the default values.
func NewAccessReviewInstanceDecisionItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessReviewInstanceDecisionItemItemRequestBuilder) {
    m := &AccessReviewInstanceDecisionItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/pendingAccessReviewInstances/{accessReviewInstance_id}/stages/{accessReviewStage_id}/decisions/{accessReviewInstanceDecisionItem_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessReviewInstanceDecisionItemItemRequestBuilder instantiates a new AccessReviewInstanceDecisionItemItemRequestBuilder and sets the default values.
func NewAccessReviewInstanceDecisionItemItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessReviewInstanceDecisionItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewInstanceDecisionItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property decisions for users
func (m *AccessReviewInstanceDecisionItemItemRequestBuilder) CreateDeleteRequestInformation(options *AccessReviewInstanceDecisionItemItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation each user reviewed in an accessReviewStage has a decision item representing if they were approved, denied, or not yet reviewed.
func (m *AccessReviewInstanceDecisionItemItemRequestBuilder) CreateGetRequestInformation(options *AccessReviewInstanceDecisionItemItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property decisions in users
func (m *AccessReviewInstanceDecisionItemItemRequestBuilder) CreatePatchRequestInformation(options *AccessReviewInstanceDecisionItemItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property decisions for users
func (m *AccessReviewInstanceDecisionItemItemRequestBuilder) Delete(options *AccessReviewInstanceDecisionItemItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get each user reviewed in an accessReviewStage has a decision item representing if they were approved, denied, or not yet reviewed.
func (m *AccessReviewInstanceDecisionItemItemRequestBuilder) Get(options *AccessReviewInstanceDecisionItemItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessReviewInstanceDecisionItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateAccessReviewInstanceDecisionItemFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessReviewInstanceDecisionItemable), nil
}
func (m *AccessReviewInstanceDecisionItemItemRequestBuilder) Insights()(*i9726a71f15a35a2afae0754bc167fdf83fa607f49db40224a5399e00e09867a5.InsightsRequestBuilder) {
    return i9726a71f15a35a2afae0754bc167fdf83fa607f49db40224a5399e00e09867a5.NewInsightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InsightsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.pendingAccessReviewInstances.item.stages.item.decisions.item.insights.item collection
func (m *AccessReviewInstanceDecisionItemItemRequestBuilder) InsightsById(id string)(*i81b099f2c33c2f064f697d68267f5f2e473d630189d67e90d2fad0a391d3fcb1.GovernanceInsightItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceInsight_id"] = id
    }
    return i81b099f2c33c2f064f697d68267f5f2e473d630189d67e90d2fad0a391d3fcb1.NewGovernanceInsightItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessReviewInstanceDecisionItemItemRequestBuilder) Instance()(*i81796c9b3bb388e047ef193315c701449a8c5498b978a31cdb67ac4a042c2efc.InstanceRequestBuilder) {
    return i81796c9b3bb388e047ef193315c701449a8c5498b978a31cdb67ac4a042c2efc.NewInstanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property decisions in users
func (m *AccessReviewInstanceDecisionItemItemRequestBuilder) Patch(options *AccessReviewInstanceDecisionItemItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
