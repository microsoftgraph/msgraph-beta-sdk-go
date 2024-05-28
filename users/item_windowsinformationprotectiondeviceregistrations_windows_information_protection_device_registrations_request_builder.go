package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder provides operations to manage the windowsInformationProtectionDeviceRegistrations property of the microsoft.graph.user entity.
type ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilderGetQueryParameters zero or more WIP device registrations that belong to the user.
type ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilderGetQueryParameters struct {
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
// ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilderGetQueryParameters
}
// ByWindowsInformationProtectionDeviceRegistrationId provides operations to manage the windowsInformationProtectionDeviceRegistrations property of the microsoft.graph.user entity.
// returns a *ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder when successful
func (m *ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder) ByWindowsInformationProtectionDeviceRegistrationId(windowsInformationProtectionDeviceRegistrationId string)(*ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if windowsInformationProtectionDeviceRegistrationId != "" {
        urlTplParams["windowsInformationProtectionDeviceRegistration%2Did"] = windowsInformationProtectionDeviceRegistrationId
    }
    return NewItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilderInternal instantiates a new ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder and sets the default values.
func NewItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder) {
    m := &ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/windowsInformationProtectionDeviceRegistrations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder instantiates a new ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder and sets the default values.
func NewItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemWindowsinformationprotectiondeviceregistrationsCountRequestBuilder when successful
func (m *ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder) Count()(*ItemWindowsinformationprotectiondeviceregistrationsCountRequestBuilder) {
    return NewItemWindowsinformationprotectiondeviceregistrationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get zero or more WIP device registrations that belong to the user.
// returns a WindowsInformationProtectionDeviceRegistrationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionDeviceRegistrationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// returns a *RequestInformation when successful
func (m *ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder when successful
func (m *ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder) WithUrl(rawUrl string)(*ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder) {
    return NewItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
