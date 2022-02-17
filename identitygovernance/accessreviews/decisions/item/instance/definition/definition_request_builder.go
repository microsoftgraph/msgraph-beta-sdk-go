package definition

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0ce831bd8561ac417fd6f28ccd4c8e38d75efb866b4e0021bb09fd2f5c9034ed "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/decisions/item/instance/definition/stop"
    ib778cc87f434c3881b79a310c65f0b500c09df529f1f6a44e344191dbc6d2c53 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/decisions/item/instance/definition/ref"
)

// DefinitionRequestBuilder builds and executes requests for operations under \identityGovernance\accessReviews\decisions\{accessReviewInstanceDecisionItem-id}\instance\definition
type DefinitionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DefinitionRequestBuilderGetOptions options for Get
type DefinitionRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DefinitionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DefinitionRequestBuilderGetQueryParameters there is exactly one accessReviewScheduleDefinition associated with each instance. It is the parent schedule for the instance, where instances are created for each recurrence of a review definition and each group selected to review by the definition.
type DefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NewDefinitionRequestBuilderInternal instantiates a new DefinitionRequestBuilder and sets the default values.
func NewDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DefinitionRequestBuilder) {
    m := &DefinitionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/accessReviews/decisions/{accessReviewInstanceDecisionItem_id}/instance/definition{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDefinitionRequestBuilder instantiates a new DefinitionRequestBuilder and sets the default values.
func NewDefinitionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation there is exactly one accessReviewScheduleDefinition associated with each instance. It is the parent schedule for the instance, where instances are created for each recurrence of a review definition and each group selected to review by the definition.
func (m *DefinitionRequestBuilder) CreateGetRequestInformation(options *DefinitionRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get there is exactly one accessReviewScheduleDefinition associated with each instance. It is the parent schedule for the instance, where instances are created for each recurrence of a review definition and each group selected to review by the definition.
func (m *DefinitionRequestBuilder) Get(options *DefinitionRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessReviewScheduleDefinition, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAccessReviewScheduleDefinition() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessReviewScheduleDefinition), nil
}
func (m *DefinitionRequestBuilder) Ref()(*ib778cc87f434c3881b79a310c65f0b500c09df529f1f6a44e344191dbc6d2c53.RefRequestBuilder) {
    return ib778cc87f434c3881b79a310c65f0b500c09df529f1f6a44e344191dbc6d2c53.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DefinitionRequestBuilder) Stop()(*i0ce831bd8561ac417fd6f28ccd4c8e38d75efb866b4e0021bb09fd2f5c9034ed.StopRequestBuilder) {
    return i0ce831bd8561ac417fd6f28ccd4c8e38d75efb866b4e0021bb09fd2f5c9034ed.NewStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
