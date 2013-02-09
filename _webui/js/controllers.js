function AboutCtrl($scope) {
  // about!
}

function RoomListCtrl($scope, $http) {
  $scope.refreshRooms = function() {
    $scope.roomNameFilter = "";

    $http.get("/api/v1/room")
      .success(function(data, status, headers, config) {
        console.log("rooms", data);
        $scope.rooms = data;
      });
  }

  $scope.createRoom = function() {
    $http( { method: "POST", url: "/api/v1/room" } )
      .success(function(data, status, headers, config) {
        console.log("created room!"); // TODO: shouldn't this just return a link? (REST)
        $scope.refreshRooms();
      });

    // create a new $http $httpF
    // that automatically does form submission
  }

  $scope.refreshRooms();
}

function ChatCtrl($scope, $http, $routeParams) {
  $scope.status = "Please allow access to your webcam!";

  $http.get("/api/room/" + $routeParams.roomId)
    .success(function(data, status, headers, config) {
      $scope.roomDetails = data;
    });

  gum(
    function(source) {
      $scope.localSource = source;
      $scope.status = "Local rendering successful!";
      $scope.$apply();
    },
    function() {
      $scope.status = "You must grant access to your webcam!";
    }
  );

  socket.joinRoom($scope.roomId, function() {
    console.log("connected!");
  });
}

function gum(success, error) {
  navigator.getUserMedia || (navigator.getUserMedia = navigator.mozGetUserMedia || navigator.webkitGetUserMedia || navigator.msGetUserMedia);
  if (navigator.getUserMedia) {
    navigator.getUserMedia(
      {video: true, audio: true, toString : function() {return "video,audio";} },
      function(stream) {
        if(window.webkitURL) {
          success(window.webkitURL.createObjectURL(stream));
        } else {
          success(stream);
        }
      },
      error);
  } else {
    $scope.status = "Your browser does not support the getUserMedia() API.";
  }
}