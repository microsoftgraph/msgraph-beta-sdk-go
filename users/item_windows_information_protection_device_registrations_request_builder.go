package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemWindowsInformationProtectionDeviceRegistrationsRequestBuilder provides operations to manage the windowsInformationProtectionDeviceRegistrations property of the microsoft.graph.user entity.
type ItemWindowsInformationProtectionDeviceRegistrationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemWindowsInformationProtectionDeviceRegistrationsRequestBuilderGetQueryParameters zero or more WIP device registrations that belong to the user.
type ItemWindowsInformationProtectionDeviceRegistrationsRequestBuilderGetQueryParameters struct {
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
// ItemWindowsInformationProtectionDeviceRegistrationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemWindowsInformationProtectionDeviceRegistrationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemWindowsInformationProtectionDeviceRegistrationsRequestBuilderGetQueryParameters
}
// ByWindowsInformationProtectionDeviceRegistrationIdString provides operations to manage the windowsInformationProtectionDeviceRegistrations property of the microsoft.graph.user entity.
func (m *ItemWindowsInformationProtectionDeviceRegistrationsRequestBuilder) ByWindowsInformationProtectionDeviceRegistrationIdString(windowsInformationProtectionDeviceRegistrationId string)(*ItemWindowsInformationProtectionDeviceRegistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if windowsInformationProtectionDeviceRegistrationId != "" {
        urlTplParams["windowsInformationProtectionDeviceRegistration%2Did"] = windowsInformationProtectionDeviceRegistrationId
    }
    return NewItemWindowsInformationProtectionDeviceRegistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemWindowsInformationProtectionDeviceRegistrationsRequestBuilderInternal instantiates a new WindowsInformationProtectionDeviceRegistrationsRequestBuilder and sets the default values.
func NewItemWindowsInformationProtectionDeviceRegistrationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWindowsInformationProtectionDeviceRegistrationsRequestBuilder) {
    m := &ItemWindowsInformationProtectionDeviceRegistrationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/windowsInformationProtectionDeviceRegistrations{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemWindowsInformationProtectionDeviceRegistrationsRequestBuilder instantiates a new WindowsInformationProtectionDeviceRegistrationsRequestBuilder and sets the default values.
func NewItemWindowsInformationProtectionDeviceRegistrationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWindowsInformationProtectionDeviceRegistrationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemWindowsInformationProtectionDeviceRegistrationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemWindowsInformationProtectionDeviceRegistrationsRequestBuilder) Count()(*ItemWindowsInformationProtectionDeviceRegistrationsCountRequestBuilder) {
    return NewItemWindowsInformationProtectionDeviceRegistrationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get zero or more WIP device registrations that belong to the user.
func (m *ItemWindowsInformationProtectionDeviceRegistrationsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemWindowsInformationProtectionDeviceRegistrationsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionDeviceRegistrationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsInformationProtectionDeviceRegistrationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionDeviceRegistrationCollectionResponseable), nil
}
// ToGetRequestInformation zero or more WIP device registrations that belong to the user.
func (m *ItemWindowsInformationProtectionDeviceRegistrationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemWindowsInformationProtectionDeviceRegistrationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
