package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ibe61fbd02c16db3e5f015184d8991d363a774ed55310ae904b00e97143389ee8 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/termsandconditions/item/groupassignments"
    ide71c0a56dd4ade2d81104d2ecb208e8db8c2910aaea2f231edb2cf623477c53 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/termsandconditions/item/assignments"
    idf88d5dcb4048ac7ff79419fd15c7d58e83a14f8b409ce16869ccf4204ab829f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/termsandconditions/item/acceptancestatuses"
    i5c99d94fd8ce128189bc114b673df4e56578860abd3b3eadccba77c18952a47d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/termsandconditions/item/groupassignments/item"
    i605e35cc4db71e7fa4db1f4759e05ab605fd4596a0a820b7a556521d22026745 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/termsandconditions/item/acceptancestatuses/item"
    idae59aa73e78d2aa3e7943286419e2c7b77cd5e337f001c8305c745502c3740f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/termsandconditions/item/assignments/item"
)

// TermsAndConditionsItemRequestBuilder provides operations to manage the termsAndConditions property of the microsoft.graph.deviceManagement entity.
type TermsAndConditionsItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TermsAndConditionsItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsAndConditionsItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TermsAndConditionsItemRequestBuilderGetQueryParameters the terms and conditions associated with device management of the company.
type TermsAndConditionsItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TermsAndConditionsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsAndConditionsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TermsAndConditionsItemRequestBuilderGetQueryParameters
}
// TermsAndConditionsItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsAndConditionsItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AcceptanceStatuses the acceptanceStatuses property
func (m *TermsAndConditionsItemRequestBuilder) AcceptanceStatuses()(*idf88d5dcb4048ac7ff79419fd15c7d58e83a14f8b409ce16869ccf4204ab829f.AcceptanceStatusesRequestBuilder) {
    return idf88d5dcb4048ac7ff79419fd15c7d58e83a14f8b409ce16869ccf4204ab829f.NewAcceptanceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AcceptanceStatusesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.termsAndConditions.item.acceptanceStatuses.item collection
func (m *TermsAndConditionsItemRequestBuilder) AcceptanceStatusesById(id string)(*i605e35cc4db71e7fa4db1f4759e05ab605fd4596a0a820b7a556521d22026745.TermsAndConditionsAcceptanceStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["termsAndConditionsAcceptanceStatus%2Did"] = id
    }
    return i605e35cc4db71e7fa4db1f4759e05ab605fd4596a0a820b7a556521d22026745.NewTermsAndConditionsAcceptanceStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Assignments the assignments property
func (m *TermsAndConditionsItemRequestBuilder) Assignments()(*ide71c0a56dd4ade2d81104d2ecb208e8db8c2910aaea2f231edb2cf623477c53.AssignmentsRequestBuilder) {
    return ide71c0a56dd4ade2d81104d2ecb208e8db8c2910aaea2f231edb2cf623477c53.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.termsAndConditions.item.assignments.item collection
func (m *TermsAndConditionsItemRequestBuilder) AssignmentsById(id string)(*idae59aa73e78d2aa3e7943286419e2c7b77cd5e337f001c8305c745502c3740f.TermsAndConditionsAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["termsAndConditionsAssignment%2Did"] = id
    }
    return idae59aa73e78d2aa3e7943286419e2c7b77cd5e337f001c8305c745502c3740f.NewTermsAndConditionsAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTermsAndConditionsItemRequestBuilderInternal instantiates a new TermsAndConditionsItemRequestBuilder and sets the default values.
func NewTermsAndConditionsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TermsAndConditionsItemRequestBuilder) {
    m := &TermsAndConditionsItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/termsAndConditions/{termsAndConditions%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTermsAndConditionsItemRequestBuilder instantiates a new TermsAndConditionsItemRequestBuilder and sets the default values.
func NewTermsAndConditionsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TermsAndConditionsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTermsAndConditionsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property termsAndConditions for deviceManagement
func (m *TermsAndConditionsItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *TermsAndConditionsItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the terms and conditions associated with device management of the company.
func (m *TermsAndConditionsItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TermsAndConditionsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property termsAndConditions in deviceManagement
func (m *TermsAndConditionsItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TermsAndConditionsable, requestConfiguration *TermsAndConditionsItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property termsAndConditions for deviceManagement
func (m *TermsAndConditionsItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TermsAndConditionsItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the terms and conditions associated with device management of the company.
func (m *TermsAndConditionsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TermsAndConditionsItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TermsAndConditionsable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTermsAndConditionsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TermsAndConditionsable), nil
}
// GroupAssignments the groupAssignments property
func (m *TermsAndConditionsItemRequestBuilder) GroupAssignments()(*ibe61fbd02c16db3e5f015184d8991d363a774ed55310ae904b00e97143389ee8.GroupAssignmentsRequestBuilder) {
    return ibe61fbd02c16db3e5f015184d8991d363a774ed55310ae904b00e97143389ee8.NewGroupAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.termsAndConditions.item.groupAssignments.item collection
func (m *TermsAndConditionsItemRequestBuilder) GroupAssignmentsById(id string)(*i5c99d94fd8ce128189bc114b673df4e56578860abd3b3eadccba77c18952a47d.TermsAndConditionsGroupAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["termsAndConditionsGroupAssignment%2Did"] = id
    }
    return i5c99d94fd8ce128189bc114b673df4e56578860abd3b3eadccba77c18952a47d.NewTermsAndConditionsGroupAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property termsAndConditions in deviceManagement
func (m *TermsAndConditionsItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TermsAndConditionsable, requestConfiguration *TermsAndConditionsItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TermsAndConditionsable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTermsAndConditionsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TermsAndConditionsable), nil
}
