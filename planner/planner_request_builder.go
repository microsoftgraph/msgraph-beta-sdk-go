package planner

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i1f734a8e118043874e700fcde1d31c1b482783dc8c8d16b4b733c5c4468193a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/buckets"
    i20268753e7c3020410bc7b257943b38d7ddbd7dac155d5aa908638a484563a0b "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/tasks"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i24498e00b94127cf52e32b44bb4a7c7b55b4996fcc8d8f7fc5588dc3f25081a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/plans"
    ib18d47314351fa2c87f6ae4b1626819d77c3e1980d8f164525db440491fe7b4e "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/rosters"
    i49d7402ded8d0526c75784ca123eb70bb195f40f6fbf171738a0e3141a53d9fe "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/tasks/item"
    i8cd70227fa2a2e73750c8e4e63440317eb3ed8404fbba0976dc45ebacacd40f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/buckets/item"
    i8de8f0901d178bd52a19035fe88edf4f26f76e989740d3eb1b20e83222def4c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/plans/item"
    ia1542c83c900b2622d1e688fb4b392f0b9a5aa87a65e4fd4f6b748c84899d0af "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/rosters/item"
)

// PlannerRequestBuilder provides operations to manage the planner singleton.
type PlannerRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PlannerRequestBuilderGetQueryParameters get planner
type PlannerRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PlannerRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PlannerRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PlannerRequestBuilderGetQueryParameters
}
// PlannerRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PlannerRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Buckets the buckets property
func (m *PlannerRequestBuilder) Buckets()(*i1f734a8e118043874e700fcde1d31c1b482783dc8c8d16b4b733c5c4468193a7.BucketsRequestBuilder) {
    return i1f734a8e118043874e700fcde1d31c1b482783dc8c8d16b4b733c5c4468193a7.NewBucketsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BucketsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.planner.buckets.item collection
func (m *PlannerRequestBuilder) BucketsById(id string)(*i8cd70227fa2a2e73750c8e4e63440317eb3ed8404fbba0976dc45ebacacd40f5.PlannerBucketItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerBucket%2Did"] = id
    }
    return i8cd70227fa2a2e73750c8e4e63440317eb3ed8404fbba0976dc45ebacacd40f5.NewPlannerBucketItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPlannerRequestBuilderInternal instantiates a new PlannerRequestBuilder and sets the default values.
func NewPlannerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PlannerRequestBuilder) {
    m := &PlannerRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/planner{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPlannerRequestBuilder instantiates a new PlannerRequestBuilder and sets the default values.
func NewPlannerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PlannerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPlannerRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get planner
func (m *PlannerRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get planner
func (m *PlannerRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *PlannerRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update planner
func (m *PlannerRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Plannerable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update planner
func (m *PlannerRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Plannerable, requestConfiguration *PlannerRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get get planner
func (m *PlannerRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Plannerable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get planner
func (m *PlannerRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *PlannerRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Plannerable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Plannerable), nil
}
// Patch update planner
func (m *PlannerRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Plannerable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update planner
func (m *PlannerRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Plannerable, requestConfiguration *PlannerRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Plans the plans property
func (m *PlannerRequestBuilder) Plans()(*i24498e00b94127cf52e32b44bb4a7c7b55b4996fcc8d8f7fc5588dc3f25081a7.PlansRequestBuilder) {
    return i24498e00b94127cf52e32b44bb4a7c7b55b4996fcc8d8f7fc5588dc3f25081a7.NewPlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PlansById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.planner.plans.item collection
func (m *PlannerRequestBuilder) PlansById(id string)(*i8de8f0901d178bd52a19035fe88edf4f26f76e989740d3eb1b20e83222def4c6.PlannerPlanItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerPlan%2Did"] = id
    }
    return i8de8f0901d178bd52a19035fe88edf4f26f76e989740d3eb1b20e83222def4c6.NewPlannerPlanItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Rosters the rosters property
func (m *PlannerRequestBuilder) Rosters()(*ib18d47314351fa2c87f6ae4b1626819d77c3e1980d8f164525db440491fe7b4e.RostersRequestBuilder) {
    return ib18d47314351fa2c87f6ae4b1626819d77c3e1980d8f164525db440491fe7b4e.NewRostersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RostersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.planner.rosters.item collection
func (m *PlannerRequestBuilder) RostersById(id string)(*ia1542c83c900b2622d1e688fb4b392f0b9a5aa87a65e4fd4f6b748c84899d0af.PlannerRosterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerRoster%2Did"] = id
    }
    return ia1542c83c900b2622d1e688fb4b392f0b9a5aa87a65e4fd4f6b748c84899d0af.NewPlannerRosterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Tasks the tasks property
func (m *PlannerRequestBuilder) Tasks()(*i20268753e7c3020410bc7b257943b38d7ddbd7dac155d5aa908638a484563a0b.TasksRequestBuilder) {
    return i20268753e7c3020410bc7b257943b38d7ddbd7dac155d5aa908638a484563a0b.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TasksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.planner.tasks.item collection
func (m *PlannerRequestBuilder) TasksById(id string)(*i49d7402ded8d0526c75784ca123eb70bb195f40f6fbf171738a0e3141a53d9fe.PlannerTaskItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerTask%2Did"] = id
    }
    return i49d7402ded8d0526c75784ca123eb70bb195f40f6fbf171738a0e3141a53d9fe.NewPlannerTaskItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
