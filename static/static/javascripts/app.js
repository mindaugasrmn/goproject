(function () {
  'use strict';

  angular
    .module('application', [
      'angularScreenfull',
      'pascalprecht.translate',
      'ngSanitize',
      'angular-jwt',
      'ngMaterial',
      'angular-loading-bar',
      'ui.bootstrap',
      'application.config',
      'application.routes',
      'application.auth',
      'application.static',
      'application.profiles',
      'application.toast',
      'application.translate',
    ]).controller('AppCtrl', function ($scope, $timeout, $mdSidenav, $mdUtil, $log, jwtHelper,$window) {
    
    if ($window.localStorage.getItem('token')){
    $scope.profile = jwtHelper.decodeToken($window.localStorage.getItem('token'));
    };
    
    $scope.toggleLeft = buildToggler('left');

    function buildToggler(navID) {
      var debounceFn =  $mdUtil.debounce(function(){
            $mdSidenav(navID)
              .toggle();
          },500);
      return debounceFn;
    }
  })
  .controller('LeftCtrl', function ($scope, $timeout, $mdSidenav, $window) {
    $scope.close = function () {
      $mdSidenav('left').close()
        .then(function () {
          
        });
    };
  });   
  
  angular
    .module('application.config', []);   

  angular
    .module('application.routes', ['ngRoute']);  
})();
