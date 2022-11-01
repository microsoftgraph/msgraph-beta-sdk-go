package reports

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i003173036331dffc584bd55a4a31f9eca0b7cc880c4006f990ef8ec3e90f8d61 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getbrowserdistributionusercountswithperiod"
    i00a0e14c627310f3b6a8da8157ec57ebbdae7fc5ecb2a65e248053e4da1956de "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/dailyprintusage"
    i036a1b2a2d29c6542eb2d7482b6087baec6694c844689f95bde300288b5d9af8 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getoffice365activeusercountswithperiod"
    i04f6ee74f0733278b8c6e9678d187269b49dcb665a9958b4ae6257f0aaec8dbf "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getskypeforbusinessdeviceusageuserdetailwithperiod"
    i088887eceaa935280d9592f89bc4b0a045fd924ff8d44ab2b855e031094ea540 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getteamsuseractivitydistributionusercountswithperiod"
    i0ee449f331a7fd89049ddd3e1a2b7b309f269639673e8bd0aa048ffdf3f4e37b "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getemailappusageversionsusercountswithperiod"
    i103d63d4193b9f3ab75fe6b66304aaeb64366749ff445bc88c2dc5beed1269cb "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getyammergroupsactivitygroupcountswithperiod"
    i125e8811c6de949618151abfadcaca28ef05c938e0252591bfa83691e13c2070 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getcredentialuserregistrationcount"
    i12820273d3c44672ae92304278e27099019c4b7ff9e6dce947a674b59ff6c98e "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/deviceconfigurationdeviceactivity"
    i1618be8d311dc94393db28762cd325c59b9ccde3cb914a388979d02c7c993fb6 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getoffice365activeuserdetailwithperiod"
    i1775513b7b2429eb61cc9982fc3d0c155badba7bb5c8436afb785b645183be6a "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getteamsuseractivitydistributiontotalusercountswithperiod"
    i18f58919f6d5882740586016a8bc9807f1c45848f93ad165acc8577a78caaf9d "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getteamsuseractivitytotalcountswithperiod"
    i1a614f493010260020309e60ff2e8606cd018fda1d411bd62662caf405dbfa5b "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getonedriveusagestoragewithperiod"
    i1a7aad10a2334fe03a5643720cc6b428cdff146a4fc62c7a50a78f70f9a36bc4 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getsharepointactivityusercountswithperiod"
    i1b39eea80211466ef73aee8b2f1a50b08a56f3dce15eda5a18ede70e235f90d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getyammerdeviceusageusercountswithperiod"
    i1ffd159fb8aed8eab5637f0dcf99bc4e9560b54299bc87e4f88446b96ca862f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getemailactivityuserdetailwithperiod"
    i20836d38a167d95c34498a5116a9cd82d480e78642fcca162facb20b00fdc534 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getoffice365activationsusercounts"
    i2157001b3d1627cbb758e52e666c104b24725dfdb773f94732a2c17226613e0c "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getonedriveactivityfilecountswithperiod"
    i21b3b996d3b537f9331bc504ac14bc8a80a52a412b2e26c73bfcc4ab966b11b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getskypeforbusinessactivityusercountswithperiod"
    i224d112f4527ff133ce9d06395ccae8994e1f895836fdc4374f4a129a5a06bed "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getskypeforbusinessorganizeractivitycountswithperiod"
    i2499892f567f8728fe2de683d0f00e72080cfd5ffca87cc0785ef89eb730a931 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getteamsdeviceusagedistributionusercountswithperiod"
    i254c294d6b358b3854c3ee45f61d16c1237e0477952e917b034bbb0d8a6802b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/usercredentialusagedetails"
    i27c74777b0efbee67b16cbd84335b5228b74a547d6ffe5e338e3c1c6cc55c61b "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/monthlyprintusagesummariesbyprinter"
    i2a4862a2c04d7e8705aebac57aeab2e0bf3268230862ae87191095bb620b0409 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getsharepointsiteusagedetailwithdate"
    i2d772c8afb07a24a1205c9ff6577c716c80df2a273be74d3b989c999d09199da "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getm365appusercountswithperiod"
    i2e11d2e753b93ae2e19efcae2b9d79d3efa476d4ee469d6db9c173f35ef3b620 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/manageddeviceenrollmentfailuredetails"
    i306349847f73d7bf75fde4cca87b9e847d160d4298503961ee7173a877566811 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getemailappusageuserdetailwithperiod"
    i313577bbc7d677412230e87962a7a9319b6aabde4851f6c8739be336203bfa65 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getemailappusageusercountswithperiod"
    i328b60d96035cb1d9bd616db71d49e8436746e43b59c2ae74edb06e1c65f6158 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getoffice365groupsactivitydetailwithperiod"
    i328e3d55288de86c60aa3add79cd83a5c047661caef06da2cff1cef6c8598b34 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getskypeforbusinessparticipantactivitycountswithperiod"
    i330d7d86bcb93d59ef14fc1257795533ce121643f4e8becd8884b57bc63e7b4b "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/monthlyprintusagebyprinter"
    i35b216d7dc42f80efd855cd81dcf7debc2bdafd7ea24b32bcb1d686238dbc76d "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getemailappusageappsusercountswithperiod"
    i3772e364c43656cba18386bc364e131c7943e3a0b0461e3a162fb5eae513ac0b "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/dailyprintusagesummariesbyuser"
    i3a0d2235d08fac82d349f0d0e02e0cdce3a0c009ad1dd8bf7208224ae85f089f "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getm365appuserdetailwithdate"
    i3e7d149707f2ddd6910123604eff6dd78480edcb3f05321bb44037a85d733380 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getyammergroupsactivitycountswithperiod"
    i40dc3bbc1c58ea07c4c8b16942aaeca4da00eb3514241eef66edcd213ba084f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/monthlyprintusagesummariesbyuser"
    i4113a715b496e73819d650e2951dc88df01049bb7ee89faa36b869ccf6a2f71e "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getteamsuseractivitytotalusercountswithperiod"
    i420730d2f754eb6fcc9ad1b5310922c60bb4e377cbdde02aaac29b3fa0f0b2e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getskypeforbusinessparticipantactivityminutecountswithperiod"
    i423884002844f09964fb45ec36b7ef31192751d8bb7a78a54fb126d4206b1025 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getskypeforbusinesspeertopeeractivitycountswithperiod"
    i435102bd0023ed78a36db74d804d78f55bf2d1ef197f516b037bb3b82f517c08 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getskypeforbusinessorganizeractivityusercountswithperiod"
    i46fd699154cb21987f377181c37a09acabdc0e13ef2aec22fc2f1352022502ba "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/dailyprintusagebyuser"
    i4add8d70f1a45f4e5d237fb6c9877de552324641d4f73a6bcd482c6db9801fcd "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getteamsdeviceusagetotalusercountswithperiod"
    i4c6dc125c665e98f5e3a0f5b0343f5fdacc0d0a7e3ae60637e7e73e22e4f0ca4 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/manageddeviceenrollmentfailuredetailswithskipwithtopwithfilterwithskiptoken"
    i4d7d4fe8dd8ea3539a0a8ad2f9a8e5d271857ba237dce7b4a8ca536a616ec435 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getteamsdeviceusageuserdetailwithperiod"
    i4f829313dc734307b5536ec2ce241c78a78f218e561a61fdf90fce1e9092a56a "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetime"
    i50fd31667e7b85e595d3060b1d746877ba37f9b2c6fae47e8ee37f5c2f965515 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getoffice365groupsactivitystoragewithperiod"
    i5125d764ebc3d13de453dc8ae6109633072221c35bf64d9e7655b7445242f381 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getyammerdeviceusagedistributionusercountswithperiod"
    i567056fba35e6edab31f5aa9994989fd7742a56af8e01ec0f240852321374891 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getteamsuseractivitycountswithperiod"
    i597d5f57aae856b81874eda3863201bedf21b61ece754b9b079304509ca2068a "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getonedriveactivityuserdetailwithperiod"
    i5ba6893546e6aaea0f4a67b2dd033135639babe968232c6fcc5163ab253c7847 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/dailyprintusagesummariesbyprinter"
    i5c1ccbcf17345644477286c949022f98da0c3dbfc12ebc979ef7ba6edea3b759 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getemailactivitycountswithperiod"
    i5df2c32f3d48e07033ebe1c8ef2dd7d264320343ff53c690d8fa229633336de2 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getsharepointsiteusagefilecountswithperiod"
    i6532f3214bfaa5d2aac58a9d9ffe42fb56e5e2242405e98165287c807365bcda "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/monthlyprintusagebyuser"
    i6bb9a3b51383aadb75ab3111f39327408462a143d79cb9106e695674d98c226a "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getoffice365activationsuserdetail"
    i6cba3dd24e78bc27f2b14a43b1489601f5f50bf4febd3101a2bdfcce9eb1c31e "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getemailappusageuserdetailwithdate"
    i72fab10b39cc1890375f583eea81a5d54930bcd29915bbd6d1b82e9b92ae1e64 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getsharepointactivityuserdetailwithdate"
    i7355efd26cb65bce9c72b3b63c77e2dba43919588a8e86369b208aefe72c690c "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getoffice365groupsactivityfilecountswithperiod"
    i75ebd72e88d0574c533352a054cd82b52553ab8eb695b568d8aaa890ce02579f "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getteamsuseractivityusercountswithperiod"
    i7a5b000346e54a3c0bb5a39921329be717069bb5aa6e2bb74c97770c0a1f79ce "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getoffice365groupsactivitycountswithperiod"
    i7ab3166690500d92e204211f333c4697c5188e65cef5efac6c698b404cea6b48 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getemailactivityuserdetailwithdate"
    i7b88d33e4fa5702328c128d0552680d05cd2414a20ea062fc43f3458e04682c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getskypeforbusinessactivitycountswithperiod"
    i804fbe3d687c5d168bcbde36b31d895d3817ace737cf5ae5896736def6e2624e "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/manageddeviceenrollmentfailuretrends"
    i808e335fa07739775bf1a553e8afacccf25cd8047be1de0c119ce5caa8554634 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getteamsdeviceusageusercountswithperiod"
    i817285b5e7fea10651e0d4b79aa6611ca57de96e79203aa153ccde9c90f07dd1 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getonedriveusageaccountdetailwithdate"
    i81aca7de3573f1e665c768b347600ba5c3fb9fede4a50466b75c3b9b5586bf78 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getskypeforbusinessorganizeractivityminutecountswithperiod"
    i859de82968bd758e08878a67874d4d3a4d19758bd7bcbe1ad402adc066dbaee2 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getteamsuseractivityuserdetailwithperiod"
    i8a126339410006480a905b751a907b1e7c5cd226076269437ad939d10cb7a71f "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetime"
    i8a1d439fc44307034a531d7f5d9324bb3c5e5b3c14bbacf473e28c1319852fcf "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getteamsuseractivitytotaldistributioncountswithperiod"
    i8fc0ad0d8f4c5cb82e70d74504fd7555d8235044ab2b40bc6e68738d416134ae "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getazureadapplicationsigninsummarywithperiod"
    i9023ee3c3a9e19cbce82a481a61ad1ced99450d1c7561a3d19d00e038450a05a "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getmailboxusagemailboxcountswithperiod"
    i932396fbe8ecd019f26fa2a7bce06d747e0bebe4e79cc653e70943aa11c9355d "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getteamsteamactivitydistributioncountswithperiod"
    i95b95e9f1f7725e5d0fd693531d9686b616e391472cd4b61b5f11da6c9cc5526 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getteamsdeviceusagedistributiontotalusercountswithperiod"
    i9631ad5d28f0b71849e4f20d19d1958e0d952c9631900bea7802308681c1759b "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getonedriveactivityusercountswithperiod"
    i96fe7fb8deb79d07ea5e417d63df350e76d28fdd0491378b2801b6e248d018c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getattacksimulationsimulationusercoverage"
    i97727cca9b8c936e00dd03bd736857d99fcb279322f6334f8075719eee6f11b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getonedriveusageaccountdetailwithperiod"
    i983e4c56868ed790407309ebf34264a39588d5f319fbc8ea49bc989533975765 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getyammeractivityuserdetailwithperiod"
    i9b2d2e6f90f5f9da6a8f69e6d4e9b3e43b80147df7cbcd17ae490b6d7b0a4194 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getteamsuseractivityuserdetailwithdate"
    i9fd0210ad4ae49215efc7cf22136268ddeef1c0bff45a84f042e4597a495e2c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getsharepointsiteusagestoragewithperiod"
    ia11c96eb7fc01f9081a8c45eff690c39a5c79c0b439f8e4687f3fd30ab5b707c "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getemailactivityusercountswithperiod"
    ia135a492a1ea4fe1e4e4616ad09038e519f2792c9d98bbdca0bab0ae0ea58926 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getteamsteamactivitydetailwithdate"
    ia290e37e65f4257fc2cba6045f8fae24a9fcb0902933df2bbcd9ea8ac77cb2ef "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getskypeforbusinessactivityuserdetailwithdate"
    ia337a32890551b1081a608ae43abf61c4497dfd42048dd808bb7cf09832281c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getskypeforbusinessdeviceusageusercountswithperiod"
    ia3c04ca51a3e25cdb708373964025e87292d13441d9e90fc357ca0a1621ebd95 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getmailboxusagequotastatusmailboxcountswithperiod"
    ia3e34a08513d2e3ca7ffcfad8b20f3382618d2a7552421e981f109b27cad2170 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/manageddeviceenrollmenttopfailureswithperiod"
    ia46b9d8d0e59a725b189c3202c5343f9f35bf2a8a04d4785295f7a71e6be9590 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getyammerdeviceusageuserdetailwithdate"
    ia58f01f9d5b233aa2565b0e4538e84e12cb2a0d586ef5a56fbbb816826c2bba2 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getcredentialusagesummarywithperiod"
    ia74d4f7cd8fb70eb99b3c94cccde51ec24c099b4b64a46531b40bf30619e07c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/authenticationmethods"
    iaa160176f699a99c070a1f1f49a0887c07919a5c689d21c8bed8f52de4a89023 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getonedriveusagefilecountswithperiod"
    iaaafb6c998e8e1eedbeeb89216037d37a8bc2089fc1568d394cdf1e95cec3c91 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getyammeractivitycountswithperiod"
    iad427df3a804c5ea6bf927439cbace855bd9edf35dec361d7cc896c86a3c3e76 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getrelyingpartydetailedsummarywithperiod"
    iae029900a0da9f48c34f75ee5d793b8dc3931fe12b0f7a985040c1d5ec7c5523 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getskypeforbusinesspeertopeeractivityminutecountswithperiod"
    iae0b95416e4f0115ded4c5680f5308d53f4ab1191730356055fa78cb04fdb3a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getsharepointsiteusagedetailwithperiod"
    ib179175741f9263f3b9b2d70c2568f051b6ea034ea8268e68071c8691030b092 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getoffice365servicesusercountswithperiod"
    ib66df0db90f0ef258b3536c6659fe0d6ede8fd889a2c38cf8c63f72fb5f1ffaa "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getyammergroupsactivitydetailwithdate"
    ib749d3394cf27f1397536acbbae8cf02a899e5a39e5596471124bf4d71501852 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getonedriveusageaccountcountswithperiod"
    ib9134a9409b3aeb31e9bbde8b92be81f399cc5028fd2d006006539e98ca706e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getskypeforbusinessactivityuserdetailwithperiod"
    iba529b34d22e6d4619a98e3a64864d4ce62443496139f8d961a28f3388b1c06b "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getuserarchivedprintjobswithuseridwithstartdatetimewithenddatetime"
    ibb4d5055652a4d2b25c6f9515b9c6f5b7c8a1bb709ad42af6daac709a178858f "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getteamsteamactivitydetailwithperiod"
    ibe3b90074f3fd6bd15c064326b107884d8dcff2d4f739ea9969038710d070c78 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getsharepointactivityuserdetailwithperiod"
    ibf41f160c65ef1194ee9e34f2d2b3bbf34a32e8ceb9026adff36a62f7067a2ac "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/dailyprintusagebyprinter"
    ibf5abd7cd37fc0ab2d9ca27fab459ff2a81a89b886033d344b5d0d36bfa92d6a "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getsharepointactivitypageswithperiod"
    ibfa03aafca1026d6cb462cc2b4a03d52cfcd50b0378e62a671000fb1c4cae155 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/manageddeviceenrollmentabandonmentsummarywithskipwithtopwithfilterwithskiptoken"
    ic05da5b8d4cf94558ff5262e866399feadd45bc394ba1441e1fc14d600fde57a "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getskypeforbusinessdeviceusageuserdetailwithdate"
    ic2e883868a9a28d28dfcc1e662d0e24885b4e5da00486835cbbe6d165c5798a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/manageddeviceenrollmenttopfailures"
    ic65cd1a979952b5550b39c84cadcef5f570c546d773d487f9bbe744c4a2ead72 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getsharepointsiteusagepageswithperiod"
    ic753e229b298240003ae8ca9aff000f32462080e1d16dc2def476877d5707d80 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/deviceconfigurationuseractivity"
    ic7a209730cd8537604c6db9508d277bc2e323d048bba634bc00a0ff07184baac "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getskypeforbusinessparticipantactivityusercountswithperiod"
    ic8b36546821464e310499bcde1e62e440cf6796bae121c31bff3e399487901dc "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getm365appplatformusercountswithperiod"
    ica973266b26481209852c53969fd134273a48e20a686d98f2c7e4ac2bb684ece "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getyammeractivityusercountswithperiod"
    icb9fbfdf1a799b2b262a2d5e2e5889891e3c531a929c01dc751359da88307334 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getskypeforbusinessdeviceusagedistributionusercountswithperiod"
    icc5ddf4fe31a3f4d96b22a2a2e64d31f88f217b06f59946a62eccd7032aa763d "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getsharepointactivityfilecountswithperiod"
    iccd91022582f2d8dde0cb9b960990f411b495b6ff93387a25fe59ca8a9961562 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getonedriveactivityuserdetailwithdate"
    icd1b1e4da27d3da9a9df7b33ef242fd4e79c8a35c3cddfe9d580b7e401250239 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getmailboxusagedetailwithperiod"
    icdcce25b5e3ff04e70a3cf70bfbfae9ec805346c7c065f15b4fc6b4b165bbf48 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getoffice365activeuserdetailwithdate"
    icddea0ed347960de17929a877880af4ac0bf296fb68f5c0e35eebd457fc0734f "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/applicationsignindetailedsummary"
    icddeafe68b5aefb5586975a9974bb7370a97219ada28c815ec9b8e3b7f284748 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getmailboxusagestoragewithperiod"
    ice018542cd139d6164abff37b45f7b6a1d3055a473b3658086e14293b59be16e "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/credentialuserregistrationdetails"
    id3ee30e7ad89e8fb7901e0c16cfe88b41322719fa7407bf70b021ab61e421a6e "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getm365appuserdetailwithperiod"
    id4c53d2ba82691363ac26d457be4a4cdcf13b6e128592549cf1224910bd3751c "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/manageddeviceenrollmentabandonmentdetailswithskipwithtopwithfilterwithskiptoken"
    id7392fa684625fd57bfdfc7c7dfbc56d0065c1cc7093f8a6d32a219c2ba1dfac "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getskypeforbusinesspeertopeeractivityusercountswithperiod"
    idf5d61f783c212fb53a8df5f35707485128386131f3ae007235e92b0eb00404c "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getteamsteamactivitycountswithperiod"
    ie050dab9fab3553491bfd68cf2523de38e35cfaa538b49fdcf301f61cb13e9e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getyammergroupsactivitydetailwithperiod"
    ie591696fdc474e68dec24db3dde3496dbb37168bbafe4d0a2b02ab16f890232e "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getyammerdeviceusageuserdetailwithperiod"
    ie5d033fba5fb2263cfd1bff0d67c57e993013c03f754b188e52f7a6cc45ba107 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getattacksimulationrepeatoffenders"
    ie7a68204cc4982a6b52ad00381436561da8774cf91ceef439cd495875cd84e9d "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getoffice365groupsactivitydetailwithdate"
    ie822df7440dd5b07f2cb5b9ae611cdc7ec11b9262d220f5682bd2958e02a643e "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getteamsdeviceusageuserdetailwithdate"
    ie8d93a68bfc5d77cb89a3bcaa58d306cf9c1eb664f150283d9e7c99bb68fdd94 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getattacksimulationtrainingusercoverage"
    if278596fba20816d527a904fbb7b5d78e3c177aa4b09c4e15657da9d6d06a2db "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getbrowserusercountswithperiod"
    if31836807723ff5bdd6b6e4e26ed08acbebd30a643cc4b8b53069f1facbbe742 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getoffice365groupsactivitygroupcountswithperiod"
    if420f5f31b1d54d7e335630814dcebdb92a9d2ca57cafb5c1aee864810f1d4c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getsharepointsiteusagesitecountswithperiod"
    if4dd226e0f2c5d03d114b08150ac66aef29652e888dbf65028360ee02623bdf9 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getbrowseruserdetailwithperiod"
    if6320aed42bd57df6018ebfaa8490e375970ae1b65838874224d2524e3133811 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getoffice365activationcounts"
    ifda5935468b5ae9cd912cd8886c9d4cca1e9ab52a6c082d9530ad21d01d3c3ad "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/getyammeractivityuserdetailwithdate"
    ifece25d39559e83e159712abc947b9fa9ef645ac39026b276f4c512910960f48 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/security"
    i1453ad3549b1a27e47913c2d4a29d5625c5b5e04deb3a2a6f264f5cd24a2e8da "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/dailyprintusagesummariesbyprinter/item"
    i1702ca3f0c998e29512738841125456c0b4e81357e5f7ca9f4176443002d28d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/monthlyprintusagebyuser/item"
    i2268eb94e47ea5d81b666c3310292b49733c18b0196efe2957b868503a9da7ff "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/monthlyprintusagesummariesbyuser/item"
    i22a3d34b6ea91c582f9826d88dc039aae3e5f8f62146be9df6713c324659e396 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/credentialuserregistrationdetails/item"
    i3808d92a196a8e86ccfe0455bb6aaf56bf5956b11460f112fbc1dbcd8375551d "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/dailyprintusagebyprinter/item"
    i4ca23ea1fafc22df57a5d1150889ab294f6d1582a79d145b961145b60dbf4c5b "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/dailyprintusagebyuser/item"
    i4f1c6062b781770b431d75b41deb0a2ac227468563ebac0623bef14f3efe351d "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/usercredentialusagedetails/item"
    i55a12b578425084e14afca840f9f8ad56ad8296a76d84139beb67b9619a01ae2 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/monthlyprintusagebyprinter/item"
    i6581ea85ca9063d302a537e4fdb5af97c20fd90311aa357f5482935e84bcd517 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/applicationsignindetailedsummary/item"
    ic1e71dc755b571267687f63b13312b758bc574de61144bfa50cd81af3b6d3626 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/monthlyprintusagesummariesbyprinter/item"
    ic59c0950801f0cba051b02aec519c855e67a10d949f0b0313e158f5867985f61 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/dailyprintusage/item"
    iffede79819a5c0828784daecd91c1db9829729eb545a0c3b1948096142482a96 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/dailyprintusagesummariesbyuser/item"
)

// ReportsRequestBuilder provides operations to manage the reports property of the microsoft.graph.print entity.
type ReportsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ReportsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ReportsRequestBuilderGetQueryParameters get reports from print
type ReportsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ReportsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ReportsRequestBuilderGetQueryParameters
}
// ReportsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApplicationSignInDetailedSummary the applicationSignInDetailedSummary property
func (m *ReportsRequestBuilder) ApplicationSignInDetailedSummary()(*icddea0ed347960de17929a877880af4ac0bf296fb68f5c0e35eebd457fc0734f.ApplicationSignInDetailedSummaryRequestBuilder) {
    return icddea0ed347960de17929a877880af4ac0bf296fb68f5c0e35eebd457fc0734f.NewApplicationSignInDetailedSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApplicationSignInDetailedSummaryById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.applicationSignInDetailedSummary.item collection
func (m *ReportsRequestBuilder) ApplicationSignInDetailedSummaryById(id string)(*i6581ea85ca9063d302a537e4fdb5af97c20fd90311aa357f5482935e84bcd517.ApplicationSignInDetailedSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["applicationSignInDetailedSummary%2Did"] = id
    }
    return i6581ea85ca9063d302a537e4fdb5af97c20fd90311aa357f5482935e84bcd517.NewApplicationSignInDetailedSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AuthenticationMethods the authenticationMethods property
func (m *ReportsRequestBuilder) AuthenticationMethods()(*ia74d4f7cd8fb70eb99b3c94cccde51ec24c099b4b64a46531b40bf30619e07c4.AuthenticationMethodsRequestBuilder) {
    return ia74d4f7cd8fb70eb99b3c94cccde51ec24c099b4b64a46531b40bf30619e07c4.NewAuthenticationMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewReportsRequestBuilderInternal instantiates a new ReportsRequestBuilder and sets the default values.
func NewReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsRequestBuilder) {
    m := &ReportsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/reports{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewReportsRequestBuilder instantiates a new ReportsRequestBuilder and sets the default values.
func NewReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property reports for print
func (m *ReportsRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ReportsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get reports from print
func (m *ReportsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ReportsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property reports in print
func (m *ReportsRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable, requestConfiguration *ReportsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CredentialUserRegistrationDetails the credentialUserRegistrationDetails property
func (m *ReportsRequestBuilder) CredentialUserRegistrationDetails()(*ice018542cd139d6164abff37b45f7b6a1d3055a473b3658086e14293b59be16e.CredentialUserRegistrationDetailsRequestBuilder) {
    return ice018542cd139d6164abff37b45f7b6a1d3055a473b3658086e14293b59be16e.NewCredentialUserRegistrationDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CredentialUserRegistrationDetailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.credentialUserRegistrationDetails.item collection
func (m *ReportsRequestBuilder) CredentialUserRegistrationDetailsById(id string)(*i22a3d34b6ea91c582f9826d88dc039aae3e5f8f62146be9df6713c324659e396.CredentialUserRegistrationDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["credentialUserRegistrationDetails%2Did"] = id
    }
    return i22a3d34b6ea91c582f9826d88dc039aae3e5f8f62146be9df6713c324659e396.NewCredentialUserRegistrationDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DailyPrintUsage the dailyPrintUsage property
func (m *ReportsRequestBuilder) DailyPrintUsage()(*i00a0e14c627310f3b6a8da8157ec57ebbdae7fc5ecb2a65e248053e4da1956de.DailyPrintUsageRequestBuilder) {
    return i00a0e14c627310f3b6a8da8157ec57ebbdae7fc5ecb2a65e248053e4da1956de.NewDailyPrintUsageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.dailyPrintUsage.item collection
func (m *ReportsRequestBuilder) DailyPrintUsageById(id string)(*ic59c0950801f0cba051b02aec519c855e67a10d949f0b0313e158f5867985f61.PrintUsageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsage%2Did"] = id
    }
    return ic59c0950801f0cba051b02aec519c855e67a10d949f0b0313e158f5867985f61.NewPrintUsageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DailyPrintUsageByPrinter the dailyPrintUsageByPrinter property
func (m *ReportsRequestBuilder) DailyPrintUsageByPrinter()(*ibf41f160c65ef1194ee9e34f2d2b3bbf34a32e8ceb9026adff36a62f7067a2ac.DailyPrintUsageByPrinterRequestBuilder) {
    return ibf41f160c65ef1194ee9e34f2d2b3bbf34a32e8ceb9026adff36a62f7067a2ac.NewDailyPrintUsageByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageByPrinterById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.dailyPrintUsageByPrinter.item collection
func (m *ReportsRequestBuilder) DailyPrintUsageByPrinterById(id string)(*i3808d92a196a8e86ccfe0455bb6aaf56bf5956b11460f112fbc1dbcd8375551d.PrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter%2Did"] = id
    }
    return i3808d92a196a8e86ccfe0455bb6aaf56bf5956b11460f112fbc1dbcd8375551d.NewPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DailyPrintUsageByUser the dailyPrintUsageByUser property
func (m *ReportsRequestBuilder) DailyPrintUsageByUser()(*i46fd699154cb21987f377181c37a09acabdc0e13ef2aec22fc2f1352022502ba.DailyPrintUsageByUserRequestBuilder) {
    return i46fd699154cb21987f377181c37a09acabdc0e13ef2aec22fc2f1352022502ba.NewDailyPrintUsageByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageByUserById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.dailyPrintUsageByUser.item collection
func (m *ReportsRequestBuilder) DailyPrintUsageByUserById(id string)(*i4ca23ea1fafc22df57a5d1150889ab294f6d1582a79d145b961145b60dbf4c5b.PrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser%2Did"] = id
    }
    return i4ca23ea1fafc22df57a5d1150889ab294f6d1582a79d145b961145b60dbf4c5b.NewPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DailyPrintUsageSummariesByPrinter the dailyPrintUsageSummariesByPrinter property
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByPrinter()(*i5ba6893546e6aaea0f4a67b2dd033135639babe968232c6fcc5163ab253c7847.DailyPrintUsageSummariesByPrinterRequestBuilder) {
    return i5ba6893546e6aaea0f4a67b2dd033135639babe968232c6fcc5163ab253c7847.NewDailyPrintUsageSummariesByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageSummariesByPrinterById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.dailyPrintUsageSummariesByPrinter.item collection
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByPrinterById(id string)(*i1453ad3549b1a27e47913c2d4a29d5625c5b5e04deb3a2a6f264f5cd24a2e8da.PrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter%2Did"] = id
    }
    return i1453ad3549b1a27e47913c2d4a29d5625c5b5e04deb3a2a6f264f5cd24a2e8da.NewPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DailyPrintUsageSummariesByUser the dailyPrintUsageSummariesByUser property
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByUser()(*i3772e364c43656cba18386bc364e131c7943e3a0b0461e3a162fb5eae513ac0b.DailyPrintUsageSummariesByUserRequestBuilder) {
    return i3772e364c43656cba18386bc364e131c7943e3a0b0461e3a162fb5eae513ac0b.NewDailyPrintUsageSummariesByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageSummariesByUserById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.dailyPrintUsageSummariesByUser.item collection
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByUserById(id string)(*iffede79819a5c0828784daecd91c1db9829729eb545a0c3b1948096142482a96.PrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser%2Did"] = id
    }
    return iffede79819a5c0828784daecd91c1db9829729eb545a0c3b1948096142482a96.NewPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property reports for print
func (m *ReportsRequestBuilder) Delete(ctx context.Context, requestConfiguration *ReportsRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceConfigurationDeviceActivity provides operations to call the deviceConfigurationDeviceActivity method.
func (m *ReportsRequestBuilder) DeviceConfigurationDeviceActivity()(*i12820273d3c44672ae92304278e27099019c4b7ff9e6dce947a674b59ff6c98e.DeviceConfigurationDeviceActivityRequestBuilder) {
    return i12820273d3c44672ae92304278e27099019c4b7ff9e6dce947a674b59ff6c98e.NewDeviceConfigurationDeviceActivityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurationUserActivity provides operations to call the deviceConfigurationUserActivity method.
func (m *ReportsRequestBuilder) DeviceConfigurationUserActivity()(*ic753e229b298240003ae8ca9aff000f32462080e1d16dc2def476877d5707d80.DeviceConfigurationUserActivityRequestBuilder) {
    return ic753e229b298240003ae8ca9aff000f32462080e1d16dc2def476877d5707d80.NewDeviceConfigurationUserActivityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get reports from print
func (m *ReportsRequestBuilder) Get(ctx context.Context, requestConfiguration *ReportsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateReportRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable), nil
}
// GetAttackSimulationRepeatOffenders provides operations to call the getAttackSimulationRepeatOffenders method.
func (m *ReportsRequestBuilder) GetAttackSimulationRepeatOffenders()(*ie5d033fba5fb2263cfd1bff0d67c57e993013c03f754b188e52f7a6cc45ba107.GetAttackSimulationRepeatOffendersRequestBuilder) {
    return ie5d033fba5fb2263cfd1bff0d67c57e993013c03f754b188e52f7a6cc45ba107.NewGetAttackSimulationRepeatOffendersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetAttackSimulationSimulationUserCoverage provides operations to call the getAttackSimulationSimulationUserCoverage method.
func (m *ReportsRequestBuilder) GetAttackSimulationSimulationUserCoverage()(*i96fe7fb8deb79d07ea5e417d63df350e76d28fdd0491378b2801b6e248d018c3.GetAttackSimulationSimulationUserCoverageRequestBuilder) {
    return i96fe7fb8deb79d07ea5e417d63df350e76d28fdd0491378b2801b6e248d018c3.NewGetAttackSimulationSimulationUserCoverageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetAttackSimulationTrainingUserCoverage provides operations to call the getAttackSimulationTrainingUserCoverage method.
func (m *ReportsRequestBuilder) GetAttackSimulationTrainingUserCoverage()(*ie8d93a68bfc5d77cb89a3bcaa58d306cf9c1eb664f150283d9e7c99bb68fdd94.GetAttackSimulationTrainingUserCoverageRequestBuilder) {
    return ie8d93a68bfc5d77cb89a3bcaa58d306cf9c1eb664f150283d9e7c99bb68fdd94.NewGetAttackSimulationTrainingUserCoverageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetAzureADApplicationSignInSummaryWithPeriod provides operations to call the getAzureADApplicationSignInSummary method.
func (m *ReportsRequestBuilder) GetAzureADApplicationSignInSummaryWithPeriod(period *string)(*i8fc0ad0d8f4c5cb82e70d74504fd7555d8235044ab2b40bc6e68738d416134ae.GetAzureADApplicationSignInSummaryWithPeriodRequestBuilder) {
    return i8fc0ad0d8f4c5cb82e70d74504fd7555d8235044ab2b40bc6e68738d416134ae.NewGetAzureADApplicationSignInSummaryWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetBrowserDistributionUserCountsWithPeriod provides operations to call the getBrowserDistributionUserCounts method.
func (m *ReportsRequestBuilder) GetBrowserDistributionUserCountsWithPeriod(period *string)(*i003173036331dffc584bd55a4a31f9eca0b7cc880c4006f990ef8ec3e90f8d61.GetBrowserDistributionUserCountsWithPeriodRequestBuilder) {
    return i003173036331dffc584bd55a4a31f9eca0b7cc880c4006f990ef8ec3e90f8d61.NewGetBrowserDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetBrowserUserCountsWithPeriod provides operations to call the getBrowserUserCounts method.
func (m *ReportsRequestBuilder) GetBrowserUserCountsWithPeriod(period *string)(*if278596fba20816d527a904fbb7b5d78e3c177aa4b09c4e15657da9d6d06a2db.GetBrowserUserCountsWithPeriodRequestBuilder) {
    return if278596fba20816d527a904fbb7b5d78e3c177aa4b09c4e15657da9d6d06a2db.NewGetBrowserUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetBrowserUserDetailWithPeriod provides operations to call the getBrowserUserDetail method.
func (m *ReportsRequestBuilder) GetBrowserUserDetailWithPeriod(period *string)(*if4dd226e0f2c5d03d114b08150ac66aef29652e888dbf65028360ee02623bdf9.GetBrowserUserDetailWithPeriodRequestBuilder) {
    return if4dd226e0f2c5d03d114b08150ac66aef29652e888dbf65028360ee02623bdf9.NewGetBrowserUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetCredentialUsageSummaryWithPeriod provides operations to call the getCredentialUsageSummary method.
func (m *ReportsRequestBuilder) GetCredentialUsageSummaryWithPeriod(period *string)(*ia58f01f9d5b233aa2565b0e4538e84e12cb2a0d586ef5a56fbbb816826c2bba2.GetCredentialUsageSummaryWithPeriodRequestBuilder) {
    return ia58f01f9d5b233aa2565b0e4538e84e12cb2a0d586ef5a56fbbb816826c2bba2.NewGetCredentialUsageSummaryWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetCredentialUserRegistrationCount provides operations to call the getCredentialUserRegistrationCount method.
func (m *ReportsRequestBuilder) GetCredentialUserRegistrationCount()(*i125e8811c6de949618151abfadcaca28ef05c938e0252591bfa83691e13c2070.GetCredentialUserRegistrationCountRequestBuilder) {
    return i125e8811c6de949618151abfadcaca28ef05c938e0252591bfa83691e13c2070.NewGetCredentialUserRegistrationCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetEmailActivityCountsWithPeriod provides operations to call the getEmailActivityCounts method.
func (m *ReportsRequestBuilder) GetEmailActivityCountsWithPeriod(period *string)(*i5c1ccbcf17345644477286c949022f98da0c3dbfc12ebc979ef7ba6edea3b759.GetEmailActivityCountsWithPeriodRequestBuilder) {
    return i5c1ccbcf17345644477286c949022f98da0c3dbfc12ebc979ef7ba6edea3b759.NewGetEmailActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailActivityUserCountsWithPeriod provides operations to call the getEmailActivityUserCounts method.
func (m *ReportsRequestBuilder) GetEmailActivityUserCountsWithPeriod(period *string)(*ia11c96eb7fc01f9081a8c45eff690c39a5c79c0b439f8e4687f3fd30ab5b707c.GetEmailActivityUserCountsWithPeriodRequestBuilder) {
    return ia11c96eb7fc01f9081a8c45eff690c39a5c79c0b439f8e4687f3fd30ab5b707c.NewGetEmailActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailActivityUserDetailWithDate provides operations to call the getEmailActivityUserDetail method.
func (m *ReportsRequestBuilder) GetEmailActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*i7ab3166690500d92e204211f333c4697c5188e65cef5efac6c698b404cea6b48.GetEmailActivityUserDetailWithDateRequestBuilder) {
    return i7ab3166690500d92e204211f333c4697c5188e65cef5efac6c698b404cea6b48.NewGetEmailActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetEmailActivityUserDetailWithPeriod provides operations to call the getEmailActivityUserDetail method.
func (m *ReportsRequestBuilder) GetEmailActivityUserDetailWithPeriod(period *string)(*i1ffd159fb8aed8eab5637f0dcf99bc4e9560b54299bc87e4f88446b96ca862f1.GetEmailActivityUserDetailWithPeriodRequestBuilder) {
    return i1ffd159fb8aed8eab5637f0dcf99bc4e9560b54299bc87e4f88446b96ca862f1.NewGetEmailActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailAppUsageAppsUserCountsWithPeriod provides operations to call the getEmailAppUsageAppsUserCounts method.
func (m *ReportsRequestBuilder) GetEmailAppUsageAppsUserCountsWithPeriod(period *string)(*i35b216d7dc42f80efd855cd81dcf7debc2bdafd7ea24b32bcb1d686238dbc76d.GetEmailAppUsageAppsUserCountsWithPeriodRequestBuilder) {
    return i35b216d7dc42f80efd855cd81dcf7debc2bdafd7ea24b32bcb1d686238dbc76d.NewGetEmailAppUsageAppsUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailAppUsageUserCountsWithPeriod provides operations to call the getEmailAppUsageUserCounts method.
func (m *ReportsRequestBuilder) GetEmailAppUsageUserCountsWithPeriod(period *string)(*i313577bbc7d677412230e87962a7a9319b6aabde4851f6c8739be336203bfa65.GetEmailAppUsageUserCountsWithPeriodRequestBuilder) {
    return i313577bbc7d677412230e87962a7a9319b6aabde4851f6c8739be336203bfa65.NewGetEmailAppUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailAppUsageUserDetailWithDate provides operations to call the getEmailAppUsageUserDetail method.
func (m *ReportsRequestBuilder) GetEmailAppUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*i6cba3dd24e78bc27f2b14a43b1489601f5f50bf4febd3101a2bdfcce9eb1c31e.GetEmailAppUsageUserDetailWithDateRequestBuilder) {
    return i6cba3dd24e78bc27f2b14a43b1489601f5f50bf4febd3101a2bdfcce9eb1c31e.NewGetEmailAppUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetEmailAppUsageUserDetailWithPeriod provides operations to call the getEmailAppUsageUserDetail method.
func (m *ReportsRequestBuilder) GetEmailAppUsageUserDetailWithPeriod(period *string)(*i306349847f73d7bf75fde4cca87b9e847d160d4298503961ee7173a877566811.GetEmailAppUsageUserDetailWithPeriodRequestBuilder) {
    return i306349847f73d7bf75fde4cca87b9e847d160d4298503961ee7173a877566811.NewGetEmailAppUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailAppUsageVersionsUserCountsWithPeriod provides operations to call the getEmailAppUsageVersionsUserCounts method.
func (m *ReportsRequestBuilder) GetEmailAppUsageVersionsUserCountsWithPeriod(period *string)(*i0ee449f331a7fd89049ddd3e1a2b7b309f269639673e8bd0aa048ffdf3f4e37b.GetEmailAppUsageVersionsUserCountsWithPeriodRequestBuilder) {
    return i0ee449f331a7fd89049ddd3e1a2b7b309f269639673e8bd0aa048ffdf3f4e37b.NewGetEmailAppUsageVersionsUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTime provides operations to call the getGroupArchivedPrintJobs method.
func (m *ReportsRequestBuilder) GetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, groupId *string, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*i8a126339410006480a905b751a907b1e7c5cd226076269437ad939d10cb7a71f.GetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i8a126339410006480a905b751a907b1e7c5cd226076269437ad939d10cb7a71f.NewGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, groupId, startDateTime);
}
// GetM365AppPlatformUserCountsWithPeriod provides operations to call the getM365AppPlatformUserCounts method.
func (m *ReportsRequestBuilder) GetM365AppPlatformUserCountsWithPeriod(period *string)(*ic8b36546821464e310499bcde1e62e440cf6796bae121c31bff3e399487901dc.GetM365AppPlatformUserCountsWithPeriodRequestBuilder) {
    return ic8b36546821464e310499bcde1e62e440cf6796bae121c31bff3e399487901dc.NewGetM365AppPlatformUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetM365AppUserCountsWithPeriod provides operations to call the getM365AppUserCounts method.
func (m *ReportsRequestBuilder) GetM365AppUserCountsWithPeriod(period *string)(*i2d772c8afb07a24a1205c9ff6577c716c80df2a273be74d3b989c999d09199da.GetM365AppUserCountsWithPeriodRequestBuilder) {
    return i2d772c8afb07a24a1205c9ff6577c716c80df2a273be74d3b989c999d09199da.NewGetM365AppUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetM365AppUserDetailWithDate provides operations to call the getM365AppUserDetail method.
func (m *ReportsRequestBuilder) GetM365AppUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*i3a0d2235d08fac82d349f0d0e02e0cdce3a0c009ad1dd8bf7208224ae85f089f.GetM365AppUserDetailWithDateRequestBuilder) {
    return i3a0d2235d08fac82d349f0d0e02e0cdce3a0c009ad1dd8bf7208224ae85f089f.NewGetM365AppUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetM365AppUserDetailWithPeriod provides operations to call the getM365AppUserDetail method.
func (m *ReportsRequestBuilder) GetM365AppUserDetailWithPeriod(period *string)(*id3ee30e7ad89e8fb7901e0c16cfe88b41322719fa7407bf70b021ab61e421a6e.GetM365AppUserDetailWithPeriodRequestBuilder) {
    return id3ee30e7ad89e8fb7901e0c16cfe88b41322719fa7407bf70b021ab61e421a6e.NewGetM365AppUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetMailboxUsageDetailWithPeriod provides operations to call the getMailboxUsageDetail method.
func (m *ReportsRequestBuilder) GetMailboxUsageDetailWithPeriod(period *string)(*icd1b1e4da27d3da9a9df7b33ef242fd4e79c8a35c3cddfe9d580b7e401250239.GetMailboxUsageDetailWithPeriodRequestBuilder) {
    return icd1b1e4da27d3da9a9df7b33ef242fd4e79c8a35c3cddfe9d580b7e401250239.NewGetMailboxUsageDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetMailboxUsageMailboxCountsWithPeriod provides operations to call the getMailboxUsageMailboxCounts method.
func (m *ReportsRequestBuilder) GetMailboxUsageMailboxCountsWithPeriod(period *string)(*i9023ee3c3a9e19cbce82a481a61ad1ced99450d1c7561a3d19d00e038450a05a.GetMailboxUsageMailboxCountsWithPeriodRequestBuilder) {
    return i9023ee3c3a9e19cbce82a481a61ad1ced99450d1c7561a3d19d00e038450a05a.NewGetMailboxUsageMailboxCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetMailboxUsageQuotaStatusMailboxCountsWithPeriod provides operations to call the getMailboxUsageQuotaStatusMailboxCounts method.
func (m *ReportsRequestBuilder) GetMailboxUsageQuotaStatusMailboxCountsWithPeriod(period *string)(*ia3c04ca51a3e25cdb708373964025e87292d13441d9e90fc357ca0a1621ebd95.GetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilder) {
    return ia3c04ca51a3e25cdb708373964025e87292d13441d9e90fc357ca0a1621ebd95.NewGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetMailboxUsageStorageWithPeriod provides operations to call the getMailboxUsageStorage method.
func (m *ReportsRequestBuilder) GetMailboxUsageStorageWithPeriod(period *string)(*icddeafe68b5aefb5586975a9974bb7370a97219ada28c815ec9b8e3b7f284748.GetMailboxUsageStorageWithPeriodRequestBuilder) {
    return icddeafe68b5aefb5586975a9974bb7370a97219ada28c815ec9b8e3b7f284748.NewGetMailboxUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365ActivationCounts provides operations to call the getOffice365ActivationCounts method.
func (m *ReportsRequestBuilder) GetOffice365ActivationCounts()(*if6320aed42bd57df6018ebfaa8490e375970ae1b65838874224d2524e3133811.GetOffice365ActivationCountsRequestBuilder) {
    return if6320aed42bd57df6018ebfaa8490e375970ae1b65838874224d2524e3133811.NewGetOffice365ActivationCountsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOffice365ActivationsUserCounts provides operations to call the getOffice365ActivationsUserCounts method.
func (m *ReportsRequestBuilder) GetOffice365ActivationsUserCounts()(*i20836d38a167d95c34498a5116a9cd82d480e78642fcca162facb20b00fdc534.GetOffice365ActivationsUserCountsRequestBuilder) {
    return i20836d38a167d95c34498a5116a9cd82d480e78642fcca162facb20b00fdc534.NewGetOffice365ActivationsUserCountsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOffice365ActivationsUserDetail provides operations to call the getOffice365ActivationsUserDetail method.
func (m *ReportsRequestBuilder) GetOffice365ActivationsUserDetail()(*i6bb9a3b51383aadb75ab3111f39327408462a143d79cb9106e695674d98c226a.GetOffice365ActivationsUserDetailRequestBuilder) {
    return i6bb9a3b51383aadb75ab3111f39327408462a143d79cb9106e695674d98c226a.NewGetOffice365ActivationsUserDetailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOffice365ActiveUserCountsWithPeriod provides operations to call the getOffice365ActiveUserCounts method.
func (m *ReportsRequestBuilder) GetOffice365ActiveUserCountsWithPeriod(period *string)(*i036a1b2a2d29c6542eb2d7482b6087baec6694c844689f95bde300288b5d9af8.GetOffice365ActiveUserCountsWithPeriodRequestBuilder) {
    return i036a1b2a2d29c6542eb2d7482b6087baec6694c844689f95bde300288b5d9af8.NewGetOffice365ActiveUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365ActiveUserDetailWithDate provides operations to call the getOffice365ActiveUserDetail method.
func (m *ReportsRequestBuilder) GetOffice365ActiveUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*icdcce25b5e3ff04e70a3cf70bfbfae9ec805346c7c065f15b4fc6b4b165bbf48.GetOffice365ActiveUserDetailWithDateRequestBuilder) {
    return icdcce25b5e3ff04e70a3cf70bfbfae9ec805346c7c065f15b4fc6b4b165bbf48.NewGetOffice365ActiveUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetOffice365ActiveUserDetailWithPeriod provides operations to call the getOffice365ActiveUserDetail method.
func (m *ReportsRequestBuilder) GetOffice365ActiveUserDetailWithPeriod(period *string)(*i1618be8d311dc94393db28762cd325c59b9ccde3cb914a388979d02c7c993fb6.GetOffice365ActiveUserDetailWithPeriodRequestBuilder) {
    return i1618be8d311dc94393db28762cd325c59b9ccde3cb914a388979d02c7c993fb6.NewGetOffice365ActiveUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365GroupsActivityCountsWithPeriod provides operations to call the getOffice365GroupsActivityCounts method.
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityCountsWithPeriod(period *string)(*i7a5b000346e54a3c0bb5a39921329be717069bb5aa6e2bb74c97770c0a1f79ce.GetOffice365GroupsActivityCountsWithPeriodRequestBuilder) {
    return i7a5b000346e54a3c0bb5a39921329be717069bb5aa6e2bb74c97770c0a1f79ce.NewGetOffice365GroupsActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365GroupsActivityDetailWithDate provides operations to call the getOffice365GroupsActivityDetail method.
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ie7a68204cc4982a6b52ad00381436561da8774cf91ceef439cd495875cd84e9d.GetOffice365GroupsActivityDetailWithDateRequestBuilder) {
    return ie7a68204cc4982a6b52ad00381436561da8774cf91ceef439cd495875cd84e9d.NewGetOffice365GroupsActivityDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetOffice365GroupsActivityDetailWithPeriod provides operations to call the getOffice365GroupsActivityDetail method.
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityDetailWithPeriod(period *string)(*i328b60d96035cb1d9bd616db71d49e8436746e43b59c2ae74edb06e1c65f6158.GetOffice365GroupsActivityDetailWithPeriodRequestBuilder) {
    return i328b60d96035cb1d9bd616db71d49e8436746e43b59c2ae74edb06e1c65f6158.NewGetOffice365GroupsActivityDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365GroupsActivityFileCountsWithPeriod provides operations to call the getOffice365GroupsActivityFileCounts method.
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityFileCountsWithPeriod(period *string)(*i7355efd26cb65bce9c72b3b63c77e2dba43919588a8e86369b208aefe72c690c.GetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder) {
    return i7355efd26cb65bce9c72b3b63c77e2dba43919588a8e86369b208aefe72c690c.NewGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365GroupsActivityGroupCountsWithPeriod provides operations to call the getOffice365GroupsActivityGroupCounts method.
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityGroupCountsWithPeriod(period *string)(*if31836807723ff5bdd6b6e4e26ed08acbebd30a643cc4b8b53069f1facbbe742.GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder) {
    return if31836807723ff5bdd6b6e4e26ed08acbebd30a643cc4b8b53069f1facbbe742.NewGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365GroupsActivityStorageWithPeriod provides operations to call the getOffice365GroupsActivityStorage method.
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityStorageWithPeriod(period *string)(*i50fd31667e7b85e595d3060b1d746877ba37f9b2c6fae47e8ee37f5c2f965515.GetOffice365GroupsActivityStorageWithPeriodRequestBuilder) {
    return i50fd31667e7b85e595d3060b1d746877ba37f9b2c6fae47e8ee37f5c2f965515.NewGetOffice365GroupsActivityStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365ServicesUserCountsWithPeriod provides operations to call the getOffice365ServicesUserCounts method.
func (m *ReportsRequestBuilder) GetOffice365ServicesUserCountsWithPeriod(period *string)(*ib179175741f9263f3b9b2d70c2568f051b6ea034ea8268e68071c8691030b092.GetOffice365ServicesUserCountsWithPeriodRequestBuilder) {
    return ib179175741f9263f3b9b2d70c2568f051b6ea034ea8268e68071c8691030b092.NewGetOffice365ServicesUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveActivityFileCountsWithPeriod provides operations to call the getOneDriveActivityFileCounts method.
func (m *ReportsRequestBuilder) GetOneDriveActivityFileCountsWithPeriod(period *string)(*i2157001b3d1627cbb758e52e666c104b24725dfdb773f94732a2c17226613e0c.GetOneDriveActivityFileCountsWithPeriodRequestBuilder) {
    return i2157001b3d1627cbb758e52e666c104b24725dfdb773f94732a2c17226613e0c.NewGetOneDriveActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveActivityUserCountsWithPeriod provides operations to call the getOneDriveActivityUserCounts method.
func (m *ReportsRequestBuilder) GetOneDriveActivityUserCountsWithPeriod(period *string)(*i9631ad5d28f0b71849e4f20d19d1958e0d952c9631900bea7802308681c1759b.GetOneDriveActivityUserCountsWithPeriodRequestBuilder) {
    return i9631ad5d28f0b71849e4f20d19d1958e0d952c9631900bea7802308681c1759b.NewGetOneDriveActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveActivityUserDetailWithDate provides operations to call the getOneDriveActivityUserDetail method.
func (m *ReportsRequestBuilder) GetOneDriveActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*iccd91022582f2d8dde0cb9b960990f411b495b6ff93387a25fe59ca8a9961562.GetOneDriveActivityUserDetailWithDateRequestBuilder) {
    return iccd91022582f2d8dde0cb9b960990f411b495b6ff93387a25fe59ca8a9961562.NewGetOneDriveActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetOneDriveActivityUserDetailWithPeriod provides operations to call the getOneDriveActivityUserDetail method.
func (m *ReportsRequestBuilder) GetOneDriveActivityUserDetailWithPeriod(period *string)(*i597d5f57aae856b81874eda3863201bedf21b61ece754b9b079304509ca2068a.GetOneDriveActivityUserDetailWithPeriodRequestBuilder) {
    return i597d5f57aae856b81874eda3863201bedf21b61ece754b9b079304509ca2068a.NewGetOneDriveActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveUsageAccountCountsWithPeriod provides operations to call the getOneDriveUsageAccountCounts method.
func (m *ReportsRequestBuilder) GetOneDriveUsageAccountCountsWithPeriod(period *string)(*ib749d3394cf27f1397536acbbae8cf02a899e5a39e5596471124bf4d71501852.GetOneDriveUsageAccountCountsWithPeriodRequestBuilder) {
    return ib749d3394cf27f1397536acbbae8cf02a899e5a39e5596471124bf4d71501852.NewGetOneDriveUsageAccountCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveUsageAccountDetailWithDate provides operations to call the getOneDriveUsageAccountDetail method.
func (m *ReportsRequestBuilder) GetOneDriveUsageAccountDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*i817285b5e7fea10651e0d4b79aa6611ca57de96e79203aa153ccde9c90f07dd1.GetOneDriveUsageAccountDetailWithDateRequestBuilder) {
    return i817285b5e7fea10651e0d4b79aa6611ca57de96e79203aa153ccde9c90f07dd1.NewGetOneDriveUsageAccountDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetOneDriveUsageAccountDetailWithPeriod provides operations to call the getOneDriveUsageAccountDetail method.
func (m *ReportsRequestBuilder) GetOneDriveUsageAccountDetailWithPeriod(period *string)(*i97727cca9b8c936e00dd03bd736857d99fcb279322f6334f8075719eee6f11b2.GetOneDriveUsageAccountDetailWithPeriodRequestBuilder) {
    return i97727cca9b8c936e00dd03bd736857d99fcb279322f6334f8075719eee6f11b2.NewGetOneDriveUsageAccountDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveUsageFileCountsWithPeriod provides operations to call the getOneDriveUsageFileCounts method.
func (m *ReportsRequestBuilder) GetOneDriveUsageFileCountsWithPeriod(period *string)(*iaa160176f699a99c070a1f1f49a0887c07919a5c689d21c8bed8f52de4a89023.GetOneDriveUsageFileCountsWithPeriodRequestBuilder) {
    return iaa160176f699a99c070a1f1f49a0887c07919a5c689d21c8bed8f52de4a89023.NewGetOneDriveUsageFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveUsageStorageWithPeriod provides operations to call the getOneDriveUsageStorage method.
func (m *ReportsRequestBuilder) GetOneDriveUsageStorageWithPeriod(period *string)(*i1a614f493010260020309e60ff2e8606cd018fda1d411bd62662caf405dbfa5b.GetOneDriveUsageStorageWithPeriodRequestBuilder) {
    return i1a614f493010260020309e60ff2e8606cd018fda1d411bd62662caf405dbfa5b.NewGetOneDriveUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTime provides operations to call the getPrinterArchivedPrintJobs method.
func (m *ReportsRequestBuilder) GetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, printerId *string, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*i4f829313dc734307b5536ec2ce241c78a78f218e561a61fdf90fce1e9092a56a.GetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i4f829313dc734307b5536ec2ce241c78a78f218e561a61fdf90fce1e9092a56a.NewGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, printerId, startDateTime);
}
// GetRelyingPartyDetailedSummaryWithPeriod provides operations to call the getRelyingPartyDetailedSummary method.
func (m *ReportsRequestBuilder) GetRelyingPartyDetailedSummaryWithPeriod(period *string)(*iad427df3a804c5ea6bf927439cbace855bd9edf35dec361d7cc896c86a3c3e76.GetRelyingPartyDetailedSummaryWithPeriodRequestBuilder) {
    return iad427df3a804c5ea6bf927439cbace855bd9edf35dec361d7cc896c86a3c3e76.NewGetRelyingPartyDetailedSummaryWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointActivityFileCountsWithPeriod provides operations to call the getSharePointActivityFileCounts method.
func (m *ReportsRequestBuilder) GetSharePointActivityFileCountsWithPeriod(period *string)(*icc5ddf4fe31a3f4d96b22a2a2e64d31f88f217b06f59946a62eccd7032aa763d.GetSharePointActivityFileCountsWithPeriodRequestBuilder) {
    return icc5ddf4fe31a3f4d96b22a2a2e64d31f88f217b06f59946a62eccd7032aa763d.NewGetSharePointActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointActivityPagesWithPeriod provides operations to call the getSharePointActivityPages method.
func (m *ReportsRequestBuilder) GetSharePointActivityPagesWithPeriod(period *string)(*ibf5abd7cd37fc0ab2d9ca27fab459ff2a81a89b886033d344b5d0d36bfa92d6a.GetSharePointActivityPagesWithPeriodRequestBuilder) {
    return ibf5abd7cd37fc0ab2d9ca27fab459ff2a81a89b886033d344b5d0d36bfa92d6a.NewGetSharePointActivityPagesWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointActivityUserCountsWithPeriod provides operations to call the getSharePointActivityUserCounts method.
func (m *ReportsRequestBuilder) GetSharePointActivityUserCountsWithPeriod(period *string)(*i1a7aad10a2334fe03a5643720cc6b428cdff146a4fc62c7a50a78f70f9a36bc4.GetSharePointActivityUserCountsWithPeriodRequestBuilder) {
    return i1a7aad10a2334fe03a5643720cc6b428cdff146a4fc62c7a50a78f70f9a36bc4.NewGetSharePointActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointActivityUserDetailWithDate provides operations to call the getSharePointActivityUserDetail method.
func (m *ReportsRequestBuilder) GetSharePointActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*i72fab10b39cc1890375f583eea81a5d54930bcd29915bbd6d1b82e9b92ae1e64.GetSharePointActivityUserDetailWithDateRequestBuilder) {
    return i72fab10b39cc1890375f583eea81a5d54930bcd29915bbd6d1b82e9b92ae1e64.NewGetSharePointActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetSharePointActivityUserDetailWithPeriod provides operations to call the getSharePointActivityUserDetail method.
func (m *ReportsRequestBuilder) GetSharePointActivityUserDetailWithPeriod(period *string)(*ibe3b90074f3fd6bd15c064326b107884d8dcff2d4f739ea9969038710d070c78.GetSharePointActivityUserDetailWithPeriodRequestBuilder) {
    return ibe3b90074f3fd6bd15c064326b107884d8dcff2d4f739ea9969038710d070c78.NewGetSharePointActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointSiteUsageDetailWithDate provides operations to call the getSharePointSiteUsageDetail method.
func (m *ReportsRequestBuilder) GetSharePointSiteUsageDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*i2a4862a2c04d7e8705aebac57aeab2e0bf3268230862ae87191095bb620b0409.GetSharePointSiteUsageDetailWithDateRequestBuilder) {
    return i2a4862a2c04d7e8705aebac57aeab2e0bf3268230862ae87191095bb620b0409.NewGetSharePointSiteUsageDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetSharePointSiteUsageDetailWithPeriod provides operations to call the getSharePointSiteUsageDetail method.
func (m *ReportsRequestBuilder) GetSharePointSiteUsageDetailWithPeriod(period *string)(*iae0b95416e4f0115ded4c5680f5308d53f4ab1191730356055fa78cb04fdb3a3.GetSharePointSiteUsageDetailWithPeriodRequestBuilder) {
    return iae0b95416e4f0115ded4c5680f5308d53f4ab1191730356055fa78cb04fdb3a3.NewGetSharePointSiteUsageDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointSiteUsageFileCountsWithPeriod provides operations to call the getSharePointSiteUsageFileCounts method.
func (m *ReportsRequestBuilder) GetSharePointSiteUsageFileCountsWithPeriod(period *string)(*i5df2c32f3d48e07033ebe1c8ef2dd7d264320343ff53c690d8fa229633336de2.GetSharePointSiteUsageFileCountsWithPeriodRequestBuilder) {
    return i5df2c32f3d48e07033ebe1c8ef2dd7d264320343ff53c690d8fa229633336de2.NewGetSharePointSiteUsageFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointSiteUsagePagesWithPeriod provides operations to call the getSharePointSiteUsagePages method.
func (m *ReportsRequestBuilder) GetSharePointSiteUsagePagesWithPeriod(period *string)(*ic65cd1a979952b5550b39c84cadcef5f570c546d773d487f9bbe744c4a2ead72.GetSharePointSiteUsagePagesWithPeriodRequestBuilder) {
    return ic65cd1a979952b5550b39c84cadcef5f570c546d773d487f9bbe744c4a2ead72.NewGetSharePointSiteUsagePagesWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointSiteUsageSiteCountsWithPeriod provides operations to call the getSharePointSiteUsageSiteCounts method.
func (m *ReportsRequestBuilder) GetSharePointSiteUsageSiteCountsWithPeriod(period *string)(*if420f5f31b1d54d7e335630814dcebdb92a9d2ca57cafb5c1aee864810f1d4c2.GetSharePointSiteUsageSiteCountsWithPeriodRequestBuilder) {
    return if420f5f31b1d54d7e335630814dcebdb92a9d2ca57cafb5c1aee864810f1d4c2.NewGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointSiteUsageStorageWithPeriod provides operations to call the getSharePointSiteUsageStorage method.
func (m *ReportsRequestBuilder) GetSharePointSiteUsageStorageWithPeriod(period *string)(*i9fd0210ad4ae49215efc7cf22136268ddeef1c0bff45a84f042e4597a495e2c2.GetSharePointSiteUsageStorageWithPeriodRequestBuilder) {
    return i9fd0210ad4ae49215efc7cf22136268ddeef1c0bff45a84f042e4597a495e2c2.NewGetSharePointSiteUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessActivityCountsWithPeriod provides operations to call the getSkypeForBusinessActivityCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessActivityCountsWithPeriod(period *string)(*i7b88d33e4fa5702328c128d0552680d05cd2414a20ea062fc43f3458e04682c7.GetSkypeForBusinessActivityCountsWithPeriodRequestBuilder) {
    return i7b88d33e4fa5702328c128d0552680d05cd2414a20ea062fc43f3458e04682c7.NewGetSkypeForBusinessActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessActivityUserCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessActivityUserCountsWithPeriod(period *string)(*i21b3b996d3b537f9331bc504ac14bc8a80a52a412b2e26c73bfcc4ab966b11b3.GetSkypeForBusinessActivityUserCountsWithPeriodRequestBuilder) {
    return i21b3b996d3b537f9331bc504ac14bc8a80a52a412b2e26c73bfcc4ab966b11b3.NewGetSkypeForBusinessActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessActivityUserDetailWithDate provides operations to call the getSkypeForBusinessActivityUserDetail method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ia290e37e65f4257fc2cba6045f8fae24a9fcb0902933df2bbcd9ea8ac77cb2ef.GetSkypeForBusinessActivityUserDetailWithDateRequestBuilder) {
    return ia290e37e65f4257fc2cba6045f8fae24a9fcb0902933df2bbcd9ea8ac77cb2ef.NewGetSkypeForBusinessActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetSkypeForBusinessActivityUserDetailWithPeriod provides operations to call the getSkypeForBusinessActivityUserDetail method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessActivityUserDetailWithPeriod(period *string)(*ib9134a9409b3aeb31e9bbde8b92be81f399cc5028fd2d006006539e98ca706e8.GetSkypeForBusinessActivityUserDetailWithPeriodRequestBuilder) {
    return ib9134a9409b3aeb31e9bbde8b92be81f399cc5028fd2d006006539e98ca706e8.NewGetSkypeForBusinessActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriod provides operations to call the getSkypeForBusinessDeviceUsageDistributionUserCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriod(period *string)(*icb9fbfdf1a799b2b262a2d5e2e5889891e3c531a929c01dc751359da88307334.GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return icb9fbfdf1a799b2b262a2d5e2e5889891e3c531a929c01dc751359da88307334.NewGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessDeviceUsageUserCountsWithPeriod provides operations to call the getSkypeForBusinessDeviceUsageUserCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessDeviceUsageUserCountsWithPeriod(period *string)(*ia337a32890551b1081a608ae43abf61c4497dfd42048dd808bb7cf09832281c3.GetSkypeForBusinessDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return ia337a32890551b1081a608ae43abf61c4497dfd42048dd808bb7cf09832281c3.NewGetSkypeForBusinessDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessDeviceUsageUserDetailWithDate provides operations to call the getSkypeForBusinessDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessDeviceUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ic05da5b8d4cf94558ff5262e866399feadd45bc394ba1441e1fc14d600fde57a.GetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilder) {
    return ic05da5b8d4cf94558ff5262e866399feadd45bc394ba1441e1fc14d600fde57a.NewGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetSkypeForBusinessDeviceUsageUserDetailWithPeriod provides operations to call the getSkypeForBusinessDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessDeviceUsageUserDetailWithPeriod(period *string)(*i04f6ee74f0733278b8c6e9678d187269b49dcb665a9958b4ae6257f0aaec8dbf.GetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return i04f6ee74f0733278b8c6e9678d187269b49dcb665a9958b4ae6257f0aaec8dbf.NewGetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessOrganizerActivityCountsWithPeriod provides operations to call the getSkypeForBusinessOrganizerActivityCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessOrganizerActivityCountsWithPeriod(period *string)(*i224d112f4527ff133ce9d06395ccae8994e1f895836fdc4374f4a129a5a06bed.GetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilder) {
    return i224d112f4527ff133ce9d06395ccae8994e1f895836fdc4374f4a129a5a06bed.NewGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod provides operations to call the getSkypeForBusinessOrganizerActivityMinuteCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod(period *string)(*i81aca7de3573f1e665c768b347600ba5c3fb9fede4a50466b75c3b9b5586bf78.GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriodRequestBuilder) {
    return i81aca7de3573f1e665c768b347600ba5c3fb9fede4a50466b75c3b9b5586bf78.NewGetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessOrganizerActivityUserCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod(period *string)(*i435102bd0023ed78a36db74d804d78f55bf2d1ef197f516b037bb3b82f517c08.GetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilder) {
    return i435102bd0023ed78a36db74d804d78f55bf2d1ef197f516b037bb3b82f517c08.NewGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessParticipantActivityCountsWithPeriod provides operations to call the getSkypeForBusinessParticipantActivityCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessParticipantActivityCountsWithPeriod(period *string)(*i328e3d55288de86c60aa3add79cd83a5c047661caef06da2cff1cef6c8598b34.GetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder) {
    return i328e3d55288de86c60aa3add79cd83a5c047661caef06da2cff1cef6c8598b34.NewGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessParticipantActivityMinuteCountsWithPeriod provides operations to call the getSkypeForBusinessParticipantActivityMinuteCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessParticipantActivityMinuteCountsWithPeriod(period *string)(*i420730d2f754eb6fcc9ad1b5310922c60bb4e377cbdde02aaac29b3fa0f0b2e5.GetSkypeForBusinessParticipantActivityMinuteCountsWithPeriodRequestBuilder) {
    return i420730d2f754eb6fcc9ad1b5310922c60bb4e377cbdde02aaac29b3fa0f0b2e5.NewGetSkypeForBusinessParticipantActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessParticipantActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessParticipantActivityUserCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessParticipantActivityUserCountsWithPeriod(period *string)(*ic7a209730cd8537604c6db9508d277bc2e323d048bba634bc00a0ff07184baac.GetSkypeForBusinessParticipantActivityUserCountsWithPeriodRequestBuilder) {
    return ic7a209730cd8537604c6db9508d277bc2e323d048bba634bc00a0ff07184baac.NewGetSkypeForBusinessParticipantActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod provides operations to call the getSkypeForBusinessPeerToPeerActivityCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod(period *string)(*i423884002844f09964fb45ec36b7ef31192751d8bb7a78a54fb126d4206b1025.GetSkypeForBusinessPeerToPeerActivityCountsWithPeriodRequestBuilder) {
    return i423884002844f09964fb45ec36b7ef31192751d8bb7a78a54fb126d4206b1025.NewGetSkypeForBusinessPeerToPeerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod provides operations to call the getSkypeForBusinessPeerToPeerActivityMinuteCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod(period *string)(*iae029900a0da9f48c34f75ee5d793b8dc3931fe12b0f7a985040c1d5ec7c5523.GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriodRequestBuilder) {
    return iae029900a0da9f48c34f75ee5d793b8dc3931fe12b0f7a985040c1d5ec7c5523.NewGetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessPeerToPeerActivityUserCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod(period *string)(*id7392fa684625fd57bfdfc7c7dfbc56d0065c1cc7093f8a6d32a219c2ba1dfac.GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriodRequestBuilder) {
    return id7392fa684625fd57bfdfc7c7dfbc56d0065c1cc7093f8a6d32a219c2ba1dfac.NewGetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageDistributionTotalUserCounts method.
func (m *ReportsRequestBuilder) GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod(period *string)(*i95b95e9f1f7725e5d0fd693531d9686b616e391472cd4b61b5f11da6c9cc5526.GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder) {
    return i95b95e9f1f7725e5d0fd693531d9686b616e391472cd4b61b5f11da6c9cc5526.NewGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsDeviceUsageDistributionUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageDistributionUserCounts method.
func (m *ReportsRequestBuilder) GetTeamsDeviceUsageDistributionUserCountsWithPeriod(period *string)(*i2499892f567f8728fe2de683d0f00e72080cfd5ffca87cc0785ef89eb730a931.GetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return i2499892f567f8728fe2de683d0f00e72080cfd5ffca87cc0785ef89eb730a931.NewGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsDeviceUsageTotalUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageTotalUserCounts method.
func (m *ReportsRequestBuilder) GetTeamsDeviceUsageTotalUserCountsWithPeriod(period *string)(*i4add8d70f1a45f4e5d237fb6c9877de552324641d4f73a6bcd482c6db9801fcd.GetTeamsDeviceUsageTotalUserCountsWithPeriodRequestBuilder) {
    return i4add8d70f1a45f4e5d237fb6c9877de552324641d4f73a6bcd482c6db9801fcd.NewGetTeamsDeviceUsageTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsDeviceUsageUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageUserCounts method.
func (m *ReportsRequestBuilder) GetTeamsDeviceUsageUserCountsWithPeriod(period *string)(*i808e335fa07739775bf1a553e8afacccf25cd8047be1de0c119ce5caa8554634.GetTeamsDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return i808e335fa07739775bf1a553e8afacccf25cd8047be1de0c119ce5caa8554634.NewGetTeamsDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsDeviceUsageUserDetailWithDate provides operations to call the getTeamsDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) GetTeamsDeviceUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ie822df7440dd5b07f2cb5b9ae611cdc7ec11b9262d220f5682bd2958e02a643e.GetTeamsDeviceUsageUserDetailWithDateRequestBuilder) {
    return ie822df7440dd5b07f2cb5b9ae611cdc7ec11b9262d220f5682bd2958e02a643e.NewGetTeamsDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetTeamsDeviceUsageUserDetailWithPeriod provides operations to call the getTeamsDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) GetTeamsDeviceUsageUserDetailWithPeriod(period *string)(*i4d7d4fe8dd8ea3539a0a8ad2f9a8e5d271857ba237dce7b4a8ca536a616ec435.GetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return i4d7d4fe8dd8ea3539a0a8ad2f9a8e5d271857ba237dce7b4a8ca536a616ec435.NewGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsTeamActivityCountsWithPeriod provides operations to call the getTeamsTeamActivityCounts method.
func (m *ReportsRequestBuilder) GetTeamsTeamActivityCountsWithPeriod(period *string)(*idf5d61f783c212fb53a8df5f35707485128386131f3ae007235e92b0eb00404c.GetTeamsTeamActivityCountsWithPeriodRequestBuilder) {
    return idf5d61f783c212fb53a8df5f35707485128386131f3ae007235e92b0eb00404c.NewGetTeamsTeamActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsTeamActivityDetailWithDate provides operations to call the getTeamsTeamActivityDetail method.
func (m *ReportsRequestBuilder) GetTeamsTeamActivityDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ia135a492a1ea4fe1e4e4616ad09038e519f2792c9d98bbdca0bab0ae0ea58926.GetTeamsTeamActivityDetailWithDateRequestBuilder) {
    return ia135a492a1ea4fe1e4e4616ad09038e519f2792c9d98bbdca0bab0ae0ea58926.NewGetTeamsTeamActivityDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetTeamsTeamActivityDetailWithPeriod provides operations to call the getTeamsTeamActivityDetail method.
func (m *ReportsRequestBuilder) GetTeamsTeamActivityDetailWithPeriod(period *string)(*ibb4d5055652a4d2b25c6f9515b9c6f5b7c8a1bb709ad42af6daac709a178858f.GetTeamsTeamActivityDetailWithPeriodRequestBuilder) {
    return ibb4d5055652a4d2b25c6f9515b9c6f5b7c8a1bb709ad42af6daac709a178858f.NewGetTeamsTeamActivityDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsTeamActivityDistributionCountsWithPeriod provides operations to call the getTeamsTeamActivityDistributionCounts method.
func (m *ReportsRequestBuilder) GetTeamsTeamActivityDistributionCountsWithPeriod(period *string)(*i932396fbe8ecd019f26fa2a7bce06d747e0bebe4e79cc653e70943aa11c9355d.GetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilder) {
    return i932396fbe8ecd019f26fa2a7bce06d747e0bebe4e79cc653e70943aa11c9355d.NewGetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityCountsWithPeriod provides operations to call the getTeamsUserActivityCounts method.
func (m *ReportsRequestBuilder) GetTeamsUserActivityCountsWithPeriod(period *string)(*i567056fba35e6edab31f5aa9994989fd7742a56af8e01ec0f240852321374891.GetTeamsUserActivityCountsWithPeriodRequestBuilder) {
    return i567056fba35e6edab31f5aa9994989fd7742a56af8e01ec0f240852321374891.NewGetTeamsUserActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityDistributionTotalUserCountsWithPeriod provides operations to call the getTeamsUserActivityDistributionTotalUserCounts method.
func (m *ReportsRequestBuilder) GetTeamsUserActivityDistributionTotalUserCountsWithPeriod(period *string)(*i1775513b7b2429eb61cc9982fc3d0c155badba7bb5c8436afb785b645183be6a.GetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder) {
    return i1775513b7b2429eb61cc9982fc3d0c155badba7bb5c8436afb785b645183be6a.NewGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityDistributionUserCountsWithPeriod provides operations to call the getTeamsUserActivityDistributionUserCounts method.
func (m *ReportsRequestBuilder) GetTeamsUserActivityDistributionUserCountsWithPeriod(period *string)(*i088887eceaa935280d9592f89bc4b0a045fd924ff8d44ab2b855e031094ea540.GetTeamsUserActivityDistributionUserCountsWithPeriodRequestBuilder) {
    return i088887eceaa935280d9592f89bc4b0a045fd924ff8d44ab2b855e031094ea540.NewGetTeamsUserActivityDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityTotalCountsWithPeriod provides operations to call the getTeamsUserActivityTotalCounts method.
func (m *ReportsRequestBuilder) GetTeamsUserActivityTotalCountsWithPeriod(period *string)(*i18f58919f6d5882740586016a8bc9807f1c45848f93ad165acc8577a78caaf9d.GetTeamsUserActivityTotalCountsWithPeriodRequestBuilder) {
    return i18f58919f6d5882740586016a8bc9807f1c45848f93ad165acc8577a78caaf9d.NewGetTeamsUserActivityTotalCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityTotalDistributionCountsWithPeriod provides operations to call the getTeamsUserActivityTotalDistributionCounts method.
func (m *ReportsRequestBuilder) GetTeamsUserActivityTotalDistributionCountsWithPeriod(period *string)(*i8a1d439fc44307034a531d7f5d9324bb3c5e5b3c14bbacf473e28c1319852fcf.GetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilder) {
    return i8a1d439fc44307034a531d7f5d9324bb3c5e5b3c14bbacf473e28c1319852fcf.NewGetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityTotalUserCountsWithPeriod provides operations to call the getTeamsUserActivityTotalUserCounts method.
func (m *ReportsRequestBuilder) GetTeamsUserActivityTotalUserCountsWithPeriod(period *string)(*i4113a715b496e73819d650e2951dc88df01049bb7ee89faa36b869ccf6a2f71e.GetTeamsUserActivityTotalUserCountsWithPeriodRequestBuilder) {
    return i4113a715b496e73819d650e2951dc88df01049bb7ee89faa36b869ccf6a2f71e.NewGetTeamsUserActivityTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityUserCountsWithPeriod provides operations to call the getTeamsUserActivityUserCounts method.
func (m *ReportsRequestBuilder) GetTeamsUserActivityUserCountsWithPeriod(period *string)(*i75ebd72e88d0574c533352a054cd82b52553ab8eb695b568d8aaa890ce02579f.GetTeamsUserActivityUserCountsWithPeriodRequestBuilder) {
    return i75ebd72e88d0574c533352a054cd82b52553ab8eb695b568d8aaa890ce02579f.NewGetTeamsUserActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityUserDetailWithDate provides operations to call the getTeamsUserActivityUserDetail method.
func (m *ReportsRequestBuilder) GetTeamsUserActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*i9b2d2e6f90f5f9da6a8f69e6d4e9b3e43b80147df7cbcd17ae490b6d7b0a4194.GetTeamsUserActivityUserDetailWithDateRequestBuilder) {
    return i9b2d2e6f90f5f9da6a8f69e6d4e9b3e43b80147df7cbcd17ae490b6d7b0a4194.NewGetTeamsUserActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetTeamsUserActivityUserDetailWithPeriod provides operations to call the getTeamsUserActivityUserDetail method.
func (m *ReportsRequestBuilder) GetTeamsUserActivityUserDetailWithPeriod(period *string)(*i859de82968bd758e08878a67874d4d3a4d19758bd7bcbe1ad402adc066dbaee2.GetTeamsUserActivityUserDetailWithPeriodRequestBuilder) {
    return i859de82968bd758e08878a67874d4d3a4d19758bd7bcbe1ad402adc066dbaee2.NewGetTeamsUserActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime provides operations to call the getUserArchivedPrintJobs method.
func (m *ReportsRequestBuilder) GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, userId *string)(*iba529b34d22e6d4619a98e3a64864d4ce62443496139f8d961a28f3388b1c06b.GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return iba529b34d22e6d4619a98e3a64864d4ce62443496139f8d961a28f3388b1c06b.NewGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime, userId);
}
// GetYammerActivityCountsWithPeriod provides operations to call the getYammerActivityCounts method.
func (m *ReportsRequestBuilder) GetYammerActivityCountsWithPeriod(period *string)(*iaaafb6c998e8e1eedbeeb89216037d37a8bc2089fc1568d394cdf1e95cec3c91.GetYammerActivityCountsWithPeriodRequestBuilder) {
    return iaaafb6c998e8e1eedbeeb89216037d37a8bc2089fc1568d394cdf1e95cec3c91.NewGetYammerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerActivityUserCountsWithPeriod provides operations to call the getYammerActivityUserCounts method.
func (m *ReportsRequestBuilder) GetYammerActivityUserCountsWithPeriod(period *string)(*ica973266b26481209852c53969fd134273a48e20a686d98f2c7e4ac2bb684ece.GetYammerActivityUserCountsWithPeriodRequestBuilder) {
    return ica973266b26481209852c53969fd134273a48e20a686d98f2c7e4ac2bb684ece.NewGetYammerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerActivityUserDetailWithDate provides operations to call the getYammerActivityUserDetail method.
func (m *ReportsRequestBuilder) GetYammerActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ifda5935468b5ae9cd912cd8886c9d4cca1e9ab52a6c082d9530ad21d01d3c3ad.GetYammerActivityUserDetailWithDateRequestBuilder) {
    return ifda5935468b5ae9cd912cd8886c9d4cca1e9ab52a6c082d9530ad21d01d3c3ad.NewGetYammerActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetYammerActivityUserDetailWithPeriod provides operations to call the getYammerActivityUserDetail method.
func (m *ReportsRequestBuilder) GetYammerActivityUserDetailWithPeriod(period *string)(*i983e4c56868ed790407309ebf34264a39588d5f319fbc8ea49bc989533975765.GetYammerActivityUserDetailWithPeriodRequestBuilder) {
    return i983e4c56868ed790407309ebf34264a39588d5f319fbc8ea49bc989533975765.NewGetYammerActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerDeviceUsageDistributionUserCountsWithPeriod provides operations to call the getYammerDeviceUsageDistributionUserCounts method.
func (m *ReportsRequestBuilder) GetYammerDeviceUsageDistributionUserCountsWithPeriod(period *string)(*i5125d764ebc3d13de453dc8ae6109633072221c35bf64d9e7655b7445242f381.GetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return i5125d764ebc3d13de453dc8ae6109633072221c35bf64d9e7655b7445242f381.NewGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerDeviceUsageUserCountsWithPeriod provides operations to call the getYammerDeviceUsageUserCounts method.
func (m *ReportsRequestBuilder) GetYammerDeviceUsageUserCountsWithPeriod(period *string)(*i1b39eea80211466ef73aee8b2f1a50b08a56f3dce15eda5a18ede70e235f90d7.GetYammerDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return i1b39eea80211466ef73aee8b2f1a50b08a56f3dce15eda5a18ede70e235f90d7.NewGetYammerDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerDeviceUsageUserDetailWithDate provides operations to call the getYammerDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) GetYammerDeviceUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ia46b9d8d0e59a725b189c3202c5343f9f35bf2a8a04d4785295f7a71e6be9590.GetYammerDeviceUsageUserDetailWithDateRequestBuilder) {
    return ia46b9d8d0e59a725b189c3202c5343f9f35bf2a8a04d4785295f7a71e6be9590.NewGetYammerDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetYammerDeviceUsageUserDetailWithPeriod provides operations to call the getYammerDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) GetYammerDeviceUsageUserDetailWithPeriod(period *string)(*ie591696fdc474e68dec24db3dde3496dbb37168bbafe4d0a2b02ab16f890232e.GetYammerDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return ie591696fdc474e68dec24db3dde3496dbb37168bbafe4d0a2b02ab16f890232e.NewGetYammerDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerGroupsActivityCountsWithPeriod provides operations to call the getYammerGroupsActivityCounts method.
func (m *ReportsRequestBuilder) GetYammerGroupsActivityCountsWithPeriod(period *string)(*i3e7d149707f2ddd6910123604eff6dd78480edcb3f05321bb44037a85d733380.GetYammerGroupsActivityCountsWithPeriodRequestBuilder) {
    return i3e7d149707f2ddd6910123604eff6dd78480edcb3f05321bb44037a85d733380.NewGetYammerGroupsActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerGroupsActivityDetailWithDate provides operations to call the getYammerGroupsActivityDetail method.
func (m *ReportsRequestBuilder) GetYammerGroupsActivityDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ib66df0db90f0ef258b3536c6659fe0d6ede8fd889a2c38cf8c63f72fb5f1ffaa.GetYammerGroupsActivityDetailWithDateRequestBuilder) {
    return ib66df0db90f0ef258b3536c6659fe0d6ede8fd889a2c38cf8c63f72fb5f1ffaa.NewGetYammerGroupsActivityDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetYammerGroupsActivityDetailWithPeriod provides operations to call the getYammerGroupsActivityDetail method.
func (m *ReportsRequestBuilder) GetYammerGroupsActivityDetailWithPeriod(period *string)(*ie050dab9fab3553491bfd68cf2523de38e35cfaa538b49fdcf301f61cb13e9e0.GetYammerGroupsActivityDetailWithPeriodRequestBuilder) {
    return ie050dab9fab3553491bfd68cf2523de38e35cfaa538b49fdcf301f61cb13e9e0.NewGetYammerGroupsActivityDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerGroupsActivityGroupCountsWithPeriod provides operations to call the getYammerGroupsActivityGroupCounts method.
func (m *ReportsRequestBuilder) GetYammerGroupsActivityGroupCountsWithPeriod(period *string)(*i103d63d4193b9f3ab75fe6b66304aaeb64366749ff445bc88c2dc5beed1269cb.GetYammerGroupsActivityGroupCountsWithPeriodRequestBuilder) {
    return i103d63d4193b9f3ab75fe6b66304aaeb64366749ff445bc88c2dc5beed1269cb.NewGetYammerGroupsActivityGroupCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// ManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipToken provides operations to call the managedDeviceEnrollmentAbandonmentDetails method.
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipToken(filter *string, skip *int32, skipToken *string, top *int32)(*id4c53d2ba82691363ac26d457be4a4cdcf13b6e128592549cf1224910bd3751c.ManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    return id4c53d2ba82691363ac26d457be4a4cdcf13b6e128592549cf1224910bd3751c.NewManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, filter, skip, skipToken, top);
}
// ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipToken provides operations to call the managedDeviceEnrollmentAbandonmentSummary method.
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipToken(filter *string, skip *int32, skipToken *string, top *int32)(*ibfa03aafca1026d6cb462cc2b4a03d52cfcd50b0378e62a671000fb1c4cae155.ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    return ibfa03aafca1026d6cb462cc2b4a03d52cfcd50b0378e62a671000fb1c4cae155.NewManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, filter, skip, skipToken, top);
}
// ManagedDeviceEnrollmentFailureDetails provides operations to call the managedDeviceEnrollmentFailureDetails method.
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentFailureDetails()(*i2e11d2e753b93ae2e19efcae2b9d79d3efa476d4ee469d6db9c173f35ef3b620.ManagedDeviceEnrollmentFailureDetailsRequestBuilder) {
    return i2e11d2e753b93ae2e19efcae2b9d79d3efa476d4ee469d6db9c173f35ef3b620.NewManagedDeviceEnrollmentFailureDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipToken provides operations to call the managedDeviceEnrollmentFailureDetails method.
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipToken(filter *string, skip *int32, skipToken *string, top *int32)(*i4c6dc125c665e98f5e3a0f5b0343f5fdacc0d0a7e3ae60637e7e73e22e4f0ca4.ManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    return i4c6dc125c665e98f5e3a0f5b0343f5fdacc0d0a7e3ae60637e7e73e22e4f0ca4.NewManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, filter, skip, skipToken, top);
}
// ManagedDeviceEnrollmentFailureTrends provides operations to call the managedDeviceEnrollmentFailureTrends method.
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentFailureTrends()(*i804fbe3d687c5d168bcbde36b31d895d3817ace737cf5ae5896736def6e2624e.ManagedDeviceEnrollmentFailureTrendsRequestBuilder) {
    return i804fbe3d687c5d168bcbde36b31d895d3817ace737cf5ae5896736def6e2624e.NewManagedDeviceEnrollmentFailureTrendsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceEnrollmentTopFailures provides operations to call the managedDeviceEnrollmentTopFailures method.
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentTopFailures()(*ic2e883868a9a28d28dfcc1e662d0e24885b4e5da00486835cbbe6d165c5798a6.ManagedDeviceEnrollmentTopFailuresRequestBuilder) {
    return ic2e883868a9a28d28dfcc1e662d0e24885b4e5da00486835cbbe6d165c5798a6.NewManagedDeviceEnrollmentTopFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceEnrollmentTopFailuresWithPeriod provides operations to call the managedDeviceEnrollmentTopFailures method.
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentTopFailuresWithPeriod(period *string)(*ia3e34a08513d2e3ca7ffcfad8b20f3382618d2a7552421e981f109b27cad2170.ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder) {
    return ia3e34a08513d2e3ca7ffcfad8b20f3382618d2a7552421e981f109b27cad2170.NewManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// MonthlyPrintUsageByPrinter the monthlyPrintUsageByPrinter property
func (m *ReportsRequestBuilder) MonthlyPrintUsageByPrinter()(*i330d7d86bcb93d59ef14fc1257795533ce121643f4e8becd8884b57bc63e7b4b.MonthlyPrintUsageByPrinterRequestBuilder) {
    return i330d7d86bcb93d59ef14fc1257795533ce121643f4e8becd8884b57bc63e7b4b.NewMonthlyPrintUsageByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MonthlyPrintUsageByPrinterById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.monthlyPrintUsageByPrinter.item collection
func (m *ReportsRequestBuilder) MonthlyPrintUsageByPrinterById(id string)(*i55a12b578425084e14afca840f9f8ad56ad8296a76d84139beb67b9619a01ae2.PrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter%2Did"] = id
    }
    return i55a12b578425084e14afca840f9f8ad56ad8296a76d84139beb67b9619a01ae2.NewPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MonthlyPrintUsageByUser the monthlyPrintUsageByUser property
func (m *ReportsRequestBuilder) MonthlyPrintUsageByUser()(*i6532f3214bfaa5d2aac58a9d9ffe42fb56e5e2242405e98165287c807365bcda.MonthlyPrintUsageByUserRequestBuilder) {
    return i6532f3214bfaa5d2aac58a9d9ffe42fb56e5e2242405e98165287c807365bcda.NewMonthlyPrintUsageByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MonthlyPrintUsageByUserById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.monthlyPrintUsageByUser.item collection
func (m *ReportsRequestBuilder) MonthlyPrintUsageByUserById(id string)(*i1702ca3f0c998e29512738841125456c0b4e81357e5f7ca9f4176443002d28d5.PrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser%2Did"] = id
    }
    return i1702ca3f0c998e29512738841125456c0b4e81357e5f7ca9f4176443002d28d5.NewPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MonthlyPrintUsageSummariesByPrinter the monthlyPrintUsageSummariesByPrinter property
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByPrinter()(*i27c74777b0efbee67b16cbd84335b5228b74a547d6ffe5e338e3c1c6cc55c61b.MonthlyPrintUsageSummariesByPrinterRequestBuilder) {
    return i27c74777b0efbee67b16cbd84335b5228b74a547d6ffe5e338e3c1c6cc55c61b.NewMonthlyPrintUsageSummariesByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MonthlyPrintUsageSummariesByPrinterById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.monthlyPrintUsageSummariesByPrinter.item collection
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByPrinterById(id string)(*ic1e71dc755b571267687f63b13312b758bc574de61144bfa50cd81af3b6d3626.PrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter%2Did"] = id
    }
    return ic1e71dc755b571267687f63b13312b758bc574de61144bfa50cd81af3b6d3626.NewPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MonthlyPrintUsageSummariesByUser the monthlyPrintUsageSummariesByUser property
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByUser()(*i40dc3bbc1c58ea07c4c8b16942aaeca4da00eb3514241eef66edcd213ba084f7.MonthlyPrintUsageSummariesByUserRequestBuilder) {
    return i40dc3bbc1c58ea07c4c8b16942aaeca4da00eb3514241eef66edcd213ba084f7.NewMonthlyPrintUsageSummariesByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MonthlyPrintUsageSummariesByUserById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.monthlyPrintUsageSummariesByUser.item collection
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByUserById(id string)(*i2268eb94e47ea5d81b666c3310292b49733c18b0196efe2957b868503a9da7ff.PrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser%2Did"] = id
    }
    return i2268eb94e47ea5d81b666c3310292b49733c18b0196efe2957b868503a9da7ff.NewPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property reports in print
func (m *ReportsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable, requestConfiguration *ReportsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateReportRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable), nil
}
// Security the security property
func (m *ReportsRequestBuilder) Security()(*ifece25d39559e83e159712abc947b9fa9ef645ac39026b276f4c512910960f48.SecurityRequestBuilder) {
    return ifece25d39559e83e159712abc947b9fa9ef645ac39026b276f4c512910960f48.NewSecurityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserCredentialUsageDetails the userCredentialUsageDetails property
func (m *ReportsRequestBuilder) UserCredentialUsageDetails()(*i254c294d6b358b3854c3ee45f61d16c1237e0477952e917b034bbb0d8a6802b1.UserCredentialUsageDetailsRequestBuilder) {
    return i254c294d6b358b3854c3ee45f61d16c1237e0477952e917b034bbb0d8a6802b1.NewUserCredentialUsageDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserCredentialUsageDetailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.userCredentialUsageDetails.item collection
func (m *ReportsRequestBuilder) UserCredentialUsageDetailsById(id string)(*i4f1c6062b781770b431d75b41deb0a2ac227468563ebac0623bef14f3efe351d.UserCredentialUsageDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userCredentialUsageDetails%2Did"] = id
    }
    return i4f1c6062b781770b431d75b41deb0a2ac227468563ebac0623bef14f3efe351d.NewUserCredentialUsageDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
