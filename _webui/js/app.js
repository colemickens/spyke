var spyke = angular.module('spyke', []);

spyke.config(function($routeProvider, $locationProvider) {
    //$locationProvider.html5Mode(true);

    $routeProvider.
        when('/',                {templateUrl: '/partials/about.html',    controller: AboutCtrl    }).
        when('/chat/:roomId',    {templateUrl: '/partials/chat.html',     controller: ChatCtrl     }).
        when('/rooms',           {templateUrl: '/partials/rooms.html',    controller: RoomListCtrl }).

        otherwise({redirectTo: '/'});
});