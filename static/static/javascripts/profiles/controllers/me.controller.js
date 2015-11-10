(function () {
  'use strict';

  angular
    .module('application.profiles.controllers')
    .controller('MeController', MeController);

  MeController.$inject = [
    '$location', 'Auth', 'Profile','ToastService'
  ];

  function MeController($location, Auth, Profile, ToastService) {


      var vm = this;
      vm.updateme = updateme;
      activate();

      function activate() {
            if (!Auth.getToken()) {
              $location.path('/login');
            } else {
              Profile.me().then(meSuccessFn);
            }
          
            function meSuccessFn(data, status, headers, config) {
              vm.profile = data.data;
            }  
          }
      function updateme() {
              
              activate();

              function activate() {
                if (!Auth.getToken()) {
                  $location.path('/login');
                } else {
                  Profile.updateme(vm.profile).then(profileSuccessFn, profileErrorFn);
                }
              }
              
              function profileSuccessFn(data, status, headers, config) {
                ToastService.error("SUCCESS")
                $location.path ('/me');
                };


              function profileErrorFn(data, status, headers, config) {
                ToastService.show("SUCCESS")
                $location.path ('/me');

              }
      }

  }

})();
