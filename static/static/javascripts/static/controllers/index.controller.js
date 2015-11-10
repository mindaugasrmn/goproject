(function () {
  'use strict';

  angular
    .module('application.static.controllers')
    .controller('IndexController', IndexController);

  IndexController.$inject = ['$location', 'Auth', 'Profile', 'ToastService'];

  function IndexController($location, Auth, Profile, ToastService) {
    var vm = this;
  


    activate();

    function activate() {
      if (!Auth.getToken()) {
        $location.path('/login');
      } else {
        Profile.me().then(usersSuccessFn);
      }
    }

    function usersSuccessFn(data, status, headers, config) {
      ToastService.show("SUCCESS")
      vm.profile = data.data;
    }  
  }
})();
