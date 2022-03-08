package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i49ec1b54a19ce17f31feaddd613b4715b173af48df3cdd9ae28ad65e94f03408 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/stop"
    i5591bb267508ecb75034e78615aaee45cc17284223ea50adf429b06c79b258f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances"
    idac0b3111695ecd4dc16b6628acb0b3c5ae374573f87c852ffc058f95d7abdab "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item"
)

// AccessReviewScheduleDefinitionItemRequestBuilder provides operations to manage the definitions property of the microsoft.graph.accessReviewSet entity.
type AccessReviewScheduleDefinitionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AccessReviewScheduleDefinitionItemRequestBuilderDeleteOptions options for Delete
type AccessReviewScheduleDefinitionItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessReviewScheduleDefinitionItemRequestBuilderGetOptions options for Get
type AccessReviewScheduleDefinitionItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AccessReviewScheduleDefinitionItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessReviewScheduleDefinitionItemRequestBuilderGetQueryParameters represents the template and scheduling for an access review.
type AccessReviewScheduleDefinitionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AccessReviewScheduleDefinitionItemRequestBuilderPatchOptions options for Patch
type AccessReviewScheduleDefinitionItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessReviewScheduleDefinitionable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewAccessReviewScheduleDefinitionItemRequestBuilderInternal instantiates a new AccessReviewScheduleDefinitionItemRequestBuilder and sets the default values.
func NewAccessReviewScheduleDefinitionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessReviewScheduleDefinitionItemRequestBuilder) {
    m := &AccessReviewScheduleDefinitionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessReviewScheduleDefinitionItemRequestBuilder instantiates a new AccessReviewScheduleDefinitionItemRequestBuilder and sets the default values.
func NewAccessReviewScheduleDefinitionItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessReviewScheduleDefinitionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewScheduleDefinitionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property definitions for identityGovernance
func (m *AccessReviewScheduleDefinitionItemRequestBuilder) CreateDeleteRequestInformation(options *AccessReviewScheduleDefinitionItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation represents the template and scheduling for an access review.
func (m *AccessReviewScheduleDefinitionItemRequestBuilder) CreateGetRequestInformation(options *AccessReviewScheduleDefinitionItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property definitions in identityGovernance
func (m *AccessReviewScheduleDefinitionItemRequestBuilder) CreatePatchRequestInformation(options *AccessReviewScheduleDefinitionItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property definitions for identityGovernance
func (m *AccessReviewScheduleDefinitionItemRequestBuilder) Delete(options *AccessReviewScheduleDefinitionItemRequestBuilderDeleteOptions)(error) {
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
// Get represents the template and scheduling for an access review.
func (m *AccessReviewScheduleDefinitionItemRequestBuilder) Get(options *AccessReviewScheduleDefinitionItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessReviewScheduleDefinitionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateAccessReviewScheduleDefinitionFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessReviewScheduleDefinitionable), nil
}
func (m *AccessReviewScheduleDefinitionItemRequestBuilder) Instances()(*i5591bb267508ecb75034e78615aaee45cc17284223ea50adf429b06c79b258f9.InstancesRequestBuilder) {
    return i5591bb267508ecb75034e78615aaee45cc17284223ea50adf429b06c79b258f9.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.accessReviews.definitions.item.instances.item collection
func (m *AccessReviewScheduleDefinitionItemRequestBuilder) InstancesById(id string)(*idac0b3111695ecd4dc16b6628acb0b3c5ae374573f87c852ffc058f95d7abdab.AccessReviewInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewInstance_id"] = id
    }
    return idac0b3111695ecd4dc16b6628acb0b3c5ae374573f87c852ffc058f95d7abdab.NewAccessReviewInstanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property definitions in identityGovernance
func (m *AccessReviewScheduleDefinitionItemRequestBuilder) Patch(options *AccessReviewScheduleDefinitionItemRequestBuilderPatchOptions)(error) {
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
func (m *AccessReviewScheduleDefinitionItemRequestBuilder) Stop()(*i49ec1b54a19ce17f31feaddd613b4715b173af48df3cdd9ae28ad65e94f03408.StopRequestBuilder) {
    return i49ec1b54a19ce17f31feaddd613b4715b173af48df3cdd9ae28ad65e94f03408.NewStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
