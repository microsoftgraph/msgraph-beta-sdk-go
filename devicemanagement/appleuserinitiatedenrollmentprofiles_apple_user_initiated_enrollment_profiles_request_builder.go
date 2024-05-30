package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilder provides operations to manage the appleUserInitiatedEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
type AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilderGetQueryParameters apple user initiated enrollment profiles
type AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilderGetQueryParameters struct {
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
// AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilderGetQueryParameters
}
// AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAppleUserInitiatedEnrollmentProfileId provides operations to manage the appleUserInitiatedEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
// returns a *AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder when successful
func (m *AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilder) ByAppleUserInitiatedEnrollmentProfileId(appleUserInitiatedEnrollmentProfileId string)(*AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if appleUserInitiatedEnrollmentProfileId != "" {
        urlTplParams["appleUserInitiatedEnrollmentProfile%2Did"] = appleUserInitiatedEnrollmentProfileId
    }
    return NewAppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilderInternal instantiates a new AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilder and sets the default values.
func NewAppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilder) {
    m := &AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/appleUserInitiatedEnrollmentProfiles{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilder instantiates a new AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilder and sets the default values.
func NewAppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AppleuserinitiatedenrollmentprofilesCountRequestBuilder when successful
func (m *AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilder) Count()(*AppleuserinitiatedenrollmentprofilesCountRequestBuilder) {
    return NewAppleuserinitiatedenrollmentprofilesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get apple user initiated enrollment profiles
// returns a AppleUserInitiatedEnrollmentProfileCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilder) Get(ctx context.Context, requestConfiguration *AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleUserInitiatedEnrollmentProfileCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAppleUserInitiatedEnrollmentProfileCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleUserInitiatedEnrollmentProfileCollectionResponseable), nil
}
// Post create new navigation property to appleUserInitiatedEnrollmentProfiles for deviceManagement
// returns a AppleUserInitiatedEnrollmentProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleUserInitiatedEnrollmentProfileable, requestConfiguration *AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleUserInitiatedEnrollmentProfileable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAppleUserInitiatedEnrollmentProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleUserInitiatedEnrollmentProfileable), nil
}
// ToGetRequestInformation apple user initiated enrollment profiles
// returns a *RequestInformation when successful
func (m *AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to appleUserInitiatedEnrollmentProfiles for deviceManagement
// returns a *RequestInformation when successful
func (m *AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleUserInitiatedEnrollmentProfileable, requestConfiguration *AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilder when successful
func (m *AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilder) WithUrl(rawUrl string)(*AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilder) {
    return NewAppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
