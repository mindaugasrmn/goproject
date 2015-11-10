(function () {
  'use strict';

  angular
    .module('application.routes')
    .config(config);

  config.$inject = ['$routeProvider'];

  function config($routeProvider) {
    $routeProvider.when('/', {
      controller: 'IndexController',
      controllerAs: 'vm',
      templateUrl: '/static/templates/static/dashboard/index.html'
    }).when('/login', {
      controller: 'LoginController',
      controllerAs: 'vm',
      templateUrl: '/static/templates/static/login.html'
    }).when('/register', {
      controller: 'RegisterController',
      controllerAs: 'vm',
      templateUrl: '/static/templates/static/register.html'
    }).when('/me', {
      controller: 'MeController',
      controllerAs: 'vm',
      templateUrl: '/static/templates/static/profile/profile.html'
    }).when('/me/edit', {
      controller: 'MeController',
      controllerAs: 'vm',
      templateUrl: '/static/templates/static/profile/profile-edit.html'
    }).when('/:username', {
      controller: 'ProfileController',
      controllerAs: 'vm',
      templateUrl: '/static/templates/profiles/profile.html'
    }).when('/:username/settings', {
      controller: 'ProfileSettingsController',
      controllerAs: 'vm',
      templateUrl: '/static/templates/profiles/settings.html'
    }).otherwise('/');
  }
})();
