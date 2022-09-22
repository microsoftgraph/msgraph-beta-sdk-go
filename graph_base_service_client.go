package msgraphbetasdkgo

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
	i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
	i424c8587488111febed6b9b4d2ad6d5226fb83e28676c38366dd47f76f319ffb "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviewdecisions"
	i9568b23d9e69e9a4defd55b61d890372f14eadc42ca82c111ec03da32b7ae304 "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviewdecisions/item"
	i1dfcb6e17563ae78b6dbaf02d32cee89099a7795106760d7d401df42ce73b8fc "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews"
	i1013b3624b41d6ff1f6ebeba5ecd89cd7627a75e8f2716bcd854eca9c4f128bf "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item"
	i8d3c03812535daaab5e9e28f499097b08f09a8a7ab62e664ebf24dd8af17e77c "github.com/microsoftgraph/msgraph-beta-sdk-go/activitystatistics"
	i55744e9ed825874e254c58dfc594879e36123cdec77313d1b77ce713a84fd52f "github.com/microsoftgraph/msgraph-beta-sdk-go/activitystatistics/item"
	i65ed27543dee9887d3df7d7d883303dfead48cba6be4e357fa7d5c21332b4626 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin"
	i895a6f3a85eea8558280363fb928ce992d75c89f1c602b57f1d6352b0af78e5f "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits"
	i1838d03f06f3481fda4d2c9d067a6dae8c67a7a45256e9fc560c1aad4b954b44 "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item"
	i15e329825c659329448e12b30278e3a09efd68996edb65e6eb37bbba528b21d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/agreementacceptances"
	i2bb1d5c6cd7c743a7653900816a2fe176fb6e05046db53e58a4851c68cc0baa9 "github.com/microsoftgraph/msgraph-beta-sdk-go/agreementacceptances/item"
	i493c694f665c6b8116f1d28cef9c35839e2b3810e4a8c9f326bfc1b2caa30afa "github.com/microsoftgraph/msgraph-beta-sdk-go/agreements"
	i4a1aa16a38122b30b82f053fe001f17ca01cdd4d01f908800c1e275735405be9 "github.com/microsoftgraph/msgraph-beta-sdk-go/agreements/item"
	i8283dbd0b9136717b100e1ef6f4ac05d85e3412714dd0dea4b38d6de05ef4349 "github.com/microsoftgraph/msgraph-beta-sdk-go/alloweddatalocations"
	ib59058cb28b3ce2c94afbcff270a3ef835dbaf3a5098531014b61463524855b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/alloweddatalocations/item"
	ifc59747dbaa83f8f51942823114f4abfa41e0c0a64d67957f17e6b60407ce219 "github.com/microsoftgraph/msgraph-beta-sdk-go/app"
	i2130b9a37453c245bc87d9a83666a92560714fc5bb3c0f5a77e999639d2f4e45 "github.com/microsoftgraph/msgraph-beta-sdk-go/appcatalogs"
	ie1b2fd35e4b1f7cbc7bd808e462c966c4ec16a274923b50216bdd8a2ae0a3129 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications"
	i8aa59152dd26b67bc666e64bae862e596a1b291f95d9cd6ec0569c024334c2ea "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item"
	i9fb9a4d9d99571d2cc1de51809c0dfccf1dae8bd81c7eb39e51d1382c2ec81ba "github.com/microsoftgraph/msgraph-beta-sdk-go/applicationtemplates"
	ifb9dfb0b8ef1c8b2922bb3079994ba58d4aea3ca66033b8bb4dd80598f8f22ee "github.com/microsoftgraph/msgraph-beta-sdk-go/applicationtemplates/item"
	i35277464b5f866fcf2cb5324cd283216c9f6e9fc22956c71cb5b11c4ab649a29 "github.com/microsoftgraph/msgraph-beta-sdk-go/approleassignments"
	i5ba1440a4f16b81498d14796ed766b1d72d556daecb751300394b56c30f71708 "github.com/microsoftgraph/msgraph-beta-sdk-go/approleassignments/item"
	if4a9faac580b9d5510ead8eac155e0cecb2222152b913f0bedc9a44bbe2ee79e "github.com/microsoftgraph/msgraph-beta-sdk-go/approvalworkflowproviders"
	ie93af5fe517f51cf4801d2bf147043b2ea15ac3c26d73444ee01180255ba4c17 "github.com/microsoftgraph/msgraph-beta-sdk-go/approvalworkflowproviders/item"
	i6c3f8c4b4b571cf0fbb7c7c8791ae736e28cc3f4bb62262698b6291c13e127b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/auditlogs"
	i8e667c6208be96da3103b8806ff97028502c18052414fe99a224c1565834ca0f "github.com/microsoftgraph/msgraph-beta-sdk-go/authenticationmethodconfigurations"
	i9f2a6fa6c0a6eea0dc75c47461a955c3bf5d0e4f31c4695f5fda45cd2ac85e37 "github.com/microsoftgraph/msgraph-beta-sdk-go/authenticationmethodconfigurations/item"
	i39dbae52481ac3c9530d9fae0a2292348b8f7327bab28ea21183045324adadbc "github.com/microsoftgraph/msgraph-beta-sdk-go/authenticationmethodspolicy"
	ibf9394d7c54feda53ca523241dde659e8725041c25384ede68e72731d68d5abe "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses"
	ice28ccdda7ef9b32e133796796ceb464781909e74f15c614ef119ccb7cda6317 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses/item"
	ie573fcc25112f624100d67f5c4380ef23bdf060e2876e90cec1bf404deffc3dd "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingcurrencies"
	id4f763cc022952ed4430efa30c2908af332dea6f57f936d1cf7ef11fa10b8be8 "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingcurrencies/item"
	icbfa8075dc5e1ad04a3bc48d231d9b422f250ca1e2b74477a880e0db7b0e7f95 "github.com/microsoftgraph/msgraph-beta-sdk-go/branding"
	i97c9750160852aa25d52a4c6fa196b644ce728c6645ca520427ff4d85c76afa0 "github.com/microsoftgraph/msgraph-beta-sdk-go/businessflowtemplates"
	iced651caca0254342ce35ea2308d19515b0c471b819960bd801a22997ef6f386 "github.com/microsoftgraph/msgraph-beta-sdk-go/businessflowtemplates/item"
	id53bdaa191b823f3e2f4009f4cc095b46d1c7a433bde3b6d09ef0bd8df3514c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/certificatebasedauthconfiguration"
	id14b905eaea3ef6f00efe71e29906d015dab23238eb870ee09debe6eab8eea9b "github.com/microsoftgraph/msgraph-beta-sdk-go/certificatebasedauthconfiguration/item"
	i2bd7b88e5d2b5a20231449a09cd2703774075c53c19dc28ca3829e91d51dfd73 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats"
	ifc0fc132a789c973aa4786767e2f987d59e214f4abef3a4e82a606459d4b00c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item"
	i81513db44a7796569414887f1110e1d158078cb199b3565d960425af9c2904ba "github.com/microsoftgraph/msgraph-beta-sdk-go/commands"
	ie39c059add8b730dc85b7d2449e4b26b98f9a83d3807fabe7cb9009271179970 "github.com/microsoftgraph/msgraph-beta-sdk-go/commands/item"
	if7bcb57951e8f2ae550fcf781dc209ed777854429fcfe2465a71b03112dfc346 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications"
	ia93a0d553904ea998c5c0e30cf6d6574ba3a42be402a2d1eb88d5bd221d1a0d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance"
	i29b3182886b8fc3d309db2628f3911e671f6745e9fdab71d07cdeb487c4453ad "github.com/microsoftgraph/msgraph-beta-sdk-go/connections"
	ie9b2b2a73483057a6575063ce71086fbc54437022d4c53761e7a4a058d7dd091 "github.com/microsoftgraph/msgraph-beta-sdk-go/connections/item"
	i716e3204a4c47d24737c05f3b4c2ef2462fa5a1df29b57365f338e8f68ee16ef "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts"
	i6e94fa7602e5cf30291b967fc83a01373300f1d5183cc458d161d85fa3dc2baf "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item"
	i5e88b3e8025e26d21777096f2c7525a182545bb8dc0137634b047a10fe14ab54 "github.com/microsoftgraph/msgraph-beta-sdk-go/contracts"
	i09e36d06bea20a5d152cc52ffa7ef3c082c8cdfbf6d2116b226488cbb8f849ed "github.com/microsoftgraph/msgraph-beta-sdk-go/contracts/item"
	if3c2dea1db099d8f9ce8b623a12f6291f276e2bbb50259658f584a3a85cf71b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification"
	i0f747ff1f24810ff51160697ed4229c9ca192f7b84644311b88fa3b475cc340d "github.com/microsoftgraph/msgraph-beta-sdk-go/datapolicyoperations"
	i5cb8828c84548ec02b63ac985f24d0e6038f8ba987f6d69cc0a62c4e8c7d7df4 "github.com/microsoftgraph/msgraph-beta-sdk-go/datapolicyoperations/item"
	idc4afe653def183ef95500aa004f556fdf3c3747771f17c8472ca3cad61cebf4 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement"
	i09893664b20e7c846b2bc7aaaf1cd7f554ed3d2c00ac11336bea4c3c3d859e09 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement"
	i0b4892b2f92a31e44541567b8065e8e7760cb336e17d7dacb9120a865d5b0a37 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices"
	i3629807faa94b36aa56a99f728c28616ef4be3a1816bf7413ce360f3f6ef1c9d "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item"
	i6e398703c86ec3814400d80161079e7253c4e25f4ba1adb0c8d31da236f7bcd7 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory"
	iaec68a3d2c3ba0a78ebb66cd93fd1c5d2a6e0450b97a0cf19d94cb58956bec1d "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryobjects"
	i901e1ca524c6a851c60abc23bfdb47050e754ebcf438922b25817ff66d56276a "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryobjects/item"
	i20621ebd49d2bb1ed6c592ae35dfa701db30564a91ff100d25b0dcdb142bd942 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles"
	i993beddae5d86c1977387e8b393295fbee6f591297fb6a91fc3501c6cb5ae0d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item"
	i5b4eb770497618728398e41e6ed415ad2b92d20f7ad45ba75277a5800d9a2a12 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroletemplates"
	i98e280807a93b84e9092eeb27b7840b536491d6fa3c9de876973f203933ecf7f "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroletemplates/item"
	iabd30d9ba6ae302fd2d3145e0da70036496a2167b9e0cbc049bab96d9d9a29b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/directorysettingtemplates"
	i942fb2a58a06777c286167881e19f5c2b00590a6cc7831e7a502f112db658679 "github.com/microsoftgraph/msgraph-beta-sdk-go/directorysettingtemplates/item"
	i5db103717157eeba829a6a8580002d8b9adfdfd549628f3c84152bf3164a7b53 "github.com/microsoftgraph/msgraph-beta-sdk-go/domaindnsrecords"
	if2c2cfdc1fc8e4c78fc654fd9ff30fce9a231a7ccf8c16c510d1dbb98e3aa7ab "github.com/microsoftgraph/msgraph-beta-sdk-go/domaindnsrecords/item"
	i9e22e53c888822daa9264a72be4d11f335e3170e9198aae4bba758214e319857 "github.com/microsoftgraph/msgraph-beta-sdk-go/domains"
	ia82ac058caa6740adb23fd487affca202963b9786c4890c80eaec27b6d7ced24 "github.com/microsoftgraph/msgraph-beta-sdk-go/domains/item"
	idc403d989fa3b67787cd15ceedea7e45400694aa996d133f9bf61a54fe5a497a "github.com/microsoftgraph/msgraph-beta-sdk-go/drive"
	i2b9cd6123f9a7ca2d7c253973e81bb3965ee0b78a350919aafe66d7b73c70433 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives"
	ib8b674690e8be27a974ea55202f0d976b3d187306aca08f5642097d6cfed398b "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item"
	id90d135edac1f1a3e952db4ad985001105d2e7c0133f8cc410765eb1af789cc0 "github.com/microsoftgraph/msgraph-beta-sdk-go/education"
	i4b923a0ffb143d5def24980fce5d55c68c9634d5b55c33bb0b3029ac68415dd1 "github.com/microsoftgraph/msgraph-beta-sdk-go/employeeexperience"
	i453454978d87486fd201df62ea4b775c5b4907e2a36395fb6fb1e9060fc3f1bb "github.com/microsoftgraph/msgraph-beta-sdk-go/external"
	i0b9d70018d3c267f9f34a818ce43cc889d06d87749a70e1ad1d45eead0c735e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/filteroperators"
	i330b70aeb45059ed2410e099c3692c88ec3d2474fbb8363fffbecba6f900c494 "github.com/microsoftgraph/msgraph-beta-sdk-go/filteroperators/item"
	i0281e791d6eedd5b4b6d3c18b07336722b6152b5db6a2e9dc8385e98f565f677 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials"
	i054b68521cee54ec767d07cf7a8a6d50a0d24b6e6fc43b8296c34730fd8ca465 "github.com/microsoftgraph/msgraph-beta-sdk-go/functions"
	ib3f85f7711d9ea4057703e0171ba196ac1d884a1e2389602111d38d84fb1380f "github.com/microsoftgraph/msgraph-beta-sdk-go/functions/item"
	i3b7da1b693d5428b20b0bf3340acb4b879042a9393e45df9349b04a5b2830acb "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceresources"
	iaeceabcb9ce892827c1e040d7ed5dec4e2fb051c4a1cb3b857ac26d621e94638 "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceresources/item"
	i3de07fd1833246d183a4d60c2329c7467381716c2cfdfc096ff627e4e9479cf8 "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroleassignmentrequests"
	if03454ccae6aab82c8d7c01d50047eceb6218bf085d9949f9f4ff37652bab5a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroleassignmentrequests/item"
	i14752cfec59ab915e7c63922270765abf65744437d9135c191cef3986f08c3bb "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroleassignments"
	i5e0d644f8c8abdc66bf30911d1b935134f92408badf2ef700cf692be1de56856 "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroleassignments/item"
	i8a18cb7418541221b2c3fd213a484d9e3029fab916358b16fb24015b078b8eba "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroledefinitions"
	i005472295f05513a4d4bcabd67693926f332f4d4740482ce99613ae21d4ea87a "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroledefinitions/item"
	ib71e32ed3a7f0f8a512aa55c1428492116ff2d1bae5015a9b89f910ecbc7c6bd "github.com/microsoftgraph/msgraph-beta-sdk-go/governancerolesettings"
	i2a5979c2bdbc99c804ec546b7cec503201936edc913e5d36d239a54c3ad002ad "github.com/microsoftgraph/msgraph-beta-sdk-go/governancerolesettings/item"
	ifcac309012d761a79a74e6d79fad6979f2117e7af36ff6e5ad131093412afcc7 "github.com/microsoftgraph/msgraph-beta-sdk-go/governancesubjects"
	i6f49755cbc32687ce5fd8e0bc2cdb1808c45f60ba20889d053bac06929b7d9ea "github.com/microsoftgraph/msgraph-beta-sdk-go/governancesubjects/item"
	i71438b4a3f9d4a17f8c873a44b8ac76600403f5ce0cce2423bde35e0191f2c17 "github.com/microsoftgraph/msgraph-beta-sdk-go/grouplifecyclepolicies"
	if2d4db22e4ec9d990102702a0f1b228fbfa6a90e349c6dfb2f3c798d2c5ec77f "github.com/microsoftgraph/msgraph-beta-sdk-go/grouplifecyclepolicies/item"
	i79f2b866e8bec3ee9349dc885ecb3691e94b20459995d83b3dbf9f05341c7a89 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups"
	ic6c550f1e18a80bf95a0e1377fab585951c941746c2a2739d2b1d77b17b80a94 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item"
	i20702653f98186060bd39b9fe8136743eafc0ddaa43435e527665ac75229a33a "github.com/microsoftgraph/msgraph-beta-sdk-go/identity"
	i7ed0a0c7cc963cb1cf6c282e5e5c04cea2d4959e5f1bfd12698c0196858b1b52 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance"
	i444aa080f7aab1ac35ae55318eca861a558a57b4c5e3cbd071222f511baac5fe "github.com/microsoftgraph/msgraph-beta-sdk-go/identityprotection"
	i531b1efd1768fd272d51921ff5812bdeba5b46e0eeec0e4c818250cb7116aed5 "github.com/microsoftgraph/msgraph-beta-sdk-go/identityproviders"
	ib66730b67f84ecbb45dd98e1ce34bd456667b7437238322804d692193cbb7cd6 "github.com/microsoftgraph/msgraph-beta-sdk-go/identityproviders/item"
	ida1e35de3db89946d53da45357ba0e3da7b3e3a1d46b511191582f0d2a917caa "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection"
	i4cb6edb865a0e38bb1799dcb0c7881b92feed59596c1912cfe5e6142b61f9c91 "github.com/microsoftgraph/msgraph-beta-sdk-go/invitations"
	i5cfcfa8bc4a6d5fed3fa8ff67c53050a2b9493ff4ba9481077f03d4df2e84db6 "github.com/microsoftgraph/msgraph-beta-sdk-go/invitations/item"
	i5a50ec7cbb7e52feaacf8d45f5e247e765502c09622e2fa44cd68445eab876eb "github.com/microsoftgraph/msgraph-beta-sdk-go/me"
	i19717748912ff29c95998304f534371c35864a4579ea92608b32aeecf7d18cc4 "github.com/microsoftgraph/msgraph-beta-sdk-go/messageevents"
	i2737b8a250008797e3665bb01cdf0bfbf42bedbb23c7d023eadf66c5dc36b26e "github.com/microsoftgraph/msgraph-beta-sdk-go/messageevents/item"
	i56ad7deac03a612015589ab4fade2313d6df08086c7ee8d46177fc8ddc5b0053 "github.com/microsoftgraph/msgraph-beta-sdk-go/messagerecipients"
	ia62d215a2c5a4beae9bb26027fb9b0f0438b6e34a48aa59b64ddddde97e1e648 "github.com/microsoftgraph/msgraph-beta-sdk-go/messagerecipients/item"
	i43a1701c7f8902247fd60e3d6ff36be02d1a59f02884ce5734adb8deb69a47ef "github.com/microsoftgraph/msgraph-beta-sdk-go/messagetraces"
	iaa37a57ee21d5046d5b74ae373a2d314357f709507fd2fb2116379f39879b88a "github.com/microsoftgraph/msgraph-beta-sdk-go/messagetraces/item"
	i05bd1def68419ff4c77b9bdafb42d256c4c19b479003ef8adcd868ec38673e84 "github.com/microsoftgraph/msgraph-beta-sdk-go/mobilitymanagementpolicies"
	i0740547a5aa3dd54a8b262a3d1cd5e0bb409b1a645795e041ffa10dcb9019e0a "github.com/microsoftgraph/msgraph-beta-sdk-go/mobilitymanagementpolicies/item"
	i3c2d5b7a05b0c58ca1d3b72343c5f03221f2bd894e75f89f62ccf8020a48250b "github.com/microsoftgraph/msgraph-beta-sdk-go/monitoring"
	ie2e0818e93fcfbb33fde071a9354c3c22bedab0ec20855b7d5232d29bcc65bad "github.com/microsoftgraph/msgraph-beta-sdk-go/oauth2permissiongrants"
	if237177d98e26841b842fcf602048c102ab53f50428ff601c17ea0795d7fb828 "github.com/microsoftgraph/msgraph-beta-sdk-go/oauth2permissiongrants/item"
	ibcc2ba1bd11f45f4381ae4007e619ef47b74dbf731b922816aff1ed4d5d47a0f "github.com/microsoftgraph/msgraph-beta-sdk-go/officeconfiguration"
	i0d38f6e6ea6126fff7bb7a5c3c2d82fe471d00233209e2b6b2ce6ccb21ce50f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles"
	ie1c23a9b73355bb18364b81d44931ba2a9c3f409f2e091f8c44725c8075f8cf9 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item"
	ie58948149bb028757a64f16376df00cc5a99ad93e0d57affa0ac24ff6d096aaf "github.com/microsoftgraph/msgraph-beta-sdk-go/organization"
	ic9a01566deb7ac7de88728b70eec29525d2a94adff4181ea235c30369d4c108c "github.com/microsoftgraph/msgraph-beta-sdk-go/organization/item"
	i22f037221f506c5b5751f13095cc17caaf248e93588f883123c452cb4f1ec6a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/payloadresponse"
	ie1b5a10a367968f2f9e397679c2e84cb9c8ec8475b94ece386d4dcd129bcefe8 "github.com/microsoftgraph/msgraph-beta-sdk-go/payloadresponse/item"
	i48a68a7c83dc874f9d9fdf942afed70a34b11f92d6b2ccb439359753116f65cc "github.com/microsoftgraph/msgraph-beta-sdk-go/permissiongrants"
	i79b93b64b0d0d7fc163598b550442694212cae76f99c0fc0cc0b570453d3decb "github.com/microsoftgraph/msgraph-beta-sdk-go/permissiongrants/item"
	ic21cf429efd6fc3199e67b5b4288a3193ff5e9cfb4e97a5e442e02ccd7748ec1 "github.com/microsoftgraph/msgraph-beta-sdk-go/places"
	ia6735253ec4d9448b4b9541b92358386ce1936ac15c0298ba04aaff0a1e74e3d "github.com/microsoftgraph/msgraph-beta-sdk-go/places/item"
	i2c16e7aacab2db540fbc4e11916f266c9c46936118144cc1383e798af27b1b6f "github.com/microsoftgraph/msgraph-beta-sdk-go/planner"
	i5f794bd004f1cc95da776bcb1947ffabf97b71aae1f5c9f15255b24451e2929b "github.com/microsoftgraph/msgraph-beta-sdk-go/policies"
	i75b6dc07087cda1a9afc465878b0aa56ca3703a3ed530d5a22119b0960d159d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/print"
	i7daddf90c53ab17a62cc6612f32636f5bea354e4e22c20c5656606d1b491bd03 "github.com/microsoftgraph/msgraph-beta-sdk-go/privacy"
	i2f87335be6a07866636e1f602f5beda6b47bc99969e216fac59efff432ff2345 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess"
	ibc6cc10ee3264e7addfcd3086a9a46a8a89142b950c482b08cde11f41f711bc2 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item"
	i613e9d4b5596faf0dba59087b1a65d06d17aab1ca9d69170b292a4e0d90063ab "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedapproval"
	i3056c9ce4b136c43e7e58e79400328464aedd7296804666295056c525cd5afe9 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedapproval/item"
	i82672a497b8f66c510d59f2a80cb96da07ffad912f0c452f2d22f4a282e37720 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedoperationevents"
	i34d9be3c24f9832412eaf838a313f099ae8c0f497b918534db057b1bd9c8316f "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedoperationevents/item"
	if11203f2b5e6319285361e1998b4a25572cc1950d617d10ffd84e91e5f477349 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroleassignmentrequests"
	i7732e0def4b0366de3c1a4a0c7517a0f56ccce5a90b647e1a0e72997c0060fb1 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroleassignmentrequests/item"
	i85eafc30b6e6233aefbca2b4e35ced0ac8836418de9401407873193666d812d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroleassignments"
	i7a3a499cefaec6ced318a596ac37e721c173d3074f10b53b0db7be2eb1d91d1a "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroleassignments/item"
	i32f7b810493504f82bd6c97020faab5a8ff5f541a46dd3d6b9cc2aa77fc22fe4 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroles"
	i83a555d9eb9470ebf76ece54f97987bafd2a508dc0aa9e61e7d07311c57975c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroles/item"
	i6b96a96c52bbdff1731b8a5490cd5f342e33866e0931912944d323bc79f663e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedsignupstatus"
	ibe5c8d6d0c440fd8b6655bd846527fdabaf68dc6c79eb41da97daa4a59fd9b9f "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedsignupstatus/item"
	i178f0aa4e5987fcbfe2e98cbb6dd777ebcdcdf124dd3478d2bf40f83912ca030 "github.com/microsoftgraph/msgraph-beta-sdk-go/programcontrols"
	i9897ce060c8c140b6319d9eec7b10f1133ddeadce443b5c08d9e6e0b1a90a1fa "github.com/microsoftgraph/msgraph-beta-sdk-go/programcontrols/item"
	if5cd0cf36bc86d9253920d73c41189ad8a30342e678d4f0138afa5095fd31538 "github.com/microsoftgraph/msgraph-beta-sdk-go/programcontroltypes"
	if01cb9f5e03c6cceecfc9647206bb0dc6e0bb97d9f416977b2683028defae8c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/programcontroltypes/item"
	ib85b32f0384596c14f04b8d0f3dc8737da4b97428d7af145db2f1b06d7d9444e "github.com/microsoftgraph/msgraph-beta-sdk-go/programs"
	i5b0c38b2eba003d29699607d25825d237b1b42b0380942484eba07941fd7ff1d "github.com/microsoftgraph/msgraph-beta-sdk-go/programs/item"
	ia5f30c11d37332bb7dfb48303b8e8880e60fc3560f60a813c66b9c1c3f2ff3f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports"
	i24b998459b0fbcdc6fbd83b90048a098ff6bbdcdaff773a2886f5cf8b3d5545e "github.com/microsoftgraph/msgraph-beta-sdk-go/riskdetections"
	i4dba4c746f2c06c2fec91dc54833c1da81c07726a80b85b29e72221e93c811ff "github.com/microsoftgraph/msgraph-beta-sdk-go/riskdetections/item"
	iac77c9b5b86109e8ad626e30830db719efb3cc77c7babab332b409d84ae324a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/riskyusers"
	i4f10f292ecab93ead1a3540c8285847ac93e9fd17579d21e923687140834b553 "github.com/microsoftgraph/msgraph-beta-sdk-go/riskyusers/item"
	i310cda3e9f244aa61f9c9c78de433773f341b91c4a2310b8991671fe773be16e "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement"
	i1b84a2c37ba0bbd175c6da40c8679db7d04dfcb044d8421d26d024db45218e4a "github.com/microsoftgraph/msgraph-beta-sdk-go/schemaextensions"
	i03c5278b23fd725e2f49b358180d02b8204a81547ad7b8391ef4d31869b4e77c "github.com/microsoftgraph/msgraph-beta-sdk-go/schemaextensions/item"
	i712907ad27a66d6ac32a26e01f88de1ad6484585eb7ed65f84b3a30571cec55b "github.com/microsoftgraph/msgraph-beta-sdk-go/scopedrolememberships"
	i9febca1a7c0d482f5cfd0d294e0b42f41284df1fc0f26b63341fd7066192a1d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/scopedrolememberships/item"
	i5840582f75a8eb78900edf3bd78566223ffee7aa1dc2f4cdca943ef635f6503e "github.com/microsoftgraph/msgraph-beta-sdk-go/search"
	i761e9f0dec20dbf36c7fd626d107fb81ef94cafa7369422d2b2af143ffa16184 "github.com/microsoftgraph/msgraph-beta-sdk-go/security"
	ibd3e65bb14e91a8a05d902c54fadec2c1b6931676c97f76da4969c975770aab2 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals"
	i1486770fbd64ddb3f673e750b91158c408009c45bd8e48463d5e06498271cd9c "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item"
	i714cbeb65962cb4d3e58007792fa4832d175c04614ba3aa7efb22871aea885bf "github.com/microsoftgraph/msgraph-beta-sdk-go/settings"
	ie532fdcd6a215be3d5886dcd2e4608af81583d060c7156a283880daf04ed9903 "github.com/microsoftgraph/msgraph-beta-sdk-go/settings/item"
	i73583652789c7aab226ac5bae66bc7b5fd924607d28350c4478c2a20524fd624 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares"
	id8327d9e69ee76052759df500e96668a03b01b09bae354eaa67bbcf83427dc00 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item"
	i2817c6849b286be20c56215e039110b08a4109a12669579f4597fbab6f720ed9 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites"
	ic557b65175a90fb05406286b014b07b096a266f27fc59e78636733800d66ff09 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item"
	i7282246c788c7b44fcd028d8a1500ce1ffd32f655bee3b580b6a33d9ec10f610 "github.com/microsoftgraph/msgraph-beta-sdk-go/solutions"
	id14bafb4ee71e1257662cacab67dd479e54eba65ce10c6816ee4676804caf821 "github.com/microsoftgraph/msgraph-beta-sdk-go/storage"
	ie934faa615fb56652e5964395b3dc205321ac84e8cf244796ebe59ba3713fbd9 "github.com/microsoftgraph/msgraph-beta-sdk-go/subscribedskus"
	i3e0ee15b77f1a6df2686711e1700bf6812443daf08105eb0b15f4b72d6077b93 "github.com/microsoftgraph/msgraph-beta-sdk-go/subscribedskus/item"
	i60392bb6eb86abe7a3079e3b2b1e202f7dcf3452adf026db62ec93e2fafe3957 "github.com/microsoftgraph/msgraph-beta-sdk-go/subscriptions"
	ide40baaefabe61fc7133751aae85320ce5f6bc70e3dde81b419c4861d1e1bfe0 "github.com/microsoftgraph/msgraph-beta-sdk-go/subscriptions/item"
	iff395ba1da21566390b02b5bed781aecf3bb849fc71f2359410792d1d1b67079 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams"
	i86ffda5c6f886142a63c7cd1c0f489d73a5f8c36310f880ba8dd82807977d46e "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item"
	i66f18ccab4e34309d26d1056f0e7dd8b563a5f8ee6f8d9c6e8e77c5fac50f8b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamstemplates"
	i44e397fa023407b3a9722e619d02140543ef889c8c03de376d8341cec5a43794 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamstemplates/item"
	i1c7e7a5d0708841f8c98ec910d583f348cbffaad386ef9a24d3ee4eba285ea21 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition"
	i9edbc585c67673e5bfc1e312a7fb5f8fca5cb966cef31e6fdfe8ab8ab7d5780d "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition/item"
	icb4f253cb1cd35435f5752b611229032c618bbcfeb3be80ee4d6a06d404114fc "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork"
	id5c2ef977a00dd1757d258dbbbfb4080031771e62e6c6b3b1339a0f03fc1c1f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships"
	i98a1471d41b15330865bc87691830281af9ecf479bfc797e54f02448790b1e4e "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore"
	i231953b44a7f4aace0800ccf375537d423e6f60e82f2a6552d9613626e39aba5 "github.com/microsoftgraph/msgraph-beta-sdk-go/threatsubmission"
	i312c0a09d8ded5436957205a14adfc7e2facbcc6f26ef9872a5b5eb79228375f "github.com/microsoftgraph/msgraph-beta-sdk-go/trustframework"
	icd01c84a90833c55ac2309fd7034cb1962c60f59eb1ee2b2cf7b04c708402b6a "github.com/microsoftgraph/msgraph-beta-sdk-go/users"
	id9f276010196e2d79e634ec622333b5c53dc0fbc407a6c9aa27ca92d4f388ed3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item"
	i0c8acb61c4c4d82e7ac6f352bd7cdb24dcd0de64d904e31cf17ca8c6b1d44202 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks"
	i36b4c1bee4db70277ed8b9217e494480565564fd187f7a35b920807b10c82740 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item"
)

// GraphBaseServiceClient the main entry point of the SDK, exposes the configuration and the fluent API.
type GraphBaseServiceClient struct {
	// Path parameters for the request
	pathParameters map[string]string
	// The request adapter to use to execute the requests.
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
	// Url template to use to build the URL for the current request builder
	urlTemplate string
}

// AccessReviewDecisions the accessReviewDecisions property
func (m *GraphBaseServiceClient) AccessReviewDecisions() *i424c8587488111febed6b9b4d2ad6d5226fb83e28676c38366dd47f76f319ffb.AccessReviewDecisionsRequestBuilder {
	return i424c8587488111febed6b9b4d2ad6d5226fb83e28676c38366dd47f76f319ffb.NewAccessReviewDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// AccessReviewDecisionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.accessReviewDecisions.item collection
func (m *GraphBaseServiceClient) AccessReviewDecisionsById(id string) *i9568b23d9e69e9a4defd55b61d890372f14eadc42ca82c111ec03da32b7ae304.AccessReviewDecisionItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["accessReviewDecision%2Did"] = id
	}
	return i9568b23d9e69e9a4defd55b61d890372f14eadc42ca82c111ec03da32b7ae304.NewAccessReviewDecisionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// AccessReviews the accessReviews property
func (m *GraphBaseServiceClient) AccessReviews() *i1dfcb6e17563ae78b6dbaf02d32cee89099a7795106760d7d401df42ce73b8fc.AccessReviewsRequestBuilder {
	return i1dfcb6e17563ae78b6dbaf02d32cee89099a7795106760d7d401df42ce73b8fc.NewAccessReviewsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// AccessReviewsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.accessReviews.item collection
func (m *GraphBaseServiceClient) AccessReviewsById(id string) *i1013b3624b41d6ff1f6ebeba5ecd89cd7627a75e8f2716bcd854eca9c4f128bf.AccessReviewItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["accessReview%2Did"] = id
	}
	return i1013b3624b41d6ff1f6ebeba5ecd89cd7627a75e8f2716bcd854eca9c4f128bf.NewAccessReviewItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Activitystatistics the activitystatistics property
func (m *GraphBaseServiceClient) Activitystatistics() *i8d3c03812535daaab5e9e28f499097b08f09a8a7ab62e664ebf24dd8af17e77c.ActivitystatisticsRequestBuilder {
	return i8d3c03812535daaab5e9e28f499097b08f09a8a7ab62e664ebf24dd8af17e77c.NewActivitystatisticsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// ActivitystatisticsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.activitystatistics.item collection
func (m *GraphBaseServiceClient) ActivitystatisticsById(id string) *i55744e9ed825874e254c58dfc594879e36123cdec77313d1b77ce713a84fd52f.ActivityStatisticsItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["activityStatistics%2Did"] = id
	}
	return i55744e9ed825874e254c58dfc594879e36123cdec77313d1b77ce713a84fd52f.NewActivityStatisticsItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Admin the admin property
func (m *GraphBaseServiceClient) Admin() *i65ed27543dee9887d3df7d7d883303dfead48cba6be4e357fa7d5c21332b4626.AdminRequestBuilder {
	return i65ed27543dee9887d3df7d7d883303dfead48cba6be4e357fa7d5c21332b4626.NewAdminRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// AdministrativeUnits the administrativeUnits property
func (m *GraphBaseServiceClient) AdministrativeUnits() *i895a6f3a85eea8558280363fb928ce992d75c89f1c602b57f1d6352b0af78e5f.AdministrativeUnitsRequestBuilder {
	return i895a6f3a85eea8558280363fb928ce992d75c89f1c602b57f1d6352b0af78e5f.NewAdministrativeUnitsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// AdministrativeUnitsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.administrativeUnits.item collection
func (m *GraphBaseServiceClient) AdministrativeUnitsById(id string) *i1838d03f06f3481fda4d2c9d067a6dae8c67a7a45256e9fc560c1aad4b954b44.AdministrativeUnitItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["administrativeUnit%2Did"] = id
	}
	return i1838d03f06f3481fda4d2c9d067a6dae8c67a7a45256e9fc560c1aad4b954b44.NewAdministrativeUnitItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// AgreementAcceptances the agreementAcceptances property
func (m *GraphBaseServiceClient) AgreementAcceptances() *i15e329825c659329448e12b30278e3a09efd68996edb65e6eb37bbba528b21d7.AgreementAcceptancesRequestBuilder {
	return i15e329825c659329448e12b30278e3a09efd68996edb65e6eb37bbba528b21d7.NewAgreementAcceptancesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// AgreementAcceptancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.agreementAcceptances.item collection
func (m *GraphBaseServiceClient) AgreementAcceptancesById(id string) *i2bb1d5c6cd7c743a7653900816a2fe176fb6e05046db53e58a4851c68cc0baa9.AgreementAcceptanceItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["agreementAcceptance%2Did"] = id
	}
	return i2bb1d5c6cd7c743a7653900816a2fe176fb6e05046db53e58a4851c68cc0baa9.NewAgreementAcceptanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Agreements the agreements property
func (m *GraphBaseServiceClient) Agreements() *i493c694f665c6b8116f1d28cef9c35839e2b3810e4a8c9f326bfc1b2caa30afa.AgreementsRequestBuilder {
	return i493c694f665c6b8116f1d28cef9c35839e2b3810e4a8c9f326bfc1b2caa30afa.NewAgreementsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// AgreementsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.agreements.item collection
func (m *GraphBaseServiceClient) AgreementsById(id string) *i4a1aa16a38122b30b82f053fe001f17ca01cdd4d01f908800c1e275735405be9.AgreementItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["agreement%2Did"] = id
	}
	return i4a1aa16a38122b30b82f053fe001f17ca01cdd4d01f908800c1e275735405be9.NewAgreementItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// AllowedDataLocations the allowedDataLocations property
func (m *GraphBaseServiceClient) AllowedDataLocations() *i8283dbd0b9136717b100e1ef6f4ac05d85e3412714dd0dea4b38d6de05ef4349.AllowedDataLocationsRequestBuilder {
	return i8283dbd0b9136717b100e1ef6f4ac05d85e3412714dd0dea4b38d6de05ef4349.NewAllowedDataLocationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// AllowedDataLocationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.allowedDataLocations.item collection
func (m *GraphBaseServiceClient) AllowedDataLocationsById(id string) *ib59058cb28b3ce2c94afbcff270a3ef835dbaf3a5098531014b61463524855b8.AllowedDataLocationItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["allowedDataLocation%2Did"] = id
	}
	return ib59058cb28b3ce2c94afbcff270a3ef835dbaf3a5098531014b61463524855b8.NewAllowedDataLocationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// App the app property
func (m *GraphBaseServiceClient) App() *ifc59747dbaa83f8f51942823114f4abfa41e0c0a64d67957f17e6b60407ce219.AppRequestBuilder {
	return ifc59747dbaa83f8f51942823114f4abfa41e0c0a64d67957f17e6b60407ce219.NewAppRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// AppCatalogs the appCatalogs property
func (m *GraphBaseServiceClient) AppCatalogs() *i2130b9a37453c245bc87d9a83666a92560714fc5bb3c0f5a77e999639d2f4e45.AppCatalogsRequestBuilder {
	return i2130b9a37453c245bc87d9a83666a92560714fc5bb3c0f5a77e999639d2f4e45.NewAppCatalogsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Applications the applications property
func (m *GraphBaseServiceClient) Applications() *ie1b2fd35e4b1f7cbc7bd808e462c966c4ec16a274923b50216bdd8a2ae0a3129.ApplicationsRequestBuilder {
	return ie1b2fd35e4b1f7cbc7bd808e462c966c4ec16a274923b50216bdd8a2ae0a3129.NewApplicationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// ApplicationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.applications.item collection
func (m *GraphBaseServiceClient) ApplicationsById(id string) *i8aa59152dd26b67bc666e64bae862e596a1b291f95d9cd6ec0569c024334c2ea.ApplicationItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["application%2Did"] = id
	}
	return i8aa59152dd26b67bc666e64bae862e596a1b291f95d9cd6ec0569c024334c2ea.NewApplicationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// ApplicationTemplates the applicationTemplates property
func (m *GraphBaseServiceClient) ApplicationTemplates() *i9fb9a4d9d99571d2cc1de51809c0dfccf1dae8bd81c7eb39e51d1382c2ec81ba.ApplicationTemplatesRequestBuilder {
	return i9fb9a4d9d99571d2cc1de51809c0dfccf1dae8bd81c7eb39e51d1382c2ec81ba.NewApplicationTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// ApplicationTemplatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.applicationTemplates.item collection
func (m *GraphBaseServiceClient) ApplicationTemplatesById(id string) *ifb9dfb0b8ef1c8b2922bb3079994ba58d4aea3ca66033b8bb4dd80598f8f22ee.ApplicationTemplateItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["applicationTemplate%2Did"] = id
	}
	return ifb9dfb0b8ef1c8b2922bb3079994ba58d4aea3ca66033b8bb4dd80598f8f22ee.NewApplicationTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// AppRoleAssignments the appRoleAssignments property
func (m *GraphBaseServiceClient) AppRoleAssignments() *i35277464b5f866fcf2cb5324cd283216c9f6e9fc22956c71cb5b11c4ab649a29.AppRoleAssignmentsRequestBuilder {
	return i35277464b5f866fcf2cb5324cd283216c9f6e9fc22956c71cb5b11c4ab649a29.NewAppRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// AppRoleAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.appRoleAssignments.item collection
func (m *GraphBaseServiceClient) AppRoleAssignmentsById(id string) *i5ba1440a4f16b81498d14796ed766b1d72d556daecb751300394b56c30f71708.AppRoleAssignmentItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["appRoleAssignment%2Did"] = id
	}
	return i5ba1440a4f16b81498d14796ed766b1d72d556daecb751300394b56c30f71708.NewAppRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// ApprovalWorkflowProviders the approvalWorkflowProviders property
func (m *GraphBaseServiceClient) ApprovalWorkflowProviders() *if4a9faac580b9d5510ead8eac155e0cecb2222152b913f0bedc9a44bbe2ee79e.ApprovalWorkflowProvidersRequestBuilder {
	return if4a9faac580b9d5510ead8eac155e0cecb2222152b913f0bedc9a44bbe2ee79e.NewApprovalWorkflowProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// ApprovalWorkflowProvidersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.approvalWorkflowProviders.item collection
func (m *GraphBaseServiceClient) ApprovalWorkflowProvidersById(id string) *ie93af5fe517f51cf4801d2bf147043b2ea15ac3c26d73444ee01180255ba4c17.ApprovalWorkflowProviderItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["approvalWorkflowProvider%2Did"] = id
	}
	return ie93af5fe517f51cf4801d2bf147043b2ea15ac3c26d73444ee01180255ba4c17.NewApprovalWorkflowProviderItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// AuditLogs the auditLogs property
func (m *GraphBaseServiceClient) AuditLogs() *i6c3f8c4b4b571cf0fbb7c7c8791ae736e28cc3f4bb62262698b6291c13e127b9.AuditLogsRequestBuilder {
	return i6c3f8c4b4b571cf0fbb7c7c8791ae736e28cc3f4bb62262698b6291c13e127b9.NewAuditLogsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// AuthenticationMethodConfigurations the authenticationMethodConfigurations property
func (m *GraphBaseServiceClient) AuthenticationMethodConfigurations() *i8e667c6208be96da3103b8806ff97028502c18052414fe99a224c1565834ca0f.AuthenticationMethodConfigurationsRequestBuilder {
	return i8e667c6208be96da3103b8806ff97028502c18052414fe99a224c1565834ca0f.NewAuthenticationMethodConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// AuthenticationMethodConfigurationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.authenticationMethodConfigurations.item collection
func (m *GraphBaseServiceClient) AuthenticationMethodConfigurationsById(id string) *i9f2a6fa6c0a6eea0dc75c47461a955c3bf5d0e4f31c4695f5fda45cd2ac85e37.AuthenticationMethodConfigurationItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["authenticationMethodConfiguration%2Did"] = id
	}
	return i9f2a6fa6c0a6eea0dc75c47461a955c3bf5d0e4f31c4695f5fda45cd2ac85e37.NewAuthenticationMethodConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// AuthenticationMethodsPolicy the authenticationMethodsPolicy property
func (m *GraphBaseServiceClient) AuthenticationMethodsPolicy() *i39dbae52481ac3c9530d9fae0a2292348b8f7327bab28ea21183045324adadbc.AuthenticationMethodsPolicyRequestBuilder {
	return i39dbae52481ac3c9530d9fae0a2292348b8f7327bab28ea21183045324adadbc.NewAuthenticationMethodsPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// BookingBusinesses the bookingBusinesses property
func (m *GraphBaseServiceClient) BookingBusinesses() *ibf9394d7c54feda53ca523241dde659e8725041c25384ede68e72731d68d5abe.BookingBusinessesRequestBuilder {
	return ibf9394d7c54feda53ca523241dde659e8725041c25384ede68e72731d68d5abe.NewBookingBusinessesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// BookingBusinessesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.bookingBusinesses.item collection
func (m *GraphBaseServiceClient) BookingBusinessesById(id string) *ice28ccdda7ef9b32e133796796ceb464781909e74f15c614ef119ccb7cda6317.BookingBusinessItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["bookingBusiness%2Did"] = id
	}
	return ice28ccdda7ef9b32e133796796ceb464781909e74f15c614ef119ccb7cda6317.NewBookingBusinessItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// BookingCurrencies the bookingCurrencies property
func (m *GraphBaseServiceClient) BookingCurrencies() *ie573fcc25112f624100d67f5c4380ef23bdf060e2876e90cec1bf404deffc3dd.BookingCurrenciesRequestBuilder {
	return ie573fcc25112f624100d67f5c4380ef23bdf060e2876e90cec1bf404deffc3dd.NewBookingCurrenciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// BookingCurrenciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.bookingCurrencies.item collection
func (m *GraphBaseServiceClient) BookingCurrenciesById(id string) *id4f763cc022952ed4430efa30c2908af332dea6f57f936d1cf7ef11fa10b8be8.BookingCurrencyItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["bookingCurrency%2Did"] = id
	}
	return id4f763cc022952ed4430efa30c2908af332dea6f57f936d1cf7ef11fa10b8be8.NewBookingCurrencyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Branding the branding property
func (m *GraphBaseServiceClient) Branding() *icbfa8075dc5e1ad04a3bc48d231d9b422f250ca1e2b74477a880e0db7b0e7f95.BrandingRequestBuilder {
	return icbfa8075dc5e1ad04a3bc48d231d9b422f250ca1e2b74477a880e0db7b0e7f95.NewBrandingRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// BusinessFlowTemplates the businessFlowTemplates property
func (m *GraphBaseServiceClient) BusinessFlowTemplates() *i97c9750160852aa25d52a4c6fa196b644ce728c6645ca520427ff4d85c76afa0.BusinessFlowTemplatesRequestBuilder {
	return i97c9750160852aa25d52a4c6fa196b644ce728c6645ca520427ff4d85c76afa0.NewBusinessFlowTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// BusinessFlowTemplatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.businessFlowTemplates.item collection
func (m *GraphBaseServiceClient) BusinessFlowTemplatesById(id string) *iced651caca0254342ce35ea2308d19515b0c471b819960bd801a22997ef6f386.BusinessFlowTemplateItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["businessFlowTemplate%2Did"] = id
	}
	return iced651caca0254342ce35ea2308d19515b0c471b819960bd801a22997ef6f386.NewBusinessFlowTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// CertificateBasedAuthConfiguration the certificateBasedAuthConfiguration property
func (m *GraphBaseServiceClient) CertificateBasedAuthConfiguration() *id53bdaa191b823f3e2f4009f4cc095b46d1c7a433bde3b6d09ef0bd8df3514c2.CertificateBasedAuthConfigurationRequestBuilder {
	return id53bdaa191b823f3e2f4009f4cc095b46d1c7a433bde3b6d09ef0bd8df3514c2.NewCertificateBasedAuthConfigurationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// CertificateBasedAuthConfigurationById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.certificateBasedAuthConfiguration.item collection
func (m *GraphBaseServiceClient) CertificateBasedAuthConfigurationById(id string) *id14b905eaea3ef6f00efe71e29906d015dab23238eb870ee09debe6eab8eea9b.CertificateBasedAuthConfigurationItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["certificateBasedAuthConfiguration%2Did"] = id
	}
	return id14b905eaea3ef6f00efe71e29906d015dab23238eb870ee09debe6eab8eea9b.NewCertificateBasedAuthConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Chats the chats property
func (m *GraphBaseServiceClient) Chats() *i2bd7b88e5d2b5a20231449a09cd2703774075c53c19dc28ca3829e91d51dfd73.ChatsRequestBuilder {
	return i2bd7b88e5d2b5a20231449a09cd2703774075c53c19dc28ca3829e91d51dfd73.NewChatsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// ChatsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.chats.item collection
func (m *GraphBaseServiceClient) ChatsById(id string) *ifc0fc132a789c973aa4786767e2f987d59e214f4abef3a4e82a606459d4b00c2.ChatItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["chat%2Did"] = id
	}
	return ifc0fc132a789c973aa4786767e2f987d59e214f4abef3a4e82a606459d4b00c2.NewChatItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Commands the commands property
func (m *GraphBaseServiceClient) Commands() *i81513db44a7796569414887f1110e1d158078cb199b3565d960425af9c2904ba.CommandsRequestBuilder {
	return i81513db44a7796569414887f1110e1d158078cb199b3565d960425af9c2904ba.NewCommandsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// CommandsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.commands.item collection
func (m *GraphBaseServiceClient) CommandsById(id string) *ie39c059add8b730dc85b7d2449e4b26b98f9a83d3807fabe7cb9009271179970.CommandItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["command%2Did"] = id
	}
	return ie39c059add8b730dc85b7d2449e4b26b98f9a83d3807fabe7cb9009271179970.NewCommandItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Communications the communications property
func (m *GraphBaseServiceClient) Communications() *if7bcb57951e8f2ae550fcf781dc209ed777854429fcfe2465a71b03112dfc346.CommunicationsRequestBuilder {
	return if7bcb57951e8f2ae550fcf781dc209ed777854429fcfe2465a71b03112dfc346.NewCommunicationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Compliance the compliance property
func (m *GraphBaseServiceClient) Compliance() *ia93a0d553904ea998c5c0e30cf6d6574ba3a42be402a2d1eb88d5bd221d1a0d6.ComplianceRequestBuilder {
	return ia93a0d553904ea998c5c0e30cf6d6574ba3a42be402a2d1eb88d5bd221d1a0d6.NewComplianceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Connections the connections property
func (m *GraphBaseServiceClient) Connections() *i29b3182886b8fc3d309db2628f3911e671f6745e9fdab71d07cdeb487c4453ad.ConnectionsRequestBuilder {
	return i29b3182886b8fc3d309db2628f3911e671f6745e9fdab71d07cdeb487c4453ad.NewConnectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// ConnectionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.connections.item collection
func (m *GraphBaseServiceClient) ConnectionsById(id string) *ie9b2b2a73483057a6575063ce71086fbc54437022d4c53761e7a4a058d7dd091.ExternalConnectionItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["externalConnection%2Did"] = id
	}
	return ie9b2b2a73483057a6575063ce71086fbc54437022d4c53761e7a4a058d7dd091.NewExternalConnectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// NewBaseGraphServiceClient instantiates a new GraphBaseServiceClient and sets the default values.
func NewBaseGraphServiceClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *GraphBaseServiceClient {
	m := &GraphBaseServiceClient{}
	m.pathParameters = make(map[string]string)
	m.urlTemplate = "{+baseurl}"
	m.requestAdapter = requestAdapter
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory {
		return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory()
	})
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory {
		return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory()
	})
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory {
		return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory()
	})
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory {
		return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory()
	})
	if m.requestAdapter.GetBaseUrl() == "" {
		m.requestAdapter.SetBaseUrl("https://graph.microsoft.com/beta")
	}
	return m
}

// Contacts the contacts property
func (m *GraphBaseServiceClient) Contacts() *i716e3204a4c47d24737c05f3b4c2ef2462fa5a1df29b57365f338e8f68ee16ef.ContactsRequestBuilder {
	return i716e3204a4c47d24737c05f3b4c2ef2462fa5a1df29b57365f338e8f68ee16ef.NewContactsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// ContactsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.contacts.item collection
func (m *GraphBaseServiceClient) ContactsById(id string) *i6e94fa7602e5cf30291b967fc83a01373300f1d5183cc458d161d85fa3dc2baf.OrgContactItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["orgContact%2Did"] = id
	}
	return i6e94fa7602e5cf30291b967fc83a01373300f1d5183cc458d161d85fa3dc2baf.NewOrgContactItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Contracts the contracts property
func (m *GraphBaseServiceClient) Contracts() *i5e88b3e8025e26d21777096f2c7525a182545bb8dc0137634b047a10fe14ab54.ContractsRequestBuilder {
	return i5e88b3e8025e26d21777096f2c7525a182545bb8dc0137634b047a10fe14ab54.NewContractsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// ContractsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.contracts.item collection
func (m *GraphBaseServiceClient) ContractsById(id string) *i09e36d06bea20a5d152cc52ffa7ef3c082c8cdfbf6d2116b226488cbb8f849ed.ContractItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["contract%2Did"] = id
	}
	return i09e36d06bea20a5d152cc52ffa7ef3c082c8cdfbf6d2116b226488cbb8f849ed.NewContractItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// DataClassification the dataClassification property
func (m *GraphBaseServiceClient) DataClassification() *if3c2dea1db099d8f9ce8b623a12f6291f276e2bbb50259658f584a3a85cf71b8.DataClassificationRequestBuilder {
	return if3c2dea1db099d8f9ce8b623a12f6291f276e2bbb50259658f584a3a85cf71b8.NewDataClassificationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// DataPolicyOperations the dataPolicyOperations property
func (m *GraphBaseServiceClient) DataPolicyOperations() *i0f747ff1f24810ff51160697ed4229c9ca192f7b84644311b88fa3b475cc340d.DataPolicyOperationsRequestBuilder {
	return i0f747ff1f24810ff51160697ed4229c9ca192f7b84644311b88fa3b475cc340d.NewDataPolicyOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// DataPolicyOperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.dataPolicyOperations.item collection
func (m *GraphBaseServiceClient) DataPolicyOperationsById(id string) *i5cb8828c84548ec02b63ac985f24d0e6038f8ba987f6d69cc0a62c4e8c7d7df4.DataPolicyOperationItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["dataPolicyOperation%2Did"] = id
	}
	return i5cb8828c84548ec02b63ac985f24d0e6038f8ba987f6d69cc0a62c4e8c7d7df4.NewDataPolicyOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// DeviceAppManagement the deviceAppManagement property
func (m *GraphBaseServiceClient) DeviceAppManagement() *idc4afe653def183ef95500aa004f556fdf3c3747771f17c8472ca3cad61cebf4.DeviceAppManagementRequestBuilder {
	return idc4afe653def183ef95500aa004f556fdf3c3747771f17c8472ca3cad61cebf4.NewDeviceAppManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// DeviceManagement the deviceManagement property
func (m *GraphBaseServiceClient) DeviceManagement() *i09893664b20e7c846b2bc7aaaf1cd7f554ed3d2c00ac11336bea4c3c3d859e09.DeviceManagementRequestBuilder {
	return i09893664b20e7c846b2bc7aaaf1cd7f554ed3d2c00ac11336bea4c3c3d859e09.NewDeviceManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Devices the devices property
func (m *GraphBaseServiceClient) Devices() *i0b4892b2f92a31e44541567b8065e8e7760cb336e17d7dacb9120a865d5b0a37.DevicesRequestBuilder {
	return i0b4892b2f92a31e44541567b8065e8e7760cb336e17d7dacb9120a865d5b0a37.NewDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// DevicesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.devices.item collection
func (m *GraphBaseServiceClient) DevicesById(id string) *i3629807faa94b36aa56a99f728c28616ef4be3a1816bf7413ce360f3f6ef1c9d.DeviceItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["device%2Did"] = id
	}
	return i3629807faa94b36aa56a99f728c28616ef4be3a1816bf7413ce360f3f6ef1c9d.NewDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Directory the directory property
func (m *GraphBaseServiceClient) Directory() *i6e398703c86ec3814400d80161079e7253c4e25f4ba1adb0c8d31da236f7bcd7.DirectoryRequestBuilder {
	return i6e398703c86ec3814400d80161079e7253c4e25f4ba1adb0c8d31da236f7bcd7.NewDirectoryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// DirectoryObjects the directoryObjects property
func (m *GraphBaseServiceClient) DirectoryObjects() *iaec68a3d2c3ba0a78ebb66cd93fd1c5d2a6e0450b97a0cf19d94cb58956bec1d.DirectoryObjectsRequestBuilder {
	return iaec68a3d2c3ba0a78ebb66cd93fd1c5d2a6e0450b97a0cf19d94cb58956bec1d.NewDirectoryObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// DirectoryObjectsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directoryObjects.item collection
func (m *GraphBaseServiceClient) DirectoryObjectsById(id string) *i901e1ca524c6a851c60abc23bfdb47050e754ebcf438922b25817ff66d56276a.DirectoryObjectItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["directoryObject%2Did"] = id
	}
	return i901e1ca524c6a851c60abc23bfdb47050e754ebcf438922b25817ff66d56276a.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// DirectoryRoles the directoryRoles property
func (m *GraphBaseServiceClient) DirectoryRoles() *i20621ebd49d2bb1ed6c592ae35dfa701db30564a91ff100d25b0dcdb142bd942.DirectoryRolesRequestBuilder {
	return i20621ebd49d2bb1ed6c592ae35dfa701db30564a91ff100d25b0dcdb142bd942.NewDirectoryRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// DirectoryRolesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directoryRoles.item collection
func (m *GraphBaseServiceClient) DirectoryRolesById(id string) *i993beddae5d86c1977387e8b393295fbee6f591297fb6a91fc3501c6cb5ae0d8.DirectoryRoleItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["directoryRole%2Did"] = id
	}
	return i993beddae5d86c1977387e8b393295fbee6f591297fb6a91fc3501c6cb5ae0d8.NewDirectoryRoleItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// DirectoryRoleTemplates the directoryRoleTemplates property
func (m *GraphBaseServiceClient) DirectoryRoleTemplates() *i5b4eb770497618728398e41e6ed415ad2b92d20f7ad45ba75277a5800d9a2a12.DirectoryRoleTemplatesRequestBuilder {
	return i5b4eb770497618728398e41e6ed415ad2b92d20f7ad45ba75277a5800d9a2a12.NewDirectoryRoleTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// DirectoryRoleTemplatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directoryRoleTemplates.item collection
func (m *GraphBaseServiceClient) DirectoryRoleTemplatesById(id string) *i98e280807a93b84e9092eeb27b7840b536491d6fa3c9de876973f203933ecf7f.DirectoryRoleTemplateItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["directoryRoleTemplate%2Did"] = id
	}
	return i98e280807a93b84e9092eeb27b7840b536491d6fa3c9de876973f203933ecf7f.NewDirectoryRoleTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// DirectorySettingTemplates the directorySettingTemplates property
func (m *GraphBaseServiceClient) DirectorySettingTemplates() *iabd30d9ba6ae302fd2d3145e0da70036496a2167b9e0cbc049bab96d9d9a29b3.DirectorySettingTemplatesRequestBuilder {
	return iabd30d9ba6ae302fd2d3145e0da70036496a2167b9e0cbc049bab96d9d9a29b3.NewDirectorySettingTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// DirectorySettingTemplatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directorySettingTemplates.item collection
func (m *GraphBaseServiceClient) DirectorySettingTemplatesById(id string) *i942fb2a58a06777c286167881e19f5c2b00590a6cc7831e7a502f112db658679.DirectorySettingTemplateItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["directorySettingTemplate%2Did"] = id
	}
	return i942fb2a58a06777c286167881e19f5c2b00590a6cc7831e7a502f112db658679.NewDirectorySettingTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// DomainDnsRecords the domainDnsRecords property
func (m *GraphBaseServiceClient) DomainDnsRecords() *i5db103717157eeba829a6a8580002d8b9adfdfd549628f3c84152bf3164a7b53.DomainDnsRecordsRequestBuilder {
	return i5db103717157eeba829a6a8580002d8b9adfdfd549628f3c84152bf3164a7b53.NewDomainDnsRecordsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// DomainDnsRecordsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.domainDnsRecords.item collection
func (m *GraphBaseServiceClient) DomainDnsRecordsById(id string) *if2c2cfdc1fc8e4c78fc654fd9ff30fce9a231a7ccf8c16c510d1dbb98e3aa7ab.DomainDnsRecordItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["domainDnsRecord%2Did"] = id
	}
	return if2c2cfdc1fc8e4c78fc654fd9ff30fce9a231a7ccf8c16c510d1dbb98e3aa7ab.NewDomainDnsRecordItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Domains the domains property
func (m *GraphBaseServiceClient) Domains() *i9e22e53c888822daa9264a72be4d11f335e3170e9198aae4bba758214e319857.DomainsRequestBuilder {
	return i9e22e53c888822daa9264a72be4d11f335e3170e9198aae4bba758214e319857.NewDomainsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// DomainsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.domains.item collection
func (m *GraphBaseServiceClient) DomainsById(id string) *ia82ac058caa6740adb23fd487affca202963b9786c4890c80eaec27b6d7ced24.DomainItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["domain%2Did"] = id
	}
	return ia82ac058caa6740adb23fd487affca202963b9786c4890c80eaec27b6d7ced24.NewDomainItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Drive the drive property
func (m *GraphBaseServiceClient) Drive() *idc403d989fa3b67787cd15ceedea7e45400694aa996d133f9bf61a54fe5a497a.DriveRequestBuilder {
	return idc403d989fa3b67787cd15ceedea7e45400694aa996d133f9bf61a54fe5a497a.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Drives the drives property
func (m *GraphBaseServiceClient) Drives() *i2b9cd6123f9a7ca2d7c253973e81bb3965ee0b78a350919aafe66d7b73c70433.DrivesRequestBuilder {
	return i2b9cd6123f9a7ca2d7c253973e81bb3965ee0b78a350919aafe66d7b73c70433.NewDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// DrivesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item collection
func (m *GraphBaseServiceClient) DrivesById(id string) *ib8b674690e8be27a974ea55202f0d976b3d187306aca08f5642097d6cfed398b.DriveItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["drive%2Did"] = id
	}
	return ib8b674690e8be27a974ea55202f0d976b3d187306aca08f5642097d6cfed398b.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Education the education property
func (m *GraphBaseServiceClient) Education() *id90d135edac1f1a3e952db4ad985001105d2e7c0133f8cc410765eb1af789cc0.EducationRequestBuilder {
	return id90d135edac1f1a3e952db4ad985001105d2e7c0133f8cc410765eb1af789cc0.NewEducationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// EmployeeExperience the employeeExperience property
func (m *GraphBaseServiceClient) EmployeeExperience() *i4b923a0ffb143d5def24980fce5d55c68c9634d5b55c33bb0b3029ac68415dd1.EmployeeExperienceRequestBuilder {
	return i4b923a0ffb143d5def24980fce5d55c68c9634d5b55c33bb0b3029ac68415dd1.NewEmployeeExperienceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// External the external property
func (m *GraphBaseServiceClient) External() *i453454978d87486fd201df62ea4b775c5b4907e2a36395fb6fb1e9060fc3f1bb.ExternalRequestBuilder {
	return i453454978d87486fd201df62ea4b775c5b4907e2a36395fb6fb1e9060fc3f1bb.NewExternalRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// FilterOperators the filterOperators property
func (m *GraphBaseServiceClient) FilterOperators() *i0b9d70018d3c267f9f34a818ce43cc889d06d87749a70e1ad1d45eead0c735e0.FilterOperatorsRequestBuilder {
	return i0b9d70018d3c267f9f34a818ce43cc889d06d87749a70e1ad1d45eead0c735e0.NewFilterOperatorsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// FilterOperatorsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.filterOperators.item collection
func (m *GraphBaseServiceClient) FilterOperatorsById(id string) *i330b70aeb45059ed2410e099c3692c88ec3d2474fbb8363fffbecba6f900c494.FilterOperatorSchemaItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["filterOperatorSchema%2Did"] = id
	}
	return i330b70aeb45059ed2410e099c3692c88ec3d2474fbb8363fffbecba6f900c494.NewFilterOperatorSchemaItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Financials the financials property
func (m *GraphBaseServiceClient) Financials() *i0281e791d6eedd5b4b6d3c18b07336722b6152b5db6a2e9dc8385e98f565f677.FinancialsRequestBuilder {
	return i0281e791d6eedd5b4b6d3c18b07336722b6152b5db6a2e9dc8385e98f565f677.NewFinancialsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Functions the functions property
func (m *GraphBaseServiceClient) Functions() *i054b68521cee54ec767d07cf7a8a6d50a0d24b6e6fc43b8296c34730fd8ca465.FunctionsRequestBuilder {
	return i054b68521cee54ec767d07cf7a8a6d50a0d24b6e6fc43b8296c34730fd8ca465.NewFunctionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// FunctionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.functions.item collection
func (m *GraphBaseServiceClient) FunctionsById(id string) *ib3f85f7711d9ea4057703e0171ba196ac1d884a1e2389602111d38d84fb1380f.AttributeMappingFunctionSchemaItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["attributeMappingFunctionSchema%2Did"] = id
	}
	return ib3f85f7711d9ea4057703e0171ba196ac1d884a1e2389602111d38d84fb1380f.NewAttributeMappingFunctionSchemaItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// GovernanceResources the governanceResources property
func (m *GraphBaseServiceClient) GovernanceResources() *i3b7da1b693d5428b20b0bf3340acb4b879042a9393e45df9349b04a5b2830acb.GovernanceResourcesRequestBuilder {
	return i3b7da1b693d5428b20b0bf3340acb4b879042a9393e45df9349b04a5b2830acb.NewGovernanceResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// GovernanceResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.governanceResources.item collection
func (m *GraphBaseServiceClient) GovernanceResourcesById(id string) *iaeceabcb9ce892827c1e040d7ed5dec4e2fb051c4a1cb3b857ac26d621e94638.GovernanceResourceItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["governanceResource%2Did"] = id
	}
	return iaeceabcb9ce892827c1e040d7ed5dec4e2fb051c4a1cb3b857ac26d621e94638.NewGovernanceResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// GovernanceRoleAssignmentRequests the governanceRoleAssignmentRequests property
func (m *GraphBaseServiceClient) GovernanceRoleAssignmentRequests() *i3de07fd1833246d183a4d60c2329c7467381716c2cfdfc096ff627e4e9479cf8.GovernanceRoleAssignmentRequestsRequestBuilder {
	return i3de07fd1833246d183a4d60c2329c7467381716c2cfdfc096ff627e4e9479cf8.NewGovernanceRoleAssignmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// GovernanceRoleAssignmentRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.governanceRoleAssignmentRequests.item collection
func (m *GraphBaseServiceClient) GovernanceRoleAssignmentRequestsById(id string) *if03454ccae6aab82c8d7c01d50047eceb6218bf085d9949f9f4ff37652bab5a0.GovernanceRoleAssignmentRequestItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["governanceRoleAssignmentRequest%2Did"] = id
	}
	return if03454ccae6aab82c8d7c01d50047eceb6218bf085d9949f9f4ff37652bab5a0.NewGovernanceRoleAssignmentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// GovernanceRoleAssignments the governanceRoleAssignments property
func (m *GraphBaseServiceClient) GovernanceRoleAssignments() *i14752cfec59ab915e7c63922270765abf65744437d9135c191cef3986f08c3bb.GovernanceRoleAssignmentsRequestBuilder {
	return i14752cfec59ab915e7c63922270765abf65744437d9135c191cef3986f08c3bb.NewGovernanceRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// GovernanceRoleAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.governanceRoleAssignments.item collection
func (m *GraphBaseServiceClient) GovernanceRoleAssignmentsById(id string) *i5e0d644f8c8abdc66bf30911d1b935134f92408badf2ef700cf692be1de56856.GovernanceRoleAssignmentItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["governanceRoleAssignment%2Did"] = id
	}
	return i5e0d644f8c8abdc66bf30911d1b935134f92408badf2ef700cf692be1de56856.NewGovernanceRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// GovernanceRoleDefinitions the governanceRoleDefinitions property
func (m *GraphBaseServiceClient) GovernanceRoleDefinitions() *i8a18cb7418541221b2c3fd213a484d9e3029fab916358b16fb24015b078b8eba.GovernanceRoleDefinitionsRequestBuilder {
	return i8a18cb7418541221b2c3fd213a484d9e3029fab916358b16fb24015b078b8eba.NewGovernanceRoleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// GovernanceRoleDefinitionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.governanceRoleDefinitions.item collection
func (m *GraphBaseServiceClient) GovernanceRoleDefinitionsById(id string) *i005472295f05513a4d4bcabd67693926f332f4d4740482ce99613ae21d4ea87a.GovernanceRoleDefinitionItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["governanceRoleDefinition%2Did"] = id
	}
	return i005472295f05513a4d4bcabd67693926f332f4d4740482ce99613ae21d4ea87a.NewGovernanceRoleDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// GovernanceRoleSettings the governanceRoleSettings property
func (m *GraphBaseServiceClient) GovernanceRoleSettings() *ib71e32ed3a7f0f8a512aa55c1428492116ff2d1bae5015a9b89f910ecbc7c6bd.GovernanceRoleSettingsRequestBuilder {
	return ib71e32ed3a7f0f8a512aa55c1428492116ff2d1bae5015a9b89f910ecbc7c6bd.NewGovernanceRoleSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// GovernanceRoleSettingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.governanceRoleSettings.item collection
func (m *GraphBaseServiceClient) GovernanceRoleSettingsById(id string) *i2a5979c2bdbc99c804ec546b7cec503201936edc913e5d36d239a54c3ad002ad.GovernanceRoleSettingItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["governanceRoleSetting%2Did"] = id
	}
	return i2a5979c2bdbc99c804ec546b7cec503201936edc913e5d36d239a54c3ad002ad.NewGovernanceRoleSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// GovernanceSubjects the governanceSubjects property
func (m *GraphBaseServiceClient) GovernanceSubjects() *ifcac309012d761a79a74e6d79fad6979f2117e7af36ff6e5ad131093412afcc7.GovernanceSubjectsRequestBuilder {
	return ifcac309012d761a79a74e6d79fad6979f2117e7af36ff6e5ad131093412afcc7.NewGovernanceSubjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// GovernanceSubjectsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.governanceSubjects.item collection
func (m *GraphBaseServiceClient) GovernanceSubjectsById(id string) *i6f49755cbc32687ce5fd8e0bc2cdb1808c45f60ba20889d053bac06929b7d9ea.GovernanceSubjectItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["governanceSubject%2Did"] = id
	}
	return i6f49755cbc32687ce5fd8e0bc2cdb1808c45f60ba20889d053bac06929b7d9ea.NewGovernanceSubjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// GroupLifecyclePolicies the groupLifecyclePolicies property
func (m *GraphBaseServiceClient) GroupLifecyclePolicies() *i71438b4a3f9d4a17f8c873a44b8ac76600403f5ce0cce2423bde35e0191f2c17.GroupLifecyclePoliciesRequestBuilder {
	return i71438b4a3f9d4a17f8c873a44b8ac76600403f5ce0cce2423bde35e0191f2c17.NewGroupLifecyclePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// GroupLifecyclePoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groupLifecyclePolicies.item collection
func (m *GraphBaseServiceClient) GroupLifecyclePoliciesById(id string) *if2d4db22e4ec9d990102702a0f1b228fbfa6a90e349c6dfb2f3c798d2c5ec77f.GroupLifecyclePolicyItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["groupLifecyclePolicy%2Did"] = id
	}
	return if2d4db22e4ec9d990102702a0f1b228fbfa6a90e349c6dfb2f3c798d2c5ec77f.NewGroupLifecyclePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Groups the groups property
func (m *GraphBaseServiceClient) Groups() *i79f2b866e8bec3ee9349dc885ecb3691e94b20459995d83b3dbf9f05341c7a89.GroupsRequestBuilder {
	return i79f2b866e8bec3ee9349dc885ecb3691e94b20459995d83b3dbf9f05341c7a89.NewGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// GroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item collection
func (m *GraphBaseServiceClient) GroupsById(id string) *ic6c550f1e18a80bf95a0e1377fab585951c941746c2a2739d2b1d77b17b80a94.GroupItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["group%2Did"] = id
	}
	return ic6c550f1e18a80bf95a0e1377fab585951c941746c2a2739d2b1d77b17b80a94.NewGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Identity the identity property
func (m *GraphBaseServiceClient) Identity() *i20702653f98186060bd39b9fe8136743eafc0ddaa43435e527665ac75229a33a.IdentityRequestBuilder {
	return i20702653f98186060bd39b9fe8136743eafc0ddaa43435e527665ac75229a33a.NewIdentityRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// IdentityGovernance the identityGovernance property
func (m *GraphBaseServiceClient) IdentityGovernance() *i7ed0a0c7cc963cb1cf6c282e5e5c04cea2d4959e5f1bfd12698c0196858b1b52.IdentityGovernanceRequestBuilder {
	return i7ed0a0c7cc963cb1cf6c282e5e5c04cea2d4959e5f1bfd12698c0196858b1b52.NewIdentityGovernanceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// IdentityProtection the identityProtection property
func (m *GraphBaseServiceClient) IdentityProtection() *i444aa080f7aab1ac35ae55318eca861a558a57b4c5e3cbd071222f511baac5fe.IdentityProtectionRequestBuilder {
	return i444aa080f7aab1ac35ae55318eca861a558a57b4c5e3cbd071222f511baac5fe.NewIdentityProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// IdentityProviders the identityProviders property
func (m *GraphBaseServiceClient) IdentityProviders() *i531b1efd1768fd272d51921ff5812bdeba5b46e0eeec0e4c818250cb7116aed5.IdentityProvidersRequestBuilder {
	return i531b1efd1768fd272d51921ff5812bdeba5b46e0eeec0e4c818250cb7116aed5.NewIdentityProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// IdentityProvidersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityProviders.item collection
func (m *GraphBaseServiceClient) IdentityProvidersById(id string) *ib66730b67f84ecbb45dd98e1ce34bd456667b7437238322804d692193cbb7cd6.IdentityProviderItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["identityProvider%2Did"] = id
	}
	return ib66730b67f84ecbb45dd98e1ce34bd456667b7437238322804d692193cbb7cd6.NewIdentityProviderItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// InformationProtection the informationProtection property
func (m *GraphBaseServiceClient) InformationProtection() *ida1e35de3db89946d53da45357ba0e3da7b3e3a1d46b511191582f0d2a917caa.InformationProtectionRequestBuilder {
	return ida1e35de3db89946d53da45357ba0e3da7b3e3a1d46b511191582f0d2a917caa.NewInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Invitations the invitations property
func (m *GraphBaseServiceClient) Invitations() *i4cb6edb865a0e38bb1799dcb0c7881b92feed59596c1912cfe5e6142b61f9c91.InvitationsRequestBuilder {
	return i4cb6edb865a0e38bb1799dcb0c7881b92feed59596c1912cfe5e6142b61f9c91.NewInvitationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// InvitationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.invitations.item collection
func (m *GraphBaseServiceClient) InvitationsById(id string) *i5cfcfa8bc4a6d5fed3fa8ff67c53050a2b9493ff4ba9481077f03d4df2e84db6.InvitationItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["invitation%2Did"] = id
	}
	return i5cfcfa8bc4a6d5fed3fa8ff67c53050a2b9493ff4ba9481077f03d4df2e84db6.NewInvitationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Me the me property
func (m *GraphBaseServiceClient) Me() *i5a50ec7cbb7e52feaacf8d45f5e247e765502c09622e2fa44cd68445eab876eb.MeRequestBuilder {
	return i5a50ec7cbb7e52feaacf8d45f5e247e765502c09622e2fa44cd68445eab876eb.NewMeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// MessageEvents the messageEvents property
func (m *GraphBaseServiceClient) MessageEvents() *i19717748912ff29c95998304f534371c35864a4579ea92608b32aeecf7d18cc4.MessageEventsRequestBuilder {
	return i19717748912ff29c95998304f534371c35864a4579ea92608b32aeecf7d18cc4.NewMessageEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// MessageEventsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.messageEvents.item collection
func (m *GraphBaseServiceClient) MessageEventsById(id string) *i2737b8a250008797e3665bb01cdf0bfbf42bedbb23c7d023eadf66c5dc36b26e.MessageEventItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["messageEvent%2Did"] = id
	}
	return i2737b8a250008797e3665bb01cdf0bfbf42bedbb23c7d023eadf66c5dc36b26e.NewMessageEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// MessageRecipients the messageRecipients property
func (m *GraphBaseServiceClient) MessageRecipients() *i56ad7deac03a612015589ab4fade2313d6df08086c7ee8d46177fc8ddc5b0053.MessageRecipientsRequestBuilder {
	return i56ad7deac03a612015589ab4fade2313d6df08086c7ee8d46177fc8ddc5b0053.NewMessageRecipientsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// MessageRecipientsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.messageRecipients.item collection
func (m *GraphBaseServiceClient) MessageRecipientsById(id string) *ia62d215a2c5a4beae9bb26027fb9b0f0438b6e34a48aa59b64ddddde97e1e648.MessageRecipientItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["messageRecipient%2Did"] = id
	}
	return ia62d215a2c5a4beae9bb26027fb9b0f0438b6e34a48aa59b64ddddde97e1e648.NewMessageRecipientItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// MessageTraces the messageTraces property
func (m *GraphBaseServiceClient) MessageTraces() *i43a1701c7f8902247fd60e3d6ff36be02d1a59f02884ce5734adb8deb69a47ef.MessageTracesRequestBuilder {
	return i43a1701c7f8902247fd60e3d6ff36be02d1a59f02884ce5734adb8deb69a47ef.NewMessageTracesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// MessageTracesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.messageTraces.item collection
func (m *GraphBaseServiceClient) MessageTracesById(id string) *iaa37a57ee21d5046d5b74ae373a2d314357f709507fd2fb2116379f39879b88a.MessageTraceItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["messageTrace%2Did"] = id
	}
	return iaa37a57ee21d5046d5b74ae373a2d314357f709507fd2fb2116379f39879b88a.NewMessageTraceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// MobilityManagementPolicies the mobilityManagementPolicies property
func (m *GraphBaseServiceClient) MobilityManagementPolicies() *i05bd1def68419ff4c77b9bdafb42d256c4c19b479003ef8adcd868ec38673e84.MobilityManagementPoliciesRequestBuilder {
	return i05bd1def68419ff4c77b9bdafb42d256c4c19b479003ef8adcd868ec38673e84.NewMobilityManagementPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// MobilityManagementPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.mobilityManagementPolicies.item collection
func (m *GraphBaseServiceClient) MobilityManagementPoliciesById(id string) *i0740547a5aa3dd54a8b262a3d1cd5e0bb409b1a645795e041ffa10dcb9019e0a.MobilityManagementPolicyItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["mobilityManagementPolicy%2Did"] = id
	}
	return i0740547a5aa3dd54a8b262a3d1cd5e0bb409b1a645795e041ffa10dcb9019e0a.NewMobilityManagementPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Monitoring the monitoring property
func (m *GraphBaseServiceClient) Monitoring() *i3c2d5b7a05b0c58ca1d3b72343c5f03221f2bd894e75f89f62ccf8020a48250b.MonitoringRequestBuilder {
	return i3c2d5b7a05b0c58ca1d3b72343c5f03221f2bd894e75f89f62ccf8020a48250b.NewMonitoringRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Oauth2PermissionGrants the oauth2PermissionGrants property
func (m *GraphBaseServiceClient) Oauth2PermissionGrants() *ie2e0818e93fcfbb33fde071a9354c3c22bedab0ec20855b7d5232d29bcc65bad.Oauth2PermissionGrantsRequestBuilder {
	return ie2e0818e93fcfbb33fde071a9354c3c22bedab0ec20855b7d5232d29bcc65bad.NewOauth2PermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Oauth2PermissionGrantsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.oauth2PermissionGrants.item collection
func (m *GraphBaseServiceClient) Oauth2PermissionGrantsById(id string) *if237177d98e26841b842fcf602048c102ab53f50428ff601c17ea0795d7fb828.OAuth2PermissionGrantItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["oAuth2PermissionGrant%2Did"] = id
	}
	return if237177d98e26841b842fcf602048c102ab53f50428ff601c17ea0795d7fb828.NewOAuth2PermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// OfficeConfiguration the officeConfiguration property
func (m *GraphBaseServiceClient) OfficeConfiguration() *ibcc2ba1bd11f45f4381ae4007e619ef47b74dbf731b922816aff1ed4d5d47a0f.OfficeConfigurationRequestBuilder {
	return ibcc2ba1bd11f45f4381ae4007e619ef47b74dbf731b922816aff1ed4d5d47a0f.NewOfficeConfigurationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// OnPremisesPublishingProfiles the onPremisesPublishingProfiles property
func (m *GraphBaseServiceClient) OnPremisesPublishingProfiles() *i0d38f6e6ea6126fff7bb7a5c3c2d82fe471d00233209e2b6b2ce6ccb21ce50f5.OnPremisesPublishingProfilesRequestBuilder {
	return i0d38f6e6ea6126fff7bb7a5c3c2d82fe471d00233209e2b6b2ce6ccb21ce50f5.NewOnPremisesPublishingProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// OnPremisesPublishingProfilesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.onPremisesPublishingProfiles.item collection
func (m *GraphBaseServiceClient) OnPremisesPublishingProfilesById(id string) *ie1c23a9b73355bb18364b81d44931ba2a9c3f409f2e091f8c44725c8075f8cf9.OnPremisesPublishingProfileItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["onPremisesPublishingProfile%2Did"] = id
	}
	return ie1c23a9b73355bb18364b81d44931ba2a9c3f409f2e091f8c44725c8075f8cf9.NewOnPremisesPublishingProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Organization the organization property
func (m *GraphBaseServiceClient) Organization() *ie58948149bb028757a64f16376df00cc5a99ad93e0d57affa0ac24ff6d096aaf.OrganizationRequestBuilder {
	return ie58948149bb028757a64f16376df00cc5a99ad93e0d57affa0ac24ff6d096aaf.NewOrganizationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// OrganizationById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.organization.item collection
func (m *GraphBaseServiceClient) OrganizationById(id string) *ic9a01566deb7ac7de88728b70eec29525d2a94adff4181ea235c30369d4c108c.OrganizationItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["organization%2Did"] = id
	}
	return ic9a01566deb7ac7de88728b70eec29525d2a94adff4181ea235c30369d4c108c.NewOrganizationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// PayloadResponse the payloadResponse property
func (m *GraphBaseServiceClient) PayloadResponse() *i22f037221f506c5b5751f13095cc17caaf248e93588f883123c452cb4f1ec6a9.PayloadResponseRequestBuilder {
	return i22f037221f506c5b5751f13095cc17caaf248e93588f883123c452cb4f1ec6a9.NewPayloadResponseRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// PayloadResponseById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.payloadResponse.item collection
func (m *GraphBaseServiceClient) PayloadResponseById(id string) *ie1b5a10a367968f2f9e397679c2e84cb9c8ec8475b94ece386d4dcd129bcefe8.PayloadResponseItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["payloadResponse%2Did"] = id
	}
	return ie1b5a10a367968f2f9e397679c2e84cb9c8ec8475b94ece386d4dcd129bcefe8.NewPayloadResponseItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// PermissionGrants the permissionGrants property
func (m *GraphBaseServiceClient) PermissionGrants() *i48a68a7c83dc874f9d9fdf942afed70a34b11f92d6b2ccb439359753116f65cc.PermissionGrantsRequestBuilder {
	return i48a68a7c83dc874f9d9fdf942afed70a34b11f92d6b2ccb439359753116f65cc.NewPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// PermissionGrantsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.permissionGrants.item collection
func (m *GraphBaseServiceClient) PermissionGrantsById(id string) *i79b93b64b0d0d7fc163598b550442694212cae76f99c0fc0cc0b570453d3decb.ResourceSpecificPermissionGrantItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["resourceSpecificPermissionGrant%2Did"] = id
	}
	return i79b93b64b0d0d7fc163598b550442694212cae76f99c0fc0cc0b570453d3decb.NewResourceSpecificPermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Places the places property
func (m *GraphBaseServiceClient) Places() *ic21cf429efd6fc3199e67b5b4288a3193ff5e9cfb4e97a5e442e02ccd7748ec1.PlacesRequestBuilder {
	return ic21cf429efd6fc3199e67b5b4288a3193ff5e9cfb4e97a5e442e02ccd7748ec1.NewPlacesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// PlacesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.places.item collection
func (m *GraphBaseServiceClient) PlacesById(id string) *ia6735253ec4d9448b4b9541b92358386ce1936ac15c0298ba04aaff0a1e74e3d.PlaceItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["place%2Did"] = id
	}
	return ia6735253ec4d9448b4b9541b92358386ce1936ac15c0298ba04aaff0a1e74e3d.NewPlaceItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Planner the planner property
func (m *GraphBaseServiceClient) Planner() *i2c16e7aacab2db540fbc4e11916f266c9c46936118144cc1383e798af27b1b6f.PlannerRequestBuilder {
	return i2c16e7aacab2db540fbc4e11916f266c9c46936118144cc1383e798af27b1b6f.NewPlannerRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Policies the policies property
func (m *GraphBaseServiceClient) Policies() *i5f794bd004f1cc95da776bcb1947ffabf97b71aae1f5c9f15255b24451e2929b.PoliciesRequestBuilder {
	return i5f794bd004f1cc95da776bcb1947ffabf97b71aae1f5c9f15255b24451e2929b.NewPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Print the print property
func (m *GraphBaseServiceClient) Print() *i75b6dc07087cda1a9afc465878b0aa56ca3703a3ed530d5a22119b0960d159d3.PrintRequestBuilder {
	return i75b6dc07087cda1a9afc465878b0aa56ca3703a3ed530d5a22119b0960d159d3.NewPrintRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Privacy the privacy property
func (m *GraphBaseServiceClient) Privacy() *i7daddf90c53ab17a62cc6612f32636f5bea354e4e22c20c5656606d1b491bd03.PrivacyRequestBuilder {
	return i7daddf90c53ab17a62cc6612f32636f5bea354e4e22c20c5656606d1b491bd03.NewPrivacyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// PrivilegedAccess the privilegedAccess property
func (m *GraphBaseServiceClient) PrivilegedAccess() *i2f87335be6a07866636e1f602f5beda6b47bc99969e216fac59efff432ff2345.PrivilegedAccessRequestBuilder {
	return i2f87335be6a07866636e1f602f5beda6b47bc99969e216fac59efff432ff2345.NewPrivilegedAccessRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// PrivilegedAccessById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.privilegedAccess.item collection
func (m *GraphBaseServiceClient) PrivilegedAccessById(id string) *ibc6cc10ee3264e7addfcd3086a9a46a8a89142b950c482b08cde11f41f711bc2.PrivilegedAccessItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["privilegedAccess%2Did"] = id
	}
	return ibc6cc10ee3264e7addfcd3086a9a46a8a89142b950c482b08cde11f41f711bc2.NewPrivilegedAccessItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// PrivilegedApproval the privilegedApproval property
func (m *GraphBaseServiceClient) PrivilegedApproval() *i613e9d4b5596faf0dba59087b1a65d06d17aab1ca9d69170b292a4e0d90063ab.PrivilegedApprovalRequestBuilder {
	return i613e9d4b5596faf0dba59087b1a65d06d17aab1ca9d69170b292a4e0d90063ab.NewPrivilegedApprovalRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// PrivilegedApprovalById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.privilegedApproval.item collection
func (m *GraphBaseServiceClient) PrivilegedApprovalById(id string) *i3056c9ce4b136c43e7e58e79400328464aedd7296804666295056c525cd5afe9.PrivilegedApprovalItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["privilegedApproval%2Did"] = id
	}
	return i3056c9ce4b136c43e7e58e79400328464aedd7296804666295056c525cd5afe9.NewPrivilegedApprovalItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// PrivilegedOperationEvents the privilegedOperationEvents property
func (m *GraphBaseServiceClient) PrivilegedOperationEvents() *i82672a497b8f66c510d59f2a80cb96da07ffad912f0c452f2d22f4a282e37720.PrivilegedOperationEventsRequestBuilder {
	return i82672a497b8f66c510d59f2a80cb96da07ffad912f0c452f2d22f4a282e37720.NewPrivilegedOperationEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// PrivilegedOperationEventsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.privilegedOperationEvents.item collection
func (m *GraphBaseServiceClient) PrivilegedOperationEventsById(id string) *i34d9be3c24f9832412eaf838a313f099ae8c0f497b918534db057b1bd9c8316f.PrivilegedOperationEventItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["privilegedOperationEvent%2Did"] = id
	}
	return i34d9be3c24f9832412eaf838a313f099ae8c0f497b918534db057b1bd9c8316f.NewPrivilegedOperationEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// PrivilegedRoleAssignmentRequests the privilegedRoleAssignmentRequests property
func (m *GraphBaseServiceClient) PrivilegedRoleAssignmentRequests() *if11203f2b5e6319285361e1998b4a25572cc1950d617d10ffd84e91e5f477349.PrivilegedRoleAssignmentRequestsRequestBuilder {
	return if11203f2b5e6319285361e1998b4a25572cc1950d617d10ffd84e91e5f477349.NewPrivilegedRoleAssignmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// PrivilegedRoleAssignmentRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.privilegedRoleAssignmentRequests.item collection
func (m *GraphBaseServiceClient) PrivilegedRoleAssignmentRequestsById(id string) *i7732e0def4b0366de3c1a4a0c7517a0f56ccce5a90b647e1a0e72997c0060fb1.PrivilegedRoleAssignmentRequestItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["privilegedRoleAssignmentRequest%2Did"] = id
	}
	return i7732e0def4b0366de3c1a4a0c7517a0f56ccce5a90b647e1a0e72997c0060fb1.NewPrivilegedRoleAssignmentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// PrivilegedRoleAssignments the privilegedRoleAssignments property
func (m *GraphBaseServiceClient) PrivilegedRoleAssignments() *i85eafc30b6e6233aefbca2b4e35ced0ac8836418de9401407873193666d812d9.PrivilegedRoleAssignmentsRequestBuilder {
	return i85eafc30b6e6233aefbca2b4e35ced0ac8836418de9401407873193666d812d9.NewPrivilegedRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// PrivilegedRoleAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.privilegedRoleAssignments.item collection
func (m *GraphBaseServiceClient) PrivilegedRoleAssignmentsById(id string) *i7a3a499cefaec6ced318a596ac37e721c173d3074f10b53b0db7be2eb1d91d1a.PrivilegedRoleAssignmentItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["privilegedRoleAssignment%2Did"] = id
	}
	return i7a3a499cefaec6ced318a596ac37e721c173d3074f10b53b0db7be2eb1d91d1a.NewPrivilegedRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// PrivilegedRoles the privilegedRoles property
func (m *GraphBaseServiceClient) PrivilegedRoles() *i32f7b810493504f82bd6c97020faab5a8ff5f541a46dd3d6b9cc2aa77fc22fe4.PrivilegedRolesRequestBuilder {
	return i32f7b810493504f82bd6c97020faab5a8ff5f541a46dd3d6b9cc2aa77fc22fe4.NewPrivilegedRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// PrivilegedRolesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.privilegedRoles.item collection
func (m *GraphBaseServiceClient) PrivilegedRolesById(id string) *i83a555d9eb9470ebf76ece54f97987bafd2a508dc0aa9e61e7d07311c57975c0.PrivilegedRoleItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["privilegedRole%2Did"] = id
	}
	return i83a555d9eb9470ebf76ece54f97987bafd2a508dc0aa9e61e7d07311c57975c0.NewPrivilegedRoleItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// PrivilegedSignupStatus the privilegedSignupStatus property
func (m *GraphBaseServiceClient) PrivilegedSignupStatus() *i6b96a96c52bbdff1731b8a5490cd5f342e33866e0931912944d323bc79f663e4.PrivilegedSignupStatusRequestBuilder {
	return i6b96a96c52bbdff1731b8a5490cd5f342e33866e0931912944d323bc79f663e4.NewPrivilegedSignupStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// PrivilegedSignupStatusById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.privilegedSignupStatus.item collection
func (m *GraphBaseServiceClient) PrivilegedSignupStatusById(id string) *ibe5c8d6d0c440fd8b6655bd846527fdabaf68dc6c79eb41da97daa4a59fd9b9f.PrivilegedSignupStatusItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["privilegedSignupStatus%2Did"] = id
	}
	return ibe5c8d6d0c440fd8b6655bd846527fdabaf68dc6c79eb41da97daa4a59fd9b9f.NewPrivilegedSignupStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// ProgramControls the programControls property
func (m *GraphBaseServiceClient) ProgramControls() *i178f0aa4e5987fcbfe2e98cbb6dd777ebcdcdf124dd3478d2bf40f83912ca030.ProgramControlsRequestBuilder {
	return i178f0aa4e5987fcbfe2e98cbb6dd777ebcdcdf124dd3478d2bf40f83912ca030.NewProgramControlsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// ProgramControlsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.programControls.item collection
func (m *GraphBaseServiceClient) ProgramControlsById(id string) *i9897ce060c8c140b6319d9eec7b10f1133ddeadce443b5c08d9e6e0b1a90a1fa.ProgramControlItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["programControl%2Did"] = id
	}
	return i9897ce060c8c140b6319d9eec7b10f1133ddeadce443b5c08d9e6e0b1a90a1fa.NewProgramControlItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// ProgramControlTypes the programControlTypes property
func (m *GraphBaseServiceClient) ProgramControlTypes() *if5cd0cf36bc86d9253920d73c41189ad8a30342e678d4f0138afa5095fd31538.ProgramControlTypesRequestBuilder {
	return if5cd0cf36bc86d9253920d73c41189ad8a30342e678d4f0138afa5095fd31538.NewProgramControlTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// ProgramControlTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.programControlTypes.item collection
func (m *GraphBaseServiceClient) ProgramControlTypesById(id string) *if01cb9f5e03c6cceecfc9647206bb0dc6e0bb97d9f416977b2683028defae8c7.ProgramControlTypeItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["programControlType%2Did"] = id
	}
	return if01cb9f5e03c6cceecfc9647206bb0dc6e0bb97d9f416977b2683028defae8c7.NewProgramControlTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Programs the programs property
func (m *GraphBaseServiceClient) Programs() *ib85b32f0384596c14f04b8d0f3dc8737da4b97428d7af145db2f1b06d7d9444e.ProgramsRequestBuilder {
	return ib85b32f0384596c14f04b8d0f3dc8737da4b97428d7af145db2f1b06d7d9444e.NewProgramsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// ProgramsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.programs.item collection
func (m *GraphBaseServiceClient) ProgramsById(id string) *i5b0c38b2eba003d29699607d25825d237b1b42b0380942484eba07941fd7ff1d.ProgramItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["program%2Did"] = id
	}
	return i5b0c38b2eba003d29699607d25825d237b1b42b0380942484eba07941fd7ff1d.NewProgramItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Reports the reports property
func (m *GraphBaseServiceClient) Reports() *ia5f30c11d37332bb7dfb48303b8e8880e60fc3560f60a813c66b9c1c3f2ff3f2.ReportsRequestBuilder {
	return ia5f30c11d37332bb7dfb48303b8e8880e60fc3560f60a813c66b9c1c3f2ff3f2.NewReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// RiskDetections the riskDetections property
func (m *GraphBaseServiceClient) RiskDetections() *i24b998459b0fbcdc6fbd83b90048a098ff6bbdcdaff773a2886f5cf8b3d5545e.RiskDetectionsRequestBuilder {
	return i24b998459b0fbcdc6fbd83b90048a098ff6bbdcdaff773a2886f5cf8b3d5545e.NewRiskDetectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// RiskDetectionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.riskDetections.item collection
func (m *GraphBaseServiceClient) RiskDetectionsById(id string) *i4dba4c746f2c06c2fec91dc54833c1da81c07726a80b85b29e72221e93c811ff.RiskDetectionItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["riskDetection%2Did"] = id
	}
	return i4dba4c746f2c06c2fec91dc54833c1da81c07726a80b85b29e72221e93c811ff.NewRiskDetectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// RiskyUsers the riskyUsers property
func (m *GraphBaseServiceClient) RiskyUsers() *iac77c9b5b86109e8ad626e30830db719efb3cc77c7babab332b409d84ae324a6.RiskyUsersRequestBuilder {
	return iac77c9b5b86109e8ad626e30830db719efb3cc77c7babab332b409d84ae324a6.NewRiskyUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// RiskyUsersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.riskyUsers.item collection
func (m *GraphBaseServiceClient) RiskyUsersById(id string) *i4f10f292ecab93ead1a3540c8285847ac93e9fd17579d21e923687140834b553.RiskyUserItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["riskyUser%2Did"] = id
	}
	return i4f10f292ecab93ead1a3540c8285847ac93e9fd17579d21e923687140834b553.NewRiskyUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// RoleManagement the roleManagement property
func (m *GraphBaseServiceClient) RoleManagement() *i310cda3e9f244aa61f9c9c78de433773f341b91c4a2310b8991671fe773be16e.RoleManagementRequestBuilder {
	return i310cda3e9f244aa61f9c9c78de433773f341b91c4a2310b8991671fe773be16e.NewRoleManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// SchemaExtensions the schemaExtensions property
func (m *GraphBaseServiceClient) SchemaExtensions() *i1b84a2c37ba0bbd175c6da40c8679db7d04dfcb044d8421d26d024db45218e4a.SchemaExtensionsRequestBuilder {
	return i1b84a2c37ba0bbd175c6da40c8679db7d04dfcb044d8421d26d024db45218e4a.NewSchemaExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// SchemaExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.schemaExtensions.item collection
func (m *GraphBaseServiceClient) SchemaExtensionsById(id string) *i03c5278b23fd725e2f49b358180d02b8204a81547ad7b8391ef4d31869b4e77c.SchemaExtensionItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["schemaExtension%2Did"] = id
	}
	return i03c5278b23fd725e2f49b358180d02b8204a81547ad7b8391ef4d31869b4e77c.NewSchemaExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// ScopedRoleMemberships the scopedRoleMemberships property
func (m *GraphBaseServiceClient) ScopedRoleMemberships() *i712907ad27a66d6ac32a26e01f88de1ad6484585eb7ed65f84b3a30571cec55b.ScopedRoleMembershipsRequestBuilder {
	return i712907ad27a66d6ac32a26e01f88de1ad6484585eb7ed65f84b3a30571cec55b.NewScopedRoleMembershipsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// ScopedRoleMembershipsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.scopedRoleMemberships.item collection
func (m *GraphBaseServiceClient) ScopedRoleMembershipsById(id string) *i9febca1a7c0d482f5cfd0d294e0b42f41284df1fc0f26b63341fd7066192a1d0.ScopedRoleMembershipItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["scopedRoleMembership%2Did"] = id
	}
	return i9febca1a7c0d482f5cfd0d294e0b42f41284df1fc0f26b63341fd7066192a1d0.NewScopedRoleMembershipItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Search the search property
func (m *GraphBaseServiceClient) Search() *i5840582f75a8eb78900edf3bd78566223ffee7aa1dc2f4cdca943ef635f6503e.SearchRequestBuilder {
	return i5840582f75a8eb78900edf3bd78566223ffee7aa1dc2f4cdca943ef635f6503e.NewSearchRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Security the security property
func (m *GraphBaseServiceClient) Security() *i761e9f0dec20dbf36c7fd626d107fb81ef94cafa7369422d2b2af143ffa16184.SecurityRequestBuilder {
	return i761e9f0dec20dbf36c7fd626d107fb81ef94cafa7369422d2b2af143ffa16184.NewSecurityRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// ServicePrincipals the servicePrincipals property
func (m *GraphBaseServiceClient) ServicePrincipals() *ibd3e65bb14e91a8a05d902c54fadec2c1b6931676c97f76da4969c975770aab2.ServicePrincipalsRequestBuilder {
	return ibd3e65bb14e91a8a05d902c54fadec2c1b6931676c97f76da4969c975770aab2.NewServicePrincipalsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// ServicePrincipalsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.servicePrincipals.item collection
func (m *GraphBaseServiceClient) ServicePrincipalsById(id string) *i1486770fbd64ddb3f673e750b91158c408009c45bd8e48463d5e06498271cd9c.ServicePrincipalItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["servicePrincipal%2Did"] = id
	}
	return i1486770fbd64ddb3f673e750b91158c408009c45bd8e48463d5e06498271cd9c.NewServicePrincipalItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Settings the settings property
func (m *GraphBaseServiceClient) Settings() *i714cbeb65962cb4d3e58007792fa4832d175c04614ba3aa7efb22871aea885bf.SettingsRequestBuilder {
	return i714cbeb65962cb4d3e58007792fa4832d175c04614ba3aa7efb22871aea885bf.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// SettingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.settings.item collection
func (m *GraphBaseServiceClient) SettingsById(id string) *ie532fdcd6a215be3d5886dcd2e4608af81583d060c7156a283880daf04ed9903.DirectorySettingItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["directorySetting%2Did"] = id
	}
	return ie532fdcd6a215be3d5886dcd2e4608af81583d060c7156a283880daf04ed9903.NewDirectorySettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Shares the shares property
func (m *GraphBaseServiceClient) Shares() *i73583652789c7aab226ac5bae66bc7b5fd924607d28350c4478c2a20524fd624.SharesRequestBuilder {
	return i73583652789c7aab226ac5bae66bc7b5fd924607d28350c4478c2a20524fd624.NewSharesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// SharesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.shares.item collection
func (m *GraphBaseServiceClient) SharesById(id string) *id8327d9e69ee76052759df500e96668a03b01b09bae354eaa67bbcf83427dc00.SharedDriveItemItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["sharedDriveItem%2Did"] = id
	}
	return id8327d9e69ee76052759df500e96668a03b01b09bae354eaa67bbcf83427dc00.NewSharedDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Sites the sites property
func (m *GraphBaseServiceClient) Sites() *i2817c6849b286be20c56215e039110b08a4109a12669579f4597fbab6f720ed9.SitesRequestBuilder {
	return i2817c6849b286be20c56215e039110b08a4109a12669579f4597fbab6f720ed9.NewSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// SitesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item collection
func (m *GraphBaseServiceClient) SitesById(id string) *ic557b65175a90fb05406286b014b07b096a266f27fc59e78636733800d66ff09.SiteItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["site%2Did"] = id
	}
	return ic557b65175a90fb05406286b014b07b096a266f27fc59e78636733800d66ff09.NewSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Solutions the solutions property
func (m *GraphBaseServiceClient) Solutions() *i7282246c788c7b44fcd028d8a1500ce1ffd32f655bee3b580b6a33d9ec10f610.SolutionsRequestBuilder {
	return i7282246c788c7b44fcd028d8a1500ce1ffd32f655bee3b580b6a33d9ec10f610.NewSolutionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Storage the storage property
func (m *GraphBaseServiceClient) Storage() *id14bafb4ee71e1257662cacab67dd479e54eba65ce10c6816ee4676804caf821.StorageRequestBuilder {
	return id14bafb4ee71e1257662cacab67dd479e54eba65ce10c6816ee4676804caf821.NewStorageRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// SubscribedSkus the subscribedSkus property
func (m *GraphBaseServiceClient) SubscribedSkus() *ie934faa615fb56652e5964395b3dc205321ac84e8cf244796ebe59ba3713fbd9.SubscribedSkusRequestBuilder {
	return ie934faa615fb56652e5964395b3dc205321ac84e8cf244796ebe59ba3713fbd9.NewSubscribedSkusRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// SubscribedSkusById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.subscribedSkus.item collection
func (m *GraphBaseServiceClient) SubscribedSkusById(id string) *i3e0ee15b77f1a6df2686711e1700bf6812443daf08105eb0b15f4b72d6077b93.SubscribedSkuItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["subscribedSku%2Did"] = id
	}
	return i3e0ee15b77f1a6df2686711e1700bf6812443daf08105eb0b15f4b72d6077b93.NewSubscribedSkuItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Subscriptions the subscriptions property
func (m *GraphBaseServiceClient) Subscriptions() *i60392bb6eb86abe7a3079e3b2b1e202f7dcf3452adf026db62ec93e2fafe3957.SubscriptionsRequestBuilder {
	return i60392bb6eb86abe7a3079e3b2b1e202f7dcf3452adf026db62ec93e2fafe3957.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.subscriptions.item collection
func (m *GraphBaseServiceClient) SubscriptionsById(id string) *ide40baaefabe61fc7133751aae85320ce5f6bc70e3dde81b419c4861d1e1bfe0.SubscriptionItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["subscription%2Did"] = id
	}
	return ide40baaefabe61fc7133751aae85320ce5f6bc70e3dde81b419c4861d1e1bfe0.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Teams the teams property
func (m *GraphBaseServiceClient) Teams() *iff395ba1da21566390b02b5bed781aecf3bb849fc71f2359410792d1d1b67079.TeamsRequestBuilder {
	return iff395ba1da21566390b02b5bed781aecf3bb849fc71f2359410792d1d1b67079.NewTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// TeamsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item collection
func (m *GraphBaseServiceClient) TeamsById(id string) *i86ffda5c6f886142a63c7cd1c0f489d73a5f8c36310f880ba8dd82807977d46e.TeamItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["team%2Did"] = id
	}
	return i86ffda5c6f886142a63c7cd1c0f489d73a5f8c36310f880ba8dd82807977d46e.NewTeamItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// TeamsTemplates the teamsTemplates property
func (m *GraphBaseServiceClient) TeamsTemplates() *i66f18ccab4e34309d26d1056f0e7dd8b563a5f8ee6f8d9c6e8e77c5fac50f8b5.TeamsTemplatesRequestBuilder {
	return i66f18ccab4e34309d26d1056f0e7dd8b563a5f8ee6f8d9c6e8e77c5fac50f8b5.NewTeamsTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// TeamsTemplatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamsTemplates.item collection
func (m *GraphBaseServiceClient) TeamsTemplatesById(id string) *i44e397fa023407b3a9722e619d02140543ef889c8c03de376d8341cec5a43794.TeamsTemplateItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["teamsTemplate%2Did"] = id
	}
	return i44e397fa023407b3a9722e619d02140543ef889c8c03de376d8341cec5a43794.NewTeamsTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// TeamTemplateDefinition the teamTemplateDefinition property
func (m *GraphBaseServiceClient) TeamTemplateDefinition() *i1c7e7a5d0708841f8c98ec910d583f348cbffaad386ef9a24d3ee4eba285ea21.TeamTemplateDefinitionRequestBuilder {
	return i1c7e7a5d0708841f8c98ec910d583f348cbffaad386ef9a24d3ee4eba285ea21.NewTeamTemplateDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// TeamTemplateDefinitionById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamTemplateDefinition.item collection
func (m *GraphBaseServiceClient) TeamTemplateDefinitionById(id string) *i9edbc585c67673e5bfc1e312a7fb5f8fca5cb966cef31e6fdfe8ab8ab7d5780d.TeamTemplateDefinitionItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["teamTemplateDefinition%2Did"] = id
	}
	return i9edbc585c67673e5bfc1e312a7fb5f8fca5cb966cef31e6fdfe8ab8ab7d5780d.NewTeamTemplateDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Teamwork the teamwork property
func (m *GraphBaseServiceClient) Teamwork() *icb4f253cb1cd35435f5752b611229032c618bbcfeb3be80ee4d6a06d404114fc.TeamworkRequestBuilder {
	return icb4f253cb1cd35435f5752b611229032c618bbcfeb3be80ee4d6a06d404114fc.NewTeamworkRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// TenantRelationships the tenantRelationships property
func (m *GraphBaseServiceClient) TenantRelationships() *id5c2ef977a00dd1757d258dbbbfb4080031771e62e6c6b3b1339a0f03fc1c1f1.TenantRelationshipsRequestBuilder {
	return id5c2ef977a00dd1757d258dbbbfb4080031771e62e6c6b3b1339a0f03fc1c1f1.NewTenantRelationshipsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// TermStore the termStore property
func (m *GraphBaseServiceClient) TermStore() *i98a1471d41b15330865bc87691830281af9ecf479bfc797e54f02448790b1e4e.TermStoreRequestBuilder {
	return i98a1471d41b15330865bc87691830281af9ecf479bfc797e54f02448790b1e4e.NewTermStoreRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// ThreatSubmission the threatSubmission property
func (m *GraphBaseServiceClient) ThreatSubmission() *i231953b44a7f4aace0800ccf375537d423e6f60e82f2a6552d9613626e39aba5.ThreatSubmissionRequestBuilder {
	return i231953b44a7f4aace0800ccf375537d423e6f60e82f2a6552d9613626e39aba5.NewThreatSubmissionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// TrustFramework the trustFramework property
func (m *GraphBaseServiceClient) TrustFramework() *i312c0a09d8ded5436957205a14adfc7e2facbcc6f26ef9872a5b5eb79228375f.TrustFrameworkRequestBuilder {
	return i312c0a09d8ded5436957205a14adfc7e2facbcc6f26ef9872a5b5eb79228375f.NewTrustFrameworkRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// Users the users property
func (m *GraphBaseServiceClient) Users() *icd01c84a90833c55ac2309fd7034cb1962c60f59eb1ee2b2cf7b04c708402b6a.UsersRequestBuilder {
	return icd01c84a90833c55ac2309fd7034cb1962c60f59eb1ee2b2cf7b04c708402b6a.NewUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// UsersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item collection
func (m *GraphBaseServiceClient) UsersById(id string) *id9f276010196e2d79e634ec622333b5c53dc0fbc407a6c9aa27ca92d4f388ed3.UserItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["user%2Did"] = id
	}
	return id9f276010196e2d79e634ec622333b5c53dc0fbc407a6c9aa27ca92d4f388ed3.NewUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}

// Workbooks the workbooks property
func (m *GraphBaseServiceClient) Workbooks() *i0c8acb61c4c4d82e7ac6f352bd7cdb24dcd0de64d904e31cf17ca8c6b1d44202.WorkbooksRequestBuilder {
	return i0c8acb61c4c4d82e7ac6f352bd7cdb24dcd0de64d904e31cf17ca8c6b1d44202.NewWorkbooksRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// WorkbooksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.workbooks.item collection
func (m *GraphBaseServiceClient) WorkbooksById(id string) *i36b4c1bee4db70277ed8b9217e494480565564fd187f7a35b920807b10c82740.DriveItemItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["driveItem%2Did"] = id
	}
	return i36b4c1bee4db70277ed8b9217e494480565564fd187f7a35b920807b10c82740.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
