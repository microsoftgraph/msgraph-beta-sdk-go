package getattacksimulationrepeatoffenders

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GetAttackSimulationRepeatOffendersRequestBuilder builds and executes requests for operations under \reports\microsoft.graph.getAttackSimulationRepeatOffenders()
type GetAttackSimulationRepeatOffendersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetAttackSimulationRepeatOffendersRequestBuilderGetOptions options for Get
type GetAttackSimulationRepeatOffendersRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetAttackSimulationRepeatOffendersRequestBuilderInternal instantiates a new GetAttackSimulationRepeatOffendersRequestBuilder and sets the default values.
func NewGetAttackSimulationRepeatOffendersRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetAttackSimulationRepeatOffendersRequestBuilder) {
    m := &GetAttackSimulationRepeatOffendersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getAttackSimulationRepeatOffenders()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetAttackSimulationRepeatOffendersRequestBuilder instantiates a new GetAttackSimulationRepeatOffendersRequestBuilder and sets the default values.
func NewGetAttackSimulationRepeatOffendersRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetAttackSimulationRepeatOffendersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetAttackSimulationRepeatOffendersRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function getAttackSimulationRepeatOffenders
func (m *GetAttackSimulationRepeatOffendersRequestBuilder) CreateGetRequestInformation(options *GetAttackSimulationRepeatOffendersRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
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
// Get invoke function getAttackSimulationRepeatOffenders
func (m *GetAttackSimulationRepeatOffendersRequestBuilder) Get(options *GetAttackSimulationRepeatOffendersRequestBuilderGetOptions)([]GetAttackSimulationRepeatOffenders, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendCollectionAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGetAttackSimulationRepeatOffenders() }, nil, nil)
    if err != nil {
        return nil, err
    }
    val := make([]GetAttackSimulationRepeatOffenders, len(res))
    for i, v := range res {
        val[i] = *(v.(*GetAttackSimulationRepeatOffenders))
    }
    return val, nil
}
