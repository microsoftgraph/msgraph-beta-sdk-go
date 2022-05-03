package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i82b6968db7b15686cbc987e9af89e40d6258324a85852e48f8f9c19a6729f0f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/connectedorganizations/item/externalsponsors"
    iff0b6eadf2514d5b5b3edd06d302b534eec4f14682b130a9eb1d95c1c6ef4489 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/connectedorganizations/item/internalsponsors"
    i0dc4c794f51a132efa630bcc7fe954fbd3f2af62e62f7c1557759f789dc562fc "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/connectedorganizations/item/externalsponsors/item"
    iff3bb963dff2078fab9881a67d7525c7d4a7f645e97c901f5515f1e0a89e28b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/connectedorganizations/item/internalsponsors/item"
)

// ConnectedOrganizationItemRequestBuilder provides operations to manage the connectedOrganizations property of the microsoft.graph.entitlementManagement entity.
type ConnectedOrganizationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ConnectedOrganizationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectedOrganizationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ConnectedOrganizationItemRequestBuilderGetQueryParameters represents references to a directory or domain of another organization whose users can request access.
type ConnectedOrganizationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ConnectedOrganizationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectedOrganizationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConnectedOrganizationItemRequestBuilderGetQueryParameters
}
// ConnectedOrganizationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectedOrganizationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConnectedOrganizationItemRequestBuilderInternal instantiates a new ConnectedOrganizationItemRequestBuilder and sets the default values.
func NewConnectedOrganizationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectedOrganizationItemRequestBuilder) {
    m := &ConnectedOrganizationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/connectedOrganizations/{connectedOrganization%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewConnectedOrganizationItemRequestBuilder instantiates a new ConnectedOrganizationItemRequestBuilder and sets the default values.
func NewConnectedOrganizationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectedOrganizationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConnectedOrganizationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property connectedOrganizations for identityGovernance
func (m *ConnectedOrganizationItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property connectedOrganizations for identityGovernance
func (m *ConnectedOrganizationItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *ConnectedOrganizationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformationWithRequestConfiguration represents references to a directory or domain of another organization whose users can request access.
func (m *ConnectedOrganizationItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration represents references to a directory or domain of another organization whose users can request access.
func (m *ConnectedOrganizationItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ConnectedOrganizationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property connectedOrganizations in identityGovernance
func (m *ConnectedOrganizationItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectedOrganizationable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property connectedOrganizations in identityGovernance
func (m *ConnectedOrganizationItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectedOrganizationable, requestConfiguration *ConnectedOrganizationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DeleteWithResponseHandler delete navigation property connectedOrganizations for identityGovernance
func (m *ConnectedOrganizationItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *ConnectedOrganizationItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property connectedOrganizations for identityGovernance
func (m *ConnectedOrganizationItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *ConnectedOrganizationItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// ExternalSponsors the externalSponsors property
func (m *ConnectedOrganizationItemRequestBuilder) ExternalSponsors()(*i82b6968db7b15686cbc987e9af89e40d6258324a85852e48f8f9c19a6729f0f0.ExternalSponsorsRequestBuilder) {
    return i82b6968db7b15686cbc987e9af89e40d6258324a85852e48f8f9c19a6729f0f0.NewExternalSponsorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExternalSponsorsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.connectedOrganizations.item.externalSponsors.item collection
func (m *ConnectedOrganizationItemRequestBuilder) ExternalSponsorsById(id string)(*i0dc4c794f51a132efa630bcc7fe954fbd3f2af62e62f7c1557759f789dc562fc.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i0dc4c794f51a132efa630bcc7fe954fbd3f2af62e62f7c1557759f789dc562fc.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GetWithResponseHandler represents references to a directory or domain of another organization whose users can request access.
func (m *ConnectedOrganizationItemRequestBuilder) GetWithResponseHandler(requestConfiguration *ConnectedOrganizationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectedOrganizationable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler represents references to a directory or domain of another organization whose users can request access.
func (m *ConnectedOrganizationItemRequestBuilder) GetWithResponseHandler(requestConfiguration *ConnectedOrganizationItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectedOrganizationable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateConnectedOrganizationFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectedOrganizationable), nil
}
// InternalSponsors the internalSponsors property
func (m *ConnectedOrganizationItemRequestBuilder) InternalSponsors()(*iff0b6eadf2514d5b5b3edd06d302b534eec4f14682b130a9eb1d95c1c6ef4489.InternalSponsorsRequestBuilder) {
    return iff0b6eadf2514d5b5b3edd06d302b534eec4f14682b130a9eb1d95c1c6ef4489.NewInternalSponsorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InternalSponsorsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.connectedOrganizations.item.internalSponsors.item collection
func (m *ConnectedOrganizationItemRequestBuilder) InternalSponsorsById(id string)(*iff3bb963dff2078fab9881a67d7525c7d4a7f645e97c901f5515f1e0a89e28b6.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return iff3bb963dff2078fab9881a67d7525c7d4a7f645e97c901f5515f1e0a89e28b6.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PatchWithResponseHandler update the navigation property connectedOrganizations in identityGovernance
func (m *ConnectedOrganizationItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectedOrganizationable, requestConfiguration *ConnectedOrganizationItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property connectedOrganizations in identityGovernance
func (m *ConnectedOrganizationItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConnectedOrganizationable, requestConfiguration *ConnectedOrganizationItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
