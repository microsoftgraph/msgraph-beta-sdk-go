package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i007cac44b1ad1a90ea316a1c0c51e045017f34e30cc3b64342ae3537775aaca4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/activateserviceplan"
    i00eee5f4e45952d7af6d7dbf02a5397275cbfa69444a9cd54a450f1cc9c5baf6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives"
    i028859f982626d2101b8d8b68caff0f04ef34796aed805e993e21fefb9152760 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/getmanageddeviceswithappfailures"
    i099d32137e6ab567a080a6d574abebbf5b3e8b218fa6fda8a2a0fce52a347329 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/wipemanagedappregistrationbydevicetag"
    i09b3ed8449f0eb014b1a1c391d9057c99dca194b7d9cee6d67e587c2b4d66363 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/findrooms"
    i0bb45fabb973cc1726799c26601b61fb5d4ceef85239469bf8a48d255ddcab77 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events"
    i0fbc6bb6d8e908a7211206db1e3f715af4ae676f093155141c801afcbfd4a468 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/deviceenrollmentconfigurations"
    i10209655736fb29f6fe332f57065da684a4bb7e75afca36c8b0cf9215d9600cf "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/exportpersonaldata"
    i163ca73c2d1426f49f134d87416999c0644a5bf7f78d9912f8098a9785dc6ec1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/getloggedonmanageddevices"
    i1b6f500d8aa60be2a2a1e7586a07867f5e80bc6da6a793388efc074011aa9f1b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/getmanagedapppolicies"
    i1c8bab6c329a747e2f1a7bc6de2b493e83d072621195eaffdf5348edafc8d0c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/chats"
    i1d3f5d5cc5119c3710246df1436375c7f067bf4348a4cf5c5decdf0052f3b667 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/unblockmanagedapps"
    i1e2ed1761df19ea44d2d90f6fdfba02f95eb595495e7804f87126e72ee9345e3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/reminderviewwithstartdatetimewithenddatetime"
    i1e8437839d244c8d93e242e55290d8df8b03ca76cf9b009f9b9a19ef5d5922ac "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports"
    i1faa710c227b8e4b8d7abef5e9611e6604d3f4265db3da621d2ddd654a3a33f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/people"
    i21282294d2d1bb2a0ab74695618e809021209e6abac85e175b427b95d4891cc5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders"
    i235dd448b749bb3a1187ec0970cdeb5c226f7aee1c9e6c72733a9d1a15182ffe "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups"
    i23fd320d20dc894e875e4225a9c0031362d1c191fcba8ad855a39c641e46a084 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/wipemanagedappregistrationsbydevicetag"
    i25e2af6a5262c1ae60708cc14abab6c2ec1e90563db9f056ec390327d5d23afe "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/checkmembergroups"
    i2a5b8dfcb3b418ee303ee5f5b01aebd3209cdfa9122c52e81bd0a9e16a726f9d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/presence"
    i2c42211f5ece49fe404bcc5fe4998da5209b6fd5a9e6ef29c937593bb6d65f32 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview"
    i2d035ce91a466eb27cdd0da2fa61e30e90dd33e1f08803145645a2bbbd70f1df "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/approvals"
    i31922605eeb11427d2cf37ddfee7f2740fc6e17cd21d0e9ee6f853fe1b3e4632 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar"
    i35ae79fec43abc52b415d013c5ba088edde69bf589391525fec20b1f062c22a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof"
    i35b254e720ef3a77057a7b6736d2d9841a2d361b2b095e7a189fd80751427060 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/reprocesslicenseassignment"
    i396aa46f4b0d835f500bf48230128cbf9b0b575b29db3501c7da923e86cbb886 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/createdobjects"
    i3aa2a5875223f6b19560c0365147205688fff76fe30e6bbf724b3b25f7b00a43 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/ownedobjects"
    i3be6b130d8bf10a5df104376e126ee2c3584be236591382ff88662fd6a574158 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances"
    i3d9c184215c67e8ba108a3159afbd942d4eaa93fc6fdec9593468297b72346e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/agreementacceptances"
    i3e6a9a935ee43b6bd89f378e203c324f1d517ccc9d1c30202303eafc13d5bb25 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drive"
    i3eaf8bc51db6f4bd95448da3951002ebce4db39125b63eaa1cb9987f99da858f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/windowsinformationprotectiondeviceregistrations"
    i44184067b9a5d3386d10a8d3d78383c99a1f94c6db7a1c78fecc6277422e9842 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/geteffectivedeviceenrollmentconfigurations"
    i441da951c572a925ed6c67355cefdd1ddc3a435432abc802841187bcd7831f08 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups"
    i44dc4d21a51a483da0eb7ac0bc6c38fa258f1e982bb140f0bd9282a87d39015c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/managedappregistrations"
    i49c0a0951cef8133c800eac5f2bf2ddd282c66d3a73fc9f9c652fc532a3331be "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/extensions"
    i4c553c01d10dc5ea91f5e8443a25576af37822a1a9e3b20c39f7d69613ca1a5d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote"
    i4cd644380b450eedb89634c647645afa6ede7ee4df886587de78ccf69b3a9a42 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices"
    i4e66b80d61d13c58a5325be0d73998d6c528f982ee63fc876b66fdffd4ec5473 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/todo"
    i4e8ad44991b185083366e107199ff13f6438a07cafd58ae36e0e862692bf9826 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedteams"
    i52ea864f44c605939329148c5a95626ead1b771dff88c6072de654caee1b281e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/getmanageddeviceswithfailedorpendingapps"
    i53e4d773c077f49b4590da13b904e6a502ef1f8f43b7a7fdfaba256a667df083 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/licensedetails"
    i55ac7591042853708e32005320f053aac164b91987142fe88da44995054ce50a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/approleassignedresources"
    i58ea06623816d885f9dc5730c5fd341b9dc32ca959bf1ffe118d49d00bdfbc62 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/photos"
    i5c0e1410fbc34cac508ae33c37b1e9dd9a4ca3540ad1525e02f20329f8a264f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/findmeetingtimes"
    i5c911b116a9e04c78f811a8141c00da4569a7dc37a1d8590843bc9d14e30377e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/notifications"
    i5f232e4d902d95d06ec20d5128cce05b56d53758222f6b28e307a6b909082dac "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/insights"
    i5f9e5aa5c251065cc59079614790107bc87c70c65435cea5312a9ad92e27cad3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/photo"
    i633ead85dd41f6f9d94e9cf71e43f604b18195cef363d66f743366cdb820fc8b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/getmembergroups"
    i64d17e545a98d6600890650331d10cc8ac9bafd9f1a9888ceee3f7c9932f9d68 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/planner"
    i6545b8651bf39e658fd8159fa1123831d6f44a855626360fac6873d3a90581c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/revokesigninsessions"
    i67e4cd2b85517c9aa8f5256b8ca64c3daab00417eaeedb8528d78204b2257eb1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/approleassignments"
    i6bc016bb2cb82b3c7d804b1443cbf5f98e5cb3ae65246c4f7658c00f6dde5db6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders"
    i6c6abf877b483e321abb71bac38e5c8c592f7ee6915c9ad28447c321b2979b4d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/exportdeviceandappmanagementdata"
    i6d8c6f6a8d72bfe9260edd2b3126823191856ab8f78dd9cd5a1c01fb0fff9c47 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/restore"
    i6f957db73ce02e759548f419eb79a2bfdfb04d6c5206523cd438181cc5619df8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof"
    i71ece4d9b62444d7cd9edeac6c2b5bdaf0f35379d154d089d602c3d246239739 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/owneddevices"
    i7358087738ea2e8de75562a191c6ccb72c9a24d79c1f9a33de77f140cbb5ce74 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mobileapptroubleshootingevents"
    i73641486333af15741c085f42cd9ba7b087800533a2f2a8781b636a965272506 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/inferenceclassification"
    i751d3667ec77ce3a4b6544375d44e850f38ac8db37e3237ce66245e949ea015e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/getmemberobjects"
    i77129a8685763d465ea62e5de870203f0d38a41c496b1f718e0c695d6df2b42c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/profile"
    i7984c072c770ec17745d737738e8ab409b4f442b183221a451997db094cd8b13 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/settings"
    i7d4d710e1c8a8fc885d164dc49223af3ba44bf6a062cfa0019953387e5b65644 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contacts"
    i7d78da96ae198e74cdcd846ebd00d688295a8b346c939e4e49d56b88b5f72711 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/removealldevicesfrommanagement"
    i7f0c66e611b8af22dffef0886d0dc8468dc8cac8bb0b577016f3d4b1daf70528 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/invalidateallrefreshtokens"
    i807859e90e02be96c7951593aa65cbf418e7a30d0f62e3a47b8382296d35e2d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/tasks"
    i87f894b87c2947e4dad88e845ac91e8a8ddf1ca5ba6b426384ab8873fe897913 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/translateexchangeids"
    i89f9ebf7c167c75aff01142f1e2b30d94fcb3ccf3d4af1752e74dbc3f979ecd6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/getmailtips"
    i8df6ede5cb15274ffb3a3341b8e1b560c67d6eca03c4eafb47c23266ca7b8809 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/checkmemberobjects"
    i92c37b67685c58d9f60f7203b77a560478fb28fb9753bb4bee0d7e274eeb3e16 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/assignlicense"
    i933c9dccaf05a6a7fcf066da46cae8d6c4a0940e841426651ace8702209f1b9d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/findroomswithroomlist"
    i992f3c4694021035f11bae1cb737eb6020f9d1221bc7676d2685a415cab74221 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection"
    i9e9fd6de702827c484559c38e811567b4bfa805d681aca1a1d7d8d8ba07aae9c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/exportdeviceandappmanagementdatawithskipwithtop"
    i9f79453a6cb43e9e5e68087f1d6e774414d8ba83449733d805265bf8cb3d0b5d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/registereddevices"
    i9fa6f9d397d0318d3e3fc1137d90e8302a95a72c11d6237f67fc7bd8dc1625b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devicemanagementtroubleshootingevents"
    ia1e1519710e771d8df5e22da39eeb50d37180582bb2065916e5cd0703a4d1cce "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivereports"
    ia6d6def3248379e9169fe56c62736c8c3d2b618252c7aac3b61a8424fce68894 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/cloudpcs"
    iaf2011a15870df19ddf40701bc4ca449214192204d76149e38fd507c87344ec5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/activities"
    ib2e97a67ea231c420ac9fa60c240e9a1314dc844268c406a4c1600c693ad6811 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices"
    ib674e5ebc7e9a35ea83a7a23737e2e850d21d009460f455398f4235cd71eb2d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/oauth2permissiongrants"
    ib9ea7c893835d89864cc7cd939d24f579ba191ab540031694fe5202498a4240c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook"
    ice23d35988441664e0a80c4309242575aeb7cc7c84c7bfc6e9ff86ce92c7cde5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manager"
    ice64ec2e30acd089e9604f4d076891112981966d167a6e899c5e617a13771654 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/wipemanagedappregistrationsbyazureaddeviceid"
    id5b765e74db4b3dcebe232c27216355fdc0568a1094952e7fc7cbdc858f3d81a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/findroomlists"
    id6953e096a9515e55975655319314bf3955019e0bded78bb5cb6fb32eb4e8e94 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/sendmail"
    id904363322fca79446fa0f0f2231760c6649433886f1c8e79503338b004b288c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/followedsites"
    idaeea547224afeb26277747dcbdab987622f0aec250f6cb9dcebbc9e6d05df1e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/getmanagedappdiagnosticstatuses"
    iddc68c193cc623cc513e0675dadbd063a05b9d5ce8584eb5a7e8ebf24faa98e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/usagerights"
    idddb11dcacd5585b4b6cb080f05aa0ce5c62d39e09bc0ffcd012d5c8c731810b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onlinemeetings"
    ide4da4ea5e9d5b4a3ed2b815f5032fc4321b390d6e1c0ba686af0a35ec886a84 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mobileappintentandstates"
    ie33da1fb1c72ec4e3eff642e23951718af609226626243167de6d0dd6df287f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/changepassword"
    ieac84d2e3111366253e033fed8ced42323b1927cd610b6c93eb3eea6d066fd20 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/wipeandblockmanagedapps"
    ieafc69423dcb93faca91a98c5fff1ec743e967e077b14aad34d9518731effa7f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars"
    ieee72b0b162c07d3dde8ae1950af04718a5801d0d3e6575e0cba2faa3f56157a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/ismanagedappuserblocked"
    iefb47a3635416080293f661822201727ab9fba4b375bd5a3f60f0b9a80cca8b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication"
    if0acf31f23813f467eba440f3275db3949170732f975822128618d2ec5b19f15 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/appconsentrequestsforapproval"
    if3e915183aa6002a59a5c947830462a91684160d445bda34aa2c4cf6886b82b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/scopedrolememberof"
    if429d108f0ce6b9f644fc8ece21d69b6a892d9788590f39dc3d9d21bda91a159 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/teamwork"
    if6f18d24cca69003f155fe07812d8a9ade476b4b7cb0560b8869ba314a7a79aa "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/analytics"
    ifc1c29cd3428806f56c209014f087cf657f124d2250271baf855225bc6fb7c3d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/security"
    ifc52bce07b756c55dabe5cd8f6a0d8c749809ddcc3831f5bded6c944d449dbd8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages"
    i0796e8f7b9f803857a438202c35d9fccd548d4779e6f2a341287440bf41acc70 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item"
    i11a0d65a9c2024ff5c04a9b3cb79486d6f2277859e51094a083537f2289c0a27 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item"
    i1a9eb6fb24429a79e1e7397bd609e240ae9eef7f551f559e57f6c0050aa81ec7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item"
    i1fb04953ffbe8174524af5da51d2e85d0602b9611bf6179631d63cd31623aa9c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item"
    i20d5bc3e0beaac35f1cb30b0824d071deea9c98aa2da0ea95797426ec3391a8b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/deviceenrollmentconfigurations/item"
    i2c5057cc6c2035899f3804d4bce915570738d9af43d8e9a3f1ea425e7c0130ff "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivereports/item"
    i2cf542b96fcffbafb31190c7c7e24f2b652b0249d3c7d0e6deaf71e333abf373 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item"
    i322200ffec4334c71536e3a74319b0afb2071fae07124a70501b4a91e5d3efc2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item"
    i404e4499afbbdf92e8ee2c947bd2f30954ca765dc8d5deeda689c119c7e7eff7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/registereddevices/item"
    i41bcf6e933152d51a0364a3166ccb57d9e6cc7316937216a235e9f7a7973b36c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/licensedetails/item"
    i4d1bb8d1a6e444b2e99d643949896393d13341c23c9e59170d7cf747778a84f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devicemanagementtroubleshootingevents/item"
    i532345ba6b7192bf733f1b4b3de662c47c1e7e0d3989701bcdc0d8f857dbe580 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/extensions/item"
    i53afa1534ef724ece2eabf4449fabc9c4904d986ecd4d6a0c11e3784f58eaab4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item"
    i5d09ee5a147e62e93172ba3c3003b1e37657e5078803e7fa2be2756f303a9607 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/notifications/item"
    i662c61ababf046e658d661a256ad38f5e55026354d35305f3f7fbd8b59a76654 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/agreementacceptances/item"
    i66ea18aba758406b3b710707db096aa499334ab8141ff00b5817c0160d62f849 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contacts/item"
    i7f7b81e49212c0fd4aff0640f7a76daf8d10816c9da8bbbf24c63c6d3002dac7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/approleassignments/item"
    i7fdae8345bbd328abce3eca8f17ca05dfb12b3a6a3adcb8b7b394f3697581dae "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mobileapptroubleshootingevents/item"
    i84c9cc6d33fc548d6b6b6210a170d3ce1e97e608f9af2ff36d66a133895c1e8b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/ownedobjects/item"
    i86e9fa8da4e357b5840f59b78cd49a5abce31f7a90096cc8b85047571e5a2b06 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item"
    i89586fbadaf2720611403fc97537d705bca9891589c65eb565dbae3f293f06c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/owneddevices/item"
    i8a37b6fd0dca54e1e3129d6fa85360b47ae2206831aad7d8bf6c3823b1fb540b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item"
    i8a8732491a2ce7d94fc7e324e6a38276076b912c30ea0094d5fb671dd11250f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/createdobjects/item"
    i8bc97b715d2b1f741be45e9db03b6f18317426bb77532976ccdc40d366389a11 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/chats/item"
    i8c655ab8b24fc37660a4329441f6c9936205ed7192659a7b410b5c1b127f0ebd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/activities/item"
    i8e9bfc4860807cf4727d9197394851201c63af67494fa725e469b158152f82c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onlinemeetings/item"
    i92f204ee1ec183e477ad6aae362cad151b598b7439eda0c58c4750c295bc84e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages/item"
    i9b0b4ddaf41e0ea5ab6a02cd93c94b2fff19122c19dfa3cf7586e383ecd47e8f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/cloudpcs/item"
    i9bb18abda2da0890104c07ce4b11694592a3ff71b3436063a29752ff07e0f525 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/followedsites/item"
    i9e371fd6357698c3cd300fbb8012addd3ec298d5df1be3dfa58412fb63c3cae8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mobileappintentandstates/item"
    i9ff5415484e41f91043b46ae2894ad1a1ddead74f139bbcb66a0ea970caba35b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedteams/item"
    ia16782564e6f61796daddf42e21afd9855b9a1c1c41eef4e65e545763ac7c8d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/people/item"
    iabe64332e133a9af5a6f5aafde3a47ea985f385ae58a2f524f86fc1a3ef2f539 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item"
    ib1c924a2a83336c2eee6c85a7c0edc6e41999bab4aa8e2db2755b64356506e43 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item"
    ib39beceac49e7dd17ed1c0885887dae465a45afa34b3eebb70aa5fc08ca691e7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/photos/item"
    ib9adfa359ccdf30972cf1f2c9fb0c0a8edde337e58fe26f0a95800e7ce3ba9db "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/usagerights/item"
    iba97b8ef76c5efeab4d738316d6e0d7c1266515ff3632bf4789418105475f2bd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/windowsinformationprotectiondeviceregistrations/item"
    ic60383690980d3eba5846d285791ffa2fb94867c1850c6c8da9aaa03aca3c008 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/appconsentrequestsforapproval/item"
    id15759ba2587bc2a0cac47e8d7da5e4bd208929d38688edb78df7122421e531f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item"
    id62ce646a00a4421b260562c6eb4464e0a6e23540ac62847efa7af52e9578392 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/managedappregistrations/item"
    idbb158da6d78d676ea0c3d14f6ae7004f016a1bdc89a3da430f82196981f35af "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item"
    ideab25898675089d3e174a0e5f0bde76aa4e227661b4518a256695e6a0500360 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/oauth2permissiongrants/item"
    ie132362db0baae23ad9567ff4032a62c165d8386308f8baec97ecbb60c10a021 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/scopedrolememberof/item"
    if35b9a09bc16de74d0bfce124707aacc1fdc8b5743fe1048d957f0e359ebd013 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/approleassignedresources/item"
    ifeea6837d24ad8765a5cb2092fe2e730b4403b36aa4e813a9cdfdb831a8f7974 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/approvals/item"
)

// UserItemRequestBuilder provides operations to manage the collection of user entities.
type UserItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UserItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserItemRequestBuilderGetQueryParameters retrieve the properties and relationships of user object. This operation returns by default only a subset of the more commonly used properties for each user. These _default_ properties are noted in the Properties section. To get properties that are _not_ returned by default, do a GET operation for the user and specify the properties in a `$select` OData query option. Because the **user** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in a **user** instance.
type UserItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserItemRequestBuilderGetQueryParameters
}
// UserItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivateServicePlan the activateServicePlan property
func (m *UserItemRequestBuilder) ActivateServicePlan()(*i007cac44b1ad1a90ea316a1c0c51e045017f34e30cc3b64342ae3537775aaca4.ActivateServicePlanRequestBuilder) {
    return i007cac44b1ad1a90ea316a1c0c51e045017f34e30cc3b64342ae3537775aaca4.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Activities the activities property
func (m *UserItemRequestBuilder) Activities()(*iaf2011a15870df19ddf40701bc4ca449214192204d76149e38fd507c87344ec5.ActivitiesRequestBuilder) {
    return iaf2011a15870df19ddf40701bc4ca449214192204d76149e38fd507c87344ec5.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.activities.item collection
func (m *UserItemRequestBuilder) ActivitiesById(id string)(*i8c655ab8b24fc37660a4329441f6c9936205ed7192659a7b410b5c1b127f0ebd.UserActivityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userActivity%2Did"] = id
    }
    return i8c655ab8b24fc37660a4329441f6c9936205ed7192659a7b410b5c1b127f0ebd.NewUserActivityItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AgreementAcceptances the agreementAcceptances property
func (m *UserItemRequestBuilder) AgreementAcceptances()(*i3d9c184215c67e8ba108a3159afbd942d4eaa93fc6fdec9593468297b72346e5.AgreementAcceptancesRequestBuilder) {
    return i3d9c184215c67e8ba108a3159afbd942d4eaa93fc6fdec9593468297b72346e5.NewAgreementAcceptancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgreementAcceptancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.agreementAcceptances.item collection
func (m *UserItemRequestBuilder) AgreementAcceptancesById(id string)(*i662c61ababf046e658d661a256ad38f5e55026354d35305f3f7fbd8b59a76654.AgreementAcceptanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreementAcceptance%2Did"] = id
    }
    return i662c61ababf046e658d661a256ad38f5e55026354d35305f3f7fbd8b59a76654.NewAgreementAcceptanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics the analytics property
func (m *UserItemRequestBuilder) Analytics()(*if6f18d24cca69003f155fe07812d8a9ade476b4b7cb0560b8869ba314a7a79aa.AnalyticsRequestBuilder) {
    return if6f18d24cca69003f155fe07812d8a9ade476b4b7cb0560b8869ba314a7a79aa.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppConsentRequestsForApproval the appConsentRequestsForApproval property
func (m *UserItemRequestBuilder) AppConsentRequestsForApproval()(*if0acf31f23813f467eba440f3275db3949170732f975822128618d2ec5b19f15.AppConsentRequestsForApprovalRequestBuilder) {
    return if0acf31f23813f467eba440f3275db3949170732f975822128618d2ec5b19f15.NewAppConsentRequestsForApprovalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppConsentRequestsForApprovalById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.appConsentRequestsForApproval.item collection
func (m *UserItemRequestBuilder) AppConsentRequestsForApprovalById(id string)(*ic60383690980d3eba5846d285791ffa2fb94867c1850c6c8da9aaa03aca3c008.AppConsentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appConsentRequest%2Did"] = id
    }
    return ic60383690980d3eba5846d285791ffa2fb94867c1850c6c8da9aaa03aca3c008.NewAppConsentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AppRoleAssignedResources the appRoleAssignedResources property
func (m *UserItemRequestBuilder) AppRoleAssignedResources()(*i55ac7591042853708e32005320f053aac164b91987142fe88da44995054ce50a.AppRoleAssignedResourcesRequestBuilder) {
    return i55ac7591042853708e32005320f053aac164b91987142fe88da44995054ce50a.NewAppRoleAssignedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignedResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.appRoleAssignedResources.item collection
func (m *UserItemRequestBuilder) AppRoleAssignedResourcesById(id string)(*if35b9a09bc16de74d0bfce124707aacc1fdc8b5743fe1048d957f0e359ebd013.ServicePrincipalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["servicePrincipal%2Did"] = id
    }
    return if35b9a09bc16de74d0bfce124707aacc1fdc8b5743fe1048d957f0e359ebd013.NewServicePrincipalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AppRoleAssignments the appRoleAssignments property
func (m *UserItemRequestBuilder) AppRoleAssignments()(*i67e4cd2b85517c9aa8f5256b8ca64c3daab00417eaeedb8528d78204b2257eb1.AppRoleAssignmentsRequestBuilder) {
    return i67e4cd2b85517c9aa8f5256b8ca64c3daab00417eaeedb8528d78204b2257eb1.NewAppRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.appRoleAssignments.item collection
func (m *UserItemRequestBuilder) AppRoleAssignmentsById(id string)(*i7f7b81e49212c0fd4aff0640f7a76daf8d10816c9da8bbbf24c63c6d3002dac7.AppRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appRoleAssignment%2Did"] = id
    }
    return i7f7b81e49212c0fd4aff0640f7a76daf8d10816c9da8bbbf24c63c6d3002dac7.NewAppRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Approvals the approvals property
func (m *UserItemRequestBuilder) Approvals()(*i2d035ce91a466eb27cdd0da2fa61e30e90dd33e1f08803145645a2bbbd70f1df.ApprovalsRequestBuilder) {
    return i2d035ce91a466eb27cdd0da2fa61e30e90dd33e1f08803145645a2bbbd70f1df.NewApprovalsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApprovalsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.approvals.item collection
func (m *UserItemRequestBuilder) ApprovalsById(id string)(*ifeea6837d24ad8765a5cb2092fe2e730b4403b36aa4e813a9cdfdb831a8f7974.ApprovalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["approval%2Did"] = id
    }
    return ifeea6837d24ad8765a5cb2092fe2e730b4403b36aa4e813a9cdfdb831a8f7974.NewApprovalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserItemRequestBuilder) AssignLicense()(*i92c37b67685c58d9f60f7203b77a560478fb28fb9753bb4bee0d7e274eeb3e16.AssignLicenseRequestBuilder) {
    return i92c37b67685c58d9f60f7203b77a560478fb28fb9753bb4bee0d7e274eeb3e16.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Authentication the authentication property
func (m *UserItemRequestBuilder) Authentication()(*iefb47a3635416080293f661822201727ab9fba4b375bd5a3f60f0b9a80cca8b0.AuthenticationRequestBuilder) {
    return iefb47a3635416080293f661822201727ab9fba4b375bd5a3f60f0b9a80cca8b0.NewAuthenticationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Calendar the calendar property
func (m *UserItemRequestBuilder) Calendar()(*i31922605eeb11427d2cf37ddfee7f2740fc6e17cd21d0e9ee6f853fe1b3e4632.CalendarRequestBuilder) {
    return i31922605eeb11427d2cf37ddfee7f2740fc6e17cd21d0e9ee6f853fe1b3e4632.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarGroups the calendarGroups property
func (m *UserItemRequestBuilder) CalendarGroups()(*i441da951c572a925ed6c67355cefdd1ddc3a435432abc802841187bcd7831f08.CalendarGroupsRequestBuilder) {
    return i441da951c572a925ed6c67355cefdd1ddc3a435432abc802841187bcd7831f08.NewCalendarGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item collection
func (m *UserItemRequestBuilder) CalendarGroupsById(id string)(*i322200ffec4334c71536e3a74319b0afb2071fae07124a70501b4a91e5d3efc2.CalendarGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarGroup%2Did"] = id
    }
    return i322200ffec4334c71536e3a74319b0afb2071fae07124a70501b4a91e5d3efc2.NewCalendarGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendars the calendars property
func (m *UserItemRequestBuilder) Calendars()(*ieafc69423dcb93faca91a98c5fff1ec743e967e077b14aad34d9518731effa7f.CalendarsRequestBuilder) {
    return ieafc69423dcb93faca91a98c5fff1ec743e967e077b14aad34d9518731effa7f.NewCalendarsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item collection
func (m *UserItemRequestBuilder) CalendarsById(id string)(*i1a9eb6fb24429a79e1e7397bd609e240ae9eef7f551f559e57f6c0050aa81ec7.CalendarItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendar%2Did"] = id
    }
    return i1a9eb6fb24429a79e1e7397bd609e240ae9eef7f551f559e57f6c0050aa81ec7.NewCalendarItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CalendarView the calendarView property
func (m *UserItemRequestBuilder) CalendarView()(*i2c42211f5ece49fe404bcc5fe4998da5209b6fd5a9e6ef29c937593bb6d65f32.CalendarViewRequestBuilder) {
    return i2c42211f5ece49fe404bcc5fe4998da5209b6fd5a9e6ef29c937593bb6d65f32.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarView.item collection
func (m *UserItemRequestBuilder) CalendarViewById(id string)(*i86e9fa8da4e357b5840f59b78cd49a5abce31f7a90096cc8b85047571e5a2b06.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return i86e9fa8da4e357b5840f59b78cd49a5abce31f7a90096cc8b85047571e5a2b06.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserItemRequestBuilder) ChangePassword()(*ie33da1fb1c72ec4e3eff642e23951718af609226626243167de6d0dd6df287f4.ChangePasswordRequestBuilder) {
    return ie33da1fb1c72ec4e3eff642e23951718af609226626243167de6d0dd6df287f4.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Chats the chats property
func (m *UserItemRequestBuilder) Chats()(*i1c8bab6c329a747e2f1a7bc6de2b493e83d072621195eaffdf5348edafc8d0c7.ChatsRequestBuilder) {
    return i1c8bab6c329a747e2f1a7bc6de2b493e83d072621195eaffdf5348edafc8d0c7.NewChatsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChatsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.chats.item collection
func (m *UserItemRequestBuilder) ChatsById(id string)(*i8bc97b715d2b1f741be45e9db03b6f18317426bb77532976ccdc40d366389a11.ChatItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chat%2Did"] = id
    }
    return i8bc97b715d2b1f741be45e9db03b6f18317426bb77532976ccdc40d366389a11.NewChatItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserItemRequestBuilder) CheckMemberGroups()(*i25e2af6a5262c1ae60708cc14abab6c2ec1e90563db9f056ec390327d5d23afe.CheckMemberGroupsRequestBuilder) {
    return i25e2af6a5262c1ae60708cc14abab6c2ec1e90563db9f056ec390327d5d23afe.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserItemRequestBuilder) CheckMemberObjects()(*i8df6ede5cb15274ffb3a3341b8e1b560c67d6eca03c4eafb47c23266ca7b8809.CheckMemberObjectsRequestBuilder) {
    return i8df6ede5cb15274ffb3a3341b8e1b560c67d6eca03c4eafb47c23266ca7b8809.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudPCs the cloudPCs property
func (m *UserItemRequestBuilder) CloudPCs()(*ia6d6def3248379e9169fe56c62736c8c3d2b618252c7aac3b61a8424fce68894.CloudPCsRequestBuilder) {
    return ia6d6def3248379e9169fe56c62736c8c3d2b618252c7aac3b61a8424fce68894.NewCloudPCsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudPCsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.cloudPCs.item collection
func (m *UserItemRequestBuilder) CloudPCsById(id string)(*i9b0b4ddaf41e0ea5ab6a02cd93c94b2fff19122c19dfa3cf7586e383ecd47e8f.CloudPCItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPC%2Did"] = id
    }
    return i9b0b4ddaf41e0ea5ab6a02cd93c94b2fff19122c19dfa3cf7586e383ecd47e8f.NewCloudPCItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewUserItemRequestBuilderInternal instantiates a new UserItemRequestBuilder and sets the default values.
func NewUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserItemRequestBuilder) {
    m := &UserItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserItemRequestBuilder instantiates a new UserItemRequestBuilder and sets the default values.
func NewUserItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContactFolders the contactFolders property
func (m *UserItemRequestBuilder) ContactFolders()(*i6bc016bb2cb82b3c7d804b1443cbf5f98e5cb3ae65246c4f7658c00f6dde5db6.ContactFoldersRequestBuilder) {
    return i6bc016bb2cb82b3c7d804b1443cbf5f98e5cb3ae65246c4f7658c00f6dde5db6.NewContactFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactFoldersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.contactFolders.item collection
func (m *UserItemRequestBuilder) ContactFoldersById(id string)(*ib1c924a2a83336c2eee6c85a7c0edc6e41999bab4aa8e2db2755b64356506e43.ContactFolderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contactFolder%2Did"] = id
    }
    return ib1c924a2a83336c2eee6c85a7c0edc6e41999bab4aa8e2db2755b64356506e43.NewContactFolderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Contacts the contacts property
func (m *UserItemRequestBuilder) Contacts()(*i7d4d710e1c8a8fc885d164dc49223af3ba44bf6a062cfa0019953387e5b65644.ContactsRequestBuilder) {
    return i7d4d710e1c8a8fc885d164dc49223af3ba44bf6a062cfa0019953387e5b65644.NewContactsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.contacts.item collection
func (m *UserItemRequestBuilder) ContactsById(id string)(*i66ea18aba758406b3b710707db096aa499334ab8141ff00b5817c0160d62f849.ContactItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contact%2Did"] = id
    }
    return i66ea18aba758406b3b710707db096aa499334ab8141ff00b5817c0160d62f849.NewContactItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete user.   When deleted, user resources are moved to a temporary container and can be restored within 30 days.  After that time, they are permanently deleted.  To learn more, see deletedItems.
func (m *UserItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete user.   When deleted, user resources are moved to a temporary container and can be restored within 30 days.  After that time, they are permanently deleted.  To learn more, see deletedItems.
func (m *UserItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *UserItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatedObjects the createdObjects property
func (m *UserItemRequestBuilder) CreatedObjects()(*i396aa46f4b0d835f500bf48230128cbf9b0b575b29db3501c7da923e86cbb886.CreatedObjectsRequestBuilder) {
    return i396aa46f4b0d835f500bf48230128cbf9b0b575b29db3501c7da923e86cbb886.NewCreatedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatedObjectsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.createdObjects.item collection
func (m *UserItemRequestBuilder) CreatedObjectsById(id string)(*i8a8732491a2ce7d94fc7e324e6a38276076b912c30ea0094d5fb671dd11250f2.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i8a8732491a2ce7d94fc7e324e6a38276076b912c30ea0094d5fb671dd11250f2.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateGetRequestInformation retrieve the properties and relationships of user object. This operation returns by default only a subset of the more commonly used properties for each user. These _default_ properties are noted in the Properties section. To get properties that are _not_ returned by default, do a GET operation for the user and specify the properties in a `$select` OData query option. Because the **user** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in a **user** instance.
func (m *UserItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration retrieve the properties and relationships of user object. This operation returns by default only a subset of the more commonly used properties for each user. These _default_ properties are noted in the Properties section. To get properties that are _not_ returned by default, do a GET operation for the user and specify the properties in a `$select` OData query option. Because the **user** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in a **user** instance.
func (m *UserItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *UserItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the properties of a user object. Not all properties can be updated by Member or Guest users with their default permissions without Administrator roles. Compare member and guest default permissions to see properties they can manage.
func (m *UserItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the properties of a user object. Not all properties can be updated by Member or Guest users with their default permissions without Administrator roles. Compare member and guest default permissions to see properties they can manage.
func (m *UserItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, requestConfiguration *UserItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete user.   When deleted, user resources are moved to a temporary container and can be restored within 30 days.  After that time, they are permanently deleted.  To learn more, see deletedItems.
func (m *UserItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete user.   When deleted, user resources are moved to a temporary container and can be restored within 30 days.  After that time, they are permanently deleted.  To learn more, see deletedItems.
func (m *UserItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *UserItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// DeviceEnrollmentConfigurations the deviceEnrollmentConfigurations property
func (m *UserItemRequestBuilder) DeviceEnrollmentConfigurations()(*i0fbc6bb6d8e908a7211206db1e3f715af4ae676f093155141c801afcbfd4a468.DeviceEnrollmentConfigurationsRequestBuilder) {
    return i0fbc6bb6d8e908a7211206db1e3f715af4ae676f093155141c801afcbfd4a468.NewDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceEnrollmentConfigurationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.deviceEnrollmentConfigurations.item collection
func (m *UserItemRequestBuilder) DeviceEnrollmentConfigurationsById(id string)(*i20d5bc3e0beaac35f1cb30b0824d071deea9c98aa2da0ea95797426ec3391a8b.DeviceEnrollmentConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceEnrollmentConfiguration%2Did"] = id
    }
    return i20d5bc3e0beaac35f1cb30b0824d071deea9c98aa2da0ea95797426ec3391a8b.NewDeviceEnrollmentConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceManagementTroubleshootingEvents the deviceManagementTroubleshootingEvents property
func (m *UserItemRequestBuilder) DeviceManagementTroubleshootingEvents()(*i9fa6f9d397d0318d3e3fc1137d90e8302a95a72c11d6237f67fc7bd8dc1625b7.DeviceManagementTroubleshootingEventsRequestBuilder) {
    return i9fa6f9d397d0318d3e3fc1137d90e8302a95a72c11d6237f67fc7bd8dc1625b7.NewDeviceManagementTroubleshootingEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceManagementTroubleshootingEventsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.deviceManagementTroubleshootingEvents.item collection
func (m *UserItemRequestBuilder) DeviceManagementTroubleshootingEventsById(id string)(*i4d1bb8d1a6e444b2e99d643949896393d13341c23c9e59170d7cf747778a84f1.DeviceManagementTroubleshootingEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementTroubleshootingEvent%2Did"] = id
    }
    return i4d1bb8d1a6e444b2e99d643949896393d13341c23c9e59170d7cf747778a84f1.NewDeviceManagementTroubleshootingEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Devices the devices property
func (m *UserItemRequestBuilder) Devices()(*ib2e97a67ea231c420ac9fa60c240e9a1314dc844268c406a4c1600c693ad6811.DevicesRequestBuilder) {
    return ib2e97a67ea231c420ac9fa60c240e9a1314dc844268c406a4c1600c693ad6811.NewDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DevicesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.devices.item collection
func (m *UserItemRequestBuilder) DevicesById(id string)(*i8a37b6fd0dca54e1e3129d6fa85360b47ae2206831aad7d8bf6c3823b1fb540b.DeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["device%2Did"] = id
    }
    return i8a37b6fd0dca54e1e3129d6fa85360b47ae2206831aad7d8bf6c3823b1fb540b.NewDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DirectReports the directReports property
func (m *UserItemRequestBuilder) DirectReports()(*i1e8437839d244c8d93e242e55290d8df8b03ca76cf9b009f9b9a19ef5d5922ac.DirectReportsRequestBuilder) {
    return i1e8437839d244c8d93e242e55290d8df8b03ca76cf9b009f9b9a19ef5d5922ac.NewDirectReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectReportsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.directReports.item collection
func (m *UserItemRequestBuilder) DirectReportsById(id string)(*i53afa1534ef724ece2eabf4449fabc9c4904d986ecd4d6a0c11e3784f58eaab4.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i53afa1534ef724ece2eabf4449fabc9c4904d986ecd4d6a0c11e3784f58eaab4.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Drive the drive property
func (m *UserItemRequestBuilder) Drive()(*i3e6a9a935ee43b6bd89f378e203c324f1d517ccc9d1c30202303eafc13d5bb25.DriveRequestBuilder) {
    return i3e6a9a935ee43b6bd89f378e203c324f1d517ccc9d1c30202303eafc13d5bb25.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Drives the drives property
func (m *UserItemRequestBuilder) Drives()(*i00eee5f4e45952d7af6d7dbf02a5397275cbfa69444a9cd54a450f1cc9c5baf6.DrivesRequestBuilder) {
    return i00eee5f4e45952d7af6d7dbf02a5397275cbfa69444a9cd54a450f1cc9c5baf6.NewDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DrivesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item collection
func (m *UserItemRequestBuilder) DrivesById(id string)(*id15759ba2587bc2a0cac47e8d7da5e4bd208929d38688edb78df7122421e531f.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive%2Did"] = id
    }
    return id15759ba2587bc2a0cac47e8d7da5e4bd208929d38688edb78df7122421e531f.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Events the events property
func (m *UserItemRequestBuilder) Events()(*i0bb45fabb973cc1726799c26601b61fb5d4ceef85239469bf8a48d255ddcab77.EventsRequestBuilder) {
    return i0bb45fabb973cc1726799c26601b61fb5d4ceef85239469bf8a48d255ddcab77.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.events.item collection
func (m *UserItemRequestBuilder) EventsById(id string)(*i1fb04953ffbe8174524af5da51d2e85d0602b9611bf6179631d63cd31623aa9c.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return i1fb04953ffbe8174524af5da51d2e85d0602b9611bf6179631d63cd31623aa9c.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ExportDeviceAndAppManagementData provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserItemRequestBuilder) ExportDeviceAndAppManagementData()(*i6c6abf877b483e321abb71bac38e5c8c592f7ee6915c9ad28447c321b2979b4d.ExportDeviceAndAppManagementDataRequestBuilder) {
    return i6c6abf877b483e321abb71bac38e5c8c592f7ee6915c9ad28447c321b2979b4d.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserItemRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*i9e9fd6de702827c484559c38e811567b4bfa805d681aca1a1d7d8d8ba07aae9c.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return i9e9fd6de702827c484559c38e811567b4bfa805d681aca1a1d7d8d8ba07aae9c.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserItemRequestBuilder) ExportPersonalData()(*i10209655736fb29f6fe332f57065da684a4bb7e75afca36c8b0cf9215d9600cf.ExportPersonalDataRequestBuilder) {
    return i10209655736fb29f6fe332f57065da684a4bb7e75afca36c8b0cf9215d9600cf.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *UserItemRequestBuilder) Extensions()(*i49c0a0951cef8133c800eac5f2bf2ddd282c66d3a73fc9f9c652fc532a3331be.ExtensionsRequestBuilder) {
    return i49c0a0951cef8133c800eac5f2bf2ddd282c66d3a73fc9f9c652fc532a3331be.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.extensions.item collection
func (m *UserItemRequestBuilder) ExtensionsById(id string)(*i532345ba6b7192bf733f1b4b3de662c47c1e7e0d3989701bcdc0d8f857dbe580.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i532345ba6b7192bf733f1b4b3de662c47c1e7e0d3989701bcdc0d8f857dbe580.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserItemRequestBuilder) FindMeetingTimes()(*i5c0e1410fbc34cac508ae33c37b1e9dd9a4ca3540ad1525e02f20329f8a264f0.FindMeetingTimesRequestBuilder) {
    return i5c0e1410fbc34cac508ae33c37b1e9dd9a4ca3540ad1525e02f20329f8a264f0.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserItemRequestBuilder) FindRoomLists()(*id5b765e74db4b3dcebe232c27216355fdc0568a1094952e7fc7cbdc858f3d81a.FindRoomListsRequestBuilder) {
    return id5b765e74db4b3dcebe232c27216355fdc0568a1094952e7fc7cbdc858f3d81a.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserItemRequestBuilder) FindRooms()(*i09b3ed8449f0eb014b1a1c391d9057c99dca194b7d9cee6d67e587c2b4d66363.FindRoomsRequestBuilder) {
    return i09b3ed8449f0eb014b1a1c391d9057c99dca194b7d9cee6d67e587c2b4d66363.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserItemRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i933c9dccaf05a6a7fcf066da46cae8d6c4a0940e841426651ace8702209f1b9d.FindRoomsWithRoomListRequestBuilder) {
    return i933c9dccaf05a6a7fcf066da46cae8d6c4a0940e841426651ace8702209f1b9d.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// FollowedSites the followedSites property
func (m *UserItemRequestBuilder) FollowedSites()(*id904363322fca79446fa0f0f2231760c6649433886f1c8e79503338b004b288c.FollowedSitesRequestBuilder) {
    return id904363322fca79446fa0f0f2231760c6649433886f1c8e79503338b004b288c.NewFollowedSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FollowedSitesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.followedSites.item collection
func (m *UserItemRequestBuilder) FollowedSitesById(id string)(*i9bb18abda2da0890104c07ce4b11694592a3ff71b3436063a29752ff07e0f525.SiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["site%2Did"] = id
    }
    return i9bb18abda2da0890104c07ce4b11694592a3ff71b3436063a29752ff07e0f525.NewSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get retrieve the properties and relationships of user object. This operation returns by default only a subset of the more commonly used properties for each user. These _default_ properties are noted in the Properties section. To get properties that are _not_ returned by default, do a GET operation for the user and specify the properties in a `$select` OData query option. Because the **user** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in a **user** instance.
func (m *UserItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserItemRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*i44184067b9a5d3386d10a8d3d78383c99a1f94c6db7a1c78fecc6277422e9842.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return i44184067b9a5d3386d10a8d3d78383c99a1f94c6db7a1c78fecc6277422e9842.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserItemRequestBuilder) GetLoggedOnManagedDevices()(*i163ca73c2d1426f49f134d87416999c0644a5bf7f78d9912f8098a9785dc6ec1.GetLoggedOnManagedDevicesRequestBuilder) {
    return i163ca73c2d1426f49f134d87416999c0644a5bf7f78d9912f8098a9785dc6ec1.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserItemRequestBuilder) GetMailTips()(*i89f9ebf7c167c75aff01142f1e2b30d94fcb3ccf3d4af1752e74dbc3f979ecd6.GetMailTipsRequestBuilder) {
    return i89f9ebf7c167c75aff01142f1e2b30d94fcb3ccf3d4af1752e74dbc3f979ecd6.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserItemRequestBuilder) GetManagedAppDiagnosticStatuses()(*idaeea547224afeb26277747dcbdab987622f0aec250f6cb9dcebbc9e6d05df1e.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return idaeea547224afeb26277747dcbdab987622f0aec250f6cb9dcebbc9e6d05df1e.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserItemRequestBuilder) GetManagedAppPolicies()(*i1b6f500d8aa60be2a2a1e7586a07867f5e80bc6da6a793388efc074011aa9f1b.GetManagedAppPoliciesRequestBuilder) {
    return i1b6f500d8aa60be2a2a1e7586a07867f5e80bc6da6a793388efc074011aa9f1b.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserItemRequestBuilder) GetManagedDevicesWithAppFailures()(*i028859f982626d2101b8d8b68caff0f04ef34796aed805e993e21fefb9152760.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return i028859f982626d2101b8d8b68caff0f04ef34796aed805e993e21fefb9152760.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserItemRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*i52ea864f44c605939329148c5a95626ead1b771dff88c6072de654caee1b281e.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return i52ea864f44c605939329148c5a95626ead1b771dff88c6072de654caee1b281e.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserItemRequestBuilder) GetMemberGroups()(*i633ead85dd41f6f9d94e9cf71e43f604b18195cef363d66f743366cdb820fc8b.GetMemberGroupsRequestBuilder) {
    return i633ead85dd41f6f9d94e9cf71e43f604b18195cef363d66f743366cdb820fc8b.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserItemRequestBuilder) GetMemberObjects()(*i751d3667ec77ce3a4b6544375d44e850f38ac8db37e3237ce66245e949ea015e.GetMemberObjectsRequestBuilder) {
    return i751d3667ec77ce3a4b6544375d44e850f38ac8db37e3237ce66245e949ea015e.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithRequestConfigurationAndResponseHandler retrieve the properties and relationships of user object. This operation returns by default only a subset of the more commonly used properties for each user. These _default_ properties are noted in the Properties section. To get properties that are _not_ returned by default, do a GET operation for the user and specify the properties in a `$select` OData query option. Because the **user** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in a **user** instance.
func (m *UserItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *UserItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable), nil
}
// InferenceClassification the inferenceClassification property
func (m *UserItemRequestBuilder) InferenceClassification()(*i73641486333af15741c085f42cd9ba7b087800533a2f2a8781b636a965272506.InferenceClassificationRequestBuilder) {
    return i73641486333af15741c085f42cd9ba7b087800533a2f2a8781b636a965272506.NewInferenceClassificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InformationProtection the informationProtection property
func (m *UserItemRequestBuilder) InformationProtection()(*i992f3c4694021035f11bae1cb737eb6020f9d1221bc7676d2685a415cab74221.InformationProtectionRequestBuilder) {
    return i992f3c4694021035f11bae1cb737eb6020f9d1221bc7676d2685a415cab74221.NewInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Insights the insights property
func (m *UserItemRequestBuilder) Insights()(*i5f232e4d902d95d06ec20d5128cce05b56d53758222f6b28e307a6b909082dac.InsightsRequestBuilder) {
    return i5f232e4d902d95d06ec20d5128cce05b56d53758222f6b28e307a6b909082dac.NewInsightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InvalidateAllRefreshTokens the invalidateAllRefreshTokens property
func (m *UserItemRequestBuilder) InvalidateAllRefreshTokens()(*i7f0c66e611b8af22dffef0886d0dc8468dc8cac8bb0b577016f3d4b1daf70528.InvalidateAllRefreshTokensRequestBuilder) {
    return i7f0c66e611b8af22dffef0886d0dc8468dc8cac8bb0b577016f3d4b1daf70528.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserItemRequestBuilder) IsManagedAppUserBlocked()(*ieee72b0b162c07d3dde8ae1950af04718a5801d0d3e6575e0cba2faa3f56157a.IsManagedAppUserBlockedRequestBuilder) {
    return ieee72b0b162c07d3dde8ae1950af04718a5801d0d3e6575e0cba2faa3f56157a.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JoinedGroups the joinedGroups property
func (m *UserItemRequestBuilder) JoinedGroups()(*i235dd448b749bb3a1187ec0970cdeb5c226f7aee1c9e6c72733a9d1a15182ffe.JoinedGroupsRequestBuilder) {
    return i235dd448b749bb3a1187ec0970cdeb5c226f7aee1c9e6c72733a9d1a15182ffe.NewJoinedGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JoinedTeams the joinedTeams property
func (m *UserItemRequestBuilder) JoinedTeams()(*i4e8ad44991b185083366e107199ff13f6438a07cafd58ae36e0e862692bf9826.JoinedTeamsRequestBuilder) {
    return i4e8ad44991b185083366e107199ff13f6438a07cafd58ae36e0e862692bf9826.NewJoinedTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JoinedTeamsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedTeams.item collection
func (m *UserItemRequestBuilder) JoinedTeamsById(id string)(*i9ff5415484e41f91043b46ae2894ad1a1ddead74f139bbcb66a0ea970caba35b.TeamItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["team%2Did"] = id
    }
    return i9ff5415484e41f91043b46ae2894ad1a1ddead74f139bbcb66a0ea970caba35b.NewTeamItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// LicenseDetails the licenseDetails property
func (m *UserItemRequestBuilder) LicenseDetails()(*i53e4d773c077f49b4590da13b904e6a502ef1f8f43b7a7fdfaba256a667df083.LicenseDetailsRequestBuilder) {
    return i53e4d773c077f49b4590da13b904e6a502ef1f8f43b7a7fdfaba256a667df083.NewLicenseDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LicenseDetailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.licenseDetails.item collection
func (m *UserItemRequestBuilder) LicenseDetailsById(id string)(*i41bcf6e933152d51a0364a3166ccb57d9e6cc7316937216a235e9f7a7973b36c.LicenseDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["licenseDetails%2Did"] = id
    }
    return i41bcf6e933152d51a0364a3166ccb57d9e6cc7316937216a235e9f7a7973b36c.NewLicenseDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MailFolders the mailFolders property
func (m *UserItemRequestBuilder) MailFolders()(*i21282294d2d1bb2a0ab74695618e809021209e6abac85e175b427b95d4891cc5.MailFoldersRequestBuilder) {
    return i21282294d2d1bb2a0ab74695618e809021209e6abac85e175b427b95d4891cc5.NewMailFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MailFoldersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mailFolders.item collection
func (m *UserItemRequestBuilder) MailFoldersById(id string)(*i11a0d65a9c2024ff5c04a9b3cb79486d6f2277859e51094a083537f2289c0a27.MailFolderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mailFolder%2Did"] = id
    }
    return i11a0d65a9c2024ff5c04a9b3cb79486d6f2277859e51094a083537f2289c0a27.NewMailFolderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedAppRegistrations the managedAppRegistrations property
func (m *UserItemRequestBuilder) ManagedAppRegistrations()(*i44dc4d21a51a483da0eb7ac0bc6c38fa258f1e982bb140f0bd9282a87d39015c.ManagedAppRegistrationsRequestBuilder) {
    return i44dc4d21a51a483da0eb7ac0bc6c38fa258f1e982bb140f0bd9282a87d39015c.NewManagedAppRegistrationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedAppRegistrationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.managedAppRegistrations.item collection
func (m *UserItemRequestBuilder) ManagedAppRegistrationsById(id string)(*id62ce646a00a4421b260562c6eb4464e0a6e23540ac62847efa7af52e9578392.ManagedAppRegistrationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppRegistration%2Did"] = id
    }
    return id62ce646a00a4421b260562c6eb4464e0a6e23540ac62847efa7af52e9578392.NewManagedAppRegistrationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedDevices the managedDevices property
func (m *UserItemRequestBuilder) ManagedDevices()(*i4cd644380b450eedb89634c647645afa6ede7ee4df886587de78ccf69b3a9a42.ManagedDevicesRequestBuilder) {
    return i4cd644380b450eedb89634c647645afa6ede7ee4df886587de78ccf69b3a9a42.NewManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDevicesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.managedDevices.item collection
func (m *UserItemRequestBuilder) ManagedDevicesById(id string)(*i0796e8f7b9f803857a438202c35d9fccd548d4779e6f2a341287440bf41acc70.ManagedDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDevice%2Did"] = id
    }
    return i0796e8f7b9f803857a438202c35d9fccd548d4779e6f2a341287440bf41acc70.NewManagedDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Manager the manager property
func (m *UserItemRequestBuilder) Manager()(*ice23d35988441664e0a80c4309242575aeb7cc7c84c7bfc6e9ff86ce92c7cde5.ManagerRequestBuilder) {
    return ice23d35988441664e0a80c4309242575aeb7cc7c84c7bfc6e9ff86ce92c7cde5.NewManagerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOf the memberOf property
func (m *UserItemRequestBuilder) MemberOf()(*i35ae79fec43abc52b415d013c5ba088edde69bf589391525fec20b1f062c22a7.MemberOfRequestBuilder) {
    return i35ae79fec43abc52b415d013c5ba088edde69bf589391525fec20b1f062c22a7.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOfById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.memberOf.item collection
func (m *UserItemRequestBuilder) MemberOfById(id string)(*idbb158da6d78d676ea0c3d14f6ae7004f016a1bdc89a3da430f82196981f35af.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return idbb158da6d78d676ea0c3d14f6ae7004f016a1bdc89a3da430f82196981f35af.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages the messages property
func (m *UserItemRequestBuilder) Messages()(*ifc52bce07b756c55dabe5cd8f6a0d8c749809ddcc3831f5bded6c944d449dbd8.MessagesRequestBuilder) {
    return ifc52bce07b756c55dabe5cd8f6a0d8c749809ddcc3831f5bded6c944d449dbd8.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.messages.item collection
func (m *UserItemRequestBuilder) MessagesById(id string)(*i92f204ee1ec183e477ad6aae362cad151b598b7439eda0c58c4750c295bc84e8.MessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["message%2Did"] = id
    }
    return i92f204ee1ec183e477ad6aae362cad151b598b7439eda0c58c4750c295bc84e8.NewMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MobileAppIntentAndStates the mobileAppIntentAndStates property
func (m *UserItemRequestBuilder) MobileAppIntentAndStates()(*ide4da4ea5e9d5b4a3ed2b815f5032fc4321b390d6e1c0ba686af0a35ec886a84.MobileAppIntentAndStatesRequestBuilder) {
    return ide4da4ea5e9d5b4a3ed2b815f5032fc4321b390d6e1c0ba686af0a35ec886a84.NewMobileAppIntentAndStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileAppIntentAndStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mobileAppIntentAndStates.item collection
func (m *UserItemRequestBuilder) MobileAppIntentAndStatesById(id string)(*i9e371fd6357698c3cd300fbb8012addd3ec298d5df1be3dfa58412fb63c3cae8.MobileAppIntentAndStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppIntentAndState%2Did"] = id
    }
    return i9e371fd6357698c3cd300fbb8012addd3ec298d5df1be3dfa58412fb63c3cae8.NewMobileAppIntentAndStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MobileAppTroubleshootingEvents the mobileAppTroubleshootingEvents property
func (m *UserItemRequestBuilder) MobileAppTroubleshootingEvents()(*i7358087738ea2e8de75562a191c6ccb72c9a24d79c1f9a33de77f140cbb5ce74.MobileAppTroubleshootingEventsRequestBuilder) {
    return i7358087738ea2e8de75562a191c6ccb72c9a24d79c1f9a33de77f140cbb5ce74.NewMobileAppTroubleshootingEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileAppTroubleshootingEventsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mobileAppTroubleshootingEvents.item collection
func (m *UserItemRequestBuilder) MobileAppTroubleshootingEventsById(id string)(*i7fdae8345bbd328abce3eca8f17ca05dfb12b3a6a3adcb8b7b394f3697581dae.MobileAppTroubleshootingEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppTroubleshootingEvent%2Did"] = id
    }
    return i7fdae8345bbd328abce3eca8f17ca05dfb12b3a6a3adcb8b7b394f3697581dae.NewMobileAppTroubleshootingEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Notifications the notifications property
func (m *UserItemRequestBuilder) Notifications()(*i5c911b116a9e04c78f811a8141c00da4569a7dc37a1d8590843bc9d14e30377e.NotificationsRequestBuilder) {
    return i5c911b116a9e04c78f811a8141c00da4569a7dc37a1d8590843bc9d14e30377e.NewNotificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotificationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.notifications.item collection
func (m *UserItemRequestBuilder) NotificationsById(id string)(*i5d09ee5a147e62e93172ba3c3003b1e37657e5078803e7fa2be2756f303a9607.NotificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["notification%2Did"] = id
    }
    return i5d09ee5a147e62e93172ba3c3003b1e37657e5078803e7fa2be2756f303a9607.NewNotificationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Oauth2PermissionGrants the oauth2PermissionGrants property
func (m *UserItemRequestBuilder) Oauth2PermissionGrants()(*ib674e5ebc7e9a35ea83a7a23737e2e850d21d009460f455398f4235cd71eb2d8.Oauth2PermissionGrantsRequestBuilder) {
    return ib674e5ebc7e9a35ea83a7a23737e2e850d21d009460f455398f4235cd71eb2d8.NewOauth2PermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Oauth2PermissionGrantsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.oauth2PermissionGrants.item collection
func (m *UserItemRequestBuilder) Oauth2PermissionGrantsById(id string)(*ideab25898675089d3e174a0e5f0bde76aa4e227661b4518a256695e6a0500360.OAuth2PermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["oAuth2PermissionGrant%2Did"] = id
    }
    return ideab25898675089d3e174a0e5f0bde76aa4e227661b4518a256695e6a0500360.NewOAuth2PermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Onenote the onenote property
func (m *UserItemRequestBuilder) Onenote()(*i4c553c01d10dc5ea91f5e8443a25576af37822a1a9e3b20c39f7d69613ca1a5d.OnenoteRequestBuilder) {
    return i4c553c01d10dc5ea91f5e8443a25576af37822a1a9e3b20c39f7d69613ca1a5d.NewOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnlineMeetings the onlineMeetings property
func (m *UserItemRequestBuilder) OnlineMeetings()(*idddb11dcacd5585b4b6cb080f05aa0ce5c62d39e09bc0ffcd012d5c8c731810b.OnlineMeetingsRequestBuilder) {
    return idddb11dcacd5585b4b6cb080f05aa0ce5c62d39e09bc0ffcd012d5c8c731810b.NewOnlineMeetingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnlineMeetingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.onlineMeetings.item collection
func (m *UserItemRequestBuilder) OnlineMeetingsById(id string)(*i8e9bfc4860807cf4727d9197394851201c63af67494fa725e469b158152f82c0.OnlineMeetingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onlineMeeting%2Did"] = id
    }
    return i8e9bfc4860807cf4727d9197394851201c63af67494fa725e469b158152f82c0.NewOnlineMeetingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Outlook the outlook property
func (m *UserItemRequestBuilder) Outlook()(*ib9ea7c893835d89864cc7cd939d24f579ba191ab540031694fe5202498a4240c.OutlookRequestBuilder) {
    return ib9ea7c893835d89864cc7cd939d24f579ba191ab540031694fe5202498a4240c.NewOutlookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnedDevices the ownedDevices property
func (m *UserItemRequestBuilder) OwnedDevices()(*i71ece4d9b62444d7cd9edeac6c2b5bdaf0f35379d154d089d602c3d246239739.OwnedDevicesRequestBuilder) {
    return i71ece4d9b62444d7cd9edeac6c2b5bdaf0f35379d154d089d602c3d246239739.NewOwnedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnedDevicesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.ownedDevices.item collection
func (m *UserItemRequestBuilder) OwnedDevicesById(id string)(*i89586fbadaf2720611403fc97537d705bca9891589c65eb565dbae3f293f06c1.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i89586fbadaf2720611403fc97537d705bca9891589c65eb565dbae3f293f06c1.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OwnedObjects the ownedObjects property
func (m *UserItemRequestBuilder) OwnedObjects()(*i3aa2a5875223f6b19560c0365147205688fff76fe30e6bbf724b3b25f7b00a43.OwnedObjectsRequestBuilder) {
    return i3aa2a5875223f6b19560c0365147205688fff76fe30e6bbf724b3b25f7b00a43.NewOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnedObjectsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.ownedObjects.item collection
func (m *UserItemRequestBuilder) OwnedObjectsById(id string)(*i84c9cc6d33fc548d6b6b6210a170d3ce1e97e608f9af2ff36d66a133895c1e8b.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i84c9cc6d33fc548d6b6b6210a170d3ce1e97e608f9af2ff36d66a133895c1e8b.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the properties of a user object. Not all properties can be updated by Member or Guest users with their default permissions without Administrator roles. Compare member and guest default permissions to see properties they can manage.
func (m *UserItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the properties of a user object. Not all properties can be updated by Member or Guest users with their default permissions without Administrator roles. Compare member and guest default permissions to see properties they can manage.
func (m *UserItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, requestConfiguration *UserItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// PendingAccessReviewInstances the pendingAccessReviewInstances property
func (m *UserItemRequestBuilder) PendingAccessReviewInstances()(*i3be6b130d8bf10a5df104376e126ee2c3584be236591382ff88662fd6a574158.PendingAccessReviewInstancesRequestBuilder) {
    return i3be6b130d8bf10a5df104376e126ee2c3584be236591382ff88662fd6a574158.NewPendingAccessReviewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PendingAccessReviewInstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.pendingAccessReviewInstances.item collection
func (m *UserItemRequestBuilder) PendingAccessReviewInstancesById(id string)(*i2cf542b96fcffbafb31190c7c7e24f2b652b0249d3c7d0e6deaf71e333abf373.AccessReviewInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewInstance%2Did"] = id
    }
    return i2cf542b96fcffbafb31190c7c7e24f2b652b0249d3c7d0e6deaf71e333abf373.NewAccessReviewInstanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// People the people property
func (m *UserItemRequestBuilder) People()(*i1faa710c227b8e4b8d7abef5e9611e6604d3f4265db3da621d2ddd654a3a33f4.PeopleRequestBuilder) {
    return i1faa710c227b8e4b8d7abef5e9611e6604d3f4265db3da621d2ddd654a3a33f4.NewPeopleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PeopleById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.people.item collection
func (m *UserItemRequestBuilder) PeopleById(id string)(*ia16782564e6f61796daddf42e21afd9855b9a1c1c41eef4e65e545763ac7c8d5.PersonItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["person%2Did"] = id
    }
    return ia16782564e6f61796daddf42e21afd9855b9a1c1c41eef4e65e545763ac7c8d5.NewPersonItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Photo the photo property
func (m *UserItemRequestBuilder) Photo()(*i5f9e5aa5c251065cc59079614790107bc87c70c65435cea5312a9ad92e27cad3.PhotoRequestBuilder) {
    return i5f9e5aa5c251065cc59079614790107bc87c70c65435cea5312a9ad92e27cad3.NewPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Photos the photos property
func (m *UserItemRequestBuilder) Photos()(*i58ea06623816d885f9dc5730c5fd341b9dc32ca959bf1ffe118d49d00bdfbc62.PhotosRequestBuilder) {
    return i58ea06623816d885f9dc5730c5fd341b9dc32ca959bf1ffe118d49d00bdfbc62.NewPhotosRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PhotosById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.photos.item collection
func (m *UserItemRequestBuilder) PhotosById(id string)(*ib39beceac49e7dd17ed1c0885887dae465a45afa34b3eebb70aa5fc08ca691e7.ProfilePhotoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["profilePhoto%2Did"] = id
    }
    return ib39beceac49e7dd17ed1c0885887dae465a45afa34b3eebb70aa5fc08ca691e7.NewProfilePhotoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Planner the planner property
func (m *UserItemRequestBuilder) Planner()(*i64d17e545a98d6600890650331d10cc8ac9bafd9f1a9888ceee3f7c9932f9d68.PlannerRequestBuilder) {
    return i64d17e545a98d6600890650331d10cc8ac9bafd9f1a9888ceee3f7c9932f9d68.NewPlannerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Presence the presence property
func (m *UserItemRequestBuilder) Presence()(*i2a5b8dfcb3b418ee303ee5f5b01aebd3209cdfa9122c52e81bd0a9e16a726f9d.PresenceRequestBuilder) {
    return i2a5b8dfcb3b418ee303ee5f5b01aebd3209cdfa9122c52e81bd0a9e16a726f9d.NewPresenceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Profile the profile property
func (m *UserItemRequestBuilder) Profile()(*i77129a8685763d465ea62e5de870203f0d38a41c496b1f718e0c695d6df2b42c.ProfileRequestBuilder) {
    return i77129a8685763d465ea62e5de870203f0d38a41c496b1f718e0c695d6df2b42c.NewProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredDevices the registeredDevices property
func (m *UserItemRequestBuilder) RegisteredDevices()(*i9f79453a6cb43e9e5e68087f1d6e774414d8ba83449733d805265bf8cb3d0b5d.RegisteredDevicesRequestBuilder) {
    return i9f79453a6cb43e9e5e68087f1d6e774414d8ba83449733d805265bf8cb3d0b5d.NewRegisteredDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredDevicesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.registeredDevices.item collection
func (m *UserItemRequestBuilder) RegisteredDevicesById(id string)(*i404e4499afbbdf92e8ee2c947bd2f30954ca765dc8d5deeda689c119c7e7eff7.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i404e4499afbbdf92e8ee2c947bd2f30954ca765dc8d5deeda689c119c7e7eff7.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserItemRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i1e2ed1761df19ea44d2d90f6fdfba02f95eb595495e7804f87126e72ee9345e3.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i1e2ed1761df19ea44d2d90f6fdfba02f95eb595495e7804f87126e72ee9345e3.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserItemRequestBuilder) RemoveAllDevicesFromManagement()(*i7d78da96ae198e74cdcd846ebd00d688295a8b346c939e4e49d56b88b5f72711.RemoveAllDevicesFromManagementRequestBuilder) {
    return i7d78da96ae198e74cdcd846ebd00d688295a8b346c939e4e49d56b88b5f72711.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserItemRequestBuilder) ReprocessLicenseAssignment()(*i35b254e720ef3a77057a7b6736d2d9841a2d361b2b095e7a189fd80751427060.ReprocessLicenseAssignmentRequestBuilder) {
    return i35b254e720ef3a77057a7b6736d2d9841a2d361b2b095e7a189fd80751427060.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserItemRequestBuilder) Restore()(*i6d8c6f6a8d72bfe9260edd2b3126823191856ab8f78dd9cd5a1c01fb0fff9c47.RestoreRequestBuilder) {
    return i6d8c6f6a8d72bfe9260edd2b3126823191856ab8f78dd9cd5a1c01fb0fff9c47.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserItemRequestBuilder) RevokeSignInSessions()(*i6545b8651bf39e658fd8159fa1123831d6f44a855626360fac6873d3a90581c8.RevokeSignInSessionsRequestBuilder) {
    return i6545b8651bf39e658fd8159fa1123831d6f44a855626360fac6873d3a90581c8.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScopedRoleMemberOf the scopedRoleMemberOf property
func (m *UserItemRequestBuilder) ScopedRoleMemberOf()(*if3e915183aa6002a59a5c947830462a91684160d445bda34aa2c4cf6886b82b2.ScopedRoleMemberOfRequestBuilder) {
    return if3e915183aa6002a59a5c947830462a91684160d445bda34aa2c4cf6886b82b2.NewScopedRoleMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScopedRoleMemberOfById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.scopedRoleMemberOf.item collection
func (m *UserItemRequestBuilder) ScopedRoleMemberOfById(id string)(*ie132362db0baae23ad9567ff4032a62c165d8386308f8baec97ecbb60c10a021.ScopedRoleMembershipItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["scopedRoleMembership%2Did"] = id
    }
    return ie132362db0baae23ad9567ff4032a62c165d8386308f8baec97ecbb60c10a021.NewScopedRoleMembershipItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Security the security property
func (m *UserItemRequestBuilder) Security()(*ifc1c29cd3428806f56c209014f087cf657f124d2250271baf855225bc6fb7c3d.SecurityRequestBuilder) {
    return ifc1c29cd3428806f56c209014f087cf657f124d2250271baf855225bc6fb7c3d.NewSecurityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserItemRequestBuilder) SendMail()(*id6953e096a9515e55975655319314bf3955019e0bded78bb5cb6fb32eb4e8e94.SendMailRequestBuilder) {
    return id6953e096a9515e55975655319314bf3955019e0bded78bb5cb6fb32eb4e8e94.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Settings the settings property
func (m *UserItemRequestBuilder) Settings()(*i7984c072c770ec17745d737738e8ab409b4f442b183221a451997db094cd8b13.SettingsRequestBuilder) {
    return i7984c072c770ec17745d737738e8ab409b4f442b183221a451997db094cd8b13.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Tasks the tasks property
func (m *UserItemRequestBuilder) Tasks()(*i807859e90e02be96c7951593aa65cbf418e7a30d0f62e3a47b8382296d35e2d2.TasksRequestBuilder) {
    return i807859e90e02be96c7951593aa65cbf418e7a30d0f62e3a47b8382296d35e2d2.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Teamwork the teamwork property
func (m *UserItemRequestBuilder) Teamwork()(*if429d108f0ce6b9f644fc8ece21d69b6a892d9788590f39dc3d9d21bda91a159.TeamworkRequestBuilder) {
    return if429d108f0ce6b9f644fc8ece21d69b6a892d9788590f39dc3d9d21bda91a159.NewTeamworkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Todo the todo property
func (m *UserItemRequestBuilder) Todo()(*i4e66b80d61d13c58a5325be0d73998d6c528f982ee63fc876b66fdffd4ec5473.TodoRequestBuilder) {
    return i4e66b80d61d13c58a5325be0d73998d6c528f982ee63fc876b66fdffd4ec5473.NewTodoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOf the transitiveMemberOf property
func (m *UserItemRequestBuilder) TransitiveMemberOf()(*i6f957db73ce02e759548f419eb79a2bfdfb04d6c5206523cd438181cc5619df8.TransitiveMemberOfRequestBuilder) {
    return i6f957db73ce02e759548f419eb79a2bfdfb04d6c5206523cd438181cc5619df8.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOfById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.transitiveMemberOf.item collection
func (m *UserItemRequestBuilder) TransitiveMemberOfById(id string)(*iabe64332e133a9af5a6f5aafde3a47ea985f385ae58a2f524f86fc1a3ef2f539.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return iabe64332e133a9af5a6f5aafde3a47ea985f385ae58a2f524f86fc1a3ef2f539.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TransitiveReports the transitiveReports property
func (m *UserItemRequestBuilder) TransitiveReports()(*ia1e1519710e771d8df5e22da39eeb50d37180582bb2065916e5cd0703a4d1cce.TransitiveReportsRequestBuilder) {
    return ia1e1519710e771d8df5e22da39eeb50d37180582bb2065916e5cd0703a4d1cce.NewTransitiveReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveReportsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.transitiveReports.item collection
func (m *UserItemRequestBuilder) TransitiveReportsById(id string)(*i2c5057cc6c2035899f3804d4bce915570738d9af43d8e9a3f1ea425e7c0130ff.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i2c5057cc6c2035899f3804d4bce915570738d9af43d8e9a3f1ea425e7c0130ff.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserItemRequestBuilder) TranslateExchangeIds()(*i87f894b87c2947e4dad88e845ac91e8a8ddf1ca5ba6b426384ab8873fe897913.TranslateExchangeIdsRequestBuilder) {
    return i87f894b87c2947e4dad88e845ac91e8a8ddf1ca5ba6b426384ab8873fe897913.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserItemRequestBuilder) UnblockManagedApps()(*i1d3f5d5cc5119c3710246df1436375c7f067bf4348a4cf5c5decdf0052f3b667.UnblockManagedAppsRequestBuilder) {
    return i1d3f5d5cc5119c3710246df1436375c7f067bf4348a4cf5c5decdf0052f3b667.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsageRights the usageRights property
func (m *UserItemRequestBuilder) UsageRights()(*iddc68c193cc623cc513e0675dadbd063a05b9d5ce8584eb5a7e8ebf24faa98e0.UsageRightsRequestBuilder) {
    return iddc68c193cc623cc513e0675dadbd063a05b9d5ce8584eb5a7e8ebf24faa98e0.NewUsageRightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsageRightsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.usageRights.item collection
func (m *UserItemRequestBuilder) UsageRightsById(id string)(*ib9adfa359ccdf30972cf1f2c9fb0c0a8edde337e58fe26f0a95800e7ce3ba9db.UsageRightItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["usageRight%2Did"] = id
    }
    return ib9adfa359ccdf30972cf1f2c9fb0c0a8edde337e58fe26f0a95800e7ce3ba9db.NewUsageRightItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsInformationProtectionDeviceRegistrations the windowsInformationProtectionDeviceRegistrations property
func (m *UserItemRequestBuilder) WindowsInformationProtectionDeviceRegistrations()(*i3eaf8bc51db6f4bd95448da3951002ebce4db39125b63eaa1cb9987f99da858f.WindowsInformationProtectionDeviceRegistrationsRequestBuilder) {
    return i3eaf8bc51db6f4bd95448da3951002ebce4db39125b63eaa1cb9987f99da858f.NewWindowsInformationProtectionDeviceRegistrationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsInformationProtectionDeviceRegistrationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.windowsInformationProtectionDeviceRegistrations.item collection
func (m *UserItemRequestBuilder) WindowsInformationProtectionDeviceRegistrationsById(id string)(*iba97b8ef76c5efeab4d738316d6e0d7c1266515ff3632bf4789418105475f2bd.WindowsInformationProtectionDeviceRegistrationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsInformationProtectionDeviceRegistration%2Did"] = id
    }
    return iba97b8ef76c5efeab4d738316d6e0d7c1266515ff3632bf4789418105475f2bd.NewWindowsInformationProtectionDeviceRegistrationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserItemRequestBuilder) WipeAndBlockManagedApps()(*ieac84d2e3111366253e033fed8ced42323b1927cd610b6c93eb3eea6d066fd20.WipeAndBlockManagedAppsRequestBuilder) {
    return ieac84d2e3111366253e033fed8ced42323b1927cd610b6c93eb3eea6d066fd20.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserItemRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*i099d32137e6ab567a080a6d574abebbf5b3e8b218fa6fda8a2a0fce52a347329.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return i099d32137e6ab567a080a6d574abebbf5b3e8b218fa6fda8a2a0fce52a347329.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserItemRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*ice64ec2e30acd089e9604f4d076891112981966d167a6e899c5e617a13771654.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return ice64ec2e30acd089e9604f4d076891112981966d167a6e899c5e617a13771654.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserItemRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i23fd320d20dc894e875e4225a9c0031362d1c191fcba8ad855a39c641e46a084.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i23fd320d20dc894e875e4225a9c0031362d1c191fcba8ad855a39c641e46a084.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
