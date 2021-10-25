package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i015798be98f5354e44d25cd3fc8c761a04c2de377021e2334f9f574c0b9d9d50 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/connectorgroups"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4c54046d8ca87c2d4d86246e027b33b58926116766af2be55ea00fe765476f06 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/agents"
    i4f4a2aad6de6e818a6b03ab06593dcd209fd37f696b126badf1ccf2cbb487d98 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/agentgroups"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i86c4967a04a253d001f5fd0f5c1dd90bb7ff7141cd43dab8638d3f8ddaff0d58 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/connectors"
    if08fda1a25bbc0d180dcba2e00721696a53ac148a9d4209b397df9ba1382eb6f "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/publishedresources"
    i04768df19176575e14a7cad44b0f47526794459b68db53ab71668e3192948132 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/publishedresources/item"
    i4e613208e03a8604c73cff8fea0c650341897927012dab6031821bf1b1ae2305 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/connectorgroups/item"
    i50192dbecdd2d205018f1c9f0d43e5faa1231c597c1e3d9879fb3bfcceb8f7bf "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/agents/item"
    i7be5eb505c9819935bb7ecd0914ce542aff26ebb696b731e5bb2ef462b429317 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/agentgroups/item"
    ia6c25f4c749e8beb990221d7c4363a8d4383190388b2db9197c94a520cdfffde "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/connectors/item"
)

type OnPremisesPublishingProfileRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type OnPremisesPublishingProfileRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *OnPremisesPublishingProfileRequestBuilder) AgentGroups()(*i4f4a2aad6de6e818a6b03ab06593dcd209fd37f696b126badf1ccf2cbb487d98.AgentGroupsRequestBuilder) {
    return i4f4a2aad6de6e818a6b03ab06593dcd209fd37f696b126badf1ccf2cbb487d98.NewAgentGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnPremisesPublishingProfileRequestBuilder) AgentGroupsById(id string)(*i7be5eb505c9819935bb7ecd0914ce542aff26ebb696b731e5bb2ef462b429317.OnPremisesAgentGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onPremisesAgentGroup_id"] = id
    }
    return i7be5eb505c9819935bb7ecd0914ce542aff26ebb696b731e5bb2ef462b429317.NewOnPremisesAgentGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnPremisesPublishingProfileRequestBuilder) Agents()(*i4c54046d8ca87c2d4d86246e027b33b58926116766af2be55ea00fe765476f06.AgentsRequestBuilder) {
    return i4c54046d8ca87c2d4d86246e027b33b58926116766af2be55ea00fe765476f06.NewAgentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnPremisesPublishingProfileRequestBuilder) AgentsById(id string)(*i50192dbecdd2d205018f1c9f0d43e5faa1231c597c1e3d9879fb3bfcceb8f7bf.OnPremisesAgentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onPremisesAgent_id"] = id
    }
    return i50192dbecdd2d205018f1c9f0d43e5faa1231c597c1e3d9879fb3bfcceb8f7bf.NewOnPremisesAgentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnPremisesPublishingProfileRequestBuilder) ConnectorGroups()(*i015798be98f5354e44d25cd3fc8c761a04c2de377021e2334f9f574c0b9d9d50.ConnectorGroupsRequestBuilder) {
    return i015798be98f5354e44d25cd3fc8c761a04c2de377021e2334f9f574c0b9d9d50.NewConnectorGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnPremisesPublishingProfileRequestBuilder) ConnectorGroupsById(id string)(*i4e613208e03a8604c73cff8fea0c650341897927012dab6031821bf1b1ae2305.ConnectorGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["connectorGroup_id"] = id
    }
    return i4e613208e03a8604c73cff8fea0c650341897927012dab6031821bf1b1ae2305.NewConnectorGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnPremisesPublishingProfileRequestBuilder) Connectors()(*i86c4967a04a253d001f5fd0f5c1dd90bb7ff7141cd43dab8638d3f8ddaff0d58.ConnectorsRequestBuilder) {
    return i86c4967a04a253d001f5fd0f5c1dd90bb7ff7141cd43dab8638d3f8ddaff0d58.NewConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnPremisesPublishingProfileRequestBuilder) ConnectorsById(id string)(*ia6c25f4c749e8beb990221d7c4363a8d4383190388b2db9197c94a520cdfffde.ConnectorRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["connector_id"] = id
    }
    return ia6c25f4c749e8beb990221d7c4363a8d4383190388b2db9197c94a520cdfffde.NewConnectorRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewOnPremisesPublishingProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnPremisesPublishingProfileRequestBuilder) {
    m := &OnPremisesPublishingProfileRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/onPremisesPublishingProfiles/{onPremisesPublishingProfile_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewOnPremisesPublishingProfileRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnPremisesPublishingProfileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnPremisesPublishingProfileRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *OnPremisesPublishingProfileRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *OnPremisesPublishingProfileRequestBuilder) CreateGetRequestInformation(q func (value *OnPremisesPublishingProfileRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(OnPremisesPublishingProfileRequestBuilderGetQueryParameters)
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
func (m *OnPremisesPublishingProfileRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnPremisesPublishingProfile, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *OnPremisesPublishingProfileRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *OnPremisesPublishingProfileRequestBuilder) Get(q func (value *OnPremisesPublishingProfileRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnPremisesPublishingProfile, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOnPremisesPublishingProfile() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnPremisesPublishingProfile), nil
}
func (m *OnPremisesPublishingProfileRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnPremisesPublishingProfile, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *OnPremisesPublishingProfileRequestBuilder) PublishedResources()(*if08fda1a25bbc0d180dcba2e00721696a53ac148a9d4209b397df9ba1382eb6f.PublishedResourcesRequestBuilder) {
    return if08fda1a25bbc0d180dcba2e00721696a53ac148a9d4209b397df9ba1382eb6f.NewPublishedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnPremisesPublishingProfileRequestBuilder) PublishedResourcesById(id string)(*i04768df19176575e14a7cad44b0f47526794459b68db53ab71668e3192948132.PublishedResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["publishedResource_id"] = id
    }
    return i04768df19176575e14a7cad44b0f47526794459b68db53ab71668e3192948132.NewPublishedResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
