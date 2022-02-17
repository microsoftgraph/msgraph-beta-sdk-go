package managedtenants

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i0abc6f978704af5e99b7eb130a6c0d2e5c1e092e355b17b5af5cdc81f3564b2f "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/windowsdevicemalwarestates"
    i0b126875f3df8ff7b29e49628de663aa712a3121ba720093bf9eedc172a2f466 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/manageddevicecompliancetrends"
    i0f79005843a77cfe19ad60950bc394afee89e8d9bc4b81f6b21b1a8097fbaf3b "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/conditionalaccesspolicycoverages"
    i1d7cc37e7e8e702a054a8907daa331abe670d4a4d9e3e68367c96a95472d6186 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/cloudpcdevices"
    i23841075703fcb028662b4d36c5564ffa62324dd12cbd9935ed7d56f0dc0ca9f "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenantsdetailedinformation"
    i2b66c9972458318dcd18b57d227c1397d3704eed4ceba1aab769bae8473fef7c "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplates"
    i351272dacb7afa33f5595b8bebab11592453eab40a95aa2158a10f1fa59e9972 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/devicecompliancepolicysettingstatesummaries"
    i39d20a98967701e034567f952771787304d37316148dcd57a17798ecd8e6b196 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementactiontenantdeploymentstatuses"
    i507510374fef5a86fcce6eb7ecdb044e4a7bc928ee54765d0620b2b896a746e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenants"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i65beab9150aa7abd814be391125d812f4d3c3bdfde9d2016db1f22a146166aee "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/windowsprotectionstates"
    i824a81edbead3d58bdec21d61a960b3c5f03f44f0e89355db9cc5414e1d67e10 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenantscustomizedinformation"
    i83b6ac56968781119a7493f0b33b04d11b021d336f7c7003dc04aacb7b9c9bb0 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/auditevents"
    i86ccee4069fa2a8887f6c01c1bd59869700242c19224ee8b7aa8d4ad7bfbb89f "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplatecollections"
    i9aa242f436094e02f161a3df7791adfa9932ceaf20ebf440b6069847018656e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementactions"
    i9c7ec006c2124aad1f6bd0b710b33ddb680a37848847978f567ac657d103fd3a "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplatesteps"
    ia20d7fc1ac1c1bc3fdda49bc096fa4ae6d7e86dfce8d76ba7d6ac1a863264fd1 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/credentialuserregistrationssummaries"
    iab13ec14c24b8dac8fd45f37a3723dad71088c5ec497397272541acbbb27e30e "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenanttags"
    ic9cc0fcc7a3eb22d3fce4f7ca12420fe56e916b98e995471ceacfe67707ebaf2 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/cloudpcsoverview"
    icc9f761cb74e72b6d07de77b6cf9c0fe4d404bf9718587ba757ff7c9a01b8c9c "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/cloudpcconnections"
    icf1f2046c60f5e428ce03f4d34ebbf94d9da108c7a89ad5560774e04115aee81 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenantgroups"
    icfac7b17b942815ef506ee4f3661445d6355fd31ed1673ad9d6ef8ebe8b57487 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementintents"
    id0d999b2f9f0ee37ec597aead92f01550dd34dbc53eee9ed3fe40e49f0d53bf7 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplatestepversions"
    idbc8f4984d2cb262d1942923c269bd190c01b953b808d1fdb9cf4be97280aaa9 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/aggregatedpolicycompliances"
    ie2be96252bd3e7e355c4cc7dbf41dd9251cda2dfdea6235dbd43916501b1bb8e "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/manageddevicecompliances"
    i186e85376ac4af3f7e13d1b03ae3ecca0403b582ff4761f086be79d2ceacdf75 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/windowsdevicemalwarestates/item"
    i1ec6b65aff2fcd462f8f45f48ce513ea76dde61b1ec3f1cf01d65b849b740b8b "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementintents/item"
    i22d7e4fc2a16cf517b6c379250cc88fc88b4ed4261b52362648bc93110b5f230 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenantscustomizedinformation/item"
    i24e026309e2f5ec4fcde15b1735de3d790a689d3408b9e14ae04f9328e84c7bd "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/cloudpcdevices/item"
    i36371b16b14a3a26bf98be71cdf018d4cde82965ba2327d328aeac3639125ea0 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplates/item"
    i39c4093a2570401121e930eab4cc05bf3a3f5879dba69aa644fa464b46e8aa51 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/devicecompliancepolicysettingstatesummaries/item"
    i3d3e0cdf3a9fbbfd02cd0827a97a9c93f4ebb2715ac2ee966e30acaf51299aff "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/manageddevicecompliancetrends/item"
    i44e5d1aa3329234fb97b49ec0893e801904dd6a9768684f18a00bd816da70fd6 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenanttags/item"
    i52c320bc8af30828e6ec0c42efaa3e1776f96666e820434df3a8666acde77551 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/aggregatedpolicycompliances/item"
    i60191c74f9be183cc909ce75260c09c09ec76d902f83e1b4f0c161cf0afa63b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/auditevents/item"
    i7002a802b3e3d4667259a74b2f46e5190bce103c7bacef3e3630f3808a1f4256 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplatesteps/item"
    i79c491f13216abf413ae87d0ff1c465e6a8a4809a3f15a3c0d1e120851ef9a5c "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenantgroups/item"
    i8612235b849d43a7141e3638c7b76b5ba29594d12c5272cfc9e32333a038808a "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplatecollections/item"
    i902764218ad5577611d10586599ccdb34a8aa22b15a8c9043e035e89db4d4c7b "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenants/item"
    ib007e276e70fb157eccc427006d3a6fdf591c84e085e79a6ae7d6a7149a17922 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementactiontenantdeploymentstatuses/item"
    ib77c9f0cca75f19b263afc4ae06ae033ceee00cde50068f4a1ad949ad4e69035 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/windowsprotectionstates/item"
    ibb2dbd75a715fb8c4c67773499667f80d218a5b875d91d6bb702f87d5c5815f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/credentialuserregistrationssummaries/item"
    ibeef58461c8396867818d36875356dc3e16db01443724532b324a94c46ca8d52 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplatestepversions/item"
    ic4d32f9937154ecb1938a14d8397e02282c93535909cb4eb03a94798fefb2623 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/conditionalaccesspolicycoverages/item"
    ic98453c829b49eb4768ddc3245c452c7c68722a915539ad220764ab7c19350c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenantsdetailedinformation/item"
    idf8987c670699b4ba0e9e1d460cbd94442affa23953d4b483a7d372543e0f818 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/manageddevicecompliances/item"
    ief008158d809c9b248f0635950f9a44dacadb584a83e78751c8b87051d125c7a "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementactions/item"
    ief1a51d13b4379383242b889e4364a07fc548b896256031dfb57941100b7e3b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/cloudpcconnections/item"
    if6d60fc8097e10a44124bc34967688bbbcf720e50c2b3034d3ecedc559f02741 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/cloudpcsoverview/item"
)

// ManagedTenantsRequestBuilder builds and executes requests for operations under \tenantRelationships\managedTenants
type ManagedTenantsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ManagedTenantsRequestBuilderDeleteOptions options for Delete
type ManagedTenantsRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagedTenantsRequestBuilderGetOptions options for Get
type ManagedTenantsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ManagedTenantsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagedTenantsRequestBuilderGetQueryParameters the operations available to interact with the multi-tenant management platform.
type ManagedTenantsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ManagedTenantsRequestBuilderPatchOptions options for Patch
type ManagedTenantsRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedTenant;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ManagedTenantsRequestBuilder) AggregatedPolicyCompliances()(*idbc8f4984d2cb262d1942923c269bd190c01b953b808d1fdb9cf4be97280aaa9.AggregatedPolicyCompliancesRequestBuilder) {
    return idbc8f4984d2cb262d1942923c269bd190c01b953b808d1fdb9cf4be97280aaa9.NewAggregatedPolicyCompliancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AggregatedPolicyCompliancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.managedTenants.aggregatedPolicyCompliances.item collection
func (m *ManagedTenantsRequestBuilder) AggregatedPolicyCompliancesById(id string)(*i52c320bc8af30828e6ec0c42efaa3e1776f96666e820434df3a8666acde77551.AggregatedPolicyComplianceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["aggregatedPolicyCompliance_id"] = id
    }
    return i52c320bc8af30828e6ec0c42efaa3e1776f96666e820434df3a8666acde77551.NewAggregatedPolicyComplianceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) AuditEvents()(*i83b6ac56968781119a7493f0b33b04d11b021d336f7c7003dc04aacb7b9c9bb0.AuditEventsRequestBuilder) {
    return i83b6ac56968781119a7493f0b33b04d11b021d336f7c7003dc04aacb7b9c9bb0.NewAuditEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuditEventsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.managedTenants.auditEvents.item collection
func (m *ManagedTenantsRequestBuilder) AuditEventsById(id string)(*i60191c74f9be183cc909ce75260c09c09ec76d902f83e1b4f0c161cf0afa63b9.AuditEventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["auditEvent_id"] = id
    }
    return i60191c74f9be183cc909ce75260c09c09ec76d902f83e1b4f0c161cf0afa63b9.NewAuditEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) CloudPcConnections()(*icc9f761cb74e72b6d07de77b6cf9c0fe4d404bf9718587ba757ff7c9a01b8c9c.CloudPcConnectionsRequestBuilder) {
    return icc9f761cb74e72b6d07de77b6cf9c0fe4d404bf9718587ba757ff7c9a01b8c9c.NewCloudPcConnectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudPcConnectionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.managedTenants.cloudPcConnections.item collection
func (m *ManagedTenantsRequestBuilder) CloudPcConnectionsById(id string)(*ief1a51d13b4379383242b889e4364a07fc548b896256031dfb57941100b7e3b5.CloudPcConnectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcConnection_id"] = id
    }
    return ief1a51d13b4379383242b889e4364a07fc548b896256031dfb57941100b7e3b5.NewCloudPcConnectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) CloudPcDevices()(*i1d7cc37e7e8e702a054a8907daa331abe670d4a4d9e3e68367c96a95472d6186.CloudPcDevicesRequestBuilder) {
    return i1d7cc37e7e8e702a054a8907daa331abe670d4a4d9e3e68367c96a95472d6186.NewCloudPcDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudPcDevicesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.managedTenants.cloudPcDevices.item collection
func (m *ManagedTenantsRequestBuilder) CloudPcDevicesById(id string)(*i24e026309e2f5ec4fcde15b1735de3d790a689d3408b9e14ae04f9328e84c7bd.CloudPcDeviceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcDevice_id"] = id
    }
    return i24e026309e2f5ec4fcde15b1735de3d790a689d3408b9e14ae04f9328e84c7bd.NewCloudPcDeviceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) CloudPcsOverview()(*ic9cc0fcc7a3eb22d3fce4f7ca12420fe56e916b98e995471ceacfe67707ebaf2.CloudPcsOverviewRequestBuilder) {
    return ic9cc0fcc7a3eb22d3fce4f7ca12420fe56e916b98e995471ceacfe67707ebaf2.NewCloudPcsOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudPcsOverviewById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.managedTenants.cloudPcsOverview.item collection
func (m *ManagedTenantsRequestBuilder) CloudPcsOverviewById(id string)(*if6d60fc8097e10a44124bc34967688bbbcf720e50c2b3034d3ecedc559f02741.CloudPcOverviewRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcOverview_tenantId"] = id
    }
    return if6d60fc8097e10a44124bc34967688bbbcf720e50c2b3034d3ecedc559f02741.NewCloudPcOverviewRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) ConditionalAccessPolicyCoverages()(*i0f79005843a77cfe19ad60950bc394afee89e8d9bc4b81f6b21b1a8097fbaf3b.ConditionalAccessPolicyCoveragesRequestBuilder) {
    return i0f79005843a77cfe19ad60950bc394afee89e8d9bc4b81f6b21b1a8097fbaf3b.NewConditionalAccessPolicyCoveragesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConditionalAccessPolicyCoveragesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.managedTenants.conditionalAccessPolicyCoverages.item collection
func (m *ManagedTenantsRequestBuilder) ConditionalAccessPolicyCoveragesById(id string)(*ic4d32f9937154ecb1938a14d8397e02282c93535909cb4eb03a94798fefb2623.ConditionalAccessPolicyCoverageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conditionalAccessPolicyCoverage_id"] = id
    }
    return ic4d32f9937154ecb1938a14d8397e02282c93535909cb4eb03a94798fefb2623.NewConditionalAccessPolicyCoverageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewManagedTenantsRequestBuilderInternal instantiates a new ManagedTenantsRequestBuilder and sets the default values.
func NewManagedTenantsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedTenantsRequestBuilder) {
    m := &ManagedTenantsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedTenantsRequestBuilder instantiates a new ManagedTenantsRequestBuilder and sets the default values.
func NewManagedTenantsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedTenantsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the operations available to interact with the multi-tenant management platform.
func (m *ManagedTenantsRequestBuilder) CreateDeleteRequestInformation(options *ManagedTenantsRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the operations available to interact with the multi-tenant management platform.
func (m *ManagedTenantsRequestBuilder) CreateGetRequestInformation(options *ManagedTenantsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation the operations available to interact with the multi-tenant management platform.
func (m *ManagedTenantsRequestBuilder) CreatePatchRequestInformation(options *ManagedTenantsRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *ManagedTenantsRequestBuilder) CredentialUserRegistrationsSummaries()(*ia20d7fc1ac1c1bc3fdda49bc096fa4ae6d7e86dfce8d76ba7d6ac1a863264fd1.CredentialUserRegistrationsSummariesRequestBuilder) {
    return ia20d7fc1ac1c1bc3fdda49bc096fa4ae6d7e86dfce8d76ba7d6ac1a863264fd1.NewCredentialUserRegistrationsSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CredentialUserRegistrationsSummariesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.managedTenants.credentialUserRegistrationsSummaries.item collection
func (m *ManagedTenantsRequestBuilder) CredentialUserRegistrationsSummariesById(id string)(*ibb2dbd75a715fb8c4c67773499667f80d218a5b875d91d6bb702f87d5c5815f3.CredentialUserRegistrationsSummaryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["credentialUserRegistrationsSummary_id"] = id
    }
    return ibb2dbd75a715fb8c4c67773499667f80d218a5b875d91d6bb702f87d5c5815f3.NewCredentialUserRegistrationsSummaryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete the operations available to interact with the multi-tenant management platform.
func (m *ManagedTenantsRequestBuilder) Delete(options *ManagedTenantsRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ManagedTenantsRequestBuilder) DeviceCompliancePolicySettingStateSummaries()(*i351272dacb7afa33f5595b8bebab11592453eab40a95aa2158a10f1fa59e9972.DeviceCompliancePolicySettingStateSummariesRequestBuilder) {
    return i351272dacb7afa33f5595b8bebab11592453eab40a95aa2158a10f1fa59e9972.NewDeviceCompliancePolicySettingStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePolicySettingStateSummariesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.managedTenants.deviceCompliancePolicySettingStateSummaries.item collection
func (m *ManagedTenantsRequestBuilder) DeviceCompliancePolicySettingStateSummariesById(id string)(*i39c4093a2570401121e930eab4cc05bf3a3f5879dba69aa644fa464b46e8aa51.DeviceCompliancePolicySettingStateSummaryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicySettingStateSummary_id"] = id
    }
    return i39c4093a2570401121e930eab4cc05bf3a3f5879dba69aa644fa464b46e8aa51.NewDeviceCompliancePolicySettingStateSummaryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the operations available to interact with the multi-tenant management platform.
func (m *ManagedTenantsRequestBuilder) Get(options *ManagedTenantsRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedTenant, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewManagedTenant() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedTenant), nil
}
func (m *ManagedTenantsRequestBuilder) ManagedDeviceCompliances()(*ie2be96252bd3e7e355c4cc7dbf41dd9251cda2dfdea6235dbd43916501b1bb8e.ManagedDeviceCompliancesRequestBuilder) {
    return ie2be96252bd3e7e355c4cc7dbf41dd9251cda2dfdea6235dbd43916501b1bb8e.NewManagedDeviceCompliancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceCompliancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.managedTenants.managedDeviceCompliances.item collection
func (m *ManagedTenantsRequestBuilder) ManagedDeviceCompliancesById(id string)(*idf8987c670699b4ba0e9e1d460cbd94442affa23953d4b483a7d372543e0f818.ManagedDeviceComplianceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceCompliance_id"] = id
    }
    return idf8987c670699b4ba0e9e1d460cbd94442affa23953d4b483a7d372543e0f818.NewManagedDeviceComplianceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) ManagedDeviceComplianceTrends()(*i0b126875f3df8ff7b29e49628de663aa712a3121ba720093bf9eedc172a2f466.ManagedDeviceComplianceTrendsRequestBuilder) {
    return i0b126875f3df8ff7b29e49628de663aa712a3121ba720093bf9eedc172a2f466.NewManagedDeviceComplianceTrendsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceComplianceTrendsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.managedTenants.managedDeviceComplianceTrends.item collection
func (m *ManagedTenantsRequestBuilder) ManagedDeviceComplianceTrendsById(id string)(*i3d3e0cdf3a9fbbfd02cd0827a97a9c93f4ebb2715ac2ee966e30acaf51299aff.ManagedDeviceComplianceTrendRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceComplianceTrend_id"] = id
    }
    return i3d3e0cdf3a9fbbfd02cd0827a97a9c93f4ebb2715ac2ee966e30acaf51299aff.NewManagedDeviceComplianceTrendRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) ManagementActions()(*i9aa242f436094e02f161a3df7791adfa9932ceaf20ebf440b6069847018656e0.ManagementActionsRequestBuilder) {
    return i9aa242f436094e02f161a3df7791adfa9932ceaf20ebf440b6069847018656e0.NewManagementActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementActionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.managedTenants.managementActions.item collection
func (m *ManagedTenantsRequestBuilder) ManagementActionsById(id string)(*ief008158d809c9b248f0635950f9a44dacadb584a83e78751c8b87051d125c7a.ManagementActionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementAction_id"] = id
    }
    return ief008158d809c9b248f0635950f9a44dacadb584a83e78751c8b87051d125c7a.NewManagementActionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) ManagementActionTenantDeploymentStatuses()(*i39d20a98967701e034567f952771787304d37316148dcd57a17798ecd8e6b196.ManagementActionTenantDeploymentStatusesRequestBuilder) {
    return i39d20a98967701e034567f952771787304d37316148dcd57a17798ecd8e6b196.NewManagementActionTenantDeploymentStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementActionTenantDeploymentStatusesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.managedTenants.managementActionTenantDeploymentStatuses.item collection
func (m *ManagedTenantsRequestBuilder) ManagementActionTenantDeploymentStatusesById(id string)(*ib007e276e70fb157eccc427006d3a6fdf591c84e085e79a6ae7d6a7149a17922.ManagementActionTenantDeploymentStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementActionTenantDeploymentStatus_id"] = id
    }
    return ib007e276e70fb157eccc427006d3a6fdf591c84e085e79a6ae7d6a7149a17922.NewManagementActionTenantDeploymentStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) ManagementIntents()(*icfac7b17b942815ef506ee4f3661445d6355fd31ed1673ad9d6ef8ebe8b57487.ManagementIntentsRequestBuilder) {
    return icfac7b17b942815ef506ee4f3661445d6355fd31ed1673ad9d6ef8ebe8b57487.NewManagementIntentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementIntentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.managedTenants.managementIntents.item collection
func (m *ManagedTenantsRequestBuilder) ManagementIntentsById(id string)(*i1ec6b65aff2fcd462f8f45f48ce513ea76dde61b1ec3f1cf01d65b849b740b8b.ManagementIntentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementIntent_id"] = id
    }
    return i1ec6b65aff2fcd462f8f45f48ce513ea76dde61b1ec3f1cf01d65b849b740b8b.NewManagementIntentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) ManagementTemplateCollections()(*i86ccee4069fa2a8887f6c01c1bd59869700242c19224ee8b7aa8d4ad7bfbb89f.ManagementTemplateCollectionsRequestBuilder) {
    return i86ccee4069fa2a8887f6c01c1bd59869700242c19224ee8b7aa8d4ad7bfbb89f.NewManagementTemplateCollectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementTemplateCollectionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.managedTenants.managementTemplateCollections.item collection
func (m *ManagedTenantsRequestBuilder) ManagementTemplateCollectionsById(id string)(*i8612235b849d43a7141e3638c7b76b5ba29594d12c5272cfc9e32333a038808a.ManagementTemplateCollectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementTemplateCollection_id"] = id
    }
    return i8612235b849d43a7141e3638c7b76b5ba29594d12c5272cfc9e32333a038808a.NewManagementTemplateCollectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) ManagementTemplates()(*i2b66c9972458318dcd18b57d227c1397d3704eed4ceba1aab769bae8473fef7c.ManagementTemplatesRequestBuilder) {
    return i2b66c9972458318dcd18b57d227c1397d3704eed4ceba1aab769bae8473fef7c.NewManagementTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementTemplatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.managedTenants.managementTemplates.item collection
func (m *ManagedTenantsRequestBuilder) ManagementTemplatesById(id string)(*i36371b16b14a3a26bf98be71cdf018d4cde82965ba2327d328aeac3639125ea0.ManagementTemplateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementTemplate_id"] = id
    }
    return i36371b16b14a3a26bf98be71cdf018d4cde82965ba2327d328aeac3639125ea0.NewManagementTemplateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) ManagementTemplateSteps()(*i9c7ec006c2124aad1f6bd0b710b33ddb680a37848847978f567ac657d103fd3a.ManagementTemplateStepsRequestBuilder) {
    return i9c7ec006c2124aad1f6bd0b710b33ddb680a37848847978f567ac657d103fd3a.NewManagementTemplateStepsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementTemplateStepsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.managedTenants.managementTemplateSteps.item collection
func (m *ManagedTenantsRequestBuilder) ManagementTemplateStepsById(id string)(*i7002a802b3e3d4667259a74b2f46e5190bce103c7bacef3e3630f3808a1f4256.ManagementTemplateStepRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementTemplateStep_id"] = id
    }
    return i7002a802b3e3d4667259a74b2f46e5190bce103c7bacef3e3630f3808a1f4256.NewManagementTemplateStepRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) ManagementTemplateStepVersions()(*id0d999b2f9f0ee37ec597aead92f01550dd34dbc53eee9ed3fe40e49f0d53bf7.ManagementTemplateStepVersionsRequestBuilder) {
    return id0d999b2f9f0ee37ec597aead92f01550dd34dbc53eee9ed3fe40e49f0d53bf7.NewManagementTemplateStepVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementTemplateStepVersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.managedTenants.managementTemplateStepVersions.item collection
func (m *ManagedTenantsRequestBuilder) ManagementTemplateStepVersionsById(id string)(*ibeef58461c8396867818d36875356dc3e16db01443724532b324a94c46ca8d52.ManagementTemplateStepVersionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementTemplateStepVersion_id"] = id
    }
    return ibeef58461c8396867818d36875356dc3e16db01443724532b324a94c46ca8d52.NewManagementTemplateStepVersionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the operations available to interact with the multi-tenant management platform.
func (m *ManagedTenantsRequestBuilder) Patch(options *ManagedTenantsRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ManagedTenantsRequestBuilder) TenantGroups()(*icf1f2046c60f5e428ce03f4d34ebbf94d9da108c7a89ad5560774e04115aee81.TenantGroupsRequestBuilder) {
    return icf1f2046c60f5e428ce03f4d34ebbf94d9da108c7a89ad5560774e04115aee81.NewTenantGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TenantGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.managedTenants.tenantGroups.item collection
func (m *ManagedTenantsRequestBuilder) TenantGroupsById(id string)(*i79c491f13216abf413ae87d0ff1c465e6a8a4809a3f15a3c0d1e120851ef9a5c.TenantGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tenantGroup_id"] = id
    }
    return i79c491f13216abf413ae87d0ff1c465e6a8a4809a3f15a3c0d1e120851ef9a5c.NewTenantGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) Tenants()(*i507510374fef5a86fcce6eb7ecdb044e4a7bc928ee54765d0620b2b896a746e1.TenantsRequestBuilder) {
    return i507510374fef5a86fcce6eb7ecdb044e4a7bc928ee54765d0620b2b896a746e1.NewTenantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TenantsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.managedTenants.tenants.item collection
func (m *ManagedTenantsRequestBuilder) TenantsById(id string)(*i902764218ad5577611d10586599ccdb34a8aa22b15a8c9043e035e89db4d4c7b.TenantRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tenant_id"] = id
    }
    return i902764218ad5577611d10586599ccdb34a8aa22b15a8c9043e035e89db4d4c7b.NewTenantRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) TenantsCustomizedInformation()(*i824a81edbead3d58bdec21d61a960b3c5f03f44f0e89355db9cc5414e1d67e10.TenantsCustomizedInformationRequestBuilder) {
    return i824a81edbead3d58bdec21d61a960b3c5f03f44f0e89355db9cc5414e1d67e10.NewTenantsCustomizedInformationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TenantsCustomizedInformationById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.managedTenants.tenantsCustomizedInformation.item collection
func (m *ManagedTenantsRequestBuilder) TenantsCustomizedInformationById(id string)(*i22d7e4fc2a16cf517b6c379250cc88fc88b4ed4261b52362648bc93110b5f230.TenantCustomizedInformationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tenantCustomizedInformation_id"] = id
    }
    return i22d7e4fc2a16cf517b6c379250cc88fc88b4ed4261b52362648bc93110b5f230.NewTenantCustomizedInformationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) TenantsDetailedInformation()(*i23841075703fcb028662b4d36c5564ffa62324dd12cbd9935ed7d56f0dc0ca9f.TenantsDetailedInformationRequestBuilder) {
    return i23841075703fcb028662b4d36c5564ffa62324dd12cbd9935ed7d56f0dc0ca9f.NewTenantsDetailedInformationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TenantsDetailedInformationById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.managedTenants.tenantsDetailedInformation.item collection
func (m *ManagedTenantsRequestBuilder) TenantsDetailedInformationById(id string)(*ic98453c829b49eb4768ddc3245c452c7c68722a915539ad220764ab7c19350c3.TenantDetailedInformationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tenantDetailedInformation_id"] = id
    }
    return ic98453c829b49eb4768ddc3245c452c7c68722a915539ad220764ab7c19350c3.NewTenantDetailedInformationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) TenantTags()(*iab13ec14c24b8dac8fd45f37a3723dad71088c5ec497397272541acbbb27e30e.TenantTagsRequestBuilder) {
    return iab13ec14c24b8dac8fd45f37a3723dad71088c5ec497397272541acbbb27e30e.NewTenantTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TenantTagsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.managedTenants.tenantTags.item collection
func (m *ManagedTenantsRequestBuilder) TenantTagsById(id string)(*i44e5d1aa3329234fb97b49ec0893e801904dd6a9768684f18a00bd816da70fd6.TenantTagRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tenantTag_id"] = id
    }
    return i44e5d1aa3329234fb97b49ec0893e801904dd6a9768684f18a00bd816da70fd6.NewTenantTagRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) WindowsDeviceMalwareStates()(*i0abc6f978704af5e99b7eb130a6c0d2e5c1e092e355b17b5af5cdc81f3564b2f.WindowsDeviceMalwareStatesRequestBuilder) {
    return i0abc6f978704af5e99b7eb130a6c0d2e5c1e092e355b17b5af5cdc81f3564b2f.NewWindowsDeviceMalwareStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsDeviceMalwareStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.managedTenants.windowsDeviceMalwareStates.item collection
func (m *ManagedTenantsRequestBuilder) WindowsDeviceMalwareStatesById(id string)(*i186e85376ac4af3f7e13d1b03ae3ecca0403b582ff4761f086be79d2ceacdf75.WindowsDeviceMalwareStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsDeviceMalwareState_id"] = id
    }
    return i186e85376ac4af3f7e13d1b03ae3ecca0403b582ff4761f086be79d2ceacdf75.NewWindowsDeviceMalwareStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedTenantsRequestBuilder) WindowsProtectionStates()(*i65beab9150aa7abd814be391125d812f4d3c3bdfde9d2016db1f22a146166aee.WindowsProtectionStatesRequestBuilder) {
    return i65beab9150aa7abd814be391125d812f4d3c3bdfde9d2016db1f22a146166aee.NewWindowsProtectionStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsProtectionStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.managedTenants.windowsProtectionStates.item collection
func (m *ManagedTenantsRequestBuilder) WindowsProtectionStatesById(id string)(*ib77c9f0cca75f19b263afc4ae06ae033ceee00cde50068f4a1ad949ad4e69035.WindowsProtectionStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsProtectionState_id"] = id
    }
    return ib77c9f0cca75f19b263afc4ae06ae033ceee00cde50068f4a1ad949ad4e69035.NewWindowsProtectionStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
