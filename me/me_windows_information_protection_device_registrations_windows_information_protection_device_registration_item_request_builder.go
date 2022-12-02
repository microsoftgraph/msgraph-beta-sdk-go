package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeWindowsInformationProtectionDeviceRegistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder provides operations to manage the windowsInformationProtectionDeviceRegistrations property of the microsoft.graph.user entity.
type MeWindowsInformationProtectionDeviceRegistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeWindowsInformationProtectionDeviceRegistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderGetQueryParameters zero or more WIP device registrations that belong to the user.
type MeWindowsInformationProtectionDeviceRegistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeWindowsInformationProtectionDeviceRegistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeWindowsInformationProtectionDeviceRegistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeWindowsInformationProtectionDeviceRegistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderGetQueryParameters
}
// NewMeWindowsInformationProtectionDeviceRegistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderInternal instantiates a new WindowsInformationProtectionDeviceRegistrationItemRequestBuilder and sets the default values.
func NewMeWindowsInformationProtectionDeviceRegistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeWindowsInformationProtectionDeviceRegistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder) {
    m := &MeWindowsInformationProtectionDeviceRegistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/windowsInformationProtectionDeviceRegistrations/{windowsInformationProtectionDeviceRegistration%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeWindowsInformationProtectionDeviceRegistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder instantiates a new WindowsInformationProtectionDeviceRegistrationItemRequestBuilder and sets the default values.
func NewMeWindowsInformationProtectionDeviceRegistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeWindowsInformationProtectionDeviceRegistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeWindowsInformationProtectionDeviceRegistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation zero or more WIP device registrations that belong to the user.
func (m *MeWindowsInformationProtectionDeviceRegistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MeWindowsInformationProtectionDeviceRegistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get zero or more WIP device registrations that belong to the user.
func (m *MeWindowsInformationProtectionDeviceRegistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MeWindowsInformationProtectionDeviceRegistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionDeviceRegistrationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsInformationProtectionDeviceRegistrationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionDeviceRegistrationable), nil
}
