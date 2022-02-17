package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i007cac44b1ad1a90ea316a1c0c51e045017f34e30cc3b64342ae3537775aaca4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/activateserviceplan"
    i00eee5f4e45952d7af6d7dbf02a5397275cbfa69444a9cd54a450f1cc9c5baf6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives"
    i028859f982626d2101b8d8b68caff0f04ef34796aed805e993e21fefb9152760 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/getmanageddeviceswithappfailures"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
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
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i53e4d773c077f49b4590da13b904e6a502ef1f8f43b7a7fdfaba256a667df083 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/licensedetails"
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
    iaf2011a15870df19ddf40701bc4ca449214192204d76149e38fd507c87344ec5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/activities"
    ib2e97a67ea231c420ac9fa60c240e9a1314dc844268c406a4c1600c693ad6811 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices"
    ib674e5ebc7e9a35ea83a7a23737e2e850d21d009460f455398f4235cd71eb2d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/oauth2permissiongrants"
    ib9ea7c893835d89864cc7cd939d24f579ba191ab540031694fe5202498a4240c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/outlook"
    ice23d35988441664e0a80c4309242575aeb7cc7c84c7bfc6e9ff86ce92c7cde5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manager"
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
    ifc52bce07b756c55dabe5cd8f6a0d8c749809ddcc3831f5bded6c944d449dbd8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages"
    i0796e8f7b9f803857a438202c35d9fccd548d4779e6f2a341287440bf41acc70 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item"
    i11a0d65a9c2024ff5c04a9b3cb79486d6f2277859e51094a083537f2289c0a27 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item"
    i1a9eb6fb24429a79e1e7397bd609e240ae9eef7f551f559e57f6c0050aa81ec7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item"
    i1fb04953ffbe8174524af5da51d2e85d0602b9611bf6179631d63cd31623aa9c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item"
    i20d5bc3e0beaac35f1cb30b0824d071deea9c98aa2da0ea95797426ec3391a8b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/deviceenrollmentconfigurations/item"
    i2cf542b96fcffbafb31190c7c7e24f2b652b0249d3c7d0e6deaf71e333abf373 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item"
    i322200ffec4334c71536e3a74319b0afb2071fae07124a70501b4a91e5d3efc2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item"
    i4d1bb8d1a6e444b2e99d643949896393d13341c23c9e59170d7cf747778a84f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devicemanagementtroubleshootingevents/item"
    i532345ba6b7192bf733f1b4b3de662c47c1e7e0d3989701bcdc0d8f857dbe580 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/extensions/item"
    i5d09ee5a147e62e93172ba3c3003b1e37657e5078803e7fa2be2756f303a9607 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/notifications/item"
    i66ea18aba758406b3b710707db096aa499334ab8141ff00b5817c0160d62f849 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contacts/item"
    i7f7b81e49212c0fd4aff0640f7a76daf8d10816c9da8bbbf24c63c6d3002dac7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/approleassignments/item"
    i7fdae8345bbd328abce3eca8f17ca05dfb12b3a6a3adcb8b7b394f3697581dae "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mobileapptroubleshootingevents/item"
    i86e9fa8da4e357b5840f59b78cd49a5abce31f7a90096cc8b85047571e5a2b06 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item"
    i8a37b6fd0dca54e1e3129d6fa85360b47ae2206831aad7d8bf6c3823b1fb540b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item"
    i8bc97b715d2b1f741be45e9db03b6f18317426bb77532976ccdc40d366389a11 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/chats/item"
    i8c655ab8b24fc37660a4329441f6c9936205ed7192659a7b410b5c1b127f0ebd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/activities/item"
    i8e9bfc4860807cf4727d9197394851201c63af67494fa725e469b158152f82c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onlinemeetings/item"
    i92f204ee1ec183e477ad6aae362cad151b598b7439eda0c58c4750c295bc84e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages/item"
    i9e371fd6357698c3cd300fbb8012addd3ec298d5df1be3dfa58412fb63c3cae8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mobileappintentandstates/item"
    ia16782564e6f61796daddf42e21afd9855b9a1c1c41eef4e65e545763ac7c8d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/people/item"
    ib1c924a2a83336c2eee6c85a7c0edc6e41999bab4aa8e2db2755b64356506e43 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/contactfolders/item"
    ib39beceac49e7dd17ed1c0885887dae465a45afa34b3eebb70aa5fc08ca691e7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/photos/item"
    ib9adfa359ccdf30972cf1f2c9fb0c0a8edde337e58fe26f0a95800e7ce3ba9db "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/usagerights/item"
    ic60383690980d3eba5846d285791ffa2fb94867c1850c6c8da9aaa03aca3c008 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/appconsentrequestsforapproval/item"
    id15759ba2587bc2a0cac47e8d7da5e4bd208929d38688edb78df7122421e531f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item"
    ie132362db0baae23ad9567ff4032a62c165d8386308f8baec97ecbb60c10a021 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/scopedrolememberof/item"
    ie39c35482440e2a813ce2612b19ea08745ca523bb7a24088144c3428169a006d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item"
    ifeea6837d24ad8765a5cb2092fe2e730b4403b36aa4e813a9cdfdb831a8f7974 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/approvals/item"
)

// UserRequestBuilder builds and executes requests for operations under \users\{user-id}
type UserRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UserRequestBuilderDeleteOptions options for Delete
type UserRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserRequestBuilderGetOptions options for Get
type UserRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UserRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserRequestBuilderGetQueryParameters get entity from users by key
type UserRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UserRequestBuilderPatchOptions options for Patch
type UserRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.User;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *UserRequestBuilder) ActivateServicePlan()(*i007cac44b1ad1a90ea316a1c0c51e045017f34e30cc3b64342ae3537775aaca4.ActivateServicePlanRequestBuilder) {
    return i007cac44b1ad1a90ea316a1c0c51e045017f34e30cc3b64342ae3537775aaca4.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Activities()(*iaf2011a15870df19ddf40701bc4ca449214192204d76149e38fd507c87344ec5.ActivitiesRequestBuilder) {
    return iaf2011a15870df19ddf40701bc4ca449214192204d76149e38fd507c87344ec5.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.activities.item collection
func (m *UserRequestBuilder) ActivitiesById(id string)(*i8c655ab8b24fc37660a4329441f6c9936205ed7192659a7b410b5c1b127f0ebd.UserActivityRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userActivity_id"] = id
    }
    return i8c655ab8b24fc37660a4329441f6c9936205ed7192659a7b410b5c1b127f0ebd.NewUserActivityRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) AgreementAcceptances()(*i3d9c184215c67e8ba108a3159afbd942d4eaa93fc6fdec9593468297b72346e5.AgreementAcceptancesRequestBuilder) {
    return i3d9c184215c67e8ba108a3159afbd942d4eaa93fc6fdec9593468297b72346e5.NewAgreementAcceptancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Analytics()(*if6f18d24cca69003f155fe07812d8a9ade476b4b7cb0560b8869ba314a7a79aa.AnalyticsRequestBuilder) {
    return if6f18d24cca69003f155fe07812d8a9ade476b4b7cb0560b8869ba314a7a79aa.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) AppConsentRequestsForApproval()(*if0acf31f23813f467eba440f3275db3949170732f975822128618d2ec5b19f15.AppConsentRequestsForApprovalRequestBuilder) {
    return if0acf31f23813f467eba440f3275db3949170732f975822128618d2ec5b19f15.NewAppConsentRequestsForApprovalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppConsentRequestsForApprovalById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.appConsentRequestsForApproval.item collection
func (m *UserRequestBuilder) AppConsentRequestsForApprovalById(id string)(*ic60383690980d3eba5846d285791ffa2fb94867c1850c6c8da9aaa03aca3c008.AppConsentRequestRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appConsentRequest_id"] = id
    }
    return ic60383690980d3eba5846d285791ffa2fb94867c1850c6c8da9aaa03aca3c008.NewAppConsentRequestRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) AppRoleAssignments()(*i67e4cd2b85517c9aa8f5256b8ca64c3daab00417eaeedb8528d78204b2257eb1.AppRoleAssignmentsRequestBuilder) {
    return i67e4cd2b85517c9aa8f5256b8ca64c3daab00417eaeedb8528d78204b2257eb1.NewAppRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.appRoleAssignments.item collection
func (m *UserRequestBuilder) AppRoleAssignmentsById(id string)(*i7f7b81e49212c0fd4aff0640f7a76daf8d10816c9da8bbbf24c63c6d3002dac7.AppRoleAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appRoleAssignment_id"] = id
    }
    return i7f7b81e49212c0fd4aff0640f7a76daf8d10816c9da8bbbf24c63c6d3002dac7.NewAppRoleAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) Approvals()(*i2d035ce91a466eb27cdd0da2fa61e30e90dd33e1f08803145645a2bbbd70f1df.ApprovalsRequestBuilder) {
    return i2d035ce91a466eb27cdd0da2fa61e30e90dd33e1f08803145645a2bbbd70f1df.NewApprovalsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApprovalsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.approvals.item collection
func (m *UserRequestBuilder) ApprovalsById(id string)(*ifeea6837d24ad8765a5cb2092fe2e730b4403b36aa4e813a9cdfdb831a8f7974.ApprovalRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["approval_id"] = id
    }
    return ifeea6837d24ad8765a5cb2092fe2e730b4403b36aa4e813a9cdfdb831a8f7974.NewApprovalRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) AssignLicense()(*i92c37b67685c58d9f60f7203b77a560478fb28fb9753bb4bee0d7e274eeb3e16.AssignLicenseRequestBuilder) {
    return i92c37b67685c58d9f60f7203b77a560478fb28fb9753bb4bee0d7e274eeb3e16.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Authentication()(*iefb47a3635416080293f661822201727ab9fba4b375bd5a3f60f0b9a80cca8b0.AuthenticationRequestBuilder) {
    return iefb47a3635416080293f661822201727ab9fba4b375bd5a3f60f0b9a80cca8b0.NewAuthenticationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Calendar()(*i31922605eeb11427d2cf37ddfee7f2740fc6e17cd21d0e9ee6f853fe1b3e4632.CalendarRequestBuilder) {
    return i31922605eeb11427d2cf37ddfee7f2740fc6e17cd21d0e9ee6f853fe1b3e4632.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) CalendarGroups()(*i441da951c572a925ed6c67355cefdd1ddc3a435432abc802841187bcd7831f08.CalendarGroupsRequestBuilder) {
    return i441da951c572a925ed6c67355cefdd1ddc3a435432abc802841187bcd7831f08.NewCalendarGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item collection
func (m *UserRequestBuilder) CalendarGroupsById(id string)(*i322200ffec4334c71536e3a74319b0afb2071fae07124a70501b4a91e5d3efc2.CalendarGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarGroup_id"] = id
    }
    return i322200ffec4334c71536e3a74319b0afb2071fae07124a70501b4a91e5d3efc2.NewCalendarGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) Calendars()(*ieafc69423dcb93faca91a98c5fff1ec743e967e077b14aad34d9518731effa7f.CalendarsRequestBuilder) {
    return ieafc69423dcb93faca91a98c5fff1ec743e967e077b14aad34d9518731effa7f.NewCalendarsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item collection
func (m *UserRequestBuilder) CalendarsById(id string)(*i1a9eb6fb24429a79e1e7397bd609e240ae9eef7f551f559e57f6c0050aa81ec7.CalendarRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendar_id"] = id
    }
    return i1a9eb6fb24429a79e1e7397bd609e240ae9eef7f551f559e57f6c0050aa81ec7.NewCalendarRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) CalendarView()(*i2c42211f5ece49fe404bcc5fe4998da5209b6fd5a9e6ef29c937593bb6d65f32.CalendarViewRequestBuilder) {
    return i2c42211f5ece49fe404bcc5fe4998da5209b6fd5a9e6ef29c937593bb6d65f32.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarView.item collection
func (m *UserRequestBuilder) CalendarViewById(id string)(*i86e9fa8da4e357b5840f59b78cd49a5abce31f7a90096cc8b85047571e5a2b06.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i86e9fa8da4e357b5840f59b78cd49a5abce31f7a90096cc8b85047571e5a2b06.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) ChangePassword()(*ie33da1fb1c72ec4e3eff642e23951718af609226626243167de6d0dd6df287f4.ChangePasswordRequestBuilder) {
    return ie33da1fb1c72ec4e3eff642e23951718af609226626243167de6d0dd6df287f4.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Chats()(*i1c8bab6c329a747e2f1a7bc6de2b493e83d072621195eaffdf5348edafc8d0c7.ChatsRequestBuilder) {
    return i1c8bab6c329a747e2f1a7bc6de2b493e83d072621195eaffdf5348edafc8d0c7.NewChatsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChatsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.chats.item collection
func (m *UserRequestBuilder) ChatsById(id string)(*i8bc97b715d2b1f741be45e9db03b6f18317426bb77532976ccdc40d366389a11.ChatRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chat_id"] = id
    }
    return i8bc97b715d2b1f741be45e9db03b6f18317426bb77532976ccdc40d366389a11.NewChatRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) CheckMemberGroups()(*i25e2af6a5262c1ae60708cc14abab6c2ec1e90563db9f056ec390327d5d23afe.CheckMemberGroupsRequestBuilder) {
    return i25e2af6a5262c1ae60708cc14abab6c2ec1e90563db9f056ec390327d5d23afe.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) CheckMemberObjects()(*i8df6ede5cb15274ffb3a3341b8e1b560c67d6eca03c4eafb47c23266ca7b8809.CheckMemberObjectsRequestBuilder) {
    return i8df6ede5cb15274ffb3a3341b8e1b560c67d6eca03c4eafb47c23266ca7b8809.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserRequestBuilder instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *UserRequestBuilder) ContactFolders()(*i6bc016bb2cb82b3c7d804b1443cbf5f98e5cb3ae65246c4f7658c00f6dde5db6.ContactFoldersRequestBuilder) {
    return i6bc016bb2cb82b3c7d804b1443cbf5f98e5cb3ae65246c4f7658c00f6dde5db6.NewContactFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactFoldersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.contactFolders.item collection
func (m *UserRequestBuilder) ContactFoldersById(id string)(*ib1c924a2a83336c2eee6c85a7c0edc6e41999bab4aa8e2db2755b64356506e43.ContactFolderRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contactFolder_id"] = id
    }
    return ib1c924a2a83336c2eee6c85a7c0edc6e41999bab4aa8e2db2755b64356506e43.NewContactFolderRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) Contacts()(*i7d4d710e1c8a8fc885d164dc49223af3ba44bf6a062cfa0019953387e5b65644.ContactsRequestBuilder) {
    return i7d4d710e1c8a8fc885d164dc49223af3ba44bf6a062cfa0019953387e5b65644.NewContactsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.contacts.item collection
func (m *UserRequestBuilder) ContactsById(id string)(*i66ea18aba758406b3b710707db096aa499334ab8141ff00b5817c0160d62f849.ContactRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contact_id"] = id
    }
    return i66ea18aba758406b3b710707db096aa499334ab8141ff00b5817c0160d62f849.NewContactRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete entity from users
func (m *UserRequestBuilder) CreateDeleteRequestInformation(options *UserRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *UserRequestBuilder) CreatedObjects()(*i396aa46f4b0d835f500bf48230128cbf9b0b575b29db3501c7da923e86cbb886.CreatedObjectsRequestBuilder) {
    return i396aa46f4b0d835f500bf48230128cbf9b0b575b29db3501c7da923e86cbb886.NewCreatedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get entity from users by key
func (m *UserRequestBuilder) CreateGetRequestInformation(options *UserRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in users
func (m *UserRequestBuilder) CreatePatchRequestInformation(options *UserRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from users
func (m *UserRequestBuilder) Delete(options *UserRequestBuilderDeleteOptions)(error) {
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
func (m *UserRequestBuilder) DeviceEnrollmentConfigurations()(*i0fbc6bb6d8e908a7211206db1e3f715af4ae676f093155141c801afcbfd4a468.DeviceEnrollmentConfigurationsRequestBuilder) {
    return i0fbc6bb6d8e908a7211206db1e3f715af4ae676f093155141c801afcbfd4a468.NewDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceEnrollmentConfigurationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.deviceEnrollmentConfigurations.item collection
func (m *UserRequestBuilder) DeviceEnrollmentConfigurationsById(id string)(*i20d5bc3e0beaac35f1cb30b0824d071deea9c98aa2da0ea95797426ec3391a8b.DeviceEnrollmentConfigurationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceEnrollmentConfiguration_id"] = id
    }
    return i20d5bc3e0beaac35f1cb30b0824d071deea9c98aa2da0ea95797426ec3391a8b.NewDeviceEnrollmentConfigurationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) DeviceManagementTroubleshootingEvents()(*i9fa6f9d397d0318d3e3fc1137d90e8302a95a72c11d6237f67fc7bd8dc1625b7.DeviceManagementTroubleshootingEventsRequestBuilder) {
    return i9fa6f9d397d0318d3e3fc1137d90e8302a95a72c11d6237f67fc7bd8dc1625b7.NewDeviceManagementTroubleshootingEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceManagementTroubleshootingEventsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.deviceManagementTroubleshootingEvents.item collection
func (m *UserRequestBuilder) DeviceManagementTroubleshootingEventsById(id string)(*i4d1bb8d1a6e444b2e99d643949896393d13341c23c9e59170d7cf747778a84f1.DeviceManagementTroubleshootingEventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementTroubleshootingEvent_id"] = id
    }
    return i4d1bb8d1a6e444b2e99d643949896393d13341c23c9e59170d7cf747778a84f1.NewDeviceManagementTroubleshootingEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) Devices()(*ib2e97a67ea231c420ac9fa60c240e9a1314dc844268c406a4c1600c693ad6811.DevicesRequestBuilder) {
    return ib2e97a67ea231c420ac9fa60c240e9a1314dc844268c406a4c1600c693ad6811.NewDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DevicesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.devices.item collection
func (m *UserRequestBuilder) DevicesById(id string)(*i8a37b6fd0dca54e1e3129d6fa85360b47ae2206831aad7d8bf6c3823b1fb540b.DeviceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["device_id"] = id
    }
    return i8a37b6fd0dca54e1e3129d6fa85360b47ae2206831aad7d8bf6c3823b1fb540b.NewDeviceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) DirectReports()(*i1e8437839d244c8d93e242e55290d8df8b03ca76cf9b009f9b9a19ef5d5922ac.DirectReportsRequestBuilder) {
    return i1e8437839d244c8d93e242e55290d8df8b03ca76cf9b009f9b9a19ef5d5922ac.NewDirectReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Drive()(*i3e6a9a935ee43b6bd89f378e203c324f1d517ccc9d1c30202303eafc13d5bb25.DriveRequestBuilder) {
    return i3e6a9a935ee43b6bd89f378e203c324f1d517ccc9d1c30202303eafc13d5bb25.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Drives()(*i00eee5f4e45952d7af6d7dbf02a5397275cbfa69444a9cd54a450f1cc9c5baf6.DrivesRequestBuilder) {
    return i00eee5f4e45952d7af6d7dbf02a5397275cbfa69444a9cd54a450f1cc9c5baf6.NewDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DrivesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item collection
func (m *UserRequestBuilder) DrivesById(id string)(*id15759ba2587bc2a0cac47e8d7da5e4bd208929d38688edb78df7122421e531f.DriveRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive_id"] = id
    }
    return id15759ba2587bc2a0cac47e8d7da5e4bd208929d38688edb78df7122421e531f.NewDriveRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) Events()(*i0bb45fabb973cc1726799c26601b61fb5d4ceef85239469bf8a48d255ddcab77.EventsRequestBuilder) {
    return i0bb45fabb973cc1726799c26601b61fb5d4ceef85239469bf8a48d255ddcab77.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.events.item collection
func (m *UserRequestBuilder) EventsById(id string)(*i1fb04953ffbe8174524af5da51d2e85d0602b9611bf6179631d63cd31623aa9c.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i1fb04953ffbe8174524af5da51d2e85d0602b9611bf6179631d63cd31623aa9c.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ExportDeviceAndAppManagementData builds and executes requests for operations under \users\{user-id}\microsoft.graph.exportDeviceAndAppManagementData()
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*i6c6abf877b483e321abb71bac38e5c8c592f7ee6915c9ad28447c321b2979b4d.ExportDeviceAndAppManagementDataRequestBuilder) {
    return i6c6abf877b483e321abb71bac38e5c8c592f7ee6915c9ad28447c321b2979b4d.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop builds and executes requests for operations under \users\{user-id}\microsoft.graph.exportDeviceAndAppManagementData(skip={skip},top={top})
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*i9e9fd6de702827c484559c38e811567b4bfa805d681aca1a1d7d8d8ba07aae9c.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return i9e9fd6de702827c484559c38e811567b4bfa805d681aca1a1d7d8d8ba07aae9c.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
func (m *UserRequestBuilder) ExportPersonalData()(*i10209655736fb29f6fe332f57065da684a4bb7e75afca36c8b0cf9215d9600cf.ExportPersonalDataRequestBuilder) {
    return i10209655736fb29f6fe332f57065da684a4bb7e75afca36c8b0cf9215d9600cf.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Extensions()(*i49c0a0951cef8133c800eac5f2bf2ddd282c66d3a73fc9f9c652fc532a3331be.ExtensionsRequestBuilder) {
    return i49c0a0951cef8133c800eac5f2bf2ddd282c66d3a73fc9f9c652fc532a3331be.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.extensions.item collection
func (m *UserRequestBuilder) ExtensionsById(id string)(*i532345ba6b7192bf733f1b4b3de662c47c1e7e0d3989701bcdc0d8f857dbe580.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i532345ba6b7192bf733f1b4b3de662c47c1e7e0d3989701bcdc0d8f857dbe580.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) FindMeetingTimes()(*i5c0e1410fbc34cac508ae33c37b1e9dd9a4ca3540ad1525e02f20329f8a264f0.FindMeetingTimesRequestBuilder) {
    return i5c0e1410fbc34cac508ae33c37b1e9dd9a4ca3540ad1525e02f20329f8a264f0.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists builds and executes requests for operations under \users\{user-id}\microsoft.graph.findRoomLists()
func (m *UserRequestBuilder) FindRoomLists()(*id5b765e74db4b3dcebe232c27216355fdc0568a1094952e7fc7cbdc858f3d81a.FindRoomListsRequestBuilder) {
    return id5b765e74db4b3dcebe232c27216355fdc0568a1094952e7fc7cbdc858f3d81a.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms builds and executes requests for operations under \users\{user-id}\microsoft.graph.findRooms()
func (m *UserRequestBuilder) FindRooms()(*i09b3ed8449f0eb014b1a1c391d9057c99dca194b7d9cee6d67e587c2b4d66363.FindRoomsRequestBuilder) {
    return i09b3ed8449f0eb014b1a1c391d9057c99dca194b7d9cee6d67e587c2b4d66363.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList builds and executes requests for operations under \users\{user-id}\microsoft.graph.findRooms(RoomList='{RoomList}')
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i933c9dccaf05a6a7fcf066da46cae8d6c4a0940e841426651ace8702209f1b9d.FindRoomsWithRoomListRequestBuilder) {
    return i933c9dccaf05a6a7fcf066da46cae8d6c4a0940e841426651ace8702209f1b9d.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
func (m *UserRequestBuilder) FollowedSites()(*id904363322fca79446fa0f0f2231760c6649433886f1c8e79503338b004b288c.FollowedSitesRequestBuilder) {
    return id904363322fca79446fa0f0f2231760c6649433886f1c8e79503338b004b288c.NewFollowedSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get entity from users by key
func (m *UserRequestBuilder) Get(options *UserRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.User, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUser() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.User), nil
}
// GetEffectiveDeviceEnrollmentConfigurations builds and executes requests for operations under \users\{user-id}\microsoft.graph.getEffectiveDeviceEnrollmentConfigurations()
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*i44184067b9a5d3386d10a8d3d78383c99a1f94c6db7a1c78fecc6277422e9842.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return i44184067b9a5d3386d10a8d3d78383c99a1f94c6db7a1c78fecc6277422e9842.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices builds and executes requests for operations under \users\{user-id}\microsoft.graph.getLoggedOnManagedDevices()
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*i163ca73c2d1426f49f134d87416999c0644a5bf7f78d9912f8098a9785dc6ec1.GetLoggedOnManagedDevicesRequestBuilder) {
    return i163ca73c2d1426f49f134d87416999c0644a5bf7f78d9912f8098a9785dc6ec1.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) GetMailTips()(*i89f9ebf7c167c75aff01142f1e2b30d94fcb3ccf3d4af1752e74dbc3f979ecd6.GetMailTipsRequestBuilder) {
    return i89f9ebf7c167c75aff01142f1e2b30d94fcb3ccf3d4af1752e74dbc3f979ecd6.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses builds and executes requests for operations under \users\{user-id}\microsoft.graph.getManagedAppDiagnosticStatuses()
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*idaeea547224afeb26277747dcbdab987622f0aec250f6cb9dcebbc9e6d05df1e.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return idaeea547224afeb26277747dcbdab987622f0aec250f6cb9dcebbc9e6d05df1e.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies builds and executes requests for operations under \users\{user-id}\microsoft.graph.getManagedAppPolicies()
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i1b6f500d8aa60be2a2a1e7586a07867f5e80bc6da6a793388efc074011aa9f1b.GetManagedAppPoliciesRequestBuilder) {
    return i1b6f500d8aa60be2a2a1e7586a07867f5e80bc6da6a793388efc074011aa9f1b.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures builds and executes requests for operations under \users\{user-id}\microsoft.graph.getManagedDevicesWithAppFailures()
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*i028859f982626d2101b8d8b68caff0f04ef34796aed805e993e21fefb9152760.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return i028859f982626d2101b8d8b68caff0f04ef34796aed805e993e21fefb9152760.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps builds and executes requests for operations under \users\{user-id}\microsoft.graph.getManagedDevicesWithFailedOrPendingApps()
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*i52ea864f44c605939329148c5a95626ead1b771dff88c6072de654caee1b281e.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return i52ea864f44c605939329148c5a95626ead1b771dff88c6072de654caee1b281e.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) GetMemberGroups()(*i633ead85dd41f6f9d94e9cf71e43f604b18195cef363d66f743366cdb820fc8b.GetMemberGroupsRequestBuilder) {
    return i633ead85dd41f6f9d94e9cf71e43f604b18195cef363d66f743366cdb820fc8b.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) GetMemberObjects()(*i751d3667ec77ce3a4b6544375d44e850f38ac8db37e3237ce66245e949ea015e.GetMemberObjectsRequestBuilder) {
    return i751d3667ec77ce3a4b6544375d44e850f38ac8db37e3237ce66245e949ea015e.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) InferenceClassification()(*i73641486333af15741c085f42cd9ba7b087800533a2f2a8781b636a965272506.InferenceClassificationRequestBuilder) {
    return i73641486333af15741c085f42cd9ba7b087800533a2f2a8781b636a965272506.NewInferenceClassificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) InformationProtection()(*i992f3c4694021035f11bae1cb737eb6020f9d1221bc7676d2685a415cab74221.InformationProtectionRequestBuilder) {
    return i992f3c4694021035f11bae1cb737eb6020f9d1221bc7676d2685a415cab74221.NewInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Insights()(*i5f232e4d902d95d06ec20d5128cce05b56d53758222f6b28e307a6b909082dac.InsightsRequestBuilder) {
    return i5f232e4d902d95d06ec20d5128cce05b56d53758222f6b28e307a6b909082dac.NewInsightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*i7f0c66e611b8af22dffef0886d0dc8468dc8cac8bb0b577016f3d4b1daf70528.InvalidateAllRefreshTokensRequestBuilder) {
    return i7f0c66e611b8af22dffef0886d0dc8468dc8cac8bb0b577016f3d4b1daf70528.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked builds and executes requests for operations under \users\{user-id}\microsoft.graph.isManagedAppUserBlocked()
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*ieee72b0b162c07d3dde8ae1950af04718a5801d0d3e6575e0cba2faa3f56157a.IsManagedAppUserBlockedRequestBuilder) {
    return ieee72b0b162c07d3dde8ae1950af04718a5801d0d3e6575e0cba2faa3f56157a.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) JoinedGroups()(*i235dd448b749bb3a1187ec0970cdeb5c226f7aee1c9e6c72733a9d1a15182ffe.JoinedGroupsRequestBuilder) {
    return i235dd448b749bb3a1187ec0970cdeb5c226f7aee1c9e6c72733a9d1a15182ffe.NewJoinedGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JoinedGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item collection
func (m *UserRequestBuilder) JoinedGroupsById(id string)(*ie39c35482440e2a813ce2612b19ea08745ca523bb7a24088144c3428169a006d.GroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["group_id"] = id
    }
    return ie39c35482440e2a813ce2612b19ea08745ca523bb7a24088144c3428169a006d.NewGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) JoinedTeams()(*i4e8ad44991b185083366e107199ff13f6438a07cafd58ae36e0e862692bf9826.JoinedTeamsRequestBuilder) {
    return i4e8ad44991b185083366e107199ff13f6438a07cafd58ae36e0e862692bf9826.NewJoinedTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) LicenseDetails()(*i53e4d773c077f49b4590da13b904e6a502ef1f8f43b7a7fdfaba256a667df083.LicenseDetailsRequestBuilder) {
    return i53e4d773c077f49b4590da13b904e6a502ef1f8f43b7a7fdfaba256a667df083.NewLicenseDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LicenseDetailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.licenseDetails.item collection
func (m *UserRequestBuilder) LicenseDetailsById(id string)(*i53e4d773c077f49b4590da13b904e6a502ef1f8f43b7a7fdfaba256a667df083.LicenseDetailsRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["licenseDetails_id"] = id
    }
    return i53e4d773c077f49b4590da13b904e6a502ef1f8f43b7a7fdfaba256a667df083.NewLicenseDetailsRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) MailFolders()(*i21282294d2d1bb2a0ab74695618e809021209e6abac85e175b427b95d4891cc5.MailFoldersRequestBuilder) {
    return i21282294d2d1bb2a0ab74695618e809021209e6abac85e175b427b95d4891cc5.NewMailFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MailFoldersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mailFolders.item collection
func (m *UserRequestBuilder) MailFoldersById(id string)(*i11a0d65a9c2024ff5c04a9b3cb79486d6f2277859e51094a083537f2289c0a27.MailFolderRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mailFolder_id"] = id
    }
    return i11a0d65a9c2024ff5c04a9b3cb79486d6f2277859e51094a083537f2289c0a27.NewMailFolderRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) ManagedAppRegistrations()(*i44dc4d21a51a483da0eb7ac0bc6c38fa258f1e982bb140f0bd9282a87d39015c.ManagedAppRegistrationsRequestBuilder) {
    return i44dc4d21a51a483da0eb7ac0bc6c38fa258f1e982bb140f0bd9282a87d39015c.NewManagedAppRegistrationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) ManagedDevices()(*i4cd644380b450eedb89634c647645afa6ede7ee4df886587de78ccf69b3a9a42.ManagedDevicesRequestBuilder) {
    return i4cd644380b450eedb89634c647645afa6ede7ee4df886587de78ccf69b3a9a42.NewManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDevicesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.managedDevices.item collection
func (m *UserRequestBuilder) ManagedDevicesById(id string)(*i0796e8f7b9f803857a438202c35d9fccd548d4779e6f2a341287440bf41acc70.ManagedDeviceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDevice_id"] = id
    }
    return i0796e8f7b9f803857a438202c35d9fccd548d4779e6f2a341287440bf41acc70.NewManagedDeviceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) Manager()(*ice23d35988441664e0a80c4309242575aeb7cc7c84c7bfc6e9ff86ce92c7cde5.ManagerRequestBuilder) {
    return ice23d35988441664e0a80c4309242575aeb7cc7c84c7bfc6e9ff86ce92c7cde5.NewManagerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) MemberOf()(*i35ae79fec43abc52b415d013c5ba088edde69bf589391525fec20b1f062c22a7.MemberOfRequestBuilder) {
    return i35ae79fec43abc52b415d013c5ba088edde69bf589391525fec20b1f062c22a7.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Messages()(*ifc52bce07b756c55dabe5cd8f6a0d8c749809ddcc3831f5bded6c944d449dbd8.MessagesRequestBuilder) {
    return ifc52bce07b756c55dabe5cd8f6a0d8c749809ddcc3831f5bded6c944d449dbd8.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.messages.item collection
func (m *UserRequestBuilder) MessagesById(id string)(*i92f204ee1ec183e477ad6aae362cad151b598b7439eda0c58c4750c295bc84e8.MessageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["message_id"] = id
    }
    return i92f204ee1ec183e477ad6aae362cad151b598b7439eda0c58c4750c295bc84e8.NewMessageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) MobileAppIntentAndStates()(*ide4da4ea5e9d5b4a3ed2b815f5032fc4321b390d6e1c0ba686af0a35ec886a84.MobileAppIntentAndStatesRequestBuilder) {
    return ide4da4ea5e9d5b4a3ed2b815f5032fc4321b390d6e1c0ba686af0a35ec886a84.NewMobileAppIntentAndStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileAppIntentAndStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mobileAppIntentAndStates.item collection
func (m *UserRequestBuilder) MobileAppIntentAndStatesById(id string)(*i9e371fd6357698c3cd300fbb8012addd3ec298d5df1be3dfa58412fb63c3cae8.MobileAppIntentAndStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppIntentAndState_id"] = id
    }
    return i9e371fd6357698c3cd300fbb8012addd3ec298d5df1be3dfa58412fb63c3cae8.NewMobileAppIntentAndStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) MobileAppTroubleshootingEvents()(*i7358087738ea2e8de75562a191c6ccb72c9a24d79c1f9a33de77f140cbb5ce74.MobileAppTroubleshootingEventsRequestBuilder) {
    return i7358087738ea2e8de75562a191c6ccb72c9a24d79c1f9a33de77f140cbb5ce74.NewMobileAppTroubleshootingEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileAppTroubleshootingEventsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mobileAppTroubleshootingEvents.item collection
func (m *UserRequestBuilder) MobileAppTroubleshootingEventsById(id string)(*i7fdae8345bbd328abce3eca8f17ca05dfb12b3a6a3adcb8b7b394f3697581dae.MobileAppTroubleshootingEventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppTroubleshootingEvent_id"] = id
    }
    return i7fdae8345bbd328abce3eca8f17ca05dfb12b3a6a3adcb8b7b394f3697581dae.NewMobileAppTroubleshootingEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) Notifications()(*i5c911b116a9e04c78f811a8141c00da4569a7dc37a1d8590843bc9d14e30377e.NotificationsRequestBuilder) {
    return i5c911b116a9e04c78f811a8141c00da4569a7dc37a1d8590843bc9d14e30377e.NewNotificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotificationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.notifications.item collection
func (m *UserRequestBuilder) NotificationsById(id string)(*i5d09ee5a147e62e93172ba3c3003b1e37657e5078803e7fa2be2756f303a9607.NotificationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["notification_id"] = id
    }
    return i5d09ee5a147e62e93172ba3c3003b1e37657e5078803e7fa2be2756f303a9607.NewNotificationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) Oauth2PermissionGrants()(*ib674e5ebc7e9a35ea83a7a23737e2e850d21d009460f455398f4235cd71eb2d8.Oauth2PermissionGrantsRequestBuilder) {
    return ib674e5ebc7e9a35ea83a7a23737e2e850d21d009460f455398f4235cd71eb2d8.NewOauth2PermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Onenote()(*i4c553c01d10dc5ea91f5e8443a25576af37822a1a9e3b20c39f7d69613ca1a5d.OnenoteRequestBuilder) {
    return i4c553c01d10dc5ea91f5e8443a25576af37822a1a9e3b20c39f7d69613ca1a5d.NewOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) OnlineMeetings()(*idddb11dcacd5585b4b6cb080f05aa0ce5c62d39e09bc0ffcd012d5c8c731810b.OnlineMeetingsRequestBuilder) {
    return idddb11dcacd5585b4b6cb080f05aa0ce5c62d39e09bc0ffcd012d5c8c731810b.NewOnlineMeetingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnlineMeetingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.onlineMeetings.item collection
func (m *UserRequestBuilder) OnlineMeetingsById(id string)(*i8e9bfc4860807cf4727d9197394851201c63af67494fa725e469b158152f82c0.OnlineMeetingRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onlineMeeting_id"] = id
    }
    return i8e9bfc4860807cf4727d9197394851201c63af67494fa725e469b158152f82c0.NewOnlineMeetingRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) Outlook()(*ib9ea7c893835d89864cc7cd939d24f579ba191ab540031694fe5202498a4240c.OutlookRequestBuilder) {
    return ib9ea7c893835d89864cc7cd939d24f579ba191ab540031694fe5202498a4240c.NewOutlookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) OwnedDevices()(*i71ece4d9b62444d7cd9edeac6c2b5bdaf0f35379d154d089d602c3d246239739.OwnedDevicesRequestBuilder) {
    return i71ece4d9b62444d7cd9edeac6c2b5bdaf0f35379d154d089d602c3d246239739.NewOwnedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) OwnedObjects()(*i3aa2a5875223f6b19560c0365147205688fff76fe30e6bbf724b3b25f7b00a43.OwnedObjectsRequestBuilder) {
    return i3aa2a5875223f6b19560c0365147205688fff76fe30e6bbf724b3b25f7b00a43.NewOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in users
func (m *UserRequestBuilder) Patch(options *UserRequestBuilderPatchOptions)(error) {
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
func (m *UserRequestBuilder) PendingAccessReviewInstances()(*i3be6b130d8bf10a5df104376e126ee2c3584be236591382ff88662fd6a574158.PendingAccessReviewInstancesRequestBuilder) {
    return i3be6b130d8bf10a5df104376e126ee2c3584be236591382ff88662fd6a574158.NewPendingAccessReviewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PendingAccessReviewInstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.pendingAccessReviewInstances.item collection
func (m *UserRequestBuilder) PendingAccessReviewInstancesById(id string)(*i2cf542b96fcffbafb31190c7c7e24f2b652b0249d3c7d0e6deaf71e333abf373.AccessReviewInstanceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewInstance_id"] = id
    }
    return i2cf542b96fcffbafb31190c7c7e24f2b652b0249d3c7d0e6deaf71e333abf373.NewAccessReviewInstanceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) People()(*i1faa710c227b8e4b8d7abef5e9611e6604d3f4265db3da621d2ddd654a3a33f4.PeopleRequestBuilder) {
    return i1faa710c227b8e4b8d7abef5e9611e6604d3f4265db3da621d2ddd654a3a33f4.NewPeopleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PeopleById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.people.item collection
func (m *UserRequestBuilder) PeopleById(id string)(*ia16782564e6f61796daddf42e21afd9855b9a1c1c41eef4e65e545763ac7c8d5.PersonRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["person_id"] = id
    }
    return ia16782564e6f61796daddf42e21afd9855b9a1c1c41eef4e65e545763ac7c8d5.NewPersonRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) Photo()(*i5f9e5aa5c251065cc59079614790107bc87c70c65435cea5312a9ad92e27cad3.PhotoRequestBuilder) {
    return i5f9e5aa5c251065cc59079614790107bc87c70c65435cea5312a9ad92e27cad3.NewPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Photos()(*i58ea06623816d885f9dc5730c5fd341b9dc32ca959bf1ffe118d49d00bdfbc62.PhotosRequestBuilder) {
    return i58ea06623816d885f9dc5730c5fd341b9dc32ca959bf1ffe118d49d00bdfbc62.NewPhotosRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PhotosById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.photos.item collection
func (m *UserRequestBuilder) PhotosById(id string)(*ib39beceac49e7dd17ed1c0885887dae465a45afa34b3eebb70aa5fc08ca691e7.ProfilePhotoRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["profilePhoto_id"] = id
    }
    return ib39beceac49e7dd17ed1c0885887dae465a45afa34b3eebb70aa5fc08ca691e7.NewProfilePhotoRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) Planner()(*i64d17e545a98d6600890650331d10cc8ac9bafd9f1a9888ceee3f7c9932f9d68.PlannerRequestBuilder) {
    return i64d17e545a98d6600890650331d10cc8ac9bafd9f1a9888ceee3f7c9932f9d68.NewPlannerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Presence()(*i2a5b8dfcb3b418ee303ee5f5b01aebd3209cdfa9122c52e81bd0a9e16a726f9d.PresenceRequestBuilder) {
    return i2a5b8dfcb3b418ee303ee5f5b01aebd3209cdfa9122c52e81bd0a9e16a726f9d.NewPresenceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Profile()(*i77129a8685763d465ea62e5de870203f0d38a41c496b1f718e0c695d6df2b42c.ProfileRequestBuilder) {
    return i77129a8685763d465ea62e5de870203f0d38a41c496b1f718e0c695d6df2b42c.NewProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) RegisteredDevices()(*i9f79453a6cb43e9e5e68087f1d6e774414d8ba83449733d805265bf8cb3d0b5d.RegisteredDevicesRequestBuilder) {
    return i9f79453a6cb43e9e5e68087f1d6e774414d8ba83449733d805265bf8cb3d0b5d.NewRegisteredDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime builds and executes requests for operations under \users\{user-id}\microsoft.graph.reminderView(StartDateTime='{StartDateTime}',EndDateTime='{EndDateTime}')
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(startDateTime *string, endDateTime *string)(*i1e2ed1761df19ea44d2d90f6fdfba02f95eb595495e7804f87126e72ee9345e3.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i1e2ed1761df19ea44d2d90f6fdfba02f95eb595495e7804f87126e72ee9345e3.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, startDateTime, endDateTime);
}
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i7d78da96ae198e74cdcd846ebd00d688295a8b346c939e4e49d56b88b5f72711.RemoveAllDevicesFromManagementRequestBuilder) {
    return i7d78da96ae198e74cdcd846ebd00d688295a8b346c939e4e49d56b88b5f72711.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i35b254e720ef3a77057a7b6736d2d9841a2d361b2b095e7a189fd80751427060.ReprocessLicenseAssignmentRequestBuilder) {
    return i35b254e720ef3a77057a7b6736d2d9841a2d361b2b095e7a189fd80751427060.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Restore()(*i6d8c6f6a8d72bfe9260edd2b3126823191856ab8f78dd9cd5a1c01fb0fff9c47.RestoreRequestBuilder) {
    return i6d8c6f6a8d72bfe9260edd2b3126823191856ab8f78dd9cd5a1c01fb0fff9c47.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) RevokeSignInSessions()(*i6545b8651bf39e658fd8159fa1123831d6f44a855626360fac6873d3a90581c8.RevokeSignInSessionsRequestBuilder) {
    return i6545b8651bf39e658fd8159fa1123831d6f44a855626360fac6873d3a90581c8.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) ScopedRoleMemberOf()(*if3e915183aa6002a59a5c947830462a91684160d445bda34aa2c4cf6886b82b2.ScopedRoleMemberOfRequestBuilder) {
    return if3e915183aa6002a59a5c947830462a91684160d445bda34aa2c4cf6886b82b2.NewScopedRoleMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScopedRoleMemberOfById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.scopedRoleMemberOf.item collection
func (m *UserRequestBuilder) ScopedRoleMemberOfById(id string)(*ie132362db0baae23ad9567ff4032a62c165d8386308f8baec97ecbb60c10a021.ScopedRoleMembershipRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["scopedRoleMembership_id"] = id
    }
    return ie132362db0baae23ad9567ff4032a62c165d8386308f8baec97ecbb60c10a021.NewScopedRoleMembershipRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) SendMail()(*id6953e096a9515e55975655319314bf3955019e0bded78bb5cb6fb32eb4e8e94.SendMailRequestBuilder) {
    return id6953e096a9515e55975655319314bf3955019e0bded78bb5cb6fb32eb4e8e94.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Settings()(*i7984c072c770ec17745d737738e8ab409b4f442b183221a451997db094cd8b13.SettingsRequestBuilder) {
    return i7984c072c770ec17745d737738e8ab409b4f442b183221a451997db094cd8b13.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Tasks()(*i807859e90e02be96c7951593aa65cbf418e7a30d0f62e3a47b8382296d35e2d2.TasksRequestBuilder) {
    return i807859e90e02be96c7951593aa65cbf418e7a30d0f62e3a47b8382296d35e2d2.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Teamwork()(*if429d108f0ce6b9f644fc8ece21d69b6a892d9788590f39dc3d9d21bda91a159.TeamworkRequestBuilder) {
    return if429d108f0ce6b9f644fc8ece21d69b6a892d9788590f39dc3d9d21bda91a159.NewTeamworkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) Todo()(*i4e66b80d61d13c58a5325be0d73998d6c528f982ee63fc876b66fdffd4ec5473.TodoRequestBuilder) {
    return i4e66b80d61d13c58a5325be0d73998d6c528f982ee63fc876b66fdffd4ec5473.NewTodoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) TransitiveMemberOf()(*i6f957db73ce02e759548f419eb79a2bfdfb04d6c5206523cd438181cc5619df8.TransitiveMemberOfRequestBuilder) {
    return i6f957db73ce02e759548f419eb79a2bfdfb04d6c5206523cd438181cc5619df8.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) TransitiveReports()(*ia1e1519710e771d8df5e22da39eeb50d37180582bb2065916e5cd0703a4d1cce.TransitiveReportsRequestBuilder) {
    return ia1e1519710e771d8df5e22da39eeb50d37180582bb2065916e5cd0703a4d1cce.NewTransitiveReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) TranslateExchangeIds()(*i87f894b87c2947e4dad88e845ac91e8a8ddf1ca5ba6b426384ab8873fe897913.TranslateExchangeIdsRequestBuilder) {
    return i87f894b87c2947e4dad88e845ac91e8a8ddf1ca5ba6b426384ab8873fe897913.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) UnblockManagedApps()(*i1d3f5d5cc5119c3710246df1436375c7f067bf4348a4cf5c5decdf0052f3b667.UnblockManagedAppsRequestBuilder) {
    return i1d3f5d5cc5119c3710246df1436375c7f067bf4348a4cf5c5decdf0052f3b667.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) UsageRights()(*iddc68c193cc623cc513e0675dadbd063a05b9d5ce8584eb5a7e8ebf24faa98e0.UsageRightsRequestBuilder) {
    return iddc68c193cc623cc513e0675dadbd063a05b9d5ce8584eb5a7e8ebf24faa98e0.NewUsageRightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsageRightsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.usageRights.item collection
func (m *UserRequestBuilder) UsageRightsById(id string)(*ib9adfa359ccdf30972cf1f2c9fb0c0a8edde337e58fe26f0a95800e7ce3ba9db.UsageRightRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["usageRight_id"] = id
    }
    return ib9adfa359ccdf30972cf1f2c9fb0c0a8edde337e58fe26f0a95800e7ce3ba9db.NewUsageRightRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) WindowsInformationProtectionDeviceRegistrations()(*i3eaf8bc51db6f4bd95448da3951002ebce4db39125b63eaa1cb9987f99da858f.WindowsInformationProtectionDeviceRegistrationsRequestBuilder) {
    return i3eaf8bc51db6f4bd95448da3951002ebce4db39125b63eaa1cb9987f99da858f.NewWindowsInformationProtectionDeviceRegistrationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*ieac84d2e3111366253e033fed8ced42323b1927cd610b6c93eb3eea6d066fd20.WipeAndBlockManagedAppsRequestBuilder) {
    return ieac84d2e3111366253e033fed8ced42323b1927cd610b6c93eb3eea6d066fd20.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*i099d32137e6ab567a080a6d574abebbf5b3e8b218fa6fda8a2a0fce52a347329.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return i099d32137e6ab567a080a6d574abebbf5b3e8b218fa6fda8a2a0fce52a347329.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i23fd320d20dc894e875e4225a9c0031362d1c191fcba8ad855a39c641e46a084.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i23fd320d20dc894e875e4225a9c0031362d1c191fcba8ad855a39c641e46a084.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
