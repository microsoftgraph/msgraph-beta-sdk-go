package organization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OrganizationItemSetMobileDeviceManagementAuthorityRequestBuilder provides operations to call the setMobileDeviceManagementAuthority method.
type OrganizationItemSetMobileDeviceManagementAuthorityRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OrganizationItemSetMobileDeviceManagementAuthorityRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrganizationItemSetMobileDeviceManagementAuthorityRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOrganizationItemSetMobileDeviceManagementAuthorityRequestBuilderInternal instantiates a new SetMobileDeviceManagementAuthorityRequestBuilder and sets the default values.
func NewOrganizationItemSetMobileDeviceManagementAuthorityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationItemSetMobileDeviceManagementAuthorityRequestBuilder) {
    m := &OrganizationItemSetMobileDeviceManagementAuthorityRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/organization/{organization%2Did}/microsoft.graph.setMobileDeviceManagementAuthority";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOrganizationItemSetMobileDeviceManagementAuthorityRequestBuilder instantiates a new SetMobileDeviceManagementAuthorityRequestBuilder and sets the default values.
func NewOrganizationItemSetMobileDeviceManagementAuthorityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationItemSetMobileDeviceManagementAuthorityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrganizationItemSetMobileDeviceManagementAuthorityRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation set mobile device management authority
func (m *OrganizationItemSetMobileDeviceManagementAuthorityRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *OrganizationItemSetMobileDeviceManagementAuthorityRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post set mobile device management authority
func (m *OrganizationItemSetMobileDeviceManagementAuthorityRequestBuilder) Post(ctx context.Context, requestConfiguration *OrganizationItemSetMobileDeviceManagementAuthorityRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationItemSetMobileDeviceManagementAuthorityResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOrganizationItemSetMobileDeviceManagementAuthorityResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrganizationItemSetMobileDeviceManagementAuthorityResponseable), nil
}
