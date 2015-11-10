(function () {
  'use strict';

  angular
    .module('application.profiles.services')
    .factory('Profile', Profile);

  Profile.$inject = ['$http','$window'];

  function Profile($http,$window) {

    var Profile = {
      me: me,
      updateme: updateme,
      destroy: destroy,
      get: get,
      update: update
    };

    return Profile;


    function destroy(profile) {
      return $http.delete('/api/v1/users/' + profile.id + '/');
    }


    function get(username) {
      return $http.get('/api/v1/users/' + username + '/');
    }

    function me() {
      return $http.get('/api/v1/me/');
    }

    function update(profile) {
      return $http.put('/api/v1/users/' + profile.username + '/', profile);

    }

    function updateme(profile) {
      return $http.put('/api/v1/me/', profile);
      
    }
  }
})();