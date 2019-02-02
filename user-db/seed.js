
db.users.insertMany([
    {"firstname":"Walter","lastname":"White","email":"walter@acmefitness.com","username":"walter","password":"6837ea9b06409112a824d113927ad74fabc5c76e","salt":""}
    ,{"firstname":"Dwight","lastname":"Schrute","email":"dwight@acmefitness.com","username":"dwight","password":"6837ea9b06409112a824d113927ad74fabc5c76e","salt":""}
    ,{"firstname":"Eric","lastname":"Cartman","email":"eric@acmefitness.com","username":"eric","password":"6837ea9b06409112a824d113927ad74fabc5c76e","salt":""}
    ,{"firstname":"Han","lastname":"Solo","email":"han@acmefitness.com","username":"han","password":"6837ea9b06409112a824d113927ad74fabc5c76e","salt":""}
    ,{"firstname":"Phoebe","lastname":"Buffay","email":"phoebe@acmefitness.com","username":"phoebe","password":"6837ea9b06409112a824d113927ad74fabc5c76e","salt":""}
    ,{"firstname":"Elaine","lastname":"Benes","email":"elaine@acmefitness.com","username":"elaine","password":"6837ea9b06409112a824d113927ad74fabc5c76e","salt":""}    
]);

db.catalog.insertMany([
    {"name":"Fitness special","description":"Limited edition fitness watch","picture":"/static/images/weave1.jpg","price":249.99000549316406,"tags":["watch"]}
    ,{"name":"Infuser water bottle","description":"For all those althelets out there, a perfect bottle to enrich you","picture":"/static/images/puma_1.jpeg","price":34.59000015258789,"tags":["bottle"]}
    ,{"name":"Classic Tread Mill","description":"Classic Tread Mill","picture":"/static/images/youtube_1.jpeg","price":800.0,"tags":["running"]}
    ,{"name":"Baby Tracker","description":"Baby Tracker","picture":"/static/images/catsocks_1.jpg","price":99.0,"tags":["baby"]}
    ,{"name":"Smart watch","description":"Smartest Watch Ever","picture":"/static/images/cross_1.jpeg","price":169.59,"tags":["watch"]}
    ,{"name":"Running shoes","description":"Best shoes","picture":"/static/images/cross_1.jpeg","price":120.00,"tags":["running"]}
    ,{"name":"Weights","description":"Lift weights","picture":"/static/images/cross_1.jpeg","price":49.99,"tags":["weight"]}
    ,{"name":"Folding Bicycle","description":"Bicycle that fold","picture":"/static/images/cross_1.jpeg","price":299.99,"tags":["bicycle"]}
]);