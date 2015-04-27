(function() {
	'use strict';
	var log = function(str) { console.log(str); };

	app.controller('IndexController', function($scope) {
		$scope.db = {}
		log("called!")
		$scope.db.items = [
  {
    "id":1,
    "name":{
      "first":"John",
      "last":"Schmidt"
    },
    "address":"45024 France",
    "price":760.41,
    "isActive":"Yes",
    "product":{
      "description":"Fried Potatoes",
      "options":[
        {
          "description":"Fried Potatoes",
          "image":"//a248.e.akamai.net/assets.github.com/images/icons/emoji/fries.png",
          "Pick$":null
        },
        {
          "description":"Fried Onions",
          "image":"//a248.e.akamai.net/assets.github.com/images/icons/emoji/fries.png",
          "Pick$":null
        }
      ]
    }
  },
  //more items go here
];
	})
}())
