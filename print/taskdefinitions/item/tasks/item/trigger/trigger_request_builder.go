package trigger

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    icc127f4bce7134dab72fc19d229a08e815c56f3d276ba464c4afd61792bd74af "github.com/microsoftgraph/msgraph-beta-sdk-go/print/taskdefinitions/item/tasks/item/trigger/ref"
)

// TriggerRequestBuilder builds and executes requests for operations under \print\taskDefinitions\{printTaskDefinition-id}\tasks\{printTask-id}\trigger
type TriggerRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TriggerRequestBuilderGetOptions options for Get
type TriggerRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TriggerRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TriggerRequestBuilderGetQueryParameters the printTaskTrigger that triggered this task's execution. Read-only.
type TriggerRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NewTriggerRequestBuilderInternal instantiates a new TriggerRequestBuilder and sets the default values.
func NewTriggerRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TriggerRequestBuilder) {
    m := &TriggerRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/taskDefinitions/{printTaskDefinition_id}/tasks/{printTask_id}/trigger{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTriggerRequestBuilder instantiates a new TriggerRequestBuilder and sets the default values.
func NewTriggerRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TriggerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTriggerRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the printTaskTrigger that triggered this task's execution. Read-only.
func (m *TriggerRequestBuilder) CreateGetRequestInformation(options *TriggerRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the printTaskTrigger that triggered this task's execution. Read-only.
func (m *TriggerRequestBuilder) Get(options *TriggerRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrintTaskTrigger, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPrintTaskTrigger() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrintTaskTrigger), nil
}
func (m *TriggerRequestBuilder) Ref()(*icc127f4bce7134dab72fc19d229a08e815c56f3d276ba464c4afd61792bd74af.RefRequestBuilder) {
    return icc127f4bce7134dab72fc19d229a08e815c56f3d276ba464c4afd61792bd74af.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
