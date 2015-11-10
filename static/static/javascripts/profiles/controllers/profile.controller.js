(function () {
  'use strict';

  angular
    .module('application.profiles.controllers')
    .controller('ProfileController', ProfileController);

  ProfileController.$inject = ['$location', '$routeParams', 'Profile','Auth'];

  function ProfileController($location, $routeParams,  Profile, Auth) {
    var vm = this;

    vm.profile = undefined;
    vm.posts = [];

    activate();


    function activate() {
      //var username = $routeParams.username.substr(1);

      //Profile.get(username).then(profileSuccessFn, profileErrorFn);

      function profileSuccessFn(data, status, headers, config) {
        vm.profile = data.data;

      }


      function profileErrorFn(data, status, headers, config) {
        $location.url('/');
        //Snackbar.error('That user does not exist.');
      }


    }
  }
})();