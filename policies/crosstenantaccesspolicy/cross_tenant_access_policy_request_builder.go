package crosstenantaccesspolicy

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i12ec00c0cda085cf3d19e747ae27322079cbb18e5b556d17b1a8a5d083c42b4f "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/crosstenantaccesspolicy/partners"
    i9012de60632d1258f2b5642de666977d54e3da254d584ad456e64c40f801e738 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/crosstenantaccesspolicy/default_escaped"
    i671445f12e1144f0707e713947da70fbe18307579db947c865515610f6068a76 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/crosstenantaccesspolicy/partners/item"
)

// CrossTenantAccessPolicyRequestBuilder provides operations to manage the crossTenantAccessPolicy property of the microsoft.graph.policyRoot entity.
type CrossTenantAccessPolicyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CrossTenantAccessPolicyRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CrossTenantAccessPolicyRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CrossTenantAccessPolicyRequestBuilderGetQueryParameters the custom rules that define an access scenario when interacting with external Azure AD tenants.
type CrossTenantAccessPolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CrossTenantAccessPolicyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CrossTenantAccessPolicyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CrossTenantAccessPolicyRequestBuilderGetQueryParameters
}
// CrossTenantAccessPolicyRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CrossTenantAccessPolicyRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCrossTenantAccessPolicyRequestBuilderInternal instantiates a new CrossTenantAccessPolicyRequestBuilder and sets the default values.
func NewCrossTenantAccessPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CrossTenantAccessPolicyRequestBuilder) {
    m := &CrossTenantAccessPolicyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/crossTenantAccessPolicy{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCrossTenantAccessPolicyRequestBuilder instantiates a new CrossTenantAccessPolicyRequestBuilder and sets the default values.
func NewCrossTenantAccessPolicyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CrossTenantAccessPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCrossTenantAccessPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property crossTenantAccessPolicy for policies
func (m *CrossTenantAccessPolicyRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property crossTenantAccessPolicy for policies
func (m *CrossTenantAccessPolicyRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *CrossTenantAccessPolicyRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the custom rules that define an access scenario when interacting with external Azure AD tenants.
func (m *CrossTenantAccessPolicyRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the custom rules that define an access scenario when interacting with external Azure AD tenants.
func (m *CrossTenantAccessPolicyRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *CrossTenantAccessPolicyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property crossTenantAccessPolicy in policies
func (m *CrossTenantAccessPolicyRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CrossTenantAccessPolicyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property crossTenantAccessPolicy in policies
func (m *CrossTenantAccessPolicyRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CrossTenantAccessPolicyable, requestConfiguration *CrossTenantAccessPolicyRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Default_escaped the default property
func (m *CrossTenantAccessPolicyRequestBuilder) Default_escaped()(*i9012de60632d1258f2b5642de666977d54e3da254d584ad456e64c40f801e738.DefaultRequestBuilder) {
    return i9012de60632d1258f2b5642de666977d54e3da254d584ad456e64c40f801e738.NewDefaultRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property crossTenantAccessPolicy for policies
func (m *CrossTenantAccessPolicyRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property crossTenantAccessPolicy for policies
func (m *CrossTenantAccessPolicyRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *CrossTenantAccessPolicyRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the custom rules that define an access scenario when interacting with external Azure AD tenants.
func (m *CrossTenantAccessPolicyRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CrossTenantAccessPolicyable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the custom rules that define an access scenario when interacting with external Azure AD tenants.
func (m *CrossTenantAccessPolicyRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *CrossTenantAccessPolicyRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CrossTenantAccessPolicyable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCrossTenantAccessPolicyFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CrossTenantAccessPolicyable), nil
}
// Partners the partners property
func (m *CrossTenantAccessPolicyRequestBuilder) Partners()(*i12ec00c0cda085cf3d19e747ae27322079cbb18e5b556d17b1a8a5d083c42b4f.PartnersRequestBuilder) {
    return i12ec00c0cda085cf3d19e747ae27322079cbb18e5b556d17b1a8a5d083c42b4f.NewPartnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PartnersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.policies.crossTenantAccessPolicy.partners.item collection
func (m *CrossTenantAccessPolicyRequestBuilder) PartnersById(id string)(*i671445f12e1144f0707e713947da70fbe18307579db947c865515610f6068a76.CrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["crossTenantAccessPolicyConfigurationPartner%2DtenantId"] = id
    }
    return i671445f12e1144f0707e713947da70fbe18307579db947c865515610f6068a76.NewCrossTenantAccessPolicyConfigurationPartnerTenantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property crossTenantAccessPolicy in policies
func (m *CrossTenantAccessPolicyRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CrossTenantAccessPolicyable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property crossTenantAccessPolicy in policies
func (m *CrossTenantAccessPolicyRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CrossTenantAccessPolicyable, requestConfiguration *CrossTenantAccessPolicyRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
