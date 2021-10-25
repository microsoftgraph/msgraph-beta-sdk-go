package planner

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i276465e5f4c886d6a61356d326a2f9f84555d806812828064e6703be56a5db80 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/recentplans"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ia8b4b566a909edf1d2a6703ac30b7e3b8deeee1d18d9a70abe370e3f4c39de38 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/tasks"
    ie9512dfe2baec3e9b57eeec21a36064d1ca4f6c60bcf3659f5ae7cc90af786fd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/rosterplans"
    iec266a670c9004b3548436ef84f6d344262593ebfb8a5931109f8ae25d2ed114 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/plans"
    if951d99d6353794a9a240c80019dac3c903ab51eb2b16b3ffd092788641abe57 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/favoriteplans"
    ifbcaf5cdca7ba00b1339240d3736a4bb960b1974633dd806eb7b6b176c1ff1a4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/all"
    ia18da0e36a14ee15225f84a02a84b46591fb9f7c1fe1ea17a3e0bd47b603f5de "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/all/item"
    ie8e0af54c5710406728d2f0cf561480ba34965894549bc0ef8a5a97e6516e83b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/plans/item"
    if89caed7c6cf9e56d6d5335f6c6908c367d55a0c198a5eef8530de7a51ed03f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/tasks/item"
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
func (m *PlannerRequestBuilder) All()(*ifbcaf5cdca7ba00b1339240d3736a4bb960b1974633dd806eb7b6b176c1ff1a4.AllRequestBuilder) {
    return ifbcaf5cdca7ba00b1339240d3736a4bb960b1974633dd806eb7b6b176c1ff1a4.NewAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PlannerRequestBuilder) AllById(id string)(*ia18da0e36a14ee15225f84a02a84b46591fb9f7c1fe1ea17a3e0bd47b603f5de.PlannerDeltaRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerDelta_id"] = id
    }
    return ia18da0e36a14ee15225f84a02a84b46591fb9f7c1fe1ea17a3e0bd47b603f5de.NewPlannerDeltaRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewPlannerRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PlannerRequestBuilder) {
    m := &PlannerRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/me/planner{?select,expand}";
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
func (m *PlannerRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
func (m *PlannerRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PlannerUser, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *PlannerRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *PlannerRequestBuilder) FavoritePlans()(*if951d99d6353794a9a240c80019dac3c903ab51eb2b16b3ffd092788641abe57.FavoritePlansRequestBuilder) {
    return if951d99d6353794a9a240c80019dac3c903ab51eb2b16b3ffd092788641abe57.NewFavoritePlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PlannerRequestBuilder) Get(q func (value *PlannerRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PlannerUser, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPlannerUser() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PlannerUser), nil
}
func (m *PlannerRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PlannerUser, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *PlannerRequestBuilder) Plans()(*iec266a670c9004b3548436ef84f6d344262593ebfb8a5931109f8ae25d2ed114.PlansRequestBuilder) {
    return iec266a670c9004b3548436ef84f6d344262593ebfb8a5931109f8ae25d2ed114.NewPlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PlannerRequestBuilder) PlansById(id string)(*ie8e0af54c5710406728d2f0cf561480ba34965894549bc0ef8a5a97e6516e83b.PlannerPlanRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerPlan_id"] = id
    }
    return ie8e0af54c5710406728d2f0cf561480ba34965894549bc0ef8a5a97e6516e83b.NewPlannerPlanRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PlannerRequestBuilder) RecentPlans()(*i276465e5f4c886d6a61356d326a2f9f84555d806812828064e6703be56a5db80.RecentPlansRequestBuilder) {
    return i276465e5f4c886d6a61356d326a2f9f84555d806812828064e6703be56a5db80.NewRecentPlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PlannerRequestBuilder) RosterPlans()(*ie9512dfe2baec3e9b57eeec21a36064d1ca4f6c60bcf3659f5ae7cc90af786fd.RosterPlansRequestBuilder) {
    return ie9512dfe2baec3e9b57eeec21a36064d1ca4f6c60bcf3659f5ae7cc90af786fd.NewRosterPlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PlannerRequestBuilder) Tasks()(*ia8b4b566a909edf1d2a6703ac30b7e3b8deeee1d18d9a70abe370e3f4c39de38.TasksRequestBuilder) {
    return ia8b4b566a909edf1d2a6703ac30b7e3b8deeee1d18d9a70abe370e3f4c39de38.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PlannerRequestBuilder) TasksById(id string)(*if89caed7c6cf9e56d6d5335f6c6908c367d55a0c198a5eef8530de7a51ed03f9.PlannerTaskRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerTask_id"] = id
    }
    return if89caed7c6cf9e56d6d5335f6c6908c367d55a0c198a5eef8530de7a51ed03f9.NewPlannerTaskRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
