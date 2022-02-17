package activatedusing

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i777e57c5945f707956e43c0d527f28b108efe2544d22ed8e6adda464c97dbed4 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/roleassignmentscheduleinstances/item/activatedusing/ref"
)

// ActivatedUsingRequestBuilder builds and executes requests for operations under \roleManagement\directory\roleAssignmentScheduleInstances\{unifiedRoleAssignmentScheduleInstance-id}\activatedUsing
type ActivatedUsingRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ActivatedUsingRequestBuilderGetOptions options for Get
type ActivatedUsingRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ActivatedUsingRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ActivatedUsingRequestBuilderGetQueryParameters if the roleAssignmentScheduleInstance is activated by a roleEligibilityScheduleRequest, this is the link to the related schedule instance.
type ActivatedUsingRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NewActivatedUsingRequestBuilderInternal instantiates a new ActivatedUsingRequestBuilder and sets the default values.
func NewActivatedUsingRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ActivatedUsingRequestBuilder) {
    m := &ActivatedUsingRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/directory/roleAssignmentScheduleInstances/{unifiedRoleAssignmentScheduleInstance_id}/activatedUsing{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewActivatedUsingRequestBuilder instantiates a new ActivatedUsingRequestBuilder and sets the default values.
func NewActivatedUsingRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ActivatedUsingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewActivatedUsingRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation if the roleAssignmentScheduleInstance is activated by a roleEligibilityScheduleRequest, this is the link to the related schedule instance.
func (m *ActivatedUsingRequestBuilder) CreateGetRequestInformation(options *ActivatedUsingRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get if the roleAssignmentScheduleInstance is activated by a roleEligibilityScheduleRequest, this is the link to the related schedule instance.
func (m *ActivatedUsingRequestBuilder) Get(options *ActivatedUsingRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleEligibilityScheduleInstance, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUnifiedRoleEligibilityScheduleInstance() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleEligibilityScheduleInstance), nil
}
func (m *ActivatedUsingRequestBuilder) Ref()(*i777e57c5945f707956e43c0d527f28b108efe2544d22ed8e6adda464c97dbed4.RefRequestBuilder) {
    return i777e57c5945f707956e43c0d527f28b108efe2544d22ed8e6adda464c97dbed4.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
