package schedule

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i192f2bd8dcbb891ee1b03bc21aca5614e6bc710d6956e97702f0e8753f8340df "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/schedule/share"
    i1f0211dce1bdc2026dd77c9d4899d384b5952a056bdb914a8d847a78beb06b15 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/schedule/shifts"
    i2051f1be0f800c80a75d0e2247e4b8628195ce935f0b3c7e76ed0823336ee564 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/schedule/schedulinggroups"
    i4e8aa376ec3ef917782f2c351caed28ebf970c082d97233a9c96ec9466e08cff "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/schedule/timesoff"
    ia47dc33a63c4fcd6048c398eaa7e78c7cc2c18e9bb3db29edde4ec891a6b9cde "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/schedule/timeoffreasons"
    iaa5a235f9610e7046adc9f3f07b08ea98bc4ae6415beb10f12bb4387b253b0fd "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/schedule/timecards"
    iadaad397ac60b5737d07949890db17a8e72367980b7f2ccd4d09d93764bb84f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/schedule/swapshiftschangerequests"
    iae514e2420d44701d960ee0cff9eba20fc0214f2502ba28e2fd99c27758f2236 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/schedule/offershiftrequests"
    ibc1ac2610b5009a4e1c2d55fb50b8f0e69387fff13f818d8a5936d1d06677307 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/schedule/openshiftchangerequests"
    icfdafd08ca92b2e9fe870ec574f79a5b094b9e7885f10ef4603065b6f979596d "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/schedule/openshifts"
    id6c776397328fa3098991c0c215826da1b83aac65f56d6b2cd8f707074257d2c "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/schedule/timeoffrequests"
    i12d05e669b2e0306dc50086dbfff1ca0ae135aa78484385c383ebfb7d6f2cc39 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/schedule/shifts/item"
    i16f1bd133e03e4abc89c3268b9e9cf0e7e99010e9a82a52e4f40498cd7e8068c "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/schedule/schedulinggroups/item"
    i3ff9d61bef110902fcb8495506c96f2de5968639146e334d1fd0a28ba8170cd6 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/schedule/timeoffreasons/item"
    i6475a2e7fb7598fa782eb2d2d3ec1a48b17bb49c7c4c7180b58e1e44562bd5eb "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/schedule/timesoff/item"
    i794c6c2369b0653fa2bca3dae3dd8a256bcaa2a45d7ee95e181869cebd948aca "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/schedule/offershiftrequests/item"
    i8d5b7a84ec5bb85b5efb6f4a458428c46adc99e6d3ea9f2537fcc95263d594b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/schedule/timecards/item"
    ia2173e3d77cdb5087b7bbe5518699a07700a1478803049b800b985170179f5f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/schedule/openshiftchangerequests/item"
    iaad52a606569bcbb3fa9e8f51189f2bd38bfd65b07d1af085748fe435acc24de "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/schedule/timeoffrequests/item"
    idb4c9dfa757bfbe745484cbdeed9170bb3dd48a1d27b1a50c3839b2b637299c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/schedule/swapshiftschangerequests/item"
    ieb22aef658bc353a92503b3c132359b1ed96975b7d294ad7eb057d17b05f8c0c "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/teamtemplates/item/definitions/item/teamdefinition/schedule/openshifts/item"
)

// ScheduleRequestBuilder provides operations to manage the schedule property of the microsoft.graph.team entity.
type ScheduleRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ScheduleRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ScheduleRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ScheduleRequestBuilderGetQueryParameters retrieve the properties and relationships of a schedule object. The schedule creation process conforms to the One API guideline for resource based long running operations (RELO).When clients use the PUT method, if the schedule is provisioned, the operation updates the schedule; otherwise, the operation starts the schedule provisioning process in the background. During schedule provisioning, clients can use the GET method to get the schedule and look at the `provisionStatus` property for the current state of the provisioning. If the provisioning failed, clients can get additional information from the `provisionStatusCode` property. Clients can also inspect the configuration of the schedule.
type ScheduleRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ScheduleRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ScheduleRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ScheduleRequestBuilderGetQueryParameters
}
// ScheduleRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ScheduleRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewScheduleRequestBuilderInternal instantiates a new ScheduleRequestBuilder and sets the default values.
func NewScheduleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScheduleRequestBuilder) {
    m := &ScheduleRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/schedule{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewScheduleRequestBuilder instantiates a new ScheduleRequestBuilder and sets the default values.
func NewScheduleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScheduleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewScheduleRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property schedule for teamwork
func (m *ScheduleRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ScheduleRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation retrieve the properties and relationships of a schedule object. The schedule creation process conforms to the One API guideline for resource based long running operations (RELO).When clients use the PUT method, if the schedule is provisioned, the operation updates the schedule; otherwise, the operation starts the schedule provisioning process in the background. During schedule provisioning, clients can use the GET method to get the schedule and look at the `provisionStatus` property for the current state of the provisioning. If the provisioning failed, clients can get additional information from the `provisionStatusCode` property. Clients can also inspect the configuration of the schedule.
func (m *ScheduleRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ScheduleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePutRequestInformation update the navigation property schedule in teamwork
func (m *ScheduleRequestBuilder) CreatePutRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, requestConfiguration *ScheduleRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property schedule for teamwork
func (m *ScheduleRequestBuilder) Delete(ctx context.Context, requestConfiguration *ScheduleRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
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
// Get retrieve the properties and relationships of a schedule object. The schedule creation process conforms to the One API guideline for resource based long running operations (RELO).When clients use the PUT method, if the schedule is provisioned, the operation updates the schedule; otherwise, the operation starts the schedule provisioning process in the background. During schedule provisioning, clients can use the GET method to get the schedule and look at the `provisionStatus` property for the current state of the provisioning. If the provisioning failed, clients can get additional information from the `provisionStatusCode` property. Clients can also inspect the configuration of the schedule.
func (m *ScheduleRequestBuilder) Get(ctx context.Context, requestConfiguration *ScheduleRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable), nil
}
// OfferShiftRequests provides operations to manage the offerShiftRequests property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) OfferShiftRequests()(*iae514e2420d44701d960ee0cff9eba20fc0214f2502ba28e2fd99c27758f2236.OfferShiftRequestsRequestBuilder) {
    return iae514e2420d44701d960ee0cff9eba20fc0214f2502ba28e2fd99c27758f2236.NewOfferShiftRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OfferShiftRequestsById provides operations to manage the offerShiftRequests property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) OfferShiftRequestsById(id string)(*i794c6c2369b0653fa2bca3dae3dd8a256bcaa2a45d7ee95e181869cebd948aca.OfferShiftRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["offerShiftRequest%2Did"] = id
    }
    return i794c6c2369b0653fa2bca3dae3dd8a256bcaa2a45d7ee95e181869cebd948aca.NewOfferShiftRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OpenShiftChangeRequests provides operations to manage the openShiftChangeRequests property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) OpenShiftChangeRequests()(*ibc1ac2610b5009a4e1c2d55fb50b8f0e69387fff13f818d8a5936d1d06677307.OpenShiftChangeRequestsRequestBuilder) {
    return ibc1ac2610b5009a4e1c2d55fb50b8f0e69387fff13f818d8a5936d1d06677307.NewOpenShiftChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OpenShiftChangeRequestsById provides operations to manage the openShiftChangeRequests property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) OpenShiftChangeRequestsById(id string)(*ia2173e3d77cdb5087b7bbe5518699a07700a1478803049b800b985170179f5f9.OpenShiftChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShiftChangeRequest%2Did"] = id
    }
    return ia2173e3d77cdb5087b7bbe5518699a07700a1478803049b800b985170179f5f9.NewOpenShiftChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OpenShifts provides operations to manage the openShifts property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) OpenShifts()(*icfdafd08ca92b2e9fe870ec574f79a5b094b9e7885f10ef4603065b6f979596d.OpenShiftsRequestBuilder) {
    return icfdafd08ca92b2e9fe870ec574f79a5b094b9e7885f10ef4603065b6f979596d.NewOpenShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OpenShiftsById provides operations to manage the openShifts property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) OpenShiftsById(id string)(*ieb22aef658bc353a92503b3c132359b1ed96975b7d294ad7eb057d17b05f8c0c.OpenShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShift%2Did"] = id
    }
    return ieb22aef658bc353a92503b3c132359b1ed96975b7d294ad7eb057d17b05f8c0c.NewOpenShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Put update the navigation property schedule in teamwork
func (m *ScheduleRequestBuilder) Put(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, requestConfiguration *ScheduleRequestBuilderPutRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, error) {
    requestInfo, err := m.CreatePutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable), nil
}
// SchedulingGroups provides operations to manage the schedulingGroups property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) SchedulingGroups()(*i2051f1be0f800c80a75d0e2247e4b8628195ce935f0b3c7e76ed0823336ee564.SchedulingGroupsRequestBuilder) {
    return i2051f1be0f800c80a75d0e2247e4b8628195ce935f0b3c7e76ed0823336ee564.NewSchedulingGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SchedulingGroupsById provides operations to manage the schedulingGroups property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) SchedulingGroupsById(id string)(*i16f1bd133e03e4abc89c3268b9e9cf0e7e99010e9a82a52e4f40498cd7e8068c.SchedulingGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["schedulingGroup%2Did"] = id
    }
    return i16f1bd133e03e4abc89c3268b9e9cf0e7e99010e9a82a52e4f40498cd7e8068c.NewSchedulingGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Share provides operations to call the share method.
func (m *ScheduleRequestBuilder) Share()(*i192f2bd8dcbb891ee1b03bc21aca5614e6bc710d6956e97702f0e8753f8340df.ShareRequestBuilder) {
    return i192f2bd8dcbb891ee1b03bc21aca5614e6bc710d6956e97702f0e8753f8340df.NewShareRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Shifts provides operations to manage the shifts property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) Shifts()(*i1f0211dce1bdc2026dd77c9d4899d384b5952a056bdb914a8d847a78beb06b15.ShiftsRequestBuilder) {
    return i1f0211dce1bdc2026dd77c9d4899d384b5952a056bdb914a8d847a78beb06b15.NewShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ShiftsById provides operations to manage the shifts property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) ShiftsById(id string)(*i12d05e669b2e0306dc50086dbfff1ca0ae135aa78484385c383ebfb7d6f2cc39.ShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["shift%2Did"] = id
    }
    return i12d05e669b2e0306dc50086dbfff1ca0ae135aa78484385c383ebfb7d6f2cc39.NewShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SwapShiftsChangeRequests provides operations to manage the swapShiftsChangeRequests property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) SwapShiftsChangeRequests()(*iadaad397ac60b5737d07949890db17a8e72367980b7f2ccd4d09d93764bb84f6.SwapShiftsChangeRequestsRequestBuilder) {
    return iadaad397ac60b5737d07949890db17a8e72367980b7f2ccd4d09d93764bb84f6.NewSwapShiftsChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SwapShiftsChangeRequestsById provides operations to manage the swapShiftsChangeRequests property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) SwapShiftsChangeRequestsById(id string)(*idb4c9dfa757bfbe745484cbdeed9170bb3dd48a1d27b1a50c3839b2b637299c6.SwapShiftsChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["swapShiftsChangeRequest%2Did"] = id
    }
    return idb4c9dfa757bfbe745484cbdeed9170bb3dd48a1d27b1a50c3839b2b637299c6.NewSwapShiftsChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimeCards provides operations to manage the timeCards property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) TimeCards()(*iaa5a235f9610e7046adc9f3f07b08ea98bc4ae6415beb10f12bb4387b253b0fd.TimeCardsRequestBuilder) {
    return iaa5a235f9610e7046adc9f3f07b08ea98bc4ae6415beb10f12bb4387b253b0fd.NewTimeCardsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeCardsById provides operations to manage the timeCards property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) TimeCardsById(id string)(*i8d5b7a84ec5bb85b5efb6f4a458428c46adc99e6d3ea9f2537fcc95263d594b8.TimeCardItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeCard%2Did"] = id
    }
    return i8d5b7a84ec5bb85b5efb6f4a458428c46adc99e6d3ea9f2537fcc95263d594b8.NewTimeCardItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimeOffReasons provides operations to manage the timeOffReasons property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) TimeOffReasons()(*ia47dc33a63c4fcd6048c398eaa7e78c7cc2c18e9bb3db29edde4ec891a6b9cde.TimeOffReasonsRequestBuilder) {
    return ia47dc33a63c4fcd6048c398eaa7e78c7cc2c18e9bb3db29edde4ec891a6b9cde.NewTimeOffReasonsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeOffReasonsById provides operations to manage the timeOffReasons property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) TimeOffReasonsById(id string)(*i3ff9d61bef110902fcb8495506c96f2de5968639146e334d1fd0a28ba8170cd6.TimeOffReasonItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffReason%2Did"] = id
    }
    return i3ff9d61bef110902fcb8495506c96f2de5968639146e334d1fd0a28ba8170cd6.NewTimeOffReasonItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimeOffRequests provides operations to manage the timeOffRequests property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) TimeOffRequests()(*id6c776397328fa3098991c0c215826da1b83aac65f56d6b2cd8f707074257d2c.TimeOffRequestsRequestBuilder) {
    return id6c776397328fa3098991c0c215826da1b83aac65f56d6b2cd8f707074257d2c.NewTimeOffRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeOffRequestsById provides operations to manage the timeOffRequests property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) TimeOffRequestsById(id string)(*iaad52a606569bcbb3fa9e8f51189f2bd38bfd65b07d1af085748fe435acc24de.TimeOffRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffRequest%2Did"] = id
    }
    return iaad52a606569bcbb3fa9e8f51189f2bd38bfd65b07d1af085748fe435acc24de.NewTimeOffRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimesOff provides operations to manage the timesOff property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) TimesOff()(*i4e8aa376ec3ef917782f2c351caed28ebf970c082d97233a9c96ec9466e08cff.TimesOffRequestBuilder) {
    return i4e8aa376ec3ef917782f2c351caed28ebf970c082d97233a9c96ec9466e08cff.NewTimesOffRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimesOffById provides operations to manage the timesOff property of the microsoft.graph.schedule entity.
func (m *ScheduleRequestBuilder) TimesOffById(id string)(*i6475a2e7fb7598fa782eb2d2d3ec1a48b17bb49c7c4c7180b58e1e44562bd5eb.TimeOffItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOff%2Did"] = id
    }
    return i6475a2e7fb7598fa782eb2d2d3ec1a48b17bb49c7c4c7180b58e1e44562bd5eb.NewTimeOffItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
