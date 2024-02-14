package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilder provides operations to manage the assignments property of the microsoft.graph.appleUserInitiatedEnrollmentProfile entity.
type AppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilderGetQueryParameters the list of assignments for this profile.
type AppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// AppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilderGetQueryParameters
}
// AppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAppleEnrollmentProfileAssignmentId provides operations to manage the assignments property of the microsoft.graph.appleUserInitiatedEnrollmentProfile entity.
// returns a *AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder when successful
func (m *AppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilder) ByAppleEnrollmentProfileAssignmentId(appleEnrollmentProfileAssignmentId string)(*AppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if appleEnrollmentProfileAssignmentId != "" {
        urlTplParams["appleEnrollmentProfileAssignment%2Did"] = appleEnrollmentProfileAssignmentId
    }
    return NewAppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilderInternal instantiates a new AppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilder and sets the default values.
func NewAppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilder) {
    m := &AppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/appleUserInitiatedEnrollmentProfiles/{appleUserInitiatedEnrollmentProfile%2Did}/assignments{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilder instantiates a new AppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilder and sets the default values.
func NewAppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AppleUserInitiatedEnrollmentProfilesItemAssignmentsCountRequestBuilder when successful
func (m *AppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilder) Count()(*AppleUserInitiatedEnrollmentProfilesItemAssignmentsCountRequestBuilder) {
    return NewAppleUserInitiatedEnrollmentProfilesItemAssignmentsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of assignments for this profile.
// returns a AppleEnrollmentProfileAssignmentCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilder) Get(ctx context.Context, requestConfiguration *AppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleEnrollmentProfileAssignmentCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAppleEnrollmentProfileAssignmentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleEnrollmentProfileAssignmentCollectionResponseable), nil
}
// Post create new navigation property to assignments for deviceManagement
// returns a AppleEnrollmentProfileAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleEnrollmentProfileAssignmentable, requestConfiguration *AppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleEnrollmentProfileAssignmentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAppleEnrollmentProfileAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleEnrollmentProfileAssignmentable), nil
}
// ToGetRequestInformation the list of assignments for this profile.
// returns a *RequestInformation when successful
func (m *AppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation create new navigation property to assignments for deviceManagement
// returns a *RequestInformation when successful
func (m *AppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleEnrollmentProfileAssignmentable, requestConfiguration *AppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/deviceManagement/appleUserInitiatedEnrollmentProfiles/{appleUserInitiatedEnrollmentProfile%2Did}/assignments", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilder when successful
func (m *AppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilder) WithUrl(rawUrl string)(*AppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilder) {
    return NewAppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
