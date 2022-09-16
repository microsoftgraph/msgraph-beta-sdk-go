package mobileapps

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i086ce4f71a728d4b38b48a7b436f5263afb5f1bd2c43d24188ace8c90c850640 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/managedmobilelobapp"
    i149429e3eafc16419f958b84dc22e5575cc97c96b553de6bc0dbce71d8dadd3c "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/count"
    i2f59ddb12d33939283eeb5185ecff0dbbce661e9e5c21c6508e809c2b051b944 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/haspayloadlinks"
    i3f1ef83e5a3c636aadac73e14a41959080d06f3c1f972ef3fc82445c7fc94f98 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/getmobileappcountwithstatus"
    i4125e00a7eb603812efbd75551680674e01278ff362db14d1fd91505cb275200 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/gettopmobileappswithstatuswithcount"
    i953b63f91b8c549b6ab097bad0bfd0ac3087a3511b8dc0e7fc54c77eba2c6e1e "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/mobilelobapp"
    id2162c282756356f4cf3805ba782e3632722a8ad528000abd031c5fae5bc20fe "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/validatexml"
)

// MobileAppsRequestBuilder provides operations to manage the mobileApps property of the microsoft.graph.deviceAppManagement entity.
type MobileAppsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MobileAppsRequestBuilderGetQueryParameters the mobile apps.
type MobileAppsRequestBuilderGetQueryParameters struct {
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
// MobileAppsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileAppsRequestBuilderGetQueryParameters
}
// MobileAppsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileAppsRequestBuilderInternal instantiates a new MobileAppsRequestBuilder and sets the default values.
func NewMobileAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsRequestBuilder) {
    m := &MobileAppsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileApps{?%24top*,%24skip*,%24search*,%24filter*,%24count*,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMobileAppsRequestBuilder instantiates a new MobileAppsRequestBuilder and sets the default values.
func NewMobileAppsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *MobileAppsRequestBuilder) Count()(*i149429e3eafc16419f958b84dc22e5575cc97c96b553de6bc0dbce71d8dadd3c.CountRequestBuilder) {
    return i149429e3eafc16419f958b84dc22e5575cc97c96b553de6bc0dbce71d8dadd3c.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the mobile apps.
func (m *MobileAppsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the mobile apps.
func (m *MobileAppsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *MobileAppsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to mobileApps for deviceAppManagement
func (m *MobileAppsRequestBuilder) CreatePostRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to mobileApps for deviceAppManagement
func (m *MobileAppsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable, requestConfiguration *MobileAppsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get the mobile apps.
func (m *MobileAppsRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppCollectionResponseable), nil
}
// GetMobileAppCountWithStatus provides operations to call the getMobileAppCount method.
func (m *MobileAppsRequestBuilder) GetMobileAppCountWithStatus(status *string)(*i3f1ef83e5a3c636aadac73e14a41959080d06f3c1f972ef3fc82445c7fc94f98.GetMobileAppCountWithStatusRequestBuilder) {
    return i3f1ef83e5a3c636aadac73e14a41959080d06f3c1f972ef3fc82445c7fc94f98.NewGetMobileAppCountWithStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter, status);
}
// GetTopMobileAppsWithStatusWithCount provides operations to call the getTopMobileApps method.
func (m *MobileAppsRequestBuilder) GetTopMobileAppsWithStatusWithCount(count *int64, status *string)(*i4125e00a7eb603812efbd75551680674e01278ff362db14d1fd91505cb275200.GetTopMobileAppsWithStatusWithCountRequestBuilder) {
    return i4125e00a7eb603812efbd75551680674e01278ff362db14d1fd91505cb275200.NewGetTopMobileAppsWithStatusWithCountRequestBuilderInternal(m.pathParameters, m.requestAdapter, count, status);
}
// HasPayloadLinks the hasPayloadLinks property
func (m *MobileAppsRequestBuilder) HasPayloadLinks()(*i2f59ddb12d33939283eeb5185ecff0dbbce661e9e5c21c6508e809c2b051b944.HasPayloadLinksRequestBuilder) {
    return i2f59ddb12d33939283eeb5185ecff0dbbce661e9e5c21c6508e809c2b051b944.NewHasPayloadLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedMobileLobApp the managedMobileLobApp property
func (m *MobileAppsRequestBuilder) ManagedMobileLobApp()(*i086ce4f71a728d4b38b48a7b436f5263afb5f1bd2c43d24188ace8c90c850640.ManagedMobileLobAppRequestBuilder) {
    return i086ce4f71a728d4b38b48a7b436f5263afb5f1bd2c43d24188ace8c90c850640.NewManagedMobileLobAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileLobApp the mobileLobApp property
func (m *MobileAppsRequestBuilder) MobileLobApp()(*i953b63f91b8c549b6ab097bad0bfd0ac3087a3511b8dc0e7fc54c77eba2c6e1e.MobileLobAppRequestBuilder) {
    return i953b63f91b8c549b6ab097bad0bfd0ac3087a3511b8dc0e7fc54c77eba2c6e1e.NewMobileLobAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to mobileApps for deviceAppManagement
func (m *MobileAppsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable, requestConfiguration *MobileAppsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable), nil
}
// ValidateXml the validateXml property
func (m *MobileAppsRequestBuilder) ValidateXml()(*id2162c282756356f4cf3805ba782e3632722a8ad528000abd031c5fae5bc20fe.ValidateXmlRequestBuilder) {
    return id2162c282756356f4cf3805ba782e3632722a8ad528000abd031c5fae5bc20fe.NewValidateXmlRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
