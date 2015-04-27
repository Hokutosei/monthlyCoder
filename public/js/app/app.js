'use strict';

var app = angular.module('monthlyCoder', ['ngRoute', 'ngResource', 'react', 'ngHandsontable']);

app.config(function($routeProvider, $locationProvider) {

	$routeProvider
		.when('/', {
			templateUrl: 'js/app/index/template/' + 'index.html',
			controller: 'IndexController'
		})
		.when('/delay', {
			templateUrl: 'js/app/delay/template/index.html',
			controller: 'DelayController'
		})
		.otherwise({
			redirectTo: '/'
		})

	$locationProvider.html5Mode({
		enabled: true,
		requireBase: true
	})
})
