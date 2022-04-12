package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2243091c8fa68ba9e7a97bce329cf8353df512034105a6cdb60022d8a3462d3f "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/rosters/item/plans"
    i9ad7de99f1b757c4ff16ccef4d347a843f114b88ad7c9579c76ffa1dd8f293f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/rosters/item/members"
    ie53712fa7eb3bcb49fdf48153b4145c0918f55b3d8a9851bfbf8ce2045244c3c "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/rosters/item/members/item"
    if82e900011d72e8f7a8566012bb928604f336d22e9e7bd8049c43486ff3d7ff3 "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/rosters/item/plans/item"
)

// PlannerRosterItemRequestBuilder provides operations to manage the rosters property of the microsoft.graph.planner entity.
type PlannerRosterItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PlannerRosterItemRequestBuilderDeleteOptions options for Delete
type PlannerRosterItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// PlannerRosterItemRequestBuilderGetOptions options for Get
type PlannerRosterItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PlannerRosterItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// PlannerRosterItemRequestBuilderGetQueryParameters read-only. Nullable. Returns a collection of the specified rosters
type PlannerRosterItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PlannerRosterItemRequestBuilderPatchOptions options for Patch
type PlannerRosterItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerRosterable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewPlannerRosterItemRequestBuilderInternal instantiates a new PlannerRosterItemRequestBuilder and sets the default values.
func NewPlannerRosterItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PlannerRosterItemRequestBuilder) {
    m := &PlannerRosterItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/planner/rosters/{plannerRoster%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPlannerRosterItemRequestBuilder instantiates a new PlannerRosterItemRequestBuilder and sets the default values.
func NewPlannerRosterItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PlannerRosterItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPlannerRosterItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property rosters for planner
func (m *PlannerRosterItemRequestBuilder) CreateDeleteRequestInformation(options *PlannerRosterItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation read-only. Nullable. Returns a collection of the specified rosters
func (m *PlannerRosterItemRequestBuilder) CreateGetRequestInformation(options *PlannerRosterItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property rosters in planner
func (m *PlannerRosterItemRequestBuilder) CreatePatchRequestInformation(options *PlannerRosterItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property rosters for planner
func (m *PlannerRosterItemRequestBuilder) Delete(options *PlannerRosterItemRequestBuilderDeleteOptions)(error) {
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
// Get read-only. Nullable. Returns a collection of the specified rosters
func (m *PlannerRosterItemRequestBuilder) Get(options *PlannerRosterItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerRosterable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerRosterFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerRosterable), nil
}
// Members the members property
func (m *PlannerRosterItemRequestBuilder) Members()(*i9ad7de99f1b757c4ff16ccef4d347a843f114b88ad7c9579c76ffa1dd8f293f9.MembersRequestBuilder) {
    return i9ad7de99f1b757c4ff16ccef4d347a843f114b88ad7c9579c76ffa1dd8f293f9.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.planner.rosters.item.members.item collection
func (m *PlannerRosterItemRequestBuilder) MembersById(id string)(*ie53712fa7eb3bcb49fdf48153b4145c0918f55b3d8a9851bfbf8ce2045244c3c.PlannerRosterMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerRosterMember%2Did"] = id
    }
    return ie53712fa7eb3bcb49fdf48153b4145c0918f55b3d8a9851bfbf8ce2045244c3c.NewPlannerRosterMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property rosters in planner
func (m *PlannerRosterItemRequestBuilder) Patch(options *PlannerRosterItemRequestBuilderPatchOptions)(error) {
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
// Plans the plans property
func (m *PlannerRosterItemRequestBuilder) Plans()(*i2243091c8fa68ba9e7a97bce329cf8353df512034105a6cdb60022d8a3462d3f.PlansRequestBuilder) {
    return i2243091c8fa68ba9e7a97bce329cf8353df512034105a6cdb60022d8a3462d3f.NewPlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PlansById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.planner.rosters.item.plans.item collection
func (m *PlannerRosterItemRequestBuilder) PlansById(id string)(*if82e900011d72e8f7a8566012bb928604f336d22e9e7bd8049c43486ff3d7ff3.PlannerPlanItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerPlan%2Did"] = id
    }
    return if82e900011d72e8f7a8566012bb928604f336d22e9e7bd8049c43486ff3d7ff3.NewPlannerPlanItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
