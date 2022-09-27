package attacksimulation

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i91ea136d6db7c9ecbe5c19aaf6f9924fa0ccbcf63930c3fadff4429c1bbff683 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/attacksimulation/simulations"
    ie8a063a943fef4e082163058dd2e72338195ff1de79a742eb6b73521689f5884 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/attacksimulation/simulationautomations"
    i50ec8564a3573ab218480a77da34dfd3ec06dcd6cc87f5f81ab4cc9acfcfd9d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/attacksimulation/simulations/item"
    ieb8d3f2d145f8b20bedef89bf1b341cbd6abe71dc63c7e852454de2baf05bb5a "github.com/microsoftgraph/msgraph-beta-sdk-go/security/attacksimulation/simulationautomations/item"
)

// AttackSimulationRequestBuilder provides operations to manage the attackSimulation property of the microsoft.graph.security entity.
type AttackSimulationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AttackSimulationRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttackSimulationRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AttackSimulationRequestBuilderGetQueryParameters provides tenants capability to launch a simulated and realistic phishing attack and learn from it.
type AttackSimulationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AttackSimulationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttackSimulationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AttackSimulationRequestBuilderGetQueryParameters
}
// AttackSimulationRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttackSimulationRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAttackSimulationRequestBuilderInternal instantiates a new AttackSimulationRequestBuilder and sets the default values.
func NewAttackSimulationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttackSimulationRequestBuilder) {
    m := &AttackSimulationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/attackSimulation{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAttackSimulationRequestBuilder instantiates a new AttackSimulationRequestBuilder and sets the default values.
func NewAttackSimulationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttackSimulationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAttackSimulationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property attackSimulation for security
func (m *AttackSimulationRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property attackSimulation for security
func (m *AttackSimulationRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *AttackSimulationRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation provides tenants capability to launch a simulated and realistic phishing attack and learn from it.
func (m *AttackSimulationRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration provides tenants capability to launch a simulated and realistic phishing attack and learn from it.
func (m *AttackSimulationRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *AttackSimulationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property attackSimulation in security
func (m *AttackSimulationRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttackSimulationRootable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property attackSimulation in security
func (m *AttackSimulationRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttackSimulationRootable, requestConfiguration *AttackSimulationRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property attackSimulation for security
func (m *AttackSimulationRequestBuilder) Delete(ctx context.Context, requestConfiguration *AttackSimulationRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get provides tenants capability to launch a simulated and realistic phishing attack and learn from it.
func (m *AttackSimulationRequestBuilder) Get(ctx context.Context, requestConfiguration *AttackSimulationRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttackSimulationRootable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAttackSimulationRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttackSimulationRootable), nil
}
// Patch update the navigation property attackSimulation in security
func (m *AttackSimulationRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttackSimulationRootable, requestConfiguration *AttackSimulationRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttackSimulationRootable, error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAttackSimulationRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttackSimulationRootable), nil
}
// SimulationAutomations the simulationAutomations property
func (m *AttackSimulationRequestBuilder) SimulationAutomations()(*ie8a063a943fef4e082163058dd2e72338195ff1de79a742eb6b73521689f5884.SimulationAutomationsRequestBuilder) {
    return ie8a063a943fef4e082163058dd2e72338195ff1de79a742eb6b73521689f5884.NewSimulationAutomationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SimulationAutomationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.attackSimulation.simulationAutomations.item collection
func (m *AttackSimulationRequestBuilder) SimulationAutomationsById(id string)(*ieb8d3f2d145f8b20bedef89bf1b341cbd6abe71dc63c7e852454de2baf05bb5a.SimulationAutomationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["simulationAutomation%2Did"] = id
    }
    return ieb8d3f2d145f8b20bedef89bf1b341cbd6abe71dc63c7e852454de2baf05bb5a.NewSimulationAutomationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Simulations the simulations property
func (m *AttackSimulationRequestBuilder) Simulations()(*i91ea136d6db7c9ecbe5c19aaf6f9924fa0ccbcf63930c3fadff4429c1bbff683.SimulationsRequestBuilder) {
    return i91ea136d6db7c9ecbe5c19aaf6f9924fa0ccbcf63930c3fadff4429c1bbff683.NewSimulationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SimulationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.attackSimulation.simulations.item collection
func (m *AttackSimulationRequestBuilder) SimulationsById(id string)(*i50ec8564a3573ab218480a77da34dfd3ec06dcd6cc87f5f81ab4cc9acfcfd9d2.SimulationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["simulation%2Did"] = id
    }
    return i50ec8564a3573ab218480a77da34dfd3ec06dcd6cc87f5f81ab4cc9acfcfd9d2.NewSimulationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
