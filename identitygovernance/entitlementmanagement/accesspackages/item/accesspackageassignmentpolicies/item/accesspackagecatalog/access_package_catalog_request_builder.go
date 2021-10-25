package accesspackagecatalog

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5d568d21b7fe91619b09f9b4ca34ecbcc735854efc2624376c1d8f5bfa5159d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/accesspackageassignmentpolicies/item/accesspackagecatalog/accesspackages"
    i76ec8ccbad7aee0d645249a9ca5df55114b4d9942d2ce225006a0ffa47878bb8 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/accesspackageassignmentpolicies/item/accesspackagecatalog/accesspackageresourcescopes"
    ibe013da09e7701ba00a47c90a6b10900a06ed1137e6129a7c9226e3a07f79403 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/accesspackageassignmentpolicies/item/accesspackagecatalog/accesspackageresources"
    if690c3853d6776a6f085a19f85d182249f323cebba342ab176f22fd85a7b8f83 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/accesspackageassignmentpolicies/item/accesspackagecatalog/accesspackageresourceroles"
    ib16fad2499a8a76bd9087a69c65f1fb2f5c8e4d0a12a18be523e8674a8e6005f "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/accesspackageassignmentpolicies/item/accesspackagecatalog/accesspackageresourcescopes/item"
    icc355b947ff31982a2b7459e8a793ab47ac4d684947820c8068459683087d33f "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/accesspackageassignmentpolicies/item/accesspackagecatalog/accesspackageresourceroles/item"
    ie521e61ad9ddcbf9b88a5ce7ed78eb522c057303d2cf11e73eaeaa1190c1381f "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/accesspackageassignmentpolicies/item/accesspackagecatalog/accesspackageresources/item"
    ifb5b38577d6ef99f3a620e02f147001d976e2f601e621227325e8da2c39f3dae "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/accesspackageassignmentpolicies/item/accesspackagecatalog/accesspackages/item"
)

type AccessPackageCatalogRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type AccessPackageCatalogRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *AccessPackageCatalogRequestBuilder) AccessPackageResourceRoles()(*if690c3853d6776a6f085a19f85d182249f323cebba342ab176f22fd85a7b8f83.AccessPackageResourceRolesRequestBuilder) {
    return if690c3853d6776a6f085a19f85d182249f323cebba342ab176f22fd85a7b8f83.NewAccessPackageResourceRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageCatalogRequestBuilder) AccessPackageResourceRolesById(id string)(*icc355b947ff31982a2b7459e8a793ab47ac4d684947820c8068459683087d33f.AccessPackageResourceRoleRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceRole_id"] = id
    }
    return icc355b947ff31982a2b7459e8a793ab47ac4d684947820c8068459683087d33f.NewAccessPackageResourceRoleRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageCatalogRequestBuilder) AccessPackageResources()(*ibe013da09e7701ba00a47c90a6b10900a06ed1137e6129a7c9226e3a07f79403.AccessPackageResourcesRequestBuilder) {
    return ibe013da09e7701ba00a47c90a6b10900a06ed1137e6129a7c9226e3a07f79403.NewAccessPackageResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageCatalogRequestBuilder) AccessPackageResourcesById(id string)(*ie521e61ad9ddcbf9b88a5ce7ed78eb522c057303d2cf11e73eaeaa1190c1381f.AccessPackageResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResource_id"] = id
    }
    return ie521e61ad9ddcbf9b88a5ce7ed78eb522c057303d2cf11e73eaeaa1190c1381f.NewAccessPackageResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageCatalogRequestBuilder) AccessPackageResourceScopes()(*i76ec8ccbad7aee0d645249a9ca5df55114b4d9942d2ce225006a0ffa47878bb8.AccessPackageResourceScopesRequestBuilder) {
    return i76ec8ccbad7aee0d645249a9ca5df55114b4d9942d2ce225006a0ffa47878bb8.NewAccessPackageResourceScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageCatalogRequestBuilder) AccessPackageResourceScopesById(id string)(*ib16fad2499a8a76bd9087a69c65f1fb2f5c8e4d0a12a18be523e8674a8e6005f.AccessPackageResourceScopeRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceScope_id"] = id
    }
    return ib16fad2499a8a76bd9087a69c65f1fb2f5c8e4d0a12a18be523e8674a8e6005f.NewAccessPackageResourceScopeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageCatalogRequestBuilder) AccessPackages()(*i5d568d21b7fe91619b09f9b4ca34ecbcc735854efc2624376c1d8f5bfa5159d1.AccessPackagesRequestBuilder) {
    return i5d568d21b7fe91619b09f9b4ca34ecbcc735854efc2624376c1d8f5bfa5159d1.NewAccessPackagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageCatalogRequestBuilder) AccessPackagesById(id string)(*ifb5b38577d6ef99f3a620e02f147001d976e2f601e621227325e8da2c39f3dae.AccessPackageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackage_id1"] = id
    }
    return ifb5b38577d6ef99f3a620e02f147001d976e2f601e621227325e8da2c39f3dae.NewAccessPackageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewAccessPackageCatalogRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageCatalogRequestBuilder) {
    m := &AccessPackageCatalogRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/identityGovernance/entitlementManagement/accessPackages/{accessPackage_id}/accessPackageAssignmentPolicies/{accessPackageAssignmentPolicy_id}/accessPackageCatalog{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewAccessPackageCatalogRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageCatalogRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessPackageCatalogRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *AccessPackageCatalogRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *AccessPackageCatalogRequestBuilder) CreateGetRequestInformation(q func (value *AccessPackageCatalogRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(AccessPackageCatalogRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *AccessPackageCatalogRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageCatalog, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *AccessPackageCatalogRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *AccessPackageCatalogRequestBuilder) Get(q func (value *AccessPackageCatalogRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageCatalog, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAccessPackageCatalog() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageCatalog), nil
}
func (m *AccessPackageCatalogRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageCatalog, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
