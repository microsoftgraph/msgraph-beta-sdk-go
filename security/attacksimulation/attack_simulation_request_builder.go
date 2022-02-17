package attacksimulation

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i91ea136d6db7c9ecbe5c19aaf6f9924fa0ccbcf63930c3fadff4429c1bbff683 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/attacksimulation/simulations"
    ie8a063a943fef4e082163058dd2e72338195ff1de79a742eb6b73521689f5884 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/attacksimulation/simulationautomations"
    i50ec8564a3573ab218480a77da34dfd3ec06dcd6cc87f5f81ab4cc9acfcfd9d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/attacksimulation/simulations/item"
    ieb8d3f2d145f8b20bedef89bf1b341cbd6abe71dc63c7e852454de2baf05bb5a "github.com/microsoftgraph/msgraph-beta-sdk-go/security/attacksimulation/simulationautomations/item"
)

// AttackSimulationRequestBuilder builds and executes requests for operations under \security\attackSimulation
type AttackSimulationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AttackSimulationRequestBuilderDeleteOptions options for Delete
type AttackSimulationRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AttackSimulationRequestBuilderGetOptions options for Get
type AttackSimulationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AttackSimulationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AttackSimulationRequestBuilderGetQueryParameters provides tenants capability to launch a simulated and realistic phishing attack and learn from it.
type AttackSimulationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AttackSimulationRequestBuilderPatchOptions options for Patch
type AttackSimulationRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AttackSimulationRoot;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewAttackSimulationRequestBuilderInternal instantiates a new AttackSimulationRequestBuilder and sets the default values.
func NewAttackSimulationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AttackSimulationRequestBuilder) {
    m := &AttackSimulationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/attackSimulation{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAttackSimulationRequestBuilder instantiates a new AttackSimulationRequestBuilder and sets the default values.
func NewAttackSimulationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AttackSimulationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAttackSimulationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation provides tenants capability to launch a simulated and realistic phishing attack and learn from it.
func (m *AttackSimulationRequestBuilder) CreateDeleteRequestInformation(options *AttackSimulationRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation provides tenants capability to launch a simulated and realistic phishing attack and learn from it.
func (m *AttackSimulationRequestBuilder) CreateGetRequestInformation(options *AttackSimulationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation provides tenants capability to launch a simulated and realistic phishing attack and learn from it.
func (m *AttackSimulationRequestBuilder) CreatePatchRequestInformation(options *AttackSimulationRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete provides tenants capability to launch a simulated and realistic phishing attack and learn from it.
func (m *AttackSimulationRequestBuilder) Delete(options *AttackSimulationRequestBuilderDeleteOptions)(error) {
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
// Get provides tenants capability to launch a simulated and realistic phishing attack and learn from it.
func (m *AttackSimulationRequestBuilder) Get(options *AttackSimulationRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AttackSimulationRoot, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAttackSimulationRoot() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AttackSimulationRoot), nil
}
// Patch provides tenants capability to launch a simulated and realistic phishing attack and learn from it.
func (m *AttackSimulationRequestBuilder) Patch(options *AttackSimulationRequestBuilderPatchOptions)(error) {
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
func (m *AttackSimulationRequestBuilder) SimulationAutomations()(*ie8a063a943fef4e082163058dd2e72338195ff1de79a742eb6b73521689f5884.SimulationAutomationsRequestBuilder) {
    return ie8a063a943fef4e082163058dd2e72338195ff1de79a742eb6b73521689f5884.NewSimulationAutomationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SimulationAutomationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.attackSimulation.simulationAutomations.item collection
func (m *AttackSimulationRequestBuilder) SimulationAutomationsById(id string)(*ieb8d3f2d145f8b20bedef89bf1b341cbd6abe71dc63c7e852454de2baf05bb5a.SimulationAutomationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["simulationAutomation_id"] = id
    }
    return ieb8d3f2d145f8b20bedef89bf1b341cbd6abe71dc63c7e852454de2baf05bb5a.NewSimulationAutomationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AttackSimulationRequestBuilder) Simulations()(*i91ea136d6db7c9ecbe5c19aaf6f9924fa0ccbcf63930c3fadff4429c1bbff683.SimulationsRequestBuilder) {
    return i91ea136d6db7c9ecbe5c19aaf6f9924fa0ccbcf63930c3fadff4429c1bbff683.NewSimulationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SimulationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.attackSimulation.simulations.item collection
func (m *AttackSimulationRequestBuilder) SimulationsById(id string)(*i50ec8564a3573ab218480a77da34dfd3ec06dcd6cc87f5f81ab4cc9acfcfd9d2.SimulationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["simulation_id"] = id
    }
    return i50ec8564a3573ab218480a77da34dfd3ec06dcd6cc87f5f81ab4cc9acfcfd9d2.NewSimulationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
