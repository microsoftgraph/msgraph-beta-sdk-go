package managedtenants

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
    i05a7b0ceeebbeaca3b0e58708628e4f815ac617a23ec0c52f6ad99de33c98651 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managedtenantapinotifications"
    i0abc6f978704af5e99b7eb130a6c0d2e5c1e092e355b17b5af5cdc81f3564b2f "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/windowsdevicemalwarestates"
    i0b126875f3df8ff7b29e49628de663aa712a3121ba720093bf9eedc172a2f466 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/manageddevicecompliancetrends"
    i0f79005843a77cfe19ad60950bc394afee89e8d9bc4b81f6b21b1a8097fbaf3b "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/conditionalaccesspolicycoverages"
    i1d7cc37e7e8e702a054a8907daa331abe670d4a4d9e3e68367c96a95472d6186 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/cloudpcdevices"
    i1e0d9f981a2f7ec6977de1f608176ed15d12b8f4c41d22354dc9c1610b826a35 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplatecollectiontenantsummaries"
    i23841075703fcb028662b4d36c5564ffa62324dd12cbd9935ed7d56f0dc0ca9f "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenantsdetailedinformation"
    i2b66c9972458318dcd18b57d227c1397d3704eed4ceba1aab769bae8473fef7c "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplates"
    i2cdcdc509df3082b25195c0a43c6d5f07ffa1bd0852bda0e8d272dc9a577f670 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managedtenantalertrules"
    i351272dacb7afa33f5595b8bebab11592453eab40a95aa2158a10f1fa59e9972 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/devicecompliancepolicysettingstatesummaries"
    i39d20a98967701e034567f952771787304d37316148dcd57a17798ecd8e6b196 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementactiontenantdeploymentstatuses"
    i412f9c2c45d441afb79881179a35f22a85e609a2de25dab29ff47d663d739c03 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managedtenantticketingendpoints"
    i507510374fef5a86fcce6eb7ecdb044e4a7bc928ee54765d0620b2b896a746e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenants"
    i65beab9150aa7abd814be391125d812f4d3c3bdfde9d2016db1f22a146166aee "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/windowsprotectionstates"
    i6c8ccc9d63db57b24eef1a198f4256d2b75c2f3d924dfb09f7dc3934cfc5dc56 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplatesteptenantsummaries"
    i71e724937f2266c62830468204cfd18d9303a53fb8e960124d76a13f5477fb03 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/myroles"
    i824a81edbead3d58bdec21d61a960b3c5f03f44f0e89355db9cc5414e1d67e10 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenantscustomizedinformation"
    i83b6ac56968781119a7493f0b33b04d11b021d336f7c7003dc04aacb7b9c9bb0 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/auditevents"
    i86ccee4069fa2a8887f6c01c1bd59869700242c19224ee8b7aa8d4ad7bfbb89f "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplatecollections"
    i8c93d5be69544696015ab3c63cfaa50d7f46a5bd17a5a56e709bfe4fc098793e "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managedtenantalertruledefinitions"
    i8dc3739f8028fe4b6369c5c2685536150d09cc78772593ddd13bc3c414f4bd29 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managedtenantalerts"
    i9aa242f436094e02f161a3df7791adfa9932ceaf20ebf440b6069847018656e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementactions"
    i9c7ec006c2124aad1f6bd0b710b33ddb680a37848847978f567ac657d103fd3a "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplatesteps"
    i9db702ace2de141bc95e2d3b64dd51c37fc5476b1b1061278955b8b7a285f99d "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managedtenantalertlogs"
    ia20d7fc1ac1c1bc3fdda49bc096fa4ae6d7e86dfce8d76ba7d6ac1a863264fd1 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/credentialuserregistrationssummaries"
    iab13ec14c24b8dac8fd45f37a3723dad71088c5ec497397272541acbbb27e30e "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenanttags"
    ic9cc0fcc7a3eb22d3fce4f7ca12420fe56e916b98e995471ceacfe67707ebaf2 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/cloudpcsoverview"
    icc9f761cb74e72b6d07de77b6cf9c0fe4d404bf9718587ba757ff7c9a01b8c9c "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/cloudpcconnections"
    icf1f2046c60f5e428ce03f4d34ebbf94d9da108c7a89ad5560774e04115aee81 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenantgroups"
    icfac7b17b942815ef506ee4f3661445d6355fd31ed1673ad9d6ef8ebe8b57487 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementintents"
    id0d999b2f9f0ee37ec597aead92f01550dd34dbc53eee9ed3fe40e49f0d53bf7 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplatestepversions"
    id38271415c05281407c0442f61c0fa978c6d3d451589b71dc5cf915930b23e2a "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managedtenantemailnotifications"
    idbc8f4984d2cb262d1942923c269bd190c01b953b808d1fdb9cf4be97280aaa9 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/aggregatedpolicycompliances"
    ie2be96252bd3e7e355c4cc7dbf41dd9251cda2dfdea6235dbd43916501b1bb8e "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/manageddevicecompliances"
    i11da6ef4eeeb8786fb93d5972f290e02c6b190e32733d30dff04b7a103677aff "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplatesteptenantsummaries/item"
    i186e85376ac4af3f7e13d1b03ae3ecca0403b582ff4761f086be79d2ceacdf75 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/windowsdevicemalwarestates/item"
    i1ec6b65aff2fcd462f8f45f48ce513ea76dde61b1ec3f1cf01d65b849b740b8b "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementintents/item"
    i22d7e4fc2a16cf517b6c379250cc88fc88b4ed4261b52362648bc93110b5f230 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenantscustomizedinformation/item"
    i24e026309e2f5ec4fcde15b1735de3d790a689d3408b9e14ae04f9328e84c7bd "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/cloudpcdevices/item"
    i36371b16b14a3a26bf98be71cdf018d4cde82965ba2327d328aeac3639125ea0 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplates/item"
    i39c4093a2570401121e930eab4cc05bf3a3f5879dba69aa644fa464b46e8aa51 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/devicecompliancepolicysettingstatesummaries/item"
    i3d3e0cdf3a9fbbfd02cd0827a97a9c93f4ebb2715ac2ee966e30acaf51299aff "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/manageddevicecompliancetrends/item"
    i44e5d1aa3329234fb97b49ec0893e801904dd6a9768684f18a00bd816da70fd6 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenanttags/item"
    i5002f242d66f13834e9f5ef44d7dbf8f7992424177c38b829295838867dae36f "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managedtenantapinotifications/item"
    i52c320bc8af30828e6ec0c42efaa3e1776f96666e820434df3a8666acde77551 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/aggregatedpolicycompliances/item"
    i60191c74f9be183cc909ce75260c09c09ec76d902f83e1b4f0c161cf0afa63b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/auditevents/item"
    i66506e2a89f3ae47e0d61d5a939ab1798a7d4fb9012b322b2c509d0c7ee44761 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managedtenantticketingendpoints/item"
    i7002a802b3e3d4667259a74b2f46e5190bce103c7bacef3e3630f3808a1f4256 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplatesteps/item"
    i7207a349b16af43bf0ab0f93a50be8bca6caf9e62ce2c90215ae9ed4137863b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managedtenantalerts/item"
    i79c491f13216abf413ae87d0ff1c465e6a8a4809a3f15a3c0d1e120851ef9a5c "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenantgroups/item"
    i8612235b849d43a7141e3638c7b76b5ba29594d12c5272cfc9e32333a038808a "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplatecollections/item"
    i902764218ad5577611d10586599ccdb34a8aa22b15a8c9043e035e89db4d4c7b "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenants/item"
    ib007e276e70fb157eccc427006d3a6fdf591c84e085e79a6ae7d6a7149a17922 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementactiontenantdeploymentstatuses/item"
    ib38449635f2df87835a316e9d35bdd8ac3d8641ecd90714b1eacf60a43539b76 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplatecollectiontenantsummaries/item"
    ib38de1bc2342da950cccc8e38397f843bd594e1cd1797be00f0b3b3cb7ac6f55 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managedtenantalertlogs/item"
    ib3de4c41c301db1496cc9d3bc09f4fa6822595a1f06c6de4ddc82599b19d147a "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managedtenantalertruledefinitions/item"
    ib6a72645824ea588f45be0386a20b709a95dddebe610e26a32be605ca2493073 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managedtenantalertrules/item"
    ib77c9f0cca75f19b263afc4ae06ae033ceee00cde50068f4a1ad949ad4e69035 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/windowsprotectionstates/item"
    ibb2dbd75a715fb8c4c67773499667f80d218a5b875d91d6bb702f87d5c5815f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/credentialuserregistrationssummaries/item"
    ibeef58461c8396867818d36875356dc3e16db01443724532b324a94c46ca8d52 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplatestepversions/item"
    ic4d32f9937154ecb1938a14d8397e02282c93535909cb4eb03a94798fefb2623 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/conditionalaccesspolicycoverages/item"
    ic7759d4bd32132833dc490b8b83ddc960dd8a9a6c0dcbf66abde2430a861a585 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managedtenantemailnotifications/item"
    ic98453c829b49eb4768ddc3245c452c7c68722a915539ad220764ab7c19350c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/tenantsdetailedinformation/item"
    id8a75758c62dfd2efe128ae3797671432b468fe56c4706706dfeb1c3655c9799 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/myroles/item"
    idf8987c670699b4ba0e9e1d460cbd94442affa23953d4b483a7d372543e0f818 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/manageddevicecompliances/item"
    ief008158d809c9b248f0635950f9a44dacadb584a83e78751c8b87051d125c7a "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementactions/item"
    ief1a51d13b4379383242b889e4364a07fc548b896256031dfb57941100b7e3b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/cloudpcconnections/item"
    if6d60fc8097e10a44124bc34967688bbbcf720e50c2b3034d3ecedc559f02741 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/cloudpcsoverview/item"
)

// ManagedTenantsRequestBuilder provides operations to manage the managedTenants property of the microsoft.graph.tenantRelationship entity.
type ManagedTenantsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedTenantsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedTenantsRequestBuilderGetQueryParameters the operations available to interact with the multi-tenant management platform.
type ManagedTenantsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedTenantsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedTenantsRequestBuilderGetQueryParameters
}
// ManagedTenantsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AggregatedPolicyCompliances provides operations to manage the aggregatedPolicyCompliances property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) AggregatedPolicyCompliances()(*idbc8f4984d2cb262d1942923c269bd190c01b953b808d1fdb9cf4be97280aaa9.AggregatedPolicyCompliancesRequestBuilder) {
    return idbc8f4984d2cb262d1942923c269bd190c01b953b808d1fdb9cf4be97280aaa9.NewAggregatedPolicyCompliancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AggregatedPolicyCompliancesById provides operations to manage the aggregatedPolicyCompliances property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) AggregatedPolicyCompliancesById(id string)(*i52c320bc8af30828e6ec0c42efaa3e1776f96666e820434df3a8666acde77551.AggregatedPolicyComplianceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["aggregatedPolicyCompliance%2Did"] = id
    }
    return i52c320bc8af30828e6ec0c42efaa3e1776f96666e820434df3a8666acde77551.NewAggregatedPolicyComplianceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AuditEvents provides operations to manage the auditEvents property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) AuditEvents()(*i83b6ac56968781119a7493f0b33b04d11b021d336f7c7003dc04aacb7b9c9bb0.AuditEventsRequestBuilder) {
    return i83b6ac56968781119a7493f0b33b04d11b021d336f7c7003dc04aacb7b9c9bb0.NewAuditEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuditEventsById provides operations to manage the auditEvents property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) AuditEventsById(id string)(*i60191c74f9be183cc909ce75260c09c09ec76d902f83e1b4f0c161cf0afa63b9.AuditEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["auditEvent%2Did"] = id
    }
    return i60191c74f9be183cc909ce75260c09c09ec76d902f83e1b4f0c161cf0afa63b9.NewAuditEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CloudPcConnections provides operations to manage the cloudPcConnections property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) CloudPcConnections()(*icc9f761cb74e72b6d07de77b6cf9c0fe4d404bf9718587ba757ff7c9a01b8c9c.CloudPcConnectionsRequestBuilder) {
    return icc9f761cb74e72b6d07de77b6cf9c0fe4d404bf9718587ba757ff7c9a01b8c9c.NewCloudPcConnectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudPcConnectionsById provides operations to manage the cloudPcConnections property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) CloudPcConnectionsById(id string)(*ief1a51d13b4379383242b889e4364a07fc548b896256031dfb57941100b7e3b5.CloudPcConnectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcConnection%2Did"] = id
    }
    return ief1a51d13b4379383242b889e4364a07fc548b896256031dfb57941100b7e3b5.NewCloudPcConnectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CloudPcDevices provides operations to manage the cloudPcDevices property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) CloudPcDevices()(*i1d7cc37e7e8e702a054a8907daa331abe670d4a4d9e3e68367c96a95472d6186.CloudPcDevicesRequestBuilder) {
    return i1d7cc37e7e8e702a054a8907daa331abe670d4a4d9e3e68367c96a95472d6186.NewCloudPcDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudPcDevicesById provides operations to manage the cloudPcDevices property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) CloudPcDevicesById(id string)(*i24e026309e2f5ec4fcde15b1735de3d790a689d3408b9e14ae04f9328e84c7bd.CloudPcDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcDevice%2Did"] = id
    }
    return i24e026309e2f5ec4fcde15b1735de3d790a689d3408b9e14ae04f9328e84c7bd.NewCloudPcDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CloudPcsOverview provides operations to manage the cloudPcsOverview property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) CloudPcsOverview()(*ic9cc0fcc7a3eb22d3fce4f7ca12420fe56e916b98e995471ceacfe67707ebaf2.CloudPcsOverviewRequestBuilder) {
    return ic9cc0fcc7a3eb22d3fce4f7ca12420fe56e916b98e995471ceacfe67707ebaf2.NewCloudPcsOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudPcsOverviewById provides operations to manage the cloudPcsOverview property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) CloudPcsOverviewById(id string)(*if6d60fc8097e10a44124bc34967688bbbcf720e50c2b3034d3ecedc559f02741.CloudPcOverviewTenantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcOverview%2DtenantId"] = id
    }
    return if6d60fc8097e10a44124bc34967688bbbcf720e50c2b3034d3ecedc559f02741.NewCloudPcOverviewTenantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ConditionalAccessPolicyCoverages provides operations to manage the conditionalAccessPolicyCoverages property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ConditionalAccessPolicyCoverages()(*i0f79005843a77cfe19ad60950bc394afee89e8d9bc4b81f6b21b1a8097fbaf3b.ConditionalAccessPolicyCoveragesRequestBuilder) {
    return i0f79005843a77cfe19ad60950bc394afee89e8d9bc4b81f6b21b1a8097fbaf3b.NewConditionalAccessPolicyCoveragesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConditionalAccessPolicyCoveragesById provides operations to manage the conditionalAccessPolicyCoverages property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ConditionalAccessPolicyCoveragesById(id string)(*ic4d32f9937154ecb1938a14d8397e02282c93535909cb4eb03a94798fefb2623.ConditionalAccessPolicyCoverageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conditionalAccessPolicyCoverage%2Did"] = id
    }
    return ic4d32f9937154ecb1938a14d8397e02282c93535909cb4eb03a94798fefb2623.NewConditionalAccessPolicyCoverageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewManagedTenantsRequestBuilderInternal instantiates a new ManagedTenantsRequestBuilder and sets the default values.
func NewManagedTenantsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsRequestBuilder) {
    m := &ManagedTenantsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedTenantsRequestBuilder instantiates a new ManagedTenantsRequestBuilder and sets the default values.
func NewManagedTenantsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property managedTenants for tenantRelationships
func (m *ManagedTenantsRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the operations available to interact with the multi-tenant management platform.
func (m *ManagedTenantsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property managedTenants in tenantRelationships
func (m *ManagedTenantsRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantable, requestConfiguration *ManagedTenantsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CredentialUserRegistrationsSummaries provides operations to manage the credentialUserRegistrationsSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) CredentialUserRegistrationsSummaries()(*ia20d7fc1ac1c1bc3fdda49bc096fa4ae6d7e86dfce8d76ba7d6ac1a863264fd1.CredentialUserRegistrationsSummariesRequestBuilder) {
    return ia20d7fc1ac1c1bc3fdda49bc096fa4ae6d7e86dfce8d76ba7d6ac1a863264fd1.NewCredentialUserRegistrationsSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CredentialUserRegistrationsSummariesById provides operations to manage the credentialUserRegistrationsSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) CredentialUserRegistrationsSummariesById(id string)(*ibb2dbd75a715fb8c4c67773499667f80d218a5b875d91d6bb702f87d5c5815f3.CredentialUserRegistrationsSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["credentialUserRegistrationsSummary%2Did"] = id
    }
    return ibb2dbd75a715fb8c4c67773499667f80d218a5b875d91d6bb702f87d5c5815f3.NewCredentialUserRegistrationsSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property managedTenants for tenantRelationships
func (m *ManagedTenantsRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedTenantsRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceCompliancePolicySettingStateSummaries provides operations to manage the deviceCompliancePolicySettingStateSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) DeviceCompliancePolicySettingStateSummaries()(*i351272dacb7afa33f5595b8bebab11592453eab40a95aa2158a10f1fa59e9972.DeviceCompliancePolicySettingStateSummariesRequestBuilder) {
    return i351272dacb7afa33f5595b8bebab11592453eab40a95aa2158a10f1fa59e9972.NewDeviceCompliancePolicySettingStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePolicySettingStateSummariesById provides operations to manage the deviceCompliancePolicySettingStateSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) DeviceCompliancePolicySettingStateSummariesById(id string)(*i39c4093a2570401121e930eab4cc05bf3a3f5879dba69aa644fa464b46e8aa51.DeviceCompliancePolicySettingStateSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicySettingStateSummary%2Did"] = id
    }
    return i39c4093a2570401121e930eab4cc05bf3a3f5879dba69aa644fa464b46e8aa51.NewDeviceCompliancePolicySettingStateSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the operations available to interact with the multi-tenant management platform.
func (m *ManagedTenantsRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedTenantsRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagedTenantFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantable), nil
}
// ManagedDeviceCompliances provides operations to manage the managedDeviceCompliances property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedDeviceCompliances()(*ie2be96252bd3e7e355c4cc7dbf41dd9251cda2dfdea6235dbd43916501b1bb8e.ManagedDeviceCompliancesRequestBuilder) {
    return ie2be96252bd3e7e355c4cc7dbf41dd9251cda2dfdea6235dbd43916501b1bb8e.NewManagedDeviceCompliancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceCompliancesById provides operations to manage the managedDeviceCompliances property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedDeviceCompliancesById(id string)(*idf8987c670699b4ba0e9e1d460cbd94442affa23953d4b483a7d372543e0f818.ManagedDeviceComplianceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceCompliance%2Did"] = id
    }
    return idf8987c670699b4ba0e9e1d460cbd94442affa23953d4b483a7d372543e0f818.NewManagedDeviceComplianceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedDeviceComplianceTrends provides operations to manage the managedDeviceComplianceTrends property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedDeviceComplianceTrends()(*i0b126875f3df8ff7b29e49628de663aa712a3121ba720093bf9eedc172a2f466.ManagedDeviceComplianceTrendsRequestBuilder) {
    return i0b126875f3df8ff7b29e49628de663aa712a3121ba720093bf9eedc172a2f466.NewManagedDeviceComplianceTrendsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceComplianceTrendsById provides operations to manage the managedDeviceComplianceTrends property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedDeviceComplianceTrendsById(id string)(*i3d3e0cdf3a9fbbfd02cd0827a97a9c93f4ebb2715ac2ee966e30acaf51299aff.ManagedDeviceComplianceTrendItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceComplianceTrend%2Did"] = id
    }
    return i3d3e0cdf3a9fbbfd02cd0827a97a9c93f4ebb2715ac2ee966e30acaf51299aff.NewManagedDeviceComplianceTrendItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedTenantAlertLogs provides operations to manage the managedTenantAlertLogs property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantAlertLogs()(*i9db702ace2de141bc95e2d3b64dd51c37fc5476b1b1061278955b8b7a285f99d.ManagedTenantAlertLogsRequestBuilder) {
    return i9db702ace2de141bc95e2d3b64dd51c37fc5476b1b1061278955b8b7a285f99d.NewManagedTenantAlertLogsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedTenantAlertLogsById provides operations to manage the managedTenantAlertLogs property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantAlertLogsById(id string)(*ib38de1bc2342da950cccc8e38397f843bd594e1cd1797be00f0b3b3cb7ac6f55.ManagedTenantAlertLogItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantAlertLog%2Did"] = id
    }
    return ib38de1bc2342da950cccc8e38397f843bd594e1cd1797be00f0b3b3cb7ac6f55.NewManagedTenantAlertLogItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedTenantAlertRuleDefinitions provides operations to manage the managedTenantAlertRuleDefinitions property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantAlertRuleDefinitions()(*i8c93d5be69544696015ab3c63cfaa50d7f46a5bd17a5a56e709bfe4fc098793e.ManagedTenantAlertRuleDefinitionsRequestBuilder) {
    return i8c93d5be69544696015ab3c63cfaa50d7f46a5bd17a5a56e709bfe4fc098793e.NewManagedTenantAlertRuleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedTenantAlertRuleDefinitionsById provides operations to manage the managedTenantAlertRuleDefinitions property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantAlertRuleDefinitionsById(id string)(*ib3de4c41c301db1496cc9d3bc09f4fa6822595a1f06c6de4ddc82599b19d147a.ManagedTenantAlertRuleDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantAlertRuleDefinition%2Did"] = id
    }
    return ib3de4c41c301db1496cc9d3bc09f4fa6822595a1f06c6de4ddc82599b19d147a.NewManagedTenantAlertRuleDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedTenantAlertRules provides operations to manage the managedTenantAlertRules property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantAlertRules()(*i2cdcdc509df3082b25195c0a43c6d5f07ffa1bd0852bda0e8d272dc9a577f670.ManagedTenantAlertRulesRequestBuilder) {
    return i2cdcdc509df3082b25195c0a43c6d5f07ffa1bd0852bda0e8d272dc9a577f670.NewManagedTenantAlertRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedTenantAlertRulesById provides operations to manage the managedTenantAlertRules property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantAlertRulesById(id string)(*ib6a72645824ea588f45be0386a20b709a95dddebe610e26a32be605ca2493073.ManagedTenantAlertRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantAlertRule%2Did"] = id
    }
    return ib6a72645824ea588f45be0386a20b709a95dddebe610e26a32be605ca2493073.NewManagedTenantAlertRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedTenantAlerts provides operations to manage the managedTenantAlerts property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantAlerts()(*i8dc3739f8028fe4b6369c5c2685536150d09cc78772593ddd13bc3c414f4bd29.ManagedTenantAlertsRequestBuilder) {
    return i8dc3739f8028fe4b6369c5c2685536150d09cc78772593ddd13bc3c414f4bd29.NewManagedTenantAlertsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedTenantAlertsById provides operations to manage the managedTenantAlerts property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantAlertsById(id string)(*i7207a349b16af43bf0ab0f93a50be8bca6caf9e62ce2c90215ae9ed4137863b8.ManagedTenantAlertItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantAlert%2Did"] = id
    }
    return i7207a349b16af43bf0ab0f93a50be8bca6caf9e62ce2c90215ae9ed4137863b8.NewManagedTenantAlertItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedTenantApiNotifications provides operations to manage the managedTenantApiNotifications property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantApiNotifications()(*i05a7b0ceeebbeaca3b0e58708628e4f815ac617a23ec0c52f6ad99de33c98651.ManagedTenantApiNotificationsRequestBuilder) {
    return i05a7b0ceeebbeaca3b0e58708628e4f815ac617a23ec0c52f6ad99de33c98651.NewManagedTenantApiNotificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedTenantApiNotificationsById provides operations to manage the managedTenantApiNotifications property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantApiNotificationsById(id string)(*i5002f242d66f13834e9f5ef44d7dbf8f7992424177c38b829295838867dae36f.ManagedTenantApiNotificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantApiNotification%2Did"] = id
    }
    return i5002f242d66f13834e9f5ef44d7dbf8f7992424177c38b829295838867dae36f.NewManagedTenantApiNotificationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedTenantEmailNotifications provides operations to manage the managedTenantEmailNotifications property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantEmailNotifications()(*id38271415c05281407c0442f61c0fa978c6d3d451589b71dc5cf915930b23e2a.ManagedTenantEmailNotificationsRequestBuilder) {
    return id38271415c05281407c0442f61c0fa978c6d3d451589b71dc5cf915930b23e2a.NewManagedTenantEmailNotificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedTenantEmailNotificationsById provides operations to manage the managedTenantEmailNotifications property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantEmailNotificationsById(id string)(*ic7759d4bd32132833dc490b8b83ddc960dd8a9a6c0dcbf66abde2430a861a585.ManagedTenantEmailNotificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantEmailNotification%2Did"] = id
    }
    return ic7759d4bd32132833dc490b8b83ddc960dd8a9a6c0dcbf66abde2430a861a585.NewManagedTenantEmailNotificationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedTenantTicketingEndpoints provides operations to manage the managedTenantTicketingEndpoints property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantTicketingEndpoints()(*i412f9c2c45d441afb79881179a35f22a85e609a2de25dab29ff47d663d739c03.ManagedTenantTicketingEndpointsRequestBuilder) {
    return i412f9c2c45d441afb79881179a35f22a85e609a2de25dab29ff47d663d739c03.NewManagedTenantTicketingEndpointsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedTenantTicketingEndpointsById provides operations to manage the managedTenantTicketingEndpoints property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantTicketingEndpointsById(id string)(*i66506e2a89f3ae47e0d61d5a939ab1798a7d4fb9012b322b2c509d0c7ee44761.ManagedTenantTicketingEndpointItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantTicketingEndpoint%2Did"] = id
    }
    return i66506e2a89f3ae47e0d61d5a939ab1798a7d4fb9012b322b2c509d0c7ee44761.NewManagedTenantTicketingEndpointItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagementActions provides operations to manage the managementActions property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementActions()(*i9aa242f436094e02f161a3df7791adfa9932ceaf20ebf440b6069847018656e0.ManagementActionsRequestBuilder) {
    return i9aa242f436094e02f161a3df7791adfa9932ceaf20ebf440b6069847018656e0.NewManagementActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementActionsById provides operations to manage the managementActions property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementActionsById(id string)(*ief008158d809c9b248f0635950f9a44dacadb584a83e78751c8b87051d125c7a.ManagementActionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementAction%2Did"] = id
    }
    return ief008158d809c9b248f0635950f9a44dacadb584a83e78751c8b87051d125c7a.NewManagementActionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagementActionTenantDeploymentStatuses provides operations to manage the managementActionTenantDeploymentStatuses property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementActionTenantDeploymentStatuses()(*i39d20a98967701e034567f952771787304d37316148dcd57a17798ecd8e6b196.ManagementActionTenantDeploymentStatusesRequestBuilder) {
    return i39d20a98967701e034567f952771787304d37316148dcd57a17798ecd8e6b196.NewManagementActionTenantDeploymentStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementActionTenantDeploymentStatusesById provides operations to manage the managementActionTenantDeploymentStatuses property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementActionTenantDeploymentStatusesById(id string)(*ib007e276e70fb157eccc427006d3a6fdf591c84e085e79a6ae7d6a7149a17922.ManagementActionTenantDeploymentStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementActionTenantDeploymentStatus%2Did"] = id
    }
    return ib007e276e70fb157eccc427006d3a6fdf591c84e085e79a6ae7d6a7149a17922.NewManagementActionTenantDeploymentStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagementIntents provides operations to manage the managementIntents property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementIntents()(*icfac7b17b942815ef506ee4f3661445d6355fd31ed1673ad9d6ef8ebe8b57487.ManagementIntentsRequestBuilder) {
    return icfac7b17b942815ef506ee4f3661445d6355fd31ed1673ad9d6ef8ebe8b57487.NewManagementIntentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementIntentsById provides operations to manage the managementIntents property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementIntentsById(id string)(*i1ec6b65aff2fcd462f8f45f48ce513ea76dde61b1ec3f1cf01d65b849b740b8b.ManagementIntentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementIntent%2Did"] = id
    }
    return i1ec6b65aff2fcd462f8f45f48ce513ea76dde61b1ec3f1cf01d65b849b740b8b.NewManagementIntentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagementTemplateCollections provides operations to manage the managementTemplateCollections property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplateCollections()(*i86ccee4069fa2a8887f6c01c1bd59869700242c19224ee8b7aa8d4ad7bfbb89f.ManagementTemplateCollectionsRequestBuilder) {
    return i86ccee4069fa2a8887f6c01c1bd59869700242c19224ee8b7aa8d4ad7bfbb89f.NewManagementTemplateCollectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementTemplateCollectionsById provides operations to manage the managementTemplateCollections property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplateCollectionsById(id string)(*i8612235b849d43a7141e3638c7b76b5ba29594d12c5272cfc9e32333a038808a.ManagementTemplateCollectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementTemplateCollection%2Did"] = id
    }
    return i8612235b849d43a7141e3638c7b76b5ba29594d12c5272cfc9e32333a038808a.NewManagementTemplateCollectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagementTemplateCollectionTenantSummaries provides operations to manage the managementTemplateCollectionTenantSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplateCollectionTenantSummaries()(*i1e0d9f981a2f7ec6977de1f608176ed15d12b8f4c41d22354dc9c1610b826a35.ManagementTemplateCollectionTenantSummariesRequestBuilder) {
    return i1e0d9f981a2f7ec6977de1f608176ed15d12b8f4c41d22354dc9c1610b826a35.NewManagementTemplateCollectionTenantSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementTemplateCollectionTenantSummariesById provides operations to manage the managementTemplateCollectionTenantSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplateCollectionTenantSummariesById(id string)(*ib38449635f2df87835a316e9d35bdd8ac3d8641ecd90714b1eacf60a43539b76.ManagementTemplateCollectionTenantSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementTemplateCollectionTenantSummary%2Did"] = id
    }
    return ib38449635f2df87835a316e9d35bdd8ac3d8641ecd90714b1eacf60a43539b76.NewManagementTemplateCollectionTenantSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagementTemplates provides operations to manage the managementTemplates property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplates()(*i2b66c9972458318dcd18b57d227c1397d3704eed4ceba1aab769bae8473fef7c.ManagementTemplatesRequestBuilder) {
    return i2b66c9972458318dcd18b57d227c1397d3704eed4ceba1aab769bae8473fef7c.NewManagementTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementTemplatesById provides operations to manage the managementTemplates property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplatesById(id string)(*i36371b16b14a3a26bf98be71cdf018d4cde82965ba2327d328aeac3639125ea0.ManagementTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementTemplate%2Did"] = id
    }
    return i36371b16b14a3a26bf98be71cdf018d4cde82965ba2327d328aeac3639125ea0.NewManagementTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagementTemplateSteps provides operations to manage the managementTemplateSteps property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplateSteps()(*i9c7ec006c2124aad1f6bd0b710b33ddb680a37848847978f567ac657d103fd3a.ManagementTemplateStepsRequestBuilder) {
    return i9c7ec006c2124aad1f6bd0b710b33ddb680a37848847978f567ac657d103fd3a.NewManagementTemplateStepsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementTemplateStepsById provides operations to manage the managementTemplateSteps property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplateStepsById(id string)(*i7002a802b3e3d4667259a74b2f46e5190bce103c7bacef3e3630f3808a1f4256.ManagementTemplateStepItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementTemplateStep%2Did"] = id
    }
    return i7002a802b3e3d4667259a74b2f46e5190bce103c7bacef3e3630f3808a1f4256.NewManagementTemplateStepItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagementTemplateStepTenantSummaries provides operations to manage the managementTemplateStepTenantSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplateStepTenantSummaries()(*i6c8ccc9d63db57b24eef1a198f4256d2b75c2f3d924dfb09f7dc3934cfc5dc56.ManagementTemplateStepTenantSummariesRequestBuilder) {
    return i6c8ccc9d63db57b24eef1a198f4256d2b75c2f3d924dfb09f7dc3934cfc5dc56.NewManagementTemplateStepTenantSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementTemplateStepTenantSummariesById provides operations to manage the managementTemplateStepTenantSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplateStepTenantSummariesById(id string)(*i11da6ef4eeeb8786fb93d5972f290e02c6b190e32733d30dff04b7a103677aff.ManagementTemplateStepTenantSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementTemplateStepTenantSummary%2Did"] = id
    }
    return i11da6ef4eeeb8786fb93d5972f290e02c6b190e32733d30dff04b7a103677aff.NewManagementTemplateStepTenantSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagementTemplateStepVersions provides operations to manage the managementTemplateStepVersions property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplateStepVersions()(*id0d999b2f9f0ee37ec597aead92f01550dd34dbc53eee9ed3fe40e49f0d53bf7.ManagementTemplateStepVersionsRequestBuilder) {
    return id0d999b2f9f0ee37ec597aead92f01550dd34dbc53eee9ed3fe40e49f0d53bf7.NewManagementTemplateStepVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementTemplateStepVersionsById provides operations to manage the managementTemplateStepVersions property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplateStepVersionsById(id string)(*ibeef58461c8396867818d36875356dc3e16db01443724532b324a94c46ca8d52.ManagementTemplateStepVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementTemplateStepVersion%2Did"] = id
    }
    return ibeef58461c8396867818d36875356dc3e16db01443724532b324a94c46ca8d52.NewManagementTemplateStepVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MyRoles provides operations to manage the myRoles property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) MyRoles()(*i71e724937f2266c62830468204cfd18d9303a53fb8e960124d76a13f5477fb03.MyRolesRequestBuilder) {
    return i71e724937f2266c62830468204cfd18d9303a53fb8e960124d76a13f5477fb03.NewMyRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MyRolesById provides operations to manage the myRoles property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) MyRolesById(id string)(*id8a75758c62dfd2efe128ae3797671432b468fe56c4706706dfeb1c3655c9799.MyRoleTenantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["myRole%2DtenantId"] = id
    }
    return id8a75758c62dfd2efe128ae3797671432b468fe56c4706706dfeb1c3655c9799.NewMyRoleTenantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property managedTenants in tenantRelationships
func (m *ManagedTenantsRequestBuilder) Patch(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantable, requestConfiguration *ManagedTenantsRequestBuilderPatchRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagedTenantFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantable), nil
}
// TenantGroups provides operations to manage the tenantGroups property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) TenantGroups()(*icf1f2046c60f5e428ce03f4d34ebbf94d9da108c7a89ad5560774e04115aee81.TenantGroupsRequestBuilder) {
    return icf1f2046c60f5e428ce03f4d34ebbf94d9da108c7a89ad5560774e04115aee81.NewTenantGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TenantGroupsById provides operations to manage the tenantGroups property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) TenantGroupsById(id string)(*i79c491f13216abf413ae87d0ff1c465e6a8a4809a3f15a3c0d1e120851ef9a5c.TenantGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tenantGroup%2Did"] = id
    }
    return i79c491f13216abf413ae87d0ff1c465e6a8a4809a3f15a3c0d1e120851ef9a5c.NewTenantGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Tenants provides operations to manage the tenants property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) Tenants()(*i507510374fef5a86fcce6eb7ecdb044e4a7bc928ee54765d0620b2b896a746e1.TenantsRequestBuilder) {
    return i507510374fef5a86fcce6eb7ecdb044e4a7bc928ee54765d0620b2b896a746e1.NewTenantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TenantsById provides operations to manage the tenants property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) TenantsById(id string)(*i902764218ad5577611d10586599ccdb34a8aa22b15a8c9043e035e89db4d4c7b.TenantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tenant%2Did"] = id
    }
    return i902764218ad5577611d10586599ccdb34a8aa22b15a8c9043e035e89db4d4c7b.NewTenantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TenantsCustomizedInformation provides operations to manage the tenantsCustomizedInformation property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) TenantsCustomizedInformation()(*i824a81edbead3d58bdec21d61a960b3c5f03f44f0e89355db9cc5414e1d67e10.TenantsCustomizedInformationRequestBuilder) {
    return i824a81edbead3d58bdec21d61a960b3c5f03f44f0e89355db9cc5414e1d67e10.NewTenantsCustomizedInformationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TenantsCustomizedInformationById provides operations to manage the tenantsCustomizedInformation property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) TenantsCustomizedInformationById(id string)(*i22d7e4fc2a16cf517b6c379250cc88fc88b4ed4261b52362648bc93110b5f230.TenantCustomizedInformationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tenantCustomizedInformation%2Did"] = id
    }
    return i22d7e4fc2a16cf517b6c379250cc88fc88b4ed4261b52362648bc93110b5f230.NewTenantCustomizedInformationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TenantsDetailedInformation provides operations to manage the tenantsDetailedInformation property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) TenantsDetailedInformation()(*i23841075703fcb028662b4d36c5564ffa62324dd12cbd9935ed7d56f0dc0ca9f.TenantsDetailedInformationRequestBuilder) {
    return i23841075703fcb028662b4d36c5564ffa62324dd12cbd9935ed7d56f0dc0ca9f.NewTenantsDetailedInformationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TenantsDetailedInformationById provides operations to manage the tenantsDetailedInformation property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) TenantsDetailedInformationById(id string)(*ic98453c829b49eb4768ddc3245c452c7c68722a915539ad220764ab7c19350c3.TenantDetailedInformationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tenantDetailedInformation%2Did"] = id
    }
    return ic98453c829b49eb4768ddc3245c452c7c68722a915539ad220764ab7c19350c3.NewTenantDetailedInformationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TenantTags provides operations to manage the tenantTags property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) TenantTags()(*iab13ec14c24b8dac8fd45f37a3723dad71088c5ec497397272541acbbb27e30e.TenantTagsRequestBuilder) {
    return iab13ec14c24b8dac8fd45f37a3723dad71088c5ec497397272541acbbb27e30e.NewTenantTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TenantTagsById provides operations to manage the tenantTags property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) TenantTagsById(id string)(*i44e5d1aa3329234fb97b49ec0893e801904dd6a9768684f18a00bd816da70fd6.TenantTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tenantTag%2Did"] = id
    }
    return i44e5d1aa3329234fb97b49ec0893e801904dd6a9768684f18a00bd816da70fd6.NewTenantTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsDeviceMalwareStates provides operations to manage the windowsDeviceMalwareStates property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) WindowsDeviceMalwareStates()(*i0abc6f978704af5e99b7eb130a6c0d2e5c1e092e355b17b5af5cdc81f3564b2f.WindowsDeviceMalwareStatesRequestBuilder) {
    return i0abc6f978704af5e99b7eb130a6c0d2e5c1e092e355b17b5af5cdc81f3564b2f.NewWindowsDeviceMalwareStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsDeviceMalwareStatesById provides operations to manage the windowsDeviceMalwareStates property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) WindowsDeviceMalwareStatesById(id string)(*i186e85376ac4af3f7e13d1b03ae3ecca0403b582ff4761f086be79d2ceacdf75.WindowsDeviceMalwareStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsDeviceMalwareState%2Did"] = id
    }
    return i186e85376ac4af3f7e13d1b03ae3ecca0403b582ff4761f086be79d2ceacdf75.NewWindowsDeviceMalwareStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsProtectionStates provides operations to manage the windowsProtectionStates property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) WindowsProtectionStates()(*i65beab9150aa7abd814be391125d812f4d3c3bdfde9d2016db1f22a146166aee.WindowsProtectionStatesRequestBuilder) {
    return i65beab9150aa7abd814be391125d812f4d3c3bdfde9d2016db1f22a146166aee.NewWindowsProtectionStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsProtectionStatesById provides operations to manage the windowsProtectionStates property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) WindowsProtectionStatesById(id string)(*ib77c9f0cca75f19b263afc4ae06ae033ceee00cde50068f4a1ad949ad4e69035.WindowsProtectionStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsProtectionState%2Did"] = id
    }
    return ib77c9f0cca75f19b263afc4ae06ae033ceee00cde50068f4a1ad949ad4e69035.NewWindowsProtectionStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
