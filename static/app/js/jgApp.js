'use strict';

// Declare app level module which depends on filters, and services
angular.module('jgApp', [/*'jgApp.filters',*/ 'jgApp.services', /*'jgApp.directives',*/ 'jgApp.controllers']).
  config(['$routeProvider', function($routeProvider) {
    $routeProvider.when('/', {
      templateUrl: 'partials/index.html',
      controller: 'IndexCtrl'});
    $routeProvider.when('/tags', {
      templateUrl: 'partials/tags.html',
      controller: 'TagsCtrl'});
    $routeProvider.when('/tag/:tagId', {
      templateUrl: 'partials/tag.html',
      controller: 'TagCtrl'});
    $routeProvider.when('/problem/:problemId', {
      templateUrl: 'partials/problem.html',
      controller: 'ProblemCtrl'});
    $routeProvider.otherwise({redirectTo: '/'});
  }]);

angular.module('jgApp.controllers', []).
  controller('IndexCtrl', function() {
    // TODO dealing with login
  }).
  controller('TagsCtrl', function($scope, $http) {
    $scope.tags = ["Loading..."];
    $http.post('/rpc/json', {
      method: "Jg.PrimaryTags",
      params: [],
      id: 0,
    }).
      success(function(data, status, headers, config) {
        $scope.tags = angular.fromJson(data).result.tags;
      }).
      error(function(data, status, headers, config) {
        $scope.tags = ["Load failed."];
      });
  }).
  controller('TagCtrl', function($scope, $http, $routeParams) {
    $http.post('/rpc/json', {
      method: "Jg.TagById",
      params: [{"id": $routeParams.tagId}],
      id: 0
    }).
      success(function(data, status, headers, config) {
        var obj = angular.fromJson(data);
        $scope.tag = {
          id: $routeParams.tagId,
          name: obj.result.name,
        };
        $scope.subtags = obj.result.tags;
        $scope.problems = obj.result.problems;
        $scope.tagged = obj.result.tagged;
      }).
      error(function(data, status, headers, config) {
        // TODO on error
      });
  }).
  controller('ProblemCtrl', function($scope, $http, $routeParams) {
    $scope.problem = {id: "np", name: "Name placeholder", brief: "blabla"};
    // TODO
  });

/* angular.module('jgApp.directives', []).
  directive('appVersion', ['version', function(version) {
    return function(scope, elm, attrs) {
      elm.text(version);
    };
  }]);

angular.module('jgApp.filters', []).
  filter('interpolate', ['version', function(v) {
    return function(text) {
      return String(text).replace(/\%VERSION\%/mg, v);
    }
  }]); */

angular.module('jgApp.services', []).
  value('version', '0.1');
