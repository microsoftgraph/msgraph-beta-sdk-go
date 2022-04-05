package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i5127ef050fde02b8b8f1a5f5fe72b8694bf708f24b9f1e0b8352fb67d87e9683 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule/timecards/item/endbreak"
    i643863d7379eb2ab6adf49e281ac4151461d6911ca95ffa366ad2d8c5e9f4ce2 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule/timecards/item/startbreak"
    i68e4a9fcef652132c7bebba9ebe08cb4fd8be16b5de9cd2ac842a8f4d53cc5ce "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule/timecards/item/confirm"
    i928cffe1d44a7c917c1e83d9a4105e827fa065d83a7a505a4f7efa3c64266fdb "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule/timecards/item/clockout"
)

// TimeCardItemRequestBuilder provides operations to manage the timeCards property of the microsoft.graph.schedule entity.
type TimeCardItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TimeCardItemRequestBuilderDeleteOptions options for Delete
type TimeCardItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// TimeCardItemRequestBuilderGetOptions options for Get
type TimeCardItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *TimeCardItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// TimeCardItemRequestBuilderGetQueryParameters get timeCards from teams
type TimeCardItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// TimeCardItemRequestBuilderPatchOptions options for Patch
type TimeCardItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeCardable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// ClockOut the clockOut property
func (m *TimeCardItemRequestBuilder) ClockOut()(*i928cffe1d44a7c917c1e83d9a4105e827fa065d83a7a505a4f7efa3c64266fdb.ClockOutRequestBuilder) {
    return i928cffe1d44a7c917c1e83d9a4105e827fa065d83a7a505a4f7efa3c64266fdb.NewClockOutRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Confirm the confirm property
func (m *TimeCardItemRequestBuilder) Confirm()(*i68e4a9fcef652132c7bebba9ebe08cb4fd8be16b5de9cd2ac842a8f4d53cc5ce.ConfirmRequestBuilder) {
    return i68e4a9fcef652132c7bebba9ebe08cb4fd8be16b5de9cd2ac842a8f4d53cc5ce.NewConfirmRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTimeCardItemRequestBuilderInternal instantiates a new TimeCardItemRequestBuilder and sets the default values.
func NewTimeCardItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TimeCardItemRequestBuilder) {
    m := &TimeCardItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teams/{team_id}/schedule/timeCards/{timeCard_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTimeCardItemRequestBuilder instantiates a new TimeCardItemRequestBuilder and sets the default values.
func NewTimeCardItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TimeCardItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTimeCardItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property timeCards for teams
func (m *TimeCardItemRequestBuilder) CreateDeleteRequestInformation(options *TimeCardItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get timeCards from teams
func (m *TimeCardItemRequestBuilder) CreateGetRequestInformation(options *TimeCardItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property timeCards in teams
func (m *TimeCardItemRequestBuilder) CreatePatchRequestInformation(options *TimeCardItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property timeCards for teams
func (m *TimeCardItemRequestBuilder) Delete(options *TimeCardItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// EndBreak the endBreak property
func (m *TimeCardItemRequestBuilder) EndBreak()(*i5127ef050fde02b8b8f1a5f5fe72b8694bf708f24b9f1e0b8352fb67d87e9683.EndBreakRequestBuilder) {
    return i5127ef050fde02b8b8f1a5f5fe72b8694bf708f24b9f1e0b8352fb67d87e9683.NewEndBreakRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get timeCards from teams
func (m *TimeCardItemRequestBuilder) Get(options *TimeCardItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeCardable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTimeCardFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeCardable), nil
}
// Patch update the navigation property timeCards in teams
func (m *TimeCardItemRequestBuilder) Patch(options *TimeCardItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// StartBreak the startBreak property
func (m *TimeCardItemRequestBuilder) StartBreak()(*i643863d7379eb2ab6adf49e281ac4151461d6911ca95ffa366ad2d8c5e9f4ce2.StartBreakRequestBuilder) {
    return i643863d7379eb2ab6adf49e281ac4151461d6911ca95ffa366ad2d8c5e9f4ce2.NewStartBreakRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
