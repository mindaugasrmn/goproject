(function(){
    "use strict";

    angular.module("application.toast.services").factory('ToastService', function($mdToast){

        var delay = 2000,
            position = 'top right',
            action = 'OK';

        return {
            show: function(content){
                if (!content){
                    return false;
                }

                return $mdToast.show({
                      template: '<md-toast>\
                                <span flex>{{ message | translate}}</span>\
                                <md-button ng-click="closeToast()" class="md-action md-button ng-scope md-default-theme md-ink-ripple">\
                                OK\
                                </md-button>\
                                </md-toast>',
                      
                      hideDelay: delay,
                      position: position,
                      controller: function($scope) {
                          $scope.message = content;
                          $scope["class"] = 'error';
                          $scope.buttonLabel = 'close';
                          return $scope.closeToast = function() {
                            return $mdToast.hide();
                              };
                            }
                });
            },
            error: function(content){
                console.log(content)
                if (!content){
                    return false;
                }

                return $mdToast.show({
                  
                  template: '<md-toast>\
                                <span flex>{{ toast.content | translate }}</span>\
                                <md-button ng-click="toast.resolve()">\
                                OK\
                                </md-button>\
                            </md-toast>',
                  
                  hideDelay: delay,
                  position: position,
                  controller: [function ($scope, $mdToast) {
                    this.content = content;
                }],
                controllerAs: 'toast'
                });
            }
        };
    });
})();