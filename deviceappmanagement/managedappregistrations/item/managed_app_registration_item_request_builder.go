package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i56070fa6720bc0b45941309da76f8b33eaeab41361bc4d8deae7cf9b20365caa "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item/appliedpolicies"
    id957f7cf90d5e3cc168823ce7a8346428ac2a3905e6406e2ee0dc61e8312d433 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item/operations"
    if8c50fd03afc54f53eb9adc01674971598fcc5e91f81ad3c1573605a4ba9a43c "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item/intendedpolicies"
    i1c13dfabdff5dc012fd7d12bf25406e90dff95d5a2f3f2db4980cf2baccac50e "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item/appliedpolicies/item"
    i4188f43081529892f83979f26fa2b9dca74e78bf1260dae80736bc3d6815fb36 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item/operations/item"
    id0b04ce4d5e56adfc6f008d1080876eacead82cead17447895ea3eb8353b96ac "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item/intendedpolicies/item"
)

// ManagedAppRegistrationItemRequestBuilder provides operations to manage the managedAppRegistrations property of the microsoft.graph.deviceAppManagement entity.
type ManagedAppRegistrationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedAppRegistrationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedAppRegistrationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedAppRegistrationItemRequestBuilderGetQueryParameters the managed app registrations.
type ManagedAppRegistrationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedAppRegistrationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedAppRegistrationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedAppRegistrationItemRequestBuilderGetQueryParameters
}
// ManagedAppRegistrationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedAppRegistrationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppliedPolicies the appliedPolicies property
func (m *ManagedAppRegistrationItemRequestBuilder) AppliedPolicies()(*i56070fa6720bc0b45941309da76f8b33eaeab41361bc4d8deae7cf9b20365caa.AppliedPoliciesRequestBuilder) {
    return i56070fa6720bc0b45941309da76f8b33eaeab41361bc4d8deae7cf9b20365caa.NewAppliedPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppliedPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.managedAppRegistrations.item.appliedPolicies.item collection
func (m *ManagedAppRegistrationItemRequestBuilder) AppliedPoliciesById(id string)(*i1c13dfabdff5dc012fd7d12bf25406e90dff95d5a2f3f2db4980cf2baccac50e.ManagedAppPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppPolicy%2Did"] = id
    }
    return i1c13dfabdff5dc012fd7d12bf25406e90dff95d5a2f3f2db4980cf2baccac50e.NewManagedAppPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewManagedAppRegistrationItemRequestBuilderInternal instantiates a new ManagedAppRegistrationItemRequestBuilder and sets the default values.
func NewManagedAppRegistrationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedAppRegistrationItemRequestBuilder) {
    m := &ManagedAppRegistrationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/managedAppRegistrations/{managedAppRegistration%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedAppRegistrationItemRequestBuilder instantiates a new ManagedAppRegistrationItemRequestBuilder and sets the default values.
func NewManagedAppRegistrationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedAppRegistrationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedAppRegistrationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property managedAppRegistrations for deviceAppManagement
func (m *ManagedAppRegistrationItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property managedAppRegistrations for deviceAppManagement
func (m *ManagedAppRegistrationItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *ManagedAppRegistrationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformationWithRequestConfiguration the managed app registrations.
func (m *ManagedAppRegistrationItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the managed app registrations.
func (m *ManagedAppRegistrationItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ManagedAppRegistrationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property managedAppRegistrations in deviceAppManagement
func (m *ManagedAppRegistrationItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAppRegistrationable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property managedAppRegistrations in deviceAppManagement
func (m *ManagedAppRegistrationItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAppRegistrationable, requestConfiguration *ManagedAppRegistrationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DeleteWithResponseHandler delete navigation property managedAppRegistrations for deviceAppManagement
func (m *ManagedAppRegistrationItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *ManagedAppRegistrationItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property managedAppRegistrations for deviceAppManagement
func (m *ManagedAppRegistrationItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *ManagedAppRegistrationItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// GetWithResponseHandler the managed app registrations.
func (m *ManagedAppRegistrationItemRequestBuilder) GetWithResponseHandler(requestConfiguration *ManagedAppRegistrationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAppRegistrationable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler the managed app registrations.
func (m *ManagedAppRegistrationItemRequestBuilder) GetWithResponseHandler(requestConfiguration *ManagedAppRegistrationItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAppRegistrationable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedAppRegistrationFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAppRegistrationable), nil
}
// IntendedPolicies the intendedPolicies property
func (m *ManagedAppRegistrationItemRequestBuilder) IntendedPolicies()(*if8c50fd03afc54f53eb9adc01674971598fcc5e91f81ad3c1573605a4ba9a43c.IntendedPoliciesRequestBuilder) {
    return if8c50fd03afc54f53eb9adc01674971598fcc5e91f81ad3c1573605a4ba9a43c.NewIntendedPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IntendedPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.managedAppRegistrations.item.intendedPolicies.item collection
func (m *ManagedAppRegistrationItemRequestBuilder) IntendedPoliciesById(id string)(*id0b04ce4d5e56adfc6f008d1080876eacead82cead17447895ea3eb8353b96ac.ManagedAppPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppPolicy%2Did"] = id
    }
    return id0b04ce4d5e56adfc6f008d1080876eacead82cead17447895ea3eb8353b96ac.NewManagedAppPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations the operations property
func (m *ManagedAppRegistrationItemRequestBuilder) Operations()(*id957f7cf90d5e3cc168823ce7a8346428ac2a3905e6406e2ee0dc61e8312d433.OperationsRequestBuilder) {
    return id957f7cf90d5e3cc168823ce7a8346428ac2a3905e6406e2ee0dc61e8312d433.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.managedAppRegistrations.item.operations.item collection
func (m *ManagedAppRegistrationItemRequestBuilder) OperationsById(id string)(*i4188f43081529892f83979f26fa2b9dca74e78bf1260dae80736bc3d6815fb36.ManagedAppOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppOperation%2Did"] = id
    }
    return i4188f43081529892f83979f26fa2b9dca74e78bf1260dae80736bc3d6815fb36.NewManagedAppOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PatchWithResponseHandler update the navigation property managedAppRegistrations in deviceAppManagement
func (m *ManagedAppRegistrationItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAppRegistrationable, requestConfiguration *ManagedAppRegistrationItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property managedAppRegistrations in deviceAppManagement
func (m *ManagedAppRegistrationItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAppRegistrationable, requestConfiguration *ManagedAppRegistrationItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
