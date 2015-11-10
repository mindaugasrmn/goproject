(function () {
  'use strict';

  angular
    .module('application.translate.controllers')
	.controller('TranslateController', function($translate, $scope) {
	  $scope.changeLanguage = function (langKey) {
	    $translate.use(langKey);
	  };
	});
})();

