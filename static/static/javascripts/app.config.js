(function () {
  'use strict';

  angular
    .module('application.config')
    .config(config);

  config.$inject = ['$httpProvider', '$locationProvider','$mdThemingProvider','$translateProvider'];

  function config($httpProvider, $locationProvider,$mdThemingProvider,$translateProvider,$scope) {
    $httpProvider.interceptors.push('AuthInterceptor');

    $locationProvider.html5Mode(true).hashPrefix('!'); 

    $mdThemingProvider.definePalette('amazingPaletteName', {
    '50': 'ffebee',
    '100': 'ffcdd2',
    '200': 'ef9a9a',
    '300': 'e57373',
    '400': 'ef5350',
    '500': 'f44336',
    '600': 'e53935',
    '700': 'd32f2f',
    '800': 'c62828',
    '900': 'b71c1c',
    'A100': 'ff8a80',
    'A200': 'ff5252',
    'A400': 'ff1744',
    'A700': 'd50000',
    'contrastDefaultColor': 'light',    // whether, by default, text (contrast)
                                        // on this palette should be dark or light
    'contrastDarkColors': ['50', '100', //hues which contrast should be 'dark' by default
     '200', '300', '400', 'A100'],
    'contrastLightColors': undefined    // could also specify this if default was 'dark'
  });

  $mdThemingProvider.theme('default')
    .primaryPalette('amazingPaletteName')
  $mdThemingProvider.theme('altTheme')
    .primaryPalette('purple')
  $mdThemingProvider.alwaysWatchTheme(true);
  $mdThemingProvider.setDefaultTheme('default');
  
  $translateProvider.useSanitizeValueStrategy('escape');
  $translateProvider.preferredLanguage('en');

  $translateProvider.translations('en', {
    MENU: 'Menu',
    FULLSCREEN: 'Toggle Full Screen',
    PROFILE: 'Profile',
    SETTINGS: 'Settings',
    LOGOUT: 'Logout',
    LOGIN: 'Login',
    USERNAME: 'Username',
    PASSWORD: 'Password',
    SUCCESS: 'Action succedded!',

    BUTTON_TEXT_EN: 'English',
    BUTTON_TEXT_LT: 'Lithuanian'
  }).translations('lt', {
    MENU: 'Meniu',
    FULLSCREEN: 'Pilno Ekrano Režimas',
    PROFILE: 'Profilis',
    SETTINGS: 'Nustatymai',
    LOGOUT: 'Atsijungti',
    LOGIN: 'Prisijungti',
    USERNAME: 'Vartotojo vardas',
    PASSWORD: 'Slaptažodis',
    SUCCESS: 'Ivykis pavyko!',

    BUTTON_TEXT_LT: 'Lietuviškai',
    BUTTON_TEXT_EN: 'Angliškai'
  });

  }

})();
