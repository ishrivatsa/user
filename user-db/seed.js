
db.users.insertMany([
    {"firstname":"Walter","lastname":"White","email":"walter@acmefitness.com","username":"walter","password":"6837ea9b06409112a824d113927ad74fabc5c76e","salt":""}
    ,{"firstname":"Dwight","lastname":"Schrute","email":"dwight@acmefitness.com","username":"dwight","password":"6837ea9b06409112a824d113927ad74fabc5c76e","salt":""}
    ,{"firstname":"Eric","lastname":"Cartman","email":"eric@acmefitness.com","username":"eric","password":"6837ea9b06409112a824d113927ad74fabc5c76e","salt":""}
    ,{"firstname":"Han","lastname":"Solo","email":"han@acmefitness.com","username":"han","password":"6837ea9b06409112a824d113927ad74fabc5c76e","salt":""}
    ,{"firstname":"Phoebe","lastname":"Buffay","email":"phoebe@acmefitness.com","username":"phoebe","password":"6837ea9b06409112a824d113927ad74fabc5c76e","salt":""}
    ,{"firstname":"Elaine","lastname":"Benes","email":"elaine@acmefitness.com","username":"elaine","password":"6837ea9b06409112a824d113927ad74fabc5c76e","salt":""}    
]);

db.catalog.insertMany([
    {"name":"Yoga Mat","shortdescription":"Limited Edition Mat","description":"Limited edition yoga mat","imageurl1":"/static/images/yogamat_square.jpg","imageurl2":"/static/images/yogamat_square.jpg","imageurl3":"/static/images/bottle_square.jpg","price":62.5,"tags":["mat"]}
    ,{"name":"Water Bottle","shortdescription":"Best water bottle ever","description":"For all those athletes out there, a perfect bottle to enrich you","imageurl1":"/static/images/bottle_square.jpg","imageurl2":"/static/images/yogamat_square.jpg","imageurl3":"/static/images/bottle_square.jpg","price":34.9900016784668,"tags":["bottle"]}
    ,{"name":"Tread Mill","shortdescription":"Tread Mill","description":"Tread Mill","imageurl1":"/static/images/treadmill_square.jpg","imageurl2":"/static/images/yogamat_square.jpg","imageurl3":"/static/images/bottle_square.jpg","price":800.0,"tags":["running"]}
    ,{"name":"Smart Watch","shortdescription":"Smart watch","description":"Smart watch to keep you fit","imageurl1":"/static/images/smartwatch_square.jpg","imageurl2":"/static/images/yogamat_square.jpg","imageurl3":"/static/images/bottle_square.jpg","price":399.5899963378906,"tags":["watch"]}
    ,{"name":"Red Pant","shortdescription":"Red pant", "description":"Special Item on Sale","imageurl1":"/static/images/redpants_square.jpg","imageurl2":"/static/images/yogamat_square.jpg","imageurl3":"/static/images/bottle_square.jpg", "price":99.0,"tags":["clothing"]}
    ,{"name":"Running shoes","shortdescription":"Running Shoes", "description":"Best shoes","imageurl1":"/static/images/shoes_square.jpg","imageurl2":"/static/images/yogamat_square.jpg","imageurl3":"/static/images/bottle_square.jpg", "price":120.00,"tags":["running"]}
    ,{"name":"Weights","shortdescription":"Weights","description":"Lift weights","imageurl1":"/static/images/weights_square.jpg","imageurl2":"/static/images/yogamat_square.jpg","imageurl3":"/static/images/bottle_square.jpg", "price":49.99,"tags":["weight"]}
    ,{"name":"Fit Bike","shortdescription":"Bicycle", "description":"Amazing Bicycle","imageurl1":"/static/images/bicycle_square.jpg","imageurl2":"/static/images/bottle_square.jpg","imageurl3":"/static/images/bottle_square.jpg", "price":499.99,"tags":["bicycle"]}
]);