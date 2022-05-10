package schedule

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0b785e0cce8b4b59962a4c4b294420e8577400865a8a1e5ca10918c1d2b3808d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/schedule/share"
    i2a7b55a5587d79d60e8aeeee131aea97ac618ad6f9ce81d5435713649d26fae0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/schedule/schedulinggroups"
    i2e0dc6a915f3d75516b79aee2705adf45b64ffe64646a7246e11f37d1e6e25de "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/schedule/offershiftrequests"
    i2e31dd6fd4933b648e386e8652eccf7859b14fba40269a3c1316f3aaf79e49a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/schedule/timecards"
    i67d90b78bfcc6c3eaba45a8e83f3e3f791fd5c2e34f5c5caf8eb2c34a0f7969d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/schedule/swapshiftschangerequests"
    i7aa035ef1c10e42f7aa18afb0525da0752f5258fee84456a4904db5512266a41 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/schedule/timeoffreasons"
    i89ab6e2437cf21737bb38b737c1b3245cb6c59dc8d63b03f38d357118efc0175 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/schedule/shifts"
    i89fadbd74817e84dcea2c7d828a04d9ef7964daf22dc08cd01dc2aee37278ee7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/schedule/timeoffrequests"
    i9865b25035659693d4011b081e72ee37c30c9e3b78147587e517a6a8ab6f87f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/schedule/timesoff"
    ia4b66b1183c3b19a7a78162eea9860134ae534bd4dc7fe4af8b3725c261be4fc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/schedule/openshifts"
    ia8bb663241b110d3ca32f325c7e918871ec8b1b2d1b8917b06582419325e4aa0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/schedule/openshiftchangerequests"
    i0375816b9f63fe6e18b95e445cd08656118e4197a7c0867b681ce6e17a222c4e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/schedule/shifts/item"
    i2349c093cb925287640b9b71bf8f70fb89a8038d50883d37846b641582364532 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/schedule/openshiftchangerequests/item"
    i4197a3cb378603a7e7d28b7e679305d7fd5f576ba0d6e0fcbad886fe6f6cb49a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/schedule/schedulinggroups/item"
    i53ac935f2b877f6c397e4c20092aa44ec0e62279d97b55e097458b7ea2e83f86 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/schedule/timecards/item"
    i544b30f76955dc7269117c4c8ae30906407333e138d6ea7decd696c4a09be94a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/schedule/offershiftrequests/item"
    i703a7a2d5596180ec434321178937406cb551e2633b873c4abdfdc6cbc118da9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/schedule/openshifts/item"
    iafb736e280074aecde4e3d55603cf3b2289e6542b45d0ac751eb3a8a8e27c51a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/schedule/timesoff/item"
    iba95a8e7cb83aa3d9ad22f8412dd662207d05a2d1ef9b4ab75631d4e2130c93f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/schedule/timeoffreasons/item"
    id16ec07bfcda24a8b2218453f49551dfdba879b934b27e9f6ad5d001c5160b56 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/schedule/timeoffrequests/item"
    ie28c30e60594c0acf8a7363775f8a042204fd6ad5b16b3db3d45f3aeee44cf92 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/schedule/swapshiftschangerequests/item"
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
// ScheduleRequestBuilderGetQueryParameters the schedule of shifts for this team.
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
// ScheduleRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ScheduleRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewScheduleRequestBuilderInternal instantiates a new ScheduleRequestBuilder and sets the default values.
func NewScheduleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScheduleRequestBuilder) {
    m := &ScheduleRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedGroups/{group%2Did}/team/schedule{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property schedule for users
func (m *ScheduleRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property schedule for users
func (m *ScheduleRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *ScheduleRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the schedule of shifts for this team.
func (m *ScheduleRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the schedule of shifts for this team.
func (m *ScheduleRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ScheduleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property schedule in users
func (m *ScheduleRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property schedule in users
func (m *ScheduleRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, requestConfiguration *ScheduleRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property schedule for users
func (m *ScheduleRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property schedule for users
func (m *ScheduleRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *ScheduleRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
// Get the schedule of shifts for this team.
func (m *ScheduleRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the schedule of shifts for this team.
func (m *ScheduleRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *ScheduleRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateScheduleFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable), nil
}
// OfferShiftRequests the offerShiftRequests property
func (m *ScheduleRequestBuilder) OfferShiftRequests()(*i2e0dc6a915f3d75516b79aee2705adf45b64ffe64646a7246e11f37d1e6e25de.OfferShiftRequestsRequestBuilder) {
    return i2e0dc6a915f3d75516b79aee2705adf45b64ffe64646a7246e11f37d1e6e25de.NewOfferShiftRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OfferShiftRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.team.schedule.offerShiftRequests.item collection
func (m *ScheduleRequestBuilder) OfferShiftRequestsById(id string)(*i544b30f76955dc7269117c4c8ae30906407333e138d6ea7decd696c4a09be94a.OfferShiftRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["offerShiftRequest%2Did"] = id
    }
    return i544b30f76955dc7269117c4c8ae30906407333e138d6ea7decd696c4a09be94a.NewOfferShiftRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OpenShiftChangeRequests the openShiftChangeRequests property
func (m *ScheduleRequestBuilder) OpenShiftChangeRequests()(*ia8bb663241b110d3ca32f325c7e918871ec8b1b2d1b8917b06582419325e4aa0.OpenShiftChangeRequestsRequestBuilder) {
    return ia8bb663241b110d3ca32f325c7e918871ec8b1b2d1b8917b06582419325e4aa0.NewOpenShiftChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OpenShiftChangeRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.team.schedule.openShiftChangeRequests.item collection
func (m *ScheduleRequestBuilder) OpenShiftChangeRequestsById(id string)(*i2349c093cb925287640b9b71bf8f70fb89a8038d50883d37846b641582364532.OpenShiftChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShiftChangeRequest%2Did"] = id
    }
    return i2349c093cb925287640b9b71bf8f70fb89a8038d50883d37846b641582364532.NewOpenShiftChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OpenShifts the openShifts property
func (m *ScheduleRequestBuilder) OpenShifts()(*ia4b66b1183c3b19a7a78162eea9860134ae534bd4dc7fe4af8b3725c261be4fc.OpenShiftsRequestBuilder) {
    return ia4b66b1183c3b19a7a78162eea9860134ae534bd4dc7fe4af8b3725c261be4fc.NewOpenShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OpenShiftsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.team.schedule.openShifts.item collection
func (m *ScheduleRequestBuilder) OpenShiftsById(id string)(*i703a7a2d5596180ec434321178937406cb551e2633b873c4abdfdc6cbc118da9.OpenShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShift%2Did"] = id
    }
    return i703a7a2d5596180ec434321178937406cb551e2633b873c4abdfdc6cbc118da9.NewOpenShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property schedule in users
func (m *ScheduleRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property schedule in users
func (m *ScheduleRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable, requestConfiguration *ScheduleRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// SchedulingGroups the schedulingGroups property
func (m *ScheduleRequestBuilder) SchedulingGroups()(*i2a7b55a5587d79d60e8aeeee131aea97ac618ad6f9ce81d5435713649d26fae0.SchedulingGroupsRequestBuilder) {
    return i2a7b55a5587d79d60e8aeeee131aea97ac618ad6f9ce81d5435713649d26fae0.NewSchedulingGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SchedulingGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.team.schedule.schedulingGroups.item collection
func (m *ScheduleRequestBuilder) SchedulingGroupsById(id string)(*i4197a3cb378603a7e7d28b7e679305d7fd5f576ba0d6e0fcbad886fe6f6cb49a.SchedulingGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["schedulingGroup%2Did"] = id
    }
    return i4197a3cb378603a7e7d28b7e679305d7fd5f576ba0d6e0fcbad886fe6f6cb49a.NewSchedulingGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Share the share property
func (m *ScheduleRequestBuilder) Share()(*i0b785e0cce8b4b59962a4c4b294420e8577400865a8a1e5ca10918c1d2b3808d.ShareRequestBuilder) {
    return i0b785e0cce8b4b59962a4c4b294420e8577400865a8a1e5ca10918c1d2b3808d.NewShareRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Shifts the shifts property
func (m *ScheduleRequestBuilder) Shifts()(*i89ab6e2437cf21737bb38b737c1b3245cb6c59dc8d63b03f38d357118efc0175.ShiftsRequestBuilder) {
    return i89ab6e2437cf21737bb38b737c1b3245cb6c59dc8d63b03f38d357118efc0175.NewShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ShiftsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.team.schedule.shifts.item collection
func (m *ScheduleRequestBuilder) ShiftsById(id string)(*i0375816b9f63fe6e18b95e445cd08656118e4197a7c0867b681ce6e17a222c4e.ShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["shift%2Did"] = id
    }
    return i0375816b9f63fe6e18b95e445cd08656118e4197a7c0867b681ce6e17a222c4e.NewShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SwapShiftsChangeRequests the swapShiftsChangeRequests property
func (m *ScheduleRequestBuilder) SwapShiftsChangeRequests()(*i67d90b78bfcc6c3eaba45a8e83f3e3f791fd5c2e34f5c5caf8eb2c34a0f7969d.SwapShiftsChangeRequestsRequestBuilder) {
    return i67d90b78bfcc6c3eaba45a8e83f3e3f791fd5c2e34f5c5caf8eb2c34a0f7969d.NewSwapShiftsChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SwapShiftsChangeRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.team.schedule.swapShiftsChangeRequests.item collection
func (m *ScheduleRequestBuilder) SwapShiftsChangeRequestsById(id string)(*ie28c30e60594c0acf8a7363775f8a042204fd6ad5b16b3db3d45f3aeee44cf92.SwapShiftsChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["swapShiftsChangeRequest%2Did"] = id
    }
    return ie28c30e60594c0acf8a7363775f8a042204fd6ad5b16b3db3d45f3aeee44cf92.NewSwapShiftsChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimeCards the timeCards property
func (m *ScheduleRequestBuilder) TimeCards()(*i2e31dd6fd4933b648e386e8652eccf7859b14fba40269a3c1316f3aaf79e49a8.TimeCardsRequestBuilder) {
    return i2e31dd6fd4933b648e386e8652eccf7859b14fba40269a3c1316f3aaf79e49a8.NewTimeCardsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeCardsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.team.schedule.timeCards.item collection
func (m *ScheduleRequestBuilder) TimeCardsById(id string)(*i53ac935f2b877f6c397e4c20092aa44ec0e62279d97b55e097458b7ea2e83f86.TimeCardItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeCard%2Did"] = id
    }
    return i53ac935f2b877f6c397e4c20092aa44ec0e62279d97b55e097458b7ea2e83f86.NewTimeCardItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimeOffReasons the timeOffReasons property
func (m *ScheduleRequestBuilder) TimeOffReasons()(*i7aa035ef1c10e42f7aa18afb0525da0752f5258fee84456a4904db5512266a41.TimeOffReasonsRequestBuilder) {
    return i7aa035ef1c10e42f7aa18afb0525da0752f5258fee84456a4904db5512266a41.NewTimeOffReasonsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeOffReasonsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.team.schedule.timeOffReasons.item collection
func (m *ScheduleRequestBuilder) TimeOffReasonsById(id string)(*iba95a8e7cb83aa3d9ad22f8412dd662207d05a2d1ef9b4ab75631d4e2130c93f.TimeOffReasonItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffReason%2Did"] = id
    }
    return iba95a8e7cb83aa3d9ad22f8412dd662207d05a2d1ef9b4ab75631d4e2130c93f.NewTimeOffReasonItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimeOffRequests the timeOffRequests property
func (m *ScheduleRequestBuilder) TimeOffRequests()(*i89fadbd74817e84dcea2c7d828a04d9ef7964daf22dc08cd01dc2aee37278ee7.TimeOffRequestsRequestBuilder) {
    return i89fadbd74817e84dcea2c7d828a04d9ef7964daf22dc08cd01dc2aee37278ee7.NewTimeOffRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeOffRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.team.schedule.timeOffRequests.item collection
func (m *ScheduleRequestBuilder) TimeOffRequestsById(id string)(*id16ec07bfcda24a8b2218453f49551dfdba879b934b27e9f6ad5d001c5160b56.TimeOffRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffRequest%2Did"] = id
    }
    return id16ec07bfcda24a8b2218453f49551dfdba879b934b27e9f6ad5d001c5160b56.NewTimeOffRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimesOff the timesOff property
func (m *ScheduleRequestBuilder) TimesOff()(*i9865b25035659693d4011b081e72ee37c30c9e3b78147587e517a6a8ab6f87f5.TimesOffRequestBuilder) {
    return i9865b25035659693d4011b081e72ee37c30c9e3b78147587e517a6a8ab6f87f5.NewTimesOffRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimesOffById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.team.schedule.timesOff.item collection
func (m *ScheduleRequestBuilder) TimesOffById(id string)(*iafb736e280074aecde4e3d55603cf3b2289e6542b45d0ac751eb3a8a8e27c51a.TimeOffItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOff%2Did"] = id
    }
    return iafb736e280074aecde4e3d55603cf3b2289e6542b45d0ac751eb3a8a8e27c51a.NewTimeOffItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
