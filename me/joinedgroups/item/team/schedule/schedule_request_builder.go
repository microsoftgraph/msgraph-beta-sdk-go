package schedule

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0a6c72d5372d465ade643d6521bd39ecc21a356e5cb796e2398f30ea2bc16f57 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/schedule/openshifts"
    i0ada0e90819f550b70ac813afc10c067f9f71c1facc9786af76fed1132e42f56 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/schedule/timesoff"
    i50ec41e89cf99e92988e7d0b1158b124c654ded088ec37579e0219358f115971 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/schedule/swapshiftschangerequests"
    i7c89bec16884299ed47c3e006a4ff852556ebcdf887d9f8e83ba1a4f853331d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/schedule/share"
    i9405c0647929fbbc4aba6ffe9bae1481af0103a7dcd8f3d837a7eb2db0d85422 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/schedule/offershiftrequests"
    ib1a2cd9d10c137fb9d8aa4e2a5177bac13599aa47502ef63f19b2dfa7805480f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/schedule/schedulinggroups"
    ic5ea6a12c1cf6bb39a0f03e9d17ddf01aa3cf21ad3668898ea57534a9ff5dd60 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/schedule/timeoffrequests"
    icb2e925840a1e88a80c8e1474c675bf0c3002e8437c362774c34c8501355a010 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/schedule/timeoffreasons"
    id0e7a4128acd7fc6931368a3017e4a2e69f2a8e9869e8ae73da274b1db1c5ced "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/schedule/shifts"
    id0ee69b1a3a26284950674930d979f7b2fc7de9edf13db20300153e2f7d76a09 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/schedule/openshiftchangerequests"
    ifda6671ab5ab48587c794683f8e1bdd50fbf18fbfb4dc6400080cb356e992e3a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/schedule/timecards"
    i0273cf487f10fc133f8e26938284ed83283c834d1cfc06782b7f71dd985f5ded "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/schedule/swapshiftschangerequests/item"
    i1b63ae81337e3432033759d6e4146cced534c4497419b7960fca6209850c2207 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/schedule/timecards/item"
    i2b7b98b015249ce9d461bb5a087f75f2fc7b21b628de26691cc27f595ffdd3da "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/schedule/timesoff/item"
    i2de009fc6bd1745764885b57676e12b7db4f1985c5fe8da942d7024c2f627bab "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/schedule/schedulinggroups/item"
    i3292d8544ab0838f203e42bc6e2b4b6cef22101bc2f099d783643c67673687a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/schedule/openshiftchangerequests/item"
    i3cc911c54d2813f10d461b0ebec6d422aa0dfe3783b4f934f3fc089b34689cf4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/schedule/offershiftrequests/item"
    i4a01c170e0944fcaf515f4cdf1f42337abf44857f9050bd270523b8ab2375c34 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/schedule/timeoffreasons/item"
    i6b7e255d7c80adbd70c8fb9450d7a9b5ffacb43eefb887be8c1f23079d0a1120 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/schedule/shifts/item"
    id84ebfa63f11e4287cda2b005b404bae94571b98d1b07b74867515036e47317b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/schedule/openshifts/item"
    if0c32c86a629925cc794ec749a26a6acf40053e71e069418e046ee1d084134a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/team/schedule/timeoffrequests/item"
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
    m.urlTemplate = "{+baseurl}/me/joinedGroups/{group%2Did}/team/schedule{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property schedule for me
func (m *ScheduleRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property schedule for me
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
// CreatePatchRequestInformation update the navigation property schedule in me
func (m *ScheduleRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property schedule in me
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
// Delete delete navigation property schedule for me
func (m *ScheduleRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property schedule for me
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
func (m *ScheduleRequestBuilder) OfferShiftRequests()(*i9405c0647929fbbc4aba6ffe9bae1481af0103a7dcd8f3d837a7eb2db0d85422.OfferShiftRequestsRequestBuilder) {
    return i9405c0647929fbbc4aba6ffe9bae1481af0103a7dcd8f3d837a7eb2db0d85422.NewOfferShiftRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OfferShiftRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.team.schedule.offerShiftRequests.item collection
func (m *ScheduleRequestBuilder) OfferShiftRequestsById(id string)(*i3cc911c54d2813f10d461b0ebec6d422aa0dfe3783b4f934f3fc089b34689cf4.OfferShiftRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["offerShiftRequest%2Did"] = id
    }
    return i3cc911c54d2813f10d461b0ebec6d422aa0dfe3783b4f934f3fc089b34689cf4.NewOfferShiftRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OpenShiftChangeRequests the openShiftChangeRequests property
func (m *ScheduleRequestBuilder) OpenShiftChangeRequests()(*id0ee69b1a3a26284950674930d979f7b2fc7de9edf13db20300153e2f7d76a09.OpenShiftChangeRequestsRequestBuilder) {
    return id0ee69b1a3a26284950674930d979f7b2fc7de9edf13db20300153e2f7d76a09.NewOpenShiftChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OpenShiftChangeRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.team.schedule.openShiftChangeRequests.item collection
func (m *ScheduleRequestBuilder) OpenShiftChangeRequestsById(id string)(*i3292d8544ab0838f203e42bc6e2b4b6cef22101bc2f099d783643c67673687a9.OpenShiftChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShiftChangeRequest%2Did"] = id
    }
    return i3292d8544ab0838f203e42bc6e2b4b6cef22101bc2f099d783643c67673687a9.NewOpenShiftChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OpenShifts the openShifts property
func (m *ScheduleRequestBuilder) OpenShifts()(*i0a6c72d5372d465ade643d6521bd39ecc21a356e5cb796e2398f30ea2bc16f57.OpenShiftsRequestBuilder) {
    return i0a6c72d5372d465ade643d6521bd39ecc21a356e5cb796e2398f30ea2bc16f57.NewOpenShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OpenShiftsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.team.schedule.openShifts.item collection
func (m *ScheduleRequestBuilder) OpenShiftsById(id string)(*id84ebfa63f11e4287cda2b005b404bae94571b98d1b07b74867515036e47317b.OpenShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShift%2Did"] = id
    }
    return id84ebfa63f11e4287cda2b005b404bae94571b98d1b07b74867515036e47317b.NewOpenShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property schedule in me
func (m *ScheduleRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Scheduleable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property schedule in me
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
func (m *ScheduleRequestBuilder) SchedulingGroups()(*ib1a2cd9d10c137fb9d8aa4e2a5177bac13599aa47502ef63f19b2dfa7805480f.SchedulingGroupsRequestBuilder) {
    return ib1a2cd9d10c137fb9d8aa4e2a5177bac13599aa47502ef63f19b2dfa7805480f.NewSchedulingGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SchedulingGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.team.schedule.schedulingGroups.item collection
func (m *ScheduleRequestBuilder) SchedulingGroupsById(id string)(*i2de009fc6bd1745764885b57676e12b7db4f1985c5fe8da942d7024c2f627bab.SchedulingGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["schedulingGroup%2Did"] = id
    }
    return i2de009fc6bd1745764885b57676e12b7db4f1985c5fe8da942d7024c2f627bab.NewSchedulingGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Share the share property
func (m *ScheduleRequestBuilder) Share()(*i7c89bec16884299ed47c3e006a4ff852556ebcdf887d9f8e83ba1a4f853331d5.ShareRequestBuilder) {
    return i7c89bec16884299ed47c3e006a4ff852556ebcdf887d9f8e83ba1a4f853331d5.NewShareRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Shifts the shifts property
func (m *ScheduleRequestBuilder) Shifts()(*id0e7a4128acd7fc6931368a3017e4a2e69f2a8e9869e8ae73da274b1db1c5ced.ShiftsRequestBuilder) {
    return id0e7a4128acd7fc6931368a3017e4a2e69f2a8e9869e8ae73da274b1db1c5ced.NewShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ShiftsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.team.schedule.shifts.item collection
func (m *ScheduleRequestBuilder) ShiftsById(id string)(*i6b7e255d7c80adbd70c8fb9450d7a9b5ffacb43eefb887be8c1f23079d0a1120.ShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["shift%2Did"] = id
    }
    return i6b7e255d7c80adbd70c8fb9450d7a9b5ffacb43eefb887be8c1f23079d0a1120.NewShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SwapShiftsChangeRequests the swapShiftsChangeRequests property
func (m *ScheduleRequestBuilder) SwapShiftsChangeRequests()(*i50ec41e89cf99e92988e7d0b1158b124c654ded088ec37579e0219358f115971.SwapShiftsChangeRequestsRequestBuilder) {
    return i50ec41e89cf99e92988e7d0b1158b124c654ded088ec37579e0219358f115971.NewSwapShiftsChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SwapShiftsChangeRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.team.schedule.swapShiftsChangeRequests.item collection
func (m *ScheduleRequestBuilder) SwapShiftsChangeRequestsById(id string)(*i0273cf487f10fc133f8e26938284ed83283c834d1cfc06782b7f71dd985f5ded.SwapShiftsChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["swapShiftsChangeRequest%2Did"] = id
    }
    return i0273cf487f10fc133f8e26938284ed83283c834d1cfc06782b7f71dd985f5ded.NewSwapShiftsChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimeCards the timeCards property
func (m *ScheduleRequestBuilder) TimeCards()(*ifda6671ab5ab48587c794683f8e1bdd50fbf18fbfb4dc6400080cb356e992e3a.TimeCardsRequestBuilder) {
    return ifda6671ab5ab48587c794683f8e1bdd50fbf18fbfb4dc6400080cb356e992e3a.NewTimeCardsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeCardsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.team.schedule.timeCards.item collection
func (m *ScheduleRequestBuilder) TimeCardsById(id string)(*i1b63ae81337e3432033759d6e4146cced534c4497419b7960fca6209850c2207.TimeCardItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeCard%2Did"] = id
    }
    return i1b63ae81337e3432033759d6e4146cced534c4497419b7960fca6209850c2207.NewTimeCardItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimeOffReasons the timeOffReasons property
func (m *ScheduleRequestBuilder) TimeOffReasons()(*icb2e925840a1e88a80c8e1474c675bf0c3002e8437c362774c34c8501355a010.TimeOffReasonsRequestBuilder) {
    return icb2e925840a1e88a80c8e1474c675bf0c3002e8437c362774c34c8501355a010.NewTimeOffReasonsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeOffReasonsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.team.schedule.timeOffReasons.item collection
func (m *ScheduleRequestBuilder) TimeOffReasonsById(id string)(*i4a01c170e0944fcaf515f4cdf1f42337abf44857f9050bd270523b8ab2375c34.TimeOffReasonItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffReason%2Did"] = id
    }
    return i4a01c170e0944fcaf515f4cdf1f42337abf44857f9050bd270523b8ab2375c34.NewTimeOffReasonItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimeOffRequests the timeOffRequests property
func (m *ScheduleRequestBuilder) TimeOffRequests()(*ic5ea6a12c1cf6bb39a0f03e9d17ddf01aa3cf21ad3668898ea57534a9ff5dd60.TimeOffRequestsRequestBuilder) {
    return ic5ea6a12c1cf6bb39a0f03e9d17ddf01aa3cf21ad3668898ea57534a9ff5dd60.NewTimeOffRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeOffRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.team.schedule.timeOffRequests.item collection
func (m *ScheduleRequestBuilder) TimeOffRequestsById(id string)(*if0c32c86a629925cc794ec749a26a6acf40053e71e069418e046ee1d084134a8.TimeOffRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffRequest%2Did"] = id
    }
    return if0c32c86a629925cc794ec749a26a6acf40053e71e069418e046ee1d084134a8.NewTimeOffRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimesOff the timesOff property
func (m *ScheduleRequestBuilder) TimesOff()(*i0ada0e90819f550b70ac813afc10c067f9f71c1facc9786af76fed1132e42f56.TimesOffRequestBuilder) {
    return i0ada0e90819f550b70ac813afc10c067f9f71c1facc9786af76fed1132e42f56.NewTimesOffRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimesOffById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.team.schedule.timesOff.item collection
func (m *ScheduleRequestBuilder) TimesOffById(id string)(*i2b7b98b015249ce9d461bb5a087f75f2fc7b21b628de26691cc27f595ffdd3da.TimeOffItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOff%2Did"] = id
    }
    return i2b7b98b015249ce9d461bb5a087f75f2fc7b21b628de26691cc27f595ffdd3da.NewTimeOffItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
