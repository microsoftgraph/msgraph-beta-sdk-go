package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileAppsItemGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder provides operations to call the getRelatedAppStates method.
type MobileAppsItemGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileAppsItemGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderGetQueryParameters invoke function getRelatedAppStates
type MobileAppsItemGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// MobileAppsItemGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileAppsItemGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderGetQueryParameters
}
// NewMobileAppsItemGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderInternal instantiates a new GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder and sets the default values.
func NewMobileAppsItemGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, deviceId *string, userPrincipalName *string)(*MobileAppsItemGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder) {
    m := &MobileAppsItemGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/getRelatedAppStates(userPrincipalName='{userPrincipalName}',deviceId='{deviceId}'){?%24top,%24skip,%24search,%24filter,%24count}", pathParameters),
    }
    if deviceId != nil {
        urlTplParams["deviceId"] = *deviceId
    }
    if userPrincipalName != nil {
        urlTplParams["userPrincipalName"] = *userPrincipalName
    }
    return m
}
// NewMobileAppsItemGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder instantiates a new GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder and sets the default values.
func NewMobileAppsItemGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppsItemGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get invoke function getRelatedAppStates
func (m *MobileAppsItemGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppsItemGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderGetRequestConfiguration)(MobileAppsItemGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateMobileAppsItemGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MobileAppsItemGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdResponseable), nil
}
// ToGetRequestInformation invoke function getRelatedAppStates
func (m *MobileAppsItemGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileAppsItemGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
