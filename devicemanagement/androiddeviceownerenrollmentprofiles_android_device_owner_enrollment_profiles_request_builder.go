package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilder provides operations to manage the androidDeviceOwnerEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
type AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilderGetQueryParameters android device owner enrollment profile entities.
type AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilderGetQueryParameters struct {
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
// AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilderGetQueryParameters
}
// AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAndroidDeviceOwnerEnrollmentProfileId provides operations to manage the androidDeviceOwnerEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
// returns a *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder when successful
func (m *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilder) ByAndroidDeviceOwnerEnrollmentProfileId(androidDeviceOwnerEnrollmentProfileId string)(*AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if androidDeviceOwnerEnrollmentProfileId != "" {
        urlTplParams["androidDeviceOwnerEnrollmentProfile%2Did"] = androidDeviceOwnerEnrollmentProfileId
    }
    return NewAndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilderInternal instantiates a new AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilder and sets the default values.
func NewAndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilder) {
    m := &AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/androidDeviceOwnerEnrollmentProfiles{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilder instantiates a new AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilder and sets the default values.
func NewAndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AndroiddeviceownerenrollmentprofilesCountRequestBuilder when successful
func (m *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilder) Count()(*AndroiddeviceownerenrollmentprofilesCountRequestBuilder) {
    return NewAndroiddeviceownerenrollmentprofilesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get android device owner enrollment profile entities.
// returns a AndroidDeviceOwnerEnrollmentProfileCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilder) Get(ctx context.Context, requestConfiguration *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidDeviceOwnerEnrollmentProfileCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAndroidDeviceOwnerEnrollmentProfileCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidDeviceOwnerEnrollmentProfileCollectionResponseable), nil
}
// Post create new navigation property to androidDeviceOwnerEnrollmentProfiles for deviceManagement
// returns a AndroidDeviceOwnerEnrollmentProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidDeviceOwnerEnrollmentProfileable, requestConfiguration *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidDeviceOwnerEnrollmentProfileable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAndroidDeviceOwnerEnrollmentProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidDeviceOwnerEnrollmentProfileable), nil
}
// ToGetRequestInformation android device owner enrollment profile entities.
// returns a *RequestInformation when successful
func (m *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to androidDeviceOwnerEnrollmentProfiles for deviceManagement
// returns a *RequestInformation when successful
func (m *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidDeviceOwnerEnrollmentProfileable, requestConfiguration *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilder when successful
func (m *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilder) WithUrl(rawUrl string)(*AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilder) {
    return NewAndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
