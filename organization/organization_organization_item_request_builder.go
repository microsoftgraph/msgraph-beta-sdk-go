package organization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OrganizationOrganizationItemRequestBuilder provides operations to manage the collection of organization entities.
type OrganizationOrganizationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OrganizationOrganizationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrganizationOrganizationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OrganizationOrganizationItemRequestBuilderGetQueryParameters get the properties and relationships of the currently authenticated organization. Since the **organization** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in an **organization** instance.
type OrganizationOrganizationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OrganizationOrganizationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrganizationOrganizationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OrganizationOrganizationItemRequestBuilderGetQueryParameters
}
// OrganizationOrganizationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrganizationOrganizationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivateService provides operations to call the activateService method.
func (m *OrganizationOrganizationItemRequestBuilder) ActivateService()(*OrganizationItemActivateServiceRequestBuilder) {
    return NewOrganizationItemActivateServiceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Branding provides operations to manage the branding property of the microsoft.graph.organization entity.
func (m *OrganizationOrganizationItemRequestBuilder) Branding()(*OrganizationItemBrandingRequestBuilder) {
    return NewOrganizationItemBrandingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CertificateBasedAuthConfiguration provides operations to manage the certificateBasedAuthConfiguration property of the microsoft.graph.organization entity.
func (m *OrganizationOrganizationItemRequestBuilder) CertificateBasedAuthConfiguration()(*OrganizationItemCertificateBasedAuthConfigurationRequestBuilder) {
    return NewOrganizationItemCertificateBasedAuthConfigurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CertificateBasedAuthConfigurationById provides operations to manage the certificateBasedAuthConfiguration property of the microsoft.graph.organization entity.
func (m *OrganizationOrganizationItemRequestBuilder) CertificateBasedAuthConfigurationById(id string)(*OrganizationItemCertificateBasedAuthConfigurationCertificateBasedAuthConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["certificateBasedAuthConfiguration%2Did"] = id
    }
    return NewOrganizationItemCertificateBasedAuthConfigurationCertificateBasedAuthConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *OrganizationOrganizationItemRequestBuilder) CheckMemberGroups()(*OrganizationItemCheckMemberGroupsRequestBuilder) {
    return NewOrganizationItemCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *OrganizationOrganizationItemRequestBuilder) CheckMemberObjects()(*OrganizationItemCheckMemberObjectsRequestBuilder) {
    return NewOrganizationItemCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOrganizationOrganizationItemRequestBuilderInternal instantiates a new OrganizationItemRequestBuilder and sets the default values.
func NewOrganizationOrganizationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationOrganizationItemRequestBuilder) {
    m := &OrganizationOrganizationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/organization/{organization%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOrganizationOrganizationItemRequestBuilder instantiates a new OrganizationItemRequestBuilder and sets the default values.
func NewOrganizationOrganizationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationOrganizationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrganizationOrganizationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from organization
func (m *OrganizationOrganizationItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *OrganizationOrganizationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get the properties and relationships of the currently authenticated organization. Since the **organization** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in an **organization** instance.
func (m *OrganizationOrganizationItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *OrganizationOrganizationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the properties of the currently authenticated organization. In this case, `organization` is defined as a collection of exactly one record, and so its **ID** must be specified in the request.  The **ID** is also known as the **tenantId** of the organization.
func (m *OrganizationOrganizationItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Organizationable, requestConfiguration *OrganizationOrganizationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete entity from organization
func (m *OrganizationOrganizationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *OrganizationOrganizationItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.organization entity.
func (m *OrganizationOrganizationItemRequestBuilder) Extensions()(*OrganizationItemExtensionsRequestBuilder) {
    return NewOrganizationItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.organization entity.
func (m *OrganizationOrganizationItemRequestBuilder) ExtensionsById(id string)(*OrganizationItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewOrganizationItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get the properties and relationships of the currently authenticated organization. Since the **organization** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in an **organization** instance.
func (m *OrganizationOrganizationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *OrganizationOrganizationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Organizationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOrganizationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Organizationable), nil
}
// GetMemberGroups provides operations to call the getMemberGroups method.
func (m *OrganizationOrganizationItemRequestBuilder) GetMemberGroups()(*OrganizationItemGetMemberGroupsRequestBuilder) {
    return NewOrganizationItemGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects provides operations to call the getMemberObjects method.
func (m *OrganizationOrganizationItemRequestBuilder) GetMemberObjects()(*OrganizationItemGetMemberObjectsRequestBuilder) {
    return NewOrganizationItemGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the properties of the currently authenticated organization. In this case, `organization` is defined as a collection of exactly one record, and so its **ID** must be specified in the request.  The **ID** is also known as the **tenantId** of the organization.
func (m *OrganizationOrganizationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Organizationable, requestConfiguration *OrganizationOrganizationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Organizationable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOrganizationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Organizationable), nil
}
// Restore provides operations to call the restore method.
func (m *OrganizationOrganizationItemRequestBuilder) Restore()(*OrganizationItemRestoreRequestBuilder) {
    return NewOrganizationItemRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetMobileDeviceManagementAuthority provides operations to call the setMobileDeviceManagementAuthority method.
func (m *OrganizationOrganizationItemRequestBuilder) SetMobileDeviceManagementAuthority()(*OrganizationItemSetMobileDeviceManagementAuthorityRequestBuilder) {
    return NewOrganizationItemSetMobileDeviceManagementAuthorityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Settings provides operations to manage the settings property of the microsoft.graph.organization entity.
func (m *OrganizationOrganizationItemRequestBuilder) Settings()(*OrganizationItemSettingsRequestBuilder) {
    return NewOrganizationItemSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
