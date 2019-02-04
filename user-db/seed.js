
db.users.insertMany([
    {"firstname":"Walter","lastname":"White","email":"walter@acmefitness.com","username":"walter","password":"6837ea9b06409112a824d113927ad74fabc5c76e","salt":""}
    ,{"firstname":"Dwight","lastname":"Schrute","email":"dwight@acmefitness.com","username":"dwight","password":"6837ea9b06409112a824d113927ad74fabc5c76e","salt":""}
    ,{"firstname":"Eric","lastname":"Cartman","email":"eric@acmefitness.com","username":"eric","password":"6837ea9b06409112a824d113927ad74fabc5c76e","salt":""}
    ,{"firstname":"Han","lastname":"Solo","email":"han@acmefitness.com","username":"han","password":"6837ea9b06409112a824d113927ad74fabc5c76e","salt":""}
    ,{"firstname":"Phoebe","lastname":"Buffay","email":"phoebe@acmefitness.com","username":"phoebe","password":"6837ea9b06409112a824d113927ad74fabc5c76e","salt":""}
    ,{"firstname":"Elaine","lastname":"Benes","email":"elaine@acmefitness.com","username":"elaine","password":"6837ea9b06409112a824d113927ad74fabc5c76e","salt":""}    
]);

db.catalog.insertMany([
    {"name":"Yoga Mat","description":"Limited edition Mat","picture":"/static/images/yogamat_square.jpg","price":249.99000549316406,"tags":["mat"]}
    ,{"name":"Water bottle","description":"For all those althelets out there, a perfect bottle to enrich you","picture":"/static/images/bottle_square.jpg","price":34.59000015258789,"tags":["bottle"]}
    ,{"name":"Tread Mill","description":"Tread Mill","picture":"/static/images/treadmill_square.jpg","price":800.0,"tags":["running"]}
    ,{"name":"Red Pant","description":"Special Item on Sale","picture":"/static/images/redpants_square.jpg","price":99.0,"tags":["clothing"]}
    ,{"name":"Smart watch","description":"Smartest Watch Ever","picture":"/static/images/smartwatch_square.jpg","price":399.59,"tags":["watch"]}
    ,{"name":"Running shoes","description":"Best shoes","picture":"/static/images/shoes_square.jpg","price":120.00,"tags":["running"]}
    ,{"name":"Weights","description":"Lift weights","picture":"/static/images/weights_square.jpg","price":49.99,"tags":["weight"]}
    ,{"name":"Fit Bike","description":"Bicycle that fold","picture":"/static/images/bicycle_square.jpg","price":499.99,"tags":["bicycle"]}
]);