package planner

import (
    i1f734a8e118043874e700fcde1d31c1b482783dc8c8d16b4b733c5c4468193a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/buckets"
    i20268753e7c3020410bc7b257943b38d7ddbd7dac155d5aa908638a484563a0b "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/tasks"
    i24498e00b94127cf52e32b44bb4a7c7b55b4996fcc8d8f7fc5588dc3f25081a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/plans"
    ib18d47314351fa2c87f6ae4b1626819d77c3e1980d8f164525db440491fe7b4e "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/rosters"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i49d7402ded8d0526c75784ca123eb70bb195f40f6fbf171738a0e3141a53d9fe "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/tasks/item"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i8cd70227fa2a2e73750c8e4e63440317eb3ed8404fbba0976dc45ebacacd40f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/buckets/item"
    i8de8f0901d178bd52a19035fe88edf4f26f76e989740d3eb1b20e83222def4c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/plans/item"
    ia1542c83c900b2622d1e688fb4b392f0b9a5aa87a65e4fd4f6b748c84899d0af "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/rosters/item"
)

type PlannerRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type PlannerRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *PlannerRequestBuilder) Buckets()(*i1f734a8e118043874e700fcde1d31c1b482783dc8c8d16b4b733c5c4468193a7.BucketsRequestBuilder) {
    return i1f734a8e118043874e700fcde1d31c1b482783dc8c8d16b4b733c5c4468193a7.NewBucketsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PlannerRequestBuilder) BucketsById(id string)(*i8cd70227fa2a2e73750c8e4e63440317eb3ed8404fbba0976dc45ebacacd40f5.PlannerBucketRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerBucket_id"] = id
    }
    return i8cd70227fa2a2e73750c8e4e63440317eb3ed8404fbba0976dc45ebacacd40f5.NewPlannerBucketRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewPlannerRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PlannerRequestBuilder) {
    m := &PlannerRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/planner{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewPlannerRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PlannerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPlannerRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *PlannerRequestBuilder) CreateGetRequestInformation(q func (value *PlannerRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(PlannerRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *PlannerRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Planner, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *PlannerRequestBuilder) Get(q func (value *PlannerRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Planner, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPlanner() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Planner), nil
}
func (m *PlannerRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Planner, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *PlannerRequestBuilder) Plans()(*i24498e00b94127cf52e32b44bb4a7c7b55b4996fcc8d8f7fc5588dc3f25081a7.PlansRequestBuilder) {
    return i24498e00b94127cf52e32b44bb4a7c7b55b4996fcc8d8f7fc5588dc3f25081a7.NewPlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PlannerRequestBuilder) PlansById(id string)(*i8de8f0901d178bd52a19035fe88edf4f26f76e989740d3eb1b20e83222def4c6.PlannerPlanRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerPlan_id"] = id
    }
    return i8de8f0901d178bd52a19035fe88edf4f26f76e989740d3eb1b20e83222def4c6.NewPlannerPlanRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PlannerRequestBuilder) Rosters()(*ib18d47314351fa2c87f6ae4b1626819d77c3e1980d8f164525db440491fe7b4e.RostersRequestBuilder) {
    return ib18d47314351fa2c87f6ae4b1626819d77c3e1980d8f164525db440491fe7b4e.NewRostersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PlannerRequestBuilder) RostersById(id string)(*ia1542c83c900b2622d1e688fb4b392f0b9a5aa87a65e4fd4f6b748c84899d0af.PlannerRosterRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerRoster_id"] = id
    }
    return ia1542c83c900b2622d1e688fb4b392f0b9a5aa87a65e4fd4f6b748c84899d0af.NewPlannerRosterRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PlannerRequestBuilder) Tasks()(*i20268753e7c3020410bc7b257943b38d7ddbd7dac155d5aa908638a484563a0b.TasksRequestBuilder) {
    return i20268753e7c3020410bc7b257943b38d7ddbd7dac155d5aa908638a484563a0b.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PlannerRequestBuilder) TasksById(id string)(*i49d7402ded8d0526c75784ca123eb70bb195f40f6fbf171738a0e3141a53d9fe.PlannerTaskRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerTask_id"] = id
    }
    return i49d7402ded8d0526c75784ca123eb70bb195f40f6fbf171738a0e3141a53d9fe.NewPlannerTaskRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
