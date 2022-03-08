package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1f36147fefe41c46cb8eb01b94e9e30db2e9ecee810b25f697f038bb59fdd8ff "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/stages/item/stop"
    ie8bb93c0d5c21facddb85b340ce5d64414de2dee85b59caccca457f2e887aa82 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/stages/item/decisions"
    i4fdaa2de2f9ec85adc70f4f7f833539dcf6a8145e40b5fb2ce3f1d1317e19b22 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/stages/item/decisions/item"
)

// AccessReviewStageItemRequestBuilder provides operations to manage the stages property of the microsoft.graph.accessReviewInstance entity.
type AccessReviewStageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AccessReviewStageItemRequestBuilderDeleteOptions options for Delete
type AccessReviewStageItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessReviewStageItemRequestBuilderGetOptions options for Get
type AccessReviewStageItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AccessReviewStageItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessReviewStageItemRequestBuilderGetQueryParameters if the instance has multiple stages, this returns the collection of stages. A new stage will only be created when the previous stage ends. The existence, number, and settings of stages on a review instance are created based on the accessReviewStageSettings on the parent accessReviewScheduleDefinition.
type AccessReviewStageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AccessReviewStageItemRequestBuilderPatchOptions options for Patch
type AccessReviewStageItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessReviewStageable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewAccessReviewStageItemRequestBuilderInternal instantiates a new AccessReviewStageItemRequestBuilder and sets the default values.
func NewAccessReviewStageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessReviewStageItemRequestBuilder) {
    m := &AccessReviewStageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/pendingAccessReviewInstances/{accessReviewInstance_id}/stages/{accessReviewStage_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessReviewStageItemRequestBuilder instantiates a new AccessReviewStageItemRequestBuilder and sets the default values.
func NewAccessReviewStageItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessReviewStageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewStageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property stages for me
func (m *AccessReviewStageItemRequestBuilder) CreateDeleteRequestInformation(options *AccessReviewStageItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation if the instance has multiple stages, this returns the collection of stages. A new stage will only be created when the previous stage ends. The existence, number, and settings of stages on a review instance are created based on the accessReviewStageSettings on the parent accessReviewScheduleDefinition.
func (m *AccessReviewStageItemRequestBuilder) CreateGetRequestInformation(options *AccessReviewStageItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property stages in me
func (m *AccessReviewStageItemRequestBuilder) CreatePatchRequestInformation(options *AccessReviewStageItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *AccessReviewStageItemRequestBuilder) Decisions()(*ie8bb93c0d5c21facddb85b340ce5d64414de2dee85b59caccca457f2e887aa82.DecisionsRequestBuilder) {
    return ie8bb93c0d5c21facddb85b340ce5d64414de2dee85b59caccca457f2e887aa82.NewDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DecisionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.pendingAccessReviewInstances.item.stages.item.decisions.item collection
func (m *AccessReviewStageItemRequestBuilder) DecisionsById(id string)(*i4fdaa2de2f9ec85adc70f4f7f833539dcf6a8145e40b5fb2ce3f1d1317e19b22.AccessReviewInstanceDecisionItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewInstanceDecisionItem_id"] = id
    }
    return i4fdaa2de2f9ec85adc70f4f7f833539dcf6a8145e40b5fb2ce3f1d1317e19b22.NewAccessReviewInstanceDecisionItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property stages for me
func (m *AccessReviewStageItemRequestBuilder) Delete(options *AccessReviewStageItemRequestBuilderDeleteOptions)(error) {
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
// Get if the instance has multiple stages, this returns the collection of stages. A new stage will only be created when the previous stage ends. The existence, number, and settings of stages on a review instance are created based on the accessReviewStageSettings on the parent accessReviewScheduleDefinition.
func (m *AccessReviewStageItemRequestBuilder) Get(options *AccessReviewStageItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessReviewStageable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateAccessReviewStageFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessReviewStageable), nil
}
// Patch update the navigation property stages in me
func (m *AccessReviewStageItemRequestBuilder) Patch(options *AccessReviewStageItemRequestBuilderPatchOptions)(error) {
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
func (m *AccessReviewStageItemRequestBuilder) Stop()(*i1f36147fefe41c46cb8eb01b94e9e30db2e9ecee810b25f697f038bb59fdd8ff.StopRequestBuilder) {
    return i1f36147fefe41c46cb8eb01b94e9e30db2e9ecee810b25f697f038bb59fdd8ff.NewStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
