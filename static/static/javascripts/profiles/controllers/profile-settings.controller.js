(function () {
  'use strict';

  angular
    .module('application.profiles.controllers')
    .controller('ProfileSettingsController', ProfileSettingsController);

  ProfileSettingsController.$inject = [
    '$location', '$routeParams', 'Auth', 'Profile','$mdToast','$scope'
  ];


  function ProfileSettingsController($location, $routeParams, Auth, Profile, $mdToast, $scope) {
    $scope.toastPosition = {
      bottom: true,
      top: false,
      left: false,
      right: true
    }; 
    $scope.getToastPosition = function() {
      return Object.keys($scope.toastPosition)
      .filter(function(pos) { return $scope.toastPosition[pos]; })
      .join(' ');
    };
    $scope.user ={
      first_name : "mrr",
      username: "mr"
    }

    var vm = this;

    vm.destroy = destroy;
    vm.update = update;

    activate();


    function activate() {
      var authenticatedAccount = Auth.getToken();
      var username = $routeParams.username.substr(1);
      console.log(authenticatedAccount);
      console.log(username);
     

      // Redirect if not logged in
      
      if (!authenticatedAccount) {
        $location.url('/');
        $mdToast.show($mdToast.simple().content('You are not authorized to view this page!').position($scope.getToastPosition()).hideDelay(3000));
        
      } else {
        // Redirect if logged in, but not the owner of this profile.
        if (authenticatedAccount.username !== username) {
          $location.url('/');
          $mdToast.show($mdToast.simple().content('You are not owner of this profile!').position($scope.getToastPosition()).hideDelay(3000));
        }
      }

      Profile.get(username).then(profileSuccessFn, profileErrorFn);


      function profileSuccessFn(data, status, headers, config) {
        vm.profile = data.data;
      }


      function profileErrorFn(data, status, headers, config) {
        $location.url('/');
        $mdToast.show($mdToast.simple().content('That user does not exist!').position($scope.getToastPosition()).hideDelay(3000));
      }
    }


    function destroy() {
      Profile.destroy(vm.profile.username).then(profileSuccessFn, profileErrorFn);

      function profileSuccessFn(data, status, headers, config) {
        Authentication.unauthenticate();
        $location.url('/');
        $mdToast.show($mdToast.simple().content('Your account has been deleted!').position($scope.getToastPosition()).hideDelay(3000));

      }


      function profileErrorFn(data, status, headers, config) {
        $mdToast.show($mdToast.simple().content(data.error).position($scope.getToastPosition()).hideDelay(3000));

      }
    }


    function update() {
      Profile.update(vm.profile).then(profileSuccessFn, profileErrorFn);


      function profileSuccessFn(data, status, headers, config) {
        $mdToast.show($mdToast.simple().content('Your profile has been updated!').position($scope.getToastPosition()).hideDelay(3000));
 
      }


      function profileErrorFn(data, status, headers, config) {
        $mdToast.show($mdToast.simple().content(data.error).position($scope.getToastPosition()).hideDelay(3000));

      }
    }
  }
})();