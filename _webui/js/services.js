var SpykeServicesModule = angular.module('SpykeServices', [])


SpykeServicesModule.factory("socket", function() {
  return {
    ws: null,
    state: "CLOSED",
    joinRoom: function(roomId, successCb) {
      ws = new WebSocket("http://spyke.io/?roomId=" + roomId);
      ws.onopen = function() { state = "CONNECTED"; successCb(); }
      ws.onclose = function() { state = "CLOSED"; }
      ws.onmessage = function() {
        // get that signal where it needs to go.
      }
    },
    sendSignal: function(successCb) {
      successCb();
    },
  };
});

SpykeServicesModule.factory("room", function($resource) {
  return $resource("api/rooms", {}, {
    // ?
  };
});